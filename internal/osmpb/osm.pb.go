// Code generated by protoc-gen-go.
// source: osm.proto
// DO NOT EDIT!

/*
Package osmpb is a generated protocol buffer package.

It is generated from these files:
	osm.proto

It has these top-level messages:
	Changeset
	Bounds
	Change
	OSM
	Node
	Info
	DenseNodes
	DenseInfo
	Way
	Relation
*/
package osmpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Relation_MemberType int32

const (
	Relation_NODE     Relation_MemberType = 0
	Relation_WAY      Relation_MemberType = 1
	Relation_RELATION Relation_MemberType = 2
)

var Relation_MemberType_name = map[int32]string{
	0: "NODE",
	1: "WAY",
	2: "RELATION",
}
var Relation_MemberType_value = map[string]int32{
	"NODE":     0,
	"WAY":      1,
	"RELATION": 2,
}

func (x Relation_MemberType) Enum() *Relation_MemberType {
	p := new(Relation_MemberType)
	*p = x
	return p
}
func (x Relation_MemberType) String() string {
	return proto.EnumName(Relation_MemberType_name, int32(x))
}
func (x *Relation_MemberType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Relation_MemberType_value, data, "Relation_MemberType")
	if err != nil {
		return err
	}
	*x = Relation_MemberType(value)
	return nil
}
func (Relation_MemberType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{9, 0} }

type Changeset struct {
	Id *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys      []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals      []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	UserId    *int32   `protobuf:"varint,5,opt,name=user_id" json:"user_id,omitempty"`
	UserSid   *uint32  `protobuf:"varint,6,opt,name=user_sid" json:"user_sid,omitempty"`
	CreatedAt *int64   `protobuf:"varint,7,opt,name=created_at" json:"created_at,omitempty"`
	ClosedAt  *int64   `protobuf:"varint,8,opt,name=closed_at" json:"closed_at,omitempty"`
	Open      *bool    `protobuf:"varint,9,opt,name=open" json:"open,omitempty"`
	Bounds    *Bounds  `protobuf:"bytes,10,opt,name=bounds" json:"bounds,omitempty"`
	Change    *Change  `protobuf:"bytes,11,opt,name=change" json:"change,omitempty"`
	// contains the tag strings for everything
	// in this entire changeset.
	Strings          []string `protobuf:"bytes,20,rep,name=strings" json:"strings,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Changeset) Reset()                    { *m = Changeset{} }
func (m *Changeset) String() string            { return proto.CompactTextString(m) }
func (*Changeset) ProtoMessage()               {}
func (*Changeset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Changeset) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Changeset) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Changeset) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Changeset) GetUserId() int32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Changeset) GetUserSid() uint32 {
	if m != nil && m.UserSid != nil {
		return *m.UserSid
	}
	return 0
}

func (m *Changeset) GetCreatedAt() int64 {
	if m != nil && m.CreatedAt != nil {
		return *m.CreatedAt
	}
	return 0
}

func (m *Changeset) GetClosedAt() int64 {
	if m != nil && m.ClosedAt != nil {
		return *m.ClosedAt
	}
	return 0
}

func (m *Changeset) GetOpen() bool {
	if m != nil && m.Open != nil {
		return *m.Open
	}
	return false
}

func (m *Changeset) GetBounds() *Bounds {
	if m != nil {
		return m.Bounds
	}
	return nil
}

func (m *Changeset) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

func (m *Changeset) GetStrings() []string {
	if m != nil {
		return m.Strings
	}
	return nil
}

type Bounds struct {
	MinLng           *int64 `protobuf:"zigzag64,1,req,name=min_lng" json:"min_lng,omitempty"`
	MaxLng           *int64 `protobuf:"zigzag64,2,req,name=max_lng" json:"max_lng,omitempty"`
	MinLat           *int64 `protobuf:"zigzag64,3,req,name=min_lat" json:"min_lat,omitempty"`
	MaxLat           *int64 `protobuf:"zigzag64,4,req,name=max_lat" json:"max_lat,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Bounds) Reset()                    { *m = Bounds{} }
func (m *Bounds) String() string            { return proto.CompactTextString(m) }
func (*Bounds) ProtoMessage()               {}
func (*Bounds) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Bounds) GetMinLng() int64 {
	if m != nil && m.MinLng != nil {
		return *m.MinLng
	}
	return 0
}

