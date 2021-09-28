// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Schema_Type int32

const (
	Schema_UNKNOWN  Schema_Type = 0
	Schema_PLAIN    Schema_Type = 1
	Schema_JSON     Schema_Type = 2
	Schema_PROTOBUF Schema_Type = 3
)

var Schema_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "PLAIN",
	2: "JSON",
	3: "PROTOBUF",
}

var Schema_Type_value = map[string]int32{
	"UNKNOWN":  0,
	"PLAIN":    1,
	"JSON":     2,
	"PROTOBUF": 3,
}

func (x Schema_Type) String() string {
	return proto.EnumName(Schema_Type_name, int32(x))
}

func (Schema_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0, 0}
}

type Schema_ElectionStatus int32

const (
	Schema_UNSET   Schema_ElectionStatus = 0
	Schema_SUCCESS Schema_ElectionStatus = 1
	Schema_FAILED  Schema_ElectionStatus = 2
)

var Schema_ElectionStatus_name = map[int32]string{
	0: "UNSET",
	1: "SUCCESS",
	2: "FAILED",
}

var Schema_ElectionStatus_value = map[string]int32{
	"UNSET":   0,
	"SUCCESS": 1,
	"FAILED":  2,
}

func (x Schema_ElectionStatus) String() string {
	return proto.EnumName(Schema_ElectionStatus_name, int32(x))
}

func (Schema_ElectionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0, 1}
}

// deprecated
type Schema_UpdateType int32

const (
	Schema_INITIAL  Schema_UpdateType = 0
	Schema_EXISTING Schema_UpdateType = 1
)

var Schema_UpdateType_name = map[int32]string{
	0: "INITIAL",
	1: "EXISTING",
}

var Schema_UpdateType_value = map[string]int32{
	"INITIAL":  0,
	"EXISTING": 1,
}

func (x Schema_UpdateType) String() string {
	return proto.EnumName(Schema_UpdateType_name, int32(x))
}

func (Schema_UpdateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0, 2}
}

