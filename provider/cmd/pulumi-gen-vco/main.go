package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	vco "github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider"
	providerVersion "github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/version"
)

func MarshalIndent(v any) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "    ")
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func main() {
	flag.Usage = func() {
		const usageFormat = "Usage: %s <schema-file>"
		fmt.Fprintf(flag.CommandLine.Output(), usageFormat, os.Args[0])
		flag.PrintDefaults()
	}

	var version string
	flag.StringVar(&version, "version", providerVersion.Version, "the provider version to record in the generated code")

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}
	s, err := vco.Schema(version)
	if err != nil {
		panic(err)
	}

	var arg map[string]any
	err = json.Unmarshal([]byte(s), &arg)
	if err != nil {
		panic(err)
	}

	delete(arg, "version")

	out, err := MarshalIndent(arg)
	if err != nil {
		panic(err)
	}

	schemaPath := args[0]
	err = os.WriteFile(schemaPath, out, 0600)
	if err != nil {
		panic(err)
	}
}
