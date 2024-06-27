package main

import (
	// "fmt"
	"baltic-sea/goprojects"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createNamespace(ctx *pulumi.Context, nameSpaceName string) error {
	_, err := yaml.NewConfigFile(ctx, nameSpaceName, &yaml.ConfigFileArgs{
		File: "yaml/namespace/pulumi-apps.yaml",
	})
	return err
}

var balticSea = func(ctx *pulumi.Context) error {
	nameSpaceName := "pulumi-apps"

	err := createNamespace(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = goprojects.CreateDebugContainer(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = goprojects.CreateTifTest(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = goprojects.CreateTheIndependentFriendDe(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = goprojects.CreateTestOlafRadickeDE(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = goprojects.CreateOlafRadickeDE(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = goprojects.CreateTestQuakerKrDE(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = goprojects.CreateIngress(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	// err = goprojects.Hugo(ctx, nameSpaceName)
	// if err != nil {
	// 	return err
	// }

	// err = createPrometheusIngress(ctx, nameSpaceName)
	// if err != nil {
	// 	return err
	// }

	// err = createPrometheus(ctx, nameSpaceName)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func main() {
	pulumi.Run(balticSea)
	// pulumi.Run(mySetup)
}
