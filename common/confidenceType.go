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

type ConfidenceType struct {
	Timestamp                string                          `json:"timestamp,omitempty"`
	TimestampPrecision       string                          `json:"timestampPrecision,omitempty"`
	Value                    *ControlledVocabularyStringType `json:"value,omitempty"`
	Description              *StructuredTextType             `json:"description,omitempty"`
	Source                   *InformationSourceType          `json:"source,omitempty"`
	ConfidenceAssertionChain []ConfidenceType                `json:"confidenceAssertionChain,omitempty"`
}

// ----------------------------------------------------------------------
// Methods ConfidenceType
// ----------------------------------------------------------------------

func (c *ConfidenceType) CreateTimeStamp() {
	c.Timestamp = time.Now().Format(time.RFC3339)
}

func (c *ConfidenceType) AddTimeStamp(t string) {
	// TODO Need to format the string in to ISO 8601 format or check that it is in the right format
	c.Timestamp = t
}

func (c *ConfidenceType) AddTimeStampPrecision(p string) {
	c.TimestampPrecision = p
}

func (c *ConfidenceType) AddValueDetail(name, ref, value string) {
	data := ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	c.Value = &data
}

func (c *ConfidenceType) AddDescription(value string) {
	data := StructuredTextType{
		StructuringFormat: "txt",
		Value:             value,
	}
	c.Description = &data
}

func (c *ConfidenceType) AddSource(source InformationSourceType) {
	c.Source = &source
}

func (c *ConfidenceType) AddConfidenceAssertion(confidence ConfidenceType) {
	if c.ConfidenceAssertionChain == nil {
		a := make([]ConfidenceType, 0)
		c.ConfidenceAssertionChain = a
	}
	c.ConfidenceAssertionChain = append(c.ConfidenceAssertionChain, confidence)
}
