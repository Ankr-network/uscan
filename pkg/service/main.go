package service

import (
	"context"

	"github.com/Ankr-network/uscan/pkg/apis"
	"github.com/spf13/cobra"
	"github.com/sunvim/utils/grace"
)

func MainRun(cmd *cobra.Command, args []string) {
	_, svc := grace.New(context.Background())

	svc.RegisterService("web service", apis.Apis)

	svc.Wait()
}
