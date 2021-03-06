// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/go-docker-api-client/models"
)

// ServiceInspectReader is a Reader for the ServiceInspect structure.
type ServiceInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServiceInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewServiceInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewServiceInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewServiceInspectServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceInspectOK creates a ServiceInspectOK with default headers values
func NewServiceInspectOK() *ServiceInspectOK {
	return &ServiceInspectOK{}
}

/*ServiceInspectOK handles this case with default header values.

no error
*/
type ServiceInspectOK struct {
	Payload *models.Service
}

func (o *ServiceInspectOK) Error() string {
	return fmt.Sprintf("[GET /services/{id}][%d] serviceInspectOK  %+v", 200, o.Payload)
}

func (o *ServiceInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInspectNotFound creates a ServiceInspectNotFound with default headers values
func NewServiceInspectNotFound() *ServiceInspectNotFound {
	return &ServiceInspectNotFound{}
}

/*ServiceInspectNotFound handles this case with default header values.

no such service
*/
type ServiceInspectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ServiceInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /services/{id}][%d] serviceInspectNotFound  %+v", 404, o.Payload)
}

func (o *ServiceInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInspectInternalServerError creates a ServiceInspectInternalServerError with default headers values
func NewServiceInspectInternalServerError() *ServiceInspectInternalServerError {
	return &ServiceInspectInternalServerError{}
}

/*ServiceInspectInternalServerError handles this case with default header values.

server error
*/
type ServiceInspectInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ServiceInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /services/{id}][%d] serviceInspectInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInspectServiceUnavailable creates a ServiceInspectServiceUnavailable with default headers values
func NewServiceInspectServiceUnavailable() *ServiceInspectServiceUnavailable {
	return &ServiceInspectServiceUnavailable{}
}

/*ServiceInspectServiceUnavailable handles this case with default header values.

node is not part of a swarm
*/
type ServiceInspectServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *ServiceInspectServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /services/{id}][%d] serviceInspectServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ServiceInspectServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
