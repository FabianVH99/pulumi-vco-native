package main

import (
	"fmt"
	vco "github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/provider"
	"github.com/fabianv-cloud/pulumi-vco-native/provider/pkg/version"
	p "github.com/pulumi/pulumi-go-provider"
	"os"
	"strings"
)

func main() {
	version := version.Version
	if strings.HasPrefix(version, "v") {
		version = version[1:]
	}

	vcoProvider := vco.NewProvider()

	err := p.RunProvider("vco", version, vcoProvider)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
