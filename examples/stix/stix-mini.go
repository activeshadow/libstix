// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/freestix/libstix/stix"
)

func main() {
	sm := stix.New()
	sm.AddIdRef("company1:package-8ef29ed2-0e87-4dfa-b5c4-00788484be09")
	sm.SetTimestampToNow()

	var data []byte
	data, _ = json.MarshalIndent(sm, "", "    ")

	fmt.Println(string(data))
}
