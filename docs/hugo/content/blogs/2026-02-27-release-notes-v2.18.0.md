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

Support for the `allowMultiEnvManagement` configuration flag allows per-namespace/per-resource Azure cloud environment settings, allowing uesrs to manage resources across different Azure clouds/environments. 
- Special thanks to [subhamrajvanshi](https://github.com/shubhamrajvanshi) for this contribution.

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

### Added support for resources under AppConfiguration

New [AppConfiguration child resources](https://github.com/Azure/azure-service-operator/pull/4948) using the latest 2024-06-01 API version have been added to address the requests for KeyValue management capabilities. These resources include `KeyValue`, `Replica`, and `Snapshot`.

### MySQL Server version flexibility

The `Version` value for `dbformysql.FlexibleServer` [has been changed](https://github.com/Azure/azure-service-operator/pull/5195) to a simple string to allow new server versions to be used without the need for an API upgrade.

## 🛡️ Operational Resilience

### Improved retry on errors

ASO no longer gives up [retrying on common Azure errors](https://github.com/Azure/azure-service-operator/pull/5092), improving reliability for long-running operations in ASO. ASO will currently still give up retrying on some well-known non-retryable errors.

### Smarter deletes

Before issuing a DELETE request to Azure, [ASO now checks](https://github.com/Azure/azure-service-operator/pull/5040) if the resource or the resource's parent has already been removed.

### Remove usage of `NewDefaultAzureCredential`

`NewDefaultAzureCredential` is replaced with `NewChainedTokenCredential` alongside a more [limited set of credentials](https://github.com/Azure/azure-service-operator/pull/5155). Authentication options that were removed were generally not recommended for running in production.

## 🐛 Bug fixes

- [The Helm chart metrics port is now fully configurable](https://github.com/Azure/azure-service-operator/pull/5156). Previously, the `metrics.port` value in `values.yaml` was ignored in favor of a hardcoded `8443`.
  - _Special thanks to [bingikarthik](https://github.com/bingikarthik) for the contribution!_
- [Replaced `NewDefaultAzureCredential` with a narrowly scoped credential chain](https://github.com/Azure/azure-service-operator/pull/5155), improving security and predictability of authentication behavior for both the controller and `asoctl`.
- [Fixed the RoleDefinition cache](https://github.com/Azure/azure-service-operator/pull/5133) to ensure deterministic test behavior when well-known RoleDefinition names like `Contributor` are referenced.
- 4 known cases where the ARM ID properties were released [without being properly tagged-as/converted-to](https://github.com/Azure/azure-service-operator/pull/4925) `KnownResourceReference` structs were fixed. To [avoid potential breakage](https://github.com/Azure/azure-service-operator/pull/4925), users can add the correct `ResourceReference` property alongside the legacy property.

## 🙏 Thank You

Thank you to all our contributors for making this release possible! A special thanks goes to [bingikarthik](https://github.com/bingikarthik) and [subhamrajvanshi](https://github.com/shubhamrajvanshi) for their contributions to this release.

**Full Changelog**: [v2.17.0...v2.18.0](https://github.com/Azure/azure-service-operator/compare/v2.17.0...v2.18.0)
