---
title: "Authentication in Azure Service Operator"
linkTitle: "Authentication"
layout: single
weight: 1 # This is the default weight if you just want to be ordered alphabetically
cascade:
- type: docs
- render: always
---

There are two key topics surrounding authentication in Azure Service Operator: The type of credential, and the credential scope.

## Credential type

Azure Service Operator supports five different styles of authentication today.

1. [Recommended for production] [Azure Workload Identity]( {{< relref "credential-format#managed-identity-via-workload-identity" >}} ) (OIDC + Managed Identity or Service Principal)
2. [Service Principal using a Client Secret]( {{< relref "credential-format#service-principal-using-a-client-secret" >}} )
3. [Service Principal using a Client Certificate]( {{< relref "credential-format#service-principal-using-a-client-certificate" >}} )
4. [Deprecated] [aad-pod-identity authentication (Managed Identity)]( {{< relref "credential-format#deprecated-managed-identity-aad-pod-identity" >}} )
5. [User Assigned Identity Credentials] ( {{< relref "credential-format#user-assigned-identity-credentials" >}} )

## Credential scope

Each supported credential type can be specified at one of three supported scopes:

1. [Not recommended] [Global]( {{< relref "credential-scope#global-scope" >}} ) - The credential applies to all resources managed by ASO.
2. [Namespace]( {{< relref "credential-scope#namespace-scope" >}} ) - The credential applies to all resources managed by ASO in that namespace.
3. [Resource]( {{< relref "credential-scope#resource-scope" >}} ) - The credential applies to only the specific resource it is referenced on.

When presented with multiple credential choices, the operator chooses the most specific one:
_resource scope_ takes precedence over _namespace scope_ which takes precedence over _global scope_.

> **Warning:** The operator identity is used to access the global, namespace, and resource scoped secrets. A user with
> access to create ASO resources but without Kubernetes secret read permissions can still direct ASO to use a secret the
> user cannot read. The user cannot access the contents of the secret, but they can manage resources in Azure
> via the identity the secret represents.
>
> The namespace is the security boundary. ASO will not allow users to read secrets from other namespaces. We recommend
> using separate namespaces for separate environments (dev, test, prod, etc) for this reason

## Using multiple operators with a single credential per operator

> **This mode is not recommended unless you _really_ need it**

ASO also supports installing multiple instances of the operator alongside one another, each
configured with different credentials.

This option exists for the most security conscious customers. 

Advantages:
* Each operator pod only has access to a single credential, reducing risk if one pod is somehow compromised.

Disadvantages:
* Significantly harder to orchestrate ASO upgrades.
* More kube-apiserver load, as there will be multiple operators running and watching/reconciling resources.

For more details about this approach, see [multitenant deployment]( {{< relref "multitenant-deployment" >}} )