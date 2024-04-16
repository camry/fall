// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0-devel
// 	protoc        v4.25.3
// source: pb/table.proto

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

// 逐个百分比掉落表
type TablePercent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AutoId     int32   `protobuf:"varint,1,opt,name=auto_id,json=autoId,proto3" json:"auto_id,omitempty" gorm:"column:auto_id" redis:"auto_id"`                         // 自动编号
	FallType   int32   `protobuf:"varint,2,opt,name=fall_type,json=fallType,proto3" json:"fall_type,omitempty" gorm:"column:fall_type" redis:"fall_type"`               // 掉落种类
	FallTypeId int32   `protobuf:"varint,3,opt,name=fall_type_id,json=fallTypeId,proto3" json:"fall_type_id,omitempty" gorm:"column:fall_type_id" redis:"fall_type_id"` // 掉落种类编号
	Prob       float32 `protobuf:"fixed32,4,opt,name=prob,proto3" json:"prob,omitempty" gorm:"column:prob" redis:"prob"`                                                // 概率浮点数
	NumMin     int32   `protobuf:"varint,5,opt,name=num_min,json=numMin,proto3" json:"num_min,omitempty" gorm:"column:num_min" redis:"num_min"`                         // 最小数量
	NumMax     int32   `protobuf:"varint,6,opt,name=num_max,json=numMax,proto3" json:"num_max,omitempty" gorm:"column:num_max" redis:"num_max"`                         // 最大数量
}

func (x *TablePercent) Reset() {
	*x = TablePercent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TablePercent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TablePercent) ProtoMessage() {}

func (x *TablePercent) ProtoReflect() protoreflect.Message {
	mi := &file_pb_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TablePercent.ProtoReflect.Descriptor instead.
func (*TablePercent) Descriptor() ([]byte, []int) {
	return file_pb_table_proto_rawDescGZIP(), []int{0}
}

func (x *TablePercent) GetAutoId() int32 {
	if x != nil {
		return x.AutoId
	}
	return 0
}

func (x *TablePercent) GetFallType() int32 {
	if x != nil {
		return x.FallType
	}
	return 0
}

func (x *TablePercent) GetFallTypeId() int32 {
	if x != nil {
		return x.FallTypeId
	}
	return 0
}

func (x *TablePercent) GetProb() float32 {
	if x != nil {
		return x.Prob
	}
	return 0
}

func (x *TablePercent) GetNumMin() int32 {
	if x != nil {
		return x.NumMin
	}
	return 0
}

func (x *TablePercent) GetNumMax() int32 {
	if x != nil {
		return x.NumMax
	}
	return 0
}

// 权重掉落组式掉落母集进阶掉落表
type TableWeightGroupMaster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MasterId      int32 `protobuf:"varint,1,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty" gorm:"column:master_id" redis:"master_id"`                              // 母集编号
	SubsetId      int32 `protobuf:"varint,2,opt,name=subset_id,json=subsetId,proto3" json:"subset_id,omitempty" gorm:"column:subset_id" redis:"subset_id"`                              // 子集编号
	IsNextSubset  bool  `protobuf:"varint,3,opt,name=is_next_subset,json=isNextSubset,proto3" json:"is_next_subset,omitempty" gorm:"column:is_next_subset" redis:"is_next_subset"`      // 是进阶子集？
	NextSubsetId  int32 `protobuf:"varint,4,opt,name=next_subset_id,json=nextSubsetId,proto3" json:"next_subset_id,omitempty" gorm:"column:next_subset_id" redis:"next_subset_id"`      // 进阶子集
	NextSubsetMin int32 `protobuf:"varint,5,opt,name=next_subset_min,json=nextSubsetMin,proto3" json:"next_subset_min,omitempty" gorm:"column:next_subset_min" redis:"next_subset_min"` // 最小进阶次数
	NextSubsetMax int32 `protobuf:"varint,6,opt,name=next_subset_max,json=nextSubsetMax,proto3" json:"next_subset_max,omitempty" gorm:"column:next_subset_max" redis:"next_subset_max"` // 最大进阶次数
	AdvanceNum    int32 `protobuf:"varint,7,opt,name=advance_num,json=advanceNum,proto3" json:"advance_num,omitempty" gorm:"column:advance_num" redis:"advance_num"`                    // 现有进阶次数
}

