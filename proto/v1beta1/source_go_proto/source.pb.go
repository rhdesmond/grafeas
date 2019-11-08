// Code generated by protoc-gen-go. DO NOT EDIT.
// source: source.proto

package source_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// The type of an alias.
type AliasContext_Kind int32

const (
	// Unknown.
	AliasContext_KIND_UNSPECIFIED AliasContext_Kind = 0
	// Git tag.
	AliasContext_FIXED AliasContext_Kind = 1
	// Git branch.
	AliasContext_MOVABLE AliasContext_Kind = 2
	// Used to specify non-standard aliases. For example, if a Git repo has a
	// ref named "refs/foo/bar".
	AliasContext_OTHER AliasContext_Kind = 4
)

var AliasContext_Kind_name = map[int32]string{
	0: "KIND_UNSPECIFIED",
	1: "FIXED",
	2: "MOVABLE",
	4: "OTHER",
}

var AliasContext_Kind_value = map[string]int32{
	"KIND_UNSPECIFIED": 0,
	"FIXED":            1,
	"MOVABLE":          2,
	"OTHER":            4,
}

func (x AliasContext_Kind) String() string {
	return proto.EnumName(AliasContext_Kind_name, int32(x))
}

func (AliasContext_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9229c9b6eb8e6b85, []int{1, 0}
}

// A SourceContext is a reference to a tree of files. A SourceContext together
// with a path point to a unique revision of a single file or directory.
type SourceContext struct {
	// A SourceContext can refer any one of the following types of repositories.
	//
	// Types that are valid to be assigned to Context:
	//	*SourceContext_CloudRepo
	//	*SourceContext_Gerrit
	//	*SourceContext_Git
	Context isSourceContext_Context `protobuf_oneof:"context"`
	// Labels with user defined metadata.
	Labels               map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SourceContext) Reset()         { *m = SourceContext{} }
func (m *SourceContext) String() string { return proto.CompactTextString(m) }
func (*SourceContext) ProtoMessage()    {}
func (*SourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_9229c9b6eb8e6b85, []int{0}
}

func (m *SourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceContext.Unmarshal(m, b)
}
func (m *SourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceContext.Marshal(b, m, deterministic)
}
func (m *SourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceContext.Merge(m, src)
}
func (m *SourceContext) XXX_Size() int {
	return xxx_messageInfo_SourceContext.Size(m)
}
func (m *SourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_SourceContext proto.InternalMessageInfo

type isSourceContext_Context interface {
	isSourceContext_Context()
}

type SourceContext_CloudRepo struct {
	CloudRepo *CloudRepoSourceContext `protobuf:"bytes,1,opt,name=cloud_repo,json=cloudRepo,proto3,oneof"`
}

type SourceContext_Gerrit struct {
	Gerrit *GerritSourceContext `protobuf:"bytes,2,opt,name=gerrit,proto3,oneof"`
}

type SourceContext_Git struct {
	Git *GitSourceContext `protobuf:"bytes,3,opt,name=git,proto3,oneof"`
}

func (*SourceContext_CloudRepo) isSourceContext_Context() {}

func (*SourceContext_Gerrit) isSourceContext_Context() {}

func (*SourceContext_Git) isSourceContext_Context() {}

func (m *SourceContext) GetContext() isSourceContext_Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *SourceContext) GetCloudRepo() *CloudRepoSourceContext {
	if x, ok := m.GetContext().(*SourceContext_CloudRepo); ok {
		return x.CloudRepo
	}
	return nil
}

func (m *SourceContext) GetGerrit() *GerritSourceContext {
	if x, ok := m.GetContext().(*SourceContext_Gerrit); ok {
		return x.Gerrit
	}
	return nil
}

func (m *SourceContext) GetGit() *GitSourceContext {
	if x, ok := m.GetContext().(*SourceContext_Git); ok {
		return x.Git
	}
	return nil
}

func (m *SourceContext) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SourceContext) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SourceContext_CloudRepo)(nil),
		(*SourceContext_Gerrit)(nil),
		(*SourceContext_Git)(nil),
	}
}

