/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"

	"github.com/Azure/azure-service-operator/v2/api"
)

func Test_ARMResourceImporter_GroupKindFromARMID(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		armId         string
		expectedGroup string
		expectedKind  string
	}{
		{
			name:          "resource group",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg",
			expectedGroup: "resources.azure.com",
			expectedKind:  "ResourceGroup",
		},
		{
			name:          "storage account",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Storage/storageAccounts/asostorage",
			expectedGroup: "storage.azure.com",
			expectedKind:  "StorageAccount",
		},
		{
			name:          "cosmos account",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.DocumentDB/databaseAccounts/aso-cosmos",
			expectedGroup: "documentdb.azure.com",
			expectedKind:  "DatabaseAccount",
		},
		{
			name:          "VMSS Scale Set",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Compute/virtualMachineScaleSets/aso-scaleset",
			expectedGroup: "compute.azure.com",
			expectedKind:  "VirtualMachineScaleSet",
		},
		{
			name:          "VMSS Scale Set",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.ContainerService/managedClusters/aso-cluster",
			expectedGroup: "containerservice.azure.com",
			expectedKind:  "ManagedCluster",
		},
		{
			name:          "Redis cache",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Cache/redis/my-cache",
			expectedGroup: "cache.azure.com",
			expectedKind:  "Redis",
		},
		{
			name:          "Redis cache different case",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Cache/REDIS/my-cache",
			expectedGroup: "cache.azure.com",
			expectedKind:  "Redis",
		},
		//{
		//	"VMSS Instance",
		//	"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Compute/virtualMachineScaleSets/aso-scaleset/virtualMachines/0",
		//	"storage",
		//	"storageaccount",
		//},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			c := c
			t.Parallel()

			g := NewGomegaWithT(t)

			rsrc := importableARMResource{}

			id, err := arm.ParseResourceID(c.armId)
			g.Expect(err).To(BeNil())

			gk, err := rsrc.groupKindFromID(id)
			g.Expect(err).To(BeNil())
			g.Expect(gk.Group).To(Equal(c.expectedGroup))
			g.Expect(gk.Kind).To(Equal(c.expectedKind))
		})
	}
}

func Test_ARMResourceImporter_GroupVersionKindFromARMID(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name            string
		armId           string
		expectedGroup   string
		expectedKind    string
		expectedVersion string
	}{
		{
			name:            "resource group",
			armId:           "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg",
			expectedGroup:   "resources.azure.com",
			expectedKind:    "ResourceGroup",
			expectedVersion: "v1api20200601",
		},
		{
			name:            "storage account",
			armId:           "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Storage/storageAccounts/aso-storage",
			expectedGroup:   "storage.azure.com",
			expectedKind:    "StorageAccount",
			expectedVersion: "v1api20230101",
		},
		{
			name:            "managed cluster",
			armId:           "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.ContainerService/managedClusters/aso-cluster",
			expectedGroup:   "containerservice.azure.com",
			expectedKind:    "ManagedCluster",
			expectedVersion: "v1api20240901",
		},
		{
			name:            "Redis cache",
			armId:           "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Cache/redis/my-cache",
			expectedGroup:   "cache.azure.com",
			expectedKind:    "Redis",
			expectedVersion: "v1api20230801",
		},
		{
			name:            "Redis cache different case",
			armId:           "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Cache/REDIS/my-cache",
			expectedGroup:   "cache.azure.com",
			expectedKind:    "Redis",
			expectedVersion: "v1api20230801",
		},
	}

	factory := &importFactory{
		scheme: api.CreateScheme(),
	}

	armRsrc := importableARMResource{}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)

			id, err := arm.ParseResourceID(c.armId)
			g.Expect(err).To(BeNil())

			gvk, err := armRsrc.groupVersionKindFromID(id, factory)
			g.Expect(err).To(BeNil())

			// If the asserts fail, check to see whether we've introduced a new version of the resource
			// specified by the ARM ID. If so, update the expected version in the test case.
			g.Expect(gvk.Group).To(Equal(c.expectedGroup))
			g.Expect(gvk.Kind).To(Equal(c.expectedKind))
			g.Expect(gvk.Version).To(Equal(c.expectedVersion))
		})
	}
}
