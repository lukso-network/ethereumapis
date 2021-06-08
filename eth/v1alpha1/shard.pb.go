// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eth/v1alpha1/shard.proto

package eth

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ShardState struct {
	Slot                 uint64   `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	GasPrice             uint64   `protobuf:"varint,2,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	LatestBlockRoot      []byte   `protobuf:"bytes,3,opt,name=latest_block_root,json=latestBlockRoot,proto3" json:"latest_block_root,omitempty" ssz-size:"32"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShardState) Reset()         { *m = ShardState{} }
func (m *ShardState) String() string { return proto.CompactTextString(m) }
func (*ShardState) ProtoMessage()    {}
func (*ShardState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d9da510e571288b, []int{0}
}
func (m *ShardState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShardState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShardState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShardState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardState.Merge(m, src)
}
func (m *ShardState) XXX_Size() int {
	return m.Size()
}
func (m *ShardState) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardState.DiscardUnknown(m)
}

var xxx_messageInfo_ShardState proto.InternalMessageInfo

func (m *ShardState) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *ShardState) GetGasPrice() uint64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *ShardState) GetLatestBlockRoot() []byte {
	if m != nil {
		return m.LatestBlockRoot
	}
	return nil
}

type ShardTransition struct {
	StartSlot                  uint64               `protobuf:"varint,1,opt,name=start_slot,json=startSlot,proto3" json:"start_slot,omitempty"`
	ShardBlockLengths          []uint64             `protobuf:"varint,2,rep,packed,name=shard_block_lengths,json=shardBlockLengths,proto3" json:"shard_block_lengths,omitempty" ssz-max:"12"`
	ShardDataRoots             [][]byte             `protobuf:"bytes,3,rep,name=shard_data_roots,json=shardDataRoots,proto3" json:"shard_data_roots,omitempty" ssz-size:"?,32" ssz-max:"12"`
	ShardStates                []*ShardState        `protobuf:"bytes,4,rep,name=shard_states,json=shardStates,proto3" json:"shard_states,omitempty" ssz-max:"12"`
	ProposerSignatureAggregate []byte               `protobuf:"bytes,5,opt,name=proposer_signature_aggregate,json=proposerSignatureAggregate,proto3" json:"proposer_signature_aggregate,omitempty" ssz-size:"96"`
	PandoraStates              []*PandoraShardState `protobuf:"bytes,6,rep,name=pandora_states,json=pandoraStates,proto3" json:"pandora_states,omitempty" ssz-max:"12"`
	XXX_NoUnkeyedLiteral       struct{}             `json:"-"`
	XXX_unrecognized           []byte               `json:"-"`
	XXX_sizecache              int32                `json:"-"`
}

func (m *ShardTransition) Reset()         { *m = ShardTransition{} }
func (m *ShardTransition) String() string { return proto.CompactTextString(m) }
func (*ShardTransition) ProtoMessage()    {}
func (*ShardTransition) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d9da510e571288b, []int{1}
}
func (m *ShardTransition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShardTransition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShardTransition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShardTransition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardTransition.Merge(m, src)
}
func (m *ShardTransition) XXX_Size() int {
	return m.Size()
}
func (m *ShardTransition) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardTransition.DiscardUnknown(m)
}

var xxx_messageInfo_ShardTransition proto.InternalMessageInfo

func (m *ShardTransition) GetStartSlot() uint64 {
	if m != nil {
		return m.StartSlot
	}
	return 0
}

func (m *ShardTransition) GetShardBlockLengths() []uint64 {
	if m != nil {
		return m.ShardBlockLengths
	}
	return nil
}

func (m *ShardTransition) GetShardDataRoots() [][]byte {
	if m != nil {
		return m.ShardDataRoots
	}
	return nil
}

func (m *ShardTransition) GetShardStates() []*ShardState {
	if m != nil {
		return m.ShardStates
	}
	return nil
}

func (m *ShardTransition) GetProposerSignatureAggregate() []byte {
	if m != nil {
		return m.ProposerSignatureAggregate
	}
	return nil
}

