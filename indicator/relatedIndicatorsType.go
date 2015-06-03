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

type RelatedIndicatorsType struct {
	Scope            string                 `json:"scope,omitempty"`
	RelatedIndicator []RelatedIndicatorType `json:"relatedIndicators,omitempty"`
}

type RelatedIndicatorType struct {
	common.GenericRelationshipType
	Indicator *IndicatorType `json:"indicator,omitempty"`
}

// ----------------------------------------------------------------------
// Methods RelatedIndicatorsType
// ----------------------------------------------------------------------

func (r *RelatedIndicatorsType) AddScope(s string) {
	r.Scope = s
}

func (r *RelatedIndicatorsType) AddRelatedIndicator(i RelatedIndicatorType) {
	if r.RelatedIndicator == nil {
		a := make([]RelatedIndicatorType, 0)
		r.RelatedIndicator = a
	}
	r.RelatedIndicator = append(r.RelatedIndicator, i)
}

// ----------------------------------------------------------------------
// Methods RelatedIndicatorType
// ----------------------------------------------------------------------

func (r *RelatedIndicatorType) AddConfidence(c common.ConfidenceType) {
	r.Confidence = &c
}

func (r *RelatedIndicatorType) AddInformationSource(i common.InformationSourceType) {
	r.InformationSource = &i
}

func (r *RelatedIndicatorType) AddRelationshipDetail(name, ref, value string) {
	data := common.ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	r.Relationship = &data
}

func (r *RelatedIndicatorType) AddIndicator(i IndicatorType) {
	r.Indicator = &i
}
