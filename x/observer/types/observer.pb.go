// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/observer/observer.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "github.com/zeta-chain/zetacore/common"
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

type ObservationType int32

const (
	ObservationType_EmptyObserverType ObservationType = 0
	ObservationType_InBoundTx         ObservationType = 1
	ObservationType_OutBoundTx        ObservationType = 2
	ObservationType_TSSKeyGen         ObservationType = 3
	ObservationType_TSSKeySign        ObservationType = 4
)

var ObservationType_name = map[int32]string{
	0: "EmptyObserverType",
	1: "InBoundTx",
	2: "OutBoundTx",
	3: "TSSKeyGen",
	4: "TSSKeySign",
}

var ObservationType_value = map[string]int32{
	"EmptyObserverType": 0,
	"InBoundTx":         1,
	"OutBoundTx":        2,
	"TSSKeyGen":         3,
	"TSSKeySign":        4,
}

func (x ObservationType) String() string {
	return proto.EnumName(ObservationType_name, int32(x))
}

func (ObservationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_12bb32b20266cc5f, []int{0}
}

type ObserverUpdateReason int32

const (
	ObserverUpdateReason_Undefined   ObserverUpdateReason = 0
	ObserverUpdateReason_Tombstoned  ObserverUpdateReason = 1
	ObserverUpdateReason_AdminUpdate ObserverUpdateReason = 2
)

var ObserverUpdateReason_name = map[int32]string{
	0: "Undefined",
	1: "Tombstoned",
	2: "AdminUpdate",
}

var ObserverUpdateReason_value = map[string]int32{
	"Undefined":   0,
	"Tombstoned":  1,
	"AdminUpdate": 2,
}

func (x ObserverUpdateReason) String() string {
	return proto.EnumName(ObserverUpdateReason_name, int32(x))
}

func (ObserverUpdateReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_12bb32b20266cc5f, []int{1}
}

