package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"AppTransportistasBack/apptransportistaspb"
	"AppTransportistasBack/config"
)

type AppTransportistasServer struct {
	apptransportistaspb.UnimplementedAppTransportistasServiceServer
}

// ObtenerDespachos envía las guías no entregadas
func (s *AppTransportistasServer) ObtenerDespachos(req *apptransportistaspb.DespachoRequest, stream apptransportistaspb.AppTransportistasService_ObtenerDespachosServer) error {
	log.Println("📡 Solicitud de guías pendientes")
	rows, err := config.DB.Query(`
		SELECT id, numero, fecha, codigo_cliente, nombre_cliente,
		       nro_comprobante, importe_x_cobrar, monto_cobrado, entregada
		FROM guia WHERE entregada = 0`)
	if err != nil {
		log.Printf("❌ Error consultando guías: %v", err)
		return err
	}
	defer rows.Close()

	var total int
	for rows.Next() {
		var (
			id                                     int
			numero, fecha, codCli, nomCli, comprob string
			impCobrar, montoCobrado                float64
			entregada                              bool
		)

		if err := rows.Scan(&id, &numero, &fecha, &codCli, &nomCli, &comprob, &impCobrar, &montoCobrado, &entregada); err != nil {
			log.Printf("⚠️ Error leyendo guía: %v", err)
			continue
		}

		prods := []*apptransportistaspb.Producto{}
		prodRows, err := config.DB.Query(`SELECT nombre, cantidad FROM producto WHERE id_guia = ?`, id)
		if err == nil {
			for prodRows.Next() {
				var n string
				var c int32
				if err := prodRows.Scan(&n, &c); err == nil {
					prods = append(prods, &apptransportistaspb.Producto{Nombre: n, Cantidad: c})
				}
			}
			prodRows.Close()
		}

		if err := stream.Send(&apptransportistaspb.Guia{
			Numero:         numero,
			Fecha:          fecha,
			CodigoCliente:  codCli,
			NombreCliente:  nomCli,
			NroComprobante: comprob,
			ImporteXCobrar: impCobrar,
			MontoCobrado:   montoCobrado,
			Entregada:      entregada,
			Productos:      prods,
		}); err != nil {
			log.Printf("❌ Error enviando guía %s: %v", numero, err)
			return err
		}
		total++
	}

	log.Printf("✅ Total guías enviadas: %d", total)
	return nil
}

