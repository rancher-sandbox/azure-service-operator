/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// InjectOriginalVersionFunctionStageID is the unique identifier for this pipeline stage
const InjectOriginalVersionFunctionStageID = "injectOriginalVersionFunction"

// InjectOriginalVersionFunction injects the function OriginalVersion() into each Spec type
// This function allows us to recover the original version used to create each custom resource, giving the operator the
// information needed to interact with ARM using the correct API version.
// We run this stage before we create any storage types, ensuring only API versions get the function.
func InjectOriginalVersionFunction(idFactory astmodel.IdentifierFactory) *Stage {
	stage := NewStage(
		InjectOriginalVersionFunctionStageID,
		"Inject the function OriginalVersion() into each Spec type",
		func(ctx context.Context, state *State) (*State, error) {
			definitions := state.Definitions()
			injector := astmodel.NewFunctionInjector()
			result := definitions.Copy()

			specs := astmodel.FindSpecDefinitions(definitions)
			for name, def := range specs {
				fn := functions.NewOriginalVersionFunction(idFactory)
				defWithFn, err := injector.Inject(def, fn)
				if err != nil {
					return nil, eris.Wrapf(err, "injecting OriginalVersion() into %s", name)
				}

				result[defWithFn.Name()] = defWithFn
			}

			return state.WithDefinitions(result), nil
		})

	stage.RequiresPostrequisiteStages(
		CreateStorageTypesStageID,
		InjectOriginalVersionPropertyStageID)

	return stage
}
