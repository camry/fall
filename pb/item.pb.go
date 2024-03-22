// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0-devel
// 	protoc        v4.25.3
// source: pb/item.proto

package pb

import (
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

// 掉落物品
type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty" gorm:"column:type" redis:"type,omitempty"` // 物品种类
	Id   int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty" gorm:"column:id" redis:"id,omitempty"`         // 物品编号
	Num  int32 `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty" gorm:"column:num" redis:"num,omitempty"`     // 物品数量
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_pb_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_pb_item_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Item) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_pb_item_proto protoreflect.FileDescriptor

var file_pb_item_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x3c, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_item_proto_rawDescOnce sync.Once
	file_pb_item_proto_rawDescData = file_pb_item_proto_rawDesc
)

func file_pb_item_proto_rawDescGZIP() []byte {
	file_pb_item_proto_rawDescOnce.Do(func() {
		file_pb_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_item_proto_rawDescData)
	})
	return file_pb_item_proto_rawDescData
}

var file_pb_item_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_item_proto_goTypes = []interface{}{
	(*Item)(nil), // 0: pb.Item
}
var file_pb_item_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_item_proto_init() }
func file_pb_item_proto_init() {
	if File_pb_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_item_proto_goTypes,
		DependencyIndexes: file_pb_item_proto_depIdxs,
		MessageInfos:      file_pb_item_proto_msgTypes,
	}.Build()
	File_pb_item_proto = out.File
	file_pb_item_proto_rawDesc = nil
	file_pb_item_proto_goTypes = nil
	file_pb_item_proto_depIdxs = nil
}
