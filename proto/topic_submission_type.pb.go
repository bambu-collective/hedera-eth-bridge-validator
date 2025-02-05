// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: topic_submission_type.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TopicSubmissionType int32

const (
	TopicSubmissionType_EthSignature   TopicSubmissionType = 0
	TopicSubmissionType_EthTransaction TopicSubmissionType = 1
)

// Enum value maps for TopicSubmissionType.
var (
	TopicSubmissionType_name = map[int32]string{
		0: "EthSignature",
		1: "EthTransaction",
	}
	TopicSubmissionType_value = map[string]int32{
		"EthSignature":   0,
		"EthTransaction": 1,
	}
)

func (x TopicSubmissionType) Enum() *TopicSubmissionType {
	p := new(TopicSubmissionType)
	*p = x
	return p
}

func (x TopicSubmissionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TopicSubmissionType) Descriptor() protoreflect.EnumDescriptor {
	return file_topic_submission_type_proto_enumTypes[0].Descriptor()
}

func (TopicSubmissionType) Type() protoreflect.EnumType {
	return &file_topic_submission_type_proto_enumTypes[0]
}

func (x TopicSubmissionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TopicSubmissionType.Descriptor instead.
func (TopicSubmissionType) EnumDescriptor() ([]byte, []int) {
	return file_topic_submission_type_proto_rawDescGZIP(), []int{0}
}

var File_topic_submission_type_proto protoreflect.FileDescriptor

var file_topic_submission_type_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x3b, 0x0a, 0x13, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x45,
	0x74, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x45, 0x74, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x69, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x2d, 0x65, 0x74, 0x68, 0x2d, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2d, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_topic_submission_type_proto_rawDescOnce sync.Once
	file_topic_submission_type_proto_rawDescData = file_topic_submission_type_proto_rawDesc
)

func file_topic_submission_type_proto_rawDescGZIP() []byte {
	file_topic_submission_type_proto_rawDescOnce.Do(func() {
		file_topic_submission_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_topic_submission_type_proto_rawDescData)
	})
	return file_topic_submission_type_proto_rawDescData
}

var file_topic_submission_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_topic_submission_type_proto_goTypes = []interface{}{
	(TopicSubmissionType)(0), // 0: proto.TopicSubmissionType
}
var file_topic_submission_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_topic_submission_type_proto_init() }
func file_topic_submission_type_proto_init() {
	if File_topic_submission_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_topic_submission_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_topic_submission_type_proto_goTypes,
		DependencyIndexes: file_topic_submission_type_proto_depIdxs,
		EnumInfos:         file_topic_submission_type_proto_enumTypes,
	}.Build()
	File_topic_submission_type_proto = out.File
	file_topic_submission_type_proto_rawDesc = nil
	file_topic_submission_type_proto_goTypes = nil
	file_topic_submission_type_proto_depIdxs = nil
}
