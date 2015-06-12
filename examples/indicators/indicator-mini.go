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
	s := stix.New()
	i1 := s.NewIndicator()

	i1.SetTimestampToNow()
	i1.AddTitle("Attack 2015-02")
	i1.AddType("IP Watchlist")
	observable_i1 := i1.NewObservable()
	properties_1 := observable_i1.GetObjectProperties()

	properties_1.AddType("URL")
	properties_1.AddEqualsUriValue("http://foo.com")
	properties_1.AddEqualsUriValue("http://bar.com")
	properties_1.AddEqualsUriValue("http://fooandbar.com")

	var data []byte
	data, _ = json.MarshalIndent(s, "", "    ")

	fmt.Println(string(data))
}