// EnviarEntregas recibe guías desde la app y las guarda en MySQL
func (s *AppTransportistasServer) EnviarEntregas(stream apptransportistaspb.AppTransportistasService_EnviarEntregasServer) error {
	log.Println("🚀 Iniciando recepción de entregas")
	var total int32

	for {
		guia, err := stream.Recv()
		if err == io.EOF {
			log.Printf("✅ Entregas completadas. Total registradas o actualizadas: %d", total)
			return stream.SendAndClose(&apptransportistaspb.EntregaResponse{
				Mensaje:          "Entregas procesadas exitosamente",
				TotalRegistradas: total,
			})
		}
		if err != nil {
			log.Printf("❌ Error recibiendo guía: %v", err)
			return err
		}

		log.Printf("📥 Guía %s - Cliente: %s", guia.Numero, guia.NombreCliente)

		// Extraer solo la fecha YYYY-MM-DD
		var fechaMysql string
		if guia.Fecha != "" {
			t, err := time.Parse(time.RFC3339, guia.Fecha) // la app envía ISO 8601
			if err != nil {
				log.Printf("⚠️ Fecha inválida '%s' para guía %s: %v", guia.Fecha, guia.Numero, err)
				fechaMysql = "0001-01-01" // fallback mínimo
			} else {
				fechaMysql = t.Format("2006-01-02") // solo fecha
			}
		} else {
			fechaMysql = "0001-01-01" // fallback si viene vacía
		}

		// Verificar si ya existe la guía
		var idGuia int64
		var entregada bool
		err = config.DB.QueryRow(`SELECT id, entregada FROM guia WHERE numero = ?`, guia.Numero).Scan(&idGuia, &entregada)

		if err == sql.ErrNoRows {
			// Insertar
			res, err := config.DB.Exec(`
				INSERT INTO guia (numero, fecha, codigo_cliente, nombre_cliente, nro_comprobante,
								  importe_x_cobrar, monto_cobrado, entregada)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
				guia.Numero, fechaMysql, guia.CodigoCliente, guia.NombreCliente,
				guia.NroComprobante, guia.ImporteXCobrar, guia.MontoCobrado, guia.Entregada)

			if err != nil {
				log.Printf("❌ Error insertando guía %s: %v", guia.Numero, err)
				continue
			}

			idGuia, _ = res.LastInsertId()
			log.Printf("🆕 Guía %s insertada", guia.Numero)

		} else if err == nil {
			if entregada {
				log.Printf("⚠️ Guía %s ya fue entregada. Se ignora.", guia.Numero)
				continue
			}

			// Actualizar
			_, err = config.DB.Exec(`
				UPDATE guia SET fecha = ?, codigo_cliente = ?, nombre_cliente = ?, nro_comprobante = ?,
								importe_x_cobrar = ?, monto_cobrado = ?, entregada = ?
				WHERE id = ?`,
				fechaMysql, guia.CodigoCliente, guia.NombreCliente, guia.NroComprobante,
				guia.ImporteXCobrar, guia.MontoCobrado, guia.Entregada, idGuia)

			if err != nil {
				log.Printf("❌ Error actualizando guía %s: %v", guia.Numero, err)
				continue
			}
			log.Printf("🔄 Guía %s actualizada", guia.Numero)

			// Borrar productos anteriores
			_, _ = config.DB.Exec("DELETE FROM producto WHERE id_guia = ?", idGuia)
		} else {
			log.Printf("❌ Error consultando guía %s: %v", guia.Numero, err)
			continue
		}

		// Insertar productos
		for _, p := range guia.Productos {
			_, err := config.DB.Exec(`INSERT INTO producto (id_guia, nombre, cantidad) VALUES (?, ?, ?)`,
				idGuia, p.Nombre, p.Cantidad)

			if err != nil {
				log.Printf("⚠️ Producto %s - Error: %v", p.Nombre, err)
			}
		}

		total++
	}
}

// EnviarPruebasEntrega recibe pruebas desde la app y las guarda en MySQL
func (s *AppTransportistasServer) EnviarPruebasEntrega(stream apptransportistaspb.AppTransportistasService_EnviarPruebasEntregaServer) error {
	log.Println("📩 Recibiendo pruebas de entrega...")

	baseFirmaPath := `C:\Users\PracSistemas\Documents\Visual Studio Code\AppTransportistasBack\img\pruebas\firmas`
	baseFotoPath := `C:\Users\PracSistemas\Documents\Visual Studio Code\AppTransportistasBack\img\pruebas\fotos`
	_ = os.MkdirAll(baseFirmaPath, os.ModePerm)
	_ = os.MkdirAll(baseFotoPath, os.ModePerm)

	var total int32

	for {
		prueba, err := stream.Recv()
		if err == io.EOF {
			log.Printf("✅ Pruebas de entrega recibidas: %d", total)
			return stream.SendAndClose(&apptransportistaspb.PruebaEntregaResponse{
				Mensaje:          "Pruebas registradas correctamente",
				TotalRegistradas: total,
			})
		}
		if err != nil {
			log.Printf("❌ Error recibiendo prueba: %v", err)
			return err
		}

		log.Printf("🖊️ Prueba para guía %s", prueba.NumeroGuia)

		var guiaID int
		err = config.DB.QueryRow("SELECT id FROM guia WHERE numero = ?", prueba.NumeroGuia).Scan(&guiaID)
		if err != nil {
			log.Printf("❌ Guía %s no encontrada: %v", prueba.NumeroGuia, err)
			continue
		}

		fecha, err := time.Parse("2006-01-02 15:04:05", prueba.FechaRegistro)
		if err != nil {
			log.Printf("⚠️ Fecha inválida '%s' para guía %s", prueba.FechaRegistro, prueba.NumeroGuia)
			continue
		}

		// Guardar firma como imagen
		firmaPath := ""
		if len(prueba.Firma) > 0 {
			firmaName := fmt.Sprintf("firma_guia_%s.jpg", prueba.NumeroGuia)
			firmaPath = filepath.Join(baseFirmaPath, firmaName)
			if err := os.WriteFile(firmaPath, prueba.Firma, 0644); err != nil {
				log.Printf("⚠️ No se pudo guardar firma para guía %s: %v", prueba.NumeroGuia, err)
				continue
			}
		}

		// Guardar imagen de comprobante
		imgPath := ""
		if len(prueba.Imagen) > 0 {
			imgName := fmt.Sprintf("img_guia_%s.jpg", prueba.NumeroGuia)
			imgPath = filepath.Join(baseFotoPath, imgName)
			if err := os.WriteFile(imgPath, prueba.Imagen, 0644); err != nil {
				log.Printf("⚠️ No se pudo guardar imagen para guía %s: %v", prueba.NumeroGuia, err)
				continue
			}
		}

		var existe int
		err = config.DB.QueryRow(`SELECT COUNT(*) FROM prueba_entrega WHERE guia_id = ?`, guiaID).Scan(&existe)
		if err != nil {
			log.Printf("❌ Error verificando existencia de prueba para guía %s: %v", prueba.NumeroGuia, err)
			continue
		}
		if existe > 0 {
			log.Printf("⏩ Prueba ya registrada para guía %s, omitiendo", prueba.NumeroGuia)
			continue
		}

		_, err = config.DB.Exec(`
			INSERT INTO prueba_entrega (guia_id, fecha_registro, firma_path, imagen_path)
			VALUES (?, ?, ?, ?)`,
			guiaID, fecha, firmaPath, imgPath)
		if err != nil {
			log.Printf("❌ Error guardando prueba para guía %s: %v", prueba.NumeroGuia, err)
			continue
		}

		total++
	}
}

// LoginAdmin verifica las credenciales del administrador
func (s *AppTransportistasServer) LoginAdmin(ctx context.Context, req *apptransportistaspb.LoginRequest) (*apptransportistaspb.LoginResponse, error) {
	usuario := req.GetUsuario()
	password := req.GetPassword()

	var count int
	err := config.DB.QueryRow(
		`SELECT COUNT(*) FROM admin WHERE usuario = ? AND password = ?`,
		usuario, password,
	).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("❌ Error verificando login: %v", err)
		return &apptransportistaspb.LoginResponse{
			Success: false,
			Mensaje: "Error interno en el servidor",
		}, nil
	}

	if count > 0 {
		return &apptransportistaspb.LoginResponse{
			Success: true,
			Mensaje: "Acceso permitido",
		}, nil
	}

	return &apptransportistaspb.LoginResponse{
		Success: false,
		Mensaje: "Usuario o contraseña incorrectos",
	}, nil
}

// EnviarTracking recibe ubicaciones desde la app y las guarda en MySQL
func (s *AppTransportistasServer) EnviarTracking(stream apptransportistaspb.AppTransportistasService_EnviarTrackingServer) error {
	log.Println("📡 Recibiendo tracking de dispositivos...")

	var total int32

	for {
		ubicacion, err := stream.Recv()
		if err == io.EOF {
			log.Printf("✅ Tracking recibido. Total registros: %d", total)
			return stream.SendAndClose(&apptransportistaspb.TrackingResponse{
				Mensaje:          "Tracking registrado correctamente",
				TotalRegistrados: total,
			})
		}
		if err != nil {
			log.Printf("❌ Error recibiendo ubicación: %v", err)
			return err
		}

		log.Printf("📍 Dispositivo %s - Lat: %f, Lon: %f, Fecha: %s",
			ubicacion.DeviceId, ubicacion.Latitud, ubicacion.Longitud, ubicacion.FechaHora)

		// Insertar en la tabla tracking
		_, err = config.DB.Exec(`
            INSERT INTO tracking (device_id, latitud, longitud, fecha_hora)
            VALUES (?, ?, ?, ?)`,
			ubicacion.DeviceId, ubicacion.Latitud, ubicacion.Longitud, ubicacion.FechaHora,
		)
		if err != nil {
			log.Printf("❌ Error guardando ubicación para dispositivo %s: %v", ubicacion.DeviceId, err)
			continue
		}

		total++
	}
}