// An alias to a repo revision.
type AliasContext struct {
	// The alias kind.
	Kind AliasContext_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=grafeas.v1beta1.source.AliasContext_Kind" json:"kind,omitempty"`
	// The alias name.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AliasContext) Reset()         { *m = AliasContext{} }
func (m *AliasContext) String() string { return proto.CompactTextString(m) }
func (*AliasContext) ProtoMessage()    {}
func (*AliasContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_9229c9b6eb8e6b85, []int{1}
}

func (m *AliasContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AliasContext.Unmarshal(m, b)
}
func (m *AliasContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AliasContext.Marshal(b, m, deterministic)
}
func (m *AliasContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AliasContext.Merge(m, src)
}
func (m *AliasContext) XXX_Size() int {
	return xxx_messageInfo_AliasContext.Size(m)
}
func (m *AliasContext) XXX_DiscardUnknown() {
	xxx_messageInfo_AliasContext.DiscardUnknown(m)
}

var xxx_messageInfo_AliasContext proto.InternalMessageInfo

func (m *AliasContext) GetKind() AliasContext_Kind {
	if m != nil {
		return m.Kind
	}
	return AliasContext_KIND_UNSPECIFIED
}

func (m *AliasContext) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A CloudRepoSourceContext denotes a particular revision in a Google Cloud
// Source Repo.
type CloudRepoSourceContext struct {
	// The ID of the repo.
	RepoId *RepoId `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	// A revision in a Cloud Repo can be identified by either its revision ID or
	// its alias.
	//
	// Types that are valid to be assigned to Revision:
	//	*CloudRepoSourceContext_RevisionId
	//	*CloudRepoSourceContext_AliasContext
	Revision             isCloudRepoSourceContext_Revision `protobuf_oneof:"revision"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *CloudRepoSourceContext) Reset()         { *m = CloudRepoSourceContext{} }
func (m *CloudRepoSourceContext) String() string { return proto.CompactTextString(m) }
func (*CloudRepoSourceContext) ProtoMessage()    {}
func (*CloudRepoSourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_9229c9b6eb8e6b85, []int{2}
}

func (m *CloudRepoSourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudRepoSourceContext.Unmarshal(m, b)
}
func (m *CloudRepoSourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudRepoSourceContext.Marshal(b, m, deterministic)
}
func (m *CloudRepoSourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudRepoSourceContext.Merge(m, src)
}
func (m *CloudRepoSourceContext) XXX_Size() int {
	return xxx_messageInfo_CloudRepoSourceContext.Size(m)
}
func (m *CloudRepoSourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudRepoSourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_CloudRepoSourceContext proto.InternalMessageInfo

func (m *CloudRepoSourceContext) GetRepoId() *RepoId {
	if m != nil {
		return m.RepoId
	}
	return nil
}

type isCloudRepoSourceContext_Revision interface {
	isCloudRepoSourceContext_Revision()
}

type CloudRepoSourceContext_RevisionId struct {
	RevisionId string `protobuf:"bytes,2,opt,name=revision_id,json=revisionId,proto3,oneof"`
}

type CloudRepoSourceContext_AliasContext struct {
	AliasContext *AliasContext `protobuf:"bytes,3,opt,name=alias_context,json=aliasContext,proto3,oneof"`
}

func (*CloudRepoSourceContext_RevisionId) isCloudRepoSourceContext_Revision() {}

func (*CloudRepoSourceContext_AliasContext) isCloudRepoSourceContext_Revision() {}

func (m *CloudRepoSourceContext) GetRevision() isCloudRepoSourceContext_Revision {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *CloudRepoSourceContext) GetRevisionId() string {
	if x, ok := m.GetRevision().(*CloudRepoSourceContext_RevisionId); ok {
		return x.RevisionId
	}
	return ""
}

func (m *CloudRepoSourceContext) GetAliasContext() *AliasContext {
	if x, ok := m.GetRevision().(*CloudRepoSourceContext_AliasContext); ok {
		return x.AliasContext
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CloudRepoSourceContext) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CloudRepoSourceContext_RevisionId)(nil),
		(*CloudRepoSourceContext_AliasContext)(nil),
	}
}

