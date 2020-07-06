package genopenapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"

	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"

	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/genswagger"
)

type generator struct {
	SwaggerGenerator *genswagger.SwaggerGenerator
	reg              *descriptor.Registry
}

// New returns a new generator which generates grpc gateway files.
func New(reg *descriptor.Registry) *generator {
	return &generator{
		SwaggerGenerator: genswagger.New(reg),
		reg:              reg,
	}
}

func (g *generator) Generate(targets []*descriptor.File) ([]*plugin.CodeGeneratorResponse_File, error) {
	files, err := g.SwaggerGenerator.Generate(targets)
	if err != nil {
		return nil, err
	}
	// TODO: Convert Swagger to OpenAPI and run `applyTemplateV3`
	for i, file := range files {
		swaggerV3, err := swaggerV2ToOpenAPI([]byte(file.GetContent()))
		if err != nil {
			return nil, err
		}
		file, err = encodeOpenAPI(file.GetName(), swaggerV3)
		if err != nil {
			return nil, err
		}
		// Overwrite v2 file with v3 file.
		files[i] = file
	}
	//fmt.Println(files[0].Content)
	return files, nil
}

func swaggerV2ToOpenAPI(swaggerV2Data []byte) (*openapi3.Swagger, error) {
	swaggerV2 := &openapi2.Swagger{}
	err  := json.Unmarshal(swaggerV2Data, swaggerV2)
	if err != nil {
		return nil, err
	}
	swaggerV3, err := openapi2conv.ToV3Swagger(swaggerV2)
	if err != nil {
		return nil, err
	}
	return swaggerV3, nil
}

// encodeOpenAPI converts OpenAPI file obj to plugin.CodeGeneratorResponse_File
func encodeOpenAPI(filename string, swaggerV3 *openapi3.Swagger) (*plugin.CodeGeneratorResponse_File, error) {
	var formatted bytes.Buffer
	enc := json.NewEncoder(&formatted)
	enc.SetIndent("", "  ")
	if err := enc.Encode(swaggerV3); err != nil {
		return nil, err
	}
	ext := filepath.Ext(filename)
	base := strings.TrimSuffix(filename, ext)
	output := fmt.Sprintf("%s_v3.json", base)
	return &plugin.CodeGeneratorResponse_File{
		Name:    proto.String(output),
		Content: proto.String(formatted.String()),
	}, nil
}
