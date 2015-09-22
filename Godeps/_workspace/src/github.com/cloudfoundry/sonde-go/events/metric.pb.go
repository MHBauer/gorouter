// Code generated by protoc-gen-gogo.
// source: metric.proto
// DO NOT EDIT!

package events

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// / A ValueMetric indicates the value of a metric at an instant in time.
type ValueMetric struct {
	Name             *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *float64 `protobuf:"fixed64,2,req,name=value" json:"value,omitempty"`
	Unit             *string  `protobuf:"bytes,3,req,name=unit" json:"unit,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ValueMetric) Reset()         { *m = ValueMetric{} }
func (m *ValueMetric) String() string { return proto.CompactTextString(m) }
func (*ValueMetric) ProtoMessage()    {}

func (m *ValueMetric) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ValueMetric) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *ValueMetric) GetUnit() string {
	if m != nil && m.Unit != nil {
		return *m.Unit
	}
	return ""
}

// / A CounterEvent represents the increment of a counter. It contains only the change in the value; it is the responsibility of downstream consumers to maintain the value of the counter.
type CounterEvent struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Delta            *uint64 `protobuf:"varint,2,req,name=delta" json:"delta,omitempty"`
	Total            *uint64 `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CounterEvent) Reset()         { *m = CounterEvent{} }
func (m *CounterEvent) String() string { return proto.CompactTextString(m) }
func (*CounterEvent) ProtoMessage()    {}

func (m *CounterEvent) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CounterEvent) GetDelta() uint64 {
	if m != nil && m.Delta != nil {
		return *m.Delta
	}
	return 0
}

func (m *CounterEvent) GetTotal() uint64 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

