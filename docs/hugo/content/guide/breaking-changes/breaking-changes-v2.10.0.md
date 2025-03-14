---
title: "v2.10.0 Breaking Changes"
linkTitle: "v2.10.0"
weight: -40  # This should be 5 lower than the previous breaking change document
---

## MachineLearningServices/WorkspacesCompute properties marked as a SecretReference

The below properties on `WorkspacesCompute` has been changed from a string to a `SecretReference`
We always try to avoid breaking changes, but in this case, allowing raw passwords in the spec is a security problem and as such we've
decided to make a break to correct this issue.

**Affected Properties:**

- SslConfiguration.Key
- SslConfiguration.Cert
- VirtualMachineSshCredentials.PrivateKeyData
- VirtualMachineSshCredentials.PublicKeyData
- DatabricksProperties.DatabricksAccessToken

**Action required:** If the `MachineLearningServices/WorkspacesCompute` resource is used in your cluster and any of the above property is set, do the following before upgrading ASO:

1. Annotate the resource with `serviceoperator.azure.com/reconcile-policy: skip` to prevent ASO from trying to reconcile the resource while you are upgrading.
2. Download the current YAML for the resource using `kubectl` if you don't have it elsewhere.
3. Create a kubernetes secret containing the value for the affected property.
4. Edit downloaded YAML in step 2, and add a secret key and name reference. Example [here](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/apimanagement/v1api20230501preview/v1api20230501preview_authorizationprovider.yaml#L12).
5. Delete the resource from your cluster using `kubectl delete`. Your Azure resource will be left untouched because of the `reconcile-policy` annotation you added above.
6. [Upgrade ASO]( {{< relref "upgrading" >}} ) in your cluster.
7. Apply the updated YAML to your cluster using `kubectl apply`. If any errors occur, address them.
8. If the `reconcile-policy` annotation is still present, remove it from the resource.