func (m *ShardTransition) GetPandoraStates() []*PandoraShardState {
	if m != nil {
		return m.PandoraStates
	}
	return nil
}

func init() {
	proto.RegisterType((*ShardState)(nil), "ethereum.eth.v1alpha1.ShardState")
	proto.RegisterType((*ShardTransition)(nil), "ethereum.eth.v1alpha1.ShardTransition")
}

func init() { proto.RegisterFile("eth/v1alpha1/shard.proto", fileDescriptor_5d9da510e571288b) }

var fileDescriptor_5d9da510e571288b = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0xc9, 0xa6, 0x2e, 0xee, 0x6c, 0xf7, 0x4f, 0x47, 0x16, 0x42, 0x5d, 0xdb, 0x98, 0x0b,
	0xc9, 0x85, 0x9b, 0xd2, 0x16, 0x16, 0x5c, 0x11, 0xb5, 0xb8, 0x37, 0xe2, 0xc5, 0xd2, 0x8a, 0x17,
	0x22, 0x84, 0xd3, 0x76, 0x9c, 0x04, 0xd3, 0x4c, 0x98, 0x73, 0x2a, 0xba, 0xf8, 0x44, 0x3e, 0x89,
	0x57, 0xe2, 0x13, 0x14, 0xe9, 0x23, 0xf4, 0x09, 0x24, 0x33, 0xa6, 0xdb, 0x85, 0xea, 0xdd, 0x9c,
	0x99, 0xef, 0x7c, 0xdf, 0x2f, 0x27, 0x87, 0x79, 0x82, 0x92, 0xce, 0xe7, 0x2e, 0x64, 0x45, 0x02,
	0xdd, 0x0e, 0x26, 0xa0, 0xa7, 0x51, 0xa1, 0x15, 0x29, 0x7e, 0x22, 0x28, 0x11, 0x5a, 0xcc, 0x67,
	0x91, 0xa0, 0x24, 0xaa, 0x24, 0xcd, 0x33, 0x99, 0x52, 0x32, 0x1f, 0x47, 0x13, 0x35, 0xeb, 0x48,
	0x25, 0x55, 0xc7, 0xa8, 0xc7, 0xf3, 0x8f, 0xa6, 0x32, 0x85, 0x39, 0x59, 0x97, 0xe6, 0xa3, 0x5b,
	0xfe, 0x05, 0xe4, 0x53, 0xa5, 0x21, 0x36, 0x39, 0x31, 0x12, 0x90, 0xb0, 0xba, 0xe0, 0x1b, 0x63,
	0xa3, 0xf2, 0x72, 0x54, 0xde, 0x71, 0xce, 0x6a, 0x98, 0x29, 0xf2, 0x1c, 0xdf, 0x09, 0x6b, 0x43,
	0x73, 0xe6, 0xf7, 0xd9, 0x9e, 0x04, 0x8c, 0x0b, 0x9d, 0x4e, 0x84, 0xb7, 0x63, 0x1e, 0xee, 0x4a,
	0xc0, 0xab, 0xb2, 0xe6, 0xcf, 0x58, 0x23, 0x03, 0x12, 0x48, 0xf1, 0x38, 0x53, 0x93, 0x4f, 0xb1,
	0x56, 0x8a, 0x3c, 0xd7, 0x77, 0xc2, 0xfa, 0xa0, 0xb1, 0x5a, 0xb4, 0x0f, 0x10, 0xaf, 0xcf, 0x30,
	0xbd, 0x16, 0x17, 0x41, 0xbf, 0x17, 0x0c, 0x8f, 0xac, 0x76, 0x50, 0x4a, 0x87, 0x4a, 0x51, 0xf0,
	0xd3, 0x65, 0x47, 0x26, 0xfe, 0xad, 0x86, 0x1c, 0x53, 0x4a, 0x55, 0xce, 0x1f, 0x30, 0x86, 0x04,
	0x9a, 0xe2, 0x0d, 0x92, 0x3d, 0x73, 0x33, 0x2a, 0x71, 0x5e, 0xb0, 0x7b, 0xf6, 0x2b, 0x6c, 0x60,
	0x26, 0x72, 0x49, 0x09, 0x7a, 0x3b, 0xbe, 0x1b, 0xd6, 0x06, 0xc7, 0xab, 0x45, 0xbb, 0x5e, 0x66,
	0xce, 0xe0, 0xcb, 0x45, 0xd0, 0xed, 0x05, 0xc3, 0x86, 0x11, 0x9b, 0xc4, 0x37, 0x56, 0xca, 0x5f,
	0xb3, 0x63, 0xeb, 0x30, 0x05, 0x02, 0x43, 0x8c, 0x9e, 0xeb, 0xbb, 0x61, 0x7d, 0xe0, 0xaf, 0x16,
	0xed, 0xd3, 0x1b, 0xe4, 0xe7, 0x8f, 0xfb, 0xbd, 0xc0, 0xbf, 0x65, 0x77, 0x68, 0x3a, 0x5f, 0x01,
	0x41, 0xc9, 0x8f, 0xfc, 0x1d, 0xab, 0x6f, 0xcc, 0x14, 0xbd, 0x9a, 0xef, 0x86, 0xfb, 0xbd, 0x87,
	0xd1, 0xd6, 0x7f, 0x18, 0xdd, 0x4c, 0x7a, 0x0b, 0xe9, 0x3e, 0xae, 0x5f, 0x91, 0x8f, 0xd8, 0x69,
	0xa1, 0x55, 0xa1, 0x50, 0xe8, 0x18, 0x53, 0x99, 0x03, 0xcd, 0xb5, 0x88, 0x41, 0x4a, 0x2d, 0x24,
	0x90, 0xf0, 0xee, 0x6c, 0x1b, 0xf1, 0x93, 0xf3, 0x60, 0xd8, 0xac, 0xda, 0x46, 0x55, 0xd7, 0xcb,
	0xaa, 0x89, 0x03, 0x3b, 0x5c, 0x2f, 0x82, 0xc5, 0xdd, 0x35, 0xb8, 0xe1, 0x3f, 0x70, 0xaf, 0xac,
	0xf8, 0xbf, 0xd4, 0x07, 0x7f, 0x1d, 0x2d, 0xf7, 0xe0, 0xc3, 0x8f, 0x65, 0xcb, 0xf9, 0xb5, 0x6c,
	0x39, 0xbf, 0x97, 0x2d, 0xe7, 0xfd, 0xf9, 0xc6, 0xce, 0x16, 0xfa, 0x2b, 0xce, 0x80, 0xd2, 0x49,
	0x06, 0x63, 0xec, 0x54, 0x81, 0x50, 0xa4, 0xa6, 0x58, 0xaf, 0xea, 0x53, 0x41, 0xc9, 0xf7, 0x9d,
	0x93, 0xcb, 0x0a, 0xe8, 0x72, 0x03, 0x68, 0xbc, 0x6b, 0x76, 0xb6, 0xff, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x6e, 0x6a, 0xeb, 0x2c, 0x3d, 0x03, 0x00, 0x00,
}

