// Code generated by protoc-gen-go.
// source: pfs/pfs.proto
// DO NOT EDIT!

/*
Package pfs is a generated protocol buffer package.

It is generated from these files:
	pfs/pfs.proto

It has these top-level messages:
	Error
	Repository
	Commit
	Path
	FileInfo
	Shard
	CommitInfo
	Version
	GetVersionResponse
	InitRepositoryRequest
	GetFileRequest
	GetFileInfoRequest
	GetFileInfoResponse
	MakeDirectoryRequest
	PutFileRequest
	ListFilesRequest
	ListFilesResponse
	BranchRequest
	BranchResponse
	CommitRequest
	GetCommitInfoRequest
	GetCommitInfoResponse
	ListCommitsRequest
	ListCommitsResponse
	PullDiffRequest
	PushDiffRequest
*/
package pfs

import proto "github.com/golang/protobuf/proto"
import google_protobuf "github.com/peter-edge/go-google-protobuf"
import google_protobuf1 "github.com/peter-edge/go-google-protobuf"
import google_protobuf2 "github.com/peter-edge/go-google-protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type ErrorCode int32

const (
	ErrorCode_ERROR_CODE_NONE                         ErrorCode = 0
	ErrorCode_ERROR_CODE_UNKNOWN                      ErrorCode = 1
	ErrorCode_ERROR_CODE_DISCOVERY_NETWORK_FAILURE    ErrorCode = 2
	ErrorCode_ERROR_CODE_DISCOVERY_NOT_FOUND          ErrorCode = 3
	ErrorCode_ERROR_CODE_DISCOVERY_NOT_DIRECTORY      ErrorCode = 4
	ErrorCode_ERROR_CODE_DISCOVERY_NOT_VALUE          ErrorCode = 5
	ErrorCode_ERROR_CODE_DISCOVERY_KEY_ALREADY_EXISTS ErrorCode = 6
	ErrorCode_ERROR_CODE_PRECONDITION_NOT_MET         ErrorCode = 7
)

