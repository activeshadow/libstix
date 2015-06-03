package main

import (
	"encoding/json"
	"fmt"
	"github.com/jordan2175/freestix/libstix/indicator"
	"github.com/jordan2175/freestix/libstix/stix"
)

func main() {
	s := stix.Create()
	i1 := indicator.Create()
	i1.CreateId()
	i1.CreateTimeStamp()
	i1.AddTitle("Attack 2015-02")
	i1.AddType("IP Watchlist")
	i1.AddType("URL Watchlist")

	h := stix.CreateHandling()
	h.AddControlledStructure("Foobar")
	h.AddTLPMarking("Red")

	i1.AddHandling(h)

	s.AddIndicator(i1)

	var data []byte
	data, _ = json.MarshalIndent(s, "", "    ")

	fmt.Println(string(data))
}
