package main

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/ssmiface"
	"github.com/benkim0414/env/pkg/cmd/root"
)

func main() {
	rootCmd := root.NewRootCmd(func() (ssmiface.ClientAPI, error) {
		cfg, err := external.LoadDefaultAWSConfig()
		if err != nil {
			return nil, err
		}
		svc := ssm.New(cfg)
		return svc, nil
	})
	if _, err := rootCmd.ExecuteC(); err != nil {
		log.Fatal(err)
	}
}