var ErrorCode_name = map[int32]string{
	0: "ERROR_CODE_NONE",
	1: "ERROR_CODE_UNKNOWN",
	2: "ERROR_CODE_DISCOVERY_NETWORK_FAILURE",
	3: "ERROR_CODE_DISCOVERY_NOT_FOUND",
	4: "ERROR_CODE_DISCOVERY_NOT_DIRECTORY",
	5: "ERROR_CODE_DISCOVERY_NOT_VALUE",
	6: "ERROR_CODE_DISCOVERY_KEY_ALREADY_EXISTS",
	7: "ERROR_CODE_PRECONDITION_NOT_MET",
}
var ErrorCode_value = map[string]int32{
	"ERROR_CODE_NONE":                         0,
	"ERROR_CODE_UNKNOWN":                      1,
	"ERROR_CODE_DISCOVERY_NETWORK_FAILURE":    2,
	"ERROR_CODE_DISCOVERY_NOT_FOUND":          3,
	"ERROR_CODE_DISCOVERY_NOT_DIRECTORY":      4,
	"ERROR_CODE_DISCOVERY_NOT_VALUE":          5,
	"ERROR_CODE_DISCOVERY_KEY_ALREADY_EXISTS": 6,
	"ERROR_CODE_PRECONDITION_NOT_MET":         7,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

// CommitType represents the type of commit.
type CommitType int32

const (
	CommitType_COMMIT_TYPE_NONE  CommitType = 0
	CommitType_COMMIT_TYPE_READ  CommitType = 1
	CommitType_COMMIT_TYPE_WRITE CommitType = 2
)

var CommitType_name = map[int32]string{
	0: "COMMIT_TYPE_NONE",
	1: "COMMIT_TYPE_READ",
	2: "COMMIT_TYPE_WRITE",
}
var CommitType_value = map[string]int32{
	"COMMIT_TYPE_NONE":  0,
	"COMMIT_TYPE_READ":  1,
	"COMMIT_TYPE_WRITE": 2,
}

func (x CommitType) String() string {
	return proto.EnumName(CommitType_name, int32(x))
}

// FileType represents a type of file from ListFiles.
type FileType int32

const (
	FileType_FILE_TYPE_NONE    FileType = 0
	FileType_FILE_TYPE_OTHER   FileType = 1
	FileType_FILE_TYPE_REGULAR FileType = 2
	FileType_FILE_TYPE_DIR     FileType = 3
)

var FileType_name = map[int32]string{
	0: "FILE_TYPE_NONE",
	1: "FILE_TYPE_OTHER",
	2: "FILE_TYPE_REGULAR",
	3: "FILE_TYPE_DIR",
}
var FileType_value = map[string]int32{
	"FILE_TYPE_NONE":    0,
	"FILE_TYPE_OTHER":   1,
	"FILE_TYPE_REGULAR": 2,
	"FILE_TYPE_DIR":     3,
}

func (x FileType) String() string {
	return proto.EnumName(FileType_name, int32(x))
}

type Error struct {
	ErrorCode ErrorCode `protobuf:"varint,1,opt,name=error_code,enum=pfs.ErrorCode" json:"error_code,omitempty"`
	Value     string    `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}

// Repository represents a repository.
type Repository struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}

// Commit represents a specific commit in a repository.
type Commit struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Id         string      `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}

func (m *Commit) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

// Path represents the full path to a file or directory within PFS.
type Path struct {
	Commit *Commit `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	Path   string  `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *Path) Reset()         { *m = Path{} }
func (m *Path) String() string { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()    {}

func (m *Path) GetCommit() *Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

// FileInfo represents information about a file.
type FileInfo struct {
	Path         *Path                       `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	FileType     FileType                    `protobuf:"varint,2,opt,name=file_type,enum=pfs.FileType" json:"file_type,omitempty"`
	SizeBytes    uint64                      `protobuf:"varint,3,opt,name=size_bytes" json:"size_bytes,omitempty"`
	Perm         uint32                      `protobuf:"varint,4,opt,name=perm" json:"perm,omitempty"`
	LastModified *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=last_modified" json:"last_modified,omitempty"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}

func (m *FileInfo) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *FileInfo) GetLastModified() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

// Shard represents a dynamic shard within PFS.
// number must always be less than modulo.
type Shard struct {
	Number uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	Modulo uint64 `protobuf:"varint,2,opt,name=modulo" json:"modulo,omitempty"`
}

func (m *Shard) Reset()         { *m = Shard{} }
func (m *Shard) String() string { return proto.CompactTextString(m) }
func (*Shard) ProtoMessage()    {}

// CommitInfo represents information about a commit.
type CommitInfo struct {
	Commit       *Commit    `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	CommitType   CommitType `protobuf:"varint,2,opt,name=commit_type,enum=pfs.CommitType" json:"commit_type,omitempty"`
	ParentCommit *Commit    `protobuf:"bytes,3,opt,name=parent_commit" json:"parent_commit,omitempty"`
}

func (m *CommitInfo) Reset()         { *m = CommitInfo{} }
func (m *CommitInfo) String() string { return proto.CompactTextString(m) }
func (*CommitInfo) ProtoMessage()    {}

func (m *CommitInfo) GetCommit() *Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *CommitInfo) GetParentCommit() *Commit {
	if m != nil {
		return m.ParentCommit
	}
	return nil
}

type Version struct {
	Major      uint32 `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor      uint32 `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
	Micro      uint32 `protobuf:"varint,3,opt,name=micro" json:"micro,omitempty"`
	Additional string `protobuf:"bytes,4,opt,name=additional" json:"additional,omitempty"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}

type GetVersionResponse struct {
	Version *Version `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
}

func (m *GetVersionResponse) Reset()         { *m = GetVersionResponse{} }
func (m *GetVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetVersionResponse) ProtoMessage()    {}

func (m *GetVersionResponse) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type InitRepositoryRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Redirect   bool        `protobuf:"varint,2,opt,name=redirect" json:"redirect,omitempty"`
}

func (m *InitRepositoryRequest) Reset()         { *m = InitRepositoryRequest{} }
func (m *InitRepositoryRequest) String() string { return proto.CompactTextString(m) }
func (*InitRepositoryRequest) ProtoMessage()    {}

