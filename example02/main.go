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

	gondpi "github.com/fs714/go-ndpi"
	"github.com/fs714/go-ndpi/types"
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

	ndpiHandle, err := gondpi.NdpiHandleInitialize(detectionBitmask)
	if err != nil {
		fmt.Printf("failed to initialize NdpiHandle with err: %s\n", err.Error())
		return
	}

	npds := ndpiHandle.GetProtoDefaults()
	npdsJson, err := json.MarshalIndent(npds, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal ndpi protocol defaults with err: %s\n", err.Error())
		return
	}

	fmt.Println(len(npds))
	fmt.Println(string(npdsJson))

	gondpi.NdpiHandleExit(ndpiHandle)
}
