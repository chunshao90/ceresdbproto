// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: prometheus.proto

package ceresprompb

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

type FilterType int32

const (
	FilterType_LITERAL_OR       FilterType = 0
	FilterType_NOT_LITERAL_OR   FilterType = 1
	FilterType_REGEXP           FilterType = 2
	FilterType_NOT_REGEXP_MATCH FilterType = 3
)

// Enum value maps for FilterType.
var (
	FilterType_name = map[int32]string{
		0: "LITERAL_OR",
		1: "NOT_LITERAL_OR",
		2: "REGEXP",
		3: "NOT_REGEXP_MATCH",
	}
	FilterType_value = map[string]int32{
		"LITERAL_OR":       0,
		"NOT_LITERAL_OR":   1,
		"REGEXP":           2,
		"NOT_REGEXP_MATCH": 3,
	}
)

func (x FilterType) Enum() *FilterType {
	p := new(FilterType)
	*p = x
	return p
}

func (x FilterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterType) Descriptor() protoreflect.EnumDescriptor {
	return file_prometheus_proto_enumTypes[0].Descriptor()
}

func (FilterType) Type() protoreflect.EnumType {
	return &file_prometheus_proto_enumTypes[0]
}

func (x FilterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterType.Descriptor instead.
func (FilterType) EnumDescriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{0}
}

type SubExpr_OperatorType int32

const (
	SubExpr_AGGR   SubExpr_OperatorType = 0
	SubExpr_FUNC   SubExpr_OperatorType = 1
	SubExpr_BINARY SubExpr_OperatorType = 2
)

// Enum value maps for SubExpr_OperatorType.
var (
	SubExpr_OperatorType_name = map[int32]string{
		0: "AGGR",
		1: "FUNC",
		2: "BINARY",
	}
	SubExpr_OperatorType_value = map[string]int32{
		"AGGR":   0,
		"FUNC":   1,
		"BINARY": 2,
	}
)

func (x SubExpr_OperatorType) Enum() *SubExpr_OperatorType {
	p := new(SubExpr_OperatorType)
	*p = x
	return p
}

func (x SubExpr_OperatorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubExpr_OperatorType) Descriptor() protoreflect.EnumDescriptor {
	return file_prometheus_proto_enumTypes[1].Descriptor()
}

func (SubExpr_OperatorType) Type() protoreflect.EnumType {
	return &file_prometheus_proto_enumTypes[1]
}

func (x SubExpr_OperatorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubExpr_OperatorType.Descriptor instead.
func (SubExpr_OperatorType) EnumDescriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{1, 0}
}

type Expr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Node:
	//
	//	*Expr_Operand
	//	*Expr_SubExpr
	Node isExpr_Node `protobuf_oneof:"Node"`
}

func (x *Expr) Reset() {
	*x = Expr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expr) ProtoMessage() {}

func (x *Expr) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expr.ProtoReflect.Descriptor instead.
func (*Expr) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{0}
}

func (m *Expr) GetNode() isExpr_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (x *Expr) GetOperand() *Operand {
	if x, ok := x.GetNode().(*Expr_Operand); ok {
		return x.Operand
	}
	return nil
}

func (x *Expr) GetSubExpr() *SubExpr {
	if x, ok := x.GetNode().(*Expr_SubExpr); ok {
		return x.SubExpr
	}
	return nil
}

type isExpr_Node interface {
	isExpr_Node()
}

type Expr_Operand struct {
	Operand *Operand `protobuf:"bytes,1,opt,name=operand,proto3,oneof"`
}

type Expr_SubExpr struct {
	SubExpr *SubExpr `protobuf:"bytes,2,opt,name=sub_expr,json=subExpr,proto3,oneof"`
}

func (*Expr_Operand) isExpr_Node() {}

func (*Expr_SubExpr) isExpr_Node() {}