type ObserverMapper struct {
	Index         string        `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	ObserverChain *common.Chain `protobuf:"bytes,2,opt,name=observer_chain,json=observerChain,proto3" json:"observer_chain,omitempty"`
	ObserverList  []string      `protobuf:"bytes,4,rep,name=observer_list,json=observerList,proto3" json:"observer_list,omitempty"`
}

func (m *ObserverMapper) Reset()         { *m = ObserverMapper{} }
func (m *ObserverMapper) String() string { return proto.CompactTextString(m) }
func (*ObserverMapper) ProtoMessage()    {}
func (*ObserverMapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_12bb32b20266cc5f, []int{0}
}
func (m *ObserverMapper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObserverMapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObserverMapper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObserverMapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObserverMapper.Merge(m, src)
}
func (m *ObserverMapper) XXX_Size() int {
	return m.Size()
}
func (m *ObserverMapper) XXX_DiscardUnknown() {
	xxx_messageInfo_ObserverMapper.DiscardUnknown(m)
}

var xxx_messageInfo_ObserverMapper proto.InternalMessageInfo

func (m *ObserverMapper) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ObserverMapper) GetObserverChain() *common.Chain {
	if m != nil {
		return m.ObserverChain
	}
	return nil
}

func (m *ObserverMapper) GetObserverList() []string {
	if m != nil {
		return m.ObserverList
	}
	return nil
}

type ObserverSet struct {
	ObserverList []string `protobuf:"bytes,1,rep,name=observer_list,json=observerList,proto3" json:"observer_list,omitempty"`
}

func (m *ObserverSet) Reset()         { *m = ObserverSet{} }
func (m *ObserverSet) String() string { return proto.CompactTextString(m) }
func (*ObserverSet) ProtoMessage()    {}
func (*ObserverSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_12bb32b20266cc5f, []int{1}
}
func (m *ObserverSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObserverSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObserverSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObserverSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObserverSet.Merge(m, src)
}
func (m *ObserverSet) XXX_Size() int {
	return m.Size()
}
func (m *ObserverSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ObserverSet.DiscardUnknown(m)
}

var xxx_messageInfo_ObserverSet proto.InternalMessageInfo

func (m *ObserverSet) GetObserverList() []string {
	if m != nil {
		return m.ObserverList
	}
	return nil
}

type LastObserverCount struct {
	Count            uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	LastChangeHeight int64  `protobuf:"varint,2,opt,name=last_change_height,json=lastChangeHeight,proto3" json:"last_change_height,omitempty"`
}

func (m *LastObserverCount) Reset()         { *m = LastObserverCount{} }
func (m *LastObserverCount) String() string { return proto.CompactTextString(m) }
func (*LastObserverCount) ProtoMessage()    {}
func (*LastObserverCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_12bb32b20266cc5f, []int{2}
}
func (m *LastObserverCount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastObserverCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastObserverCount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastObserverCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastObserverCount.Merge(m, src)
}
func (m *LastObserverCount) XXX_Size() int {
	return m.Size()
}
func (m *LastObserverCount) XXX_DiscardUnknown() {
	xxx_messageInfo_LastObserverCount.DiscardUnknown(m)
}

var xxx_messageInfo_LastObserverCount proto.InternalMessageInfo

func (m *LastObserverCount) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *LastObserverCount) GetLastChangeHeight() int64 {
	if m != nil {
		return m.LastChangeHeight
	}
	return 0
}

func init() {
	proto.RegisterEnum("zetachain.zetacore.observer.ObservationType", ObservationType_name, ObservationType_value)
	proto.RegisterEnum("zetachain.zetacore.observer.ObserverUpdateReason", ObserverUpdateReason_name, ObserverUpdateReason_value)
	proto.RegisterType((*ObserverMapper)(nil), "zetachain.zetacore.observer.ObserverMapper")
	proto.RegisterType((*ObserverSet)(nil), "zetachain.zetacore.observer.ObserverSet")
	proto.RegisterType((*LastObserverCount)(nil), "zetachain.zetacore.observer.LastObserverCount")
}

func init() { proto.RegisterFile("zetachain/observer/observer.proto", fileDescriptor_12bb32b20266cc5f) }

var fileDescriptor_12bb32b20266cc5f = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x8a, 0x13, 0x41,
	0x14, 0x9c, 0x4e, 0xa2, 0x90, 0x8e, 0xc9, 0xce, 0x36, 0x11, 0x42, 0xc4, 0x21, 0xae, 0x97, 0xb0,
	0x68, 0x06, 0x56, 0x7f, 0xc0, 0x0d, 0xa2, 0x8b, 0x91, 0x85, 0x49, 0x16, 0xc1, 0x4b, 0xe8, 0x64,
	0x9e, 0x93, 0x86, 0x4c, 0xf7, 0x30, 0xfd, 0x22, 0x19, 0x6f, 0xfe, 0x81, 0x1f, 0xe1, 0xc1, 0x4f,
	0xf1, 0xb8, 0x47, 0x8f, 0x92, 0xfc, 0x88, 0x74, 0xf7, 0x76, 0x72, 0x70, 0x4f, 0x53, 0x55, 0xaf,
	0xeb, 0xbd, 0x82, 0x29, 0xfa, 0xec, 0x1b, 0x20, 0x5f, 0xae, 0xb8, 0x90, 0xb1, 0x5a, 0x68, 0x28,
	0xbf, 0x42, 0x79, 0x00, 0xa3, 0xa2, 0x54, 0xa8, 0xd8, 0x93, 0xc3, 0x93, 0x91, 0x45, 0xaa, 0x84,
	0x91, 0x7f, 0xd2, 0x7f, 0x7a, 0xf4, 0x2f, 0x55, 0x9e, 0x2b, 0xff, 0x71, 0xde, 0x7e, 0x37, 0x53,
	0x99, 0xb2, 0x30, 0x36, 0xc8, 0xa9, 0x67, 0xdf, 0x09, 0xed, 0x5c, 0xdf, 0x6d, 0xf8, 0xc8, 0x8b,
	0x02, 0x4a, 0xd6, 0xa5, 0x0f, 0x84, 0x4c, 0x61, 0xdb, 0x23, 0x03, 0x32, 0x6c, 0x26, 0x8e, 0xb0,
	0xd7, 0xb4, 0xe3, 0x2f, 0xcd, 0xed, 0x91, 0x5e, 0x6d, 0x40, 0x86, 0xad, 0x8b, 0xf6, 0xe8, 0xee,
	0xca, 0xd8, 0x88, 0x49, 0xdb, 0x3f, 0xb2, 0x94, 0x3d, 0xa7, 0x07, 0x61, 0xbe, 0x16, 0x1a, 0x7b,
	0x8d, 0x41, 0x7d, 0xd8, 0x4c, 0x1e, 0x79, 0x71, 0x22, 0x34, 0x9e, 0x5d, 0xd0, 0x96, 0x8f, 0x30,
	0x05, 0xfc, 0xdf, 0x43, 0xee, 0xf1, 0x7c, 0xa2, 0xa7, 0x13, 0xae, 0xd1, 0xfb, 0xc6, 0x6a, 0x23,
	0xd1, 0x24, 0x5f, 0x1a, 0x60, 0x93, 0x37, 0x12, 0x47, 0xd8, 0x0b, 0xca, 0xd6, 0x5c, 0xa3, 0x49,
	0x2d, 0x33, 0x98, 0xaf, 0x40, 0x64, 0x2b, 0xb4, 0xe9, 0xeb, 0x49, 0x68, 0x26, 0x63, 0x3b, 0x78,
	0x6f, 0xf5, 0xf3, 0x35, 0x3d, 0x71, 0x4b, 0x39, 0x0a, 0x25, 0x67, 0x55, 0x01, 0xec, 0x31, 0x3d,
	0x7d, 0x9b, 0x17, 0x58, 0xf9, 0x63, 0x46, 0x0c, 0x03, 0xd6, 0xa6, 0xcd, 0x2b, 0x79, 0xa9, 0x36,
	0x32, 0x9d, 0x6d, 0x43, 0xc2, 0x3a, 0x94, 0x5e, 0x6f, 0xd0, 0xf3, 0x9a, 0x19, 0xcf, 0xa6, 0xd3,
	0x0f, 0x50, 0xbd, 0x03, 0x19, 0xd6, 0xcd, 0xd8, 0xd1, 0xa9, 0xc8, 0x64, 0xd8, 0xe8, 0x37, 0x7e,
	0xfd, 0x8c, 0xc8, 0xf9, 0x84, 0x76, 0xfd, 0xd6, 0x9b, 0x22, 0xe5, 0x08, 0x09, 0x70, 0xad, 0xa4,
	0x31, 0xdf, 0xc8, 0x14, 0xbe, 0x08, 0x09, 0x69, 0x18, 0x58, 0xb3, 0xca, 0x17, 0x1a, 0x95, 0xe1,
	0x84, 0x9d, 0xd0, 0xd6, 0x9b, 0x34, 0x17, 0xd2, 0x79, 0xc2, 0x9a, 0xdb, 0x76, 0x79, 0xf5, 0x7b,
	0x17, 0x91, 0xdb, 0x5d, 0x44, 0xfe, 0xee, 0x22, 0xf2, 0x63, 0x1f, 0x05, 0xb7, 0xfb, 0x28, 0xf8,
	0xb3, 0x8f, 0x82, 0xcf, 0x71, 0x26, 0x70, 0xb5, 0x59, 0x98, 0x7f, 0x15, 0x9b, 0x9a, 0xbc, 0x74,
	0x3d, 0xf1, 0x25, 0x8a, 0xb7, 0xc7, 0xca, 0x61, 0x55, 0x80, 0x5e, 0x3c, 0xb4, 0xf5, 0x78, 0xf5,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x19, 0xd5, 0xdd, 0x42, 0x95, 0x02, 0x00, 0x00,
}

func (m *ObserverMapper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObserverMapper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObserverMapper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ObserverList) > 0 {
		for iNdEx := len(m.ObserverList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ObserverList[iNdEx])
			copy(dAtA[i:], m.ObserverList[iNdEx])
			i = encodeVarintObserver(dAtA, i, uint64(len(m.ObserverList[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.ObserverChain != nil {
		{
			size, err := m.ObserverChain.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObserver(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintObserver(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ObserverSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObserverSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObserverSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ObserverList) > 0 {
		for iNdEx := len(m.ObserverList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ObserverList[iNdEx])
			copy(dAtA[i:], m.ObserverList[iNdEx])
			i = encodeVarintObserver(dAtA, i, uint64(len(m.ObserverList[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *LastObserverCount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastObserverCount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastObserverCount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastChangeHeight != 0 {
		i = encodeVarintObserver(dAtA, i, uint64(m.LastChangeHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Count != 0 {
		i = encodeVarintObserver(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintObserver(dAtA []byte, offset int, v uint64) int {
	offset -= sovObserver(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ObserverMapper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovObserver(uint64(l))
	}
	if m.ObserverChain != nil {
		l = m.ObserverChain.Size()
		n += 1 + l + sovObserver(uint64(l))
	}
	if len(m.ObserverList) > 0 {
		for _, s := range m.ObserverList {
			l = len(s)
			n += 1 + l + sovObserver(uint64(l))
		}
	}
	return n
}

func (m *ObserverSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ObserverList) > 0 {
		for _, s := range m.ObserverList {
			l = len(s)
			n += 1 + l + sovObserver(uint64(l))
		}
	}
	return n
}

func (m *LastObserverCount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Count != 0 {
		n += 1 + sovObserver(uint64(m.Count))
	}
	if m.LastChangeHeight != 0 {
		n += 1 + sovObserver(uint64(m.LastChangeHeight))
	}
	return n
}

func sovObserver(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozObserver(x uint64) (n int) {
	return sovObserver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ObserverMapper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObserver
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
			return fmt.Errorf("proto: ObserverMapper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObserverMapper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthObserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObserver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverChain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserver
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
				return ErrInvalidLengthObserver
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObserver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ObserverChain == nil {
				m.ObserverChain = &common.Chain{}
			}
			if err := m.ObserverChain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthObserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObserver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObserverList = append(m.ObserverList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthObserver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ObserverSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObserver
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
			return fmt.Errorf("proto: ObserverSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObserverSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserverList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthObserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthObserver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObserverList = append(m.ObserverList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthObserver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LastObserverCount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObserver
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
			return fmt.Errorf("proto: LastObserverCount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastObserverCount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastChangeHeight", wireType)
			}
			m.LastChangeHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastChangeHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipObserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthObserver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipObserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowObserver
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
					return 0, ErrIntOverflowObserver
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
					return 0, ErrIntOverflowObserver
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
				return 0, ErrInvalidLengthObserver
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupObserver
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthObserver
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthObserver        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowObserver          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupObserver = fmt.Errorf("proto: unexpected end of group")
)
