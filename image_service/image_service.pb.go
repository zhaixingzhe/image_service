// Code generated by protoc-gen-go.
// source: image_service.proto
// DO NOT EDIT!

/*
Package image_service is a generated protocol buffer package.

It is generated from these files:
	image_service.proto

It has these top-level messages:
	ImageStoreRequest
	Image
	ImageOperation
	DeleteRequest
	DeleteResponse
*/
package image_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Format int32

const (
	Format_JPEG Format = 0
	Format_WEBP Format = 1
	Format_PNG  Format = 2
)

var Format_name = map[int32]string{
	0: "JPEG",
	1: "WEBP",
	2: "PNG",
}
var Format_value = map[string]int32{
	"JPEG": 0,
	"WEBP": 1,
	"PNG":  2,
}

func (x Format) String() string {
	return proto.EnumName(Format_name, int32(x))
}
func (Format) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ImageStoreRequest struct {
	Filename string            `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Data     []byte            `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Ops      []*ImageOperation `protobuf:"bytes,3,rep,name=ops" json:"ops,omitempty"`
}

func (m *ImageStoreRequest) Reset()                    { *m = ImageStoreRequest{} }
func (m *ImageStoreRequest) String() string            { return proto.CompactTextString(m) }
func (*ImageStoreRequest) ProtoMessage()               {}
func (*ImageStoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ImageStoreRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *ImageStoreRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ImageStoreRequest) GetOps() []*ImageOperation {
	if m != nil {
		return m.Ops
	}
	return nil
}

