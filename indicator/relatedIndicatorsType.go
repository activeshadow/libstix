// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package indicator

import (
	"github.com/freestix/libstix/common"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

type RelatedIndicatorsType struct {
	Scope            string                 `json:"scope,omitempty"`
	RelatedIndicator []RelatedIndicatorType `json:"related_indicators,omitempty"`
}

type RelatedIndicatorType struct {
	common.GenericRelationshipType
	Indicator *IndicatorType `json:"indicator,omitempty"`
}

// ----------------------------------------------------------------------
// Create Functions
// ----------------------------------------------------------------------

func CreateRelatedIndicator() RelatedIndicatorType {
	var obj RelatedIndicatorType
	return obj
}

func CreateRelatedIndicators() RelatedIndicatorsType {
	var obj RelatedIndicatorsType
	return obj
}

// ----------------------------------------------------------------------
// Methods RelatedIndicatorsType
// ----------------------------------------------------------------------

func (this *RelatedIndicatorsType) AddScope(s string) {
	this.Scope = s
}

func (this *RelatedIndicatorsType) AddRelatedIndicator(i RelatedIndicatorType) {
	if this.RelatedIndicator == nil {
		a := make([]RelatedIndicatorType, 0)
		this.RelatedIndicator = a
	}
	this.RelatedIndicator = append(this.RelatedIndicator, i)
}

// ----------------------------------------------------------------------
// Methods RelatedIndicatorType
// ----------------------------------------------------------------------

func (this *RelatedIndicatorType) AddConfidence(c common.ConfidenceType) {
	this.Confidence = &c
}

func (this *RelatedIndicatorType) AddInformationSource(i common.InformationSourceType) {
	this.InformationSource = &i
}

func (this *RelatedIndicatorType) AddRelationshipDetail(name, ref, value string) {
	data := common.ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	this.Relationship = &data
}

func (this *RelatedIndicatorType) AddIndicator(i IndicatorType) {
	this.Indicator = &i
}
