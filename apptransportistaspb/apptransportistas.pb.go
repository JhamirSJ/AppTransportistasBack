// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: apptransportistas.proto

package apptransportistaspb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DespachoRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	TransportistaId int32                  `protobuf:"varint,1,opt,name=transportista_id,json=transportistaId,proto3" json:"transportista_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *DespachoRequest) Reset() {
	*x = DespachoRequest{}
	mi := &file_apptransportistas_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DespachoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DespachoRequest) ProtoMessage() {}

func (x *DespachoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apptransportistas_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DespachoRequest.ProtoReflect.Descriptor instead.
func (*DespachoRequest) Descriptor() ([]byte, []int) {
	return file_apptransportistas_proto_rawDescGZIP(), []int{0}
}

func (x *DespachoRequest) GetTransportistaId() int32 {
	if x != nil {
		return x.TransportistaId
	}
	return 0
}

type EntregaResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Mensaje          string                 `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
	TotalRegistradas int32                  `protobuf:"varint,2,opt,name=total_registradas,json=totalRegistradas,proto3" json:"total_registradas,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *EntregaResponse) Reset() {
	*x = EntregaResponse{}
	mi := &file_apptransportistas_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntregaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntregaResponse) ProtoMessage() {}

func (x *EntregaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apptransportistas_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntregaResponse.ProtoReflect.Descriptor instead.
func (*EntregaResponse) Descriptor() ([]byte, []int) {
	return file_apptransportistas_proto_rawDescGZIP(), []int{1}
}

func (x *EntregaResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

func (x *EntregaResponse) GetTotalRegistradas() int32 {
	if x != nil {
		return x.TotalRegistradas
	}
	return 0
}

type DepositoResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Mensaje          string                 `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
	TotalRegistrados int32                  `protobuf:"varint,2,opt,name=total_registrados,json=totalRegistrados,proto3" json:"total_registrados,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *DepositoResponse) Reset() {
	*x = DepositoResponse{}
	mi := &file_apptransportistas_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DepositoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositoResponse) ProtoMessage() {}

func (x *DepositoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apptransportistas_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositoResponse.ProtoReflect.Descriptor instead.
func (*DepositoResponse) Descriptor() ([]byte, []int) {
	return file_apptransportistas_proto_rawDescGZIP(), []int{2}
}

func (x *DepositoResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

func (x *DepositoResponse) GetTotalRegistrados() int32 {
	if x != nil {
		return x.TotalRegistrados
	}
	return 0
}

type PruebaEntregaResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Mensaje          string                 `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
	TotalRegistradas int32                  `protobuf:"varint,2,opt,name=total_registradas,json=totalRegistradas,proto3" json:"total_registradas,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *PruebaEntregaResponse) Reset() {
	*x = PruebaEntregaResponse{}
	mi := &file_apptransportistas_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PruebaEntregaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PruebaEntregaResponse) ProtoMessage() {}

func (x *PruebaEntregaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apptransportistas_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PruebaEntregaResponse.ProtoReflect.Descriptor instead.
func (*PruebaEntregaResponse) Descriptor() ([]byte, []int) {
	return file_apptransportistas_proto_rawDescGZIP(), []int{3}
}

func (x *PruebaEntregaResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

func (x *PruebaEntregaResponse) GetTotalRegistradas() int32 {
	if x != nil {
		return x.TotalRegistradas
	}
	return 0
}

type Guia struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Numero         string                 `protobuf:"bytes,1,opt,name=numero,proto3" json:"numero,omitempty"`
	Fecha          string                 `protobuf:"bytes,2,opt,name=fecha,proto3" json:"fecha,omitempty"` // formato: "2025-07-01"
	CodigoCliente  string                 `protobuf:"bytes,3,opt,name=codigo_cliente,json=codigoCliente,proto3" json:"codigo_cliente,omitempty"`
	NombreCliente  string                 `protobuf:"bytes,4,opt,name=nombre_cliente,json=nombreCliente,proto3" json:"nombre_cliente,omitempty"`
	NroComprobante string                 `protobuf:"bytes,5,opt,name=nro_comprobante,json=nroComprobante,proto3" json:"nro_comprobante,omitempty"`
	ImporteXCobrar float64                `protobuf:"fixed64,6,opt,name=importe_x_cobrar,json=importeXCobrar,proto3" json:"importe_x_cobrar,omitempty"`
	MontoCobrado   float64                `protobuf:"fixed64,7,opt,name=monto_cobrado,json=montoCobrado,proto3" json:"monto_cobrado,omitempty"`
	Entregada      bool                   `protobuf:"varint,8,opt,name=entregada,proto3" json:"entregada,omitempty"`
	Productos      []*Producto            `protobuf:"bytes,9,rep,name=productos,proto3" json:"productos,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Guia) Reset() {
	*x = Guia{}
	mi := &file_apptransportistas_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Guia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Guia) ProtoMessage() {}

