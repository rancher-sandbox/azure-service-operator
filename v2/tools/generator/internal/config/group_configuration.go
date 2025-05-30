/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"strings"

	"github.com/rotisserie/eris"
	"gopkg.in/yaml.v3"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/typo"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// GroupConfiguration contains additional information about an entire group and forms the top of a hierarchy containing
// information to supplement the schema and swagger sources consumed by the generator.
//
// ┌──────────────────────────┐       ╔════════════════════╗       ┌──────────────────────┐       ┌───────────────────┐       ┌───────────────────────┐
// │                          │       ║                    ║       │                      │       │                   │       │                       │
// │ ObjectModelConfiguration │───────║ GroupConfiguration ║───────│ VersionConfiguration │───────│ TypeConfiguration │───────│ PropertyConfiguration │
// │                          │1  1..n║                    ║1  1..n│                      │1  1..n│                   │1  1..n│                       │
// └──────────────────────────┘       ╚════════════════════╝       └──────────────────────┘       └───────────────────┘       └───────────────────────┘
type GroupConfiguration struct {
	name     string
	versions map[string]*VersionConfiguration
	advisor  *typo.Advisor
	// Configurable properties here (alphabetical, please)
	PayloadType configurable[PayloadType] // Specify how this property should be serialized for ARM
}

type PayloadType string

const (
	OmitEmptyProperties      PayloadType = "omitempty"                // Omit all empty properties even collections
	ExplicitCollections      PayloadType = "explicitcollections"      // Always include collections (as null), omit other empty properties
	ExplicitEmptyCollections PayloadType = "explicitemptycollections" // Always include collections (as empty map/array), omit other empty properties
	ExplicitProperties       PayloadType = "explicitproperties"       // Always include all properties
)

const (
	payloadTypeTag = "$payloadType" // Enumeration specifying what kind of payload to send to ARM.
)

// NewGroupConfiguration returns a new (empty) GroupConfiguration
func NewGroupConfiguration(name string) *GroupConfiguration {
	scope := "group " + name
	return &GroupConfiguration{
		name:     name,
		versions: make(map[string]*VersionConfiguration),
		advisor:  typo.NewAdvisor(),
		// Initialize configurable properties here (alphabetical, please)
		PayloadType: makeConfigurable[PayloadType](payloadTypeTag, scope),
	}
}

// Add includes configuration for the specified version as a part of this group configuration
// In addition to indexing by the name of the version, we also index by the local-package-name and storage-package-name
// of the version, so we can do lookups via TypeName. All indexing is lower-case to allow case-insensitive lookups (this
// makes our configuration more forgiving).
func (gc *GroupConfiguration) addVersion(name string, version *VersionConfiguration) {
	// Convert version.name into a package version
	// We do this by constructing a local package reference because this avoids replicating the logic here and risking
	// inconsistency if things are changed in the future.
	local := astmodel.MakeLocalPackageReference("prefix", "group", astmodel.GeneratorVersion, name)

	gc.versions[strings.ToLower(name)] = version
	gc.versions[strings.ToLower(local.Version())] = version
}

// visitVersion invokes the provided visitor on the specified version if present.
// Returns a NotConfiguredError if the version is not found; otherwise whatever error is returned by the visitor.
func (gc *GroupConfiguration) visitVersion(
	ref astmodel.PackageReference,
	visitor *configurationVisitor,
) error {
	vc := gc.findVersion(ref)
	if vc == nil {
		return nil
	}

	err := visitor.visitVersion(vc)
	if err != nil {
		return eris.Wrapf(err, "configuration of group %s", gc.name)
	}

	return nil
}

// visitVersions invokes the provided visitor on all versions.
func (gc *GroupConfiguration) visitVersions(visitor *configurationVisitor) error {
	// All our versions are listed under multiple keys, so we hedge against processing them multiple times
	versionsSeen := set.Make[string]()

	errs := make([]error, 0, len(gc.versions))
	for _, v := range gc.versions {
		if versionsSeen.Contains(v.name) {
			continue
		}

		err := visitor.visitVersion(v)
		err = gc.advisor.Wrapf(err, v.name, "version %s not seen", v.name)
		errs = append(errs, err)

		versionsSeen.Add(v.name)
	}

	// Both errors.Wrapf() and kerrors.NewAggregate() return nil if nothing went wrong
	return eris.Wrapf(
		kerrors.NewAggregate(errs),
		"group %s",
		gc.name)
}

// findVersion uses the provided PackageReference to work out which nested VersionConfiguration should be used
func (gc *GroupConfiguration) findVersion(ref astmodel.PackageReference) *VersionConfiguration {
	switch r := ref.(type) {
	case astmodel.DerivedPackageReference:
		return gc.findVersion(r.Base())
	case astmodel.LocalPackageReference:
		return gc.findVersionForLocalPackageReference(r)
	}

	panic(fmt.Sprintf("didn't expect PackageReference of type %T", ref))
}

// findVersion uses the provided LocalPackageReference to work out which nested VersionConfiguration should be used
func (gc *GroupConfiguration) findVersionForLocalPackageReference(ref astmodel.LocalPackageReference) *VersionConfiguration {
	gc.advisor.AddTerm(ref.APIVersion())
	gc.advisor.AddTerm(ref.PackageName())

	// Check based on the ApiVersion alone
	apiKey := strings.ToLower(ref.APIVersion())
	if version, ok := gc.versions[apiKey]; ok {
		// make sure there's an exact match on the actual version name, so we don't generate a recommendation
		gc.advisor.AddTerm(version.name)
		return version
	}

	// Also check the entire package name (allows config to specify just a particular generator version if needed)
	pkgKey := strings.ToLower(ref.PackageName())
	if version, ok := gc.versions[pkgKey]; ok {
		// make sure there's an exact match on the actual version name, so we don't generate a recommendation
		gc.advisor.AddTerm(version.name)
		return version
	}

	return nil
}

// UnmarshalYAML populates our instance from the YAML.
// The slice node.Content contains pairs of nodes, first one for an ID, then one for the value.
func (gc *GroupConfiguration) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return eris.New("expected mapping")
	}

	gc.versions = make(map[string]*VersionConfiguration)
	var lastID string

	for i, c := range value.Content {
		// Grab identifiers and loop to handle the associated value
		if i%2 == 0 {
			lastID = c.Value
			continue
		}

		// Handle nested version metadata
		if c.Kind == yaml.MappingNode {
			v := NewVersionConfiguration(lastID)
			err := c.Decode(&v)
			if err != nil {
				return eris.Wrapf(err, "decoding yaml for %q", lastID)
			}

			gc.addVersion(lastID, v)
			continue
		}

		// $payloadType: <string>
		if strings.EqualFold(lastID, payloadTypeTag) && c.Kind == yaml.ScalarNode {
			switch strings.ToLower(c.Value) {
			case string(OmitEmptyProperties):
				gc.PayloadType.Set(OmitEmptyProperties)
			case string(ExplicitCollections):
				gc.PayloadType.Set(ExplicitCollections)
			case string(ExplicitEmptyCollections):
				gc.PayloadType.Set(ExplicitEmptyCollections)
			case string(ExplicitProperties):
				gc.PayloadType.Set(ExplicitProperties)
			default:
				return eris.Errorf("unknown %s value: %s.", payloadTypeTag, c.Value)
			}

			continue
		}

		// No handler for this value, return an error
		return eris.Errorf(
			"group configuration, unexpected yaml value %s: %s (line %d col %d)", lastID, c.Value, c.Line, c.Column)

	}

	return nil
}
