package genopenapi

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/golang/protobuf/proto"
	pbdescriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	openapi_options "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapi/options"
)

func applyTemplate(file *descriptor.File, reg *descriptor.Registry) *openapi3.Swagger {
	return nil
}

//func overrideParameterWithOption(param *openapi3.Parameter, opt *Parame, field *descriptor.Field) error {
//	if opt.Name != "" {
//		param.Name = opt.Name
//	}
//	// Make sure that parameter isn't empty. If specifying a non-path parameter, there must be an in field.
//	if opt.In == swagger_options.Parameter_INVALID {
//		return fmt.Errorf("option parameter for field %s must contain a valid \"in\" field", field.GetName())
//	}
//	if opt.In == swagger_options.Parameter_PATH {
//		return fmt.Errorf("\"in\" field in option parameter for field %s must not equal \"path\"", field.GetName())
//	}
//	param.In = strings.ToLower(opt.In.String())
//	if opt.Description != "" {
//		param.Description = opt.Description
//	}
//	if opt.Required {
//		param.Required = opt.Required
//	}
//	// TODO: Implement
//	//  1. schema
//	//  2. two style/explode
//	//  3. example
//	//  4. examples
//	//  5. deprecated
//	//  6. extensions
//	//schema := opt.GetSchema()
//	//if schema != nil {
//	//	param.Type = schema.JsonSchema.GetType()
//	//}
//	//opt.GetXRef()
//	if opt.Style != "" {
//		param.Style = opt.Style
//	}
//	if opt.Explode {
//		param.Explode = opt.Explode
//	}
//	if opt.Deprecated {
//		param.Deprecated = opt.Deprecated
//	}
//	//if opt.Example != nil {
//	//	param.Example = opt.Example
//	//}
//	//if opt.Examples != nil {
//	//	param.Example = opt.Example
//	//}
//	return nil
//}
//
//extractParameterOptionFromFieldDescriptor extracts the message of type
//swagger_options.Parameter from a given proto method's descriptor.
func extractParameterOptionFromFieldDescriptor(meth *pbdescriptor.FieldDescriptorProto) (*openapi_options.Parameter, error) {
	if meth.Options == nil {
		return nil, nil
	}
	if !proto.HasExtension(meth.Options, openapi_options.E_Openapiv3Parameter) {
		return nil, nil
	}
	ext, err := proto.GetExtension(meth.Options, openapi_options.E_Openapiv3Parameter)
	if err != nil {
		return nil, err
	}
	opts, ok := ext.(*openapi_options.Parameter)
	if !ok {
		return nil, fmt.Errorf("extension is %T; want an Operation", ext)
	}
	return opts, nil
}

//paramOpt, err := extractParameterOptionFromFieldDescriptor(field.FieldDescriptorProto)
//if paramOpt == nil {
//	return nil, err
//}
//if paramOpt.In == swagger_options.Parameter_INVALID {
//	return nil, fmt.Errorf("option parameter for field %s must contain a valid \"in\" field", *field.Name)
//}
//if paramOpt.In == swagger_options.Parameter_PATH {
//	return nil, fmt.Errorf("\"in\" field in option parameter for field %s must not equal \"path\"", *field.Name)
//}
