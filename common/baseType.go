// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1

package common

type IndicatorBaseType struct {
	Id        string `json:"id,omitempty"`
	IdRef     string `json:"idref,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

type GenericRelationshipType struct {
	Confidence        *ConfidenceType                 `json:"confidence,omitempty"`
	InformationSource *InformationSourceType          `json:"informationSource,omitempty"`
	Relationship      *ControlledVocabularyStringType `json:"relationship,omitempty"`
}