func (x *TableWeightGroupMaster) Reset() {
	*x = TableWeightGroupMaster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableWeightGroupMaster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableWeightGroupMaster) ProtoMessage() {}

func (x *TableWeightGroupMaster) ProtoReflect() protoreflect.Message {
	mi := &file_pb_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableWeightGroupMaster.ProtoReflect.Descriptor instead.
func (*TableWeightGroupMaster) Descriptor() ([]byte, []int) {
	return file_pb_table_proto_rawDescGZIP(), []int{1}
}

func (x *TableWeightGroupMaster) GetMasterId() int32 {
	if x != nil {
		return x.MasterId
	}
	return 0
}

func (x *TableWeightGroupMaster) GetSubsetId() int32 {
	if x != nil {
		return x.SubsetId
	}
	return 0
}

func (x *TableWeightGroupMaster) GetIsNextSubset() bool {
	if x != nil {
		return x.IsNextSubset
	}
	return false
}

func (x *TableWeightGroupMaster) GetNextSubsetId() int32 {
	if x != nil {
		return x.NextSubsetId
	}
	return 0
}

func (x *TableWeightGroupMaster) GetNextSubsetMin() int32 {
	if x != nil {
		return x.NextSubsetMin
	}
	return 0
}

func (x *TableWeightGroupMaster) GetNextSubsetMax() int32 {
	if x != nil {
		return x.NextSubsetMax
	}
	return 0
}

func (x *TableWeightGroupMaster) GetAdvanceNum() int32 {
	if x != nil {
		return x.AdvanceNum
	}
	return 0
}

// 权重掉落组式掉落子集表
type TableWeightGroupSubset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AutoId       int32 `protobuf:"varint,1,opt,name=auto_id,json=autoId,proto3" json:"auto_id,omitempty" gorm:"column:auto_id" redis:"auto_id"`                                   // 自动编号
	SubsetId     int32 `protobuf:"varint,2,opt,name=subset_id,json=subsetId,proto3" json:"subset_id,omitempty" gorm:"column:subset_id" redis:"subset_id"`                         // 子集编号
	FallType     int32 `protobuf:"varint,3,opt,name=fall_type,json=fallType,proto3" json:"fall_type,omitempty" gorm:"column:fall_type" redis:"fall_type"`                         // 掉落种类
	FallTypeId   int32 `protobuf:"varint,4,opt,name=fall_type_id,json=fallTypeId,proto3" json:"fall_type_id,omitempty" gorm:"column:fall_type_id" redis:"fall_type_id"`           // 掉落种类编号
	SubsetNumMin int32 `protobuf:"varint,5,opt,name=subset_num_min,json=subsetNumMin,proto3" json:"subset_num_min,omitempty" gorm:"column:subset_num_min" redis:"subset_num_min"` // 最小数量
	SubsetNumMax int32 `protobuf:"varint,6,opt,name=subset_num_max,json=subsetNumMax,proto3" json:"subset_num_max,omitempty" gorm:"column:subset_num_max" redis:"subset_num_max"` // 最大数量
	SubsetWeight int32 `protobuf:"varint,7,opt,name=subset_weight,json=subsetWeight,proto3" json:"subset_weight,omitempty" gorm:"column:subset_weight" redis:"subset_weight"`     // 权重
}

func (x *TableWeightGroupSubset) Reset() {
	*x = TableWeightGroupSubset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableWeightGroupSubset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableWeightGroupSubset) ProtoMessage() {}

