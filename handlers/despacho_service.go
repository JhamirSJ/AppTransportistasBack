package handlers

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"AppTransportistasBack/apptransportistaspb"
	"AppTransportistasBack/config"
)

type AppTransportistasServer struct {
	apptransportistaspb.UnimplementedAppTransportistasServiceServer
}

// EnviarEntregas recibe guías desde la app y las guarda en MySQL
func (s *AppTransportistasServer) EnviarEntregas(stream apptransportistaspb.AppTransportistasService_EnviarEntregasServer) error {
	log.Println("🚀 Iniciando recepción de entregas")
	var total int32

	for {
		guia, err := stream.Recv()
		if err == io.EOF {
			log.Printf("✅ Entregas completadas. Total registradas: %d", total)
			return stream.SendAndClose(&apptransportistaspb.EntregaResponse{
				Mensaje:          "Entregas registradas exitosamente",
				TotalRegistradas: total,
			})
		}
		if err != nil {
			log.Printf("❌ Error recibiendo guía: %v", err)
			return err
		}

		log.Printf("📥 Guía %s - Cliente: %s", guia.Numero, guia.NombreCliente)

		res, err := config.DB.Exec(`
			INSERT INTO guia (numero, fecha, codigo_cliente, nombre_cliente, nro_comprobante,
							  importe_x_cobrar, monto_cobrado, entregada)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			guia.Numero, guia.Fecha, guia.CodigoCliente, guia.NombreCliente,
			guia.NroComprobante, guia.ImporteXCobrar, guia.MontoCobrado, guia.Entregada)
		if err != nil {
			log.Printf("❌ Error insertando guía %s: %v", guia.Numero, err)
			continue
		}

		guiaID, _ := res.LastInsertId()
		for _, p := range guia.Productos {
			if _, err := config.DB.Exec(`INSERT INTO producto (id_guia, nombre, cantidad) VALUES (?, ?, ?)`,
				guiaID, p.Nombre, p.Cantidad); err != nil {
				log.Printf("⚠️ Producto %s - Error: %v", p.Nombre, err)
			}
		}

		total++
	}
}

// ObtenerDespachos envía las guías no entregadas
func (s *AppTransportistasServer) ObtenerDespachos(req *apptransportistaspb.AppTransportistasRequest, stream apptransportistaspb.AppTransportistasService_ObtenerDespachosServer) error {
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

// EnviarDepositos guarda los depósitos y la imagen en disco
func (s *AppTransportistasServer) EnviarDepositos(stream apptransportistaspb.AppTransportistasService_EnviarDepositosServer) error {
	log.Println("🚀 Iniciando recepción de depósitos")
	basePath := `C:\Users\PracSistemas\Documents\Visual Studio Code\AppTransportistasBack\vouchers`
	_ = os.MkdirAll(basePath, os.ModePerm)

	var total int32

	for {
		depo, err := stream.Recv()
		if err == io.EOF {
			log.Printf("✅ Depósitos completados. Total registrados: %d", total)
			return stream.SendAndClose(&apptransportistaspb.DepositoResponse{
				Mensaje:          "Depósitos registrados exitosamente",
				TotalRegistrados: total,
			})
		}
		if err != nil {
			log.Printf("❌ Error recibiendo depósito: %v", err)
			return err
		}

		log.Printf("📥 Depósito recibido: %s - S/ %.2f", depo.NroOperacion, depo.Monto)

		filename := fmt.Sprintf("%s.jpg", depo.NroOperacion)
		fullPath := filepath.Join(basePath, filename)

		if err := os.WriteFile(fullPath, depo.Comprobante, 0644); err != nil {
			log.Printf("⚠️ No se pudo guardar imagen %s: %v", filename, err)
			continue
		}

		if _, err := config.DB.Exec(`
			INSERT INTO deposito (nro_operacion, fecha, id_banco, monto, comprobante_path, sincronizado)
			VALUES (?, ?, ?, ?, ?, 1)`,
			depo.NroOperacion, depo.Fecha, depo.IdBanco, depo.Monto, fullPath); err != nil {
			log.Printf("❌ Error insertando depósito %s: %v", depo.NroOperacion, err)
			continue
		}

		total++
	}
}
