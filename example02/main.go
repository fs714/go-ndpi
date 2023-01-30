/*
Example for ndpi protocol bitmask
- NewNdpiProtocolBitmask
- NdpiProtocolBitmaskAdd
- NdpiProtocolBitmaskDel
- NdpiProtocolBitmaskIsSet
- NdpiProtocolBitmaskReset
- NdpiProtocolBitmaskSetAll
- GetProtoDefaults
*/

package main

import (
	"encoding/json"
	"fmt"

	"github.com/fs714/go-ndpi/gondpi"
	"github.com/fs714/go-ndpi/gondpi/types"
)

func main() {
	detectionBitmask := gondpi.NewNdpiProtocolBitmask()
	fmt.Printf("new: %v\n", detectionBitmask)

	detectionBitmask = gondpi.NdpiProtocolBitmaskReset(detectionBitmask)
	fmt.Printf("reset: %v\n", detectionBitmask)

	fmt.Printf("is ssh enabled before add ssh: %t\n", gondpi.NdpiProtocolBitmaskIsSet(detectionBitmask, types.NDPI_PROTOCOL_SSH))
	detectionBitmask = gondpi.NdpiProtocolBitmaskAdd(detectionBitmask, types.NDPI_PROTOCOL_SSH)
	fmt.Printf("add ssh: %v\n", detectionBitmask)
	fmt.Printf("is ssh enabled after add ssh: %t\n", gondpi.NdpiProtocolBitmaskIsSet(detectionBitmask, types.NDPI_PROTOCOL_SSH))

	detectionBitmask = gondpi.NdpiProtocolBitmaskSetAll(detectionBitmask)
	fmt.Printf("set all: %v\n", detectionBitmask)

	detectionBitmask = gondpi.NdpiProtocolBitmaskDel(detectionBitmask, types.NDPI_PROTOCOL_SSH)
	fmt.Printf("del ssh: %v\n", detectionBitmask)

	ndpiDM, err := gondpi.NdpiDetectionModuleInitialize(types.NDPI_NO_PREFS, detectionBitmask)
	if err != nil {
		fmt.Printf("failed to initialize ndpi detection module with err: %s\n", err.Error())
		return
	}

	npds := ndpiDM.GetProtoDefaults()
	npdsJson, err := json.MarshalIndent(npds, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal ndpi protocol defaults with err: %s\n", err.Error())
		return
	}

	fmt.Println(len(npds))
	fmt.Println(string(npdsJson))

	gondpi.NdpiDetectionModuleExit(ndpiDM)
}
