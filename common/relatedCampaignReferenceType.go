// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

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

func (this *RelatedCampaignReferenceType) AddConfidence(c ConfidenceType) {
	this.Confidence = &c
}

func (this *RelatedCampaignReferenceType) AddInformationSource(i InformationSourceType) {
	this.InformationSource = &i
}

func (this *RelatedCampaignReferenceType) AddRelationshipDetail(name, ref, value string) {
	data := ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	this.Relationship = &data
}

// TODO Campaign CampaignReferenceType
