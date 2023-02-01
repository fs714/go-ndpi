package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PackageName = "types"
const NdpiProtocolPrefix = "NDPI_PROTOCOL_"

var ndpi_protocol_ids_header_path string

func getProtocolMap(path string) (protocolVarList []string, protocolIdList []int, protocolNameList []string, err error) {
	protocolVarList = make([]string, 0)
	protocolIdList = make([]int, 0)
	protocolNameList = make([]string, 0)

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed to open %s with err: %s\n", path, err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, NdpiProtocolPrefix) {
			ls := strings.Split(line, "=")
			varName := strings.TrimSpace(ls[0])
			name := strings.TrimPrefix(varName, NdpiProtocolPrefix)
			idString := strings.TrimSpace(strings.Split(strings.TrimSpace(ls[1]), ",")[0])

			id, eerr := strconv.Atoi(idString)
			if eerr != nil {
				fmt.Printf("failed to parse id for line: %s\n", line)
				continue
			}

			protocolVarList = append(protocolVarList, varName)
			protocolIdList = append(protocolIdList, id)
			protocolNameList = append(protocolNameList, name)
		}
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("failed to scan file with err: %s\n", err.Error())
		return
	}

	return
}

func generateGolangCode(protocolVarList []string, protocolIdList []int, protocolNameList []string) (code string) {
	var builder strings.Builder

	builder.WriteString("package " + PackageName + "\n\n")
	builder.WriteString("type NdpiProtocol uint16\n\n")

	// func (p *NdpiProtocol) ToName() string {
	// 	name, ok := NdpiProtocolNameMap[*p]
	// 	if !ok {
	// 		name = ""
	// 	}

	// 	return name
	// }

	builder.WriteString("func (p *NdpiProtocol) ToName() string {\n")
	builder.WriteString("    name, ok := NdpiProtocolNameMap[*p]\n")
	builder.WriteString("    if !ok {\n")
	builder.WriteString(`        name = ""` + "\n")
	builder.WriteString("    }\n\n")
	builder.WriteString("    return name\n")
	builder.WriteString("}\n\n")

	builder.WriteString("const (\n")
	for idx, v := range protocolVarList {
		builder.WriteString("    " + v + " NdpiProtocol = " + strconv.Itoa(protocolIdList[idx]) + "\n")
	}
	builder.WriteString(")\n\n")

	builder.WriteString("const (\n")
	for _, v := range protocolNameList {
		builder.WriteString("    PROTO_" + v + ` string = "` + v + `"` + "\n")
	}
	builder.WriteString(")\n\n")

	builder.WriteString("var NdpiProtocolNameMap = map[NdpiProtocol]string{\n")
	for idx, v := range protocolVarList {
		builder.WriteString("    " + v + ": PROTO_" + protocolNameList[idx] + ",\n")
	}
	builder.WriteString("}\n\n")

	code = builder.String()

	return
}

func main() {
	flag.StringVar(&ndpi_protocol_ids_header_path, "p", "/usr/include/ndpi/ndpi_protocol_ids.h", "path of ndpi_protocol_ids.h")
	flag.Parse()

	protocolVarList, protocolIdList, protocolNamList, err := getProtocolMap(ndpi_protocol_ids_header_path)
	if err != nil {
		return
	}

	code := generateGolangCode(protocolVarList, protocolIdList, protocolNamList)

	fmt.Println(code)
}