func (x *Guia) ProtoReflect() protoreflect.Message {
	mi := &file_apptransportistas_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Guia.ProtoReflect.Descriptor instead.
func (*Guia) Descriptor() ([]byte, []int) {
	return file_apptransportistas_proto_rawDescGZIP(), []int{4}
}

func (x *Guia) GetNumero() string {
	if x != nil {
		return x.Numero
	}
	return ""
}

func (x *Guia) GetFecha() string {
	if x != nil {
		return x.Fecha
	}
	return ""
}

func (x *Guia) GetCodigoCliente() string {
	if x != nil {
		return x.CodigoCliente
	}
	return ""
}

func (x *Guia) GetNombreCliente() string {
	if x != nil {
		return x.NombreCliente
	}
	return ""
}

func (x *Guia) GetNroComprobante() string {
	if x != nil {
		return x.NroComprobante
	}
	return ""
}

func (x *Guia) GetImporteXCobrar() float64 {
	if x != nil {
		return x.ImporteXCobrar
	}
	return 0
}

func (x *Guia) GetMontoCobrado() float64 {
	if x != nil {
		return x.MontoCobrado
	}
	return 0
}

func (x *Guia) GetEntregada() bool {
	if x != nil {
		return x.Entregada
	}
	return false
}

func (x *Guia) GetProductos() []*Producto {
	if x != nil {
		return x.Productos
	}
	return nil
}

type Producto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nombre        string                 `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Cantidad      int32                  `protobuf:"varint,2,opt,name=cantidad,proto3" json:"cantidad,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Producto) Reset() {
	*x = Producto{}
	mi := &file_apptransportistas_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Producto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Producto) ProtoMessage() {}

func (x *Producto) ProtoReflect() protoreflect.Message {
	mi := &file_apptransportistas_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Producto.ProtoReflect.Descriptor instead.
func (*Producto) Descriptor() ([]byte, []int) {
	return file_apptransportistas_proto_rawDescGZIP(), []int{5}
}

func (x *Producto) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Producto) GetCantidad() int32 {
	if x != nil {
		return x.Cantidad
	}
	return 0
}

