// Code generated by go-swagger; DO NOT EDIT.

package container

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

// NewContainerArchiveParams creates a new ContainerArchiveParams object
// with the default values initialized.
func NewContainerArchiveParams() *ContainerArchiveParams {
	var ()
	return &ContainerArchiveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContainerArchiveParamsWithTimeout creates a new ContainerArchiveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContainerArchiveParamsWithTimeout(timeout time.Duration) *ContainerArchiveParams {
	var ()
	return &ContainerArchiveParams{

		timeout: timeout,
	}
}

// NewContainerArchiveParamsWithContext creates a new ContainerArchiveParams object
// with the default values initialized, and the ability to set a context for a request
func NewContainerArchiveParamsWithContext(ctx context.Context) *ContainerArchiveParams {
	var ()
	return &ContainerArchiveParams{

		Context: ctx,
	}
}

// NewContainerArchiveParamsWithHTTPClient creates a new ContainerArchiveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContainerArchiveParamsWithHTTPClient(client *http.Client) *ContainerArchiveParams {
	var ()
	return &ContainerArchiveParams{
		HTTPClient: client,
	}
}

/*ContainerArchiveParams contains all the parameters to send to the API endpoint
for the container archive operation typically these are written to a http.Request
*/
type ContainerArchiveParams struct {

	/*ID
	  ID or name of the container

	*/
	ID string
	/*Path
	  Resource in the container’s filesystem to archive.

	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the container archive params
func (o *ContainerArchiveParams) WithTimeout(timeout time.Duration) *ContainerArchiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container archive params
func (o *ContainerArchiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container archive params
func (o *ContainerArchiveParams) WithContext(ctx context.Context) *ContainerArchiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container archive params
func (o *ContainerArchiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container archive params
func (o *ContainerArchiveParams) WithHTTPClient(client *http.Client) *ContainerArchiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container archive params
func (o *ContainerArchiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the container archive params
func (o *ContainerArchiveParams) WithID(id string) *ContainerArchiveParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the container archive params
func (o *ContainerArchiveParams) SetID(id string) {
	o.ID = id
}

// WithPath adds the path to the container archive params
func (o *ContainerArchiveParams) WithPath(path string) *ContainerArchiveParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the container archive params
func (o *ContainerArchiveParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerArchiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {
		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