// A SourceContext referring to a Gerrit project.
type GerritSourceContext struct {
	// The URI of a running Gerrit instance.
	HostUri string `protobuf:"bytes,1,opt,name=host_uri,json=hostUri,proto3" json:"host_uri,omitempty"`
	// The full project name within the host. Projects may be nested, so
	// "project/subproject" is a valid project name. The "repo name" is the
	// hostURI/project.
	GerritProject string `protobuf:"bytes,2,opt,name=gerrit_project,json=gerritProject,proto3" json:"gerrit_project,omitempty"`
	// A revision in a Gerrit project can be identified by either its revision ID
	// or its alias.
	//
	// Types that are valid to be assigned to Revision:
	//	*GerritSourceContext_RevisionId
	//	*GerritSourceContext_AliasContext
	Revision             isGerritSourceContext_Revision `protobuf_oneof:"revision"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GerritSourceContext) Reset()         { *m = GerritSourceContext{} }
func (m *GerritSourceContext) String() string { return proto.CompactTextString(m) }
func (*GerritSourceContext) ProtoMessage()    {}
func (*GerritSourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_9229c9b6eb8e6b85, []int{3}
}

func (m *GerritSourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GerritSourceContext.Unmarshal(m, b)
}
func (m *GerritSourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GerritSourceContext.Marshal(b, m, deterministic)
}
func (m *GerritSourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GerritSourceContext.Merge(m, src)
}
func (m *GerritSourceContext) XXX_Size() int {
	return xxx_messageInfo_GerritSourceContext.Size(m)
}
func (m *GerritSourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_GerritSourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_GerritSourceContext proto.InternalMessageInfo

func (m *GerritSourceContext) GetHostUri() string {
	if m != nil {
		return m.HostUri
	}
	return ""
}

func (m *GerritSourceContext) GetGerritProject() string {
	if m != nil {
		return m.GerritProject
	}
	return ""
}

type isGerritSourceContext_Revision interface {
	isGerritSourceContext_Revision()
}

type GerritSourceContext_RevisionId struct {
	RevisionId string `protobuf:"bytes,3,opt,name=revision_id,json=revisionId,proto3,oneof"`
}

type GerritSourceContext_AliasContext struct {
	AliasContext *AliasContext `protobuf:"bytes,4,opt,name=alias_context,json=aliasContext,proto3,oneof"`
}

func (*GerritSourceContext_RevisionId) isGerritSourceContext_Revision() {}

func (*GerritSourceContext_AliasContext) isGerritSourceContext_Revision() {}

func (m *GerritSourceContext) GetRevision() isGerritSourceContext_Revision {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *GerritSourceContext) GetRevisionId() string {
	if x, ok := m.GetRevision().(*GerritSourceContext_RevisionId); ok {
		return x.RevisionId
	}
	return ""
}

func (m *GerritSourceContext) GetAliasContext() *AliasContext {
	if x, ok := m.GetRevision().(*GerritSourceContext_AliasContext); ok {
		return x.AliasContext
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GerritSourceContext) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GerritSourceContext_RevisionId)(nil),
		(*GerritSourceContext_AliasContext)(nil),
	}
}

// A GitSourceContext denotes a particular revision in a third party Git
// repository (e.g., GitHub).
type GitSourceContext struct {
	// Git repository URL.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Git commit hash.
	RevisionId           string   `protobuf:"bytes,2,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GitSourceContext) Reset()         { *m = GitSourceContext{} }
func (m *GitSourceContext) String() string { return proto.CompactTextString(m) }
func (*GitSourceContext) ProtoMessage()    {}
func (*GitSourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_9229c9b6eb8e6b85, []int{4}
}

func (m *GitSourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitSourceContext.Unmarshal(m, b)
}
func (m *GitSourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitSourceContext.Marshal(b, m, deterministic)
}
func (m *GitSourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitSourceContext.Merge(m, src)
}
func (m *GitSourceContext) XXX_Size() int {
	return xxx_messageInfo_GitSourceContext.Size(m)
}
func (m *GitSourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_GitSourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_GitSourceContext proto.InternalMessageInfo

func (m *GitSourceContext) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *GitSourceContext) GetRevisionId() string {
	if m != nil {
		return m.RevisionId
	}
	return ""
}

// A unique identifier for a Cloud Repo.
type RepoId struct {
	// A cloud repo can be identified by either its project ID and repository name
	// combination, or its globally unique identifier.
	//
	// Types that are valid to be assigned to Id:
	//	*RepoId_ProjectRepoId
	//	*RepoId_Uid
	Id                   isRepoId_Id `protobuf_oneof:"id"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RepoId) Reset()         { *m = RepoId{} }
func (m *RepoId) String() string { return proto.CompactTextString(m) }
func (*RepoId) ProtoMessage()    {}
func (*RepoId) Descriptor() ([]byte, []int) {
	return fileDescriptor_9229c9b6eb8e6b85, []int{5}
}

func (m *RepoId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepoId.Unmarshal(m, b)
}
func (m *RepoId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepoId.Marshal(b, m, deterministic)
}
func (m *RepoId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepoId.Merge(m, src)
}
func (m *RepoId) XXX_Size() int {
	return xxx_messageInfo_RepoId.Size(m)
}
func (m *RepoId) XXX_DiscardUnknown() {
	xxx_messageInfo_RepoId.DiscardUnknown(m)
}

var xxx_messageInfo_RepoId proto.InternalMessageInfo

type isRepoId_Id interface {
	isRepoId_Id()
}

type RepoId_ProjectRepoId struct {
	ProjectRepoId *ProjectRepoId `protobuf:"bytes,1,opt,name=project_repo_id,json=projectRepoId,proto3,oneof"`
}

type RepoId_Uid struct {
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3,oneof"`
}

func (*RepoId_ProjectRepoId) isRepoId_Id() {}

func (*RepoId_Uid) isRepoId_Id() {}

func (m *RepoId) GetId() isRepoId_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *RepoId) GetProjectRepoId() *ProjectRepoId {
	if x, ok := m.GetId().(*RepoId_ProjectRepoId); ok {
		return x.ProjectRepoId
	}
	return nil
}

