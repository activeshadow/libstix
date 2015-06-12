// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

import (
	"github.com/freestix/libstix/defs"
	"time"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

type ConfidenceType struct {
	Timestamp                string                          `json:"timestamp,omitempty"`
	TimestampPrecision       string                          `json:"timestamp_precision,omitempty"`
	Value                    *ControlledVocabularyStringType `json:"value,omitempty"`
	Description              *StructuredTextType             `json:"description,omitempty"`
	Source                   *InformationSourceType          `json:"source,omitempty"`
	ConfidenceAssertionChain []ConfidenceType                `json:"confidence_assertion_chain,omitempty"`
}

// ----------------------------------------------------------------------
// Methods ConfidenceType
// ----------------------------------------------------------------------

func (this *ConfidenceType) CreateTimeStamp() {
	this.Timestamp = time.Now().Format(time.RFC3339)
}

func (this *ConfidenceType) AddTimeStamp(t string) {
	// TODO Need to format the string in to ISO 8601 format or check that it is in the right format
	this.Timestamp = t
}

func (this *ConfidenceType) AddTimeStampPrecision(p string) {
	this.TimestampPrecision = p
}

func (this *ConfidenceType) AddValueDetail(name, ref, value string) {
	data := ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	this.Value = &data
}

func (this *ConfidenceType) AddDescriptionText(value string) {
	this.AddDescription(defs.RFC6838TEXTPLAIN, value)
}

func (this *ConfidenceType) AddDescription(format, value string) {
	var data StructuredTextType
	data.AddFormat(format)
	data.AddValue(value)

	this.Description = &data
}

func (this *ConfidenceType) AddSource(source InformationSourceType) {
	this.Source = &source
}

func (this *ConfidenceType) AddConfidenceAssertion(confidence ConfidenceType) {
	if this.ConfidenceAssertionChain == nil {
		a := make([]ConfidenceType, 0)
		this.ConfidenceAssertionChain = a
	}
	this.ConfidenceAssertionChain = append(this.ConfidenceAssertionChain, confidence)
}
