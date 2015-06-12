// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/freestix/libstix/indicator"
	"github.com/freestix/libstix/stix"
)

func main() {

	s := stix.New()

	i1 := indicator.New()
	i1.AddIdRef("companyfoo:indicator-1234-1234-1234-1234")
	i1.AddVersion("2.0.alpha")
	i1.SetTimestampToNow()
	i1.SetNegate(false)
	i1.AddTitle("Some really neat indicator that we found")
	i1.AddType("URL Watchlist")
	i1.AddAlternativeID("CV-2014-12-12345")
	i1.AddAlternativeID("CV-2015-02-54321")
	i1.AddDescriptionText("", "Some long description")
	i1.AddShortDescriptionText("", "Some shorter description")

	s.AddIndicator(i1)

	fmt.Println("====================================")
	var data []byte
	data, _ = json.MarshalIndent(s, "", "    ")

	fmt.Println(string(data))

}
