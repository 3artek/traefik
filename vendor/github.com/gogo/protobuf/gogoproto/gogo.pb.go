// Code generated by protoc-gen-gogo.
// source: gogo.proto
// DO NOT EDIT!

/*
Package gogoproto is a generated protocol buffer package.

It is generated from these files:
	gogo.proto

It has these top-level messages:
*/
package gogoproto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

var E_GoprotoEnumPrefix = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62001,
	Name:          "gogoproto.goproto_enum_prefix",
	Tag:           "varint,62001,opt,name=goproto_enum_prefix,json=goprotoEnumPrefix",
}

var E_GoprotoEnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62021,
	Name:          "gogoproto.goproto_enum_stringer",
	Tag:           "varint,62021,opt,name=goproto_enum_stringer,json=goprotoEnumStringer",
}

var E_EnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62022,
	Name:          "gogoproto.enum_stringer",
	Tag:           "varint,62022,opt,name=enum_stringer,json=enumStringer",
}

var E_EnumCustomname = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         62023,
	Name:          "gogoproto.enum_customname",
	Tag:           "bytes,62023,opt,name=enum_customname,json=enumCustomname",
}

var E_EnumvalueCustomname = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         66001,
	Name:          "gogoproto.enumvalue_customname",
	Tag:           "bytes,66001,opt,name=enumvalue_customname,json=enumvalueCustomname",
}

var E_GoprotoGettersAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63001,
	Name:          "gogoproto.goproto_getters_all",
	Tag:           "varint,63001,opt,name=goproto_getters_all,json=goprotoGettersAll",
}

var E_GoprotoEnumPrefixAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63002,
	Name:          "gogoproto.goproto_enum_prefix_all",
	Tag:           "varint,63002,opt,name=goproto_enum_prefix_all,json=goprotoEnumPrefixAll",
}

var E_GoprotoStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63003,
	Name:          "gogoproto.goproto_stringer_all",
	Tag:           "varint,63003,opt,name=goproto_stringer_all,json=goprotoStringerAll",
}

var E_VerboseEqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63004,
	Name:          "gogoproto.verbose_equal_all",
	Tag:           "varint,63004,opt,name=verbose_equal_all,json=verboseEqualAll",
}

var E_FaceAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63005,
	Name:          "gogoproto.face_all",
	Tag:           "varint,63005,opt,name=face_all,json=faceAll",
}

var E_GostringAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63006,
	Name:          "gogoproto.gostring_all",
	Tag:           "varint,63006,opt,name=gostring_all,json=gostringAll",
}

var E_PopulateAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63007,
	Name:          "gogoproto.populate_all",
	Tag:           "varint,63007,opt,name=populate_all,json=populateAll",
}

var E_StringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63008,
	Name:          "gogoproto.stringer_all",
	Tag:           "varint,63008,opt,name=stringer_all,json=stringerAll",
}

var E_OnlyoneAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63009,
	Name:          "gogoproto.onlyone_all",
	Tag:           "varint,63009,opt,name=onlyone_all,json=onlyoneAll",
}

var E_EqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63013,
	Name:          "gogoproto.equal_all",
	Tag:           "varint,63013,opt,name=equal_all,json=equalAll",
}

var E_DescriptionAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63014,
	Name:          "gogoproto.description_all",
	Tag:           "varint,63014,opt,name=description_all,json=descriptionAll",
}

var E_TestgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63015,
	Name:          "gogoproto.testgen_all",
	Tag:           "varint,63015,opt,name=testgen_all,json=testgenAll",
}

var E_BenchgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63016,
	Name:          "gogoproto.benchgen_all",
	Tag:           "varint,63016,opt,name=benchgen_all,json=benchgenAll",
}

var E_MarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63017,
	Name:          "gogoproto.marshaler_all",
	Tag:           "varint,63017,opt,name=marshaler_all,json=marshalerAll",
}

var E_UnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63018,
	Name:          "gogoproto.unmarshaler_all",
	Tag:           "varint,63018,opt,name=unmarshaler_all,json=unmarshalerAll",
}

var E_StableMarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63019,
	Name:          "gogoproto.stable_marshaler_all",
	Tag:           "varint,63019,opt,name=stable_marshaler_all,json=stableMarshalerAll",
}

