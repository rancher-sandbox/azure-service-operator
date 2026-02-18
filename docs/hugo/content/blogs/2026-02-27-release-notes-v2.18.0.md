---
title: "ASO v2.18 Release Notes"
date: 2026-02-18
description: "Release notes for Azure Service Operator v2.18.0"
type: blog
---

We're excited to announce the release of Azure Service Operator v2.18.0! This release adds new AKS preview API support, expands simplified versioning to storage resources, strengthens reconciliation checks, fixes Helm configurability issues, and updates a broad set of dependencies and tooling.

## 🎇 Headline Features

The `storage` group now supports [hybrid API versioning](https://github.com/Azure/azure-service-operator/pull/5116). You can use all storage resources with the new `v` prefix, while the legacy `v1api` versions remain available for backward compatibility. This continues our broader transition toward simplified versioning across all resource groups.

ASO now issues a fresh `GET` on the owner ARM resource before calling `PreReconcileOwnerCheck`, ensuring the check always operates on the latest owner status ([#5140](https://github.com/Azure/azure-service-operator/pull/5140)). The `owner` parameter has also been removed from `PreReconcileCheck` to enforce a cleaner separation of concerns between the two extension points.

## ⚠️ Breaking changes

This release includes breaking changes. Please review the [breaking changes documentation](https://azure.github.io/azure-service-operator/guide/breaking-changes/) before upgrading.

### Removed AKS API versions

This release removes the `containerservice` ManagedCluster and AgentPool API versions `v1api20230201` and `v1api20231001`, as [previously announced](https://github.com/Azure/azure-service-operator/releases/tag/v2.17.0).

* If you allow the operator to manage its own CRDs via `--crd-pattern`, no action is needed—the operator handles removing these versions automatically.
* If you manage CRD versions yourself, you must run [asoctl clean crds](https://azure.github.io/azure-service-operator/tools/asoctl/#clean-crds) before upgrading.

### Upcoming breaking changes

In [ASO v2.19](https://github.com/Azure/azure-service-operator/milestone/38), we will remove `containerservice` ManagedCluster and AgentPool API version `v1api20240402preview`.

## 🎉 New and improved resource support

### New AKS preview API version

You can now use the latest AKS preview features through [containerservice v20251002preview](https://github.com/Azure/azure-service-operator/pull/5170).

### Updated insights API versions

Several `insights` resources now use their latest ARM API versions ([#5136](https://github.com/Azure/azure-service-operator/pull/5136)):

- **DataCollectionEndpoint**, **DataCollectionRule**, and **DataCollectionRuleAssociation** now target API version `2024-03-11`
- **ScheduledQueryRule** now targets `2025-01-01-preview`

## 🐛 Bug fixes

- [The Helm chart metrics port is now fully configurable](https://github.com/Azure/azure-service-operator/pull/5156). Previously, the `metrics.port` value in `values.yaml` was ignored in favor of a hardcoded `8443`.
  - _Special thanks to [bingikarthik](https://github.com/bingikarthik) for the contribution!_
- [Replaced `NewDefaultAzureCredential` with a narrowly scoped credential chain](https://github.com/Azure/azure-service-operator/pull/5155), improving security and predictability of authentication behavior for both the controller and `asoctl`.
- [Fixed the RoleDefinition cache](https://github.com/Azure/azure-service-operator/pull/5133) to ensure deterministic test behavior when well-known RoleDefinition names like `Contributor` are referenced.

## 🙏 Thank You

Thank you to all our contributors for making this release possible! We're especially grateful to [bingikarthik](https://github.com/bingikarthik)for the contribution to this release.

**Full Changelog**: [v2.17.0...v2.18.0](https://github.com/Azure/azure-service-operator/compare/v2.17.0...v2.18.0)
