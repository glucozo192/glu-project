// Code generated by protoc-gen-struct-transformer, version: <dev>. DO NOT EDIT.
// source file: common.proto
// source package: pb

package transform

import (
	"github.com/glu/shopvui/idl/pb"
	"github.com/glu/shopvui/internal/userm/models"
	"github.com/glu/shopvui/transformhelpers"
)

// "pb.Product": target: "", Omitted: true, OneofDecl: ""
// "pb.ListAllProductsResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FileOptions": target: "", Omitted: true, OneofDecl: ""
// "pb.User": target: "User", Omitted: false, OneofDecl: ""
// "pb.GetMeResponse": target: "", Omitted: true, OneofDecl: ""
// "google.api.CustomHttpPattern": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.Timestamp": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FileDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumValueDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.MethodDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.MethodOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.UninterpretedOption": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FieldDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.OneofDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.GeneratedCodeInfo": target: "", Omitted: true, OneofDecl: ""
// "pb.InsertProductResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FileDescriptorSet": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.DescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "pb.LoginRequest": target: "", Omitted: true, OneofDecl: ""
// "pb.ListProductsResponse": target: "", Omitted: true, OneofDecl: ""
// "pb.RegisterRequest": target: "", Omitted: true, OneofDecl: ""
// "pb.GetMeRequest": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.ServiceDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "pb.ListProductsRequest": target: "", Omitted: true, OneofDecl: ""
// "pb.DeleteProductRequest": target: "", Omitted: true, OneofDecl: ""
// "pb.DeleteProductResponse": target: "", Omitted: true, OneofDecl: ""
// "pb.LoginResponse": target: "", Omitted: true, OneofDecl: ""
// "google.api.Http": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.OneofOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumValueOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FieldOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.ServiceOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.SourceCodeInfo": target: "", Omitted: true, OneofDecl: ""
// "pb.InsertProductRequest": target: "", Omitted: true, OneofDecl: ""
// "pb.ListAllProductsRequest": target: "", Omitted: true, OneofDecl: ""
// "pb.RegisterResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.ExtensionRangeOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.MessageOptions": target: "", Omitted: true, OneofDecl: ""
// "google.api.HttpRule": target: "", Omitted: true, OneofDecl: ""

// Target struct fields:
// Field: "FirstName", Type: "string", isPointer: false
// Field: "LastName", Type: "string", isPointer: false
// Field: "Active", Type: "pgtype.Bool", isPointer: false
// Field: "CreatedAt", Type: "pgtype.Timestamptz", isPointer: false
// Field: "UpdatedAt", Type: "pgtype.Timestamptz", isPointer: false
// Field: "DeletedAt", Type: "pgtype.Timestamptz", isPointer: false
// Field: "UserID", Type: "string", isPointer: false
// Field: "Email", Type: "string", isPointer: false
// Field: "Password", Type: "string", isPointer: false

// ===============================
// fdp.TypeName: ".google.protobuf.Timestamp"

// fdp.Name: "created_at", mapAs: "", mapTo: ""

// ===============================
// fdp.TypeName: ".google.protobuf.Timestamp"

// fdp.Name: "updated_at", mapAs: "", mapTo: ""

// ===============================
// fdp.Name: "email", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"String", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "Email", gname: "Email"

// ===============================
// fdp.Name: "user_id", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"String", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "UserId", gname: "UserID"

// ===============================
// fdp.Name: "first_name", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"String", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "FirstName", gname: "FirstName"

// ===============================
// fdp.Name: "last_name", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"String", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "LastName", gname: "LastName"

// ===============================
// fdp.Name: "password", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"String", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "Password", gname: "Password"
func PbToUserPtr(src *pb.User, opts ...TransformParam) *models.User {
	if src == nil {
		return nil
	}

	d := PbToUser(*src, opts...)
	return &d
}

func PbToUserPtrList(src []*pb.User, opts ...TransformParam) []*models.User {
	resp := make([]*models.User, len(src))

	for i, s := range src {
		resp[i] = PbToUserPtr(s, opts...)
	}

	return resp
}

func PbToUserPtrVal(src *pb.User, opts ...TransformParam) models.User {
	if src == nil {
		return models.User{}
	}

	return PbToUser(*src, opts...)
}

func PbToUserPtrValList(src []*pb.User, opts ...TransformParam) []models.User {
	resp := make([]models.User, len(src))

	for i, s := range src {
		resp[i] = PbToUser(*s)
	}

	return resp
}

// PbToUserList is DEPRECATED. Use PbToUserPtrValList instead.
func PbToUserList(src []*pb.User, opts ...TransformParam) []models.User {
	return PbToUserPtrValList(src)
}

func PbToUser(src pb.User, opts ...TransformParam) models.User {
	s := models.User{
		CreatedAt: transformhelpers.TimePtrToPgtypeTimestamptz(src.CreatedAt),
		UpdatedAt: transformhelpers.TimePtrToPgtypeTimestamptz(src.UpdatedAt),
		Email:     src.Email,
		UserID:    src.UserId,
		FirstName: src.FirstName,
		LastName:  src.LastName,
		Password:  src.Password,
	}

	applyOptions(opts...)

	return s
}

func PbToUserValPtr(src pb.User, opts ...TransformParam) *models.User {
	d := PbToUser(src, opts...)
	return &d
}

func PbToUserValList(src []pb.User, opts ...TransformParam) []models.User {
	resp := make([]models.User, len(src))

	for i, s := range src {
		resp[i] = PbToUser(s, opts...)
	}

	return resp
}

func UserToPbPtr(src *models.User, opts ...TransformParam) *pb.User {
	if src == nil {
		return nil
	}

	d := UserToPb(*src, opts...)
	return &d
}

func UserToPbPtrList(src []*models.User, opts ...TransformParam) []*pb.User {
	resp := make([]*pb.User, len(src))

	for i, s := range src {
		resp[i] = UserToPbPtr(s, opts...)
	}

	return resp
}

func UserToPbPtrVal(src *models.User, opts ...TransformParam) pb.User {
	if src == nil {
		return pb.User{}
	}

	return UserToPb(*src, opts...)
}

func UserToPbValPtrList(src []models.User, opts ...TransformParam) []*pb.User {
	resp := make([]*pb.User, len(src))

	for i, s := range src {
		g := UserToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// UserToPbList is DEPRECATED. Use UserToPbValPtrList instead.
func UserToPbList(src []models.User, opts ...TransformParam) []*pb.User {
	return UserToPbValPtrList(src)
}

func UserToPb(src models.User, opts ...TransformParam) pb.User {
	s := pb.User{
		CreatedAt: transformhelpers.PgtypeTimestamptzToTimePtr(src.CreatedAt),
		UpdatedAt: transformhelpers.PgtypeTimestamptzToTimePtr(src.UpdatedAt),
		Email:     src.Email,
		UserId:    src.UserID,
		FirstName: src.FirstName,
		LastName:  src.LastName,
		Password:  src.Password,
	}

	applyOptions(opts...)

	return s
}

func UserToPbValPtr(src models.User, opts ...TransformParam) *pb.User {
	d := UserToPb(src, opts...)
	return &d
}

func UserToPbValList(src []models.User, opts ...TransformParam) []pb.User {
	resp := make([]pb.User, len(src))

	for i, s := range src {
		resp[i] = UserToPb(s, opts...)
	}

	return resp
}
