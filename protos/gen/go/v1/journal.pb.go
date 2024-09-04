// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: v1/journal.proto

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

type JournalEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Codigo       string  `protobuf:"bytes,2,opt,name=codigo,proto3" json:"codigo,omitempty"`
	Denominacion string  `protobuf:"bytes,3,opt,name=denominacion,proto3" json:"denominacion,omitempty"`
	EmpresaId    string  `protobuf:"bytes,4,opt,name=empresa_id,json=empresaId,proto3" json:"empresa_id,omitempty"`
	Efe          string  `protobuf:"bytes,5,opt,name=efe,proto3" json:"efe,omitempty"`
	Smvn         string  `protobuf:"bytes,6,opt,name=smvn,proto3" json:"smvn,omitempty"`
	Smve         string  `protobuf:"bytes,7,opt,name=smve,proto3" json:"smve,omitempty"`
	Tipo         string  `protobuf:"bytes,8,opt,name=tipo,proto3" json:"tipo,omitempty"`
	TipoId       string  `protobuf:"bytes,9,opt,name=tipo_id,json=tipoId,proto3" json:"tipo_id,omitempty"`
	Esf          bool    `protobuf:"varint,10,opt,name=esf,proto3" json:"esf,omitempty"`
	Erf          bool    `protobuf:"varint,11,opt,name=erf,proto3" json:"erf,omitempty"`
	Ern          bool    `protobuf:"varint,12,opt,name=ern,proto3" json:"ern,omitempty"`
	Efle         bool    `protobuf:"varint,13,opt,name=efle,proto3" json:"efle,omitempty"`
	Ecp          bool    `protobuf:"varint,14,opt,name=ecp,proto3" json:"ecp,omitempty"`
	CierreId     string  `protobuf:"bytes,15,opt,name=cierre_id,json=cierreId,proto3" json:"cierre_id,omitempty"`
	Sequence     int32   `protobuf:"varint,16,opt,name=sequence,proto3" json:"sequence,omitempty"`
	CreatedAt    string  `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string  `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt    string  `protobuf:"bytes,19,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Iidebe       float32 `protobuf:"fixed32,20,opt,name=iidebe,proto3" json:"iidebe,omitempty"`
	Iihaber      float32 `protobuf:"fixed32,21,opt,name=iihaber,proto3" json:"iihaber,omitempty"`
	Debe         float32 `protobuf:"fixed32,22,opt,name=debe,proto3" json:"debe,omitempty"`
	Haber        float32 `protobuf:"fixed32,23,opt,name=haber,proto3" json:"haber,omitempty"`
}

func (x *JournalEntry) Reset() {
	*x = JournalEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_journal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JournalEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JournalEntry) ProtoMessage() {}

func (x *JournalEntry) ProtoReflect() protoreflect.Message {
	mi := &file_v1_journal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JournalEntry.ProtoReflect.Descriptor instead.
func (*JournalEntry) Descriptor() ([]byte, []int) {
	return file_v1_journal_proto_rawDescGZIP(), []int{0}
}

func (x *JournalEntry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JournalEntry) GetCodigo() string {
	if x != nil {
		return x.Codigo
	}
	return ""
}

func (x *JournalEntry) GetDenominacion() string {
	if x != nil {
		return x.Denominacion
	}
	return ""
}

func (x *JournalEntry) GetEmpresaId() string {
	if x != nil {
		return x.EmpresaId
	}
	return ""
}

func (x *JournalEntry) GetEfe() string {
	if x != nil {
		return x.Efe
	}
	return ""
}

func (x *JournalEntry) GetSmvn() string {
	if x != nil {
		return x.Smvn
	}
	return ""
}

func (x *JournalEntry) GetSmve() string {
	if x != nil {
		return x.Smve
	}
	return ""
}

func (x *JournalEntry) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *JournalEntry) GetTipoId() string {
	if x != nil {
		return x.TipoId
	}
	return ""
}

func (x *JournalEntry) GetEsf() bool {
	if x != nil {
		return x.Esf
	}
	return false
}

func (x *JournalEntry) GetErf() bool {
	if x != nil {
		return x.Erf
	}
	return false
}

func (x *JournalEntry) GetErn() bool {
	if x != nil {
		return x.Ern
	}
	return false
}

func (x *JournalEntry) GetEfle() bool {
	if x != nil {
		return x.Efle
	}
	return false
}

func (x *JournalEntry) GetEcp() bool {
	if x != nil {
		return x.Ecp
	}
	return false
}

func (x *JournalEntry) GetCierreId() string {
	if x != nil {
		return x.CierreId
	}
	return ""
}

func (x *JournalEntry) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *JournalEntry) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *JournalEntry) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *JournalEntry) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *JournalEntry) GetIidebe() float32 {
	if x != nil {
		return x.Iidebe
	}
	return 0
}

func (x *JournalEntry) GetIihaber() float32 {
	if x != nil {
		return x.Iihaber
	}
	return 0
}

func (x *JournalEntry) GetDebe() float32 {
	if x != nil {
		return x.Debe
	}
	return 0
}

func (x *JournalEntry) GetHaber() float32 {
	if x != nil {
		return x.Haber
	}
	return 0
}

type RetrieveJournalReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Journals []*JournalEntry `protobuf:"bytes,1,rep,name=journals,proto3" json:"journals,omitempty"`
}

func (x *RetrieveJournalReportResponse) Reset() {
	*x = RetrieveJournalReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_journal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveJournalReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveJournalReportResponse) ProtoMessage() {}

func (x *RetrieveJournalReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_journal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveJournalReportResponse.ProtoReflect.Descriptor instead.
func (*RetrieveJournalReportResponse) Descriptor() ([]byte, []int) {
	return file_v1_journal_proto_rawDescGZIP(), []int{1}
}

func (x *RetrieveJournalReportResponse) GetJournals() []*JournalEntry {
	if x != nil {
		return x.Journals
	}
	return nil
}

type RetrieveJournalReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessId     string `protobuf:"bytes,1,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	Period         string `protobuf:"bytes,2,opt,name=period,proto3" json:"period,omitempty"`
	IsConsolidated bool   `protobuf:"varint,3,opt,name=is_consolidated,json=isConsolidated,proto3" json:"is_consolidated,omitempty"`
	IncludeClose   bool   `protobuf:"varint,4,opt,name=include_close,json=includeClose,proto3" json:"include_close,omitempty"`
	IncludeCuBa    bool   `protobuf:"varint,5,opt,name=include_cu_ba,json=includeCuBa,proto3" json:"include_cu_ba,omitempty"`
}