func (m *Bounds) GetMaxLng() int64 {
	if m != nil && m.MaxLng != nil {
		return *m.MaxLng
	}
	return 0
}

func (m *Bounds) GetMinLat() int64 {
	if m != nil && m.MinLat != nil {
		return *m.MinLat
	}
	return 0
}

func (m *Bounds) GetMaxLat() int64 {
	if m != nil && m.MaxLat != nil {
		return *m.MaxLat
	}
	return 0
}

type Change struct {
	Create *OSM `protobuf:"bytes,1,opt,name=create" json:"create,omitempty"`
	Modify *OSM `protobuf:"bytes,2,opt,name=modify" json:"modify,omitempty"`
	Delete *OSM `protobuf:"bytes,3,opt,name=delete" json:"delete,omitempty"`
	// elements that give the change extra context like
	// nodes of the ways, and previous versions.
	Context *OSM `protobuf:"bytes,4,opt,name=context" json:"context,omitempty"`
	// contains the tag strings if this is the root of the data.
	Strings          []string `protobuf:"bytes,20,rep,name=strings" json:"strings,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Change) Reset()                    { *m = Change{} }
func (m *Change) String() string            { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()               {}
func (*Change) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Change) GetCreate() *OSM {
	if m != nil {
		return m.Create
	}
	return nil
}

func (m *Change) GetModify() *OSM {
	if m != nil {
		return m.Modify
	}
	return nil
}

func (m *Change) GetDelete() *OSM {
	if m != nil {
		return m.Delete
	}
	return nil
}

func (m *Change) GetContext() *OSM {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Change) GetStrings() []string {
	if m != nil {
		return m.Strings
	}
	return nil
}

type OSM struct {
	Bounds *Bounds `protobuf:"bytes,1,opt,name=bounds" json:"bounds,omitempty"`
	// an encoded should have either nodes or a dense_nodes, but not both.
	Nodes      []*Node     `protobuf:"bytes,2,rep,name=nodes" json:"nodes,omitempty"`
	DenseNodes *DenseNodes `protobuf:"bytes,3,opt,name=dense_nodes" json:"dense_nodes,omitempty"`
	Ways       []*Way      `protobuf:"bytes,4,rep,name=ways" json:"ways,omitempty"`
	Relations  []*Relation `protobuf:"bytes,5,rep,name=relations" json:"relations,omitempty"`
	// contains the tag strings if this is the root of the data.
	Strings          []string `protobuf:"bytes,20,rep,name=strings" json:"strings,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *OSM) Reset()                    { *m = OSM{} }
func (m *OSM) String() string            { return proto.CompactTextString(m) }
func (*OSM) ProtoMessage()               {}
func (*OSM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OSM) GetBounds() *Bounds {
	if m != nil {
		return m.Bounds
	}
	return nil
}

func (m *OSM) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *OSM) GetDenseNodes() *DenseNodes {
	if m != nil {
		return m.DenseNodes
	}
	return nil
}

func (m *OSM) GetWays() []*Way {
	if m != nil {
		return m.Ways
	}
	return nil
}

func (m *OSM) GetRelations() []*Relation {
	if m != nil {
		return m.Relations
	}
	return nil
}

func (m *OSM) GetStrings() []string {
	if m != nil {
		return m.Strings
	}
	return nil
}

