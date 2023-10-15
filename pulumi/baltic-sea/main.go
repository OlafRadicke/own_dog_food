package main

// appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
// corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
// metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var nameOfNamespace = "my-namespace"

var myNamespace = func(ctx *pulumi.Context) error {
	_, err := yaml.NewConfigFile(ctx, nameOfNamespace, &yaml.ConfigFileArgs{
		File: "namespace.yaml",
	})
	if err != nil {
		return err
	}

	_, err = helm.NewChart(ctx, "nginx-ingress", helm.ChartArgs{
		Namespace: pulumi.String(nameOfNamespace),
		Chart:     pulumi.String("nginx-ingress"),
		Version:   pulumi.String("1.24.4"),
		FetchArgs: helm.FetchArgs{
			Repo: pulumi.String("https://charts.helm.sh/stable"),
		},
		Values: pulumi.Map{
			"controller": pulumi.Map{
				"metrics": pulumi.Map{
					"enabled": pulumi.Bool(true),
				},
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	pulumi.Run(myNamespace)
	// pulumi.Run(mySetup)
}
