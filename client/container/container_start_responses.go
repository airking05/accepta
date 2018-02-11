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

// ContainerStartReader is a Reader for the ContainerStart structure.
type ContainerStartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerStartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewContainerStartNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewContainerStartNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewContainerStartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerStartInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerStartNoContent creates a ContainerStartNoContent with default headers values
func NewContainerStartNoContent() *ContainerStartNoContent {
	return &ContainerStartNoContent{}
}

/*ContainerStartNoContent handles this case with default header values.

no error
*/
type ContainerStartNoContent struct {
}

func (o *ContainerStartNoContent) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/start][%d] containerStartNoContent ", 204)
}

func (o *ContainerStartNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerStartNotModified creates a ContainerStartNotModified with default headers values
func NewContainerStartNotModified() *ContainerStartNotModified {
	return &ContainerStartNotModified{}
}

/*ContainerStartNotModified handles this case with default header values.

container already started
*/
type ContainerStartNotModified struct {
	Payload *models.ErrorResponse
}

func (o *ContainerStartNotModified) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/start][%d] containerStartNotModified  %+v", 304, o.Payload)
}

func (o *ContainerStartNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerStartNotFound creates a ContainerStartNotFound with default headers values
func NewContainerStartNotFound() *ContainerStartNotFound {
	return &ContainerStartNotFound{}
}

/*ContainerStartNotFound handles this case with default header values.

no such container
*/
type ContainerStartNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerStartNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/start][%d] containerStartNotFound  %+v", 404, o.Payload)
}

func (o *ContainerStartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerStartInternalServerError creates a ContainerStartInternalServerError with default headers values
func NewContainerStartInternalServerError() *ContainerStartInternalServerError {
	return &ContainerStartInternalServerError{}
}

/*ContainerStartInternalServerError handles this case with default header values.

server error
*/
type ContainerStartInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerStartInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/start][%d] containerStartInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerStartInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
