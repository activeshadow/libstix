// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1

package stix

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/freestix/libstix/defs"
	// "github.com/freestix/libstix/indicator"
	// "github.com/freestix/libstix/ttp"
	"time"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type StixPackageType struct {
	StixMessage *StixMessageType `json:"stix_package,omitempty"`
}

type StixMessageType struct {
	Id        string `json:"id,omitempty"`
	IdRef     string `json:"idref,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Version   string `json:"version,omitempty"`
	// Indicators []indicator.IndicatorType `json:"indicators,omitempty"`
	// TTPs       *ttp.TTPsType             `json:"ttps,omitempty"`
}

// ----------------------------------------------------------------------
// Create Functions
// ----------------------------------------------------------------------

func New() StixPackageType {
	obj := CreateStixMessage()
	obj.CreateId()
	return obj
}

func CreateStixMessage() StixPackageType {
	var obj StixPackageType
	var msg StixMessageType
	obj.StixMessage = &msg
	return obj
}

// ----------------------------------------------------------------------
// Methods StixPackageType
// ----------------------------------------------------------------------

func (this *StixPackageType) CreateId() {
	this.StixMessage.Id = defs.COMPANY + ":package-" + uuid.New()
}

// This function will add an IDRef to the package
func (this *StixPackageType) AddIdRef(idref string) {
	this.StixMessage.IdRef = idref
}

func (this *StixPackageType) AddTimestamp(t string) {
	this.StixMessage.Timestamp = t
}

func (this *StixPackageType) SetTimestampToNow() {
	this.StixMessage.Timestamp = time.Now().Format(time.RFC3339)
}

// func (this *StixPackageType) AddIndicator(i indicator.IndicatorType) {
// 	if this.StixMessage.Indicators == nil {
// 		a := make([]indicator.IndicatorType, 0)
// 		this.StixMessage.Indicators = a
// 	}
// 	this.StixMessage.Indicators = append(this.StixMessage.Indicators, i)
// }

// func (this *StixPackageType) AddTTPs(t ttp.TTPsType) {
// 	this.StixMessage.TTPs = &t
// }