type Node struct {
	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys             []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals             []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info             *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	Lat              *int64   `protobuf:"zigzag64,8,req,name=lat" json:"lat,omitempty"`
	Lng              *int64   `protobuf:"zigzag64,9,req,name=lng" json:"lng,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Node) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Node) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Node) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Node) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Node) GetLat() int64 {
	if m != nil && m.Lat != nil {
		return *m.Lat
	}
	return 0
}

func (m *Node) GetLng() int64 {
	if m != nil && m.Lng != nil {
		return *m.Lng
	}
	return 0
}

type Info struct {
	Version   *int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Timestamp *int64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// these caan be omitted if the object represents one changeset
	// since they will be all the same. However tests on 200k changesets
	// show this saves about 17 bytes per changeset on average after gzip.
	ChangesetId *int64  `protobuf:"varint,3,opt,name=changeset_id" json:"changeset_id,omitempty"`
	UserId      *int32  `protobuf:"varint,4,opt,name=user_id" json:"user_id,omitempty"`
	UserSid     *uint32 `protobuf:"varint,5,opt,name=user_sid" json:"user_sid,omitempty"`
	// The visible flag is used to store history information. It indicates that
	// the current object version has been created by a delete operation on the
	// OSM API. This info may be omitted if it can be inferred from its group
	// ie. create, modify, delete.
	Visible          *bool  `protobuf:"varint,6,opt,name=visible" json:"visible,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Info) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Info) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Info) GetChangesetId() int64 {
	if m != nil && m.ChangesetId != nil {
		return *m.ChangesetId
	}
	return 0
}

func (m *Info) GetUserId() int32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Info) GetUserSid() uint32 {
	if m != nil && m.UserSid != nil {
		return *m.UserSid
	}
	return 0
}

func (m *Info) GetVisible() bool {
	if m != nil && m.Visible != nil {
		return *m.Visible
	}
	return false
}

type DenseNodes struct {
	Id        []int64    `protobuf:"zigzag64,1,rep,packed,name=id" json:"id,omitempty"`
	DenseInfo *DenseInfo `protobuf:"bytes,5,opt,name=dense_info" json:"dense_info,omitempty"`
	Lat       []int64    `protobuf:"zigzag64,8,rep,packed,name=lat" json:"lat,omitempty"`
	Lng       []int64    `protobuf:"zigzag64,9,rep,packed,name=lng" json:"lng,omitempty"`
	// Special packing of keys and vals into one array. We use a single stringid
	// of 0 to delimit when the tags of a node ends and the tags of the next node
	// begin. The storage pattern is: ((<keyid> <valid>)* '0' )* As an exception,
	// if no node in the current block has any key/value pairs, this array does
	// not contain any delimiters, but is simply empty.
	KeysVals         []uint32 `protobuf:"varint,10,rep,packed,name=keys_vals" json:"keys_vals,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DenseNodes) Reset()                    { *m = DenseNodes{} }
func (m *DenseNodes) String() string            { return proto.CompactTextString(m) }
func (*DenseNodes) ProtoMessage()               {}
func (*DenseNodes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DenseNodes) GetId() []int64 {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DenseNodes) GetDenseInfo() *DenseInfo {
	if m != nil {
		return m.DenseInfo
	}
	return nil
}

func (m *DenseNodes) GetLat() []int64 {
	if m != nil {
		return m.Lat
	}
	return nil
}

func (m *DenseNodes) GetLng() []int64 {
	if m != nil {
		return m.Lng
	}
	return nil
}

func (m *DenseNodes) GetKeysVals() []uint32 {
	if m != nil {
		return m.KeysVals
	}
	return nil
}

type DenseInfo struct {
	Version   []int32 `protobuf:"varint,1,rep,packed,name=version" json:"version,omitempty"`
	Timestamp []int64 `protobuf:"zigzag64,2,rep,packed,name=timestamp" json:"timestamp,omitempty"`
	// these will be omitted if the object represents one changeset
	// and these will be all the same.
	ChangesetId []int64 `protobuf:"zigzag64,3,rep,packed,name=changeset_id" json:"changeset_id,omitempty"`
	UserId      []int32 `protobuf:"zigzag32,4,rep,packed,name=user_id" json:"user_id,omitempty"`
	UserSid     []int32 `protobuf:"zigzag32,5,rep,packed,name=user_sid" json:"user_sid,omitempty"`
	// The visible flag is used to store history information. It indicates that
	// the current object version has been created by a delete operation on the
	// OSM API. This info may be omitted if it can be inferred from it's group
	// ie. create, modify, delete.
	Visible          []bool `protobuf:"varint,6,rep,packed,name=visible" json:"visible,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DenseInfo) Reset()                    { *m = DenseInfo{} }
func (m *DenseInfo) String() string            { return proto.CompactTextString(m) }
func (*DenseInfo) ProtoMessage()               {}
func (*DenseInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DenseInfo) GetVersion() []int32 {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *DenseInfo) GetTimestamp() []int64 {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *DenseInfo) GetChangesetId() []int64 {
	if m != nil {
		return m.ChangesetId
	}
	return nil
}

func (m *DenseInfo) GetUserId() []int32 {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *DenseInfo) GetUserSid() []int32 {
	if m != nil {
		return m.UserSid
	}
	return nil
}

func (m *DenseInfo) GetVisible() []bool {
	if m != nil {
		return m.Visible
	}
	return nil
}

type Way struct {
	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys             []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals             []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info             *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	Refs             []int64  `protobuf:"zigzag64,8,rep,packed,name=refs" json:"refs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Way) Reset()                    { *m = Way{} }
func (m *Way) String() string            { return proto.CompactTextString(m) }
func (*Way) ProtoMessage()               {}
func (*Way) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Way) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Way) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Way) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Way) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Way) GetRefs() []int64 {
	if m != nil {
		return m.Refs
	}
	return nil
}

type Relation struct {
	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	// Parallel arrays.
	Keys []uint32 `protobuf:"varint,2,rep,packed,name=keys" json:"keys,omitempty"`
	Vals []uint32 `protobuf:"varint,3,rep,packed,name=vals" json:"vals,omitempty"`
	Info *Info    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	// Parallel arrays
	// Roles has been changed int32 -> uint32 form the osm proto,
	// this is for consistency and backwards compatible.
	Roles            []uint32              `protobuf:"varint,8,rep,packed,name=roles" json:"roles,omitempty"`
	Refs             []int64               `protobuf:"zigzag64,9,rep,packed,name=refs" json:"refs,omitempty"`
	Types            []Relation_MemberType `protobuf:"varint,10,rep,packed,name=types,enum=osm.Relation_MemberType" json:"types,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Relation) Reset()                    { *m = Relation{} }
func (m *Relation) String() string            { return proto.CompactTextString(m) }
func (*Relation) ProtoMessage()               {}
func (*Relation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Relation) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Relation) GetKeys() []uint32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Relation) GetVals() []uint32 {
	if m != nil {
		return m.Vals
	}
	return nil
}