func (m *InitRepositoryRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type GetFileRequest struct {
	Path        *Path `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	OffsetBytes int64 `protobuf:"varint,2,opt,name=offset_bytes" json:"offset_bytes,omitempty"`
	SizeBytes   int64 `protobuf:"varint,3,opt,name=size_bytes" json:"size_bytes,omitempty"`
}

func (m *GetFileRequest) Reset()         { *m = GetFileRequest{} }
func (m *GetFileRequest) String() string { return proto.CompactTextString(m) }
func (*GetFileRequest) ProtoMessage()    {}

func (m *GetFileRequest) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

type GetFileInfoRequest struct {
	Path *Path `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *GetFileInfoRequest) Reset()         { *m = GetFileInfoRequest{} }
func (m *GetFileInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetFileInfoRequest) ProtoMessage()    {}

func (m *GetFileInfoRequest) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

type GetFileInfoResponse struct {
	FileInfo *FileInfo `protobuf:"bytes,1,opt,name=file_info" json:"file_info,omitempty"`
}

func (m *GetFileInfoResponse) Reset()         { *m = GetFileInfoResponse{} }
func (m *GetFileInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetFileInfoResponse) ProtoMessage()    {}

func (m *GetFileInfoResponse) GetFileInfo() *FileInfo {
	if m != nil {
		return m.FileInfo
	}
	return nil
}

type MakeDirectoryRequest struct {
	Path     *Path `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Redirect bool  `protobuf:"varint,2,opt,name=redirect" json:"redirect,omitempty"`
}

func (m *MakeDirectoryRequest) Reset()         { *m = MakeDirectoryRequest{} }
func (m *MakeDirectoryRequest) String() string { return proto.CompactTextString(m) }
func (*MakeDirectoryRequest) ProtoMessage()    {}

func (m *MakeDirectoryRequest) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

type PutFileRequest struct {
	Path  *Path  `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *PutFileRequest) Reset()         { *m = PutFileRequest{} }
func (m *PutFileRequest) String() string { return proto.CompactTextString(m) }
func (*PutFileRequest) ProtoMessage()    {}

func (m *PutFileRequest) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

type ListFilesRequest struct {
	Path     *Path  `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Shard    *Shard `protobuf:"bytes,2,opt,name=shard" json:"shard,omitempty"`
	Redirect bool   `protobuf:"varint,3,opt,name=redirect" json:"redirect,omitempty"`
}

func (m *ListFilesRequest) Reset()         { *m = ListFilesRequest{} }
func (m *ListFilesRequest) String() string { return proto.CompactTextString(m) }
func (*ListFilesRequest) ProtoMessage()    {}

func (m *ListFilesRequest) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *ListFilesRequest) GetShard() *Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

type ListFilesResponse struct {
	FileInfo []*FileInfo `protobuf:"bytes,1,rep,name=file_info" json:"file_info,omitempty"`
}

func (m *ListFilesResponse) Reset()         { *m = ListFilesResponse{} }
func (m *ListFilesResponse) String() string { return proto.CompactTextString(m) }
func (*ListFilesResponse) ProtoMessage()    {}

func (m *ListFilesResponse) GetFileInfo() []*FileInfo {
	if m != nil {
		return m.FileInfo
	}
	return nil
}

type BranchRequest struct {
	Commit    *Commit `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	Redirect  bool    `protobuf:"varint,2,opt,name=redirect" json:"redirect,omitempty"`
	NewCommit *Commit `protobuf:"bytes,3,opt,name=new_commit" json:"new_commit,omitempty"`
}

func (m *BranchRequest) Reset()         { *m = BranchRequest{} }
func (m *BranchRequest) String() string { return proto.CompactTextString(m) }
func (*BranchRequest) ProtoMessage()    {}

func (m *BranchRequest) GetCommit() *Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *BranchRequest) GetNewCommit() *Commit {
	if m != nil {
		return m.NewCommit
	}
	return nil
}

type BranchResponse struct {
	Commit *Commit `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
}

func (m *BranchResponse) Reset()         { *m = BranchResponse{} }
func (m *BranchResponse) String() string { return proto.CompactTextString(m) }
func (*BranchResponse) ProtoMessage()    {}

func (m *BranchResponse) GetCommit() *Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

type CommitRequest struct {
	Commit   *Commit `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	Redirect bool    `protobuf:"varint,2,opt,name=redirect" json:"redirect,omitempty"`
}

func (m *CommitRequest) Reset()         { *m = CommitRequest{} }
func (m *CommitRequest) String() string { return proto.CompactTextString(m) }
func (*CommitRequest) ProtoMessage()    {}

func (m *CommitRequest) GetCommit() *Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

