package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
)

type Config struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services" jsonschema:"title=Services,description=Container definitions"`
}

type Service struct {
	Image string            `yaml:"image" jsonschema:"required,title=Image,description=The image to run,example=alpine:latest"`
	Ports []PortMapping     `yaml:"ports" jsonschema:"title=Ports,description=Ports to expose,example=3000:3000"`
	Env   map[string]string `yaml:"env"`
}

type PortMapping string

func (PortMapping) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "string",
		Title:       "Port mapping",
		Description: "Port to expose",
		Pattern:     "^[0-9]+:[0-9]+$",
		Examples:    []interface{}{"3000:3000"},
	}
}

func main() {
	r := &jsonschema.Reflector{
		PreferYAMLSchema:           true,
		RequiredFromJSONSchemaTags: true,
	}
	schema := r.Reflect(&Config{})

	schemaStr, err := json.Marshal(schema)
	if err != nil {
		panic(err)
	}

	var indented bytes.Buffer
	if err = json.Indent(&indented, schemaStr, "", "  "); err != nil {
		panic(err)
	}

	fmt.Println(indented.String())
}
