// Code generated by go-swagger; DO NOT EDIT.

package swarm

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

// NewSwarmUnlockkeyParams creates a new SwarmUnlockkeyParams object
// with the default values initialized.
func NewSwarmUnlockkeyParams() *SwarmUnlockkeyParams {

	return &SwarmUnlockkeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSwarmUnlockkeyParamsWithTimeout creates a new SwarmUnlockkeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSwarmUnlockkeyParamsWithTimeout(timeout time.Duration) *SwarmUnlockkeyParams {

	return &SwarmUnlockkeyParams{

		timeout: timeout,
	}
}

// NewSwarmUnlockkeyParamsWithContext creates a new SwarmUnlockkeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewSwarmUnlockkeyParamsWithContext(ctx context.Context) *SwarmUnlockkeyParams {

	return &SwarmUnlockkeyParams{

		Context: ctx,
	}
}

// NewSwarmUnlockkeyParamsWithHTTPClient creates a new SwarmUnlockkeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSwarmUnlockkeyParamsWithHTTPClient(client *http.Client) *SwarmUnlockkeyParams {

	return &SwarmUnlockkeyParams{
		HTTPClient: client,
	}
}

/*SwarmUnlockkeyParams contains all the parameters to send to the API endpoint
for the swarm unlockkey operation typically these are written to a http.Request
*/
type SwarmUnlockkeyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the swarm unlockkey params
func (o *SwarmUnlockkeyParams) WithTimeout(timeout time.Duration) *SwarmUnlockkeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the swarm unlockkey params
func (o *SwarmUnlockkeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the swarm unlockkey params
func (o *SwarmUnlockkeyParams) WithContext(ctx context.Context) *SwarmUnlockkeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the swarm unlockkey params
func (o *SwarmUnlockkeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the swarm unlockkey params
func (o *SwarmUnlockkeyParams) WithHTTPClient(client *http.Client) *SwarmUnlockkeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the swarm unlockkey params
func (o *SwarmUnlockkeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SwarmUnlockkeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}