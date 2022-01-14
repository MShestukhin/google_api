module gitlab.somin.ai/analytics/platform/services/google_api

go 1.16

replace github.com/micro/go-micro/v2 => gitlab.somin.ai/analytics/go-micro/v2 v2.9.2-0.20210714093304-414789c34a58

require (
	cloud.google.com/go v0.83.0 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/opteo/google-ads-go v0.2.2-0.20210712133232-20de8242e64f // indirect
	github.com/urfave/cli/v2 v2.3.0
	gitlab.somin.ai/analytics/platform/pkg/app v0.0.0-20210823135053-92a4fc580a85
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c
	golang.org/x/sys v0.0.0-20210603125802-9665404d3644 // indirect
	google.golang.org/genproto v0.0.0-20210604141403-392c879c8b08 // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)
