/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"fmt"

	gomegaformat "github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func actualAsObj(actual interface{}) (client.Object, error) {
	obj, ok := actual.(client.Object)
	if !ok {
		return nil, eris.Errorf("expected client.Object, was: %T", actual)
	}

	return obj, nil
}

type BeDeletedMatcher struct {
	verify *Verify
	ctx    context.Context

	subsequentMissingDeleteTimestamps int
}

var _ types.GomegaMatcher = &BeDeletedMatcher{}

func (m *BeDeletedMatcher) Match(actual interface{}) (bool, error) {
	if actual == nil {
		return false, nil
	}

	obj, err := actualAsObj(actual)
	if err != nil {
		return false, err
	}

	return m.verify.Deleted(m.ctx, obj)
}

func (m *BeDeletedMatcher) message(actual interface{}, negate bool) string {
	obj, err := actualAsObj(actual)
	if err != nil {
		// Gomegas contract is that it won't call one of the message functions
		// if Match returned an error. If we make it here somehow that contract
		// has been violated
		panic(err)
	}

	notStr := ""
	if negate {
		notStr = "not "
	}

	return gomegaformat.Message(obj, fmt.Sprintf("%sto be deleted", notStr))
}

func (m *BeDeletedMatcher) FailureMessage(actual interface{}) string {
	return m.message(actual, false)
}

func (m *BeDeletedMatcher) NegatedFailureMessage(actual interface{}) string {
	return m.message(actual, true)
}

// MatchMayChangeInTheFuture implements OracleMatcher which of course isn't exported so we can't type-assert we implement it
func (m *BeDeletedMatcher) MatchMayChangeInTheFuture(actual interface{}) bool {
	if actual == nil {
		return false
	}

	obj, err := actualAsObj(actual)
	if err != nil {
		panic(err)
	}

	// Initial object may not have a deletion timestamp set yet so look instead
	// for subsequent calls that all don't have timestamp
	deletionTimestamp := obj.GetDeletionTimestamp()
	if deletionTimestamp == nil {
		m.subsequentMissingDeleteTimestamps++
	}
	return m.subsequentMissingDeleteTimestamps <= 1
}
