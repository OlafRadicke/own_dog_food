package gocode

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LocalStorage(ctx *pulumi.Context, nameSpaceName string) error {

	_, err := yaml.NewConfigFile(ctx, "local-storage", &yaml.ConfigFileArgs{
		File: "yaml/local-storage/*.yaml",
	})
	if err != nil {
		return err
	}
	return nil

}
