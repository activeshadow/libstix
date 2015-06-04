package main

import (
	"encoding/json"
	"fmt"
	"github.com/freestix/libstix/stix"
)

func main() {
	sm := stix.New()
	sm.SetTimestampToNow()

	var data []byte
	data, _ = json.MarshalIndent(sm, "", "    ")

	fmt.Println(string(data))
}