type SubExpr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpType   SubExpr_OperatorType `protobuf:"varint,1,opt,name=op_type,json=opType,proto3,enum=prometheus.SubExpr_OperatorType" json:"op_type,omitempty"`
	Operator string               `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Operands []*Expr              `protobuf:"bytes,3,rep,name=operands,proto3" json:"operands,omitempty"`
	// for aggr
	Group   []string `protobuf:"bytes,4,rep,name=group,proto3" json:"group,omitempty"`
	Without bool     `protobuf:"varint,5,opt,name=without,proto3" json:"without,omitempty"`
}

func (x *SubExpr) Reset() {
	*x = SubExpr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubExpr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubExpr) ProtoMessage() {}

func (x *SubExpr) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubExpr.ProtoReflect.Descriptor instead.
func (*SubExpr) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{1}
}

func (x *SubExpr) GetOpType() SubExpr_OperatorType {
	if x != nil {
		return x.OpType
	}
	return SubExpr_AGGR
}

func (x *SubExpr) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *SubExpr) GetOperands() []*Expr {
	if x != nil {
		return x.Operands
	}
	return nil
}

func (x *SubExpr) GetGroup() []string {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *SubExpr) GetWithout() bool {
	if x != nil {
		return x.Without
	}
	return false
}

type Operand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Operand_FloatVal
	//	*Operand_StringVal
	//	*Operand_Selector
	Value isOperand_Value `protobuf_oneof:"Value"`
}

func (x *Operand) Reset() {
	*x = Operand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operand) ProtoMessage() {}

func (x *Operand) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operand.ProtoReflect.Descriptor instead.
func (*Operand) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{2}
}

func (m *Operand) GetValue() isOperand_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Operand) GetFloatVal() float64 {
	if x, ok := x.GetValue().(*Operand_FloatVal); ok {
		return x.FloatVal
	}
	return 0
}

func (x *Operand) GetStringVal() string {
	if x, ok := x.GetValue().(*Operand_StringVal); ok {
		return x.StringVal
	}
	return ""
}

func (x *Operand) GetSelector() *Selector {
	if x, ok := x.GetValue().(*Operand_Selector); ok {
		return x.Selector
	}
	return nil
}

type isOperand_Value interface {
	isOperand_Value()
}

type Operand_FloatVal struct {
	FloatVal float64 `protobuf:"fixed64,1,opt,name=float_val,json=floatVal,proto3,oneof"`
}

type Operand_StringVal struct {
	StringVal string `protobuf:"bytes,2,opt,name=string_val,json=stringVal,proto3,oneof"`
}

type Operand_Selector struct {
	Selector *Selector `protobuf:"bytes,3,opt,name=selector,proto3,oneof"`
}

func (*Operand_FloatVal) isOperand_Value() {}

func (*Operand_StringVal) isOperand_Value() {}

func (*Operand_Selector) isOperand_Value() {}

type FilterOperator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilterType FilterType `protobuf:"varint,1,opt,name=filter_type,json=filterType,proto3,enum=prometheus.FilterType" json:"filter_type,omitempty"`
	Params     []string   `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *FilterOperator) Reset() {
	*x = FilterOperator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterOperator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterOperator) ProtoMessage() {}

func (x *FilterOperator) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterOperator.ProtoReflect.Descriptor instead.
func (*FilterOperator) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{3}
}

func (x *FilterOperator) GetFilterType() FilterType {
	if x != nil {
		return x.FilterType
	}
	return FilterType_LITERAL_OR
}

