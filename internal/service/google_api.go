package service

import (
	"github.com/micro/go-micro/v2/server"
	"gitlab.somin.ai/analytics/platform/pkg/app"
	"gitlab.somin.ai/analytics/platform/pkg/app/core"
	"gitlab.somin.ai/analytics/platform/services/google_api/internal/google"
	"gitlab.somin.ai/analytics/platform/services/google_api/internal/handlers"
	pb "gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
)

var (
	_ core.Service    = (*svc)(nil)
	_ core.RpcService = (*svc)(nil)
)

func New(a app.App) *svc {
	return &svc{a: a}
}

type svc struct {
	a app.App
}

func (s *svc) Name() string {
	return pb.ServiceName
}

func (s *svc) RegisterRpcService(srv server.Server) error {
	serviceCfg := google.GoogleAdsClientParams{}

	if err := s.a.Micro().Options().Config.Get("google_api").Scan(&serviceCfg); err != nil {
		panic(err)
	}

	handler := handlers.NewGoogleApiHandler(serviceCfg)
	return pb.RegisterGoogleApiHandler(srv, handler)
}
