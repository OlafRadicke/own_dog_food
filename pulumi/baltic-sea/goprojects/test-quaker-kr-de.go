package goprojects

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateTestQuakerKrDE(ctx *pulumi.Context, nameSpaceName string) error {

	yamlConfig := &yaml.ConfigFileArgs{
		File: "yaml/test-quaker-kr-de/*.yaml",
	}
	_, err := yaml.NewConfigFile(ctx, "test-quaker-kr-de-deployment", yamlConfig)
	if err != nil {
		return err
	}

	return nil
}
