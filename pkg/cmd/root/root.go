package root

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm/ssmiface"
	listCmd "github.com/benkim0414/env/pkg/cmd/list"
	"github.com/spf13/cobra"
)

func NewRootCmd(svc func() (ssmiface.ClientAPI, error)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "env <command> [flags]",
		Short: "env is a tool to manage environment variables with AWS SSM",
	}

	cmd.AddCommand(listCmd.NewListCmd(svc))

	return cmd
}