func (m *ShardState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShardState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShardState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.LatestBlockRoot) > 0 {
		i -= len(m.LatestBlockRoot)
		copy(dAtA[i:], m.LatestBlockRoot)
		i = encodeVarintShard(dAtA, i, uint64(len(m.LatestBlockRoot)))
		i--
		dAtA[i] = 0x1a
	}
	if m.GasPrice != 0 {
		i = encodeVarintShard(dAtA, i, uint64(m.GasPrice))
		i--
		dAtA[i] = 0x10
	}
	if m.Slot != 0 {
		i = encodeVarintShard(dAtA, i, uint64(m.Slot))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ShardTransition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShardTransition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShardTransition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PandoraStates) > 0 {
		for iNdEx := len(m.PandoraStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PandoraStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintShard(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ProposerSignatureAggregate) > 0 {
		i -= len(m.ProposerSignatureAggregate)
		copy(dAtA[i:], m.ProposerSignatureAggregate)
		i = encodeVarintShard(dAtA, i, uint64(len(m.ProposerSignatureAggregate)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ShardStates) > 0 {
		for iNdEx := len(m.ShardStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ShardStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintShard(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ShardDataRoots) > 0 {
		for iNdEx := len(m.ShardDataRoots) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ShardDataRoots[iNdEx])
			copy(dAtA[i:], m.ShardDataRoots[iNdEx])
			i = encodeVarintShard(dAtA, i, uint64(len(m.ShardDataRoots[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ShardBlockLengths) > 0 {
		dAtA2 := make([]byte, len(m.ShardBlockLengths)*10)
		var j1 int
		for _, num := range m.ShardBlockLengths {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintShard(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if m.StartSlot != 0 {
		i = encodeVarintShard(dAtA, i, uint64(m.StartSlot))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintShard(dAtA []byte, offset int, v uint64) int {
	offset -= sovShard(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ShardState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Slot != 0 {
		n += 1 + sovShard(uint64(m.Slot))
	}
	if m.GasPrice != 0 {
		n += 1 + sovShard(uint64(m.GasPrice))
	}
	l = len(m.LatestBlockRoot)
	if l > 0 {
		n += 1 + l + sovShard(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ShardTransition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartSlot != 0 {
		n += 1 + sovShard(uint64(m.StartSlot))
	}
	if len(m.ShardBlockLengths) > 0 {
		l = 0
		for _, e := range m.ShardBlockLengths {
			l += sovShard(uint64(e))
		}
		n += 1 + sovShard(uint64(l)) + l
	}
	if len(m.ShardDataRoots) > 0 {
		for _, b := range m.ShardDataRoots {
			l = len(b)
			n += 1 + l + sovShard(uint64(l))
		}
	}
	if len(m.ShardStates) > 0 {
		for _, e := range m.ShardStates {
			l = e.Size()
			n += 1 + l + sovShard(uint64(l))
		}
	}
	l = len(m.ProposerSignatureAggregate)
	if l > 0 {
		n += 1 + l + sovShard(uint64(l))
	}
	if len(m.PandoraStates) > 0 {
		for _, e := range m.PandoraStates {
			l = e.Size()
			n += 1 + l + sovShard(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovShard(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShard(x uint64) (n int) {
	return sovShard(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ShardState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShard
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ShardState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShardState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slot", wireType)
			}
			m.Slot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Slot |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			m.GasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlockRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthShard
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LatestBlockRoot = append(m.LatestBlockRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.LatestBlockRoot == nil {
				m.LatestBlockRoot = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShard(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShard
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShard
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ShardTransition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShard
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ShardTransition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShardTransition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartSlot", wireType)
			}
			m.StartSlot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartSlot |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowShard
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ShardBlockLengths = append(m.ShardBlockLengths, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowShard
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthShard
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthShard
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ShardBlockLengths) == 0 {
					m.ShardBlockLengths = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowShard
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ShardBlockLengths = append(m.ShardBlockLengths, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardBlockLengths", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardDataRoots", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthShard
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShardDataRoots = append(m.ShardDataRoots, make([]byte, postIndex-iNdEx))
			copy(m.ShardDataRoots[len(m.ShardDataRoots)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthShard
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShardStates = append(m.ShardStates, &ShardState{})
			if err := m.ShardStates[len(m.ShardStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerSignatureAggregate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthShard
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposerSignatureAggregate = append(m.ProposerSignatureAggregate[:0], dAtA[iNdEx:postIndex]...)
			if m.ProposerSignatureAggregate == nil {
				m.ProposerSignatureAggregate = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PandoraStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShard
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthShard
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShard
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PandoraStates = append(m.PandoraStates, &PandoraShardState{})
			if err := m.PandoraStates[len(m.PandoraStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShard(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShard
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShard
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipShard(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShard
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowShard
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowShard
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthShard
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShard
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShard
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShard        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShard          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShard = fmt.Errorf("proto: unexpected end of group")
)