type Deposito struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NroOperacion  string                 `protobuf:"bytes,1,opt,name=nro_operacion,json=nroOperacion,proto3" json:"nro_operacion,omitempty"`
	Fecha         string                 `protobuf:"bytes,2,opt,name=fecha,proto3" json:"fecha,omitempty"`
	IdBanco       int32                  `protobuf:"varint,3,opt,name=id_banco,json=idBanco,proto3" json:"id_banco,omitempty"`
	Monto         float64                `protobuf:"fixed64,4,opt,name=monto,proto3" json:"monto,omitempty"`
	Comprobante   []byte                 `protobuf:"bytes,5,opt,name=comprobante,proto3" json:"comprobante,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Deposito) Reset() {
	*x = Deposito{}
	mi := &file_apptransportistas_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Deposito) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deposito) ProtoMessage() {}

func (x *Deposito) ProtoReflect() protoreflect.Message {
	mi := &file_apptransportistas_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deposito.ProtoReflect.Descriptor instead.
func (*Deposito) Descriptor() ([]byte, []int) {
	return file_apptransportistas_proto_rawDescGZIP(), []int{6}
}

func (x *Deposito) GetNroOperacion() string {
	if x != nil {
		return x.NroOperacion
	}
	return ""
}

func (x *Deposito) GetFecha() string {
	if x != nil {
		return x.Fecha
	}
	return ""
}

func (x *Deposito) GetIdBanco() int32 {
	if x != nil {
		return x.IdBanco
	}
	return 0
}

func (x *Deposito) GetMonto() float64 {
	if x != nil {
		return x.Monto
	}
	return 0
}

func (x *Deposito) GetComprobante() []byte {
	if x != nil {
		return x.Comprobante
	}
	return nil
}

type PruebaEntrega struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NumeroGuia    string                 `protobuf:"bytes,1,opt,name=numero_guia,json=numeroGuia,proto3" json:"numero_guia,omitempty"`
	FechaRegistro string                 `protobuf:"bytes,2,opt,name=fecha_registro,json=fechaRegistro,proto3" json:"fecha_registro,omitempty"` // formato: "2025-07-11"
	Firma         []byte                 `protobuf:"bytes,3,opt,name=firma,proto3" json:"firma,omitempty"`                                      // Firma como imagen (PNG comprimido)
	Imagen        []byte                 `protobuf:"bytes,4,opt,name=imagen,proto3" json:"imagen,omitempty"`                                    // Imagen como comprobante (JPG o PNG comprimido)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PruebaEntrega) Reset() {
	*x = PruebaEntrega{}
	mi := &file_apptransportistas_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PruebaEntrega) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PruebaEntrega) ProtoMessage() {}

func (x *PruebaEntrega) ProtoReflect() protoreflect.Message {
	mi := &file_apptransportistas_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PruebaEntrega.ProtoReflect.Descriptor instead.
func (*PruebaEntrega) Descriptor() ([]byte, []int) {
	return file_apptransportistas_proto_rawDescGZIP(), []int{7}
}

func (x *PruebaEntrega) GetNumeroGuia() string {
	if x != nil {
		return x.NumeroGuia
	}
	return ""
}

func (x *PruebaEntrega) GetFechaRegistro() string {
	if x != nil {
		return x.FechaRegistro
	}
	return ""
}

func (x *PruebaEntrega) GetFirma() []byte {
	if x != nil {
		return x.Firma
	}
	return nil
}

func (x *PruebaEntrega) GetImagen() []byte {
	if x != nil {
		return x.Imagen
	}
	return nil
}

var File_apptransportistas_proto protoreflect.FileDescriptor

const file_apptransportistas_proto_rawDesc = "" +
	"\n" +
	"\x17apptransportistas.proto\x12\x11apptransportistas\"<\n" +
	"\x0fDespachoRequest\x12)\n" +
	"\x10transportista_id\x18\x01 \x01(\x05R\x0ftransportistaId\"X\n" +
	"\x0fEntregaResponse\x12\x18\n" +
	"\amensaje\x18\x01 \x01(\tR\amensaje\x12+\n" +
	"\x11total_registradas\x18\x02 \x01(\x05R\x10totalRegistradas\"Y\n" +
	"\x10DepositoResponse\x12\x18\n" +
	"\amensaje\x18\x01 \x01(\tR\amensaje\x12+\n" +
	"\x11total_registrados\x18\x02 \x01(\x05R\x10totalRegistrados\"^\n" +
	"\x15PruebaEntregaResponse\x12\x18\n" +
	"\amensaje\x18\x01 \x01(\tR\amensaje\x12+\n" +
	"\x11total_registradas\x18\x02 \x01(\x05R\x10totalRegistradas\"\xd3\x02\n" +
	"\x04Guia\x12\x16\n" +
	"\x06numero\x18\x01 \x01(\tR\x06numero\x12\x14\n" +
	"\x05fecha\x18\x02 \x01(\tR\x05fecha\x12%\n" +
	"\x0ecodigo_cliente\x18\x03 \x01(\tR\rcodigoCliente\x12%\n" +
	"\x0enombre_cliente\x18\x04 \x01(\tR\rnombreCliente\x12'\n" +
	"\x0fnro_comprobante\x18\x05 \x01(\tR\x0enroComprobante\x12(\n" +
	"\x10importe_x_cobrar\x18\x06 \x01(\x01R\x0eimporteXCobrar\x12#\n" +
	"\rmonto_cobrado\x18\a \x01(\x01R\fmontoCobrado\x12\x1c\n" +
	"\tentregada\x18\b \x01(\bR\tentregada\x129\n" +
	"\tproductos\x18\t \x03(\v2\x1b.apptransportistas.ProductoR\tproductos\">\n" +
	"\bProducto\x12\x16\n" +
	"\x06nombre\x18\x01 \x01(\tR\x06nombre\x12\x1a\n" +
	"\bcantidad\x18\x02 \x01(\x05R\bcantidad\"\x98\x01\n" +
	"\bDeposito\x12#\n" +
	"\rnro_operacion\x18\x01 \x01(\tR\fnroOperacion\x12\x14\n" +
	"\x05fecha\x18\x02 \x01(\tR\x05fecha\x12\x19\n" +
	"\bid_banco\x18\x03 \x01(\x05R\aidBanco\x12\x14\n" +
	"\x05monto\x18\x04 \x01(\x01R\x05monto\x12 \n" +
	"\vcomprobante\x18\x05 \x01(\fR\vcomprobante\"\x85\x01\n" +
	"\rPruebaEntrega\x12\x1f\n" +
	"\vnumero_guia\x18\x01 \x01(\tR\n" +
	"numeroGuia\x12%\n" +
	"\x0efecha_registro\x18\x02 \x01(\tR\rfechaRegistro\x12\x14\n" +
	"\x05firma\x18\x03 \x01(\fR\x05firma\x12\x16\n" +
	"\x06imagen\x18\x04 \x01(\fR\x06imagen2\xfb\x02\n" +
	"\x18AppTransportistasService\x12O\n" +
	"\x0eEnviarEntregas\x12\x17.apptransportistas.Guia\x1a\".apptransportistas.EntregaResponse(\x01\x12Q\n" +
	"\x10ObtenerDespachos\x12\".apptransportistas.DespachoRequest\x1a\x17.apptransportistas.Guia0\x01\x12U\n" +
	"\x0fEnviarDepositos\x12\x1b.apptransportistas.Deposito\x1a#.apptransportistas.DepositoResponse(\x01\x12d\n" +
	"\x14EnviarPruebasEntrega\x12 .apptransportistas.PruebaEntrega\x1a(.apptransportistas.PruebaEntregaResponse(\x01B\x18Z\x16../apptransportistaspbb\x06proto3"

var (
	file_apptransportistas_proto_rawDescOnce sync.Once
	file_apptransportistas_proto_rawDescData []byte
)

func file_apptransportistas_proto_rawDescGZIP() []byte {
	file_apptransportistas_proto_rawDescOnce.Do(func() {
		file_apptransportistas_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_apptransportistas_proto_rawDesc), len(file_apptransportistas_proto_rawDesc)))
	})
	return file_apptransportistas_proto_rawDescData
}

var file_apptransportistas_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_apptransportistas_proto_goTypes = []any{
	(*DespachoRequest)(nil),       // 0: apptransportistas.DespachoRequest
	(*EntregaResponse)(nil),       // 1: apptransportistas.EntregaResponse
	(*DepositoResponse)(nil),      // 2: apptransportistas.DepositoResponse
	(*PruebaEntregaResponse)(nil), // 3: apptransportistas.PruebaEntregaResponse
	(*Guia)(nil),                  // 4: apptransportistas.Guia
	(*Producto)(nil),              // 5: apptransportistas.Producto
	(*Deposito)(nil),              // 6: apptransportistas.Deposito
	(*PruebaEntrega)(nil),         // 7: apptransportistas.PruebaEntrega
}
var file_apptransportistas_proto_depIdxs = []int32{
	5, // 0: apptransportistas.Guia.productos:type_name -> apptransportistas.Producto
	4, // 1: apptransportistas.AppTransportistasService.EnviarEntregas:input_type -> apptransportistas.Guia
	0, // 2: apptransportistas.AppTransportistasService.ObtenerDespachos:input_type -> apptransportistas.DespachoRequest
	6, // 3: apptransportistas.AppTransportistasService.EnviarDepositos:input_type -> apptransportistas.Deposito
	7, // 4: apptransportistas.AppTransportistasService.EnviarPruebasEntrega:input_type -> apptransportistas.PruebaEntrega
	1, // 5: apptransportistas.AppTransportistasService.EnviarEntregas:output_type -> apptransportistas.EntregaResponse
	4, // 6: apptransportistas.AppTransportistasService.ObtenerDespachos:output_type -> apptransportistas.Guia
	2, // 7: apptransportistas.AppTransportistasService.EnviarDepositos:output_type -> apptransportistas.DepositoResponse
	3, // 8: apptransportistas.AppTransportistasService.EnviarPruebasEntrega:output_type -> apptransportistas.PruebaEntregaResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apptransportistas_proto_init() }
func file_apptransportistas_proto_init() {
	if File_apptransportistas_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_apptransportistas_proto_rawDesc), len(file_apptransportistas_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apptransportistas_proto_goTypes,
		DependencyIndexes: file_apptransportistas_proto_depIdxs,
		MessageInfos:      file_apptransportistas_proto_msgTypes,
	}.Build()
	File_apptransportistas_proto = out.File
	file_apptransportistas_proto_goTypes = nil
	file_apptransportistas_proto_depIdxs = nil
}
