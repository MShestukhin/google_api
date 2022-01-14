// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: services/google_api/pb/google_api/google_api.proto

package google_api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for GoogleApi service

func NewGoogleApiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GoogleApi service

type GoogleApiService interface {
	ExchangeCode(ctx context.Context, in *ExchangeCodeRequest, opts ...client.CallOption) (*ExchangeCodeResponse, error)
	GetOAuthUrl(ctx context.Context, in *GetOAuthUrlRequest, opts ...client.CallOption) (*GetOAuthUrlResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTockenRequest, opts ...client.CallOption) (*RefreshTockenResponse, error)
	GetCustomers(ctx context.Context, in *GetCustomersRequest, opts ...client.CallOption) (*GetCustomersResponse, error)
	GetCampaigns(ctx context.Context, in *GetCampaignsRequest, opts ...client.CallOption) (*GetCampaignsResponse, error)
	GetAdGroups(ctx context.Context, in *GetAdGroupsRequest, opts ...client.CallOption) (*GetAdGroupsResponse, error)
	GetAdGroupsAd(ctx context.Context, in *GetAdGroupsAdRequest, opts ...client.CallOption) (*GetAdGroupsAdResponse, error)
	GetCampaignCriteria(ctx context.Context, in *GetCampaignCriteriaRequest, opts ...client.CallOption) (*GetCampaignCriteriaResponse, error)
	GetAdGroupsCriteria(ctx context.Context, in *GetAdGroupsCriteriaRequest, opts ...client.CallOption) (*GetAdGroupsCriteriaResponse, error)
}

type googleApiService struct {
	c    client.Client
	name string
}

func NewGoogleApiService(name string, c client.Client) GoogleApiService {
	return &googleApiService{
		c:    c,
		name: name,
	}
}

