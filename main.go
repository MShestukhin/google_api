package main

import (
	"github.com/urfave/cli/v2"
	"gitlab.somin.ai/analytics/platform/pkg/app"
	"gitlab.somin.ai/analytics/platform/pkg/app/cmd"
	"gitlab.somin.ai/analytics/platform/services/google_api/internal/service"
	"gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
	"os"
)

func main() {
	err := cmd.Run(
		app.Namespace("ai.somin"),
		app.ServiceName(google_api.ServiceName),
		app.Version("latest"),
		app.Tracer(true),
		app.Action(func(_ *cli.Context, opts app.Options) error {
			a := app.New(opts)

			a.AddService(service.New(a))
			if err := a.Init(); err != nil {
				return err
			}
			defer a.Exit()

			return a.Run()
		}),
	)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