var E_SizerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63020,
	Name:          "gogoproto.sizer_all",
	Tag:           "varint,63020,opt,name=sizer_all,json=sizerAll",
}

var E_GoprotoEnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63021,
	Name:          "gogoproto.goproto_enum_stringer_all",
	Tag:           "varint,63021,opt,name=goproto_enum_stringer_all,json=goprotoEnumStringerAll",
}

var E_EnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63022,
	Name:          "gogoproto.enum_stringer_all",
	Tag:           "varint,63022,opt,name=enum_stringer_all,json=enumStringerAll",
}

var E_UnsafeMarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63023,
	Name:          "gogoproto.unsafe_marshaler_all",
	Tag:           "varint,63023,opt,name=unsafe_marshaler_all,json=unsafeMarshalerAll",
}

var E_UnsafeUnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63024,
	Name:          "gogoproto.unsafe_unmarshaler_all",
	Tag:           "varint,63024,opt,name=unsafe_unmarshaler_all,json=unsafeUnmarshalerAll",
}

var E_GoprotoExtensionsMapAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63025,
	Name:          "gogoproto.goproto_extensions_map_all",
	Tag:           "varint,63025,opt,name=goproto_extensions_map_all,json=goprotoExtensionsMapAll",
}

var E_GoprotoUnrecognizedAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63026,
	Name:          "gogoproto.goproto_unrecognized_all",
	Tag:           "varint,63026,opt,name=goproto_unrecognized_all,json=goprotoUnrecognizedAll",
}

var E_GogoprotoImport = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63027,
	Name:          "gogoproto.gogoproto_import",
	Tag:           "varint,63027,opt,name=gogoproto_import,json=gogoprotoImport",
}

var E_ProtosizerAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63028,
	Name:          "gogoproto.protosizer_all",
	Tag:           "varint,63028,opt,name=protosizer_all,json=protosizerAll",
}

var E_CompareAll = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63029,
	Name:          "gogoproto.compare_all",
	Tag:           "varint,63029,opt,name=compare_all,json=compareAll",
}

var E_GoprotoGetters = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64001,
	Name:          "gogoproto.goproto_getters",
	Tag:           "varint,64001,opt,name=goproto_getters,json=goprotoGetters",
}

var E_GoprotoStringer = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64003,
	Name:          "gogoproto.goproto_stringer",
	Tag:           "varint,64003,opt,name=goproto_stringer,json=goprotoStringer",
}

var E_VerboseEqual = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64004,
	Name:          "gogoproto.verbose_equal",
	Tag:           "varint,64004,opt,name=verbose_equal,json=verboseEqual",
}

var E_Face = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64005,
	Name:          "gogoproto.face",
	Tag:           "varint,64005,opt,name=face",
}

var E_Gostring = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64006,
	Name:          "gogoproto.gostring",
	Tag:           "varint,64006,opt,name=gostring",
}

var E_Populate = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64007,
	Name:          "gogoproto.populate",
	Tag:           "varint,64007,opt,name=populate",
}

var E_Stringer = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         67008,
	Name:          "gogoproto.stringer",
	Tag:           "varint,67008,opt,name=stringer",
}

var E_Onlyone = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64009,
	Name:          "gogoproto.onlyone",
	Tag:           "varint,64009,opt,name=onlyone",
}

var E_Equal = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64013,
	Name:          "gogoproto.equal",
	Tag:           "varint,64013,opt,name=equal",
}

var E_Description = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64014,
	Name:          "gogoproto.description",
	Tag:           "varint,64014,opt,name=description",
}

var E_Testgen = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64015,
	Name:          "gogoproto.testgen",
	Tag:           "varint,64015,opt,name=testgen",
}

var E_Benchgen = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64016,
	Name:          "gogoproto.benchgen",
	Tag:           "varint,64016,opt,name=benchgen",
}

var E_Marshaler = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64017,
	Name:          "gogoproto.marshaler",
	Tag:           "varint,64017,opt,name=marshaler",
}

var E_Unmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64018,
	Name:          "gogoproto.unmarshaler",
	Tag:           "varint,64018,opt,name=unmarshaler",
}

var E_StableMarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64019,
	Name:          "gogoproto.stable_marshaler",
	Tag:           "varint,64019,opt,name=stable_marshaler,json=stableMarshaler",
}

