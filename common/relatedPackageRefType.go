// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.3

package common

// RelatedPackageRefType identifies or characterizes a relationship to a Package.
type RelatedPackageRefType struct {
	GenericRelationshipType
	IdRef     string `json:"idref,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

// ----------------------------------------------------------------------
// Methods RelatedPackageRefType
// ----------------------------------------------------------------------

// AddConfidence adds a level of confidence in the assertion of the relationship
// between the two components.
//
// This supports 0..1 objects
func (r *RelatedPackageRefType) AddConfidence(c ConfidenceType) {
	r.Confidence = &c
}

// AddInformationSource adds the source of the information about the
// relationship between the two components.
//
// This supports 0..1 objects
func (r *RelatedPackageRefType) AddInformationSource(i InformationSourceType) {
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
func (r *RelatedPackageRefType) AddRelationshipDetail(name, ref, value string) {
	data := ControlledVocabularyStringType{
		VocabName:      name,
		VocabReference: ref,
		Value:          value,
	}
	r.Relationship = &data
}

// AddIdRef adds a specific ID value to the IDREF field. The format should be
// as follows:
//   example.com:indicator-uuid
//
// This field specifies a globally unique identifier of a STIX Package specified
// elsewhere.
//
// This is optional
func (r *RelatedPackageRefType) AddIdRef(idref string) {
	r.IdRef = idref
}

// AddTimeStamp will add a specific timestamp to the timestamp field for the
// definition of a specific version of an Indicator. The timestamp SHOULD be in
// ISO8601 format and include a timezone, example:
//   yyyy-mm-ddThh:mm:ss+-hh:mm
//
// When used in conjunction with the idref, this field may be used to reference
// a specific version of a STIX Package defined elsewhere.
//
// This is optional
func (r *RelatedPackageRefType) AddTimeStamp(t string) {
	// TODO Need to format the string in to ISO 8601 format or check that it is in the right format
	r.Timestamp = t
}