func (m *Relation) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Relation) GetRoles() []uint32 {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Relation) GetRefs() []int64 {
	if m != nil {
		return m.Refs
	}
	return nil
}

func (m *Relation) GetTypes() []Relation_MemberType {
	if m != nil {
		return m.Types
	}
	return nil
}

func init() {
	proto.RegisterType((*Changeset)(nil), "osm.Changeset")
	proto.RegisterType((*Bounds)(nil), "osm.Bounds")
	proto.RegisterType((*Change)(nil), "osm.Change")
	proto.RegisterType((*OSM)(nil), "osm.OSM")
	proto.RegisterType((*Node)(nil), "osm.Node")
	proto.RegisterType((*Info)(nil), "osm.Info")
	proto.RegisterType((*DenseNodes)(nil), "osm.DenseNodes")
	proto.RegisterType((*DenseInfo)(nil), "osm.DenseInfo")
	proto.RegisterType((*Way)(nil), "osm.Way")
	proto.RegisterType((*Relation)(nil), "osm.Relation")
	proto.RegisterEnum("osm.Relation_MemberType", Relation_MemberType_name, Relation_MemberType_value)
}

var fileDescriptor0 = []byte{
	// 648 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0xec, 0xb1, 0x13, 0xfb, 0xba, 0x3f, 0xe9, 0xb4, 0x1f, 0x0c, 0x62, 0x53, 0x59, 0x2c,
	0x2c, 0x21, 0xba, 0xc8, 0x1b, 0xb4, 0xb4, 0x8b, 0x4a, 0xfd, 0x91, 0xa0, 0x12, 0x82, 0x4d, 0x70,
	0xe2, 0x69, 0xb1, 0xb0, 0x3d, 0x91, 0xc7, 0x2d, 0xcd, 0x96, 0x37, 0xe0, 0x41, 0x78, 0x1b, 0xd6,
	0x3c, 0x0b, 0x77, 0xae, 0xed, 0xd8, 0xc1, 0x1b, 0x16, 0xdd, 0xc5, 0xe7, 0xdc, 0x7b, 0xe7, 0x9c,
	0x73, 0x6f, 0xc0, 0x57, 0x3a, 0x3f, 0x5a, 0x96, 0xaa, 0x52, 0x9c, 0xe1, 0xcf, 0xf0, 0xb7, 0x05,
	0xfe, 0xdb, 0x2f, 0x71, 0x71, 0x27, 0xb5, 0xac, 0x38, 0x80, 0x9d, 0x26, 0xc2, 0x3a, 0xb4, 0x22,
	0xc6, 0x27, 0xe0, 0x7c, 0x95, 0x2b, 0x2d, 0xec, 0x43, 0x16, 0x6d, 0x9f, 0xd8, 0x13, 0xcb, 0x20,
	0x0f, 0x71, 0xa6, 0x05, 0x5b, 0x23, 0xbb, 0x30, 0xbe, 0xd7, 0xb2, 0x9c, 0x61, 0x93, 0x8b, 0x4d,
	0x2e, 0x96, 0x78, 0x04, 0x68, 0x44, 0x46, 0x88, 0x6c, 0x73, 0x9c, 0xb9, 0x28, 0x65, 0x5c, 0xc9,
	0x64, 0x16, 0x57, 0x62, 0x4c, 0xa3, 0xf7, 0xc0, 0x5f, 0x64, 0x4a, 0xd7, 0x90, 0x47, 0xd0, 0x16,
	0x38, 0x6a, 0x29, 0x0b, 0xe1, 0xe3, 0x97, 0xc7, 0x5f, 0xc2, 0x68, 0xae, 0xee, 0x8b, 0x44, 0x0b,
	0xc0, 0xef, 0x60, 0x1a, 0x1c, 0x19, 0xd9, 0x27, 0x04, 0x19, 0x72, 0x41, 0x8a, 0x45, 0xd0, 0x23,
	0x6b, 0x13, 0x46, 0x91, 0xae, 0xca, 0xb4, 0xb8, 0xd3, 0xe2, 0x00, 0x65, 0xfa, 0xe1, 0x05, 0x8c,
	0x9a, 0x3e, 0xa4, 0xf2, 0xb4, 0x98, 0x65, 0xc5, 0x1d, 0x3a, 0xb4, 0x23, 0x4e, 0x40, 0xfc, 0x48,
	0x80, 0xbd, 0x06, 0x4c, 0x05, 0xaa, 0x62, 0x1b, 0x15, 0x08, 0x38, 0x06, 0x08, 0xbf, 0x5b, 0x30,
	0x6a, 0x5e, 0x12, 0x28, 0x83, 0x8c, 0x51, 0x5e, 0xc1, 0xd4, 0x23, 0x19, 0xd7, 0xef, 0x2f, 0x0d,
	0x93, 0xab, 0x24, 0xbd, 0x5d, 0xe1, 0xd8, 0x01, 0x93, 0xc8, 0x4c, 0x62, 0x0f, 0xfb, 0x8b, 0x79,
	0x01, 0xe3, 0x85, 0x2a, 0x2a, 0xf9, 0x68, 0x5e, 0xda, 0xa4, 0x06, 0x96, 0x7e, 0x5a, 0xc0, 0x0c,
	0xd1, 0xa5, 0x64, 0x0d, 0x53, 0x12, 0xe0, 0x16, 0x2a, 0x91, 0xf5, 0xfe, 0x82, 0xa9, 0x4f, 0xdc,
	0x15, 0x22, 0xfc, 0x15, 0x04, 0x89, 0x2c, 0xb4, 0x9c, 0xd5, 0x7c, 0xad, 0x64, 0x97, 0xf8, 0x53,
	0x83, 0x9b, 0x22, 0xcd, 0x9f, 0x81, 0xf3, 0x2d, 0xc6, 0xf5, 0x3b, 0xd4, 0x5e, 0xab, 0xf9, 0x10,
	0xaf, 0xf8, 0x21, 0xf8, 0xa5, 0xc4, 0x40, 0x52, 0x55, 0x68, 0x5c, 0xba, 0x21, 0xb7, 0x89, 0x7c,
	0xd7, 0xa0, 0x43, 0xbd, 0x29, 0x38, 0xf4, 0x70, 0x7b, 0x5d, 0xf6, 0x3f, 0x5e, 0xd7, 0x73, 0x70,
	0xd2, 0xe2, 0x56, 0x35, 0x81, 0xd4, 0x0e, 0xce, 0x11, 0xe0, 0x01, 0xb0, 0x8c, 0x2e, 0xc7, 0xec,
	0xc8, 0x7c, 0xe0, 0x06, 0x7d, 0xda, 0x4f, 0x05, 0x0e, 0x55, 0xa0, 0x86, 0x07, 0x59, 0x6a, 0x94,
	0x43, 0xd9, 0xb8, 0xe6, 0xe4, 0xaa, 0x34, 0x97, 0xba, 0x8a, 0xf3, 0x25, 0xad, 0x85, 0xf1, 0x03,
	0xd8, 0x5a, 0xb4, 0x97, 0x6f, 0x2e, 0x98, 0x11, 0xda, 0x3b, 0x69, 0x67, 0x70, 0xd2, 0x2e, 0x9d,
	0xb4, 0x19, 0x9e, 0xea, 0x74, 0x9e, 0x49, 0xba, 0x71, 0x2f, 0x5c, 0x01, 0xf4, 0x92, 0xdb, 0x69,
	0x6c, 0xb2, 0x88, 0x93, 0x8d, 0x10, 0xa0, 0xce, 0x9b, 0xcc, 0xb8, 0x64, 0x66, 0xa7, 0x8b, 0xbb,
	0xd1, 0xdb, 0x38, 0x6a, 0x9b, 0x76, 0x5b, 0x57, 0x2d, 0xf0, 0x3f, 0xf8, 0x26, 0xb0, 0x19, 0x65,
	0x04, 0x6d, 0x46, 0xe1, 0x0f, 0xfc, 0xff, 0x76, 0x63, 0xf6, 0xfb, 0xb6, 0x59, 0xe4, 0xb6, 0x9d,
	0x7d, 0xeb, 0xed, 0x40, 0x31, 0xb0, 0xdf, 0x32, 0xfb, 0xfd, 0x08, 0x58, 0xb4, 0x47, 0xe0, 0xc1,
	0x46, 0x0c, 0x2d, 0xba, 0xdf, 0x8f, 0x82, 0x45, 0x1e, 0x69, 0xfa, 0x0c, 0xcc, 0x5c, 0xca, 0x93,
	0xad, 0x1b, 0x4b, 0x4b, 0x79, 0xab, 0xbb, 0x74, 0xc2, 0x5f, 0x16, 0x78, 0xeb, 0x7b, 0x7b, 0xb2,
	0x77, 0xf6, 0xc0, 0x2d, 0x55, 0x26, 0xeb, 0x87, 0xd6, 0xdd, 0xf4, 0x74, 0xb7, 0x87, 0xd7, 0xe0,
	0x56, 0xab, 0xa5, 0xac, 0x77, 0xb0, 0x33, 0x15, 0x1b, 0xb7, 0x7f, 0x74, 0x29, 0xf3, 0xb9, 0x2c,
	0x6f, 0xb0, 0x80, 0x74, 0xbe, 0x01, 0xe8, 0x10, 0xee, 0xe1, 0xff, 0xe0, 0xfa, 0xf4, 0x6c, 0xf2,
	0x1f, 0x1f, 0x63, 0x42, 0xc7, 0x1f, 0x71, 0xda, 0x16, 0xfa, 0x38, 0xbb, 0x38, 0xbe, 0x39, 0xbf,
	0xbe, 0x9a, 0xd8, 0x27, 0xe3, 0x4f, 0x2e, 0x4e, 0x5b, 0xce, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x07, 0xe6, 0x63, 0x45, 0xa6, 0x05, 0x00, 0x00,
}
