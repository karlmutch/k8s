// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/api/coordination/v1beta1/generated.proto

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/api/coordination/v1beta1/generated.proto

	It has these top-level messages:
		Lease
		LeaseList
		LeaseSpec
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import k8s_io_apimachinery_pkg_apis_meta_v1 "github.com/karlmutch/k8s/apis/meta/v1"
import _ "github.com/karlmutch/k8s/runtime"
import _ "github.com/karlmutch/k8s/runtime/schema"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Lease defines a lease concept.
type Lease struct {
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	Metadata *k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Specification of the Lease.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec             *LeaseSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Lease) Reset()                    { *m = Lease{} }
func (m *Lease) String() string            { return proto.CompactTextString(m) }
func (*Lease) ProtoMessage()               {}
func (*Lease) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *Lease) GetMetadata() *k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Lease) GetSpec() *LeaseSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// LeaseList is a list of Lease objects.
type LeaseList struct {
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	Metadata *k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Items is a list of schema objects.
	Items            []*Lease `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *LeaseList) Reset()                    { *m = LeaseList{} }
func (m *LeaseList) String() string            { return proto.CompactTextString(m) }
func (*LeaseList) ProtoMessage()               {}
func (*LeaseList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *LeaseList) GetMetadata() *k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LeaseList) GetItems() []*Lease {
	if m != nil {
		return m.Items
	}
	return nil
}

// LeaseSpec is a specification of a Lease.
type LeaseSpec struct {
	// holderIdentity contains the identity of the holder of a current lease.
	// +optional
	HolderIdentity *string `protobuf:"bytes,1,opt,name=holderIdentity" json:"holderIdentity,omitempty"`
	// leaseDurationSeconds is a duration that candidates for a lease need
	// to wait to force acquire it. This is measure against time of last
	// observed RenewTime.
	// +optional
	LeaseDurationSeconds *int32 `protobuf:"varint,2,opt,name=leaseDurationSeconds" json:"leaseDurationSeconds,omitempty"`
	// acquireTime is a time when the current lease was acquired.
	// +optional
	AcquireTime *k8s_io_apimachinery_pkg_apis_meta_v1.MicroTime `protobuf:"bytes,3,opt,name=acquireTime" json:"acquireTime,omitempty"`
	// renewTime is a time when the current holder of a lease has last
	// updated the lease.
	// +optional
	RenewTime *k8s_io_apimachinery_pkg_apis_meta_v1.MicroTime `protobuf:"bytes,4,opt,name=renewTime" json:"renewTime,omitempty"`
	// leaseTransitions is the number of transitions of a lease between
	// holders.
	// +optional
	LeaseTransitions *int32 `protobuf:"varint,5,opt,name=leaseTransitions" json:"leaseTransitions,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LeaseSpec) Reset()                    { *m = LeaseSpec{} }
func (m *LeaseSpec) String() string            { return proto.CompactTextString(m) }
func (*LeaseSpec) ProtoMessage()               {}
func (*LeaseSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *LeaseSpec) GetHolderIdentity() string {
	if m != nil && m.HolderIdentity != nil {
		return *m.HolderIdentity
	}
	return ""
}

func (m *LeaseSpec) GetLeaseDurationSeconds() int32 {
	if m != nil && m.LeaseDurationSeconds != nil {
		return *m.LeaseDurationSeconds
	}
	return 0
}

func (m *LeaseSpec) GetAcquireTime() *k8s_io_apimachinery_pkg_apis_meta_v1.MicroTime {
	if m != nil {
		return m.AcquireTime
	}
	return nil
}

func (m *LeaseSpec) GetRenewTime() *k8s_io_apimachinery_pkg_apis_meta_v1.MicroTime {
	if m != nil {
		return m.RenewTime
	}
	return nil
}

func (m *LeaseSpec) GetLeaseTransitions() int32 {
	if m != nil && m.LeaseTransitions != nil {
		return *m.LeaseTransitions
	}
	return 0
}

func init() {
	proto.RegisterType((*Lease)(nil), "k8s.io.api.coordination.v1beta1.Lease")
	proto.RegisterType((*LeaseList)(nil), "k8s.io.api.coordination.v1beta1.LeaseList")
	proto.RegisterType((*LeaseSpec)(nil), "k8s.io.api.coordination.v1beta1.LeaseSpec")
}
func (m *Lease) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lease) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Spec != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Spec.Size()))
		n2, err := m.Spec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *LeaseList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Metadata.Size()))
		n3, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *LeaseSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.HolderIdentity != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.HolderIdentity)))
		i += copy(dAtA[i:], *m.HolderIdentity)
	}
	if m.LeaseDurationSeconds != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.LeaseDurationSeconds))
	}
	if m.AcquireTime != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.AcquireTime.Size()))
		n4, err := m.AcquireTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.RenewTime != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.RenewTime.Size()))
		n5, err := m.RenewTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.LeaseTransitions != nil {
		dAtA[i] = 0x28
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(*m.LeaseTransitions))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Lease) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseList) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LeaseSpec) Size() (n int) {
	var l int
	_ = l
	if m.HolderIdentity != nil {
		l = len(*m.HolderIdentity)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.LeaseDurationSeconds != nil {
		n += 1 + sovGenerated(uint64(*m.LeaseDurationSeconds))
	}
	if m.AcquireTime != nil {
		l = m.AcquireTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.RenewTime != nil {
		l = m.RenewTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.LeaseTransitions != nil {
		n += 1 + sovGenerated(uint64(*m.LeaseTransitions))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lease) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Lease: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lease: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &LeaseSpec{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *LeaseList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LeaseList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &k8s_io_apimachinery_pkg_apis_meta_v1.ListMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Lease{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *LeaseSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LeaseSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HolderIdentity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.HolderIdentity = &s
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseDurationSeconds", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LeaseDurationSeconds = &v
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcquireTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AcquireTime == nil {
				m.AcquireTime = &k8s_io_apimachinery_pkg_apis_meta_v1.MicroTime{}
			}
			if err := m.AcquireTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RenewTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RenewTime == nil {
				m.RenewTime = &k8s_io_apimachinery_pkg_apis_meta_v1.MicroTime{}
			}
			if err := m.RenewTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseTransitions", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LeaseTransitions = &v
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/api/coordination/v1beta1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xc1, 0x8a, 0x1a, 0x31,
	0x1c, 0x87, 0x3b, 0xea, 0xd0, 0x1a, 0xa1, 0x94, 0xd0, 0xc3, 0xd4, 0xc3, 0x54, 0x3c, 0x88, 0x78,
	0x48, 0xaa, 0x94, 0xd2, 0x43, 0xe9, 0xa1, 0xf4, 0xd2, 0xa2, 0x94, 0x8e, 0x9e, 0x7a, 0x8b, 0x99,
	0x3f, 0x9a, 0xea, 0x24, 0xb3, 0x49, 0x74, 0xf1, 0x41, 0x16, 0xf6, 0x45, 0xf6, 0x1d, 0xf6, 0xb8,
	0x8f, 0xb0, 0xb8, 0x2f, 0xb2, 0x24, 0x8a, 0x8a, 0xb3, 0x8b, 0xb2, 0xd7, 0xdf, 0xcc, 0xf7, 0xf1,
	0xfd, 0x83, 0xe8, 0xec, 0xab, 0x21, 0x42, 0x51, 0x96, 0x0b, 0xca, 0x95, 0xd2, 0xa9, 0x90, 0xcc,
	0x0a, 0x25, 0xe9, 0xb2, 0x3b, 0x06, 0xcb, 0xba, 0x74, 0x02, 0x12, 0x34, 0xb3, 0x90, 0x92, 0x5c,
	0x2b, 0xab, 0xf0, 0xc7, 0x0d, 0x40, 0x58, 0x2e, 0xc8, 0x21, 0x40, 0xb6, 0x40, 0xfd, 0xf3, 0xde,
	0x98, 0x31, 0x3e, 0x15, 0x12, 0xf4, 0x8a, 0xe6, 0xb3, 0x89, 0x1b, 0x0c, 0xcd, 0xc0, 0x32, 0xba,
	0x2c, 0x68, 0xeb, 0xf4, 0x39, 0x4a, 0x2f, 0xa4, 0x15, 0x19, 0x14, 0x80, 0x2f, 0xa7, 0x00, 0xc3,
	0xa7, 0x90, 0xb1, 0x63, 0xae, 0x79, 0x15, 0xa0, 0xb0, 0x0f, 0xcc, 0x00, 0xee, 0xa3, 0x37, 0xae,
	0x26, 0x65, 0x96, 0x45, 0x41, 0x23, 0x68, 0xd7, 0x7a, 0x9f, 0xc8, 0xfe, 0xb8, 0x9d, 0x94, 0xe4,
	0xb3, 0x89, 0x1b, 0x0c, 0x71, 0x7f, 0x93, 0x65, 0x97, 0xfc, 0x19, 0xff, 0x07, 0x6e, 0x07, 0x60,
	0x59, 0xb2, 0x33, 0xe0, 0xef, 0xa8, 0x62, 0x72, 0xe0, 0x51, 0xc9, 0x9b, 0x3a, 0xe4, 0xc4, 0x33,
	0x11, 0xdf, 0x30, 0xcc, 0x81, 0x27, 0x9e, 0x73, 0x5d, 0x55, 0xbf, 0xf5, 0x85, 0xb1, 0xf8, 0x77,
	0xa1, 0x8d, 0x9c, 0xd7, 0xe6, 0xe8, 0xa3, 0xb2, 0x6f, 0x28, 0x14, 0x16, 0x32, 0x13, 0x95, 0x1a,
	0xe5, 0x76, 0xad, 0xd7, 0x3a, 0x2f, 0x2d, 0xd9, 0x40, 0xcd, 0x9b, 0xd2, 0xb6, 0xcb, 0xb5, 0xe2,
	0x16, 0x7a, 0x3b, 0x55, 0xf3, 0x14, 0xf4, 0xaf, 0x14, 0xa4, 0x15, 0x76, 0xe5, 0xeb, 0xaa, 0xc9,
	0xd1, 0x8a, 0x7b, 0xe8, 0xfd, 0xdc, 0x41, 0x3f, 0x17, 0xda, 0xbb, 0x87, 0xc0, 0x95, 0x4c, 0x8d,
	0x7f, 0x9d, 0x30, 0x79, 0xf2, 0x1b, 0xfe, 0x8b, 0x6a, 0x8c, 0x5f, 0x2c, 0x84, 0x86, 0x91, 0xc8,
	0x20, 0x2a, 0xfb, 0xb3, 0xe9, 0x79, 0x67, 0x0f, 0x04, 0xd7, 0xca, 0x61, 0xc9, 0xa1, 0x03, 0x0f,
	0x50, 0x55, 0x83, 0x84, 0x4b, 0x2f, 0xac, 0xbc, 0x4c, 0xb8, 0x37, 0xe0, 0x0e, 0x7a, 0xe7, 0xcb,
	0x47, 0x9a, 0x49, 0x23, 0x5c, 0xbb, 0x89, 0x42, 0x7f, 0x51, 0x61, 0xff, 0xf1, 0xe1, 0x76, 0x1d,
	0x07, 0x77, 0xeb, 0x38, 0xb8, 0x5f, 0xc7, 0xc1, 0xf5, 0x43, 0xfc, 0xea, 0xdf, 0xeb, 0xed, 0xfb,
	0x3e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xe1, 0x9b, 0x21, 0x74, 0x03, 0x00, 0x00,
}
