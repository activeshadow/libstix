// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

import (
	"github.com/activeshadow/libcybox/common"
	"github.com/activeshadow/libstix/defs"
	"time"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------
// TODO Tools

type InformationSourceType struct {
	Identity            *IdentityType                   `json:"identity,omitempty"`
	Role                *ControlledVocabularyStringType `json:"role,omitempty"`
	Description         *StructuredTextType             `json:"description,omitempty"`
	References          []string                        `json:"references,omitempty"`
	Time                *common.TimeType                `json:"time,omitempty"`
	Tools               []common.ToolInformationType    `json:"tools,omitempty"`
	ContributingSources []InformationSourceType         `json:"contributing_sources,omitempty"`
}

// ----------------------------------------------------------------------
// Methods InformationSourceType
// ----------------------------------------------------------------------

func (this *InformationSourceType) AddDescriptionText(value string) {
	this.AddDescription(defs.RFC6838TEXTPLAIN, value)
}

func (this *InformationSourceType) AddDescription(format, value string) {
	var data StructuredTextType
	data.AddFormat(format)
	data.AddValue(value)

	this.Description = &data
}

func (this *InformationSourceType) AddIdentity(id IdentityType) {
	this.Identity = &id
}

func (this *InformationSourceType) AddRoleDetail(name, ref, value string) {
	data := ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	this.Role = &data
}

func (s *InformationSourceType) AddContributingSource(source InformationSourceType) {
	if s.ContributingSources == nil {
		a := make([]InformationSourceType, 0)
		s.ContributingSources = a
	}
	s.ContributingSources = append(s.ContributingSources, source)
}

// We need to initalize the time data structure if it is not there yet, otherwise we can not use it
func (s *InformationSourceType) InitTime() {
	if s.Time == nil {
		var ctime common.TimeType
		s.Time = &ctime
	}
}

func (s *InformationSourceType) AddStartTime(t string) {
	s.InitTime()
	t1 := common.DateTimeWithPrecisionType{
		Precision: "0.0",
		Value:     t,
	}
	s.Time.StartTime = &t1
}

func (s *InformationSourceType) AddEndTime(t string) {
	s.InitTime()
	t1 := common.DateTimeWithPrecisionType{
		Precision: "0.0",
		Value:     t,
	}
	s.Time.EndTime = &t1
}

func (this *InformationSourceType) SetProducedTimeToNow() {
	this.AddProducedTime(time.Now().Format(time.RFC3339))
}

func (s *InformationSourceType) AddProducedTime(t string) {
	s.InitTime()
	t1 := common.DateTimeWithPrecisionType{
		Precision: "0.0",
		Value:     t,
	}
	s.Time.ProducedTime = &t1
}

func (s *InformationSourceType) AddReceivedTime(t string) {
	s.InitTime()
	t1 := common.DateTimeWithPrecisionType{
		Precision: "0.0",
		Value:     t,
	}
	s.Time.ReceivedTime = &t1
}

func (s *InformationSourceType) AddReference(r string) {
	if s.References == nil {
		a := make([]string, 0)
		s.References = a
	}
	s.References = append(s.References, r)
}
