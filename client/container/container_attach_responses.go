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

// ContainerAttachReader is a Reader for the ContainerAttach structure.
type ContainerAttachReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerAttachReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 101:
		result := NewContainerAttachSwitchingProtocols()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 200:
		result := NewContainerAttachOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewContainerAttachBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewContainerAttachNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerAttachInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerAttachSwitchingProtocols creates a ContainerAttachSwitchingProtocols with default headers values
func NewContainerAttachSwitchingProtocols() *ContainerAttachSwitchingProtocols {
	return &ContainerAttachSwitchingProtocols{}
}

/*ContainerAttachSwitchingProtocols handles this case with default header values.

no error, hints proxy about hijacking
*/
type ContainerAttachSwitchingProtocols struct {
}

func (o *ContainerAttachSwitchingProtocols) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachSwitchingProtocols ", 101)
}

func (o *ContainerAttachSwitchingProtocols) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerAttachOK creates a ContainerAttachOK with default headers values
func NewContainerAttachOK() *ContainerAttachOK {
	return &ContainerAttachOK{}
}

/*ContainerAttachOK handles this case with default header values.

no error, no upgrade header found
*/
type ContainerAttachOK struct {
}

func (o *ContainerAttachOK) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachOK ", 200)
}

func (o *ContainerAttachOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerAttachBadRequest creates a ContainerAttachBadRequest with default headers values
func NewContainerAttachBadRequest() *ContainerAttachBadRequest {
	return &ContainerAttachBadRequest{}
}

/*ContainerAttachBadRequest handles this case with default header values.

bad parameter
*/
type ContainerAttachBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ContainerAttachBadRequest) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachBadRequest  %+v", 400, o.Payload)
}

func (o *ContainerAttachBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerAttachNotFound creates a ContainerAttachNotFound with default headers values
func NewContainerAttachNotFound() *ContainerAttachNotFound {
	return &ContainerAttachNotFound{}
}

/*ContainerAttachNotFound handles this case with default header values.

no such container
*/
type ContainerAttachNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerAttachNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachNotFound  %+v", 404, o.Payload)
}

func (o *ContainerAttachNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerAttachInternalServerError creates a ContainerAttachInternalServerError with default headers values
func NewContainerAttachInternalServerError() *ContainerAttachInternalServerError {
	return &ContainerAttachInternalServerError{}
}

/*ContainerAttachInternalServerError handles this case with default header values.

server error
*/
type ContainerAttachInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerAttachInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerAttachInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
