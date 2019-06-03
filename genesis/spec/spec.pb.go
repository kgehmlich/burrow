// Copyright 2019 Monax Industries Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spec.proto

/*
	Package spec is a generated protocol buffer package.

	It is generated from these files:
		spec.proto

	It has these top-level messages:
		TemplateAccount
*/
package spec

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import crypto "github.com/hyperledger/burrow/crypto"
import balance "github.com/hyperledger/burrow/acm/balance"

import github_com_hyperledger_burrow_crypto "github.com/hyperledger/burrow/crypto"
import github_com_hyperledger_burrow_acm "github.com/hyperledger/burrow/acm"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TemplateAccount struct {
	Name        string                                        `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Address     *github_com_hyperledger_burrow_crypto.Address `protobuf:"bytes,2,opt,name=Address,proto3,customtype=github.com/hyperledger/burrow/crypto.Address" json:",omitempty" toml:",omitempty"`
	NodeAddress *github_com_hyperledger_burrow_crypto.Address `protobuf:"bytes,3,opt,name=NodeAddress,proto3,customtype=github.com/hyperledger/burrow/crypto.Address" json:",omitempty" toml:",omitempty"`
	PublicKey   *crypto.PublicKey                             `protobuf:"bytes,4,opt,name=PublicKey" json:",omitempty" toml:",omitempty"`
	Amounts     []balance.Balance                             `protobuf:"bytes,5,rep,name=Amounts" json:",omitempty" toml:",omitempty"`
	Permissions []string                                      `protobuf:"bytes,6,rep,name=Permissions" json:",omitempty" toml:",omitempty"`
	Roles       []string                                      `protobuf:"bytes,7,rep,name=Roles" json:",omitempty" toml:",omitempty"`
	Code        *github_com_hyperledger_burrow_acm.Bytecode   `protobuf:"bytes,8,opt,name=Code,proto3,customtype=github.com/hyperledger/burrow/acm.Bytecode" json:"Code,omitempty"`
}

func (m *TemplateAccount) Reset()                    { *m = TemplateAccount{} }
func (m *TemplateAccount) String() string            { return proto.CompactTextString(m) }
func (*TemplateAccount) ProtoMessage()               {}
func (*TemplateAccount) Descriptor() ([]byte, []int) { return fileDescriptorSpec, []int{0} }

func (m *TemplateAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TemplateAccount) GetPublicKey() *crypto.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *TemplateAccount) GetAmounts() []balance.Balance {
	if m != nil {
		return m.Amounts
	}
	return nil
}

func (m *TemplateAccount) GetPermissions() []string {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *TemplateAccount) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (*TemplateAccount) XXX_MessageName() string {
	return "spec.TemplateAccount"
}
func init() {
	proto.RegisterType((*TemplateAccount)(nil), "spec.TemplateAccount")
	golang_proto.RegisterType((*TemplateAccount)(nil), "spec.TemplateAccount")
}
func (m *TemplateAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TemplateAccount) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSpec(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Address != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSpec(dAtA, i, uint64(m.Address.Size()))
		n1, err := m.Address.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.NodeAddress != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSpec(dAtA, i, uint64(m.NodeAddress.Size()))
		n2, err := m.NodeAddress.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.PublicKey != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSpec(dAtA, i, uint64(m.PublicKey.Size()))
		n3, err := m.PublicKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Amounts) > 0 {
		for _, msg := range m.Amounts {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintSpec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Permissions) > 0 {
		for _, s := range m.Permissions {
			dAtA[i] = 0x32
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			dAtA[i] = 0x3a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Code != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintSpec(dAtA, i, uint64(m.Code.Size()))
		n4, err := m.Code.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeVarintSpec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TemplateAccount) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSpec(uint64(l))
	}
	if m.Address != nil {
		l = m.Address.Size()
		n += 1 + l + sovSpec(uint64(l))
	}
	if m.NodeAddress != nil {
		l = m.NodeAddress.Size()
		n += 1 + l + sovSpec(uint64(l))
	}
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovSpec(uint64(l))
	}
	if len(m.Amounts) > 0 {
		for _, e := range m.Amounts {
			l = e.Size()
			n += 1 + l + sovSpec(uint64(l))
		}
	}
	if len(m.Permissions) > 0 {
		for _, s := range m.Permissions {
			l = len(s)
			n += 1 + l + sovSpec(uint64(l))
		}
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovSpec(uint64(l))
		}
	}
	if m.Code != nil {
		l = m.Code.Size()
		n += 1 + l + sovSpec(uint64(l))
	}
	return n
}

func sovSpec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSpec(x uint64) (n int) {
	return sovSpec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TemplateAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpec
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
			return fmt.Errorf("proto: TemplateAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TemplateAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_hyperledger_burrow_crypto.Address
			m.Address = &v
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_hyperledger_burrow_crypto.Address
			m.NodeAddress = &v
			if err := m.NodeAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &crypto.PublicKey{}
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amounts = append(m.Amounts, balance.Balance{})
			if err := m.Amounts[len(m.Amounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permissions = append(m.Permissions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
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
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSpec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_hyperledger_burrow_acm.Bytecode
			m.Code = &v
			if err := m.Code.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpec
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
func skipSpec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpec
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
					return 0, ErrIntOverflowSpec
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
					return 0, ErrIntOverflowSpec
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
				return 0, ErrInvalidLengthSpec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSpec
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
				next, err := skipSpec(dAtA[start:])
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
	ErrInvalidLengthSpec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpec   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("spec.proto", fileDescriptorSpec) }
func init() { golang_proto.RegisterFile("spec.proto", fileDescriptorSpec) }

var fileDescriptorSpec = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0x3e, 0x73, 0x77, 0xbd, 0x9e, 0xaf, 0x08, 0xea, 0x29, 0xea, 0x10, 0x47, 0xc7, 0x40, 0x84,
	0x4a, 0x22, 0x1d, 0x13, 0x9d, 0xb8, 0x20, 0x58, 0x90, 0xaa, 0x2a, 0xed, 0xc4, 0x96, 0x38, 0x8f,
	0x34, 0x52, 0x1c, 0x47, 0xb6, 0x23, 0x94, 0xbf, 0xc0, 0xc4, 0xc8, 0xdc, 0x5f, 0xc2, 0x98, 0x91,
	0xb9, 0x43, 0x84, 0xae, 0x1b, 0x23, 0xbf, 0x00, 0x9d, 0x73, 0xa1, 0x37, 0x41, 0x96, 0x4e, 0x7e,
	0x9f, 0x9f, 0xbe, 0xef, 0x7b, 0x7a, 0x9f, 0x8d, 0xb1, 0x2a, 0x81, 0x79, 0xa5, 0x14, 0x5a, 0x90,
	0xc9, 0xb6, 0x3e, 0x79, 0x99, 0x66, 0xfa, 0xba, 0x8a, 0x3d, 0x26, 0xb8, 0x9f, 0x8a, 0x54, 0xf8,
	0xa6, 0x19, 0x57, 0x9f, 0x0c, 0x32, 0xc0, 0x54, 0x1d, 0xe9, 0xe4, 0x88, 0xc9, 0xba, 0xd4, 0x3d,
	0x7a, 0x1c, 0x47, 0x79, 0x54, 0x30, 0xe8, 0xe0, 0xf2, 0xcb, 0x14, 0x3f, 0xb9, 0x02, 0x5e, 0xe6,
	0x91, 0x86, 0x35, 0x63, 0xa2, 0x2a, 0x34, 0x21, 0x78, 0x72, 0x1e, 0x71, 0xb0, 0x90, 0x83, 0xdc,
	0x79, 0x68, 0x6a, 0xc2, 0xf1, 0x6c, 0x9d, 0x24, 0x12, 0x94, 0xb2, 0x1e, 0x39, 0xc8, 0x3d, 0x0a,
	0x2e, 0x6f, 0x5b, 0x7a, 0xba, 0x37, 0xc8, 0x75, 0x5d, 0x82, 0xcc, 0x21, 0x49, 0x41, 0xfa, 0x71,
	0x25, 0xa5, 0xf8, 0xec, 0xef, 0x7c, 0x77, 0xbc, 0x5f, 0x2d, 0xc5, 0xa7, 0x82, 0x67, 0x1a, 0x78,
	0xa9, 0xeb, 0xdf, 0x2d, 0x3d, 0xd6, 0x82, 0xe7, 0x67, 0xcb, 0xfb, 0xbb, 0x65, 0xd8, 0x7b, 0x90,
	0x0a, 0x2f, 0xce, 0x45, 0x02, 0xbd, 0xe5, 0xf8, 0xe1, 0x2c, 0xf7, 0x7d, 0xc8, 0x15, 0x9e, 0x5f,
	0x54, 0x71, 0x9e, 0xb1, 0x0f, 0x50, 0x5b, 0x13, 0x07, 0xb9, 0x8b, 0xd5, 0xb1, 0xb7, 0xd3, 0xfc,
	0xdb, 0x08, 0x9e, 0x0d, 0xd1, 0xbd, 0x17, 0x22, 0x97, 0x78, 0xb6, 0xe6, 0xdb, 0xcd, 0x2a, 0x6b,
	0xea, 0x8c, 0xdd, 0xc5, 0xea, 0xa9, 0xd7, 0x87, 0x10, 0x74, 0x67, 0xf0, 0xbc, 0x69, 0xe9, 0x68,
	0xd8, 0x86, 0x3a, 0x25, 0xf2, 0x0e, 0x2f, 0x2e, 0x40, 0xf2, 0x4c, 0xa9, 0x4c, 0x14, 0xca, 0x3a,
	0x70, 0xc6, 0xee, 0x7c, 0xd8, 0x64, 0xfb, 0x3c, 0xf2, 0x1a, 0x4f, 0x43, 0x91, 0x83, 0xb2, 0x66,
	0xc3, 0x05, 0x3a, 0x06, 0x79, 0x8f, 0x27, 0x6f, 0x45, 0x02, 0xd6, 0xa1, 0x09, 0x67, 0xd5, 0xb4,
	0x14, 0xdd, 0xb6, 0xf4, 0xc5, 0xbf, 0x03, 0x8a, 0x18, 0xf7, 0x82, 0x5a, 0x03, 0x13, 0x09, 0x84,
	0x86, 0x7f, 0x76, 0xf8, 0xf5, 0x86, 0x8e, 0xbe, 0xdd, 0xd0, 0x51, 0xf0, 0xa6, 0xd9, 0xd8, 0xe8,
	0xc7, 0xc6, 0x46, 0x3f, 0x37, 0x36, 0xfa, 0x7e, 0x67, 0xa3, 0xe6, 0xce, 0x46, 0x1f, 0xff, 0xa3,
	0x98, 0x42, 0x01, 0x2a, 0x53, 0xfe, 0xf6, 0x6b, 0xc4, 0x07, 0xe6, 0x55, 0xbf, 0xfa, 0x13, 0x00,
	0x00, 0xff, 0xff, 0xcf, 0x67, 0x15, 0x5d, 0x35, 0x03, 0x00, 0x00,
}
