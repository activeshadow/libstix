// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1
// TODO Tools

package common

import (
	"github.com/freestix/libcybox/common"
)

type InformationSourceType struct {
	Description         *StructuredTextType             `json:"description,omitempty"`
	Identity            *IdentityType                   `json:"identity,omitempty"`
	Role                *ControlledVocabularyStringType `json:"role,omitempty"`
	ContributingSources []InformationSourceType         `json:"contributingSources,omitempty"`
	Time                *common.TimeType                `json:"time,omitempty"`
	Tools               []common.ToolInformationType    `json:"tools,omitempty"`
	References          []string                        `json:"references,omitempty"`
}

// ----------------------------------------------------------------------
// Methods InformationSourceType
// ----------------------------------------------------------------------

func (s *InformationSourceType) AddDescription(value string) {
	data := StructuredTextType{
		StructuringFormat: "txt",
		Value:             value,
	}
	s.Description = &data
}

func (s *InformationSourceType) AddIdentity(id IdentityType) {
	s.Identity = &id
}

func (s *InformationSourceType) AddRoleDetail(name, ref, value string) {
	data := ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	s.Role = &data
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
