// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1

package stix

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/jordan2175/freestix/libstix/defs"
	"github.com/jordan2175/freestix/libstix/indicator"
	"github.com/jordan2175/freestix/libstix/ttp"
)

// I needed this for JSON so that you could get an outer container called stixPackage
type STIXPackageType struct {
	STIXPackage STIXType `json:"stixPackage"`
}

type NamespaceType struct {
	Name string `json:"name,omitempty"`
	Ref  string `json:"ref,omitempty"`
}

type STIXType struct {
	Namespaces      []NamespaceType           `json:"namespaces,omitempty"`
	SchemaLocations []NamespaceType           `json:"schemaLocations,omitempty"`
	Id              string                    `json:"id,omitempty"`
	IdRef           string                    `json:"idref,omitempty"`
	Timestamp       string                    `json:"timestamp,omitempty"`
	Version         string                    `json:"version,omitempty"`
	Indicators      []indicator.IndicatorType `json:"indicators,omitempty"`
	TTPs            *ttp.TTPsType             `json:"ttps,omitempty"`
}

// ----------------------------------------------------------------------
// Methods
// ----------------------------------------------------------------------

func (s *STIXPackageType) AddNamespace(name, ref string) {
	ns := NamespaceType{name, ref}
	s.STIXPackage.Namespaces = append(s.STIXPackage.Namespaces, ns)
}

func (s *STIXPackageType) AddSchemaLocation(name, ref string) {
	ns := NamespaceType{name, ref}
	s.STIXPackage.SchemaLocations = append(s.STIXPackage.SchemaLocations, ns)
}

func (s *STIXPackageType) AddIndicator(i indicator.IndicatorType) {
	if s.STIXPackage.Indicators == nil {
		a := make([]indicator.IndicatorType, 0)
		s.STIXPackage.Indicators = a
	}
	s.STIXPackage.Indicators = append(s.STIXPackage.Indicators, i)
}

func (s *STIXPackageType) AddTTPs(t ttp.TTPsType) {
	s.STIXPackage.TTPs = &t
}

func (s *STIXPackageType) CreateId() {
	s.STIXPackage.Id = defs.COMPANY + ":package-" + uuid.New()
}

// This function will add an IDRef to the package
func (s *STIXPackageType) AddIdRef(idref string) {
	s.STIXPackage.IdRef = idref
}

func (s *STIXPackageType) AddCurrentTime() {
	s.STIXPackage.Timestamp = "2015"
}
