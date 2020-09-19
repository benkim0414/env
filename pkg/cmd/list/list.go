package list

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/ssmiface"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	SSMClient func() (ssmiface.ClientAPI, error)

	Path           string
	MaxResults     int64
	Recursive      bool
	WithDecryption bool
}

func NewListCmd(svc func() (ssmiface.ClientAPI, error)) *cobra.Command {
	opts := &ListOptions{
		SSMClient: svc,
	}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List parameters in a specific hierarchy",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Path, "path", "p", "/", "The hierarchy for the parameter")
	cmd.Flags().Int64VarP(&opts.MaxResults, "max", "m", 10, "The maximum number of items to list")
	cmd.Flags().BoolVarP(&opts.Recursive, "recursive", "r", true, "Retrieve all parameters within a hierarchy")
	cmd.Flags().BoolVarP(&opts.WithDecryption, "decrypt", "d", true, "Retrieve all parameters in a hierarchy with their value decrypted")

	return cmd
}

func run(opts *ListOptions) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)

	svc, err := opts.SSMClient()
	if err != nil {
		return err
	}

	req := svc.GetParametersByPathRequest(&ssm.GetParametersByPathInput{
		Path:           aws.String(opts.Path),
		MaxResults:     aws.Int64(opts.MaxResults),
		Recursive:      aws.Bool(opts.Recursive),
		WithDecryption: aws.Bool(opts.WithDecryption),
	})

	p := ssm.NewGetParametersByPathPaginator(req)
	for p.Next(context.TODO()) {
		page := p.CurrentPage()
		for _, param := range page.Parameters {
			fmt.Fprintf(w, "%s\t%s\n", aws.StringValue(param.Name), aws.StringValue(param.Value))
		}
	}
	w.Flush()

	if err := p.Err(); err != nil {
		return err
	}
	return nil
}
