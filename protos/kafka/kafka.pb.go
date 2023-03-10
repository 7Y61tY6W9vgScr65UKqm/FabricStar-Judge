// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/hyperledger/fabric_judge/protos/kafka/kafka.proto

package kafka // import "github.com/hyperledger/fabric_judge/protos/kafka"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KafkaMessageRegular_Class int32

const (
	KafkaMessageRegular_UNKNOWN KafkaMessageRegular_Class = 0
	KafkaMessageRegular_NORMAL  KafkaMessageRegular_Class = 1
	KafkaMessageRegular_CONFIG  KafkaMessageRegular_Class = 2
)

var KafkaMessageRegular_Class_name = map[int32]string{
	0: "UNKNOWN",
	1: "NORMAL",
	2: "CONFIG",
}
var KafkaMessageRegular_Class_value = map[string]int32{
	"UNKNOWN": 0,
	"NORMAL":  1,
	"CONFIG":  2,
}

func (x KafkaMessageRegular_Class) String() string {
	return proto.EnumName(KafkaMessageRegular_Class_name, int32(x))
}
func (KafkaMessageRegular_Class) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_kafka_8ca1a5cc9386535d, []int{1, 0}
}

// KafkaMessage is a wrapper type for the messages
// that the Kafka-based orderer deals with.
type KafkaMessage struct {
	// Types that are valid to be assigned to Type:
	//	*KafkaMessage_Regular
	//	*KafkaMessage_TimeToCut
	//	*KafkaMessage_Connect
	Type                 isKafkaMessage_Type `protobuf_oneof:"Type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *KafkaMessage) Reset()         { *m = KafkaMessage{} }
func (m *KafkaMessage) String() string { return proto.CompactTextString(m) }
func (*KafkaMessage) ProtoMessage()    {}
func (*KafkaMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_kafka_8ca1a5cc9386535d, []int{0}
}
func (m *KafkaMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaMessage.Unmarshal(m, b)
}
func (m *KafkaMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaMessage.Marshal(b, m, deterministic)
}
func (dst *KafkaMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaMessage.Merge(dst, src)
}
func (m *KafkaMessage) XXX_Size() int {
	return xxx_messageInfo_KafkaMessage.Size(m)
}
func (m *KafkaMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaMessage.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaMessage proto.InternalMessageInfo

type isKafkaMessage_Type interface {
	isKafkaMessage_Type()
}

type KafkaMessage_Regular struct {
	Regular *KafkaMessageRegular `protobuf:"bytes,1,opt,name=regular,proto3,oneof"`
}

type KafkaMessage_TimeToCut struct {
	TimeToCut *KafkaMessageTimeToCut `protobuf:"bytes,2,opt,name=time_to_cut,json=timeToCut,proto3,oneof"`
}

type KafkaMessage_Connect struct {
	Connect *KafkaMessageConnect `protobuf:"bytes,3,opt,name=connect,proto3,oneof"`
}

func (*KafkaMessage_Regular) isKafkaMessage_Type() {}

func (*KafkaMessage_TimeToCut) isKafkaMessage_Type() {}

func (*KafkaMessage_Connect) isKafkaMessage_Type() {}

func (m *KafkaMessage) GetType() isKafkaMessage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *KafkaMessage) GetRegular() *KafkaMessageRegular {
	if x, ok := m.GetType().(*KafkaMessage_Regular); ok {
		return x.Regular
	}
	return nil
}

func (m *KafkaMessage) GetTimeToCut() *KafkaMessageTimeToCut {
	if x, ok := m.GetType().(*KafkaMessage_TimeToCut); ok {
		return x.TimeToCut
	}
	return nil
}

func (m *KafkaMessage) GetConnect() *KafkaMessageConnect {
	if x, ok := m.GetType().(*KafkaMessage_Connect); ok {
		return x.Connect
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*KafkaMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _KafkaMessage_OneofMarshaler, _KafkaMessage_OneofUnmarshaler, _KafkaMessage_OneofSizer, []interface{}{
		(*KafkaMessage_Regular)(nil),
		(*KafkaMessage_TimeToCut)(nil),
		(*KafkaMessage_Connect)(nil),
	}
}

func _KafkaMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*KafkaMessage)
	// Type
	switch x := m.Type.(type) {
	case *KafkaMessage_Regular:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Regular); err != nil {
			return err
		}
	case *KafkaMessage_TimeToCut:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimeToCut); err != nil {
			return err
		}
	case *KafkaMessage_Connect:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Connect); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("KafkaMessage.Type has unexpected type %T", x)
	}
	return nil
}

func _KafkaMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*KafkaMessage)
	switch tag {
	case 1: // Type.regular
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KafkaMessageRegular)
		err := b.DecodeMessage(msg)
		m.Type = &KafkaMessage_Regular{msg}
		return true, err
	case 2: // Type.time_to_cut
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KafkaMessageTimeToCut)
		err := b.DecodeMessage(msg)
		m.Type = &KafkaMessage_TimeToCut{msg}
		return true, err
	case 3: // Type.connect
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KafkaMessageConnect)
		err := b.DecodeMessage(msg)
		m.Type = &KafkaMessage_Connect{msg}
		return true, err
	default:
		return false, nil
	}
}

func _KafkaMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*KafkaMessage)
	// Type
	switch x := m.Type.(type) {
	case *KafkaMessage_Regular:
		s := proto.Size(x.Regular)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *KafkaMessage_TimeToCut:
		s := proto.Size(x.TimeToCut)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *KafkaMessage_Connect:
		s := proto.Size(x.Connect)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// KafkaMessageRegular wraps a marshalled envelope.
type KafkaMessageRegular struct {
	Payload              []byte                    `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	ConfigSeq            uint64                    `protobuf:"varint,2,opt,name=config_seq,json=configSeq,proto3" json:"config_seq,omitempty"`
	Class                KafkaMessageRegular_Class `protobuf:"varint,3,opt,name=class,proto3,enum=kafka.KafkaMessageRegular_Class" json:"class,omitempty"`
	OriginalOffset       int64                     `protobuf:"varint,4,opt,name=original_offset,json=originalOffset,proto3" json:"original_offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *KafkaMessageRegular) Reset()         { *m = KafkaMessageRegular{} }
func (m *KafkaMessageRegular) String() string { return proto.CompactTextString(m) }
func (*KafkaMessageRegular) ProtoMessage()    {}
func (*KafkaMessageRegular) Descriptor() ([]byte, []int) {
	return fileDescriptor_kafka_8ca1a5cc9386535d, []int{1}
}
func (m *KafkaMessageRegular) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaMessageRegular.Unmarshal(m, b)
}
func (m *KafkaMessageRegular) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaMessageRegular.Marshal(b, m, deterministic)
}
func (dst *KafkaMessageRegular) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaMessageRegular.Merge(dst, src)
}
func (m *KafkaMessageRegular) XXX_Size() int {
	return xxx_messageInfo_KafkaMessageRegular.Size(m)
}
func (m *KafkaMessageRegular) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaMessageRegular.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaMessageRegular proto.InternalMessageInfo

func (m *KafkaMessageRegular) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *KafkaMessageRegular) GetConfigSeq() uint64 {
	if m != nil {
		return m.ConfigSeq
	}
	return 0
}

func (m *KafkaMessageRegular) GetClass() KafkaMessageRegular_Class {
	if m != nil {
		return m.Class
	}
	return KafkaMessageRegular_UNKNOWN
}

func (m *KafkaMessageRegular) GetOriginalOffset() int64 {
	if m != nil {
		return m.OriginalOffset
	}
	return 0
}

// KafkaMessageTimeToCut is used to signal to the orderers
// that it is time to cut block <block_number>.
type KafkaMessageTimeToCut struct {
	BlockNumber          uint64   `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaMessageTimeToCut) Reset()         { *m = KafkaMessageTimeToCut{} }
