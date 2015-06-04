// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.2

package common

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/freestix/libstix/defs"
)

type IdentityType struct {
	Id                string                `json:"id,omitempty"`
	IdRef             string                `json:"idref,omitempty"`
	Name              string                `json:"name,omitempty"`
	RelatedIdentities []RelatedIdentityType `json:"relatedIdentites,omitempty"`
}

type RelatedIdentityType struct {
	Confidence        *ConfidenceType                 `json:"confidence,omitempty"`
	InformationSource *InformationSourceType          `json:"informationSource,omitempty"`
	Relationship      *ControlledVocabularyStringType `json:"relationship,omitempty"`
	Identity          *IdentityType                   `json:"identity,omitempty"`
}

// ----------------------------------------------------------------------
// Methods IdentityType
// ----------------------------------------------------------------------

func (i *IdentityType) CreateId() {
	i.Id = defs.COMPANY + ":identity-" + uuid.New()
}

func (i *IdentityType) AddIdRef(idref string) {
	i.IdRef = idref
}

func (i *IdentityType) AddName(name string) {
	i.Name = name
}

func (i *IdentityType) AddRelatedIdentity(iden RelatedIdentityType) {
	if i.RelatedIdentities == nil {
		a := make([]RelatedIdentityType, 0)
		i.RelatedIdentities = a
	}
	i.RelatedIdentities = append(i.RelatedIdentities, iden)
}

// ----------------------------------------------------------------------
// Methods RelatedIdentityType
// ----------------------------------------------------------------------

func (r *RelatedIdentityType) AddConfidence(c ConfidenceType) {
	r.Confidence = &c
}

func (r *RelatedIdentityType) AddInformationSource(source InformationSourceType) {
	r.InformationSource = &source
}

func (r *RelatedIdentityType) AddRelationshipDetail(name, ref, value string) {
	data := ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	r.Relationship = &data
}

func (r *RelatedIdentityType) AddIdentity(id IdentityType) {
	r.Identity = &id
}