type GetCommitInfoRequest struct {
	Commit *Commit `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
}

func (m *GetCommitInfoRequest) Reset()         { *m = GetCommitInfoRequest{} }
func (m *GetCommitInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetCommitInfoRequest) ProtoMessage()    {}

func (m *GetCommitInfoRequest) GetCommit() *Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

type GetCommitInfoResponse struct {
	CommitInfo *CommitInfo `protobuf:"bytes,1,opt,name=commit_info" json:"commit_info,omitempty"`
}

func (m *GetCommitInfoResponse) Reset()         { *m = GetCommitInfoResponse{} }
func (m *GetCommitInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetCommitInfoResponse) ProtoMessage()    {}

func (m *GetCommitInfoResponse) GetCommitInfo() *CommitInfo {
	if m != nil {
		return m.CommitInfo
	}
	return nil
}

type ListCommitsRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *ListCommitsRequest) Reset()         { *m = ListCommitsRequest{} }
func (m *ListCommitsRequest) String() string { return proto.CompactTextString(m) }
func (*ListCommitsRequest) ProtoMessage()    {}

func (m *ListCommitsRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type ListCommitsResponse struct {
	CommitInfo []*CommitInfo `protobuf:"bytes,1,rep,name=commit_info" json:"commit_info,omitempty"`
}

func (m *ListCommitsResponse) Reset()         { *m = ListCommitsResponse{} }
func (m *ListCommitsResponse) String() string { return proto.CompactTextString(m) }
func (*ListCommitsResponse) ProtoMessage()    {}

func (m *ListCommitsResponse) GetCommitInfo() []*CommitInfo {
	if m != nil {
		return m.CommitInfo
	}
	return nil
}

type PullDiffRequest struct {
	Commit *Commit `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	Shard  uint64  `protobuf:"varint,2,opt,name=shard" json:"shard,omitempty"`
}

func (m *PullDiffRequest) Reset()         { *m = PullDiffRequest{} }
func (m *PullDiffRequest) String() string { return proto.CompactTextString(m) }
func (*PullDiffRequest) ProtoMessage()    {}

func (m *PullDiffRequest) GetCommit() *Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

type PushDiffRequest struct {
	Commit *Commit `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	Shard  uint64  `protobuf:"varint,2,opt,name=shard" json:"shard,omitempty"`
	Value  []byte  `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *PushDiffRequest) Reset()         { *m = PushDiffRequest{} }
func (m *PushDiffRequest) String() string { return proto.CompactTextString(m) }
func (*PushDiffRequest) ProtoMessage()    {}

