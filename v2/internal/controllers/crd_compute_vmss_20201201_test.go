/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	compute2020 "github.com/Azure/azure-service-operator/v2/api/compute/v1api20201201"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	network20240301 "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func newVMVirtualNetwork(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.VirtualNetwork {
	return &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetwork_Spec{
			Owner:    owner,
			Location: tc.AzureRegion,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/16"},
			},
		},
	}
}

func newVMVirtualNetwork20240301(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network20240301.VirtualNetwork {
	return &network20240301.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network20240301.VirtualNetwork_Spec{
			Owner:    owner,
			Location: tc.AzureRegion,
			AddressSpace: &network20240301.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/16"},
			},
		},
	}
}

func newVMSubnet(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.VirtualNetworksSubnet {
	return &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network.VirtualNetworksSubnet_Spec{
			Owner:         owner,
			AddressPrefix: to.Ptr("10.0.0.0/24"),
		},
	}
}

func newVMSubnet20240301(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network20240301.VirtualNetworksSubnet {
	return &network20240301.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network20240301.VirtualNetworksSubnet_Spec{
			Owner:         owner,
			AddressPrefix: to.Ptr("10.0.0.0/24"),
		},
	}
}

func newPublicIPAddressForVMSS(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.PublicIPAddress {
	publicIPAddressSku := network.PublicIPAddressSku_Name_Standard
	allocationMethod := network.IPAllocationMethod_Static
	return &network.PublicIPAddress{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddress_Spec{
			Location: tc.AzureRegion,
			Owner:    owner,
			Sku: &network.PublicIPAddressSku{
				Name: &publicIPAddressSku,
			},
			PublicIPAllocationMethod: &allocationMethod,
		},
	}
}

func newLoadBalancerForVMSS(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, publicIPAddress *network.PublicIPAddress) *network.LoadBalancer {
	loadBalancerSku := network.LoadBalancerSku_Name_Standard
	lbName := tc.Namer.GenerateName("loadbalancer")
	lbFrontendName := "LoadBalancerFrontend"
	protocol := network.TransportProtocol_Tcp

	// TODO: Getting this is SUPER awkward
	frontIPConfigurationARMID, err := genericarmclient.MakeResourceGroupScopeARMID(
		tc.AzureSubscription,
		rg.Name,
		"Microsoft.Network",
		"loadBalancers",
		lbName,
		"frontendIPConfigurations",
		lbFrontendName)
	if err != nil {
		panic(err)
	}

	return &network.LoadBalancer{
		ObjectMeta: tc.MakeObjectMetaWithName(lbName),
		Spec: network.LoadBalancer_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.LoadBalancerSku{
				Name: &loadBalancerSku,
			},
			FrontendIPConfigurations: []network.FrontendIPConfiguration_LoadBalancer_SubResourceEmbedded{
				{
					Name: &lbFrontendName,
					PublicIPAddress: &network.PublicIPAddressSpec_LoadBalancer_SubResourceEmbedded{
						Reference: tc.MakeReferenceFromResource(publicIPAddress),
					},
				},
			},
			InboundNatPools: []network.InboundNatPool{
				{
					Name: to.Ptr("MyFancyNatPool"),
					FrontendIPConfiguration: &network.SubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: frontIPConfigurationARMID,
						},
					},
					Protocol:               &protocol,
					FrontendPortRangeStart: to.Ptr(50_000),
					FrontendPortRangeEnd:   to.Ptr(51_000),
					BackendPort:            to.Ptr(22),
				},
			},
		},
	}
}

