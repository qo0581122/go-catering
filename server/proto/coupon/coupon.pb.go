// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coupon.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	product "catering/proto/product"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UseCouponReq struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	CouponId             int64    `protobuf:"varint,2,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id,omitempty"`
	UserCouponId         int64    `protobuf:"varint,3,opt,name=user_coupon_id,json=userCouponId,proto3" json:"user_coupon_id,omitempty"`
	OrderId              int64    `protobuf:"varint,4,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-,optional"`
	XXX_unrecognized     []byte   `json:"-,optional"`
	XXX_sizecache        int32    `json:"-,optional"`
}

func (m *UseCouponReq) Reset()         { *m = UseCouponReq{} }
func (m *UseCouponReq) String() string { return proto.CompactTextString(m) }
func (*UseCouponReq) ProtoMessage()    {}
func (*UseCouponReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a727a1a30518ca78, []int{0}
}

func (m *UseCouponReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UseCouponReq.Unmarshal(m, b)
}
func (m *UseCouponReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UseCouponReq.Marshal(b, m, deterministic)
}
func (m *UseCouponReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UseCouponReq.Merge(m, src)
}
func (m *UseCouponReq) XXX_Size() int {
	return xxx_messageInfo_UseCouponReq.Size(m)
}
func (m *UseCouponReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UseCouponReq.DiscardUnknown(m)
}

var xxx_messageInfo_UseCouponReq proto.InternalMessageInfo

func (m *UseCouponReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UseCouponReq) GetCouponId() int64 {
	if m != nil {
		return m.CouponId
	}
	return 0
}

func (m *UseCouponReq) GetUserCouponId() int64 {
	if m != nil {
		return m.UserCouponId
	}
	return 0
}

func (m *UseCouponReq) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type UseCouponResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-,optional"`
	XXX_unrecognized     []byte   `json:"-,optional"`
	XXX_sizecache        int32    `json:"-,optional"`
}

func (m *UseCouponResp) Reset()         { *m = UseCouponResp{} }
func (m *UseCouponResp) String() string { return proto.CompactTextString(m) }
func (*UseCouponResp) ProtoMessage()    {}
func (*UseCouponResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a727a1a30518ca78, []int{1}
}

func (m *UseCouponResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UseCouponResp.Unmarshal(m, b)
}
func (m *UseCouponResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UseCouponResp.Marshal(b, m, deterministic)
}
func (m *UseCouponResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UseCouponResp.Merge(m, src)
}
func (m *UseCouponResp) XXX_Size() int {
	return xxx_messageInfo_UseCouponResp.Size(m)
}
func (m *UseCouponResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UseCouponResp.DiscardUnknown(m)
}

var xxx_messageInfo_UseCouponResp proto.InternalMessageInfo

type GetUserCouponReq struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Begin                int32    `protobuf:"varint,2,opt,name=begin,proto3" json:"begin,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-,optional"`
	XXX_unrecognized     []byte   `json:"-,optional"`
	XXX_sizecache        int32    `json:"-,optional"`
}

func (m *GetUserCouponReq) Reset()         { *m = GetUserCouponReq{} }
func (m *GetUserCouponReq) String() string { return proto.CompactTextString(m) }
func (*GetUserCouponReq) ProtoMessage()    {}
func (*GetUserCouponReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a727a1a30518ca78, []int{2}
}