type Schema struct {
	// The collector will ONLY fill out the 'id' for incoming messages - it is
	// the responsibility of downstream consumers to lookup the corresponding
	// schema configuration by 'id'.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Indicates the "data type" - what format are the collectors expecting to
	// receive the events in?
	Type Schema_Type `protobuf:"varint,2,opt,name=type,proto3,enum=events.Schema_Type" json:"type,omitempty"`
	// Not sure what this is used for - super vague; Mark wasn't able to find any uses
	Raw map[string][]byte `protobuf:"bytes,3,rep,name=raw,proto3" json:"raw,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Deprecated: Do not use.
	// Only used when Type == PROTOBUF
	ProtobufMessageName       string `protobuf:"bytes,4,opt,name=protobuf_message_name,json=protobufMessageName,proto3" json:"protobuf_message_name,omitempty"`
	ProtobufFileDescriptorSet []byte `protobuf:"bytes,5,opt,name=protobuf_file_descriptor_set,json=protobufFileDescriptorSet,proto3" json:"protobuf_file_descriptor_set,omitempty"`
	// The following fields are used by the schema-manager to facilitate schema updates
	UpdateType          Schema_UpdateType `protobuf:"varint,6,opt,name=update_type,json=updateType,proto3,enum=events.Schema_UpdateType" json:"update_type,omitempty"` // Deprecated: Do not use.
	UpdateCollectToken  string            `protobuf:"bytes,7,opt,name=update_collect_token,json=updateCollectToken,proto3" json:"update_collect_token,omitempty"`      // Deprecated: Do not use.
	UpdateParquetSchema []byte            `protobuf:"bytes,8,opt,name=update_parquet_schema,json=updateParquetSchema,proto3" json:"update_parquet_schema,omitempty"`   // Deprecated: Do not use.
	UpdateSqlSchema     []byte            `protobuf:"bytes,9,opt,name=update_sql_schema,json=updateSqlSchema,proto3" json:"update_sql_schema,omitempty"`               // Deprecated: Do not use.
	UpdateFingerprint   string            `protobuf:"bytes,10,opt,name=update_fingerprint,json=updateFingerprint,proto3" json:"update_fingerprint,omitempty"`          // Deprecated: Do not use.
	UpdateCollectId     string            `protobuf:"bytes,11,opt,name=update_collect_id,json=updateCollectId,proto3" json:"update_collect_id,omitempty"`              // Deprecated: Do not use.
	// Schema version is used to create unique collect-update-* topics which in
	// turn allow the writer to write data using correct schema when there are
	// more than 1 schema updates in-flight. Talk to MG or DS.
	//
	// This field is incremented by the collectors on a schema update.
	SchemaVersion int64 `protobuf:"varint,12,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"` // Deprecated: Do not use.
	// The manifest message payload, as JSON, to infer the schema from.
	// Only used when Message.Type == GENERATE_SCHEMA
	SourcePayload []byte `protobuf:"bytes,13,opt,name=source_payload,json=sourcePayload,proto3" json:"source_payload,omitempty"` // Deprecated: Do not use.
	// A single election can reference multiple batches - this ID represents
	// this specific election event. Set by athena-writer, used by schema-manager
	// for the etcd key under /schema-manager/elections/$collection_id/<$election_id>
	ElectionId string `protobuf:"bytes,99,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	// What collection this schema message pertains to
	ElectionCollectionId string `protobuf:"bytes,100,opt,name=election_collection_id,json=electionCollectionId,proto3" json:"election_collection_id,omitempty"`
	// Indicates the status of a completed election; set by schema-manager
	ElectionStatus Schema_ElectionStatus `protobuf:"varint,101,opt,name=election_status,json=electionStatus,proto3,enum=events.Schema_ElectionStatus" json:"election_status,omitempty"`
	// In case of election failure, this field explains what happened
	ElectionStatusMessage string `protobuf:"bytes,102,opt,name=election_status_message,json=electionStatusMessage,proto3" json:"election_status_message,omitempty"`
	// Which batches this election pertains to.
	// Every athena-writer assigns a unique batch-id for the batch they are
	// working on/electing. If a STATUS_ELECT_SCHEMA comes in, athena-writer
	// needs to check that the STATUS_ELECT_SCHEMA pertains to both the
	// collection_id AND batch id before it uses the elected schema. If the
	// batch id is not found, then their ELECT_SCHEMA probably pertains to
	// another election (and should continue waiting for another
	// STATUS_ELECT_SCHEMA message).
	//
	// This field is used by both - athena-writer and schema-manager:
	//
	// athena-writer uses it to specify what batch this ELECT_SCHEMA pertains to
	// schema-manager uses it to indicate which batches an STATUS_ELECT_SCHEMA pertains to
	ElectionBatchIds []string `protobuf:"bytes,103,rep,name=election_batch_ids,json=electionBatchIds,proto3" json:"election_batch_ids,omitempty"`
	// Set by athena-writer IF schema election succeeded
	ElectionParquetSchema []byte   `protobuf:"bytes,104,opt,name=election_parquet_schema,json=electionParquetSchema,proto3" json:"election_parquet_schema,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Schema) Reset()         { *m = Schema{} }
func (m *Schema) String() string { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()    {}
func (*Schema) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}

func (m *Schema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schema.Unmarshal(m, b)
}
func (m *Schema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schema.Marshal(b, m, deterministic)
}
func (m *Schema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schema.Merge(m, src)
}
func (m *Schema) XXX_Size() int {
	return xxx_messageInfo_Schema.Size(m)
}
func (m *Schema) XXX_DiscardUnknown() {
	xxx_messageInfo_Schema.DiscardUnknown(m)
}

var xxx_messageInfo_Schema proto.InternalMessageInfo

func (m *Schema) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Schema) GetType() Schema_Type {
	if m != nil {
		return m.Type
	}
	return Schema_UNKNOWN
}

// Deprecated: Do not use.
func (m *Schema) GetRaw() map[string][]byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Schema) GetProtobufMessageName() string {
	if m != nil {
		return m.ProtobufMessageName
	}
	return ""
}

func (m *Schema) GetProtobufFileDescriptorSet() []byte {
	if m != nil {
		return m.ProtobufFileDescriptorSet
	}
	return nil
}

// Deprecated: Do not use.
func (m *Schema) GetUpdateType() Schema_UpdateType {
	if m != nil {
		return m.UpdateType
	}
	return Schema_INITIAL
}

// Deprecated: Do not use.
func (m *Schema) GetUpdateCollectToken() string {
	if m != nil {
		return m.UpdateCollectToken
	}
	return ""
}

// Deprecated: Do not use.
func (m *Schema) GetUpdateParquetSchema() []byte {
	if m != nil {
		return m.UpdateParquetSchema
	}
	return nil
}

// Deprecated: Do not use.
func (m *Schema) GetUpdateSqlSchema() []byte {
	if m != nil {
		return m.UpdateSqlSchema
	}
	return nil
}

// Deprecated: Do not use.
func (m *Schema) GetUpdateFingerprint() string {
	if m != nil {
		return m.UpdateFingerprint
	}
	return ""
}

// Deprecated: Do not use.
func (m *Schema) GetUpdateCollectId() string {
	if m != nil {
		return m.UpdateCollectId
	}
	return ""
}

// Deprecated: Do not use.
func (m *Schema) GetSchemaVersion() int64 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

// Deprecated: Do not use.
func (m *Schema) GetSourcePayload() []byte {
	if m != nil {
		return m.SourcePayload
	}
	return nil
}

func (m *Schema) GetElectionId() string {
	if m != nil {
		return m.ElectionId
	}
	return ""
}

func (m *Schema) GetElectionCollectionId() string {
	if m != nil {
		return m.ElectionCollectionId
	}
	return ""
}

func (m *Schema) GetElectionStatus() Schema_ElectionStatus {
	if m != nil {
		return m.ElectionStatus
	}
	return Schema_UNSET
}

func (m *Schema) GetElectionStatusMessage() string {
	if m != nil {
		return m.ElectionStatusMessage
	}
	return ""
}

func (m *Schema) GetElectionBatchIds() []string {
	if m != nil {
		return m.ElectionBatchIds
	}
	return nil
}

func (m *Schema) GetElectionParquetSchema() []byte {
	if m != nil {
		return m.ElectionParquetSchema
	}
	return nil
}

func init() {
	proto.RegisterEnum("events.Schema_Type", Schema_Type_name, Schema_Type_value)
	proto.RegisterEnum("events.Schema_ElectionStatus", Schema_ElectionStatus_name, Schema_ElectionStatus_value)
	proto.RegisterEnum("events.Schema_UpdateType", Schema_UpdateType_name, Schema_UpdateType_value)
	proto.RegisterType((*Schema)(nil), "events.Schema")
	proto.RegisterMapType((map[string][]byte)(nil), "events.Schema.RawEntry")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0x41, 0x6f, 0xda, 0x4c,
	0x10, 0x86, 0x63, 0x4c, 0x08, 0x0c, 0x84, 0xf8, 0xdb, 0x24, 0x5f, 0x36, 0x55, 0xab, 0x22, 0x2e,
	0xa1, 0x52, 0x6a, 0xd4, 0x34, 0x42, 0x55, 0x0f, 0xad, 0x02, 0x81, 0xca, 0x6d, 0x6a, 0x90, 0x0d,
	0x6d, 0xd5, 0x8b, 0x65, 0xec, 0x05, 0xac, 0x18, 0xdb, 0xb1, 0xd7, 0x89, 0xf8, 0x07, 0xfd, 0xd9,
	0xd5, 0xee, 0xda, 0x4e, 0xe1, 0xe6, 0xf5, 0xfb, 0xbc, 0x3b, 0x33, 0x3b, 0x33, 0xd0, 0x48, 0x9c,
	0x15, 0x59, 0xdb, 0x6a, 0x14, 0x87, 0x34, 0x44, 0x15, 0xf2, 0x48, 0x02, 0x9a, 0xb4, 0xff, 0xd4,
	0xa0, 0x62, 0x72, 0x01, 0x35, 0xa1, 0xe4, 0xb9, 0x58, 0x6a, 0x49, 0x9d, 0x9a, 0x51, 0xf2, 0x5c,
	0x74, 0x01, 0x65, 0xba, 0x89, 0x08, 0x2e, 0xb5, 0xa4, 0x4e, 0xf3, 0xea, 0x58, 0x15, 0x0e, 0x55,
	0xd0, 0xea, 0x74, 0x13, 0x11, 0x83, 0x03, 0xe8, 0x2d, 0xc8, 0xb1, 0xfd, 0x84, 0xe5, 0x96, 0xdc,
	0xa9, 0x5f, 0x9d, 0xed, 0x70, 0x86, 0xfd, 0x34, 0x0c, 0x68, 0xbc, 0xe9, 0x97, 0xb0, 0x64, 0x30,
	0x0e, 0x5d, 0xc1, 0x29, 0xcf, 0x61, 0x9e, 0x2e, 0xac, 0x35, 0x49, 0x12, 0x7b, 0x49, 0xac, 0xc0,
	0x5e, 0x13, 0x5c, 0xe6, 0xa1, 0x8f, 0x73, 0xf1, 0xbb, 0xd0, 0x74, 0x7b, 0x4d, 0xd0, 0x67, 0x78,
	0x59, 0x78, 0x16, 0x9e, 0x4f, 0x2c, 0x97, 0x24, 0x4e, 0xec, 0x45, 0x34, 0x8c, 0xad, 0x84, 0x50,
	0xbc, 0xdf, 0x92, 0x3a, 0x0d, 0xe3, 0x3c, 0x67, 0x46, 0x9e, 0x4f, 0x6e, 0x0b, 0xc2, 0x24, 0x14,
	0x7d, 0x82, 0x7a, 0x1a, 0xb9, 0x36, 0x25, 0x16, 0xaf, 0xa9, 0xc2, 0x6b, 0x3a, 0xdf, 0xc9, 0x75,
	0xc6, 0x09, 0x56, 0x19, 0xcf, 0x16, 0xd2, 0xe2, 0x8c, 0xae, 0xe1, 0x24, 0xf3, 0x3b, 0xa1, 0xef,
	0x13, 0x87, 0x5a, 0x34, 0xbc, 0x27, 0x01, 0x3e, 0x60, 0x39, 0x73, 0x1a, 0x09, 0x7d, 0x20, 0xe4,
	0x29, 0x53, 0x51, 0x0f, 0x4e, 0x33, 0x57, 0x64, 0xc7, 0x0f, 0x29, 0xa1, 0x96, 0x68, 0x02, 0xae,
	0xb2, 0x7c, 0xb9, 0xed, 0x58, 0x00, 0x13, 0xa1, 0x67, 0xad, 0x50, 0xe1, 0xbf, 0xcc, 0x97, 0x3c,
	0xf8, 0xb9, 0xa7, 0x56, 0x78, 0x8e, 0x84, 0x68, 0x3e, 0xf8, 0x19, 0xff, 0x0e, 0xb2, 0xe8, 0xd6,
	0xc2, 0x0b, 0x96, 0x24, 0x8e, 0x62, 0x2f, 0xa0, 0x18, 0x8a, 0xdc, 0xb2, 0xdb, 0x46, 0xcf, 0xe2,
	0x3f, 0x21, 0xf2, 0x82, 0x3c, 0x17, 0xd7, 0x0b, 0xc7, 0xd1, 0x56, 0x35, 0x9a, 0x8b, 0xde, 0x40,
	0x53, 0xe4, 0x61, 0x3d, 0x92, 0x38, 0xf1, 0xc2, 0x00, 0x37, 0x5a, 0x52, 0x47, 0xe6, 0xf0, 0xa1,
	0x50, 0x7e, 0x08, 0x81, 0xa3, 0x61, 0x1a, 0x3b, 0xac, 0xea, 0x8d, 0x1f, 0xda, 0x2e, 0x3e, 0x2c,
	0x52, 0x3f, 0x14, 0xca, 0x44, 0x08, 0xe8, 0x35, 0xd4, 0x09, 0x0b, 0xe0, 0x85, 0x01, 0x8b, 0xef,
	0xf0, 0x09, 0x80, 0xfc, 0x97, 0xe6, 0xa2, 0x6b, 0xf8, 0xbf, 0x00, 0xb2, 0x44, 0x33, 0xd6, 0xe5,
	0xec, 0x49, 0xae, 0x0e, 0x0a, 0x51, 0x73, 0xd1, 0x08, 0x8e, 0x0a, 0x57, 0x42, 0x6d, 0x9a, 0x26,
	0x98, 0xf0, 0x8e, 0xbf, 0xda, 0xe9, 0xf8, 0x30, 0xa3, 0x4c, 0x0e, 0x19, 0x4d, 0xb2, 0x75, 0x46,
	0x3d, 0x38, 0xdb, 0xb9, 0x27, 0x9f, 0x58, 0xbc, 0xe0, 0xe1, 0x4f, 0xb7, 0x0d, 0xd9, 0xc8, 0xa2,
	0x4b, 0x40, 0x85, 0x6f, 0x6e, 0x53, 0x67, 0x65, 0x79, 0x6e, 0x82, 0x97, 0x2d, 0xb9, 0x53, 0x33,
	0x94, 0x5c, 0xe9, 0x33, 0x41, 0x73, 0xb7, 0xa3, 0xec, 0xcc, 0xc9, 0x8a, 0xcf, 0x75, 0x11, 0x65,
	0x6b, 0x4a, 0x5e, 0xf4, 0xa0, 0x9a, 0x6f, 0x17, 0x52, 0x40, 0xbe, 0x27, 0x9b, 0x6c, 0x7b, 0xd9,
	0x27, 0x3a, 0x81, 0xfd, 0x47, 0xdb, 0x4f, 0xc5, 0xfe, 0x36, 0x0c, 0x71, 0xf8, 0x58, 0xfa, 0x20,
	0xb5, 0x7b, 0x50, 0xe6, 0x33, 0x5d, 0x87, 0x83, 0x99, 0xfe, 0x4d, 0x1f, 0xff, 0xd4, 0x95, 0x3d,
	0x54, 0x83, 0xfd, 0xc9, 0xdd, 0x8d, 0xa6, 0x2b, 0x12, 0xaa, 0x42, 0xf9, 0xab, 0x39, 0xd6, 0x95,
	0x12, 0x6a, 0x40, 0x75, 0x62, 0x8c, 0xa7, 0xe3, 0xfe, 0x6c, 0xa4, 0xc8, 0xed, 0x6b, 0x68, 0x6e,
	0xbf, 0x17, 0x33, 0xcd, 0x74, 0x73, 0x38, 0x55, 0xf6, 0xd8, 0x65, 0xe6, 0x6c, 0x30, 0x18, 0x9a,
	0xa6, 0x22, 0x21, 0x80, 0xca, 0xe8, 0x46, 0xbb, 0x1b, 0xde, 0x2a, 0xa5, 0xf6, 0x05, 0xc0, 0xf3,
	0x5e, 0x31, 0x4c, 0xd3, 0xb5, 0xa9, 0x76, 0x73, 0xa7, 0xec, 0xb1, 0xeb, 0x87, 0xbf, 0x34, 0x73,
	0xaa, 0xe9, 0x5f, 0x14, 0xa9, 0xaf, 0xfe, 0xbe, 0x5c, 0x7a, 0x74, 0x95, 0xce, 0x55, 0x27, 0x5c,
	0x77, 0xf9, 0xb3, 0x39, 0x61, 0x1c, 0x75, 0xc5, 0x53, 0x24, 0xdd, 0x79, 0xea, 0xf9, 0x6e, 0x77,
	0x19, 0x76, 0x45, 0x0b, 0xe7, 0x15, 0xbe, 0xed, 0xef, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xa9,
	0x12, 0x64, 0xbf, 0xd9, 0x04, 0x00, 0x00,
}