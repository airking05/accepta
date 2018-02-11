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

// ContainerChangesReader is a Reader for the ContainerChanges structure.
type ContainerChangesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerChangesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContainerChangesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewContainerChangesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerChangesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerChangesOK creates a ContainerChangesOK with default headers values
func NewContainerChangesOK() *ContainerChangesOK {
	return &ContainerChangesOK{}
}

/*ContainerChangesOK handles this case with default header values.

The list of changes
*/
type ContainerChangesOK struct {
	Payload models.ContainerChangesOKBody
}

func (o *ContainerChangesOK) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/changes][%d] containerChangesOK  %+v", 200, o.Payload)
}

func (o *ContainerChangesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerChangesNotFound creates a ContainerChangesNotFound with default headers values
func NewContainerChangesNotFound() *ContainerChangesNotFound {
	return &ContainerChangesNotFound{}
}

/*ContainerChangesNotFound handles this case with default header values.

no such container
*/
type ContainerChangesNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerChangesNotFound) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/changes][%d] containerChangesNotFound  %+v", 404, o.Payload)
}

func (o *ContainerChangesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerChangesInternalServerError creates a ContainerChangesInternalServerError with default headers values
func NewContainerChangesInternalServerError() *ContainerChangesInternalServerError {
	return &ContainerChangesInternalServerError{}
}

/*ContainerChangesInternalServerError handles this case with default header values.

server error
*/
type ContainerChangesInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerChangesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/changes][%d] containerChangesInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerChangesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
