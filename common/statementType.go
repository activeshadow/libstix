// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.2

package common

import (
	"time"
)

type StatementType struct {
	Timestamp          string                          `json:"timestamp,omitempty"`
	TimestampPrecision string                          `json:"timestampPrecision,omitempty"`
	Value              *ControlledVocabularyStringType `json:"value,omitempty"`
	Description        *StructuredTextType             `json:"description,omitempty"`
	Source             *InformationSourceType          `json:"source,omitempty"`
	Confidence         *ConfidenceType                 `json:"confidence,omitempty"`
}

// ----------------------------------------------------------------------
// Methods StatementType
// ----------------------------------------------------------------------

func (s *StatementType) CreateTimeStamp() {
	s.Timestamp = time.Now().Format(time.RFC3339)
}

func (s *StatementType) AddTimeStamp(t string) {
	// TODO Need to format the string in to ISO 8601 format or check that it is in the right format
	s.Timestamp = t
}

func (s *StatementType) AddTimeStampPrecision(p string) {
	s.TimestampPrecision = p
}

func (s *StatementType) AddValueDetail(name, ref, value string) {
	data := ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	s.Value = &data
}

func (s *StatementType) AddDescription(value string) {
	data := StructuredTextType{
		StructuringFormat: "txt",
		Value:             value,
	}
	s.Description = &data
}

func (s *StatementType) AddSource(source InformationSourceType) {
	s.Source = &source
}

func (s *StatementType) AddConfidence(c ConfidenceType) {
	s.Confidence = &c
}
