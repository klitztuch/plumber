// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opts/ps_opts_server.proto

package opts

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

type ServerOptions struct {
	// @gotags: kong:"default=plumber1,help='Unique ID that identifies this plumber node',env='PLUMBER_SERVER_NODE_ID',required"
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty" kong:"default=plumber1,help='Unique ID that identifies this plumber node',env='PLUMBER_SERVER_NODE_ID',required"`
	// @gotags: kong:"default=aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa,help='ID of the plumber cluster (has to be the same across all nodes)',env='PLUMBER_SERVER_CLUSTER_ID',required"
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" kong:"default=aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa,help='ID of the plumber cluster (has to be the same across all nodes)',env='PLUMBER_SERVER_CLUSTER_ID',required"`
	// @gotags: kong:"help='Host:port that the gRPC server will bind to',env='PLUMBER_SERVER_GRPC_LISTEN_ADDRESS',default=127.0.0.1:9090"
	GrpcListenAddress string `protobuf:"bytes,3,opt,name=grpc_listen_address,json=grpcListenAddress,proto3" json:"grpc_listen_address,omitempty" kong:"help='Host:port that the gRPC server will bind to',env='PLUMBER_SERVER_GRPC_LISTEN_ADDRESS',default=127.0.0.1:9090"`
	// @gotags: kong:"default=batchcorp,help='All gRPC requests require this auth token to be set',env='PLUMBER_SERVER_AUTH_TOKEN',required"
	AuthToken string `protobuf:"bytes,4,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty" kong:"default=batchcorp,help='All gRPC requests require this auth token to be set',env='PLUMBER_SERVER_AUTH_TOKEN',required"`
	// @gotags: kong:"help='Comma separated list of NATS server URLs (can contain user:password if using auth; only used if --enable-cluster is true)',env='PLUMBER_SERVER_NATS_URL',default='nats://localhost:4222'"
	NatsUrl []string `protobuf:"bytes,5,rep,name=nats_url,json=natsUrl,proto3" json:"nats_url,omitempty" kong:"help='Comma separated list of NATS server URLs (can contain user:password if using auth; only used if --enable-cluster is true)',env='PLUMBER_SERVER_NATS_URL',default='nats://localhost:4222'"`
	// @gotags: kong:"help='Whether to use TLS (only used if --enable-cluster is true)',env='PLUMBER_SERVER_USE_TLS'"
	UseTls bool `protobuf:"varint,500,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty" kong:"help='Whether to use TLS (only used if --enable-cluster is true)',env='PLUMBER_SERVER_USE_TLS'"`
	// @gotags: kong:"help='TLS client cert file (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_CERT_FILE'"
	TlsCertFile string `protobuf:"bytes,6,opt,name=tls_cert_file,json=tlsCertFile,proto3" json:"tls_cert_file,omitempty" kong:"help='TLS client cert file (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_CERT_FILE'"`
	// @gotags: kong:"help='TLS client key file (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_KEY_FILE'"
	TlsKeyFile string `protobuf:"bytes,7,opt,name=tls_key_file,json=tlsKeyFile,proto3" json:"tls_key_file,omitempty" kong:"help='TLS client key file (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_KEY_FILE'"`
	// @gotags: kong:"help='TLS CA certificate file (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_CA_FILE'"
	TlsCaFile string `protobuf:"bytes,8,opt,name=tls_ca_file,json=tlsCaFile,proto3" json:"tls_ca_file,omitempty" kong:"help='TLS CA certificate file (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_CA_FILE'"`
	// @gotags: kong:"help='Skip server cert verification (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_SKIP_VERIFY',default=false"
	TlsSkipVerify bool `protobuf:"varint,9,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty" kong:"help='Skip server cert verification (only used if --enable-cluster is true)',env='PLUMBER_SERVER_TLS_SKIP_VERIFY',default=false"`
	// @gotags: kong:"help='Run plumber in cluster mode (will use NATS)',env='PLUMBER_SERVER_ENABLE_CLUSTER',default=false"
	EnableCluster bool `protobuf:"varint,10,opt,name=enable_cluster,json=enableCluster,proto3" json:"enable_cluster,omitempty" kong:"help='Run plumber in cluster mode (will use NATS)',env='PLUMBER_SERVER_ENABLE_CLUSTER',default=false"`
	// @gotags: kong:"help='VC-Service gRPC server address',default='https://vc-service.batch.sh'"
	VcserviceGrpcAddress string `protobuf:"bytes,11,opt,name=vcservice_grpc_address,json=vcserviceGrpcAddress,proto3" json:"vcservice_grpc_address,omitempty" kong:"help='VC-Service gRPC server address',default='https://vc-service.batch.sh'"`
	// @gotags: kong:"help='VC-Service gRPC  server initial connection timeout',default=5"
	VcserviceGrpcTimeoutSeconds uint32 `protobuf:"varint,12,opt,name=vcservice_grpc_timeout_seconds,json=vcserviceGrpcTimeoutSeconds,proto3" json:"vcservice_grpc_timeout_seconds,omitempty" kong:"help='VC-Service gRPC  server initial connection timeout',default=5"`
	// @gotags: kong:"help='Use gRPC insecure mode when talking to VC-Service'"
	VcserviceGrpcInsecure bool `protobuf:"varint,13,opt,name=vcservice_grpc_insecure,json=vcserviceGrpcInsecure,proto3" json:"vcservice_grpc_insecure,omitempty" kong:"help='Use gRPC insecure mode when talking to VC-Service'"`
	// @gotags: kong:"help='Location to store time-series data for counters',default='./.tsdata'"
	StatsDatabasePath string `protobuf:"bytes,14,opt,name=stats_database_path,json=statsDatabasePath,proto3" json:"stats_database_path,omitempty" kong:"help='Location to store time-series data for counters',default='./.tsdata'"`
	// @gotags: kong:"help='How often to flush time-series data (in seconds) from memory to storage',default='60'"
	StatsFlushIntervalSeconds int32 `protobuf:"varint,15,opt,name=stats_flush_interval_seconds,json=statsFlushIntervalSeconds,proto3" json:"stats_flush_interval_seconds,omitempty" kong:"help='How often to flush time-series data (in seconds) from memory to storage',default='60'"`
	// @gotags: kong:"help='What address to bind the built-in HTTP server to',default='127.0.0.1:9191'"
	HttpListenAddress string `protobuf:"bytes,16,opt,name=http_listen_address,json=httpListenAddress,proto3" json:"http_listen_address,omitempty" kong:"help='What address to bind the built-in HTTP server to',default='127.0.0.1:9191'"`
	// @gotags: kong:"default='./.plumber',help='Location to store etcd and time series files.',env='PLUMBER_SERVER_STORAGE_PATH'"
	StoragePath          string   `protobuf:"bytes,17,opt,name=storage_path,json=storagePath,proto3" json:"storage_path,omitempty" kong:"default='./.plumber',help='Location to store etcd and time series files.',env='PLUMBER_SERVER_STORAGE_PATH'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerOptions) Reset()         { *m = ServerOptions{} }
func (m *ServerOptions) String() string { return proto.CompactTextString(m) }
func (*ServerOptions) ProtoMessage()    {}
func (*ServerOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_72fba0b7ae2941aa, []int{0}
}

func (m *ServerOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerOptions.Unmarshal(m, b)
}
func (m *ServerOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerOptions.Marshal(b, m, deterministic)
}
func (m *ServerOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerOptions.Merge(m, src)
}
func (m *ServerOptions) XXX_Size() int {
	return xxx_messageInfo_ServerOptions.Size(m)
}
func (m *ServerOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ServerOptions proto.InternalMessageInfo

func (m *ServerOptions) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ServerOptions) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ServerOptions) GetGrpcListenAddress() string {
	if m != nil {
		return m.GrpcListenAddress
	}
	return ""
}

func (m *ServerOptions) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func (m *ServerOptions) GetNatsUrl() []string {
	if m != nil {
		return m.NatsUrl
	}
	return nil
}

func (m *ServerOptions) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *ServerOptions) GetTlsCertFile() string {
	if m != nil {
		return m.TlsCertFile
	}
	return ""
}

func (m *ServerOptions) GetTlsKeyFile() string {
	if m != nil {
		return m.TlsKeyFile
	}
	return ""
}

func (m *ServerOptions) GetTlsCaFile() string {
	if m != nil {
		return m.TlsCaFile
	}
	return ""
}

func (m *ServerOptions) GetTlsSkipVerify() bool {
	if m != nil {
		return m.TlsSkipVerify
	}
	return false
}

func (m *ServerOptions) GetEnableCluster() bool {
	if m != nil {
		return m.EnableCluster
	}
	return false
}

func (m *ServerOptions) GetVcserviceGrpcAddress() string {
	if m != nil {
		return m.VcserviceGrpcAddress
	}
	return ""
}

func (m *ServerOptions) GetVcserviceGrpcTimeoutSeconds() uint32 {
	if m != nil {
		return m.VcserviceGrpcTimeoutSeconds
	}
	return 0
}

func (m *ServerOptions) GetVcserviceGrpcInsecure() bool {
	if m != nil {
		return m.VcserviceGrpcInsecure
	}
	return false
}

func (m *ServerOptions) GetStatsDatabasePath() string {
	if m != nil {
		return m.StatsDatabasePath
	}
	return ""
}

func (m *ServerOptions) GetStatsFlushIntervalSeconds() int32 {
	if m != nil {
		return m.StatsFlushIntervalSeconds
	}
	return 0
}

func (m *ServerOptions) GetHttpListenAddress() string {
	if m != nil {
		return m.HttpListenAddress
	}
	return ""
}

func (m *ServerOptions) GetStoragePath() string {
	if m != nil {
		return m.StoragePath
	}
	return ""
}

func init() {
	proto.RegisterType((*ServerOptions)(nil), "protos.opts.ServerOptions")
}

func init() { proto.RegisterFile("opts/ps_opts_server.proto", fileDescriptor_72fba0b7ae2941aa) }

var fileDescriptor_72fba0b7ae2941aa = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xd1, 0x6b, 0x53, 0x3f,
	0x14, 0xc7, 0xe9, 0x6f, 0xbf, 0xb5, 0x6b, 0xba, 0x6e, 0x2e, 0x53, 0x97, 0xa1, 0x8e, 0x3a, 0x50,
	0xfa, 0x62, 0xef, 0x83, 0x22, 0x88, 0x0f, 0xa2, 0x93, 0x49, 0x51, 0x50, 0xba, 0xe9, 0x83, 0x2f,
	0x21, 0x37, 0xf7, 0xac, 0x37, 0x34, 0xbd, 0x09, 0x39, 0xe7, 0x16, 0xf6, 0x7f, 0x0b, 0xbe, 0x4a,
	0x92, 0xbb, 0xc1, 0xf6, 0x14, 0x72, 0x3e, 0xdf, 0x7c, 0xf3, 0xe5, 0x9c, 0x84, 0x1d, 0x3b, 0x4f,
	0x58, 0x78, 0x94, 0x71, 0x95, 0x08, 0x61, 0x03, 0x61, 0xe6, 0x83, 0x23, 0xc7, 0x47, 0x69, 0xc1,
	0x59, 0x24, 0xa7, 0x7f, 0xb7, 0xd9, 0xf8, 0x22, 0xd1, 0xef, 0x9e, 0x8c, 0x6b, 0x90, 0x1f, 0xb1,
	0x41, 0xe3, 0x2a, 0x90, 0xa6, 0x12, 0xbd, 0x49, 0x6f, 0x3a, 0x5c, 0xf4, 0xe3, 0x76, 0x5e, 0xf1,
	0x67, 0x8c, 0x69, 0xdb, 0x22, 0x41, 0x88, 0xec, 0xbf, 0xc4, 0x86, 0x5d, 0x65, 0x5e, 0xf1, 0x19,
	0x3b, 0x5c, 0x06, 0xaf, 0xa5, 0x35, 0x48, 0xd0, 0x48, 0x55, 0x55, 0x01, 0x10, 0xc5, 0x56, 0xd2,
	0x1d, 0x44, 0xf4, 0x2d, 0x91, 0x8f, 0x19, 0x44, 0x3b, 0xd5, 0x52, 0x2d, 0xc9, 0xad, 0xa0, 0x11,
	0xff, 0x67, 0xbb, 0x58, 0xb9, 0x8c, 0x05, 0x7e, 0xcc, 0x76, 0x1a, 0x45, 0x28, 0xdb, 0x60, 0xc5,
	0xf6, 0x64, 0x6b, 0x3a, 0x5c, 0x0c, 0xe2, 0xfe, 0x67, 0xb0, 0x5c, 0xb0, 0x41, 0x8b, 0x20, 0xc9,
	0xa2, 0xf8, 0x13, 0xed, 0x77, 0x16, 0xfd, 0x16, 0xe1, 0xd2, 0x22, 0x3f, 0x65, 0x63, 0xb2, 0x28,
	0x35, 0x04, 0x92, 0x57, 0xc6, 0x82, 0xe8, 0x27, 0xdb, 0x11, 0x59, 0x3c, 0x83, 0x40, 0xe7, 0xc6,
	0x02, 0x9f, 0xb0, 0xdd, 0xa8, 0x59, 0xc1, 0x75, 0x96, 0x0c, 0x92, 0x84, 0x91, 0xc5, 0xaf, 0x70,
	0x9d, 0x14, 0x27, 0x6c, 0x94, 0x5c, 0x54, 0x16, 0xec, 0xe4, 0x68, 0xd1, 0x43, 0x25, 0xfe, 0x92,
	0xed, 0x47, 0x8e, 0x2b, 0xe3, 0xe5, 0x06, 0x82, 0xb9, 0xba, 0x16, 0xc3, 0x14, 0x23, 0x5e, 0x7e,
	0xb1, 0x32, 0xfe, 0x57, 0x2a, 0xf2, 0x17, 0x6c, 0x0f, 0x1a, 0x55, 0x5a, 0x90, 0x5d, 0x97, 0x04,
	0xcb, 0xb2, 0x5c, 0x3d, 0xcb, 0x45, 0xfe, 0x86, 0x3d, 0xde, 0xe8, 0x38, 0x21, 0xa3, 0x41, 0xa6,
	0x16, 0xde, 0xf4, 0x6e, 0x94, 0x6e, 0x7e, 0x78, 0x4b, 0xbf, 0x04, 0xaf, 0x6f, 0xda, 0x77, 0xc6,
	0x4e, 0xee, 0x9d, 0x22, 0xb3, 0x06, 0xd7, 0x92, 0x44, 0xd0, 0xae, 0xa9, 0x50, 0xec, 0x4e, 0x7a,
	0xd3, 0xf1, 0xe2, 0xc9, 0x9d, 0xd3, 0x97, 0x59, 0x73, 0x91, 0x25, 0xfc, 0x2d, 0x3b, 0xba, 0x67,
	0x62, 0x1a, 0x04, 0xdd, 0x06, 0x10, 0xe3, 0x14, 0xf5, 0xd1, 0x9d, 0xd3, 0xf3, 0x0e, 0xc6, 0x59,
	0x23, 0xc5, 0xe9, 0x54, 0x8a, 0x54, 0xa9, 0x10, 0xa4, 0x57, 0x54, 0x8b, 0xbd, 0x3c, 0xeb, 0x84,
	0x3e, 0x77, 0xe4, 0x87, 0xa2, 0x9a, 0x7f, 0x60, 0x4f, 0xb3, 0xfe, 0xca, 0xb6, 0x58, 0x4b, 0xd3,
	0x10, 0x84, 0x8d, 0xb2, 0xb7, 0x51, 0xf7, 0x27, 0xbd, 0xe9, 0xf6, 0xe2, 0x38, 0x69, 0xce, 0xa3,
	0x64, 0xde, 0x29, 0x6e, 0x82, 0xce, 0xd8, 0x61, 0x4d, 0xe4, 0xef, 0x3f, 0xae, 0x07, 0xf9, 0xc2,
	0x88, 0xee, 0x3e, 0xae, 0xe7, 0x6c, 0x17, 0xc9, 0x05, 0xb5, 0xec, 0x92, 0x1d, 0xe4, 0x77, 0xd0,
	0xd5, 0x62, 0xa6, 0x4f, 0xef, 0x7f, 0xbf, 0x5b, 0x1a, 0xaa, 0xdb, 0x72, 0xa6, 0xdd, 0xba, 0x28,
	0x15, 0xe9, 0x5a, 0xbb, 0xe0, 0x0b, 0x6f, 0xdb, 0x75, 0x09, 0xe1, 0x15, 0xea, 0x1a, 0xd6, 0x0a,
	0x8b, 0xb2, 0x35, 0xb6, 0x2a, 0x96, 0xae, 0xc8, 0xdf, 0xa6, 0x88, 0xdf, 0xa6, 0xec, 0xa7, 0xcd,
	0xeb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x86, 0xff, 0x47, 0x67, 0x03, 0x00, 0x00,
}