func (x *TableWeightGroupSubset) ProtoReflect() protoreflect.Message {
	mi := &file_pb_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableWeightGroupSubset.ProtoReflect.Descriptor instead.
func (*TableWeightGroupSubset) Descriptor() ([]byte, []int) {
	return file_pb_table_proto_rawDescGZIP(), []int{2}
}

func (x *TableWeightGroupSubset) GetAutoId() int32 {
	if x != nil {
		return x.AutoId
	}
	return 0
}

func (x *TableWeightGroupSubset) GetSubsetId() int32 {
	if x != nil {
		return x.SubsetId
	}
	return 0
}

func (x *TableWeightGroupSubset) GetFallType() int32 {
	if x != nil {
		return x.FallType
	}
	return 0
}

func (x *TableWeightGroupSubset) GetFallTypeId() int32 {
	if x != nil {
		return x.FallTypeId
	}
	return 0
}

func (x *TableWeightGroupSubset) GetSubsetNumMin() int32 {
	if x != nil {
		return x.SubsetNumMin
	}
	return 0
}

func (x *TableWeightGroupSubset) GetSubsetNumMax() int32 {
	if x != nil {
		return x.SubsetNumMax
	}
	return 0
}

func (x *TableWeightGroupSubset) GetSubsetWeight() int32 {
	if x != nil {
		return x.SubsetWeight
	}
	return 0
}

// 木桶原理掉落表
type TableVat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VatId       int32   `protobuf:"varint,1,opt,name=vat_id,json=vatId,proto3" json:"vat_id,omitempty" gorm:"column:vat_id" redis:"vat_id"`                                 // 木桶编号
	FallType    int32   `protobuf:"varint,2,opt,name=fall_type,json=fallType,proto3" json:"fall_type,omitempty" gorm:"column:fall_type" redis:"fall_type"`                  // 掉落种类
	FallTypeId  int32   `protobuf:"varint,3,opt,name=fall_type_id,json=fallTypeId,proto3" json:"fall_type_id,omitempty" gorm:"column:fall_type_id" redis:"fall_type_id"`    // 掉落种类编号
	VatNumMin   int32   `protobuf:"varint,4,opt,name=vat_num_min,json=vatNumMin,proto3" json:"vat_num_min,omitempty" gorm:"column:vat_num_min" redis:"vat_num_min"`         // 最小数量
	VatNumMax   int32   `protobuf:"varint,5,opt,name=vat_num_max,json=vatNumMax,proto3" json:"vat_num_max,omitempty" gorm:"column:vat_num_max" redis:"vat_num_max"`         // 最大数量
	VatWeight   int32   `protobuf:"varint,6,opt,name=vat_weight,json=vatWeight,proto3" json:"vat_weight,omitempty" gorm:"column:vat_weight" redis:"vat_weight"`             // 权重（空缺量*系数加成）
	ExpectNum   int32   `protobuf:"varint,7,opt,name=expect_num,json=expectNum,proto3" json:"expect_num,omitempty" gorm:"column:expect_num" redis:"expect_num"`             // 预期量
	ExistingNum int32   `protobuf:"varint,8,opt,name=existing_num,json=existingNum,proto3" json:"existing_num,omitempty" gorm:"column:existing_num" redis:"existing_num"`   // 现有量
	VacancyNum  int32   `protobuf:"varint,9,opt,name=vacancy_num,json=vacancyNum,proto3" json:"vacancy_num,omitempty" gorm:"column:vacancy_num" redis:"vacancy_num"`        // 空缺量（预期量-现有量）
	VacancyProb float32 `protobuf:"fixed32,10,opt,name=vacancy_prob,json=vacancyProb,proto3" json:"vacancy_prob,omitempty" gorm:"column:vacancy_prob" redis:"vacancy_prob"` // 空缺率（空缺量/预期量）
	Coefficient int32   `protobuf:"varint,11,opt,name=coefficient,proto3" json:"coefficient,omitempty" gorm:"column:coefficient" redis:"coefficient"`                       // 系数加成（4=空缺率>=90%,2=空缺率>=70% AND 空缺率<90%,DEFAULT 1）
}

func (x *TableVat) Reset() {
	*x = TableVat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_table_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableVat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableVat) ProtoMessage() {}

