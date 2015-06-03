// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.2

package indicator

import (
	"github.com/jordan2175/freestix/libstix/common"
	"time"
)

type SightingsType struct {
	SightingsCount int            `json:"sightingsCount,omitempty"`
	Sighting       []SightingType `json:"sighting,omitempty"`
}

type SightingType struct {
	Timestamp          string                        `json:"timestamp,omitempty"`
	TimestampPrecision string                        `json:"timestampPrecision,omitempty"`
	Source             *common.InformationSourceType `json:"source,omitempty"`
	References         []string                      `json:"references,omitempty"`
	Confidence         *common.ConfidenceType        `json:"confidence,omitempty"`
	Description        *common.StructuredTextType    `json:"description,omitempty"`
	RelatedObservables *RelatedObservablesType       `json:"relatedObservables,omitempty"`
}

// ----------------------------------------------------------------------
// Methods SightingsType
// ----------------------------------------------------------------------

func (s *SightingsType) SetSightingsCount(count int) {
	s.SightingsCount = count
}

func (s *SightingsType) AddSighting(seen SightingType) {
	if s.Sighting == nil {
		a := make([]SightingType, 0)
		s.Sighting = a
	}
	s.Sighting = append(s.Sighting, seen)
}

// ----------------------------------------------------------------------
// Methods SightingType
// ----------------------------------------------------------------------

func (s *SightingType) CreateTimeStamp() {
	s.Timestamp = time.Now().Format(time.RFC3339)
}

func (s *SightingType) AddTimeStamp(t string) {
	// TODO Need to format the string in to ISO 8601 format or check that it is in the right format
	s.Timestamp = t
}

func (s *SightingType) AddTimeStampPrecision(p string) {
	s.TimestampPrecision = p
}

func (s *SightingType) AddSource(source common.InformationSourceType) {
	s.Source = &source
}

func (s *SightingType) AddReference(r string) {
	if s.References == nil {
		a := make([]string, 0)
		s.References = a
	}
	s.References = append(s.References, r)
}

func (s *SightingType) AddConfidence(c common.ConfidenceType) {
	s.Confidence = &c
}

func (s *SightingType) AddDescription(value string) {
	data := common.StructuredTextType{
		StructuringFormat: "txt",
		Value:             value,
	}
	s.Description = &data
}

func (s *SightingType) AddRelatedObservables(r RelatedObservablesType) {
	s.RelatedObservables = &r
}
