// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: orders/orders.proto

package orders

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderRef string `protobuf:"bytes,1,opt,name=orderRef,proto3" json:"orderRef,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Notes    string `protobuf:"bytes,3,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orders_orders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_orders_orders_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderRef() string {
	if x != nil {
		return x.OrderRef
	}
	return ""
}

func (x *Order) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Order) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type OrderLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product  string `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Quantity string `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *OrderLine) Reset() {
	*x = OrderLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_orders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderLine) ProtoMessage() {}

func (x *OrderLine) ProtoReflect() protoreflect.Message {
	mi := &file_orders_orders_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderLine.ProtoReflect.Descriptor instead.
func (*OrderLine) Descriptor() ([]byte, []int) {
	return file_orders_orders_proto_rawDescGZIP(), []int{1}
}

func (x *OrderLine) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *OrderLine) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

type OrderConfirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID string  `protobuf:"bytes,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
	Total   float32 `protobuf:"fixed32,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *OrderConfirmation) Reset() {
	*x = OrderConfirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_orders_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderConfirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderConfirmation) ProtoMessage() {}

func (x *OrderConfirmation) ProtoReflect() protoreflect.Message {
	mi := &file_orders_orders_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderConfirmation.ProtoReflect.Descriptor instead.
func (*OrderConfirmation) Descriptor() ([]byte, []int) {
	return file_orders_orders_proto_rawDescGZIP(), []int{2}
}

func (x *OrderConfirmation) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *OrderConfirmation) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_orders_orders_proto protoreflect.FileDescriptor

var file_orders_orders_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x53, 0x0a,
	0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x22, 0x41, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x43, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x86, 0x01, 0x0a, 0x05, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x0d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x1a, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x65,
	0x73, 0x12, 0x11, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4c, 0x69, 0x6e, 0x65, 0x1a, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x28, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x6c, 0x69, 0x6e, 0x6e, 0x65, 0x77, 0x65, 0x6c, 0x6c, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_orders_proto_rawDescOnce sync.Once
	file_orders_orders_proto_rawDescData = file_orders_orders_proto_rawDesc
)

func file_orders_orders_proto_rawDescGZIP() []byte {
	file_orders_orders_proto_rawDescOnce.Do(func() {
		file_orders_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_orders_proto_rawDescData)
	})
	return file_orders_orders_proto_rawDescData
}

var file_orders_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orders_orders_proto_goTypes = []interface{}{
	(*Order)(nil),             // 0: orders.Order
	(*OrderLine)(nil),         // 1: orders.OrderLine
	(*OrderConfirmation)(nil), // 2: orders.OrderConfirmation
}
var file_orders_orders_proto_depIdxs = []int32{
	0, // 0: orders.Store.PlaceOrder:input_type -> orders.Order
	1, // 1: orders.Store.PlaceOrderLines:input_type -> orders.OrderLine
	2, // 2: orders.Store.PlaceOrder:output_type -> orders.OrderConfirmation
	2, // 3: orders.Store.PlaceOrderLines:output_type -> orders.OrderConfirmation
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orders_orders_proto_init() }
func file_orders_orders_proto_init() {
	if File_orders_orders_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orders_orders_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_orders_orders_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderLine); i {
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
		file_orders_orders_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderConfirmation); i {
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
			RawDescriptor: file_orders_orders_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_orders_proto_goTypes,
		DependencyIndexes: file_orders_orders_proto_depIdxs,
		MessageInfos:      file_orders_orders_proto_msgTypes,
	}.Build()
	File_orders_orders_proto = out.File
	file_orders_orders_proto_rawDesc = nil
	file_orders_orders_proto_goTypes = nil
	file_orders_orders_proto_depIdxs = nil
}
