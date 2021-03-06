// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: api/captcha.proto

package grpcIFace

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

type GetCaptchaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 验证码类型, 支持的有: audio, digit, string, chinese, match
	// 默认为string(字母+数字)
	Type string `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	// 验证码图片 宽度, 高度及字符个数
	// 默认值分别为: 150, 50, 5
	Width  int32 `protobuf:"varint,2,opt,name=Width,proto3" json:"Width,omitempty"`
	Height int32 `protobuf:"varint,3,opt,name=Height,proto3" json:"Height,omitempty"`
	Length int32 `protobuf:"varint,4,opt,name=Length,proto3" json:"Length,omitempty"`
	// 验证码有效期, 单位为秒, 默认为180秒
	MaxAge int64 `protobuf:"varint,5,opt,name=MaxAge,proto3" json:"MaxAge,omitempty"`
	// 以#rrggbbaa的方式指定, 均为十六进制, alpha的范围为:00-ff(00为全透明, ff为不透明)
	// 默认值为 #00000000
	BgColor string `protobuf:"bytes,6,opt,name=BgColor,proto3" json:"BgColor,omitempty"`
	// type为audio时, 需要设置语言, 默认为zh
	// 支持: "en", "ja", "ru", "zh"
	AudioLanguage string `protobuf:"bytes,7,opt,name=AudioLanguage,proto3" json:"AudioLanguage,omitempty"`
	// 噪音(Noise)就是在验证码文字底部添加一些干扰的文字, NoiseCount就是干扰文字的个数, 默认为50
	NoiseCount int32 `protobuf:"varint,10,opt,name=NoiseCount,proto3" json:"NoiseCount,omitempty"`
	// 干扰线设置, 默认为OptionsShowAllLines(请参考consts.go)
	ShowLineOptions int32 `protobuf:"varint,11,opt,name=ShowLineOptions,proto3" json:"ShowLineOptions,omitempty"`
	// 纯数字类型的验证码相关参数
	// 歪斜程度, 取值范围为(0到1), 默认为0
	DigitMaxSkew float64 `protobuf:"fixed64,12,opt,name=DigitMaxSkew,proto3" json:"DigitMaxSkew,omitempty"`
	// 干扰点(一些小点), 默认为0
	DigitDotCount int32 `protobuf:"varint,13,opt,name=DigitDotCount,proto3" json:"DigitDotCount,omitempty"`
	// 调试模式会返回answer，方便测试
	Debug bool `protobuf:"varint,14,opt,name=Debug,proto3" json:"Debug,omitempty"`
}

func (x *GetCaptchaRequest) Reset() {
	*x = GetCaptchaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_captcha_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCaptchaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCaptchaRequest) ProtoMessage() {}

func (x *GetCaptchaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_captcha_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCaptchaRequest.ProtoReflect.Descriptor instead.
func (*GetCaptchaRequest) Descriptor() ([]byte, []int) {
	return file_api_captcha_proto_rawDescGZIP(), []int{0}
}

func (x *GetCaptchaRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetCaptchaRequest) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *GetCaptchaRequest) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *GetCaptchaRequest) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *GetCaptchaRequest) GetMaxAge() int64 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *GetCaptchaRequest) GetBgColor() string {
	if x != nil {
		return x.BgColor
	}
	return ""
}

func (x *GetCaptchaRequest) GetAudioLanguage() string {
	if x != nil {
		return x.AudioLanguage
	}
	return ""
}

func (x *GetCaptchaRequest) GetNoiseCount() int32 {
	if x != nil {
		return x.NoiseCount
	}
	return 0
}

func (x *GetCaptchaRequest) GetShowLineOptions() int32 {
	if x != nil {
		return x.ShowLineOptions
	}
	return 0
}

func (x *GetCaptchaRequest) GetDigitMaxSkew() float64 {
	if x != nil {
		return x.DigitMaxSkew
	}
	return 0
}

func (x *GetCaptchaRequest) GetDigitDotCount() int32 {
	if x != nil {
		return x.DigitDotCount
	}
	return 0
}

func (x *GetCaptchaRequest) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

type GetCaptchaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Base64Value string `protobuf:"bytes,2,opt,name=Base64Value,proto3" json:"Base64Value,omitempty"`
	Answer      string `protobuf:"bytes,3,opt,name=Answer,proto3" json:"Answer,omitempty"`
}

func (x *GetCaptchaResponse) Reset() {
	*x = GetCaptchaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_captcha_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCaptchaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCaptchaResponse) ProtoMessage() {}

func (x *GetCaptchaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_captcha_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCaptchaResponse.ProtoReflect.Descriptor instead.
func (*GetCaptchaResponse) Descriptor() ([]byte, []int) {
	return file_api_captcha_proto_rawDescGZIP(), []int{1}
}

func (x *GetCaptchaResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *GetCaptchaResponse) GetBase64Value() string {
	if x != nil {
		return x.Base64Value
	}
	return ""
}

func (x *GetCaptchaResponse) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type VerifyCaptchaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Answer string `protobuf:"bytes,2,opt,name=Answer,proto3" json:"Answer,omitempty"`
}

func (x *VerifyCaptchaRequest) Reset() {
	*x = VerifyCaptchaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_captcha_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCaptchaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCaptchaRequest) ProtoMessage() {}

func (x *VerifyCaptchaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_captcha_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCaptchaRequest.ProtoReflect.Descriptor instead.
func (*VerifyCaptchaRequest) Descriptor() ([]byte, []int) {
	return file_api_captcha_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyCaptchaRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *VerifyCaptchaRequest) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type VerifyCaptchaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data bool `protobuf:"varint,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *VerifyCaptchaResponse) Reset() {
	*x = VerifyCaptchaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_captcha_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCaptchaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCaptchaResponse) ProtoMessage() {}