func (x *FilterOperator) GetParams() []string {
	if x != nil {
		return x.Params
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagKey    string            `protobuf:"bytes,1,opt,name=tag_key,json=tagKey,proto3" json:"tag_key,omitempty"`
	Operators []*FilterOperator `protobuf:"bytes,2,rep,name=operators,proto3" json:"operators,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{4}
}

func (x *Filter) GetTagKey() string {
	if x != nil {
		return x.TagKey
	}
	return ""
}

func (x *Filter) GetOperators() []*FilterOperator {
	if x != nil {
		return x.Operators
	}
	return nil
}

type Selector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Measurement string    `protobuf:"bytes,1,opt,name=measurement,proto3" json:"measurement,omitempty"`
	Filters     []*Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	Start       int64     `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	End         int64     `protobuf:"varint,4,opt,name=end,proto3" json:"end,omitempty"`
	AlignStart  int64     `protobuf:"varint,5,opt,name=align_start,json=alignStart,proto3" json:"align_start,omitempty"`
	AlignEnd    int64     `protobuf:"varint,6,opt,name=align_end,json=alignEnd,proto3" json:"align_end,omitempty"`
	Step        int64     `protobuf:"varint,7,opt,name=step,proto3" json:"step,omitempty"`
	Range       int64     `protobuf:"varint,8,opt,name=range,proto3" json:"range,omitempty"`
	Offset      int64     `protobuf:"varint,9,opt,name=offset,proto3" json:"offset,omitempty"`
	Field       string    `protobuf:"bytes,10,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *Selector) Reset() {
	*x = Selector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Selector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selector) ProtoMessage() {}

func (x *Selector) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selector.ProtoReflect.Descriptor instead.
func (*Selector) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{5}
}

func (x *Selector) GetMeasurement() string {
	if x != nil {
		return x.Measurement
	}
	return ""
}

func (x *Selector) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *Selector) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Selector) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *Selector) GetAlignStart() int64 {
	if x != nil {
		return x.AlignStart
	}
	return 0
}

func (x *Selector) GetAlignEnd() int64 {
	if x != nil {
		return x.AlignEnd
	}
	return 0
}

func (x *Selector) GetStep() int64 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *Selector) GetRange() int64 {
	if x != nil {
		return x.Range
	}
	return 0
}

func (x *Selector) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Selector) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp int64   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{6}
}

func (x *Sample) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Sample) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type TimeSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels  []*Label  `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	Samples []*Sample `protobuf:"bytes,2,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *TimeSeries) Reset() {
	*x = TimeSeries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeries) ProtoMessage() {}

func (x *TimeSeries) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeries.ProtoReflect.Descriptor instead.
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{7}
}

func (x *TimeSeries) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *TimeSeries) GetSamples() []*Sample {
	if x != nil {
		return x.Samples
	}
	return nil
}

type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{8}
}

func (x *Label) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Label) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Labels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels []*Label `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Labels) Reset() {
	*x = Labels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prometheus_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Labels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Labels) ProtoMessage() {}

func (x *Labels) ProtoReflect() protoreflect.Message {
	mi := &file_prometheus_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Labels.ProtoReflect.Descriptor instead.
func (*Labels) Descriptor() ([]byte, []int) {
	return file_prometheus_proto_rawDescGZIP(), []int{9}
}

func (x *Labels) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_prometheus_proto protoreflect.FileDescriptor

var file_prometheus_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x22, 0x71,
	0x0a, 0x04, 0x45, 0x78, 0x70, 0x72, 0x12, 0x2f, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74,
	0x68, 0x65, 0x75, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x07,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x65,
	0x78, 0x70, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x45, 0x78, 0x70, 0x72, 0x48, 0x00,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x45, 0x78, 0x70, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x4e, 0x6f, 0x64,
	0x65, 0x22, 0xee, 0x01, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x45, 0x78, 0x70, 0x72, 0x12, 0x39, 0x0a,
	0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x45,
	0x78, 0x70, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68,
	0x65, 0x75, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e,
	0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x69, 0x74, 0x68,
	0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x77, 0x69, 0x74, 0x68, 0x6f,
	0x75, 0x74, 0x22, 0x2e, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x47, 0x47, 0x52, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x55, 0x4e, 0x43, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59,
	0x10, 0x02, 0x22, 0x86, 0x01, 0x0a, 0x07, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x1d,
	0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x00, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x1f, 0x0a,
	0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x12, 0x32,
	0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x61, 0x0a, 0x0e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x0a,
	0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x5b,
	0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x67, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x67, 0x4b, 0x65,
	0x79, 0x12, 0x38, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75,
	0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x98, 0x02, 0x0a, 0x08,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x45, 0x6e, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x74,
	0x65, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x3c, 0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x65, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2c, 0x0a,
	0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0x31, 0x0a, 0x05, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x33,
	0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x2a, 0x52, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x4f, 0x52, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x41, 0x4c,
	0x5f, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x47, 0x45, 0x58, 0x50, 0x10,
	0x02, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x45, 0x58, 0x50, 0x5f,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x03, 0x42, 0x5f, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x63, 0x65,
	0x72, 0x65, 0x73, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x42, 0x0a, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x65, 0x72,
	0x65, 0x73, 0x44, 0x42, 0x2f, 0x63, 0x65, 0x72, 0x65, 0x73, 0x64, 0x62, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x65, 0x72,
	0x65, 0x73, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prometheus_proto_rawDescOnce sync.Once
	file_prometheus_proto_rawDescData = file_prometheus_proto_rawDesc
)

func file_prometheus_proto_rawDescGZIP() []byte {
	file_prometheus_proto_rawDescOnce.Do(func() {
		file_prometheus_proto_rawDescData = protoimpl.X.CompressGZIP(file_prometheus_proto_rawDescData)
	})
	return file_prometheus_proto_rawDescData
}

var file_prometheus_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_prometheus_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_prometheus_proto_goTypes = []interface{}{
	(FilterType)(0),           // 0: prometheus.FilterType
	(SubExpr_OperatorType)(0), // 1: prometheus.SubExpr.OperatorType
	(*Expr)(nil),              // 2: prometheus.Expr
	(*SubExpr)(nil),           // 3: prometheus.SubExpr
	(*Operand)(nil),           // 4: prometheus.Operand
	(*FilterOperator)(nil),    // 5: prometheus.FilterOperator
	(*Filter)(nil),            // 6: prometheus.Filter
	(*Selector)(nil),          // 7: prometheus.Selector
	(*Sample)(nil),            // 8: prometheus.Sample
	(*TimeSeries)(nil),        // 9: prometheus.TimeSeries
	(*Label)(nil),             // 10: prometheus.Label
	(*Labels)(nil),            // 11: prometheus.Labels
}
var file_prometheus_proto_depIdxs = []int32{
	4,  // 0: prometheus.Expr.operand:type_name -> prometheus.Operand
	3,  // 1: prometheus.Expr.sub_expr:type_name -> prometheus.SubExpr
	1,  // 2: prometheus.SubExpr.op_type:type_name -> prometheus.SubExpr.OperatorType
	2,  // 3: prometheus.SubExpr.operands:type_name -> prometheus.Expr
	7,  // 4: prometheus.Operand.selector:type_name -> prometheus.Selector
	0,  // 5: prometheus.FilterOperator.filter_type:type_name -> prometheus.FilterType
	5,  // 6: prometheus.Filter.operators:type_name -> prometheus.FilterOperator
	6,  // 7: prometheus.Selector.filters:type_name -> prometheus.Filter
	10, // 8: prometheus.TimeSeries.labels:type_name -> prometheus.Label
	8,  // 9: prometheus.TimeSeries.samples:type_name -> prometheus.Sample
	10, // 10: prometheus.Labels.labels:type_name -> prometheus.Label
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_prometheus_proto_init() }
func file_prometheus_proto_init() {
	if File_prometheus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prometheus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expr); i {
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
		file_prometheus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubExpr); i {
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
		file_prometheus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operand); i {
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
		file_prometheus_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterOperator); i {
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
		file_prometheus_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_prometheus_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Selector); i {
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
		file_prometheus_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
		file_prometheus_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeries); i {
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
		file_prometheus_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
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
		file_prometheus_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Labels); i {
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
	file_prometheus_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Expr_Operand)(nil),
		(*Expr_SubExpr)(nil),
	}
	file_prometheus_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Operand_FloatVal)(nil),
		(*Operand_StringVal)(nil),
		(*Operand_Selector)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prometheus_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prometheus_proto_goTypes,
		DependencyIndexes: file_prometheus_proto_depIdxs,
		EnumInfos:         file_prometheus_proto_enumTypes,
		MessageInfos:      file_prometheus_proto_msgTypes,
	}.Build()
	File_prometheus_proto = out.File
	file_prometheus_proto_rawDesc = nil
	file_prometheus_proto_goTypes = nil
	file_prometheus_proto_depIdxs = nil
}