func (m *PushDiffRequest) GetCommit() *Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func init() {
	proto.RegisterEnum("pfs.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("pfs.CommitType", CommitType_name, CommitType_value)
	proto.RegisterEnum("pfs.FileType", FileType_name, FileType_value)
}

// Client API for Api service

type ApiClient interface {
	GetVersion(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*GetVersionResponse, error)
	// InitRepository creates a new repository.
	// An error is returned if the specified repository already exists.
	InitRepository(ctx context.Context, in *InitRepositoryRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// GetFile returns a byte stream of the specified file.
	// An error is returned if the specified commit is a write commit.
	GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (Api_GetFileClient, error)
	// GetFileInfo returns a FileInfo for a file.
	GetFileInfo(ctx context.Context, in *GetFileInfoRequest, opts ...grpc.CallOption) (*GetFileInfoResponse, error)
	// MakeDirectory makes a directory on the file system.
	MakeDirectory(ctx context.Context, in *MakeDirectoryRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// PutFile writes the specified file to PFS.
	// An error is returned if the specified commit is not a write commit.
	PutFile(ctx context.Context, in *PutFileRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// ListFiles lists the files within a directory.
	// An error is returned if the specified path is not a directory.
	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error)
	// Branch creates a new write commit from a base commit.
	// An error is returned if the base commit is not a read commit.
	Branch(ctx context.Context, in *BranchRequest, opts ...grpc.CallOption) (*BranchResponse, error)
	// Commit turns the specified write commit into a read commit.
	// An error is returned if the specified commit is not a write commit.
	Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// GetCommitInfo returns the CommitInfo for a commit.
	GetCommitInfo(ctx context.Context, in *GetCommitInfoRequest, opts ...grpc.CallOption) (*GetCommitInfoResponse, error)
	// ListCommitInfo lists the commits on a repo
	ListCommits(ctx context.Context, in *ListCommitsRequest, opts ...grpc.CallOption) (*ListCommitsResponse, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) GetVersion(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := grpc.Invoke(ctx, "/pfs.Api/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) InitRepository(ctx context.Context, in *InitRepositoryRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pfs.Api/InitRepository", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (Api_GetFileClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Api_serviceDesc.Streams[0], c.cc, "/pfs.Api/GetFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &apiGetFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Api_GetFileClient interface {
	Recv() (*google_protobuf2.BytesValue, error)
	grpc.ClientStream
}

type apiGetFileClient struct {
	grpc.ClientStream
}

func (x *apiGetFileClient) Recv() (*google_protobuf2.BytesValue, error) {
	m := new(google_protobuf2.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *apiClient) GetFileInfo(ctx context.Context, in *GetFileInfoRequest, opts ...grpc.CallOption) (*GetFileInfoResponse, error) {
	out := new(GetFileInfoResponse)
	err := grpc.Invoke(ctx, "/pfs.Api/GetFileInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) MakeDirectory(ctx context.Context, in *MakeDirectoryRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pfs.Api/MakeDirectory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PutFile(ctx context.Context, in *PutFileRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pfs.Api/PutFile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error) {
	out := new(ListFilesResponse)
	err := grpc.Invoke(ctx, "/pfs.Api/ListFiles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Branch(ctx context.Context, in *BranchRequest, opts ...grpc.CallOption) (*BranchResponse, error) {
	out := new(BranchResponse)
	err := grpc.Invoke(ctx, "/pfs.Api/Branch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pfs.Api/Commit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetCommitInfo(ctx context.Context, in *GetCommitInfoRequest, opts ...grpc.CallOption) (*GetCommitInfoResponse, error) {
	out := new(GetCommitInfoResponse)
	err := grpc.Invoke(ctx, "/pfs.Api/GetCommitInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListCommits(ctx context.Context, in *ListCommitsRequest, opts ...grpc.CallOption) (*ListCommitsResponse, error) {
	out := new(ListCommitsResponse)
	err := grpc.Invoke(ctx, "/pfs.Api/ListCommits", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiServer interface {
	GetVersion(context.Context, *google_protobuf.Empty) (*GetVersionResponse, error)
	// InitRepository creates a new repository.
	// An error is returned if the specified repository already exists.
	InitRepository(context.Context, *InitRepositoryRequest) (*google_protobuf.Empty, error)
	// GetFile returns a byte stream of the specified file.
	// An error is returned if the specified commit is a write commit.
	GetFile(*GetFileRequest, Api_GetFileServer) error
	// GetFileInfo returns a FileInfo for a file.
	GetFileInfo(context.Context, *GetFileInfoRequest) (*GetFileInfoResponse, error)
	// MakeDirectory makes a directory on the file system.
	MakeDirectory(context.Context, *MakeDirectoryRequest) (*google_protobuf.Empty, error)
	// PutFile writes the specified file to PFS.
	// An error is returned if the specified commit is not a write commit.
	PutFile(context.Context, *PutFileRequest) (*google_protobuf.Empty, error)
	// ListFiles lists the files within a directory.
	// An error is returned if the specified path is not a directory.
	ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error)
	// Branch creates a new write commit from a base commit.
	// An error is returned if the base commit is not a read commit.
	Branch(context.Context, *BranchRequest) (*BranchResponse, error)
	// Commit turns the specified write commit into a read commit.
	// An error is returned if the specified commit is not a write commit.
	Commit(context.Context, *CommitRequest) (*google_protobuf.Empty, error)
	// GetCommitInfo returns the CommitInfo for a commit.
	GetCommitInfo(context.Context, *GetCommitInfoRequest) (*GetCommitInfoResponse, error)
	// ListCommitInfo lists the commits on a repo
	ListCommits(context.Context, *ListCommitsRequest) (*ListCommitsResponse, error)
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_GetVersion_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).GetVersion(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_InitRepository_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(InitRepositoryRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).InitRepository(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_GetFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApiServer).GetFile(m, &apiGetFileServer{stream})
}

type Api_GetFileServer interface {
	Send(*google_protobuf2.BytesValue) error
	grpc.ServerStream
}

type apiGetFileServer struct {
	grpc.ServerStream
}

func (x *apiGetFileServer) Send(m *google_protobuf2.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _Api_GetFileInfo_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(GetFileInfoRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).GetFileInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_MakeDirectory_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(MakeDirectoryRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).MakeDirectory(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_PutFile_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(PutFileRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).PutFile(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_ListFiles_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).ListFiles(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_Branch_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(BranchRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).Branch(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_Commit_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(CommitRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).Commit(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_GetCommitInfo_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(GetCommitInfoRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).GetCommitInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_ListCommits_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(ListCommitsRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).ListCommits(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pfs.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Api_GetVersion_Handler,
		},
		{
			MethodName: "InitRepository",
			Handler:    _Api_InitRepository_Handler,
		},
		{
			MethodName: "GetFileInfo",
			Handler:    _Api_GetFileInfo_Handler,
		},
		{
			MethodName: "MakeDirectory",
			Handler:    _Api_MakeDirectory_Handler,
		},
		{
			MethodName: "PutFile",
			Handler:    _Api_PutFile_Handler,
		},
		{
			MethodName: "ListFiles",
			Handler:    _Api_ListFiles_Handler,
		},
		{
			MethodName: "Branch",
			Handler:    _Api_Branch_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _Api_Commit_Handler,
		},
		{
			MethodName: "GetCommitInfo",
			Handler:    _Api_GetCommitInfo_Handler,
		},
		{
			MethodName: "ListCommits",
			Handler:    _Api_ListCommits_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFile",
			Handler:       _Api_GetFile_Handler,
			ServerStreams: true,
		},
	},
}

// Client API for InternalApi service

type InternalApiClient interface {
	// PullDiff pulls a binary stream of the diff from the specified
	// commit to the commit's parent.
	PullDiff(ctx context.Context, in *PullDiffRequest, opts ...grpc.CallOption) (InternalApi_PullDiffClient, error)
	// Push diff pushes a diff from the specified commit.
	PushDiff(ctx context.Context, in *PushDiffRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type internalApiClient struct {
	cc *grpc.ClientConn
}

func NewInternalApiClient(cc *grpc.ClientConn) InternalApiClient {
	return &internalApiClient{cc}
}

func (c *internalApiClient) PullDiff(ctx context.Context, in *PullDiffRequest, opts ...grpc.CallOption) (InternalApi_PullDiffClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_InternalApi_serviceDesc.Streams[0], c.cc, "/pfs.InternalApi/PullDiff", opts...)
	if err != nil {
		return nil, err
	}
	x := &internalApiPullDiffClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InternalApi_PullDiffClient interface {
	Recv() (*google_protobuf2.BytesValue, error)
	grpc.ClientStream
}

type internalApiPullDiffClient struct {
	grpc.ClientStream
}

func (x *internalApiPullDiffClient) Recv() (*google_protobuf2.BytesValue, error) {
	m := new(google_protobuf2.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *internalApiClient) PushDiff(ctx context.Context, in *PushDiffRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pfs.InternalApi/PushDiff", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InternalApi service

type InternalApiServer interface {
	// PullDiff pulls a binary stream of the diff from the specified
	// commit to the commit's parent.
	PullDiff(*PullDiffRequest, InternalApi_PullDiffServer) error
	// Push diff pushes a diff from the specified commit.
	PushDiff(context.Context, *PushDiffRequest) (*google_protobuf.Empty, error)
}

func RegisterInternalApiServer(s *grpc.Server, srv InternalApiServer) {
	s.RegisterService(&_InternalApi_serviceDesc, srv)
}

func _InternalApi_PullDiff_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullDiffRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InternalApiServer).PullDiff(m, &internalApiPullDiffServer{stream})
}

type InternalApi_PullDiffServer interface {
	Send(*google_protobuf2.BytesValue) error
	grpc.ServerStream
}

type internalApiPullDiffServer struct {
	grpc.ServerStream
}

func (x *internalApiPullDiffServer) Send(m *google_protobuf2.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _InternalApi_PushDiff_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(PushDiffRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(InternalApiServer).PushDiff(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _InternalApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pfs.InternalApi",
	HandlerType: (*InternalApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushDiff",
			Handler:    _InternalApi_PushDiff_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullDiff",
			Handler:       _InternalApi_PullDiff_Handler,
			ServerStreams: true,
		},
	},
}