func (x *RetrieveJournalReportRequest) Reset() {
	*x = RetrieveJournalReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_journal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveJournalReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveJournalReportRequest) ProtoMessage() {}

func (x *RetrieveJournalReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_journal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveJournalReportRequest.ProtoReflect.Descriptor instead.
func (*RetrieveJournalReportRequest) Descriptor() ([]byte, []int) {
	return file_v1_journal_proto_rawDescGZIP(), []int{2}
}

func (x *RetrieveJournalReportRequest) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *RetrieveJournalReportRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *RetrieveJournalReportRequest) GetIsConsolidated() bool {
	if x != nil {
		return x.IsConsolidated
	}
	return false
}

func (x *RetrieveJournalReportRequest) GetIncludeClose() bool {
	if x != nil {
		return x.IncludeClose
	}
	return false
}

func (x *RetrieveJournalReportRequest) GetIncludeCuBa() bool {
	if x != nil {
		return x.IncludeCuBa
	}
	return false
}

var File_v1_journal_proto protoreflect.FileDescriptor

var file_v1_journal_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0xae, 0x04, 0x0a, 0x0c, 0x4a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x64, 0x69, 0x67,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x12,
	0x22, 0x0a, 0x0c, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x63,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x61, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x61,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x66, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x65, 0x66, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6d, 0x76, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6d, 0x76, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6d, 0x76, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6d, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x70, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x69, 0x70, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x69, 0x70, 0x6f, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x73, 0x66,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x65, 0x73, 0x66, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x72, 0x66, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x65, 0x72, 0x66, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x72, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x65, 0x72, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x65, 0x66, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x65,
	0x66, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x63, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x65, 0x63, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x65, 0x72, 0x72, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x65, 0x72, 0x72, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x69, 0x64, 0x65, 0x62, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x69, 0x69, 0x64,
	0x65, 0x62, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x69, 0x68, 0x61, 0x62, 0x65, 0x72, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x69, 0x69, 0x68, 0x61, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x65, 0x62, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x64, 0x65, 0x62,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x61, 0x62, 0x65, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x68, 0x61, 0x62, 0x65, 0x72, 0x22, 0x4d, 0x0a, 0x1d, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x22, 0xc9, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x12, 0x27, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x43, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x0d, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x63, 0x75, 0x5f, 0x62, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x75,
	0x42, 0x61, 0x32, 0x6e, 0x0a, 0x0e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x15, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4a, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x73, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x6e, 0x72, 0x79, 0x62, 0x72,
	0x61, 0x76, 0x6f, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56,
	0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_journal_proto_rawDescOnce sync.Once
	file_v1_journal_proto_rawDescData = file_v1_journal_proto_rawDesc
)

func file_v1_journal_proto_rawDescGZIP() []byte {
	file_v1_journal_proto_rawDescOnce.Do(func() {
		file_v1_journal_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_journal_proto_rawDescData)
	})
	return file_v1_journal_proto_rawDescData
}

var file_v1_journal_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_journal_proto_goTypes = []any{
	(*JournalEntry)(nil),                  // 0: v1.JournalEntry
	(*RetrieveJournalReportResponse)(nil), // 1: v1.RetrieveJournalReportResponse
	(*RetrieveJournalReportRequest)(nil),  // 2: v1.RetrieveJournalReportRequest
}
var file_v1_journal_proto_depIdxs = []int32{
	0, // 0: v1.RetrieveJournalReportResponse.journals:type_name -> v1.JournalEntry
	2, // 1: v1.JournalService.RetrieveJournalReport:input_type -> v1.RetrieveJournalReportRequest
	1, // 2: v1.JournalService.RetrieveJournalReport:output_type -> v1.RetrieveJournalReportResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_journal_proto_init() }
func file_v1_journal_proto_init() {
	if File_v1_journal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_journal_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*JournalEntry); i {
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
		file_v1_journal_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RetrieveJournalReportResponse); i {
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
		file_v1_journal_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RetrieveJournalReportRequest); i {
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
			RawDescriptor: file_v1_journal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_journal_proto_goTypes,
		DependencyIndexes: file_v1_journal_proto_depIdxs,
		MessageInfos:      file_v1_journal_proto_msgTypes,
	}.Build()
	File_v1_journal_proto = out.File
	file_v1_journal_proto_rawDesc = nil
	file_v1_journal_proto_goTypes = nil
	file_v1_journal_proto_depIdxs = nil
}