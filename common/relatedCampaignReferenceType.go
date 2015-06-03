// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1

package common

// RelatedCampaignReferenceType identifies or characterizes a relationship by
// reference to a campaign.
type RelatedCampaignReferenceType struct {
	GenericRelationshipType
}

type CampaignReferenceType struct {
	IdRef     string                           `json:"idref,omitempty"`
	Timestamp string                           `json:"timestamp,omitempty"`
	Names     []ControlledVocabularyStringType `json:"names,omitempty"`
}

// ----------------------------------------------------------------------
// Methods RelatedCampaignReferenceType
// ----------------------------------------------------------------------

// AddConfidence adds a level of confidence in the assertion of the relationship
// between the two components.
//
// This supports 0..1 objects
func (r *RelatedCampaignReferenceType) AddConfidence(c ConfidenceType) {
	r.Confidence = &c
}

// AddInformationSource adds the source of the information about the
// relationship between the two components.
//
// This supports 0..1 objects
func (r *RelatedCampaignReferenceType) AddInformationSource(i InformationSourceType) {
	r.InformationSource = &i
}

// AddRelationshipDetail adds the type of the relationship between the two
// components.
//
// This field is XML is implemented through the xsi:type controlled vocabulary
// extension mechanism. No default vocabulary type has been defined for
// STIX 1.1.1. Users may either define their own vocabulary using the type
// extension mechanism, specify a vocabulary name and reference using the
// attributes, or simply use this as a free string field.
//
// This supports 0..1 objects
func (r *RelatedCampaignReferenceType) AddRelationshipDetail(name, ref, value string) {
	data := ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	r.Relationship = &data
}

// TODO Campaign CampaignReferenceType