type Image struct {
	VersionName string `protobuf:"bytes,1,opt,name=version_name,json=versionName" json:"version_name,omitempty"`
	Filename    string `protobuf:"bytes,2,opt,name=filename" json:"filename,omitempty"`
	Url         string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Image) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func (m *Image) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Image) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ImageOperation struct {
	VersionName string `protobuf:"bytes,1,opt,name=version_name,json=versionName" json:"version_name,omitempty"`
	Height      int32  `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Width       int32  `protobuf:"varint,3,opt,name=width" json:"width,omitempty"`
	Quality     int32  `protobuf:"varint,4,opt,name=quality" json:"quality,omitempty"`
	Compression int32  `protobuf:"varint,5,opt,name=compression" json:"compression,omitempty"`
	Crop        bool   `protobuf:"varint,6,opt,name=crop" json:"crop,omitempty"`
	Enlarge     bool   `protobuf:"varint,7,opt,name=enlarge" json:"enlarge,omitempty"`
	Flip        bool   `protobuf:"varint,8,opt,name=flip" json:"flip,omitempty"`
	Interlace   bool   `protobuf:"varint,9,opt,name=interlace" json:"interlace,omitempty"`
	Format      Format `protobuf:"varint,10,opt,name=format,enum=image_service.Format" json:"format,omitempty"`
}

func (m *ImageOperation) Reset()                    { *m = ImageOperation{} }
func (m *ImageOperation) String() string            { return proto.CompactTextString(m) }
func (*ImageOperation) ProtoMessage()               {}
func (*ImageOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ImageOperation) GetVersionName() string {
	if m != nil {
		return m.VersionName
	}
	return ""
}

func (m *ImageOperation) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ImageOperation) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ImageOperation) GetQuality() int32 {
	if m != nil {
		return m.Quality
	}
	return 0
}

func (m *ImageOperation) GetCompression() int32 {
	if m != nil {
		return m.Compression
	}
	return 0
}

func (m *ImageOperation) GetCrop() bool {
	if m != nil {
		return m.Crop
	}
	return false
}

func (m *ImageOperation) GetEnlarge() bool {
	if m != nil {
		return m.Enlarge
	}
	return false
}

func (m *ImageOperation) GetFlip() bool {
	if m != nil {
		return m.Flip
	}
	return false
}

func (m *ImageOperation) GetInterlace() bool {
	if m != nil {
		return m.Interlace
	}
	return false
}

func (m *ImageOperation) GetFormat() Format {
	if m != nil {
		return m.Format
	}
	return Format_JPEG
}

type DeleteRequest struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type DeleteResponse struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteResponse) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageStoreRequest)(nil), "image_service.ImageStoreRequest")
	proto.RegisterType((*Image)(nil), "image_service.Image")
	proto.RegisterType((*ImageOperation)(nil), "image_service.ImageOperation")
	proto.RegisterType((*DeleteRequest)(nil), "image_service.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "image_service.DeleteResponse")
	proto.RegisterEnum("image_service.Format", Format_name, Format_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ImageService service

type ImageServiceClient interface {
	Store(ctx context.Context, in *ImageStoreRequest, opts ...grpc.CallOption) (ImageService_StoreClient, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type imageServiceClient struct {
	cc *grpc.ClientConn
}

func NewImageServiceClient(cc *grpc.ClientConn) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) Store(ctx context.Context, in *ImageStoreRequest, opts ...grpc.CallOption) (ImageService_StoreClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ImageService_serviceDesc.Streams[0], c.cc, "/image_service.ImageService/Store", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageServiceStoreClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ImageService_StoreClient interface {
	Recv() (*Image, error)
	grpc.ClientStream
}

type imageServiceStoreClient struct {
	grpc.ClientStream
}

func (x *imageServiceStoreClient) Recv() (*Image, error) {
	m := new(Image)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imageServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/image_service.ImageService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ImageService service

type ImageServiceServer interface {
	Store(*ImageStoreRequest, ImageService_StoreServer) error
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterImageServiceServer(s *grpc.Server, srv ImageServiceServer) {
	s.RegisterService(&_ImageService_serviceDesc, srv)
}

func _ImageService_Store_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ImageStoreRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImageServiceServer).Store(m, &imageServiceStoreServer{stream})
}

type ImageService_StoreServer interface {
	Send(*Image) error
	grpc.ServerStream
}

type imageServiceStoreServer struct {
	grpc.ServerStream
}

func (x *imageServiceStoreServer) Send(m *Image) error {
	return x.ServerStream.SendMsg(m)
}

func _ImageService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image_service.ImageService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "image_service.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _ImageService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Store",
			Handler:       _ImageService_Store_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "image_service.proto",
}

func init() { proto.RegisterFile("image_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xed, 0x6e, 0xd3, 0x40,
	0x10, 0xcc, 0xc5, 0xb1, 0x93, 0x6c, 0xd2, 0x28, 0x2c, 0x05, 0x9d, 0xa2, 0x56, 0x32, 0x96, 0x90,
	0x2c, 0x3e, 0x0a, 0x0a, 0x6f, 0x00, 0x94, 0x08, 0x7e, 0x94, 0xe8, 0xf8, 0x01, 0xff, 0xaa, 0x23,
	0xdd, 0x24, 0x27, 0x39, 0x3e, 0xf7, 0x7c, 0x29, 0xe2, 0x5d, 0x78, 0x04, 0x1e, 0x12, 0xdd, 0x39,
	0x01, 0x3b, 0x2a, 0x55, 0xff, 0xcd, 0xcc, 0x8e, 0x76, 0xf6, 0xbc, 0x6b, 0x78, 0xa8, 0x36, 0x72,
	0x45, 0x97, 0x25, 0x99, 0x1b, 0xb5, 0xa0, 0xb3, 0xc2, 0x68, 0xab, 0xf1, 0xa8, 0x21, 0x26, 0x16,
	0x1e, 0x7c, 0x74, 0xc2, 0x17, 0xab, 0x0d, 0x09, 0xba, 0xde, 0x52, 0x69, 0x71, 0x02, 0xbd, 0xa5,
	0xca, 0x28, 0x97, 0x1b, 0xe2, 0x2c, 0x66, 0x69, 0x5f, 0xfc, 0xe5, 0x88, 0xd0, 0xb9, 0x92, 0x56,
	0xf2, 0x76, 0xcc, 0xd2, 0xa1, 0xf0, 0x18, 0x5f, 0x41, 0xa0, 0x8b, 0x92, 0x07, 0x71, 0x90, 0x0e,
	0xa6, 0xa7, 0x67, 0xcd, 0x58, 0xdf, 0xfe, 0x73, 0x41, 0x46, 0x5a, 0xa5, 0x73, 0xe1, 0x9c, 0xc9,
	0x37, 0x08, 0xbd, 0x8c, 0x4f, 0x60, 0x78, 0x43, 0xa6, 0x54, 0x3a, 0xbf, 0xac, 0xa5, 0x0d, 0x76,
	0xda, 0x85, 0x0b, 0xac, 0x0f, 0xd3, 0x3e, 0x18, 0x66, 0x0c, 0xc1, 0xd6, 0x64, 0x3c, 0xf0, 0xb2,
	0x83, 0xc9, 0xef, 0x36, 0x8c, 0x9a, 0x89, 0xf7, 0xc9, 0x78, 0x0c, 0xd1, 0x9a, 0xd4, 0x6a, 0x6d,
	0x7d, 0x42, 0x28, 0x76, 0x0c, 0x8f, 0x21, 0xfc, 0xa1, 0xae, 0xec, 0xda, 0x27, 0x84, 0xa2, 0x22,
	0xc8, 0xa1, 0x7b, 0xbd, 0x95, 0x99, 0xb2, 0x3f, 0x79, 0xc7, 0xeb, 0x7b, 0x8a, 0x31, 0x0c, 0x16,
	0x7a, 0x53, 0x18, 0x2a, 0x5d, 0x6b, 0x1e, 0xfa, 0x6a, 0x5d, 0x72, 0x9f, 0x6f, 0x61, 0x74, 0xc1,
	0xa3, 0x98, 0xa5, 0x3d, 0xe1, 0xb1, 0xeb, 0x47, 0x79, 0x26, 0xcd, 0x8a, 0x78, 0xd7, 0xcb, 0x7b,
	0xea, 0xdc, 0xcb, 0x4c, 0x15, 0xbc, 0x57, 0xb9, 0x1d, 0xc6, 0x13, 0xe8, 0xab, 0xdc, 0x92, 0xc9,
	0xe4, 0x82, 0x78, 0xdf, 0x17, 0xfe, 0x09, 0xf8, 0x12, 0xa2, 0xa5, 0x36, 0x1b, 0x69, 0x39, 0xc4,
	0x2c, 0x1d, 0x4d, 0x1f, 0x1d, 0x6c, 0xe3, 0x83, 0x2f, 0x8a, 0x9d, 0x29, 0x79, 0x0e, 0x47, 0xef,
	0x29, 0x23, 0x7b, 0x9f, 0xd5, 0x27, 0x2f, 0x60, 0xb4, 0x37, 0x97, 0x85, 0xce, 0x4b, 0xba, 0xcb,
	0xfd, 0xec, 0x29, 0x44, 0x55, 0x18, 0xf6, 0xa0, 0xf3, 0x69, 0x7e, 0x3e, 0x1b, 0xb7, 0x1c, 0xfa,
	0x7a, 0xfe, 0x76, 0x3e, 0x66, 0xd8, 0x85, 0x60, 0x7e, 0x31, 0x1b, 0xb7, 0xa7, 0xbf, 0x18, 0x0c,
	0xab, 0x0b, 0xac, 0x26, 0xc4, 0x77, 0x10, 0xfa, 0x63, 0xc4, 0xf8, 0xb6, 0x43, 0xaa, 0xdf, 0xe9,
	0xe4, 0xf8, 0x36, 0x47, 0xd2, 0x7a, 0xcd, 0x70, 0x06, 0x51, 0x35, 0x2a, 0x9e, 0x1c, 0x78, 0x1a,
	0xcf, 0x9d, 0x9c, 0xfe, 0xa7, 0x5a, 0xbd, 0x2f, 0x69, 0x7d, 0x8f, 0xfc, 0x5f, 0xf3, 0xe6, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xc9, 0xfb, 0xad, 0x4c, 0x03, 0x00, 0x00,
}