// / A ContainerMetric records resource usage of an app in a container.
type ContainerMetric struct {
	ApplicationId    *string  `protobuf:"bytes,1,req,name=applicationId" json:"applicationId,omitempty"`
	InstanceIndex    *int32   `protobuf:"varint,2,req,name=instanceIndex" json:"instanceIndex,omitempty"`
	CpuPercentage    *float64 `protobuf:"fixed64,3,req,name=cpuPercentage" json:"cpuPercentage,omitempty"`
	MemoryBytes      *uint64  `protobuf:"varint,4,req,name=memoryBytes" json:"memoryBytes,omitempty"`
	DiskBytes        *uint64  `protobuf:"varint,5,req,name=diskBytes" json:"diskBytes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ContainerMetric) Reset()         { *m = ContainerMetric{} }
func (m *ContainerMetric) String() string { return proto.CompactTextString(m) }
func (*ContainerMetric) ProtoMessage()    {}

func (m *ContainerMetric) GetApplicationId() string {
	if m != nil && m.ApplicationId != nil {
		return *m.ApplicationId
	}
	return ""
}

func (m *ContainerMetric) GetInstanceIndex() int32 {
	if m != nil && m.InstanceIndex != nil {
		return *m.InstanceIndex
	}
	return 0
}

func (m *ContainerMetric) GetCpuPercentage() float64 {
	if m != nil && m.CpuPercentage != nil {
		return *m.CpuPercentage
	}
	return 0
}

func (m *ContainerMetric) GetMemoryBytes() uint64 {
	if m != nil && m.MemoryBytes != nil {
		return *m.MemoryBytes
	}
	return 0
}

func (m *ContainerMetric) GetDiskBytes() uint64 {
	if m != nil && m.DiskBytes != nil {
		return *m.DiskBytes
	}
	return 0
}

func (m *ValueMetric) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ValueMetric) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	} else {
		data[i] = 0xa
		i++
		i = encodeVarintMetric(data, i, uint64(len(*m.Name)))
		i += copy(data[i:], *m.Name)
	}
	if m.Value == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("value")
	} else {
		data[i] = 0x11
		i++
		i = encodeFixed64Metric(data, i, uint64(math.Float64bits(*m.Value)))
	}
	if m.Unit == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("unit")
	} else {
		data[i] = 0x1a
		i++
		i = encodeVarintMetric(data, i, uint64(len(*m.Unit)))
		i += copy(data[i:], *m.Unit)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CounterEvent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CounterEvent) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	} else {
		data[i] = 0xa
		i++
		i = encodeVarintMetric(data, i, uint64(len(*m.Name)))
		i += copy(data[i:], *m.Name)
	}
	if m.Delta == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("delta")
	} else {
		data[i] = 0x10
		i++
		i = encodeVarintMetric(data, i, uint64(*m.Delta))
	}
	if m.Total != nil {
		data[i] = 0x18
		i++
		i = encodeVarintMetric(data, i, uint64(*m.Total))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ContainerMetric) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ContainerMetric) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ApplicationId == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("applicationId")
	} else {
		data[i] = 0xa
		i++
		i = encodeVarintMetric(data, i, uint64(len(*m.ApplicationId)))
		i += copy(data[i:], *m.ApplicationId)
	}
	if m.InstanceIndex == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("instanceIndex")
	} else {
		data[i] = 0x10
		i++
		i = encodeVarintMetric(data, i, uint64(*m.InstanceIndex))
	}
	if m.CpuPercentage == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("cpuPercentage")
	} else {
		data[i] = 0x19
		i++
		i = encodeFixed64Metric(data, i, uint64(math.Float64bits(*m.CpuPercentage)))
	}
	if m.MemoryBytes == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("memoryBytes")
	} else {
		data[i] = 0x20
		i++
		i = encodeVarintMetric(data, i, uint64(*m.MemoryBytes))
	}
	if m.DiskBytes == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("diskBytes")
	} else {
		data[i] = 0x28
		i++
		i = encodeVarintMetric(data, i, uint64(*m.DiskBytes))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Metric(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Metric(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMetric(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ValueMetric) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.Value != nil {
		n += 9
	}
	if m.Unit != nil {
		l = len(*m.Unit)
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CounterEvent) Size() (n int) {
	var l int
	_ = l
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.Delta != nil {
		n += 1 + sovMetric(uint64(*m.Delta))
	}
	if m.Total != nil {
		n += 1 + sovMetric(uint64(*m.Total))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ContainerMetric) Size() (n int) {
	var l int
	_ = l
	if m.ApplicationId != nil {
		l = len(*m.ApplicationId)
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.InstanceIndex != nil {
		n += 1 + sovMetric(uint64(*m.InstanceIndex))
	}
	if m.CpuPercentage != nil {
		n += 9
	}
	if m.MemoryBytes != nil {
		n += 1 + sovMetric(uint64(*m.MemoryBytes))
	}
	if m.DiskBytes != nil {
		n += 1 + sovMetric(uint64(*m.DiskBytes))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMetric(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetric(x uint64) (n int) {
	return sovMetric(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValueMetric) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Name = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			v2 := float64(math.Float64frombits(v))
			m.Value = &v2
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Unit = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipMetric(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("value")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("unit")
	}

	return nil
}
func (m *CounterEvent) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Name = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delta", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Delta = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Total = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipMetric(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("delta")
	}

	return nil
}
func (m *ContainerMetric) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.ApplicationId = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstanceIndex", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InstanceIndex = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field CpuPercentage", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			v2 := float64(math.Float64frombits(v))
			m.CpuPercentage = &v2
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryBytes", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MemoryBytes = &v
			hasFields[0] |= uint64(0x00000008)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiskBytes", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DiskBytes = &v
			hasFields[0] |= uint64(0x00000010)
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipMetric(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("applicationId")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("instanceIndex")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("cpuPercentage")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("memoryBytes")
	}
	if hasFields[0]&uint64(0x00000010) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("diskBytes")
	}

	return nil
}
func skipMetric(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMetric
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipMetric(data[start:])
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
	ErrInvalidLengthMetric = fmt.Errorf("proto: negative length found during unmarshaling")
)
