package gocode

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func UptimeKuma(ctx *pulumi.Context, nameSpaceName string) error {

	_, err := yaml.NewConfigFile(ctx, "uptime-kuma", &yaml.ConfigFileArgs{
		File: "yaml/uptime-kuma/*.yaml",
	})
	if err != nil {
		return err
	}
	return nil

}
