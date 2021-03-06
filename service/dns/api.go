// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dns

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opCreateHostedZone = "CreateHostedZone"

// CreateHostedZoneRequest generates a "ksc/request.Request" representing the
// client's request for the CreateHostedZone operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See CreateHostedZone for more information on using the CreateHostedZone
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the CreateHostedZoneRequest method.
//    req, resp := client.CreateHostedZoneRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/CreateHostedZone
func (c *Dns) CreateHostedZoneRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateHostedZone,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// CreateHostedZone API operation for dns.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for dns's
// API operation CreateHostedZone for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/CreateHostedZone
func (c *Dns) CreateHostedZone(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateHostedZoneRequest(input)
	return out, req.Send()
}

// CreateHostedZoneWithContext is the same as CreateHostedZone with the addition of
// the ability to pass a context and additional request options.
//
// See CreateHostedZone for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Dns) CreateHostedZoneWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateHostedZoneRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateResourceRecord = "CreateResourceRecord"

// CreateResourceRecordRequest generates a "ksc/request.Request" representing the
// client's request for the CreateResourceRecord operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See CreateResourceRecord for more information on using the CreateResourceRecord
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the CreateResourceRecordRequest method.
//    req, resp := client.CreateResourceRecordRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/CreateResourceRecord
func (c *Dns) CreateResourceRecordRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateResourceRecord,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// CreateResourceRecord API operation for dns.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for dns's
// API operation CreateResourceRecord for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/CreateResourceRecord
func (c *Dns) CreateResourceRecord(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateResourceRecordRequest(input)
	return out, req.Send()
}

// CreateResourceRecordWithContext is the same as CreateResourceRecord with the addition of
// the ability to pass a context and additional request options.
//
// See CreateResourceRecord for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Dns) CreateResourceRecordWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateResourceRecordRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteHostedZone = "DeleteHostedZone"

// DeleteHostedZoneRequest generates a "ksc/request.Request" representing the
// client's request for the DeleteHostedZone operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DeleteHostedZone for more information on using the DeleteHostedZone
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DeleteHostedZoneRequest method.
//    req, resp := client.DeleteHostedZoneRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/DeleteHostedZone
func (c *Dns) DeleteHostedZoneRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteHostedZone,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteHostedZone API operation for dns.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for dns's
// API operation DeleteHostedZone for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/DeleteHostedZone
func (c *Dns) DeleteHostedZone(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteHostedZoneRequest(input)
	return out, req.Send()
}

// DeleteHostedZoneWithContext is the same as DeleteHostedZone with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteHostedZone for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Dns) DeleteHostedZoneWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteHostedZoneRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteResourceRecord = "DeleteResourceRecord"

// DeleteResourceRecordRequest generates a "ksc/request.Request" representing the
// client's request for the DeleteResourceRecord operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DeleteResourceRecord for more information on using the DeleteResourceRecord
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DeleteResourceRecordRequest method.
//    req, resp := client.DeleteResourceRecordRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/DeleteResourceRecord
func (c *Dns) DeleteResourceRecordRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteResourceRecord,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteResourceRecord API operation for dns.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for dns's
// API operation DeleteResourceRecord for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/DeleteResourceRecord
func (c *Dns) DeleteResourceRecord(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteResourceRecordRequest(input)
	return out, req.Send()
}

// DeleteResourceRecordWithContext is the same as DeleteResourceRecord with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteResourceRecord for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Dns) DeleteResourceRecordWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteResourceRecordRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeHostedZones = "DescribeHostedZones"

// DescribeHostedZonesRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeHostedZones operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeHostedZones for more information on using the DescribeHostedZones
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeHostedZonesRequest method.
//    req, resp := client.DescribeHostedZonesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/DescribeHostedZones
func (c *Dns) DescribeHostedZonesRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeHostedZones,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeHostedZones API operation for dns.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for dns's
// API operation DescribeHostedZones for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/DescribeHostedZones
func (c *Dns) DescribeHostedZones(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeHostedZonesRequest(input)
	return out, req.Send()
}

// DescribeHostedZonesWithContext is the same as DescribeHostedZones with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeHostedZones for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Dns) DescribeHostedZonesWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeHostedZonesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeResourceRecords = "DescribeResourceRecords"

// DescribeResourceRecordsRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeResourceRecords operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeResourceRecords for more information on using the DescribeResourceRecords
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeResourceRecordsRequest method.
//    req, resp := client.DescribeResourceRecordsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/DescribeResourceRecords
func (c *Dns) DescribeResourceRecordsRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeResourceRecords,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeResourceRecords API operation for dns.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for dns's
// API operation DescribeResourceRecords for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/DescribeResourceRecords
func (c *Dns) DescribeResourceRecords(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeResourceRecordsRequest(input)
	return out, req.Send()
}

// DescribeResourceRecordsWithContext is the same as DescribeResourceRecords with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeResourceRecords for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Dns) DescribeResourceRecordsWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeResourceRecordsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetGeolocations = "GetGeolocations"

// GetGeolocationsRequest generates a "ksc/request.Request" representing the
// client's request for the GetGeolocations operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetGeolocations for more information on using the GetGeolocations
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetGeolocationsRequest method.
//    req, resp := client.GetGeolocationsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/GetGeolocations
func (c *Dns) GetGeolocationsRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetGeolocations,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// GetGeolocations API operation for dns.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for dns's
// API operation GetGeolocations for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/GetGeolocations
func (c *Dns) GetGeolocations(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetGeolocationsRequest(input)
	return out, req.Send()
}

// GetGeolocationsWithContext is the same as GetGeolocations with the addition of
// the ability to pass a context and additional request options.
//
// See GetGeolocations for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Dns) GetGeolocationsWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetGeolocationsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyResourceRecord = "ModifyResourceRecord"

// ModifyResourceRecordRequest generates a "ksc/request.Request" representing the
// client's request for the ModifyResourceRecord operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ModifyResourceRecord for more information on using the ModifyResourceRecord
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ModifyResourceRecordRequest method.
//    req, resp := client.ModifyResourceRecordRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/ModifyResourceRecord
func (c *Dns) ModifyResourceRecordRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyResourceRecord,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyResourceRecord API operation for dns.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for dns's
// API operation ModifyResourceRecord for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/ModifyResourceRecord
func (c *Dns) ModifyResourceRecord(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyResourceRecordRequest(input)
	return out, req.Send()
}

// ModifyResourceRecordWithContext is the same as ModifyResourceRecord with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyResourceRecord for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Dns) ModifyResourceRecordWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyResourceRecordRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opResourceRecordChangeType = "ResourceRecordChangeType"

// ResourceRecordChangeTypeRequest generates a "ksc/request.Request" representing the
// client's request for the ResourceRecordChangeType operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ResourceRecordChangeType for more information on using the ResourceRecordChangeType
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ResourceRecordChangeTypeRequest method.
//    req, resp := client.ResourceRecordChangeTypeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/ResourceRecordChangeType
func (c *Dns) ResourceRecordChangeTypeRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opResourceRecordChangeType,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ResourceRecordChangeType API operation for dns.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for dns's
// API operation ResourceRecordChangeType for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/dns-2016-06-07/ResourceRecordChangeType
func (c *Dns) ResourceRecordChangeType(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ResourceRecordChangeTypeRequest(input)
	return out, req.Send()
}

// ResourceRecordChangeTypeWithContext is the same as ResourceRecordChangeType with the addition of
// the ability to pass a context and additional request options.
//
// See ResourceRecordChangeType for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Dns) ResourceRecordChangeTypeWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ResourceRecordChangeTypeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}
