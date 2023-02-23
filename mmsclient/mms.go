// Code generated by goctl. DO NOT EDIT.
// Source: mms.proto

package mmsclient

import (
	"context"

	"github.com/suyuan32/simple-admin-member-rpc/mms"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseIDResp         = mms.BaseIDResp
	BaseResp           = mms.BaseResp
	BaseUUIDResp       = mms.BaseUUIDResp
	Empty              = mms.Empty
	IDReq              = mms.IDReq
	IDsReq             = mms.IDsReq
	MemberInfo         = mms.MemberInfo
	MemberListReq      = mms.MemberListReq
	MemberListResp     = mms.MemberListResp
	MemberLoginResp    = mms.MemberLoginResp
	MemberRankInfo     = mms.MemberRankInfo
	MemberRankListReq  = mms.MemberRankListReq
	MemberRankListResp = mms.MemberRankListResp
	MemberRegisterReq  = mms.MemberRegisterReq
	PageInfoReq        = mms.PageInfoReq
	UUIDReq            = mms.UUIDReq
	UUIDsReq           = mms.UUIDsReq
	UsernameReq        = mms.UsernameReq

	Mms interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// Member management
		CreateMember(ctx context.Context, in *MemberInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		UpdateMember(ctx context.Context, in *MemberInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetMemberList(ctx context.Context, in *MemberListReq, opts ...grpc.CallOption) (*MemberListResp, error)
		DeleteMember(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetMemberById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*MemberInfo, error)
		GetMemberByUsername(ctx context.Context, in *UsernameReq, opts ...grpc.CallOption) (*MemberInfo, error)
		// MemberRank management
		CreateMemberRank(ctx context.Context, in *MemberRankInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateMemberRank(ctx context.Context, in *MemberRankInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetMemberRankList(ctx context.Context, in *MemberRankListReq, opts ...grpc.CallOption) (*MemberRankListResp, error)
		GetMemberRankById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MemberRankInfo, error)
		DeleteMemberRank(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultMms struct {
		cli zrpc.Client
	}
)

func NewMms(cli zrpc.Client) Mms {
	return &defaultMms{
		cli: cli,
	}
}

func (m *defaultMms) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

// Member management
func (m *defaultMms) CreateMember(ctx context.Context, in *MemberInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.CreateMember(ctx, in, opts...)
}

func (m *defaultMms) UpdateMember(ctx context.Context, in *MemberInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.UpdateMember(ctx, in, opts...)
}

func (m *defaultMms) GetMemberList(ctx context.Context, in *MemberListReq, opts ...grpc.CallOption) (*MemberListResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetMemberList(ctx, in, opts...)
}

func (m *defaultMms) DeleteMember(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.DeleteMember(ctx, in, opts...)
}

func (m *defaultMms) GetMemberById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*MemberInfo, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetMemberById(ctx, in, opts...)
}

func (m *defaultMms) GetMemberByUsername(ctx context.Context, in *UsernameReq, opts ...grpc.CallOption) (*MemberInfo, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetMemberByUsername(ctx, in, opts...)
}

// MemberRank management
func (m *defaultMms) CreateMemberRank(ctx context.Context, in *MemberRankInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.CreateMemberRank(ctx, in, opts...)
}

func (m *defaultMms) UpdateMemberRank(ctx context.Context, in *MemberRankInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.UpdateMemberRank(ctx, in, opts...)
}

func (m *defaultMms) GetMemberRankList(ctx context.Context, in *MemberRankListReq, opts ...grpc.CallOption) (*MemberRankListResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetMemberRankList(ctx, in, opts...)
}

func (m *defaultMms) GetMemberRankById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MemberRankInfo, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetMemberRankById(ctx, in, opts...)
}

func (m *defaultMms) DeleteMemberRank(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.DeleteMemberRank(ctx, in, opts...)
}