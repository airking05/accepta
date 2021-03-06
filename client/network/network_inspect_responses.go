// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/go-docker-api-client/models"
)

// NetworkInspectReader is a Reader for the NetworkInspect structure.
type NetworkInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNetworkInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewNetworkInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewNetworkInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNetworkInspectOK creates a NetworkInspectOK with default headers values
func NewNetworkInspectOK() *NetworkInspectOK {
	return &NetworkInspectOK{}
}

/*NetworkInspectOK handles this case with default header values.

No error
*/
type NetworkInspectOK struct {
	Payload *models.Network
}

func (o *NetworkInspectOK) Error() string {
	return fmt.Sprintf("[GET /networks/{id}][%d] networkInspectOK  %+v", 200, o.Payload)
}

func (o *NetworkInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Network)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkInspectNotFound creates a NetworkInspectNotFound with default headers values
func NewNetworkInspectNotFound() *NetworkInspectNotFound {
	return &NetworkInspectNotFound{}
}

/*NetworkInspectNotFound handles this case with default header values.

Network not found
*/
type NetworkInspectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *NetworkInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /networks/{id}][%d] networkInspectNotFound  %+v", 404, o.Payload)
}

func (o *NetworkInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkInspectInternalServerError creates a NetworkInspectInternalServerError with default headers values
func NewNetworkInspectInternalServerError() *NetworkInspectInternalServerError {
	return &NetworkInspectInternalServerError{}
}

/*NetworkInspectInternalServerError handles this case with default header values.

Server error
*/
type NetworkInspectInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *NetworkInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /networks/{id}][%d] networkInspectInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
