package handlers

import (
	"io"
	"log"

	"AppTransportistasBack/config"
	"AppTransportistasBack/despachopb"
)

type DespachoServer struct {
	despachopb.UnimplementedDespachoServiceServer
}

// EnviarEntregas recibe guías entregadas desde la app y las guarda en MySQL
func (s *DespachoServer) EnviarEntregas(stream despachopb.DespachoService_EnviarEntregasServer) error {
	var total int32

	for {
		guia, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&despachopb.EntregaResponse{
				Mensaje:          "Entregas registradas exitosamente",
				TotalRegistradas: total,
			})
		}
		if err != nil {
			return err
		}

		// Insertar guía
		res, err := config.DB.Exec(`
            INSERT INTO guia (numero, fecha, codigo_cliente, nombre_cliente, nro_comprobante,
                               importe_x_cobrar, monto_cobrado, entregada)
            VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			guia.Numero,
			guia.Fecha,
			guia.CodigoCliente,
			guia.NombreCliente,
			guia.NroComprobante,
			guia.ImporteXCobrar,
			guia.MontoCobrado,
			guia.Entregada,
		)
		if err != nil {
			log.Printf("Error insertando guía: %v", err)
			continue
		}

		guiaID, _ := res.LastInsertId()

		// Insertar productos
		for _, p := range guia.Productos {
			_, err := config.DB.Exec(`
                INSERT INTO producto (id_guia, nombre, cantidad)
                VALUES (?, ?, ?)`,
				guiaID, p.Nombre, p.Cantidad,
			)
			if err != nil {
				log.Printf("Error insertando producto: %v", err)
			}
		}

		total++
	}
}

// ObtenerDespachos devuelve guías pendientes desde el backend a la app móvil
func (s *DespachoServer) ObtenerDespachos(req *despachopb.DespachoRequest, stream despachopb.DespachoService_ObtenerDespachosServer) error {
	//log.Println("📡 Recibida solicitud de obtener guías entregadas = 0")

	rows, err := config.DB.Query(`
        SELECT id, numero, fecha, codigo_cliente, nombre_cliente,
               nro_comprobante, importe_x_cobrar, monto_cobrado, entregada
        FROM guia
        WHERE entregada = 0
    `)
	if err != nil {
		log.Printf("❌ Error ejecutando query de guías: %v", err)
		return err
	}
	defer rows.Close()

	contador := 0

	for rows.Next() {
		var (
			id                                         int
			numero, fecha, codigo, nombre, comprobante string
			importe, cobrado                           float64
			entregada                                  bool
		)

		err := rows.Scan(&id, &numero, &fecha, &codigo, &nombre, &comprobante, &importe, &cobrado, &entregada)
		if err != nil {
			log.Printf("❌ Error escaneando guía: %v", err)
			continue
		}

		//log.Printf("✅ Guía encontrada: %s | Cliente: %s | Importe: %.2f", numero, nombre, importe)

		// Consultar productos
		productos := []*despachopb.Producto{}
		prodRows, err := config.DB.Query(`
            SELECT nombre, cantidad
            FROM producto
            WHERE id_guia = ?
        `, id)
		if err != nil {
			log.Printf("⚠️ Error obteniendo productos para guía %s: %v", numero, err)
		} else {
			for prodRows.Next() {
				var nombre string
				var cantidad int32
				if err := prodRows.Scan(&nombre, &cantidad); err != nil {
					log.Printf("❌ Error escaneando producto: %v", err)
					continue
				}
				productos = append(productos, &despachopb.Producto{
					Nombre:   nombre,
					Cantidad: cantidad,
				})
			}
			prodRows.Close()
		}

		guia := &despachopb.Guia{
			Numero:         numero,
			Fecha:          fecha,
			CodigoCliente:  codigo,
			NombreCliente:  nombre,
			NroComprobante: comprobante,
			ImporteXCobrar: importe,
			MontoCobrado:   cobrado,
			Entregada:      entregada,
			Productos:      productos,
		}

		log.Printf("📤 Enviando guía al cliente: %s con %d productos", numero, len(productos))

		if err := stream.Send(guia); err != nil {
			log.Printf("❌ Error enviando guía %s: %v", numero, err)
			return err
		}

		contador++
	}

	//log.Printf("Total de guías enviadas: %d", contador)
	return nil
}
