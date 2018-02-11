// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/go-docker-api-client/models"
)

// PluginCreateReader is a Reader for the PluginCreate structure.
type PluginCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPluginCreateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewPluginCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPluginCreateNoContent creates a PluginCreateNoContent with default headers values
func NewPluginCreateNoContent() *PluginCreateNoContent {
	return &PluginCreateNoContent{}
}

/*PluginCreateNoContent handles this case with default header values.

no error
*/
type PluginCreateNoContent struct {
}

func (o *PluginCreateNoContent) Error() string {
	return fmt.Sprintf("[POST /plugins/create][%d] pluginCreateNoContent ", 204)
}

func (o *PluginCreateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPluginCreateInternalServerError creates a PluginCreateInternalServerError with default headers values
func NewPluginCreateInternalServerError() *PluginCreateInternalServerError {
	return &PluginCreateInternalServerError{}
}

/*PluginCreateInternalServerError handles this case with default header values.

server error
*/
type PluginCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PluginCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /plugins/create][%d] pluginCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *PluginCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
