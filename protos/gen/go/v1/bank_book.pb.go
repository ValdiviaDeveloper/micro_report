// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: v1/bank_book.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BankBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BankBalance) Reset() {
	*x = BankBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bank_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankBalance) ProtoMessage() {}

func (x *BankBalance) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bank_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankBalance.ProtoReflect.Descriptor instead.
func (*BankBalance) Descriptor() ([]byte, []int) {
	return file_v1_bank_book_proto_rawDescGZIP(), []int{0}
}

func (x *BankBalance) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BankBalance) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type LBank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BId                  string  `protobuf:"bytes,1,opt,name=b_id,json=bId,proto3" json:"b_id,omitempty"`
	BPeriodo             string  `protobuf:"bytes,2,opt,name=b_periodo,json=bPeriodo,proto3" json:"b_periodo,omitempty"`
	OCuoTipo             string  `protobuf:"bytes,3,opt,name=o_cuo_tipo,json=oCuoTipo,proto3" json:"o_cuo_tipo,omitempty"`
	BIdentificador       string  `protobuf:"bytes,4,opt,name=b_identificador,json=bIdentificador,proto3" json:"b_identificador,omitempty"`
	TefCodigo            string  `protobuf:"bytes,5,opt,name=tef_codigo,json=tefCodigo,proto3" json:"tef_codigo,omitempty"`
	CfCuenta             string  `protobuf:"bytes,6,opt,name=cf_cuenta,json=cfCuenta,proto3" json:"cf_cuenta,omitempty"`
	OFechaEmion          string  `protobuf:"bytes,7,opt,name=o_fecha_emion,json=oFechaEmion,proto3" json:"o_fecha_emion,omitempty"`
	PgFecha              string  `protobuf:"bytes,8,opt,name=pg_fecha,json=pgFecha,proto3" json:"pg_fecha,omitempty"`
	TmpCodigo            string  `protobuf:"bytes,9,opt,name=tmp_codigo,json=tmpCodigo,proto3" json:"tmp_codigo,omitempty"`
	PgGlosa              string  `protobuf:"bytes,10,opt,name=pg_glosa,json=pgGlosa,proto3" json:"pg_glosa,omitempty"`
	OGlosa               string  `protobuf:"bytes,11,opt,name=o_glosa,json=oGlosa,proto3" json:"o_glosa,omitempty"`
	TdCodigo             string  `protobuf:"bytes,12,opt,name=td_codigo,json=tdCodigo,proto3" json:"td_codigo,omitempty"`
	PNumero              string  `protobuf:"bytes,13,opt,name=p_numero,json=pNumero,proto3" json:"p_numero,omitempty"`
	PRazonSocial         string  `protobuf:"bytes,14,opt,name=p_razon_social,json=pRazonSocial,proto3" json:"p_razon_social,omitempty"`
	BNumeroTransaccion   string  `protobuf:"bytes,15,opt,name=b_numero_transaccion,json=bNumeroTransaccion,proto3" json:"b_numero_transaccion,omitempty"`
	BDebe                float32 `protobuf:"fixed32,16,opt,name=b_debe,json=bDebe,proto3" json:"b_debe,omitempty"`
	BHaber               float32 `protobuf:"fixed32,17,opt,name=b_haber,json=bHaber,proto3" json:"b_haber,omitempty"`
	BEstado              string  `protobuf:"bytes,18,opt,name=b_estado,json=bEstado,proto3" json:"b_estado,omitempty"`
	PcCodigo             string  `protobuf:"bytes,19,opt,name=pc_codigo,json=pcCodigo,proto3" json:"pc_codigo,omitempty"`
	PcDenominacion       string  `protobuf:"bytes,20,opt,name=pc_denominacion,json=pcDenominacion,proto3" json:"pc_denominacion,omitempty"`
	OEstadoLe            string  `protobuf:"bytes,21,opt,name=o_estado_le,json=oEstadoLe,proto3" json:"o_estado_le,omitempty"`
	OObservaciones       string  `protobuf:"bytes,22,opt,name=o_observaciones,json=oObservaciones,proto3" json:"o_observaciones,omitempty"`
	OTipoCambio          float32 `protobuf:"fixed32,23,opt,name=o_tipo_cambio,json=oTipoCambio,proto3" json:"o_tipo_cambio,omitempty"`
	PgTipoCambio         float32 `protobuf:"fixed32,24,opt,name=pg_tipo_cambio,json=pgTipoCambio,proto3" json:"pg_tipo_cambio,omitempty"`
	PgFechaOFechaEmision string  `protobuf:"bytes,25,opt,name=pg_fecha_o_fecha_emision,json=pgFechaOFechaEmision,proto3" json:"pg_fecha_o_fecha_emision,omitempty"`
}

