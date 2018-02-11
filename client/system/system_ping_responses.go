// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/go-docker-api-client/models"
)

// SystemPingReader is a Reader for the SystemPing structure.
type SystemPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSystemPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewSystemPingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemPingOK creates a SystemPingOK with default headers values
func NewSystemPingOK() *SystemPingOK {
	return &SystemPingOK{}
}

/*SystemPingOK handles this case with default header values.

no error
*/
type SystemPingOK struct {
	/*Max API Version the server supports
	 */
	APIVersion string
	/*If the server is running with experimental mode enabled
	 */
	DockerExperimental bool

	Payload string
}

func (o *SystemPingOK) Error() string {
	return fmt.Sprintf("[GET /_ping][%d] systemPingOK  %+v", 200, o.Payload)
}

func (o *SystemPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header API-Version
	o.APIVersion = response.GetHeader("API-Version")

	// response header Docker-Experimental
	dockerExperimental, err := swag.ConvertBool(response.GetHeader("Docker-Experimental"))
	if err != nil {
		return errors.InvalidType("Docker-Experimental", "header", "bool", response.GetHeader("Docker-Experimental"))
	}
	o.DockerExperimental = dockerExperimental

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemPingInternalServerError creates a SystemPingInternalServerError with default headers values
func NewSystemPingInternalServerError() *SystemPingInternalServerError {
	return &SystemPingInternalServerError{}
}

/*SystemPingInternalServerError handles this case with default header values.

server error
*/
type SystemPingInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SystemPingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /_ping][%d] systemPingInternalServerError  %+v", 500, o.Payload)
}

func (o *SystemPingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