func (m *KafkaMessageTimeToCut) String() string { return proto.CompactTextString(m) }
func (*KafkaMessageTimeToCut) ProtoMessage()    {}
func (*KafkaMessageTimeToCut) Descriptor() ([]byte, []int) {
	return fileDescriptor_kafka_8ca1a5cc9386535d, []int{2}
}
func (m *KafkaMessageTimeToCut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaMessageTimeToCut.Unmarshal(m, b)
}
func (m *KafkaMessageTimeToCut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaMessageTimeToCut.Marshal(b, m, deterministic)
}
func (dst *KafkaMessageTimeToCut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaMessageTimeToCut.Merge(dst, src)
}
func (m *KafkaMessageTimeToCut) XXX_Size() int {
	return xxx_messageInfo_KafkaMessageTimeToCut.Size(m)
}
func (m *KafkaMessageTimeToCut) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaMessageTimeToCut.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaMessageTimeToCut proto.InternalMessageInfo

func (m *KafkaMessageTimeToCut) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

// KafkaMessageConnect is posted by an orderer upon booting up.
// It is used to prevent the panic that would be caused if we
// were to consume an empty partition. It is ignored by all
// orderers when processing the partition.
type KafkaMessageConnect struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaMessageConnect) Reset()         { *m = KafkaMessageConnect{} }
func (m *KafkaMessageConnect) String() string { return proto.CompactTextString(m) }
func (*KafkaMessageConnect) ProtoMessage()    {}
func (*KafkaMessageConnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_kafka_8ca1a5cc9386535d, []int{3}
}
func (m *KafkaMessageConnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaMessageConnect.Unmarshal(m, b)
}
func (m *KafkaMessageConnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaMessageConnect.Marshal(b, m, deterministic)
}
func (dst *KafkaMessageConnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaMessageConnect.Merge(dst, src)
}
func (m *KafkaMessageConnect) XXX_Size() int {
	return xxx_messageInfo_KafkaMessageConnect.Size(m)
}
func (m *KafkaMessageConnect) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaMessageConnect.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaMessageConnect proto.InternalMessageInfo

