// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewImagePushParams creates a new ImagePushParams object
// with the default values initialized.
func NewImagePushParams() *ImagePushParams {
	var ()
	return &ImagePushParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImagePushParamsWithTimeout creates a new ImagePushParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImagePushParamsWithTimeout(timeout time.Duration) *ImagePushParams {
	var ()
	return &ImagePushParams{

		timeout: timeout,
	}
}

// NewImagePushParamsWithContext creates a new ImagePushParams object
// with the default values initialized, and the ability to set a context for a request
func NewImagePushParamsWithContext(ctx context.Context) *ImagePushParams {
	var ()
	return &ImagePushParams{

		Context: ctx,
	}
}

// NewImagePushParamsWithHTTPClient creates a new ImagePushParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImagePushParamsWithHTTPClient(client *http.Client) *ImagePushParams {
	var ()
	return &ImagePushParams{
		HTTPClient: client,
	}
}

/*ImagePushParams contains all the parameters to send to the API endpoint
for the image push operation typically these are written to a http.Request
*/
type ImagePushParams struct {

	/*XRegistryAuth
	  A base64-encoded auth configuration. [See the authentication section for details.](#section/Authentication)

	*/
	XRegistryAuth string
	/*Name
	  Image name or ID.

	*/
	Name string
	/*Tag
	  The tag to associate with the image on the registry.

	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the image push params
func (o *ImagePushParams) WithTimeout(timeout time.Duration) *ImagePushParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image push params
func (o *ImagePushParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image push params
func (o *ImagePushParams) WithContext(ctx context.Context) *ImagePushParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image push params
func (o *ImagePushParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image push params
func (o *ImagePushParams) WithHTTPClient(client *http.Client) *ImagePushParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image push params
func (o *ImagePushParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRegistryAuth adds the xRegistryAuth to the image push params
func (o *ImagePushParams) WithXRegistryAuth(xRegistryAuth string) *ImagePushParams {
	o.SetXRegistryAuth(xRegistryAuth)
	return o
}

// SetXRegistryAuth adds the xRegistryAuth to the image push params
func (o *ImagePushParams) SetXRegistryAuth(xRegistryAuth string) {
	o.XRegistryAuth = xRegistryAuth
}

// WithName adds the name to the image push params
func (o *ImagePushParams) WithName(name string) *ImagePushParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the image push params
func (o *ImagePushParams) SetName(name string) {
	o.Name = name
}

// WithTag adds the tag to the image push params
func (o *ImagePushParams) WithTag(tag *string) *ImagePushParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the image push params
func (o *ImagePushParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *ImagePushParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Registry-Auth
	if err := r.SetHeaderParam("X-Registry-Auth", o.XRegistryAuth); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
