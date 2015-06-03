// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.2

package indicator

import (
	"github.com/jordan2175/freestix/libstix/common"
)

type RelatedObservablesType struct {
	Scope             string                         `json:"scope,omitempty"`
	RelatedObservable []common.RelatedObservableType `json:"relatedObservable,omitempty"`
}

// ----------------------------------------------------------------------
// Methods RelatedObservablesType
// ----------------------------------------------------------------------

func (r *RelatedObservablesType) AddScope(s string) {
	r.Scope = s
}

func (r *RelatedObservablesType) AddRelatedObservable(o common.RelatedObservableType) {
	if r.RelatedObservable == nil {
		a := make([]common.RelatedObservableType, 0)
		r.RelatedObservable = a
	}
	r.RelatedObservable = append(r.RelatedObservable, o)
}
