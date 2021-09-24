// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/mutation_record.proto

package monitoring

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Describes a change made to a configuration.
type MutationRecord struct {
	// When the change occurred.
	MutateTime *google_protobuf2.Timestamp `protobuf:"bytes,1,opt,name=mutate_time,json=mutateTime" json:"mutate_time,omitempty"`
	// The email address of the user making the change.
	MutatedBy string `protobuf:"bytes,2,opt,name=mutated_by,json=mutatedBy" json:"mutated_by,omitempty"`
}

func (m *MutationRecord) Reset()                    { *m = MutationRecord{} }
func (m *MutationRecord) String() string            { return proto.CompactTextString(m) }
func (*MutationRecord) ProtoMessage()               {}
func (*MutationRecord) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *MutationRecord) GetMutateTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.MutateTime
	}
	return nil
}

func (m *MutationRecord) GetMutatedBy() string {
	if m != nil {
		return m.MutatedBy
	}
	return ""
}

func init() {
	proto.RegisterType((*MutationRecord)(nil), "google.monitoring.v3.MutationRecord")
}

func init() { proto.RegisterFile("google/monitoring/v3/mutation_record.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0xcf, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9, 0x2f, 0xca, 0xcc, 0x4b, 0xd7, 0x2f, 0x33,
	0xd6, 0xcf, 0x2d, 0x2d, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b, 0x2f, 0x4a, 0x4d, 0xce, 0x2f, 0x4a,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x81, 0xa8, 0xd5, 0x43, 0xa8, 0xd5, 0x2b, 0x33,
	0x96, 0x92, 0x87, 0x9a, 0x00, 0x56, 0x93, 0x54, 0x9a, 0xa6, 0x5f, 0x92, 0x99, 0x9b, 0x5a, 0x5c,
	0x92, 0x98, 0x5b, 0x00, 0xd1, 0xa6, 0x94, 0xc3, 0xc5, 0xe7, 0x0b, 0x35, 0x2f, 0x08, 0x6c, 0x9c,
	0x90, 0x35, 0x17, 0x37, 0xd8, 0x86, 0xd4, 0x78, 0x90, 0x5a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e,
	0x23, 0x29, 0x3d, 0xa8, 0xf1, 0x30, 0x83, 0xf4, 0x42, 0x60, 0x06, 0x05, 0x71, 0x41, 0x94, 0x83,
	0x04, 0x84, 0x64, 0xb9, 0xa0, 0xbc, 0x94, 0xf8, 0xa4, 0x4a, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x4e, 0xa8, 0x88, 0x53, 0xa5, 0xd3, 0x6a, 0x46, 0x2e, 0x89, 0xe4, 0xfc, 0x5c, 0x3d, 0x6c,
	0x6e, 0x75, 0x12, 0x46, 0x75, 0x48, 0x00, 0xc8, 0xa6, 0x00, 0xc6, 0x28, 0x3b, 0xa8, 0xe2, 0xf4,
	0xfc, 0x9c, 0xc4, 0xbc, 0x74, 0xbd, 0xfc, 0xa2, 0x74, 0xfd, 0xf4, 0xd4, 0x3c, 0xb0, 0x3b, 0xf4,
	0x21, 0x52, 0x89, 0x05, 0x99, 0xc5, 0xa8, 0x61, 0x64, 0x8d, 0xe0, 0xad, 0x62, 0x92, 0x72, 0x87,
	0x18, 0xe0, 0x9c, 0x93, 0x5f, 0x9a, 0xa2, 0xe7, 0x8b, 0xb0, 0x33, 0xcc, 0xf8, 0x14, 0x4c, 0x32,
	0x06, 0x2c, 0x19, 0x83, 0x90, 0x8c, 0x09, 0x33, 0x4e, 0x62, 0x03, 0x5b, 0x62, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x95, 0xa7, 0xf3, 0xbd, 0x87, 0x01, 0x00, 0x00,
}
