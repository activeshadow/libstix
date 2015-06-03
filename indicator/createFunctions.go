// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1

package indicator

// ----------------------------------------------------------------------
// Top level object - stix.indicator
// ----------------------------------------------------------------------

func New() IndicatorType {
	var i IndicatorType
	return i
}

// ----------------------------------------------------------------------
// Objects from stix.indicator
// ----------------------------------------------------------------------

func CreateRelatedIndicator() RelatedIndicatorType {
	var r RelatedIndicatorType
	return r
}

func CreateRelatedIndicators() RelatedIndicatorsType {
	var r RelatedIndicatorsType
	return r
}

func CreateRelatedObservables() RelatedObservablesType {
	var r RelatedObservablesType
	return r
}

func CreateSighting() SightingType {
	var s SightingType
	return s
}

func CreateSightings() SightingsType {
	var s SightingsType
	return s
}
