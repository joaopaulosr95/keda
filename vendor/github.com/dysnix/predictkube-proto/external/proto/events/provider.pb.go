// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.1
// source: proto/events/provider.proto

package events

import (
	enums "github.com/dysnix/predictkube-proto/external/proto/enums"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QueryRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *QueryRange) Reset() {
	*x = QueryRange{}
	mi := &file_proto_events_provider_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRange) ProtoMessage() {}

func (x *QueryRange) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_provider_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRange.ProtoReflect.Descriptor instead.
func (*QueryRange) Descriptor() ([]byte, []int) {
	return file_proto_events_provider_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRange) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *QueryRange) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type RawQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body      string                 `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Range     *QueryRange            `protobuf:"bytes,2,opt,name=range,proto3" json:"range,omitempty"`
	GivenTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=given_time,json=givenTime,proto3" json:"given_time,omitempty"`
}

func (x *RawQuery) Reset() {
	*x = RawQuery{}
	mi := &file_proto_events_provider_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RawQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawQuery) ProtoMessage() {}

func (x *RawQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_provider_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawQuery.ProtoReflect.Descriptor instead.
func (*RawQuery) Descriptor() ([]byte, []int) {
	return file_proto_events_provider_proto_rawDescGZIP(), []int{1}
}

func (x *RawQuery) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *RawQuery) GetRange() *QueryRange {
	if x != nil {
		return x.Range
	}
	return nil
}

func (x *RawQuery) GetGivenTime() *timestamppb.Timestamp {
	if x != nil {
		return x.GivenTime
	}
	return nil
}

type History struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Period  []enums.PeriodType  `protobuf:"varint,1,rep,packed,name=period,proto3,enum=enums.PeriodType" json:"period,omitempty"`
	Metrics []enums.MetricsType `protobuf:"varint,2,rep,packed,name=metrics,proto3,enum=enums.MetricsType" json:"metrics,omitempty"`
}

func (x *History) Reset() {
	*x = History{}
	mi := &file_proto_events_provider_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *History) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*History) ProtoMessage() {}

func (x *History) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_provider_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use History.ProtoReflect.Descriptor instead.
func (*History) Descriptor() ([]byte, []int) {
	return file_proto_events_provider_proto_rawDescGZIP(), []int{2}
}

func (x *History) GetPeriod() []enums.PeriodType {
	if x != nil {
		return x.Period
	}
	return nil
}

func (x *History) GetMetrics() []enums.MetricsType {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_proto_events_provider_proto protoreflect.FileDescriptor

var file_proto_events_provider_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x08, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x67,
	0x69, 0x76, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x62, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x2c,
	0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x3b, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x79, 0x73, 0x6e, 0x69,
	0x78, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x6b, 0x75, 0x62, 0x65, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_events_provider_proto_rawDescOnce sync.Once
	file_proto_events_provider_proto_rawDescData = file_proto_events_provider_proto_rawDesc
)

func file_proto_events_provider_proto_rawDescGZIP() []byte {
	file_proto_events_provider_proto_rawDescOnce.Do(func() {
		file_proto_events_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_events_provider_proto_rawDescData)
	})
	return file_proto_events_provider_proto_rawDescData
}

var file_proto_events_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_events_provider_proto_goTypes = []any{
	(*QueryRange)(nil),            // 0: events.QueryRange
	(*RawQuery)(nil),              // 1: events.RawQuery
	(*History)(nil),               // 2: events.History
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(enums.PeriodType)(0),         // 4: enums.PeriodType
	(enums.MetricsType)(0),        // 5: enums.MetricsType
}
var file_proto_events_provider_proto_depIdxs = []int32{
	3, // 0: events.QueryRange.start_time:type_name -> google.protobuf.Timestamp
	3, // 1: events.QueryRange.end_time:type_name -> google.protobuf.Timestamp
	0, // 2: events.RawQuery.range:type_name -> events.QueryRange
	3, // 3: events.RawQuery.given_time:type_name -> google.protobuf.Timestamp
	4, // 4: events.History.period:type_name -> enums.PeriodType
	5, // 5: events.History.metrics:type_name -> enums.MetricsType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_events_provider_proto_init() }
func file_proto_events_provider_proto_init() {
	if File_proto_events_provider_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_events_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_events_provider_proto_goTypes,
		DependencyIndexes: file_proto_events_provider_proto_depIdxs,
		MessageInfos:      file_proto_events_provider_proto_msgTypes,
	}.Build()
	File_proto_events_provider_proto = out.File
	file_proto_events_provider_proto_rawDesc = nil
	file_proto_events_provider_proto_goTypes = nil
	file_proto_events_provider_proto_depIdxs = nil
}
