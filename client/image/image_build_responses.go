// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/go-docker-api-client/models"
)

// ImageBuildReader is a Reader for the ImageBuild structure.
type ImageBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewImageBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewImageBuildBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewImageBuildInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImageBuildOK creates a ImageBuildOK with default headers values
func NewImageBuildOK() *ImageBuildOK {
	return &ImageBuildOK{}
}

/*ImageBuildOK handles this case with default header values.

no error
*/
type ImageBuildOK struct {
}

func (o *ImageBuildOK) Error() string {
	return fmt.Sprintf("[POST /build][%d] imageBuildOK ", 200)
}

func (o *ImageBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImageBuildBadRequest creates a ImageBuildBadRequest with default headers values
func NewImageBuildBadRequest() *ImageBuildBadRequest {
	return &ImageBuildBadRequest{}
}

/*ImageBuildBadRequest handles this case with default header values.

Bad parameter
*/
type ImageBuildBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ImageBuildBadRequest) Error() string {
	return fmt.Sprintf("[POST /build][%d] imageBuildBadRequest  %+v", 400, o.Payload)
}

func (o *ImageBuildBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageBuildInternalServerError creates a ImageBuildInternalServerError with default headers values
func NewImageBuildInternalServerError() *ImageBuildInternalServerError {
	return &ImageBuildInternalServerError{}
}

/*ImageBuildInternalServerError handles this case with default header values.

server error
*/
type ImageBuildInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ImageBuildInternalServerError) Error() string {
	return fmt.Sprintf("[POST /build][%d] imageBuildInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageBuildInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
