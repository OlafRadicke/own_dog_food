package main

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createPrometheus(ctx *pulumi.Context, nameSpaceName string) error {
	_, err := helm.NewChart(ctx, "prometheus", helm.ChartArgs{
		Namespace: pulumi.String(nameSpaceName),
		Chart:     pulumi.String("prometheus"),
		Version:   pulumi.String("15.9.1"),
		FetchArgs: helm.FetchArgs{
			Repo: pulumi.String("https://prometheus-community.github.io/helm-charts"),
		},
		Values: pulumi.Map{
			"rbac": pulumi.Map{
				"enabled": pulumi.Bool(false),
			},
			"nodeExporter": pulumi.Map{
				"create": pulumi.Bool(false),
				"service": pulumi.Map{
					"hostPort": pulumi.Int(9110),
				},
			},
		},
	})
	return err
}

func createPrometheusIngress(ctx *pulumi.Context, nameSpaceName string) error {
	_, err := helm.NewChart(ctx, "prometheus-ingress", helm.ChartArgs{
		Namespace: pulumi.String(nameSpaceName),
		Chart:     pulumi.String("tls_routing"),
		Version:   pulumi.String("0.3.0"),
		FetchArgs: helm.FetchArgs{
			Repo: pulumi.String("https://olafradicke.github.io/own_dog_food/helm/charts/"),
		},
		Values: pulumi.Map{
			"mtls": pulumi.Map{
				"enabled":     pulumi.Bool(false),
				"secretnames": pulumi.String("mtls-ca"),
			},
			"app": pulumi.Map{
				"name":         pulumi.String("prometheus"),
				"fqdn":         pulumi.String("prometheus.olaf-radicke.de"),
				"service":      pulumi.String("prometheus-server"),
				"service_port": pulumi.Int(80),
			},
		},
	})
	return err
}
