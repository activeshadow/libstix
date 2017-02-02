// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package indicator

import (
	"github.com/activeshadow/libstix/common"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

type RelatedObservablesType struct {
	Scope             string                         `json:"scope,omitempty"`
	RelatedObservable []common.RelatedObservableType `json:"related_observable,omitempty"`
}

// ----------------------------------------------------------------------
// Create Functions
// ----------------------------------------------------------------------

func CreateRelatedObservables() RelatedObservablesType {
	var obj RelatedObservablesType
	return obj
}

// ----------------------------------------------------------------------
// Methods RelatedObservablesType
// ----------------------------------------------------------------------

func (this *RelatedObservablesType) AddScope(s string) {
	this.Scope = s
}

func (this *RelatedObservablesType) AddRelatedObservable(o common.RelatedObservableType) {
	if this.RelatedObservable == nil {
		a := make([]common.RelatedObservableType, 0)
		this.RelatedObservable = a
	}
	this.RelatedObservable = append(this.RelatedObservable, o)
}