func (x *TableVat) ProtoReflect() protoreflect.Message {
	mi := &file_pb_table_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableVat.ProtoReflect.Descriptor instead.
func (*TableVat) Descriptor() ([]byte, []int) {
	return file_pb_table_proto_rawDescGZIP(), []int{3}
}

func (x *TableVat) GetVatId() int32 {
	if x != nil {
		return x.VatId
	}
	return 0
}

func (x *TableVat) GetFallType() int32 {
	if x != nil {
		return x.FallType
	}
	return 0
}

func (x *TableVat) GetFallTypeId() int32 {
	if x != nil {
		return x.FallTypeId
	}
	return 0
}

func (x *TableVat) GetVatNumMin() int32 {
	if x != nil {
		return x.VatNumMin
	}
	return 0
}

func (x *TableVat) GetVatNumMax() int32 {
	if x != nil {
		return x.VatNumMax
	}
	return 0
}

func (x *TableVat) GetVatWeight() int32 {
	if x != nil {
		return x.VatWeight
	}
	return 0
}

func (x *TableVat) GetExpectNum() int32 {
	if x != nil {
		return x.ExpectNum
	}
	return 0
}

func (x *TableVat) GetExistingNum() int32 {
	if x != nil {
		return x.ExistingNum
	}
	return 0
}

func (x *TableVat) GetVacancyNum() int32 {
	if x != nil {
		return x.VacancyNum
	}
	return 0
}

func (x *TableVat) GetVacancyProb() float32 {
	if x != nil {
		return x.VacancyProb
	}
	return 0
}

func (x *TableVat) GetCoefficient() int32 {
	if x != nil {
		return x.Coefficient
	}
	return 0
}

var File_pb_table_proto protoreflect.FileDescriptor

var file_pb_table_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0xac, 0x01, 0x0a, 0x0c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x75, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x66, 0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x66,
	0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x66, 0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x72, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x70, 0x72, 0x6f,
	0x62, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x4d, 0x69, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x75,
	0x6d, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x4d, 0x61, 0x78, 0x22, 0x8f, 0x02, 0x0a, 0x16, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x75, 0x62, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x69, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x12, 0x24,
	0x0a, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x75, 0x62,
	0x73, 0x65, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x75, 0x62, 0x73, 0x65,
	0x74, 0x4d, 0x61, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0xfe, 0x01, 0x0a, 0x16, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x75, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62,
	0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x75,
	0x62, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x6c, 0x6c, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x61, 0x6c, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x61, 0x6c, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73,
	0x75, 0x62, 0x73, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x4d, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x73,
	0x75, 0x62, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x4d, 0x61,
	0x78, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x5f, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xe7, 0x02, 0x0a, 0x08, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x56, 0x61, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x76, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61,
	0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66,
	0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x61, 0x6c, 0x6c, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66,
	0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x76, 0x61, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x76, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x4d, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0b, 0x76, 0x61, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x76, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x4d, 0x61, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x74,
	0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76,
	0x61, 0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65,
	0x63, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x65, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x61,
	0x63, 0x61, 0x6e, 0x63, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x76,
	0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0b, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_table_proto_rawDescOnce sync.Once
	file_pb_table_proto_rawDescData = file_pb_table_proto_rawDesc
)

func file_pb_table_proto_rawDescGZIP() []byte {
	file_pb_table_proto_rawDescOnce.Do(func() {
		file_pb_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_table_proto_rawDescData)
	})
	return file_pb_table_proto_rawDescData
}

var file_pb_table_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_table_proto_goTypes = []interface{}{
	(*TablePercent)(nil),           // 0: pb.TablePercent
	(*TableWeightGroupMaster)(nil), // 1: pb.TableWeightGroupMaster
	(*TableWeightGroupSubset)(nil), // 2: pb.TableWeightGroupSubset
	(*TableVat)(nil),               // 3: pb.TableVat
}
var file_pb_table_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_table_proto_init() }
func file_pb_table_proto_init() {
	if File_pb_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TablePercent); i {
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
		file_pb_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableWeightGroupMaster); i {
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
		file_pb_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableWeightGroupSubset); i {
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
		file_pb_table_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableVat); i {
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
			RawDescriptor: file_pb_table_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_table_proto_goTypes,
		DependencyIndexes: file_pb_table_proto_depIdxs,
		MessageInfos:      file_pb_table_proto_msgTypes,
	}.Build()
	File_pb_table_proto = out.File
	file_pb_table_proto_rawDesc = nil
	file_pb_table_proto_goTypes = nil
	file_pb_table_proto_depIdxs = nil
}
