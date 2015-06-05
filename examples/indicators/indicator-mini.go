// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/freestix/libcybox/cybox"
	"github.com/freestix/libstix/indicator"
	"github.com/freestix/libstix/stix"
)

func main() {
	s := stix.New()
	i1 := indicator.New()
	i1.SetTimestampToNow()
	i1.AddTitle("Attack 2015-02")
	i1.AddType("IP Watchlist")

	o := cybox.New()

	obj := cybox.CreateObject("UriObj:UriObjectType", "URL")
	obj.CreateId()

	uriObj := cybox.CreateUriObject()
	uriObj.CreateId()
	uriObj.AddUriObject("http://foo.com")
	uriObj.AddUriObject("http://bar.com")
	uriObj.AddUriObject("http://fooandbar.com")

	obj.AddUriObject(uriObj)

	o.AddObject(obj)

	i1.AddObservable(o)
	s.AddIndicator(i1)

	var data []byte
	data, _ = json.MarshalIndent(s, "", "    ")

	fmt.Println(string(data))
}
