// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateRepositoryCommon = "UpdateRepository"

// UpdateRepositoryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateRepositoryCommon operation. The "output" return
// value will be populated with the UpdateRepositoryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateRepositoryCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateRepositoryCommon Send returns without error.
//
// See UpdateRepositoryCommon for more information on using the UpdateRepositoryCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateRepositoryCommonRequest method.
//    req, resp := client.UpdateRepositoryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) UpdateRepositoryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateRepositoryCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateRepositoryCommon API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation UpdateRepositoryCommon for usage and error information.
func (c *CR) UpdateRepositoryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateRepositoryCommonRequest(input)
	return out, req.Send()
}

// UpdateRepositoryCommonWithContext is the same as UpdateRepositoryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateRepositoryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) UpdateRepositoryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateRepositoryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateRepository = "UpdateRepository"

// UpdateRepositoryRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateRepository operation. The "output" return
// value will be populated with the UpdateRepositoryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateRepositoryCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateRepositoryCommon Send returns without error.
//
// See UpdateRepository for more information on using the UpdateRepository
// API call, and error handling.
//
//    // Example sending a request using the UpdateRepositoryRequest method.
//    req, resp := client.UpdateRepositoryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) UpdateRepositoryRequest(input *UpdateRepositoryInput) (req *request.Request, output *UpdateRepositoryOutput) {
	op := &request.Operation{
		Name:       opUpdateRepository,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRepositoryInput{}
	}

	output = &UpdateRepositoryOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateRepository API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation UpdateRepository for usage and error information.
func (c *CR) UpdateRepository(input *UpdateRepositoryInput) (*UpdateRepositoryOutput, error) {
	req, out := c.UpdateRepositoryRequest(input)
	return out, req.Send()
}

// UpdateRepositoryWithContext is the same as UpdateRepository with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateRepository for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) UpdateRepositoryWithContext(ctx volcengine.Context, input *UpdateRepositoryInput, opts ...request.Option) (*UpdateRepositoryOutput, error) {
	req, out := c.UpdateRepositoryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateRepositoryInput struct {
	_ struct{} `type:"structure"`

	AccessLevel *string `type:"string"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	Name *string `type:"string"`

	Namespace *string `type:"string"`

	Registry *string `type:"string"`
}

// String returns the string representation
func (s UpdateRepositoryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateRepositoryInput) GoString() string {
	return s.String()
}

// SetAccessLevel sets the AccessLevel field's value.
func (s *UpdateRepositoryInput) SetAccessLevel(v string) *UpdateRepositoryInput {
	s.AccessLevel = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *UpdateRepositoryInput) SetClientToken(v string) *UpdateRepositoryInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UpdateRepositoryInput) SetDescription(v string) *UpdateRepositoryInput {
	s.Description = &v
	return s
}

// SetName sets the Name field's value.
func (s *UpdateRepositoryInput) SetName(v string) *UpdateRepositoryInput {
	s.Name = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *UpdateRepositoryInput) SetNamespace(v string) *UpdateRepositoryInput {
	s.Namespace = &v
	return s
}

// SetRegistry sets the Registry field's value.
func (s *UpdateRepositoryInput) SetRegistry(v string) *UpdateRepositoryInput {
	s.Registry = &v
	return s
}

type UpdateRepositoryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateRepositoryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateRepositoryOutput) GoString() string {
	return s.String()
}
