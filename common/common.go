// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

type ControlledVocabularyStringType struct {
	VocabName      string `json:"vocab_name,omitempty"`
	VocabReference string `json:"vocab_reference,omitempty"`
	Value          string `json:"value,omitempty"`
}

// TODO: change the value to use time.Time and import time
// Issues: Need to support ISO 8601, specifically 2015-01-15T14:16:00-07:00
type DateTimeWithPrecisionType struct {
	Precision string `json:"precision,omitempty"`
	Value     string `json:"value,omitempty"`
}

// TODO: Need to build out the Course of Action Idiom
type RelatedCourseOfActionType struct {
	Confidence        *ConfidenceType                 `json:"confidence,omitempty"`
	InformationSource *InformationSourceType          `json:"information_source,omitempty"`
	Relationship      *ControlledVocabularyStringType `json:"relationship,omitempty"`
	//	CourseOfAction
}

// TODO : Need to figure out what we call the TTPBaseType as that is not going to be correct.
type RelatedTTPType struct {
	Confidence        *ConfidenceType                 `json:"confidence,omitempty"`
	InformationSource *InformationSourceType          `json:"information_source,omitempty"`
	Relationship      *ControlledVocabularyStringType `json:"relationship,omitempty"`
	//TTP                *TTPBaseType                    `json:"ttp,omitempty"`
}
