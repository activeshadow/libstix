package main

import (
	"encoding/json"
	"fmt"
	"github.com/jordan2175/freestix/libstix/indicator"
	"github.com/jordan2175/freestix/libstix/stix"
)

func main() {

	s := stix.New()

	i1 := indicator.New()
	i1.CreateId()
	i1.AddIdRef("companyfoo:indicator-1234-1234-1234-1234")
	i1.AddVersion("2.0")
	i1.CreateTimeStamp()
	i1.SetNegate(false)
	i1.AddTitle("Some really neat indicator that we found")
	i1.AddStandardType("URL Watchlist")
	i1.AddAlternativeID("CV-2014-12-12345")
	i1.AddAlternativeID("CV-2015-02-54321")
	i1.AddDescription("txt", "Some long description")
	i1.AddShortDescription("txt", "Some shorter description")

	s.AddIndicator(i1)

	fmt.Println("====================================")
	var data []byte
	data, _ = json.MarshalIndent(s, "", "    ")

	fmt.Println(string(data))

}
