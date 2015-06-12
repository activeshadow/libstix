// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

// RelatedPackageRefType identifies or characterizes a relationship to a Package.
type RelatedPackageRefType struct {
	GenericRelationshipType
	IdRef     string `json:"idref,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

// ----------------------------------------------------------------------
// Methods RelatedPackageRefType
// ----------------------------------------------------------------------

func (this *RelatedPackageRefType) AddConfidence(c ConfidenceType) {
	this.Confidence = &c
}

func (this *RelatedPackageRefType) AddInformationSource(i InformationSourceType) {
	this.InformationSource = &i
}

func (this *RelatedPackageRefType) AddRelationshipDetail(name, ref, value string) {
	data := ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	this.Relationship = &data
}

func (this *RelatedPackageRefType) AddIdRef(idref string) {
	this.IdRef = idref
}

func (this *RelatedPackageRefType) AddTimeStamp(t string) {
	// TODO Need to format the string in to ISO 8601 format or check that it is in the right format
	this.Timestamp = t
}
