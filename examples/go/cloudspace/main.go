package main

import (
	"github.com/pulumi/fabianv-cloud/sdk/go/vco/base"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cloudspace, err := base.NewCloudspace(ctx, "pulumi-cloudspace", base.Cloudspaceargs{})
		if err != nil {
			return err
		}

		return nil
	})
}
