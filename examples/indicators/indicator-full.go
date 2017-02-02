// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/activeshadow/libstix/stix"
)

func main() {

	s := stix.New()
	i1 := s.NewIndicator()

	s.SetTimestampToNow()

	i1.AddIdRef("companyfoo:indicator-1234-1234-1234-1234")
	i1.AddVersion("2.0")
	i1.SetTimestampToNow()
	i1.SetNegate(false)
	i1.AddTitle("Some really neat indicator that we found")
	i1.AddType("URL Watchlist")
	i1.AddAlternativeID("CV-2014-12-12345")
	i1.AddAlternativeID("CV-2015-02-54321")
	i1.AddDescriptionText("", "Some long description")
	i1.AddShortDescriptionText("", "Some shorter description")
	i1.AddValidTimePosition("2015-01-01T00:00:00-0700", "2015-02-02T23:59:59-0700")
	i1.AddValidTimePosition("2014-01-01T00:00:00-0700", "2014-02-02T23:59:59-0700")

	source1 := stix.CreateInformationSource()
	source1.AddDescription("Some details about this source")
	source1.AddStartTime("2015-01-01T11:12:13-0700")
	source1.AddEndTime("2015-01-05T11:12:13-0700")
	source1.AddProducedTime("2015-01-05T11:12:13-0700")
	source1.AddReceivedTime("2015-01-05T11:12:13-0700")
	source1.AddRoleDetail("MyRoleVocab-1.0", "http://example.com/MyRoleVocab-1.0.txt", "CIO")

	identity1 := stix.CreateIdentity()
	identity1.CreateId()
	identity1.AddIdRef("companyfoo:identity-1234-1234-1234-1234")
	identity1.AddName("Lego Guy 1")
	source1.AddIdentity(identity1)

	// Create contributing source 1
	contribSource1 := stix.CreateInformationSource()
	contribSource1.AddDescription("Special source 1")
	identity2 := stix.CreateIdentity()
	identity2.AddName("George the Lego Guy")
	contribSource1.AddIdentity(identity2)

	// Create contributing source 2
	contribSource2 := stix.CreateInformationSource()
	contribSource2.AddDescription("Special source 2")
	identity3 := stix.CreateIdentity()
	identity3.AddName("Fred the Lego Guy")
	contribSource2.AddIdentity(identity3)

	source1.AddContributingSource(contribSource1)
	source1.AddContributingSource(contribSource2)

	source1.AddReference("http://foo.com/foo.txt")
	source1.AddReference("http://bar.com/bar.txt")

	con1 := stix.CreateConfidence()
	con1.CreateTimeStamp()
	con1.AddTimeStampPrecision("0.0")
	con1.AddValueDetail("MyConfidenceVocab-1.0", "http://example.com/MyConfidenceVocab-1.0.txt", "Super Sure")
	con1.AddDescription("We are super sure about this one")
	con1.AddSource(source1)

	con2 := stix.CreateConfidence()
	con2.CreateTimeStamp()
	con2.AddTimeStampPrecision("0.0")
	con2.AddValueDetail("MyConfidenceVocab-1.0", "http://example.com/MyConfidenceVocab-1.0.txt", "Super Sure")
	con2.AddDescription("These people are more sure")
	con2.AddSource(source1)
	con1.AddConfidenceAssertion(con2)
	i1.AddConfidence(con1)

	i1.AddProducer(source1)

	i2 := s.NewIndicator()
	i2.SetNegate(true)

	// fmt.Println(s.STIXPackage.Id)
	// fmt.Println(s.STIXPackage.Indicators[0].Title)
	// fmt.Println(s.STIXPackage.Indicators[0].Type[0].Value)
	// fmt.Println(s.STIXPackage.Indicators[0].Producer.Identity.Name)
	//fmt.Println(i)

	var data []byte
	data, _ = json.MarshalIndent(s, "", "    ")

	fmt.Println(string(data))

}