func (m *KafkaMessageConnect) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// KafkaMetadata is encoded into the ORDERER block to keep track of
// Kafka-related metadata associated with this block.
type KafkaMetadata struct {
	// LastOffsetPersisted is the encoded value for the Metadata message
	// which is encoded in the ORDERER block metadata index for the case
	// of the Kafka-based orderer.
	LastOffsetPersisted int64 `protobuf:"varint,1,opt,name=last_offset_persisted,json=lastOffsetPersisted,proto3" json:"last_offset_persisted,omitempty"`
	// LastOriginalOffsetProcessed is used to keep track of the newest
	// offset processed if a message is re-validated and re-ordered.
	// This value is used to deduplicate re-submitted messages from
	// multiple orderer so that we don't bother re-processing it again.
	LastOriginalOffsetProcessed int64 `protobuf:"varint,2,opt,name=last_original_offset_processed,json=lastOriginalOffsetProcessed,proto3" json:"last_original_offset_processed,omitempty"`
	// LastResubmittedConfigOffset is used to capture the newest offset of
	// CONFIG kafka message, which is revalidated and resubmitted. By comparing
	// this with LastOriginalOffsetProcessed, we could detemine whether there
	// are still CONFIG messages that have been resubmitted but NOT processed
	// yet. It's used as condition to block ingress messages, so we could reduce
	// the overhead of repeatedly resubmitting messages as config seq keeps
	// advancing.
	LastResubmittedConfigOffset int64           `protobuf:"varint,3,opt,name=last_resubmitted_config_offset,json=lastResubmittedConfigOffset,proto3" json:"last_resubmitted_config_offset,omitempty"`
	ReceivedTTCMessage          bool            `protobuf:"varint,4,opt,name=received_t_t_c_message,json=receivedTTCMessage,proto3" json:"received_t_t_c_message,omitempty"`
	TTCPayload                  *KafkaPayload   `protobuf:"bytes,5,opt,name=t_t_c_payload,json=tTCPayload,proto3" json:"t_t_c_payload,omitempty"`
	ReceivedConnectOrTTCMessage bool            `protobuf:"varint,6,opt,name=received_connect_or_t_t_c_message,json=receivedConnectOrTTCMessage,proto3" json:"received_connect_or_t_t_c_message,omitempty"`
	ConnectOrTTCPayload         []*KafkaPayload `protobuf:"bytes,7,rep,name=connect_or_t_t_c_payload,json=connectOrTTCPayload,proto3" json:"connect_or_t_t_c_payload,omitempty"`
	IsConfigMessage             bool            `protobuf:"varint,8,opt,name=is_config_message,json=isConfigMessage,proto3" json:"is_config_message,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}        `json:"-"`
	XXX_unrecognized            []byte          `json:"-"`
	XXX_sizecache               int32           `json:"-"`
}

func (m *KafkaMetadata) Reset()         { *m = KafkaMetadata{} }
func (m *KafkaMetadata) String() string { return proto.CompactTextString(m) }
func (*KafkaMetadata) ProtoMessage()    {}
func (*KafkaMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_kafka_8ca1a5cc9386535d, []int{4}
}
func (m *KafkaMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaMetadata.Unmarshal(m, b)
}
func (m *KafkaMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaMetadata.Marshal(b, m, deterministic)
}
func (dst *KafkaMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaMetadata.Merge(dst, src)
}
func (m *KafkaMetadata) XXX_Size() int {
	return xxx_messageInfo_KafkaMetadata.Size(m)
}
func (m *KafkaMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaMetadata proto.InternalMessageInfo

func (m *KafkaMetadata) GetLastOffsetPersisted() int64 {
	if m != nil {
		return m.LastOffsetPersisted
	}
	return 0
}

func (m *KafkaMetadata) GetLastOriginalOffsetProcessed() int64 {
	if m != nil {
		return m.LastOriginalOffsetProcessed
	}
	return 0
}

func (m *KafkaMetadata) GetLastResubmittedConfigOffset() int64 {
	if m != nil {
		return m.LastResubmittedConfigOffset
	}
	return 0
}

func (m *KafkaMetadata) GetReceivedTTCMessage() bool {
	if m != nil {
		return m.ReceivedTTCMessage
	}
	return false
}

func (m *KafkaMetadata) GetTTCPayload() *KafkaPayload {
	if m != nil {
		return m.TTCPayload
	}
	return nil
}

func (m *KafkaMetadata) GetReceivedConnectOrTTCMessage() bool {
	if m != nil {
		return m.ReceivedConnectOrTTCMessage
	}
	return false
}

func (m *KafkaMetadata) GetConnectOrTTCPayload() []*KafkaPayload {
	if m != nil {
		return m.ConnectOrTTCPayload
	}
	return nil
}

func (m *KafkaMetadata) GetIsConfigMessage() bool {
	if m != nil {
		return m.IsConfigMessage
	}
	return false
}

type KafkaPayload struct {
	KafkaMerkleProofHeader []byte   `protobuf:"bytes,1,opt,name=kafka_merkle_proof_header,json=kafkaMerkleProofHeader,proto3" json:"kafka_merkle_proof_header,omitempty"`
	KafkaSignatureHeader   []byte   `protobuf:"bytes,2,opt,name=kafka_signature_header,json=kafkaSignatureHeader,proto3" json:"kafka_signature_header,omitempty"`
	ConsumerMessageBytes   []byte   `protobuf:"bytes,3,opt,name=consumer_message_bytes,json=consumerMessageBytes,proto3" json:"consumer_message_bytes,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *KafkaPayload) Reset()         { *m = KafkaPayload{} }
func (m *KafkaPayload) String() string { return proto.CompactTextString(m) }
func (*KafkaPayload) ProtoMessage()    {}
func (*KafkaPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_kafka_8ca1a5cc9386535d, []int{5}
}
func (m *KafkaPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaPayload.Unmarshal(m, b)
}
func (m *KafkaPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaPayload.Marshal(b, m, deterministic)
}
func (dst *KafkaPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaPayload.Merge(dst, src)
}
func (m *KafkaPayload) XXX_Size() int {
	return xxx_messageInfo_KafkaPayload.Size(m)
}
func (m *KafkaPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaPayload.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaPayload proto.InternalMessageInfo

func (m *KafkaPayload) GetKafkaMerkleProofHeader() []byte {
	if m != nil {
		return m.KafkaMerkleProofHeader
	}
	return nil
}

func (m *KafkaPayload) GetKafkaSignatureHeader() []byte {
	if m != nil {
		return m.KafkaSignatureHeader
	}
	return nil
}

func (m *KafkaPayload) GetConsumerMessageBytes() []byte {
	if m != nil {
		return m.ConsumerMessageBytes
	}
	return nil
}

func init() {
	proto.RegisterType((*KafkaMessage)(nil), "kafka.KafkaMessage")
	proto.RegisterType((*KafkaMessageRegular)(nil), "kafka.KafkaMessageRegular")
	proto.RegisterType((*KafkaMessageTimeToCut)(nil), "kafka.KafkaMessageTimeToCut")
	proto.RegisterType((*KafkaMessageConnect)(nil), "kafka.KafkaMessageConnect")
	proto.RegisterType((*KafkaMetadata)(nil), "kafka.KafkaMetadata")
	proto.RegisterType((*KafkaPayload)(nil), "kafka.KafkaPayload")
	proto.RegisterEnum("kafka.KafkaMessageRegular_Class", KafkaMessageRegular_Class_name, KafkaMessageRegular_Class_value)
}

func init() {
	proto.RegisterFile("github.com/hyperledger/fabric_judge/protos/kafka/kafka.proto", fileDescriptor_kafka_8ca1a5cc9386535d)
}

var fileDescriptor_kafka_8ca1a5cc9386535d = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6f, 0xd3, 0x3c,
	0x18, 0x5d, 0xd7, 0xaf, 0xcd, 0xed, 0x3e, 0x5e, 0x77, 0x9b, 0xfa, 0xbe, 0x7b, 0x41, 0x5d, 0x6e,
	0x98, 0x10, 0x6a, 0x51, 0x41, 0x93, 0x40, 0x08, 0x89, 0x45, 0x1a, 0x83, 0xb1, 0xb6, 0xf2, 0x8a,
	0x90, 0xb8, 0xb1, 0x9c, 0xe4, 0x69, 0x16, 0x9a, 0xc4, 0x9d, 0xed, 0x20, 0xf5, 0x87, 0xf1, 0x17,
	0xf8, 0x17, 0xfc, 0x10, 0xee, 0x50, 0xec, 0x78, 0xed, 0x60, 0xbd, 0xe0, 0xa6, 0xb2, 0x9f, 0x73,
	0x8e, 0x9f, 0xe3, 0xe3, 0xa7, 0x41, 0xaf, 0xc2, 0x48, 0x5d, 0x67, 0x5e, 0xd7, 0xe7, 0x49, 0xef,
	0x7a, 0x3e, 0x03, 0x11, 0x43, 0x10, 0x82, 0xe8, 0x4d, 0x98, 0x27, 0x22, 0x9f, 0x7e, 0xc9, 0x82,
	0x10, 0x7a, 0x33, 0xc1, 0x15, 0x97, 0xbd, 0x29, 0x9b, 0x4c, 0x99, 0xf9, 0xed, 0xea, 0x12, 0xae,
	0xea, 0x8d, 0xf3, 0xbd, 0x84, 0x9a, 0x17, 0xf9, 0xea, 0x12, 0xa4, 0x64, 0x21, 0xe0, 0x13, 0x54,
	0x17, 0x10, 0x66, 0x31, 0x13, 0xed, 0x52, 0xa7, 0x74, 0xdc, 0xe8, 0xff, 0xd7, 0x35, 0xb2, 0x65,
	0x16, 0x31, 0x8c, 0xf3, 0x35, 0x62, 0xc9, 0xf8, 0x35, 0x6a, 0xa8, 0x28, 0x01, 0xaa, 0x38, 0xf5,
	0x33, 0xd5, 0x5e, 0xd7, 0xda, 0xff, 0xef, 0xd1, 0x8e, 0xa3, 0x04, 0xc6, 0xdc, 0xcd, 0xd4, 0xf9,
	0x1a, 0xd9, 0x54, 0x76, 0x93, 0xf7, 0xf5, 0x79, 0x9a, 0x82, 0xaf, 0xda, 0xe5, 0x95, 0x7d, 0x5d,
	0xc3, 0xc8, 0xfb, 0x16, 0xe4, 0xd3, 0x1a, 0xaa, 0x8c, 0xe7, 0x33, 0x70, 0x7e, 0x94, 0x50, 0xeb,
	0x1e, 0x8b, 0xb8, 0x8d, 0xea, 0x33, 0x36, 0x8f, 0x39, 0x0b, 0xf4, 0x7d, 0x9a, 0xc4, 0x6e, 0xf1,
	0x03, 0x84, 0x7c, 0x9e, 0x4e, 0xa2, 0x90, 0x4a, 0xb8, 0xd1, 0x86, 0x2b, 0x64, 0xd3, 0x54, 0xae,
	0xe0, 0x06, 0x9f, 0xa0, 0xaa, 0x1f, 0x33, 0x29, 0xb5, 0x9d, 0xed, 0x7e, 0x67, 0x75, 0x0c, 0x5d,
	0x37, 0xe7, 0x11, 0x43, 0xc7, 0x8f, 0xd0, 0x0e, 0x17, 0x51, 0x18, 0xa5, 0x2c, 0xa6, 0x7c, 0x32,
	0x91, 0xa0, 0xda, 0x95, 0x4e, 0xe9, 0xb8, 0x4c, 0xb6, 0x6d, 0x79, 0xa8, 0xab, 0xce, 0x13, 0x54,
	0xd5, 0x42, 0xdc, 0x40, 0xf5, 0x8f, 0x83, 0x8b, 0xc1, 0xf0, 0xd3, 0x60, 0x77, 0x0d, 0x23, 0x54,
	0x1b, 0x0c, 0xc9, 0xe5, 0x9b, 0x0f, 0xbb, 0xa5, 0x7c, 0xed, 0x0e, 0x07, 0x67, 0xef, 0xde, 0xee,
	0xae, 0x3b, 0x2f, 0xd1, 0xfe, 0xbd, 0x29, 0xe2, 0x23, 0xd4, 0xf4, 0x62, 0xee, 0x4f, 0x69, 0x9a,
	0x25, 0x1e, 0x98, 0x57, 0xab, 0x90, 0x86, 0xae, 0x0d, 0x74, 0xc9, 0xe9, 0xdd, 0x8d, 0xa6, 0x48,
	0x71, 0x75, 0x34, 0xce, 0xcf, 0x32, 0xda, 0x2a, 0x14, 0x8a, 0x05, 0x4c, 0x31, 0xdc, 0x47, 0xfb,
	0x31, 0x93, 0xaa, 0xb8, 0x11, 0x9d, 0x81, 0x90, 0x91, 0x54, 0x60, 0x94, 0x65, 0xd2, 0xca, 0x41,
	0x73, 0xaf, 0x91, 0x85, 0xb0, 0x8b, 0x1e, 0x1a, 0xcd, 0xdd, 0x38, 0xe8, 0x4c, 0x70, 0x1f, 0xa4,
	0x84, 0x40, 0x87, 0x5e, 0x26, 0x87, 0x5a, 0x7c, 0x27, 0x9c, 0x91, 0xa5, 0xdc, 0x1e, 0x22, 0x40,
	0x66, 0x5e, 0x12, 0x29, 0x05, 0x01, 0x2d, 0x9e, 0xad, 0x48, 0xb7, 0xbc, 0x38, 0x84, 0x2c, 0x48,
	0xae, 0xe6, 0x98, 0xd3, 0x70, 0x1f, 0x1d, 0x08, 0xf0, 0x21, 0xfa, 0x0a, 0x01, 0x55, 0x54, 0x51,
	0x9f, 0x26, 0x26, 0x0a, 0xfd, 0x34, 0x1b, 0x04, 0x5b, 0x74, 0x3c, 0x76, 0x17, 0x7f, 0x84, 0x2d,
	0x43, 0xb5, 0x19, 0x55, 0xf5, 0x58, 0xb6, 0x96, 0xe7, 0x60, 0x64, 0x20, 0x82, 0xd4, 0xd8, 0x2d,
	0xd6, 0xf8, 0x0c, 0x1d, 0xdd, 0xf6, 0x2a, 0x86, 0x94, 0x72, 0xf1, 0x5b, 0xdb, 0x9a, 0x6e, 0x7b,
	0x68, 0x89, 0xc5, 0x8b, 0x0c, 0xc5, 0x52, 0xff, 0xf7, 0xa8, 0xfd, 0x87, 0xdc, 0x5a, 0xa9, 0x77,
	0xca, 0xab, 0xac, 0xb4, 0xfc, 0xa5, 0xb3, 0xac, 0xa7, 0xc7, 0xe8, 0x9f, 0x48, 0xda, 0xd8, 0xac,
	0x87, 0x0d, 0xed, 0x61, 0x27, 0x92, 0x26, 0xaa, 0xa2, 0xaf, 0xf3, 0xcd, 0x7e, 0x11, 0xac, 0xf8,
	0x05, 0xfa, 0x57, 0xf7, 0xa1, 0x09, 0x88, 0x69, 0x0c, 0xf9, 0xf3, 0xf1, 0x09, 0xbd, 0x06, 0x16,
	0x14, 0xd3, 0xd6, 0x24, 0x07, 0x53, 0x33, 0x2c, 0x39, 0x3e, 0xca, 0xe1, 0x73, 0x8d, 0xe2, 0xe7,
	0xc8, 0x20, 0x54, 0x46, 0x61, 0xca, 0x54, 0x26, 0xc0, 0xea, 0xd6, 0xb5, 0x6e, 0x4f, 0xa3, 0x57,
	0x16, 0x5c, 0xa8, 0x7c, 0x9e, 0xca, 0x2c, 0x01, 0x61, 0xcd, 0x52, 0x6f, 0xae, 0xc0, 0xfc, 0x15,
	0x9b, 0x64, 0xcf, 0xa2, 0x85, 0xe5, 0xd3, 0x1c, 0x3b, 0xed, 0x7f, 0x7e, 0xfa, 0xb7, 0x1f, 0x44,
	0xaf, 0xa6, 0x77, 0xcf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x4e, 0xb8, 0x3e, 0x4b, 0x05,
	0x00, 0x00,
}
