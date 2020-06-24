package genopenapi

import (
	"fmt"
	"os"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"

	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/genswagger"
)

type generator struct {
	SwaggerGenerator *genswagger.SwaggerGenerator
	reg              *descriptor.Registry
}

func (g *generator) Generate(targets []*descriptor.File) ([]*plugin.CodeGeneratorResponse_File, error) {
	f, err := g.SwaggerGenerator.Generate(targets)
	if err != nil {
		return nil, err
	}
	// TODO: Convert Swagger to OpenAPI and run `applyTemplateV3`
	fmt.Fprint(os.Stderr, "hi")
	return f, nil
}