func (m *GetUserCouponReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserCouponReq.Unmarshal(m, b)
}
func (m *GetUserCouponReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserCouponReq.Marshal(b, m, deterministic)
}
func (m *GetUserCouponReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserCouponReq.Merge(m, src)
}
func (m *GetUserCouponReq) XXX_Size() int {
	return xxx_messageInfo_GetUserCouponReq.Size(m)
}
func (m *GetUserCouponReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserCouponReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserCouponReq proto.InternalMessageInfo

func (m *GetUserCouponReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *GetUserCouponReq) GetBegin() int32 {
	if m != nil {
		return m.Begin
	}
	return 0
}

func (m *GetUserCouponReq) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetUserCouponResp struct {
	Coupons              []*UserCoupon `protobuf:"bytes,1,rep,name=coupons,proto3" json:"coupons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-,optional"`
	XXX_unrecognized     []byte        `json:"-,optional"`
	XXX_sizecache        int32         `json:"-,optional"`
}

func (m *GetUserCouponResp) Reset()         { *m = GetUserCouponResp{} }
func (m *GetUserCouponResp) String() string { return proto.CompactTextString(m) }
func (*GetUserCouponResp) ProtoMessage()    {}
func (*GetUserCouponResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a727a1a30518ca78, []int{3}
}

func (m *GetUserCouponResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserCouponResp.Unmarshal(m, b)
}
func (m *GetUserCouponResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserCouponResp.Marshal(b, m, deterministic)
}
func (m *GetUserCouponResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserCouponResp.Merge(m, src)
}
func (m *GetUserCouponResp) XXX_Size() int {
	return xxx_messageInfo_GetUserCouponResp.Size(m)
}
func (m *GetUserCouponResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserCouponResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserCouponResp proto.InternalMessageInfo

func (m *GetUserCouponResp) GetCoupons() []*UserCoupon {
	if m != nil {
		return m.Coupons
	}
	return nil
}

type DrawCouponReq struct {
	CouponId             int64    `protobuf:"varint,1,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id,omitempty"`
	Uid                  int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	GetType              int32    `protobuf:"varint,4,opt,name=get_type,json=getType,proto3" json:"get_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-,optional"`
	XXX_unrecognized     []byte   `json:"-,optional"`
	XXX_sizecache        int32    `json:"-,optional"`
}

func (m *DrawCouponReq) Reset()         { *m = DrawCouponReq{} }
func (m *DrawCouponReq) String() string { return proto.CompactTextString(m) }
func (*DrawCouponReq) ProtoMessage()    {}
func (*DrawCouponReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a727a1a30518ca78, []int{4}
}

func (m *DrawCouponReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DrawCouponReq.Unmarshal(m, b)
}
func (m *DrawCouponReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DrawCouponReq.Marshal(b, m, deterministic)
}
func (m *DrawCouponReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DrawCouponReq.Merge(m, src)
}
func (m *DrawCouponReq) XXX_Size() int {
	return xxx_messageInfo_DrawCouponReq.Size(m)
}
func (m *DrawCouponReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DrawCouponReq.DiscardUnknown(m)
}

var xxx_messageInfo_DrawCouponReq proto.InternalMessageInfo

func (m *DrawCouponReq) GetCouponId() int64 {
	if m != nil {
		return m.CouponId
	}
	return 0
}

func (m *DrawCouponReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *DrawCouponReq) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *DrawCouponReq) GetGetType() int32 {
	if m != nil {
		return m.GetType
	}
	return 0
}

type DrawCouponResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-,optional"`
	XXX_unrecognized     []byte   `json:"-,optional"`
	XXX_sizecache        int32    `json:"-,optional"`
}

func (m *DrawCouponResp) Reset()         { *m = DrawCouponResp{} }
func (m *DrawCouponResp) String() string { return proto.CompactTextString(m) }
func (*DrawCouponResp) ProtoMessage()    {}
func (*DrawCouponResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a727a1a30518ca78, []int{5}
}

func (m *DrawCouponResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DrawCouponResp.Unmarshal(m, b)
}
func (m *DrawCouponResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DrawCouponResp.Marshal(b, m, deterministic)
}
func (m *DrawCouponResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DrawCouponResp.Merge(m, src)
}
func (m *DrawCouponResp) XXX_Size() int {
	return xxx_messageInfo_DrawCouponResp.Size(m)
}
func (m *DrawCouponResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DrawCouponResp.DiscardUnknown(m)
}

var xxx_messageInfo_DrawCouponResp proto.InternalMessageInfo

type GetCouponInfoReq struct {
	CouponId             int64    `protobuf:"varint,1,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-,optional"`
	XXX_unrecognized     []byte   `json:"-,optional"`
	XXX_sizecache        int32    `json:"-,optional"`
}

