// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package stix

import (
	"github.com/freestix/libstix/common"
)

// ----------------------------------------------------------------------
// Objects from stix.common
// ----------------------------------------------------------------------

func CreateConfidence() common.ConfidenceType {
	var obj common.ConfidenceType
	return obj
}

func CreateKillChain() common.KillChainType {
	var obj common.KillChainType
	return obj
}

func CreateIdentity() common.IdentityType {
	var obj common.IdentityType
	return obj
}

func CreateInformationSource() common.InformationSourceType {
	var obj common.InformationSourceType
	return obj
}

func CreateHandling() common.MarkingSpecificationType {
	var obj common.MarkingSpecificationType
	return obj
}

func CreateMarkingStructure() common.MarkingStructureType {
	var obj common.MarkingStructureType
	return obj
}

func CreateRelatedCampaign() common.RelatedCampaignReferenceType {
	var obj common.RelatedCampaignReferenceType
	return obj
}

func CreateRelatedIdentity() common.RelatedIdentityType {
	var obj common.RelatedIdentityType
	return obj
}

func CreateRelatedObservable() common.RelatedObservableType {
	var obj common.RelatedObservableType
	return obj
}

func CreateRelatedPackageRef() common.RelatedPackageRefType {
	var obj common.RelatedPackageRefType
	return obj
}

func CreateStatement() common.StatementType {
	var obj common.StatementType
	return obj
}
