// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

type IdBaseType struct {
	Id string `json:"id,omitempty"`
}

type IdRefBaseType struct {
	IdRef string `json:"idref,omitempty"`
}

type MarkingIdRefBaseType struct {
	MarkingIdRef string `json:"marking_idref,omitempty"`
}

type IndicatorBaseType struct {
	IdBaseType
	IdRefBaseType
	MarkingIdRefBaseType
	Timestamp string `json:"timestamp,omitempty"`
}

type GenericRelationshipType struct {
	Confidence        *ConfidenceType                 `json:"confidence,omitempty"`
	InformationSource *InformationSourceType          `json:"information_source,omitempty"`
	Relationship      *ControlledVocabularyStringType `json:"relationship,omitempty"`
}
