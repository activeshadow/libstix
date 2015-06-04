// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.2

package common

import (
	"github.com/freestix/libcybox/cybox"
)

type RelatedObservableType struct {
	GenericRelationshipType
	Observable *cybox.ObservableType `json:"observable,omitempty"`
}

// ----------------------------------------------------------------------
// Methods RelatedObservableType
// ----------------------------------------------------------------------

func (r *RelatedObservableType) AddConfidence(c ConfidenceType) {
	r.Confidence = &c
}

func (r *RelatedObservableType) AddInformationSource(source InformationSourceType) {
	r.InformationSource = &source
}

func (r *RelatedObservableType) AddRelationshipDetail(name, ref, value string) {
	data := ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	r.Relationship = &data
}

func (r *RelatedObservableType) AddObservable(o cybox.ObservableType) {
	r.Observable = &o
}