var E_Sizer = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64020,
	Name:          "gogoproto.sizer",
	Tag:           "varint,64020,opt,name=sizer",
}

var E_UnsafeMarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64023,
	Name:          "gogoproto.unsafe_marshaler",
	Tag:           "varint,64023,opt,name=unsafe_marshaler,json=unsafeMarshaler",
}

var E_UnsafeUnmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64024,
	Name:          "gogoproto.unsafe_unmarshaler",
	Tag:           "varint,64024,opt,name=unsafe_unmarshaler,json=unsafeUnmarshaler",
}

var E_GoprotoExtensionsMap = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64025,
	Name:          "gogoproto.goproto_extensions_map",
	Tag:           "varint,64025,opt,name=goproto_extensions_map,json=goprotoExtensionsMap",
}

var E_GoprotoUnrecognized = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64026,
	Name:          "gogoproto.goproto_unrecognized",
	Tag:           "varint,64026,opt,name=goproto_unrecognized,json=goprotoUnrecognized",
}

var E_Protosizer = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64028,
	Name:          "gogoproto.protosizer",
	Tag:           "varint,64028,opt,name=protosizer",
}

var E_Compare = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64029,
	Name:          "gogoproto.compare",
	Tag:           "varint,64029,opt,name=compare",
}

var E_Nullable = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65001,
	Name:          "gogoproto.nullable",
	Tag:           "varint,65001,opt,name=nullable",
}

var E_Embed = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65002,
	Name:          "gogoproto.embed",
	Tag:           "varint,65002,opt,name=embed",
}

var E_Customtype = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65003,
	Name:          "gogoproto.customtype",
	Tag:           "bytes,65003,opt,name=customtype",
}

var E_Customname = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65004,
	Name:          "gogoproto.customname",
	Tag:           "bytes,65004,opt,name=customname",
}

var E_Jsontag = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65005,
	Name:          "gogoproto.jsontag",
	Tag:           "bytes,65005,opt,name=jsontag",
}

var E_Moretags = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65006,
	Name:          "gogoproto.moretags",
	Tag:           "bytes,65006,opt,name=moretags",
}

var E_Casttype = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65007,
	Name:          "gogoproto.casttype",
	Tag:           "bytes,65007,opt,name=casttype",
}

var E_Castkey = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65008,
	Name:          "gogoproto.castkey",
	Tag:           "bytes,65008,opt,name=castkey",
}

var E_Castvalue = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65009,
	Name:          "gogoproto.castvalue",
	Tag:           "bytes,65009,opt,name=castvalue",
}

