// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/go-docker-api-client/models"
)

// ContainerStopReader is a Reader for the ContainerStop structure.
type ContainerStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewContainerStopNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewContainerStopNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewContainerStopNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerStopInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerStopNoContent creates a ContainerStopNoContent with default headers values
func NewContainerStopNoContent() *ContainerStopNoContent {
	return &ContainerStopNoContent{}
}

/*ContainerStopNoContent handles this case with default header values.

no error
*/
type ContainerStopNoContent struct {
}

func (o *ContainerStopNoContent) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/stop][%d] containerStopNoContent ", 204)
}

func (o *ContainerStopNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerStopNotModified creates a ContainerStopNotModified with default headers values
func NewContainerStopNotModified() *ContainerStopNotModified {
	return &ContainerStopNotModified{}
}

/*ContainerStopNotModified handles this case with default header values.

container already stopped
*/
type ContainerStopNotModified struct {
	Payload *models.ErrorResponse
}

func (o *ContainerStopNotModified) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/stop][%d] containerStopNotModified  %+v", 304, o.Payload)
}

func (o *ContainerStopNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerStopNotFound creates a ContainerStopNotFound with default headers values
func NewContainerStopNotFound() *ContainerStopNotFound {
	return &ContainerStopNotFound{}
}

/*ContainerStopNotFound handles this case with default header values.

no such container
*/
type ContainerStopNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerStopNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/stop][%d] containerStopNotFound  %+v", 404, o.Payload)
}

func (o *ContainerStopNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerStopInternalServerError creates a ContainerStopInternalServerError with default headers values
func NewContainerStopInternalServerError() *ContainerStopInternalServerError {
	return &ContainerStopInternalServerError{}
}

/*ContainerStopInternalServerError handles this case with default header values.

server error
*/
type ContainerStopInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerStopInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/stop][%d] containerStopInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerStopInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
