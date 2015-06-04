package main

import (
	"encoding/json"
	"fmt"
	"github.com/freestix/cybox"
	"github.com/freestix/libstix/indicator"
	"github.com/freestix/libstix/stix"
)

func main() {
	s := stix.Create()
	i1 := indicator.Create()
	i1.CreateId()
	i1.CreateTimeStamp()
	i1.AddTitle("Attack 2015-02")
	i1.AddType("IP Watchlist")

	o := cybox.CreateObservable()
	o.CreateId()

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