func init() {
	proto.RegisterExtension(E_GoprotoEnumPrefix)
	proto.RegisterExtension(E_GoprotoEnumStringer)
	proto.RegisterExtension(E_EnumStringer)
	proto.RegisterExtension(E_EnumCustomname)
	proto.RegisterExtension(E_EnumvalueCustomname)
	proto.RegisterExtension(E_GoprotoGettersAll)
	proto.RegisterExtension(E_GoprotoEnumPrefixAll)
	proto.RegisterExtension(E_GoprotoStringerAll)
	proto.RegisterExtension(E_VerboseEqualAll)
	proto.RegisterExtension(E_FaceAll)
	proto.RegisterExtension(E_GostringAll)
	proto.RegisterExtension(E_PopulateAll)
	proto.RegisterExtension(E_StringerAll)
	proto.RegisterExtension(E_OnlyoneAll)
	proto.RegisterExtension(E_EqualAll)
	proto.RegisterExtension(E_DescriptionAll)
	proto.RegisterExtension(E_TestgenAll)
	proto.RegisterExtension(E_BenchgenAll)
	proto.RegisterExtension(E_MarshalerAll)
	proto.RegisterExtension(E_UnmarshalerAll)
	proto.RegisterExtension(E_StableMarshalerAll)
	proto.RegisterExtension(E_SizerAll)
	proto.RegisterExtension(E_GoprotoEnumStringerAll)
	proto.RegisterExtension(E_EnumStringerAll)
	proto.RegisterExtension(E_UnsafeMarshalerAll)
	proto.RegisterExtension(E_UnsafeUnmarshalerAll)
	proto.RegisterExtension(E_GoprotoExtensionsMapAll)
	proto.RegisterExtension(E_GoprotoUnrecognizedAll)
	proto.RegisterExtension(E_GogoprotoImport)
	proto.RegisterExtension(E_ProtosizerAll)
	proto.RegisterExtension(E_CompareAll)
	proto.RegisterExtension(E_GoprotoGetters)
	proto.RegisterExtension(E_GoprotoStringer)
	proto.RegisterExtension(E_VerboseEqual)
	proto.RegisterExtension(E_Face)
	proto.RegisterExtension(E_Gostring)
	proto.RegisterExtension(E_Populate)
	proto.RegisterExtension(E_Stringer)
	proto.RegisterExtension(E_Onlyone)
	proto.RegisterExtension(E_Equal)
	proto.RegisterExtension(E_Description)
	proto.RegisterExtension(E_Testgen)
	proto.RegisterExtension(E_Benchgen)
	proto.RegisterExtension(E_Marshaler)
	proto.RegisterExtension(E_Unmarshaler)
	proto.RegisterExtension(E_StableMarshaler)
	proto.RegisterExtension(E_Sizer)
	proto.RegisterExtension(E_UnsafeMarshaler)
	proto.RegisterExtension(E_UnsafeUnmarshaler)
	proto.RegisterExtension(E_GoprotoExtensionsMap)
	proto.RegisterExtension(E_GoprotoUnrecognized)
	proto.RegisterExtension(E_Protosizer)
	proto.RegisterExtension(E_Compare)
	proto.RegisterExtension(E_Nullable)
	proto.RegisterExtension(E_Embed)
	proto.RegisterExtension(E_Customtype)
	proto.RegisterExtension(E_Customname)
	proto.RegisterExtension(E_Jsontag)
	proto.RegisterExtension(E_Moretags)
	proto.RegisterExtension(E_Casttype)
	proto.RegisterExtension(E_Castkey)
	proto.RegisterExtension(E_Castvalue)
}

func init() { proto.RegisterFile("gogo.proto", fileDescriptorGogo) }