func (m *RepoId) GetUid() string {
	if x, ok := m.GetId().(*RepoId_Uid); ok {
		return x.Uid
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RepoId) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RepoId_ProjectRepoId)(nil),
		(*RepoId_Uid)(nil),
	}
}

// Selects a repo using a Google Cloud Platform project ID (e.g.,
// winged-cargo-31) and a repo name within that project.
type ProjectRepoId struct {
	// The ID of the project.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The name of the repo. Leave empty for the default repo.
	RepoName             string   `protobuf:"bytes,2,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectRepoId) Reset()         { *m = ProjectRepoId{} }
func (m *ProjectRepoId) String() string { return proto.CompactTextString(m) }
func (*ProjectRepoId) ProtoMessage()    {}
func (*ProjectRepoId) Descriptor() ([]byte, []int) {
	return fileDescriptor_9229c9b6eb8e6b85, []int{6}
}

func (m *ProjectRepoId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectRepoId.Unmarshal(m, b)
}
func (m *ProjectRepoId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectRepoId.Marshal(b, m, deterministic)
}
func (m *ProjectRepoId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectRepoId.Merge(m, src)
}
func (m *ProjectRepoId) XXX_Size() int {
	return xxx_messageInfo_ProjectRepoId.Size(m)
}
func (m *ProjectRepoId) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectRepoId.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectRepoId proto.InternalMessageInfo

func (m *ProjectRepoId) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ProjectRepoId) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func init() {
	proto.RegisterEnum("grafeas.v1beta1.source.AliasContext_Kind", AliasContext_Kind_name, AliasContext_Kind_value)
	proto.RegisterType((*SourceContext)(nil), "grafeas.v1beta1.source.SourceContext")
	proto.RegisterMapType((map[string]string)(nil), "grafeas.v1beta1.source.SourceContext.LabelsEntry")
	proto.RegisterType((*AliasContext)(nil), "grafeas.v1beta1.source.AliasContext")
	proto.RegisterType((*CloudRepoSourceContext)(nil), "grafeas.v1beta1.source.CloudRepoSourceContext")
	proto.RegisterType((*GerritSourceContext)(nil), "grafeas.v1beta1.source.GerritSourceContext")
	proto.RegisterType((*GitSourceContext)(nil), "grafeas.v1beta1.source.GitSourceContext")
	proto.RegisterType((*RepoId)(nil), "grafeas.v1beta1.source.RepoId")
	proto.RegisterType((*ProjectRepoId)(nil), "grafeas.v1beta1.source.ProjectRepoId")
}

func init() { proto.RegisterFile("source.proto", fileDescriptor_9229c9b6eb8e6b85) }

var fileDescriptor_9229c9b6eb8e6b85 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xef, 0x4f, 0xd3, 0x40,
	0x18, 0x5e, 0xd7, 0xb2, 0xd1, 0x77, 0x0c, 0x9b, 0x93, 0x90, 0xa1, 0x51, 0xb1, 0x91, 0x04, 0x63,
	0x52, 0x02, 0x7e, 0x10, 0x8d, 0xc6, 0x30, 0x28, 0xac, 0x19, 0x32, 0x72, 0x88, 0x31, 0x26, 0xa6,
	0xe9, 0xd6, 0x73, 0x9c, 0x94, 0xde, 0x72, 0x6d, 0x89, 0xfc, 0x3b, 0xfc, 0x3f, 0x7e, 0xf2, 0x7f,
	0xf1, 0xb3, 0xb9, 0xeb, 0xd5, 0x74, 0x73, 0x4d, 0xf8, 0xe0, 0xa7, 0xde, 0xfb, 0xdc, 0xfb, 0x3c,
	0xf7, 0xfe, 0x78, 0x52, 0x58, 0x4a, 0x58, 0xc6, 0x47, 0xc4, 0x99, 0x70, 0x96, 0x32, 0xb4, 0x3a,
	0xe6, 0xc1, 0x37, 0x12, 0x24, 0xce, 0xf5, 0xf6, 0x90, 0xa4, 0xc1, 0xb6, 0x93, 0xdf, 0xda, 0xbf,
	0xeb, 0xd0, 0x3e, 0x93, 0xc7, 0x7d, 0x16, 0xa7, 0xe4, 0x47, 0x8a, 0x06, 0x00, 0xa3, 0x88, 0x65,
	0xa1, 0xcf, 0xc9, 0x84, 0x75, 0xb4, 0x75, 0x6d, 0xb3, 0xb5, 0xe3, 0x38, 0xf3, 0xe9, 0xce, 0xbe,
	0xc8, 0xc4, 0x64, 0xc2, 0xa6, 0x34, 0x7a, 0x35, 0x6c, 0x8e, 0x8a, 0x1b, 0xe4, 0x42, 0x63, 0x4c,
	0x38, 0xa7, 0x69, 0xa7, 0x2e, 0xc5, 0x5e, 0x54, 0x89, 0x1d, 0xc9, 0xac, 0x59, 0x25, 0x45, 0x46,
	0x6f, 0x41, 0x1f, 0xd3, 0xb4, 0xa3, 0x4b, 0x8d, 0xcd, 0x4a, 0x8d, 0x7f, 0x05, 0x04, 0x0d, 0x79,
	0xd0, 0x88, 0x82, 0x21, 0x89, 0x92, 0x8e, 0xb1, 0xae, 0x6f, 0xb6, 0x76, 0xb6, 0xab, 0x04, 0xa6,
	0xd8, 0xce, 0xb1, 0xe4, 0xb8, 0x71, 0xca, 0x6f, 0xb0, 0x12, 0x78, 0xf0, 0x1a, 0x5a, 0x25, 0x18,
	0x59, 0xa0, 0x5f, 0x92, 0x1b, 0x39, 0x28, 0x13, 0x8b, 0x23, 0x5a, 0x81, 0x85, 0xeb, 0x20, 0xca,
	0x88, 0xec, 0xd7, 0xc4, 0x79, 0xf0, 0xa6, 0xbe, 0xab, 0x75, 0x4d, 0x68, 0x8e, 0x72, 0x65, 0xfb,
	0x56, 0x83, 0xa5, 0xbd, 0x88, 0x06, 0x49, 0x31, 0xf7, 0x77, 0x60, 0x5c, 0xd2, 0x38, 0x94, 0x42,
	0xcb, 0x3b, 0xcf, 0xab, 0xea, 0x2b, 0x73, 0x9c, 0x3e, 0x8d, 0x43, 0x2c, 0x69, 0x08, 0x81, 0x11,
	0x07, 0x57, 0xc5, 0x9b, 0xf2, 0x6c, 0xbf, 0x07, 0x43, 0x64, 0xa0, 0x15, 0xb0, 0xfa, 0xde, 0xc9,
	0x81, 0x7f, 0x7e, 0x72, 0x76, 0xea, 0xee, 0x7b, 0x87, 0x9e, 0x7b, 0x60, 0xd5, 0x90, 0x09, 0x0b,
	0x87, 0xde, 0x67, 0xf7, 0xc0, 0xd2, 0x50, 0x0b, 0x9a, 0x1f, 0x06, 0x9f, 0xf6, 0xba, 0xc7, 0xae,
	0x55, 0x17, 0xf8, 0xe0, 0x63, 0xcf, 0xc5, 0x96, 0x61, 0xff, 0xd4, 0x60, 0x75, 0xfe, 0x8a, 0xd1,
	0x2b, 0x68, 0x0a, 0x83, 0xf8, 0x34, 0x54, 0x1e, 0x79, 0x5c, 0x55, 0xb1, 0xe0, 0x7a, 0x21, 0x6e,
	0x70, 0xf9, 0x45, 0x4f, 0xa1, 0xc5, 0xc9, 0x35, 0x4d, 0x28, 0x8b, 0x05, 0x59, 0xd6, 0xdb, 0xab,
	0x61, 0x28, 0x40, 0x2f, 0x44, 0x7d, 0x68, 0x07, 0xa2, 0x4d, 0x5f, 0x0d, 0x4b, 0x2d, 0xfd, 0xd9,
	0x5d, 0x66, 0xd2, 0xab, 0xe1, 0xa5, 0xa0, 0x14, 0x77, 0x01, 0x16, 0x0b, 0x69, 0xfb, 0x97, 0x06,
	0xf7, 0xe7, 0xb8, 0x0c, 0xad, 0xc1, 0xe2, 0x05, 0x4b, 0x52, 0x3f, 0xe3, 0x54, 0x2d, 0xb2, 0x29,
	0xe2, 0x73, 0x4e, 0xd1, 0x06, 0x2c, 0xe7, 0x06, 0xf4, 0x27, 0x9c, 0x7d, 0x27, 0xa3, 0x54, 0x4d,
	0xb8, 0x9d, 0xa3, 0xa7, 0x39, 0x38, 0xdb, 0x95, 0x7e, 0x97, 0xae, 0x8c, 0xff, 0xd4, 0x95, 0x0b,
	0xd6, 0xac, 0xed, 0x85, 0x2b, 0x33, 0x1e, 0x15, 0xae, 0xcc, 0x78, 0x84, 0x9e, 0xcc, 0x99, 0x7b,
	0xb9, 0x3e, 0x3b, 0x81, 0x46, 0xbe, 0x2a, 0x34, 0x80, 0x7b, 0xaa, 0x59, 0x7f, 0x7a, 0xc7, 0x1b,
	0x55, 0xb5, 0xaa, 0x31, 0xe4, 0xfc, 0x5e, 0x0d, 0xb7, 0x27, 0x65, 0x00, 0x21, 0xd0, 0xb3, 0xd2,
	0xae, 0x45, 0xd0, 0x35, 0xa0, 0x4e, 0x43, 0xbb, 0x0f, 0xed, 0x29, 0x2e, 0x7a, 0x04, 0x50, 0xbc,
	0xad, 0x9e, 0x35, 0xb1, 0xa9, 0x10, 0x2f, 0x44, 0x0f, 0xc1, 0x94, 0x25, 0x95, 0xbc, 0xbe, 0x28,
	0x80, 0x93, 0xe0, 0x8a, 0x74, 0xbf, 0xc2, 0x1a, 0x65, 0x15, 0x25, 0x9e, 0x6a, 0x5f, 0x76, 0xc7,
	0x34, 0xbd, 0xc8, 0x86, 0xce, 0x88, 0x5d, 0x6d, 0xa9, 0xa4, 0xbf, 0x5f, 0xf9, 0x97, 0xdc, 0x52,
	0x94, 0xad, 0x9c, 0xe2, 0x8f, 0x99, 0x2f, 0xf1, 0xdb, 0xba, 0x7e, 0x84, 0xf7, 0x86, 0x0d, 0x19,
	0xbc, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x99, 0x45, 0x92, 0x73, 0x5a, 0x05, 0x00, 0x00,
}