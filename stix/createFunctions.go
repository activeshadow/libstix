// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1

package stix

import (
	"github.com/freestix/libstix/common"
)

// ----------------------------------------------------------------------
// Top level object - stix
// ----------------------------------------------------------------------

func New() STIXPackageType {
	var s STIXPackageType
	ns1 := make([]NamespaceType, 0)
	ns2 := make([]NamespaceType, 0)
	s.STIXPackage.Namespaces = ns1
	s.STIXPackage.SchemaLocations = ns2
	return s
}

// ----------------------------------------------------------------------
// Objects from stix.common
// ----------------------------------------------------------------------

func CreateConfidence() common.ConfidenceType {
	var c common.ConfidenceType
	return c
}

func CreateKillChain() common.KillChainType {
	var k common.KillChainType
	return k
}

func CreateIdentity() common.IdentityType {
	var i common.IdentityType
	return i
}

func CreateInformationSource() common.InformationSourceType {
	var s common.InformationSourceType
	return s
}

func CreateHandling() common.MarkingSpecificationType {
	var m common.MarkingSpecificationType
	return m
}

func CreateMarkingStructure() common.MarkingStructureType {
	var m common.MarkingStructureType
	return m
}

func CreateRelatedCampaign() common.RelatedCampaignReferenceType {
	var r common.RelatedCampaignReferenceType
	return r
}

func CreateRelatedIdentity() common.RelatedIdentityType {
	var r common.RelatedIdentityType
	return r
}

func CreateRelatedObservable() common.RelatedObservableType {
	var r common.RelatedObservableType
	return r
}

func CreateRelatedPackageRef() common.RelatedPackageRefType {
	var p common.RelatedPackageRefType
	return p
}

func CreateStatement() common.StatementType {
	var s common.StatementType
	return s
}