var fileDescriptorGogo = []byte{
	// 1098 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x97, 0xc9, 0x6f, 0x1c, 0x45,
	0x14, 0x87, 0x85, 0x70, 0xe4, 0x99, 0xe7, 0x0d, 0x8f, 0x8d, 0x09, 0x11, 0x88, 0xe4, 0xc6, 0xc9,
	0x39, 0x45, 0x28, 0x65, 0x45, 0x96, 0x63, 0x39, 0xa3, 0x20, 0x0c, 0x23, 0x13, 0x07, 0x10, 0x87,
	0x51, 0xcf, 0xb8, 0xdc, 0x19, 0xe8, 0xee, 0x6a, 0xba, 0xba, 0xa3, 0x38, 0x37, 0x14, 0x16, 0x21,
	0xc4, 0x8e, 0x04, 0x09, 0x09, 0xcb, 0x81, 0x7d, 0x0d, 0xcb, 0x9d, 0x0b, 0x70, 0xe6, 0x7f, 0xe0,
	0x02, 0x98, 0x4d, 0xf2, 0xcd, 0x17, 0xf4, 0xba, 0xdf, 0xeb, 0xa9, 0x1e, 0x8f, 0x54, 0x35, 0xb7,
	0xf6, 0xb8, 0xbe, 0x6f, 0xaa, 0xdf, 0xeb, 0x7a, 0xbf, 0x69, 0x00, 0x5f, 0xf9, 0x6a, 0x31, 0x4e,
	0x54, 0xaa, 0x1a, 0x75, 0xbc, 0xce, 0x2f, 0x8f, 0x1c, 0xf5, 0x95, 0xf2, 0x03, 0x79, 0x3c, 0xff,
	0xab, 0x93, 0x6d, 0x1f, 0xdf, 0x92, 0xba, 0x9b, 0xf4, 0xe2, 0x54, 0x25, 0xc5, 0x62, 0xf1, 0x20,
	0xcc, 0xd1, 0xe2, 0xb6, 0x8c, 0xb2, 0xb0, 0x1d, 0x27, 0x72, 0xbb, 0x77, 0xa9, 0x71, 0xd7, 0x62,
	0x41, 0x2e, 0x32, 0xb9, 0xb8, 0x16, 0x65, 0xe1, 0x43, 0x71, 0xda, 0x53, 0x91, 0x3e, 0x7c, 0xf3,
	0xb7, 0x5b, 0x8f, 0xde, 0x72, 0x6f, 0x6d, 0x63, 0x96, 0x50, 0xfc, 0x5f, 0x2b, 0x07, 0xc5, 0x06,
	0xdc, 0x5e, 0xf1, 0xe9, 0x34, 0xe9, 0x45, 0xbe, 0x4c, 0x2c, 0xc6, 0x9f, 0xc8, 0x38, 0x67, 0x18,
	0x1f, 0x26, 0x54, 0xac, 0xc2, 0xd4, 0x28, 0xae, 0x9f, 0xc9, 0x35, 0x29, 0x4d, 0x49, 0x13, 0x66,
	0x72, 0x49, 0x37, 0xd3, 0xa9, 0x0a, 0x23, 0x2f, 0x94, 0x16, 0xcd, 0x2f, 0xb9, 0xa6, 0xbe, 0x31,
	0x8d, 0xd8, 0x6a, 0x49, 0x89, 0xf3, 0x30, 0x8f, 0x9f, 0x5c, 0xf4, 0x82, 0x4c, 0x9a, 0xb6, 0x63,
	0x43, 0x6d, 0xe7, 0x71, 0x19, 0x2b, 0x7f, 0xbd, 0x32, 0x96, 0x2b, 0xe7, 0x4a, 0x81, 0xe1, 0x35,
	0x3a, 0xe1, 0xcb, 0x34, 0x95, 0x89, 0x6e, 0x7b, 0x41, 0x30, 0x64, 0x93, 0x67, 0x7a, 0x41, 0x69,
	0xbc, 0xba, 0x5b, 0xed, 0x44, 0xb3, 0x20, 0x57, 0x82, 0x40, 0x6c, 0xc2, 0x1d, 0x43, 0x3a, 0xeb,
	0xe0, 0xbc, 0x46, 0xce, 0xf9, 0x03, 0xdd, 0x45, 0x6d, 0x0b, 0xf8, 0xf3, 0xb2, 0x1f, 0x0e, 0xce,
	0x77, 0xc8, 0xd9, 0x20, 0x96, 0xdb, 0x82, 0xc6, 0xfb, 0x61, 0xf6, 0xa2, 0x4c, 0x3a, 0x4a, 0xcb,
	0xb6, 0x7c, 0x2a, 0xf3, 0x02, 0x07, 0xdd, 0x75, 0xd2, 0xcd, 0x10, 0xb8, 0x86, 0x1c, 0xba, 0x4e,
	0x42, 0x6d, 0xdb, 0xeb, 0x4a, 0x07, 0xc5, 0x0d, 0x52, 0x8c, 0xe3, 0x7a, 0x44, 0x57, 0x60, 0xd2,
	0x57, 0xc5, 0x2d, 0x39, 0xe0, 0xef, 0x12, 0x3e, 0xc1, 0x0c, 0x29, 0x62, 0x15, 0x67, 0x81, 0x97,
	0xba, 0xec, 0xe0, 0x3d, 0x56, 0x30, 0x43, 0x8a, 0x11, 0xca, 0xfa, 0x3e, 0x2b, 0xb4, 0x51, 0xcf,
	0x65, 0x98, 0x50, 0x51, 0xb0, 0xa3, 0x22, 0x97, 0x4d, 0x7c, 0x40, 0x06, 0x20, 0x04, 0x05, 0x4b,
	0x50, 0x77, 0x6d, 0xc4, 0x87, 0x84, 0xd7, 0x24, 0x77, 0xa0, 0x09, 0x33, 0x3c, 0x64, 0x7a, 0x2a,
	0x72, 0x50, 0x7c, 0x44, 0x8a, 0x69, 0x03, 0xa3, 0xdb, 0x48, 0xa5, 0x4e, 0x7d, 0xe9, 0x22, 0xf9,
	0x98, 0x6f, 0x83, 0x10, 0x2a, 0x65, 0x47, 0x46, 0xdd, 0x0b, 0x6e, 0x86, 0x4f, 0xb8, 0x94, 0xcc,
	0xa0, 0x62, 0x15, 0xa6, 0x42, 0x2f, 0xd1, 0x17, 0xbc, 0xc0, 0xa9, 0x1d, 0x9f, 0x92, 0x63, 0xb2,
	0x84, 0xa8, 0x22, 0x59, 0x34, 0x8a, 0xe6, 0x33, 0xae, 0x88, 0x81, 0xd1, 0xd1, 0xd3, 0xa9, 0xd7,
	0x09, 0x64, 0x7b, 0x14, 0xdb, 0xe7, 0x7c, 0xf4, 0x0a, 0x76, 0xdd, 0x34, 0x2e, 0x41, 0x5d, 0xf7,
	0x2e, 0x3b, 0x69, 0xbe, 0xe0, 0x4e, 0xe7, 0x00, 0xc2, 0x8f, 0xc1, 0x9d, 0x43, 0x47, 0xbd, 0x83,
	0xec, 0x4b, 0x92, 0x2d, 0x0c, 0x19, 0xf7, 0x34, 0x12, 0x46, 0x55, 0x7e, 0xc5, 0x23, 0x41, 0x0e,
	0xb8, 0x5a, 0x30, 0x9f, 0x45, 0xda, 0xdb, 0x1e, 0xad, 0x6a, 0x5f, 0x73, 0xd5, 0x0a, 0xb6, 0x52,
	0xb5, 0x73, 0xb0, 0x40, 0xc6, 0xd1, 0xfa, 0xfa, 0x0d, 0x0f, 0xd6, 0x82, 0xde, 0xac, 0x76, 0xf7,
	0x71, 0x38, 0x52, 0x96, 0xf3, 0x52, 0x2a, 0x23, 0x8d, 0x4c, 0x3b, 0xf4, 0x62, 0x07, 0xf3, 0x4d,
	0x32, 0xf3, 0xc4, 0x5f, 0x2b, 0x05, 0xeb, 0x5e, 0x8c, 0xf2, 0x47, 0xe1, 0x30, 0xcb, 0xb3, 0x28,
	0x91, 0x5d, 0xe5, 0x47, 0xbd, 0xcb, 0x72, 0xcb, 0x41, 0xfd, 0xed, 0x40, 0xab, 0x36, 0x0d, 0x1c,
	0xcd, 0x67, 0xe1, 0xb6, 0xf2, 0xf7, 0x46, 0xbb, 0x17, 0xc6, 0x2a, 0x49, 0x2d, 0xc6, 0xef, 0xb8,
	0x53, 0x25, 0x77, 0x36, 0xc7, 0xc4, 0x1a, 0x4c, 0xe7, 0x7f, 0xba, 0x3e, 0x92, 0xdf, 0x93, 0x68,
	0xaa, 0x4f, 0xd1, 0xe0, 0xe8, 0xaa, 0x30, 0xf6, 0x12, 0x97, 0xf9, 0xf7, 0x03, 0x0f, 0x0e, 0x42,
	0x8a, 0xa7, 0x6f, 0x66, 0x20, 0x89, 0x1b, 0xf7, 0x1c, 0x90, 0xac, 0x4b, 0xad, 0x3d, 0xbf, 0xf4,
	0x3c, 0xbd, 0x47, 0x67, 0xb6, 0x1a, 0xc4, 0xe2, 0x01, 0x2c, 0x4f, 0x35, 0x2e, 0xed, 0xb2, 0x2b,
	0x7b, 0x65, 0x85, 0x2a, 0x69, 0x29, 0xce, 0xc0, 0x54, 0x25, 0x2a, 0xed, 0xaa, 0x67, 0x48, 0x35,
	0x69, 0x26, 0xa5, 0x38, 0x01, 0x63, 0x18, 0x7b, 0x76, 0xfc, 0x59, 0xc2, 0xf3, 0xe5, 0xe2, 0x14,
	0xd4, 0x38, 0xee, 0xec, 0xe8, 0x73, 0x84, 0x96, 0x08, 0xe2, 0x1c, 0x75, 0x76, 0xfc, 0x79, 0xc6,
	0x19, 0x41, 0xdc, 0xbd, 0x84, 0x3f, 0xbe, 0x38, 0x46, 0xe3, 0x8a, 0x6b, 0xb7, 0x04, 0xe3, 0x94,
	0x71, 0x76, 0xfa, 0x05, 0xfa, 0x72, 0x26, 0xc4, 0x7d, 0x70, 0xc8, 0xb1, 0xe0, 0x2f, 0x11, 0x5a,
	0xac, 0x17, 0xab, 0x30, 0x61, 0xe4, 0x9a, 0x1d, 0x7f, 0x99, 0x70, 0x93, 0xc2, 0xad, 0x53, 0xae,
	0xd9, 0x05, 0xaf, 0xf0, 0xd6, 0x89, 0xc0, 0xb2, 0x71, 0xa4, 0xd9, 0xe9, 0x57, 0xb9, 0xea, 0x8c,
	0x88, 0x65, 0xa8, 0x97, 0x63, 0xca, 0xce, 0xbf, 0x46, 0x7c, 0x9f, 0xc1, 0x0a, 0x18, 0x63, 0xd2,
	0xae, 0x78, 0x9d, 0x2b, 0x60, 0x50, 0x78, 0x8c, 0x06, 0xa3, 0xcf, 0x6e, 0x7a, 0x83, 0x8f, 0xd1,
	0x40, 0xf2, 0x61, 0x37, 0xf3, 0x69, 0x61, 0x57, 0xbc, 0xc9, 0xdd, 0xcc, 0xd7, 0xe3, 0x36, 0x06,
	0xb3, 0xc4, 0xee, 0x78, 0x8b, 0xb7, 0x31, 0x10, 0x25, 0xa2, 0x05, 0x8d, 0x83, 0x39, 0x62, 0xf7,
	0xbd, 0x4d, 0xbe, 0xd9, 0x03, 0x31, 0x22, 0x1e, 0x81, 0x85, 0xe1, 0x19, 0x62, 0xb7, 0x5e, 0xdd,
	0x1b, 0xf8, 0xd5, 0x6f, 0x46, 0x88, 0x38, 0xd7, 0xff, 0xd5, 0x6f, 0xe6, 0x87, 0x5d, 0x7b, 0x6d,
	0xaf, 0xfa, 0x62, 0x67, 0xc6, 0x87, 0x58, 0x01, 0xe8, 0x8f, 0x6e, 0xbb, 0xeb, 0x3a, 0xb9, 0x0c,
	0x08, 0x8f, 0x06, 0x4d, 0x6e, 0x3b, 0x7f, 0x83, 0x8f, 0x06, 0x11, 0x62, 0x09, 0x6a, 0x51, 0x16,
	0x04, 0xf8, 0x70, 0x34, 0xee, 0x1e, 0x12, 0x13, 0x32, 0xd8, 0x62, 0xf6, 0xf7, 0x7d, 0x3a, 0x18,
	0x0c, 0x88, 0x13, 0x70, 0x48, 0x86, 0x1d, 0xb9, 0x65, 0x23, 0xff, 0xd8, 0xe7, 0x81, 0x80, 0xab,
	0xc5, 0x32, 0x40, 0xf1, 0xd2, 0x98, 0xee, 0xc4, 0xd6, 0x6f, 0xfd, 0x73, 0xbf, 0x78, 0x07, 0x35,
	0x90, 0xbe, 0x20, 0x7f, 0xeb, 0xb4, 0x08, 0x76, 0xab, 0x82, 0xfc, 0x45, 0xf3, 0x24, 0x8c, 0x3f,
	0xa1, 0x55, 0x94, 0x7a, 0xbe, 0x8d, 0xfe, 0x8b, 0x68, 0x5e, 0x8f, 0x05, 0x0b, 0x55, 0x22, 0x53,
	0xcf, 0xd7, 0x36, 0xf6, 0x6f, 0x62, 0x4b, 0x00, 0xe1, 0xae, 0xa7, 0x53, 0x97, 0xfb, 0xfe, 0x87,
	0x61, 0x06, 0x70, 0xd3, 0x78, 0xfd, 0xa4, 0xdc, 0xb1, 0xb1, 0xff, 0xf2, 0xa6, 0x69, 0xbd, 0x38,
	0x05, 0x75, 0xbc, 0xcc, 0xdf, 0xb7, 0x6d, 0xf0, 0x7f, 0x04, 0xf7, 0x89, 0xd3, 0xc7, 0x60, 0xae,
	0xab, 0xc2, 0x41, 0xec, 0x34, 0x34, 0x55, 0x53, 0xb5, 0xf2, 0x07, 0xf1, 0xff, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x87, 0x5c, 0xee, 0x2b, 0x7e, 0x11, 0x00, 0x00,
}
