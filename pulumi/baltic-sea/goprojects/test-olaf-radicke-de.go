package goprojects

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateTestOlafRadickeDE(ctx *pulumi.Context, nameSpaceName string) error {

	yamlConfig := &yaml.ConfigFileArgs{
		File: "yaml/test-olaf-radicke-de/*.yaml",
	}
	_, err := yaml.NewConfigFile(ctx, "test-olaf-radicke-de-deployment", yamlConfig)
	if err != nil {
		return err
	}

	return nil
}
