// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

import (
	"github.com/freestix/libstix"
	"github.com/pborman/uuid"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

type IdentityType struct {
	Id                string                `json:"id,omitempty"`
	IdRef             string                `json:"idref,omitempty"`
	Name              string                `json:"name,omitempty"`
	RelatedIdentities []RelatedIdentityType `json:"related_identites,omitempty"`
}

type RelatedIdentityType struct {
	Confidence        *ConfidenceType                 `json:"confidence,omitempty"`
	InformationSource *InformationSourceType          `json:"information_source,omitempty"`
	Relationship      *ControlledVocabularyStringType `json:"relationship,omitempty"`
	Identity          *IdentityType                   `json:"identity,omitempty"`
}

// ----------------------------------------------------------------------
// Methods IdentityType
// ----------------------------------------------------------------------

func (this *IdentityType) CreateId() {
	this.Id = libstix.Company + ":identity-" + uuid.New()
}

func (this *IdentityType) AddIdRef(idref string) {
	this.IdRef = idref
}

func (this *IdentityType) AddName(name string) {
	this.Name = name
}

func (this *IdentityType) AddRelatedIdentity(iden RelatedIdentityType) {
	if this.RelatedIdentities == nil {
		a := make([]RelatedIdentityType, 0)
		this.RelatedIdentities = a
	}
	this.RelatedIdentities = append(this.RelatedIdentities, iden)
}

// ----------------------------------------------------------------------
// Methods RelatedIdentityType
// ----------------------------------------------------------------------

func (this *RelatedIdentityType) AddConfidence(c ConfidenceType) {
	this.Confidence = &c
}

func (this *RelatedIdentityType) AddInformationSource(source InformationSourceType) {
	this.InformationSource = &source
}

func (this *RelatedIdentityType) AddRelationshipDetail(name, ref, value string) {
	data := ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	this.Relationship = &data
}

func (this *RelatedIdentityType) AddIdentity(id IdentityType) {
	this.Identity = &id
}