func (x *LBank) Reset() {
	*x = LBank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bank_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LBank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LBank) ProtoMessage() {}

func (x *LBank) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bank_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LBank.ProtoReflect.Descriptor instead.
func (*LBank) Descriptor() ([]byte, []int) {
	return file_v1_bank_book_proto_rawDescGZIP(), []int{1}
}

func (x *LBank) GetBId() string {
	if x != nil {
		return x.BId
	}
	return ""
}

func (x *LBank) GetBPeriodo() string {
	if x != nil {
		return x.BPeriodo
	}
	return ""
}

func (x *LBank) GetOCuoTipo() string {
	if x != nil {
		return x.OCuoTipo
	}
	return ""
}

func (x *LBank) GetBIdentificador() string {
	if x != nil {
		return x.BIdentificador
	}
	return ""
}

func (x *LBank) GetTefCodigo() string {
	if x != nil {
		return x.TefCodigo
	}
	return ""
}

func (x *LBank) GetCfCuenta() string {
	if x != nil {
		return x.CfCuenta
	}
	return ""
}

func (x *LBank) GetOFechaEmion() string {
	if x != nil {
		return x.OFechaEmion
	}
	return ""
}

func (x *LBank) GetPgFecha() string {
	if x != nil {
		return x.PgFecha
	}
	return ""
}

func (x *LBank) GetTmpCodigo() string {
	if x != nil {
		return x.TmpCodigo
	}
	return ""
}

func (x *LBank) GetPgGlosa() string {
	if x != nil {
		return x.PgGlosa
	}
	return ""
}

func (x *LBank) GetOGlosa() string {
	if x != nil {
		return x.OGlosa
	}
	return ""
}

func (x *LBank) GetTdCodigo() string {
	if x != nil {
		return x.TdCodigo
	}
	return ""
}

func (x *LBank) GetPNumero() string {
	if x != nil {
		return x.PNumero
	}
	return ""
}

func (x *LBank) GetPRazonSocial() string {
	if x != nil {
		return x.PRazonSocial
	}
	return ""
}

func (x *LBank) GetBNumeroTransaccion() string {
	if x != nil {
		return x.BNumeroTransaccion
	}
	return ""
}

func (x *LBank) GetBDebe() float32 {
	if x != nil {
		return x.BDebe
	}
	return 0
}

func (x *LBank) GetBHaber() float32 {
	if x != nil {
		return x.BHaber
	}
	return 0
}

func (x *LBank) GetBEstado() string {
	if x != nil {
		return x.BEstado
	}
	return ""
}

func (x *LBank) GetPcCodigo() string {
	if x != nil {
		return x.PcCodigo
	}
	return ""
}

func (x *LBank) GetPcDenominacion() string {
	if x != nil {
		return x.PcDenominacion
	}
	return ""
}

func (x *LBank) GetOEstadoLe() string {
	if x != nil {
		return x.OEstadoLe
	}
	return ""
}

func (x *LBank) GetOObservaciones() string {
	if x != nil {
		return x.OObservaciones
	}
	return ""
}

func (x *LBank) GetOTipoCambio() float32 {
	if x != nil {
		return x.OTipoCambio
	}
	return 0
}

func (x *LBank) GetPgTipoCambio() float32 {
	if x != nil {
		return x.PgTipoCambio
	}
	return 0
}

func (x *LBank) GetPgFechaOFechaEmision() string {
	if x != nil {
		return x.PgFechaOFechaEmision
	}
	return ""
}

type RetrieveBankBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BankBalances []*BankBalance `protobuf:"bytes,1,rep,name=bank_balances,json=bankBalances,proto3" json:"bank_balances,omitempty"`
	BankBooks    []*LBank       `protobuf:"bytes,2,rep,name=bank_books,json=bankBooks,proto3" json:"bank_books,omitempty"`
}

func (x *RetrieveBankBookResponse) Reset() {
	*x = RetrieveBankBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bank_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveBankBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveBankBookResponse) ProtoMessage() {}

func (x *RetrieveBankBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bank_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveBankBookResponse.ProtoReflect.Descriptor instead.
func (*RetrieveBankBookResponse) Descriptor() ([]byte, []int) {
	return file_v1_bank_book_proto_rawDescGZIP(), []int{2}
}

func (x *RetrieveBankBookResponse) GetBankBalances() []*BankBalance {
	if x != nil {
		return x.BankBalances
	}
	return nil
}

func (x *RetrieveBankBookResponse) GetBankBooks() []*LBank {
	if x != nil {
		return x.BankBooks
	}
	return nil
}

type RetrieveBankBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessId string   `protobuf:"bytes,1,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	Period     string   `protobuf:"bytes,2,opt,name=period,proto3" json:"period,omitempty"`
	AccountIds []string `protobuf:"bytes,3,rep,name=account_ids,json=accountIds,proto3" json:"account_ids,omitempty"`
}

func (x *RetrieveBankBookRequest) Reset() {
	*x = RetrieveBankBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bank_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveBankBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveBankBookRequest) ProtoMessage() {}

func (x *RetrieveBankBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bank_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveBankBookRequest.ProtoReflect.Descriptor instead.
func (*RetrieveBankBookRequest) Descriptor() ([]byte, []int) {
	return file_v1_bank_book_proto_rawDescGZIP(), []int{3}
}

func (x *RetrieveBankBookRequest) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *RetrieveBankBookRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *RetrieveBankBookRequest) GetAccountIds() []string {
	if x != nil {
		return x.AccountIds
	}
	return nil
}

var File_v1_bank_book_proto protoreflect.FileDescriptor

var file_v1_bank_book_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0x35, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x6b,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0xb8, 0x06, 0x0a, 0x05, 0x4c, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x11, 0x0a, 0x04, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x6f, 0x12, 0x1c, 0x0a, 0x0a, 0x6f, 0x5f, 0x63,
	0x75, 0x6f, 0x5f, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x43, 0x75, 0x6f, 0x54, 0x69, 0x70, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x64, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x62, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x64, 0x6f, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x66, 0x5f, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x66, 0x43, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x66, 0x5f, 0x63, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x66, 0x43, 0x75, 0x65, 0x6e, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0d,
	0x6f, 0x5f, 0x66, 0x65, 0x63, 0x68, 0x61, 0x5f, 0x65, 0x6d, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x46, 0x65, 0x63, 0x68, 0x61, 0x45, 0x6d, 0x69, 0x6f, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x67, 0x5f, 0x66, 0x65, 0x63, 0x68, 0x61, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x67, 0x46, 0x65, 0x63, 0x68, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x6d, 0x70, 0x43, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x67,
	0x5f, 0x67, 0x6c, 0x6f, 0x73, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x67,
	0x47, 0x6c, 0x6f, 0x73, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x5f, 0x67, 0x6c, 0x6f, 0x73, 0x61,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x47, 0x6c, 0x6f, 0x73, 0x61, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x64, 0x43, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x4e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x5f, 0x72, 0x61, 0x7a, 0x6f,
	0x6e, 0x5f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x52, 0x61, 0x7a, 0x6f, 0x6e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x14,
	0x62, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x63, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x62, 0x4e, 0x75, 0x6d,
	0x65, 0x72, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x15,
	0x0a, 0x06, 0x62, 0x5f, 0x64, 0x65, 0x62, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x62, 0x44, 0x65, 0x62, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x5f, 0x68, 0x61, 0x62, 0x65, 0x72,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x62, 0x48, 0x61, 0x62, 0x65, 0x72, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x5f, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x63, 0x5f,
	0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x63,
	0x43, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6e,
	0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x70, 0x63, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x0a, 0x0b, 0x6f, 0x5f, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x5f, 0x6c, 0x65, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x4c, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x6f, 0x5f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x63, 0x69, 0x6f, 0x6e,
	0x65, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x5f, 0x74, 0x69,
	0x70, 0x6f, 0x5f, 0x63, 0x61, 0x6d, 0x62, 0x69, 0x6f, 0x18, 0x17, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x6f, 0x54, 0x69, 0x70, 0x6f, 0x43, 0x61, 0x6d, 0x62, 0x69, 0x6f, 0x12, 0x24, 0x0a, 0x0e,
	0x70, 0x67, 0x5f, 0x74, 0x69, 0x70, 0x6f, 0x5f, 0x63, 0x61, 0x6d, 0x62, 0x69, 0x6f, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x70, 0x67, 0x54, 0x69, 0x70, 0x6f, 0x43, 0x61, 0x6d, 0x62,
	0x69, 0x6f, 0x12, 0x36, 0x0a, 0x18, 0x70, 0x67, 0x5f, 0x66, 0x65, 0x63, 0x68, 0x61, 0x5f, 0x6f,
	0x5f, 0x66, 0x65, 0x63, 0x68, 0x61, 0x5f, 0x65, 0x6d, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x67, 0x46, 0x65, 0x63, 0x68, 0x61, 0x4f, 0x46, 0x65,
	0x63, 0x68, 0x61, 0x45, 0x6d, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x7a, 0x0a, 0x18, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0c,
	0x62, 0x61, 0x6e, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0a,
	0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x09, 0x62, 0x61, 0x6e,
	0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x73, 0x0a, 0x17, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x32, 0x60, 0x0a, 0x0f, 0x42,
	0x61, 0x6e, 0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d,
	0x0a, 0x10, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x42, 0x61, 0x6e, 0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x61, 0x6e,
	0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x74, 0x0a,
	0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x42, 0x61, 0x6e, 0x6b, 0x42, 0x6f, 0x6f,
	0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x6e, 0x72, 0x79, 0x62, 0x72, 0x61, 0x76, 0x6f, 0x2f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_bank_book_proto_rawDescOnce sync.Once
	file_v1_bank_book_proto_rawDescData = file_v1_bank_book_proto_rawDesc
)

func file_v1_bank_book_proto_rawDescGZIP() []byte {
	file_v1_bank_book_proto_rawDescOnce.Do(func() {
		file_v1_bank_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_bank_book_proto_rawDescData)
	})
	return file_v1_bank_book_proto_rawDescData
}

var file_v1_bank_book_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_bank_book_proto_goTypes = []any{
	(*BankBalance)(nil),              // 0: v1.BankBalance
	(*LBank)(nil),                    // 1: v1.LBank
	(*RetrieveBankBookResponse)(nil), // 2: v1.RetrieveBankBookResponse
	(*RetrieveBankBookRequest)(nil),  // 3: v1.RetrieveBankBookRequest
}
var file_v1_bank_book_proto_depIdxs = []int32{
	0, // 0: v1.RetrieveBankBookResponse.bank_balances:type_name -> v1.BankBalance
	1, // 1: v1.RetrieveBankBookResponse.bank_books:type_name -> v1.LBank
	3, // 2: v1.BankBookService.RetrieveBankBook:input_type -> v1.RetrieveBankBookRequest
	2, // 3: v1.BankBookService.RetrieveBankBook:output_type -> v1.RetrieveBankBookResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_bank_book_proto_init() }
func file_v1_bank_book_proto_init() {
	if File_v1_bank_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_bank_book_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BankBalance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_bank_book_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LBank); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_bank_book_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RetrieveBankBookResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_bank_book_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RetrieveBankBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_bank_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_bank_book_proto_goTypes,
		DependencyIndexes: file_v1_bank_book_proto_depIdxs,
		MessageInfos:      file_v1_bank_book_proto_msgTypes,
	}.Build()
	File_v1_bank_book_proto = out.File
	file_v1_bank_book_proto_rawDesc = nil
	file_v1_bank_book_proto_goTypes = nil
	file_v1_bank_book_proto_depIdxs = nil
}
