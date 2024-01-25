package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func azureSetup(contex *pulumi.Context) error {
	// Create an Azure Resource Group
	resourceGroup, err := resources.NewResourceGroup(contex, "pulumi", nil)
	if err != nil {
		return err
	}
	var skuArgs = storage.SkuArgs{
		Name: pulumi.String("Standard_LRS"),
	}
	var storageAccountArgs = storage.StorageAccountArgs{
		ResourceGroupName: resourceGroup.Name,
		Sku:               &skuArgs,
		Kind:              pulumi.String("StorageV2"),
	}
	// Create an Azure resource (Storage Account)
	account, err := storage.NewStorageAccount(contex, "sa001", &storageAccountArgs)
	if err != nil {
		return err
	}

	// Export the primary key of the Storage Account
	contex.Export("primaryStorageKey", pulumi.All(resourceGroup.Name, account.Name).ApplyT(
		func(args []interface{}) (string, error) {
			resourceGroupName := args[0].(string)
			accountName := args[1].(string)
			accountKeys, err := storage.ListStorageAccountKeys(contex, &storage.ListStorageAccountKeysArgs{
				ResourceGroupName: resourceGroupName,
				AccountName:       accountName,
			})
			if err != nil {
				return "", err
			}

			return accountKeys.Keys[0].Value, nil
		},
	))

	return nil
}

func main() {
	pulumi.Run(azureSetup)
}
