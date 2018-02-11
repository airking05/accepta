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

// PutContainerArchiveReader is a Reader for the PutContainerArchive structure.
type PutContainerArchiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutContainerArchiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutContainerArchiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutContainerArchiveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutContainerArchiveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutContainerArchiveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutContainerArchiveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutContainerArchiveOK creates a PutContainerArchiveOK with default headers values
func NewPutContainerArchiveOK() *PutContainerArchiveOK {
	return &PutContainerArchiveOK{}
}

/*PutContainerArchiveOK handles this case with default header values.

The content was extracted successfully
*/
type PutContainerArchiveOK struct {
}

func (o *PutContainerArchiveOK) Error() string {
	return fmt.Sprintf("[PUT /containers/{id}/archive][%d] putContainerArchiveOK ", 200)
}

func (o *PutContainerArchiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutContainerArchiveBadRequest creates a PutContainerArchiveBadRequest with default headers values
func NewPutContainerArchiveBadRequest() *PutContainerArchiveBadRequest {
	return &PutContainerArchiveBadRequest{}
}

/*PutContainerArchiveBadRequest handles this case with default header values.

Bad parameter
*/
type PutContainerArchiveBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PutContainerArchiveBadRequest) Error() string {
	return fmt.Sprintf("[PUT /containers/{id}/archive][%d] putContainerArchiveBadRequest  %+v", 400, o.Payload)
}

func (o *PutContainerArchiveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContainerArchiveForbidden creates a PutContainerArchiveForbidden with default headers values
func NewPutContainerArchiveForbidden() *PutContainerArchiveForbidden {
	return &PutContainerArchiveForbidden{}
}

/*PutContainerArchiveForbidden handles this case with default header values.

Permission denied, the volume or container rootfs is marked as read-only.
*/
type PutContainerArchiveForbidden struct {
	Payload *models.ErrorResponse
}

func (o *PutContainerArchiveForbidden) Error() string {
	return fmt.Sprintf("[PUT /containers/{id}/archive][%d] putContainerArchiveForbidden  %+v", 403, o.Payload)
}

func (o *PutContainerArchiveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContainerArchiveNotFound creates a PutContainerArchiveNotFound with default headers values
func NewPutContainerArchiveNotFound() *PutContainerArchiveNotFound {
	return &PutContainerArchiveNotFound{}
}

/*PutContainerArchiveNotFound handles this case with default header values.

No such container or path does not exist inside the container
*/
type PutContainerArchiveNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PutContainerArchiveNotFound) Error() string {
	return fmt.Sprintf("[PUT /containers/{id}/archive][%d] putContainerArchiveNotFound  %+v", 404, o.Payload)
}

func (o *PutContainerArchiveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContainerArchiveInternalServerError creates a PutContainerArchiveInternalServerError with default headers values
func NewPutContainerArchiveInternalServerError() *PutContainerArchiveInternalServerError {
	return &PutContainerArchiveInternalServerError{}
}

/*PutContainerArchiveInternalServerError handles this case with default header values.

Server error
*/
type PutContainerArchiveInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PutContainerArchiveInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /containers/{id}/archive][%d] putContainerArchiveInternalServerError  %+v", 500, o.Payload)
}

func (o *PutContainerArchiveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
