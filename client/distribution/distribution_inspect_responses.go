// Code generated by go-swagger; DO NOT EDIT.

package distribution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/go-docker-api-client/models"
)

// DistributionInspectReader is a Reader for the DistributionInspect structure.
type DistributionInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DistributionInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDistributionInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDistributionInspectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDistributionInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDistributionInspectOK creates a DistributionInspectOK with default headers values
func NewDistributionInspectOK() *DistributionInspectOK {
	return &DistributionInspectOK{}
}

/*DistributionInspectOK handles this case with default header values.

descriptor and platform information
*/
type DistributionInspectOK struct {
	Payload *models.DistributionInspect
}

func (o *DistributionInspectOK) Error() string {
	return fmt.Sprintf("[GET /distribution/{name}/json][%d] distributionInspectOK  %+v", 200, o.Payload)
}

func (o *DistributionInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DistributionInspect)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDistributionInspectUnauthorized creates a DistributionInspectUnauthorized with default headers values
func NewDistributionInspectUnauthorized() *DistributionInspectUnauthorized {
	return &DistributionInspectUnauthorized{}
}

/*DistributionInspectUnauthorized handles this case with default header values.

Failed authentication or no image found
*/
type DistributionInspectUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DistributionInspectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /distribution/{name}/json][%d] distributionInspectUnauthorized  %+v", 401, o.Payload)
}

func (o *DistributionInspectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDistributionInspectInternalServerError creates a DistributionInspectInternalServerError with default headers values
func NewDistributionInspectInternalServerError() *DistributionInspectInternalServerError {
	return &DistributionInspectInternalServerError{}
}

/*DistributionInspectInternalServerError handles this case with default header values.

Server error
*/
type DistributionInspectInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *DistributionInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /distribution/{name}/json][%d] distributionInspectInternalServerError  %+v", 500, o.Payload)
}

func (o *DistributionInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