func (m *GetCouponInfoReq) Reset()         { *m = GetCouponInfoReq{} }
func (m *GetCouponInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetCouponInfoReq) ProtoMessage()    {}
func (*GetCouponInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a727a1a30518ca78, []int{6}
}

func (m *GetCouponInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCouponInfoReq.Unmarshal(m, b)
}
func (m *GetCouponInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCouponInfoReq.Marshal(b, m, deterministic)
}
func (m *GetCouponInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCouponInfoReq.Merge(m, src)
}
func (m *GetCouponInfoReq) XXX_Size() int {
	return xxx_messageInfo_GetCouponInfoReq.Size(m)
}
func (m *GetCouponInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCouponInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetCouponInfoReq proto.InternalMessageInfo

func (m *GetCouponInfoReq) GetCouponId() int64 {
	if m != nil {
		return m.CouponId
	}
	return 0
}

type GetCouponInfoResp struct {
	Coupon               *Coupon  `protobuf:"bytes,1,opt,name=coupon,proto3" json:"coupon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-,optional"`
	XXX_unrecognized     []byte   `json:"-,optional"`
	XXX_sizecache        int32    `json:"-,optional"`
}

func (m *GetCouponInfoResp) Reset()         { *m = GetCouponInfoResp{} }
func (m *GetCouponInfoResp) String() string { return proto.CompactTextString(m) }
func (*GetCouponInfoResp) ProtoMessage()    {}
func (*GetCouponInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a727a1a30518ca78, []int{7}
}

func (m *GetCouponInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCouponInfoResp.Unmarshal(m, b)
}
func (m *GetCouponInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCouponInfoResp.Marshal(b, m, deterministic)
}
func (m *GetCouponInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCouponInfoResp.Merge(m, src)
}
func (m *GetCouponInfoResp) XXX_Size() int {
	return xxx_messageInfo_GetCouponInfoResp.Size(m)
}
func (m *GetCouponInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCouponInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetCouponInfoResp proto.InternalMessageInfo

func (m *GetCouponInfoResp) GetCoupon() *Coupon {
	if m != nil {
		return m.Coupon
	}
	return nil
}

type Coupon struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CouponName           string               `protobuf:"bytes,2,opt,name=coupon_name,json=couponName,proto3" json:"coupon_name,omitempty"`
	LeastUsePrice        int32                `protobuf:"varint,3,opt,name=least_use_price,json=leastUsePrice,proto3" json:"least_use_price,omitempty"`
	Price                int32                `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	PicUrl               string               `protobuf:"bytes,5,opt,name=pic_url,json=picUrl,proto3" json:"pic_url,omitempty"`
	UseType              int32                `protobuf:"varint,6,opt,name=use_type,json=useType,proto3" json:"use_type,omitempty"`
	GetType              int32                `protobuf:"varint,7,opt,name=get_type,json=getType,proto3" json:"get_type,omitempty"`
	UseBeginTime         string               `protobuf:"bytes,8,opt,name=use_begin_time,json=useBeginTime,proto3" json:"use_begin_time,omitempty"`
	UseEndTime           string               `protobuf:"bytes,9,opt,name=use_end_time,json=useEndTime,proto3" json:"use_end_time,omitempty"`
	ValidTimeType        int32                `protobuf:"varint,10,opt,name=valid_time_type,json=validTimeType,proto3" json:"valid_time_type,omitempty"`
	ValidDay             int32                `protobuf:"varint,11,opt,name=valid_day,json=validDay,proto3" json:"valid_day,omitempty"`
	GetBeginTime         string               `protobuf:"bytes,12,opt,name=get_begin_time,json=getBeginTime,proto3" json:"get_begin_time,omitempty"`
	GetEndTime           string               `protobuf:"bytes,13,opt,name=get_end_time,json=getEndTime,proto3" json:"get_end_time,omitempty"`
	GetCount             int32                `protobuf:"varint,14,opt,name=get_count,json=getCount,proto3" json:"get_count,omitempty"`
	Description          string               `protobuf:"bytes,15,opt,name=description,proto3" json:"description,omitempty"`
	Product              *product.ProductBase `protobuf:"bytes,16,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-,optional"`
	XXX_unrecognized     []byte               `json:"-,optional"`
	XXX_sizecache        int32                `json:"-,optional"`
}

func (m *Coupon) Reset()         { *m = Coupon{} }
func (m *Coupon) String() string { return proto.CompactTextString(m) }
func (*Coupon) ProtoMessage()    {}
func (*Coupon) Descriptor() ([]byte, []int) {
	return fileDescriptor_a727a1a30518ca78, []int{8}
}

func (m *Coupon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coupon.Unmarshal(m, b)
}
func (m *Coupon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coupon.Marshal(b, m, deterministic)
}
func (m *Coupon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coupon.Merge(m, src)
}
func (m *Coupon) XXX_Size() int {
	return xxx_messageInfo_Coupon.Size(m)
}
func (m *Coupon) XXX_DiscardUnknown() {
	xxx_messageInfo_Coupon.DiscardUnknown(m)
}

var xxx_messageInfo_Coupon proto.InternalMessageInfo

func (m *Coupon) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Coupon) GetCouponName() string {
	if m != nil {
		return m.CouponName
	}
	return ""
}

func (m *Coupon) GetLeastUsePrice() int32 {
	if m != nil {
		return m.LeastUsePrice
	}
	return 0
}

func (m *Coupon) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Coupon) GetPicUrl() string {
	if m != nil {
		return m.PicUrl
	}
	return ""
}

func (m *Coupon) GetUseType() int32 {
	if m != nil {
		return m.UseType
	}
	return 0
}

func (m *Coupon) GetGetType() int32 {
	if m != nil {
		return m.GetType
	}
	return 0
}

func (m *Coupon) GetUseBeginTime() string {
	if m != nil {
		return m.UseBeginTime
	}
	return ""
}

func (m *Coupon) GetUseEndTime() string {
	if m != nil {
		return m.UseEndTime
	}
	return ""
}

func (m *Coupon) GetValidTimeType() int32 {
	if m != nil {
		return m.ValidTimeType
	}
	return 0
}

func (m *Coupon) GetValidDay() int32 {
	if m != nil {
		return m.ValidDay
	}
	return 0
}

func (m *Coupon) GetGetBeginTime() string {
	if m != nil {
		return m.GetBeginTime
	}
	return ""
}

func (m *Coupon) GetGetEndTime() string {
	if m != nil {
		return m.GetEndTime
	}
	return ""
}

func (m *Coupon) GetGetCount() int32 {
	if m != nil {
		return m.GetCount
	}
	return 0
}

func (m *Coupon) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Coupon) GetProduct() *product.ProductBase {
	if m != nil {
		return m.Product
	}
	return nil
}

type UserCoupon struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GetTime              string   `protobuf:"bytes,2,opt,name=get_time,json=getTime,proto3" json:"get_time,omitempty"`
	UseBeginTime         string   `protobuf:"bytes,3,opt,name=use_begin_time,json=useBeginTime,proto3" json:"use_begin_time,omitempty"`
	UseEndTime           string   `protobuf:"bytes,4,opt,name=use_end_time,json=useEndTime,proto3" json:"use_end_time,omitempty"`
	UseStatus            int32    `protobuf:"varint,5,opt,name=use_status,json=useStatus,proto3" json:"use_status,omitempty"`
	GetType              int32    `protobuf:"varint,6,opt,name=get_type,json=getType,proto3" json:"get_type,omitempty"`
	Coupon               *Coupon  `protobuf:"bytes,7,opt,name=coupon,proto3" json:"coupon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-,optional"`
	XXX_unrecognized     []byte   `json:"-,optional"`
	XXX_sizecache        int32    `json:"-,optional"`
}

func (m *UserCoupon) Reset()         { *m = UserCoupon{} }
func (m *UserCoupon) String() string { return proto.CompactTextString(m) }
func (*UserCoupon) ProtoMessage()    {}
func (*UserCoupon) Descriptor() ([]byte, []int) {
	return fileDescriptor_a727a1a30518ca78, []int{9}
}

func (m *UserCoupon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCoupon.Unmarshal(m, b)
}
func (m *UserCoupon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCoupon.Marshal(b, m, deterministic)
}
func (m *UserCoupon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCoupon.Merge(m, src)
}
func (m *UserCoupon) XXX_Size() int {
	return xxx_messageInfo_UserCoupon.Size(m)
}
func (m *UserCoupon) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCoupon.DiscardUnknown(m)
}

var xxx_messageInfo_UserCoupon proto.InternalMessageInfo

func (m *UserCoupon) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserCoupon) GetGetTime() string {
	if m != nil {
		return m.GetTime
	}
	return ""
}

func (m *UserCoupon) GetUseBeginTime() string {
	if m != nil {
		return m.UseBeginTime
	}
	return ""
}

func (m *UserCoupon) GetUseEndTime() string {
	if m != nil {
		return m.UseEndTime
	}
	return ""
}

func (m *UserCoupon) GetUseStatus() int32 {
	if m != nil {
		return m.UseStatus
	}
	return 0
}

func (m *UserCoupon) GetGetType() int32 {
	if m != nil {
		return m.GetType
	}
	return 0
}

func (m *UserCoupon) GetCoupon() *Coupon {
	if m != nil {
		return m.Coupon
	}
	return nil
}

func init() {
	proto.RegisterType((*UseCouponReq)(nil), "proto.UseCouponReq")
	proto.RegisterType((*UseCouponResp)(nil), "proto.UseCouponResp")
	proto.RegisterType((*GetUserCouponReq)(nil), "proto.GetUserCouponReq")
	proto.RegisterType((*GetUserCouponResp)(nil), "proto.GetUserCouponResp")
	proto.RegisterType((*DrawCouponReq)(nil), "proto.DrawCouponReq")
	proto.RegisterType((*DrawCouponResp)(nil), "proto.DrawCouponResp")
	proto.RegisterType((*GetCouponInfoReq)(nil), "proto.GetCouponInfoReq")
	proto.RegisterType((*GetCouponInfoResp)(nil), "proto.GetCouponInfoResp")
	proto.RegisterType((*Coupon)(nil), "proto.Coupon")
	proto.RegisterType((*UserCoupon)(nil), "proto.UserCoupon")
}

func init() { proto.RegisterFile("coupon.proto", fileDescriptor_a727a1a30518ca78) }

var fileDescriptor_a727a1a30518ca78 = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x55, 0x92, 0x3a, 0x8e, 0x27, 0x71, 0x92, 0x2e, 0xad, 0x6a, 0x5a, 0x21, 0x22, 0x0b, 0xaa,
	0x4a, 0xa0, 0x56, 0x2a, 0x12, 0x12, 0x9c, 0x50, 0x5b, 0x04, 0xbd, 0xa0, 0xca, 0x34, 0x67, 0xcb,
	0xb5, 0x87, 0xc8, 0x52, 0x62, 0x2f, 0xde, 0x35, 0x28, 0x17, 0xce, 0x88, 0xaf, 0xe4, 0x53, 0xd0,
	0xce, 0x6e, 0x62, 0x3b, 0x89, 0x80, 0x53, 0xb2, 0x6f, 0xde, 0xce, 0xbc, 0x79, 0x9e, 0x1d, 0x18,
	0xc4, 0x79, 0xc9, 0xf3, 0xec, 0x9c, 0x17, 0xb9, 0xcc, 0x99, 0x45, 0x3f, 0xc7, 0x87, 0xbc, 0xc8,
	0x93, 0x32, 0x96, 0x17, 0xe6, 0x57, 0x47, 0xfd, 0x1f, 0x30, 0x98, 0x0a, 0xbc, 0xa6, 0x0b, 0x01,
	0x7e, 0x65, 0x63, 0xe8, 0x94, 0x69, 0xe2, 0xb5, 0x26, 0xad, 0xb3, 0x4e, 0xa0, 0xfe, 0xb2, 0x13,
	0x70, 0x74, 0xbe, 0x30, 0x4d, 0xbc, 0x36, 0xe1, 0x3d, 0x0d, 0xdc, 0x26, 0xec, 0x19, 0x0c, 0x4b,
	0x81, 0x45, 0x58, 0x31, 0x3a, 0xc4, 0x18, 0x28, 0xf4, 0x7a, 0xc5, 0x7a, 0x0c, 0xbd, 0xbc, 0x48,
	0xb0, 0x50, 0xf1, 0x3d, 0x8a, 0xdb, 0x74, 0xbe, 0x4d, 0xfc, 0x11, 0xb8, 0xb5, 0xfa, 0x82, 0xfb,
	0x77, 0x30, 0xfe, 0x80, 0x72, 0xba, 0xbe, 0xbe, 0x5b, 0xd4, 0x01, 0x58, 0x0f, 0x38, 0x4b, 0x33,
	0x12, 0x64, 0x05, 0xfa, 0xa0, 0xd0, 0x38, 0x2f, 0x33, 0x49, 0x22, 0xac, 0x40, 0x1f, 0xfc, 0x77,
	0xb0, 0xbf, 0x91, 0x51, 0x70, 0xf6, 0x02, 0x6c, 0xad, 0x59, 0x78, 0xad, 0x49, 0xe7, 0xac, 0x7f,
	0xb9, 0xaf, 0x0d, 0x39, 0xaf, 0xf1, 0x56, 0x0c, 0x3f, 0x07, 0xf7, 0xa6, 0x88, 0xbe, 0x57, 0x82,
	0x1a, 0x9e, 0xb4, 0x36, 0x3c, 0x31, 0x6a, 0xdb, 0x0d, 0xb5, 0xdb, 0xba, 0x94, 0x2b, 0x33, 0x94,
	0xa1, 0x5c, 0x72, 0x24, 0x57, 0xac, 0xc0, 0x9e, 0xa1, 0xbc, 0x5f, 0x72, 0xf4, 0xc7, 0x30, 0xac,
	0x17, 0x14, 0xdc, 0xbf, 0x20, 0x5b, 0x8c, 0xa3, 0xd9, 0x97, 0xfc, 0x5f, 0x2a, 0xfc, 0xb7, 0xd4,
	0x75, 0xfd, 0x82, 0xe0, 0xec, 0x39, 0x74, 0x35, 0x81, 0xe8, 0xfd, 0x4b, 0xd7, 0x34, 0x6d, 0x0a,
	0x99, 0xa0, 0xff, 0x73, 0x0f, 0xba, 0x1a, 0x62, 0x43, 0x68, 0xaf, 0x93, 0xb7, 0xd3, 0x84, 0x3d,
	0x85, 0xbe, 0xa9, 0x99, 0x45, 0x0b, 0xa4, 0x26, 0x9d, 0x00, 0x34, 0xf4, 0x29, 0x5a, 0x20, 0x3b,
	0x85, 0xd1, 0x1c, 0x23, 0x21, 0xc3, 0x52, 0x60, 0xc8, 0x8b, 0x34, 0x46, 0xd3, 0xb5, 0x4b, 0xf0,
	0x54, 0xe0, 0x9d, 0x02, 0x95, 0x27, 0x3a, 0xaa, 0x5b, 0xd7, 0x07, 0x76, 0x04, 0x36, 0x4f, 0xe3,
	0xb0, 0x2c, 0xe6, 0x9e, 0x45, 0xa9, 0xbb, 0x3c, 0x8d, 0xa7, 0xc5, 0x5c, 0x99, 0xa5, 0x12, 0x92,
	0x59, 0x5d, 0x6d, 0x56, 0x29, 0x50, 0x99, 0xd5, 0xf0, 0xd1, 0x6e, 0xf8, 0x68, 0xc6, 0x33, 0xa4,
	0xe9, 0x08, 0x65, 0xba, 0x40, 0xaf, 0x47, 0x59, 0xd5, 0x78, 0x5e, 0x29, 0xf0, 0x3e, 0x5d, 0x20,
	0x9b, 0x80, 0x3a, 0x87, 0x98, 0x25, 0x9a, 0xe3, 0xe8, 0xa6, 0x4a, 0x81, 0xef, 0xb3, 0x84, 0x18,
	0xa7, 0x30, 0xfa, 0x16, 0xcd, 0x53, 0x1d, 0xd7, 0x95, 0x40, 0x37, 0x45, 0xb0, 0xe2, 0x50, 0xbd,
	0x13, 0x70, 0x34, 0x2f, 0x89, 0x96, 0x5e, 0x9f, 0x18, 0x3d, 0x02, 0x6e, 0xa2, 0xa5, 0x12, 0xa3,
	0x74, 0xd6, 0xc4, 0x0c, 0xb4, 0x98, 0x19, 0xca, 0x86, 0x18, 0xc5, 0x5a, 0x8b, 0x71, 0xb5, 0x98,
	0x19, 0xca, 0x95, 0x98, 0x13, 0x70, 0x14, 0x43, 0x4f, 0xd4, 0x50, 0x17, 0x99, 0xd1, 0xa7, 0xce,
	0x24, 0x9b, 0x40, 0x3f, 0x41, 0x11, 0x17, 0x29, 0x97, 0x69, 0x9e, 0x79, 0x23, 0xba, 0x5d, 0x87,
	0xd8, 0x4b, 0xb0, 0xcd, 0x0a, 0xf0, 0xc6, 0x34, 0x04, 0xcc, 0x0c, 0xc1, 0x9d, 0x46, 0xaf, 0x22,
	0x81, 0xc1, 0x8a, 0xe2, 0xff, 0x6e, 0x01, 0x54, 0x4f, 0x62, 0x6b, 0x1c, 0x56, 0xde, 0xa7, 0xeb,
	0x59, 0x20, 0xef, 0x95, 0xcc, 0x6d, 0xef, 0x3b, 0xff, 0xe1, 0xfd, 0xde, 0x96, 0xf7, 0x4f, 0x40,
	0x9d, 0x42, 0x21, 0x23, 0x59, 0x0a, 0x9a, 0x0a, 0x2b, 0x70, 0x4a, 0x81, 0x9f, 0x09, 0x68, 0x7c,
	0xfd, 0x6e, 0xf3, 0xeb, 0x57, 0xd3, 0x6e, 0xff, 0x65, 0xda, 0x2f, 0x7f, 0xb5, 0xc1, 0xd5, 0xd0,
	0xc7, 0x28, 0x4b, 0xe6, 0x58, 0xb0, 0x2b, 0x70, 0x1b, 0x1b, 0x83, 0x1d, 0x99, 0x9b, 0x9b, 0x9b,
	0xe9, 0xd8, 0xdb, 0x1d, 0x10, 0x9c, 0xbd, 0x01, 0xa8, 0x9e, 0x30, 0x3b, 0x30, 0xbc, 0xc6, 0x1a,
	0x39, 0x3e, 0xdc, 0x81, 0x0a, 0x6e, 0xca, 0x57, 0x4f, 0xb7, 0x5e, 0xbe, 0xb1, 0x01, 0xea, 0xe5,
	0x37, 0x5e, 0xfa, 0x6b, 0x70, 0xd6, 0x7b, 0x95, 0x3d, 0xaa, 0x76, 0x5b, 0x55, 0xfc, 0x60, 0x1b,
	0x14, 0xfc, 0xa1, 0x4b, 0xe0, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xb6, 0x39, 0x2f,
	0x44, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CouponHandlerClient is the client API for CouponHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CouponHandlerClient interface {
	GetUserCoupon(ctx context.Context, in *GetUserCouponReq, opts ...grpc.CallOption) (*GetUserCouponResp, error)
	DrawCoupon(ctx context.Context, in *DrawCouponReq, opts ...grpc.CallOption) (*DrawCouponResp, error)
	GetCouponInfo(ctx context.Context, in *GetCouponInfoReq, opts ...grpc.CallOption) (*GetCouponInfoResp, error)
	UseCoupon(ctx context.Context, in *UseCouponReq, opts ...grpc.CallOption) (*UseCouponResp, error)
}

type couponHandlerClient struct {
	cc *grpc.ClientConn
}

func NewCouponHandlerClient(cc *grpc.ClientConn) CouponHandlerClient {
	return &couponHandlerClient{cc}
}

func (c *couponHandlerClient) GetUserCoupon(ctx context.Context, in *GetUserCouponReq, opts ...grpc.CallOption) (*GetUserCouponResp, error) {
	out := new(GetUserCouponResp)
	err := c.cc.Invoke(ctx, "/proto.CouponHandler/GetUserCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponHandlerClient) DrawCoupon(ctx context.Context, in *DrawCouponReq, opts ...grpc.CallOption) (*DrawCouponResp, error) {
	out := new(DrawCouponResp)
	err := c.cc.Invoke(ctx, "/proto.CouponHandler/DrawCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponHandlerClient) GetCouponInfo(ctx context.Context, in *GetCouponInfoReq, opts ...grpc.CallOption) (*GetCouponInfoResp, error) {
	out := new(GetCouponInfoResp)
	err := c.cc.Invoke(ctx, "/proto.CouponHandler/GetCouponInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponHandlerClient) UseCoupon(ctx context.Context, in *UseCouponReq, opts ...grpc.CallOption) (*UseCouponResp, error) {
	out := new(UseCouponResp)
	err := c.cc.Invoke(ctx, "/proto.CouponHandler/UseCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CouponHandlerServer is the server API for CouponHandler service.
type CouponHandlerServer interface {
	GetUserCoupon(context.Context, *GetUserCouponReq) (*GetUserCouponResp, error)
	DrawCoupon(context.Context, *DrawCouponReq) (*DrawCouponResp, error)
	GetCouponInfo(context.Context, *GetCouponInfoReq) (*GetCouponInfoResp, error)
	UseCoupon(context.Context, *UseCouponReq) (*UseCouponResp, error)
}

// UnimplementedCouponHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedCouponHandlerServer struct {
}

func (*UnimplementedCouponHandlerServer) GetUserCoupon(ctx context.Context, req *GetUserCouponReq) (*GetUserCouponResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCoupon not implemented")
}
func (*UnimplementedCouponHandlerServer) DrawCoupon(ctx context.Context, req *DrawCouponReq) (*DrawCouponResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DrawCoupon not implemented")
}
func (*UnimplementedCouponHandlerServer) GetCouponInfo(ctx context.Context, req *GetCouponInfoReq) (*GetCouponInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouponInfo not implemented")
}
func (*UnimplementedCouponHandlerServer) UseCoupon(ctx context.Context, req *UseCouponReq) (*UseCouponResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseCoupon not implemented")
}

func RegisterCouponHandlerServer(s *grpc.Server, srv CouponHandlerServer) {
	s.RegisterService(&_CouponHandler_serviceDesc, srv)
}

func _CouponHandler_GetUserCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCouponReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponHandlerServer).GetUserCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CouponHandler/GetUserCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponHandlerServer).GetUserCoupon(ctx, req.(*GetUserCouponReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CouponHandler_DrawCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DrawCouponReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponHandlerServer).DrawCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CouponHandler/DrawCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponHandlerServer).DrawCoupon(ctx, req.(*DrawCouponReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CouponHandler_GetCouponInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCouponInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponHandlerServer).GetCouponInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CouponHandler/GetCouponInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponHandlerServer).GetCouponInfo(ctx, req.(*GetCouponInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CouponHandler_UseCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UseCouponReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponHandlerServer).UseCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CouponHandler/UseCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponHandlerServer).UseCoupon(ctx, req.(*UseCouponReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CouponHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CouponHandler",
	HandlerType: (*CouponHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserCoupon",
			Handler:    _CouponHandler_GetUserCoupon_Handler,
		},
		{
			MethodName: "DrawCoupon",
			Handler:    _CouponHandler_DrawCoupon_Handler,
		},
		{
			MethodName: "GetCouponInfo",
			Handler:    _CouponHandler_GetCouponInfo_Handler,
		},
		{
			MethodName: "UseCoupon",
			Handler:    _CouponHandler_UseCoupon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coupon.proto",
}