func newVMSS20201201(
	tc *testcommon.KubePerTestContext,
	rg *resources.ResourceGroup,
	loadBalancer *network.LoadBalancer,
	subnet *network.VirtualNetworksSubnet,
) *compute2020.VirtualMachineScaleSet {
	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

	upgradePolicyMode := compute2020.UpgradePolicy_Mode_Automatic
	adminUsername := "adminUser"

	inboundNATPoolRef := genruntime.ResourceReference{
		// TODO: It is the most awkward thing in the world that this is not a fully fledged resource
		ARMID: *loadBalancer.Status.InboundNatPools[0].Id,
	}

	return &compute2020.VirtualMachineScaleSet{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vmss")),
		Spec: compute2020.VirtualMachineScaleSet_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &compute2020.Sku{
				Name:     to.Ptr("STANDARD_D1_v2"),
				Capacity: to.Ptr(1),
			},
			PlatformFaultDomainCount: to.Ptr(3),
			SinglePlacementGroup:     to.Ptr(false),
			UpgradePolicy: &compute2020.UpgradePolicy{
				Mode: &upgradePolicyMode,
			},
			VirtualMachineProfile: &compute2020.VirtualMachineScaleSetVMProfile{
				StorageProfile: &compute2020.VirtualMachineScaleSetStorageProfile{
					ImageReference: &compute2020.ImageReference{
						Publisher: to.Ptr("Canonical"),
						Offer:     to.Ptr("UbuntuServer"),
						Sku:       to.Ptr("18.04-lts"),
						Version:   to.Ptr("latest"),
					},
				},
				OsProfile: &compute2020.VirtualMachineScaleSetOSProfile{
					ComputerNamePrefix: to.Ptr("computer"),
					AdminUsername:      &adminUsername,
					LinuxConfiguration: &compute2020.LinuxConfiguration{
						DisablePasswordAuthentication: to.Ptr(true),
						Ssh: &compute2020.SshConfiguration{
							PublicKeys: []compute2020.SshPublicKeySpec{
								{
									KeyData: sshPublicKey,
									Path:    to.Ptr(fmt.Sprintf("/home/%s/.ssh/authorized_keys", adminUsername)),
								},
							},
						},
					},
				},
				NetworkProfile: &compute2020.VirtualMachineScaleSetNetworkProfile{
					NetworkInterfaceConfigurations: []compute2020.VirtualMachineScaleSetNetworkConfiguration{
						{
							Name:    to.Ptr("mynicconfig"),
							Primary: to.Ptr(true),
							IpConfigurations: []compute2020.VirtualMachineScaleSetIPConfiguration{
								{
									Name: to.Ptr("myipconfiguration"),
									Subnet: &compute2020.ApiEntityReference{
										Reference: tc.MakeReferenceFromResource(subnet),
									},
									LoadBalancerInboundNatPools: []compute2020.SubResource{
										{
											Reference: &inboundNATPoolRef,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func Test_Compute_VMSS_20201201_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	// Move to a different region where we have quota
	tc.AzureRegion = to.Ptr("westeurope")

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	publicIPAddress := newPublicIPAddressForVMSS(tc, testcommon.AsOwner(rg))
	loadBalancer := newLoadBalancerForVMSS(tc, rg, publicIPAddress)
	// Have to create the vnet first there's a race between it and subnet creation that
	// can change the body of the VNET PUT (because VNET PUT contains subnets)
	tc.CreateResourceAndWait(vnet)
	tc.CreateResourcesAndWait(subnet, loadBalancer, publicIPAddress)
	vmss := newVMSS20201201(tc, rg, loadBalancer, subnet)

	tc.CreateResourceAndWait(vmss)
	tc.Expect(vmss.Status.Id).ToNot(BeNil())
	armId := *vmss.Status.Id

	// Perform a simple patch to add a basic custom script extension
	old := vmss.DeepCopy()
	extensionName := "mycustomextension"
	extensionName2 := "mycustomextension2"

	vmss.Spec.VirtualMachineProfile.ExtensionProfile = &compute2020.VirtualMachineScaleSetExtensionProfile{
		Extensions: []compute2020.VirtualMachineScaleSetExtension{
			{
				Name:               &extensionName,
				Publisher:          to.Ptr("Microsoft.Azure.Extensions"),
				Type:               to.Ptr("CustomScript"),
				TypeHandlerVersion: to.Ptr("2.0"),
				Settings: map[string]v1.JSON{
					"commandToExecute": {
						Raw: []byte(`"/bin/bash -c \"echo hello\""`),
					},
				},
			},
		},
	}
	tc.PatchResourceAndWait(old, vmss)
	tc.Expect(vmss.Status.VirtualMachineProfile).ToNot(BeNil())
	tc.Expect(vmss.Status.VirtualMachineProfile.ExtensionProfile).ToNot(BeNil())
	tc.Expect(len(vmss.Status.VirtualMachineProfile.ExtensionProfile.Extensions)).To(BeNumerically(">", 0))

	found := false
	for _, extension := range vmss.Status.VirtualMachineProfile.ExtensionProfile.Extensions {
		tc.Expect(extension.Name).ToNot(BeNil())
		if *extension.Name == extensionName {
			found = true
		}
	}
	tc.Expect(found).To(BeTrue())

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "VMSS_Extension_20201201_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				VMSS_Extension_20201201_CRUD(tc, vmss, extensionName2)
			},
		},
	)

	objectKey := client.ObjectKeyFromObject(vmss)

	tc.Eventually(func() bool {
		var updated compute2020.VirtualMachineScaleSet
		tc.GetResource(objectKey, &updated)
		return checkExtensionExists2020(tc, &updated, extensionName)
	}).Should(BeTrue())

	// Delete VMSS
	tc.DeleteResourceAndWait(vmss)

	// Ensure that the resource was really deleted in Azure
	tc.ExpectResourceIsDeletedInAzure(armId, string(compute2020.APIVersion_Value))
}

func checkExtensionExists2020(tc *testcommon.KubePerTestContext, vmss *compute2020.VirtualMachineScaleSet, extensionName string) bool {
	tc.Expect(vmss.Status.VirtualMachineProfile).ToNot(BeNil())
	tc.Expect(vmss.Status.VirtualMachineProfile.ExtensionProfile).ToNot(BeNil())
	tc.Expect(len(vmss.Status.VirtualMachineProfile.ExtensionProfile.Extensions)).To(BeNumerically(">", 0))

	found := false
	for _, extension := range vmss.Status.VirtualMachineProfile.ExtensionProfile.Extensions {
		tc.Expect(extension.Name).ToNot(BeNil())
		if *extension.Name == extensionName {
			found = true
		}
	}

	return found
}

func VMSS_Extension_20201201_CRUD(tc *testcommon.KubePerTestContext, vmss *compute2020.VirtualMachineScaleSet, extensionName string) {
	extension := &compute2020.VirtualMachineScaleSetsExtension{
		ObjectMeta: tc.MakeObjectMetaWithName(extensionName),
		Spec: compute2020.VirtualMachineScaleSetsExtension_Spec{
			Owner:              testcommon.AsOwner(vmss),
			Publisher:          to.Ptr("Microsoft.ManagedServices"),
			Type:               to.Ptr("ApplicationHealthLinux"),
			TypeHandlerVersion: to.Ptr("1.0"),
		},
	}

	tc.CreateResourceAndWait(extension)
	tc.Expect(extension.Status.Id).ToNot(BeNil())
	armId := *extension.Status.Id

	tc.DeleteResourceAndWait(extension)

	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(compute2020.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