func (c *googleApiService) ExchangeCode(ctx context.Context, in *ExchangeCodeRequest, opts ...client.CallOption) (*ExchangeCodeResponse, error) {
	req := c.c.NewRequest(c.name, "GoogleApi.ExchangeCode", in)
	out := new(ExchangeCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleApiService) GetOAuthUrl(ctx context.Context, in *GetOAuthUrlRequest, opts ...client.CallOption) (*GetOAuthUrlResponse, error) {
	req := c.c.NewRequest(c.name, "GoogleApi.GetOAuthUrl", in)
	out := new(GetOAuthUrlResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleApiService) RefreshToken(ctx context.Context, in *RefreshTockenRequest, opts ...client.CallOption) (*RefreshTockenResponse, error) {
	req := c.c.NewRequest(c.name, "GoogleApi.RefreshToken", in)
	out := new(RefreshTockenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleApiService) GetCustomers(ctx context.Context, in *GetCustomersRequest, opts ...client.CallOption) (*GetCustomersResponse, error) {
	req := c.c.NewRequest(c.name, "GoogleApi.GetCustomers", in)
	out := new(GetCustomersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleApiService) GetCampaigns(ctx context.Context, in *GetCampaignsRequest, opts ...client.CallOption) (*GetCampaignsResponse, error) {
	req := c.c.NewRequest(c.name, "GoogleApi.GetCampaigns", in)
	out := new(GetCampaignsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleApiService) GetAdGroups(ctx context.Context, in *GetAdGroupsRequest, opts ...client.CallOption) (*GetAdGroupsResponse, error) {
	req := c.c.NewRequest(c.name, "GoogleApi.GetAdGroups", in)
	out := new(GetAdGroupsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleApiService) GetAdGroupsAd(ctx context.Context, in *GetAdGroupsAdRequest, opts ...client.CallOption) (*GetAdGroupsAdResponse, error) {
	req := c.c.NewRequest(c.name, "GoogleApi.GetAdGroupsAd", in)
	out := new(GetAdGroupsAdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleApiService) GetCampaignCriteria(ctx context.Context, in *GetCampaignCriteriaRequest, opts ...client.CallOption) (*GetCampaignCriteriaResponse, error) {
	req := c.c.NewRequest(c.name, "GoogleApi.GetCampaignCriteria", in)
	out := new(GetCampaignCriteriaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleApiService) GetAdGroupsCriteria(ctx context.Context, in *GetAdGroupsCriteriaRequest, opts ...client.CallOption) (*GetAdGroupsCriteriaResponse, error) {
	req := c.c.NewRequest(c.name, "GoogleApi.GetAdGroupsCriteria", in)
	out := new(GetAdGroupsCriteriaResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoogleApi service

type GoogleApiHandler interface {
	ExchangeCode(context.Context, *ExchangeCodeRequest, *ExchangeCodeResponse) error
	GetOAuthUrl(context.Context, *GetOAuthUrlRequest, *GetOAuthUrlResponse) error
	RefreshToken(context.Context, *RefreshTockenRequest, *RefreshTockenResponse) error
	GetCustomers(context.Context, *GetCustomersRequest, *GetCustomersResponse) error
	GetCampaigns(context.Context, *GetCampaignsRequest, *GetCampaignsResponse) error
	GetAdGroups(context.Context, *GetAdGroupsRequest, *GetAdGroupsResponse) error
	GetAdGroupsAd(context.Context, *GetAdGroupsAdRequest, *GetAdGroupsAdResponse) error
	GetCampaignCriteria(context.Context, *GetCampaignCriteriaRequest, *GetCampaignCriteriaResponse) error
	GetAdGroupsCriteria(context.Context, *GetAdGroupsCriteriaRequest, *GetAdGroupsCriteriaResponse) error
}

func RegisterGoogleApiHandler(s server.Server, hdlr GoogleApiHandler, opts ...server.HandlerOption) error {
	type googleApi interface {
		ExchangeCode(ctx context.Context, in *ExchangeCodeRequest, out *ExchangeCodeResponse) error
		GetOAuthUrl(ctx context.Context, in *GetOAuthUrlRequest, out *GetOAuthUrlResponse) error
		RefreshToken(ctx context.Context, in *RefreshTockenRequest, out *RefreshTockenResponse) error
		GetCustomers(ctx context.Context, in *GetCustomersRequest, out *GetCustomersResponse) error
		GetCampaigns(ctx context.Context, in *GetCampaignsRequest, out *GetCampaignsResponse) error
		GetAdGroups(ctx context.Context, in *GetAdGroupsRequest, out *GetAdGroupsResponse) error
		GetAdGroupsAd(ctx context.Context, in *GetAdGroupsAdRequest, out *GetAdGroupsAdResponse) error
		GetCampaignCriteria(ctx context.Context, in *GetCampaignCriteriaRequest, out *GetCampaignCriteriaResponse) error
		GetAdGroupsCriteria(ctx context.Context, in *GetAdGroupsCriteriaRequest, out *GetAdGroupsCriteriaResponse) error
	}
	type GoogleApi struct {
		googleApi
	}
	h := &googleApiHandler{hdlr}
	return s.Handle(s.NewHandler(&GoogleApi{h}, opts...))
}

type googleApiHandler struct {
	GoogleApiHandler
}

func (h *googleApiHandler) ExchangeCode(ctx context.Context, in *ExchangeCodeRequest, out *ExchangeCodeResponse) error {
	return h.GoogleApiHandler.ExchangeCode(ctx, in, out)
}

func (h *googleApiHandler) GetOAuthUrl(ctx context.Context, in *GetOAuthUrlRequest, out *GetOAuthUrlResponse) error {
	return h.GoogleApiHandler.GetOAuthUrl(ctx, in, out)
}

func (h *googleApiHandler) RefreshToken(ctx context.Context, in *RefreshTockenRequest, out *RefreshTockenResponse) error {
	return h.GoogleApiHandler.RefreshToken(ctx, in, out)
}

func (h *googleApiHandler) GetCustomers(ctx context.Context, in *GetCustomersRequest, out *GetCustomersResponse) error {
	return h.GoogleApiHandler.GetCustomers(ctx, in, out)
}

func (h *googleApiHandler) GetCampaigns(ctx context.Context, in *GetCampaignsRequest, out *GetCampaignsResponse) error {
	return h.GoogleApiHandler.GetCampaigns(ctx, in, out)
}

func (h *googleApiHandler) GetAdGroups(ctx context.Context, in *GetAdGroupsRequest, out *GetAdGroupsResponse) error {
	return h.GoogleApiHandler.GetAdGroups(ctx, in, out)
}

func (h *googleApiHandler) GetAdGroupsAd(ctx context.Context, in *GetAdGroupsAdRequest, out *GetAdGroupsAdResponse) error {
	return h.GoogleApiHandler.GetAdGroupsAd(ctx, in, out)
}

func (h *googleApiHandler) GetCampaignCriteria(ctx context.Context, in *GetCampaignCriteriaRequest, out *GetCampaignCriteriaResponse) error {
	return h.GoogleApiHandler.GetCampaignCriteria(ctx, in, out)
}

func (h *googleApiHandler) GetAdGroupsCriteria(ctx context.Context, in *GetAdGroupsCriteriaRequest, out *GetAdGroupsCriteriaResponse) error {
	return h.GoogleApiHandler.GetAdGroupsCriteria(ctx, in, out)
}
