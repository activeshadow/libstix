// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

import (
	"github.com/activeshadow/libcybox/observable"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

type RelatedObservableType struct {
	GenericRelationshipType
	Observable *observable.ObservableType `json:"observable,omitempty"`
}

// ----------------------------------------------------------------------
// Methods RelatedObservableType
// ----------------------------------------------------------------------

func (this *RelatedObservableType) AddConfidence(c ConfidenceType) {
	this.Confidence = &c
}

func (this *RelatedObservableType) AddInformationSource(source InformationSourceType) {
	this.InformationSource = &source
}

func (this *RelatedObservableType) AddRelationshipDetail(name, ref, value string) {
	data := ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	this.Relationship = &data
}

func (this *RelatedObservableType) AddObservable(o observable.ObservableType) {
	this.Observable = &o
}