func (x *VerifyCaptchaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_captcha_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCaptchaResponse.ProtoReflect.Descriptor instead.
func (*VerifyCaptchaResponse) Descriptor() ([]byte, []int) {
	return file_api_captcha_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyCaptchaResponse) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

var File_api_captcha_proto protoreflect.FileDescriptor

var file_api_captcha_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x22, 0xef, 0x02, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x57, 0x69,
	0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x4d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x42,
	0x67, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x42, 0x67,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x41, 0x75,
	0x64, 0x69, 0x6f, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4e,
	0x6f, 0x69, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x53,
	0x68, 0x6f, 0x77, 0x4c, 0x69, 0x6e, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x53, 0x68, 0x6f, 0x77, 0x4c, 0x69, 0x6e, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x69, 0x67, 0x69, 0x74, 0x4d, 0x61,
	0x78, 0x53, 0x6b, 0x65, 0x77, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x44, 0x69, 0x67,
	0x69, 0x74, 0x4d, 0x61, 0x78, 0x53, 0x6b, 0x65, 0x77, 0x12, 0x24, 0x0a, 0x0d, 0x44, 0x69, 0x67,
	0x69, 0x74, 0x44, 0x6f, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x44, 0x69, 0x67, 0x69, 0x74, 0x44, 0x6f, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x44, 0x65, 0x62, 0x75, 0x67, 0x22, 0x5e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x42,
	0x61, 0x73, 0x65, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x15, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x32, 0xb9, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12,
	0x24, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x63,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f,
	0x5a, 0x0d, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x49, 0x46, 0x61, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_captcha_proto_rawDescOnce sync.Once
	file_api_captcha_proto_rawDescData = file_api_captcha_proto_rawDesc
)

func file_api_captcha_proto_rawDescGZIP() []byte {
	file_api_captcha_proto_rawDescOnce.Do(func() {
		file_api_captcha_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_captcha_proto_rawDescData)
	})
	return file_api_captcha_proto_rawDescData
}

var file_api_captcha_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_captcha_proto_goTypes = []interface{}{
	(*GetCaptchaRequest)(nil),     // 0: shared.captcha.GetCaptchaRequest
	(*GetCaptchaResponse)(nil),    // 1: shared.captcha.GetCaptchaResponse
	(*VerifyCaptchaRequest)(nil),  // 2: shared.captcha.VerifyCaptchaRequest
	(*VerifyCaptchaResponse)(nil), // 3: shared.captcha.VerifyCaptchaResponse
}
var file_api_captcha_proto_depIdxs = []int32{
	0, // 0: shared.captcha.CaptchaService.Get:input_type -> shared.captcha.GetCaptchaRequest
	2, // 1: shared.captcha.CaptchaService.Verify:input_type -> shared.captcha.VerifyCaptchaRequest
	1, // 2: shared.captcha.CaptchaService.Get:output_type -> shared.captcha.GetCaptchaResponse
	3, // 3: shared.captcha.CaptchaService.Verify:output_type -> shared.captcha.VerifyCaptchaResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_captcha_proto_init() }
func file_api_captcha_proto_init() {
	if File_api_captcha_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_captcha_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCaptchaRequest); i {
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
		file_api_captcha_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCaptchaResponse); i {
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
		file_api_captcha_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCaptchaRequest); i {
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
		file_api_captcha_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCaptchaResponse); i {
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
			RawDescriptor: file_api_captcha_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_captcha_proto_goTypes,
		DependencyIndexes: file_api_captcha_proto_depIdxs,
		MessageInfos:      file_api_captcha_proto_msgTypes,
	}.Build()
	File_api_captcha_proto = out.File
	file_api_captcha_proto_rawDesc = nil
	file_api_captcha_proto_goTypes = nil
	file_api_captcha_proto_depIdxs = nil
}
