// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPluginEnableParams creates a new PluginEnableParams object
// with the default values initialized.
func NewPluginEnableParams() *PluginEnableParams {
	var (
		timeoutDefault = int64(0)
	)
	return &PluginEnableParams{
		Timeout: &timeoutDefault,

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewPluginEnableParamsWithTimeout creates a new PluginEnableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPluginEnableParamsWithTimeout(timeout time.Duration) *PluginEnableParams {
	var (
		timeoutDefault = int64(0)
	)
	return &PluginEnableParams{
		Timeout: &timeoutDefault,

		requestTimeout: timeout,
	}
}

// NewPluginEnableParamsWithContext creates a new PluginEnableParams object
// with the default values initialized, and the ability to set a context for a request
func NewPluginEnableParamsWithContext(ctx context.Context) *PluginEnableParams {
	var (
		timeoutDefault = int64(0)
	)
	return &PluginEnableParams{
		Timeout: &timeoutDefault,

		Context: ctx,
	}
}

// NewPluginEnableParamsWithHTTPClient creates a new PluginEnableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPluginEnableParamsWithHTTPClient(client *http.Client) *PluginEnableParams {
	var (
		timeoutDefault = int64(0)
	)
	return &PluginEnableParams{
		Timeout:    &timeoutDefault,
		HTTPClient: client,
	}
}

/*PluginEnableParams contains all the parameters to send to the API endpoint
for the plugin enable operation typically these are written to a http.Request
*/
type PluginEnableParams struct {

	/*Name
	  The name of the plugin. The `:latest` tag is optional, and is the default if omitted.

	*/
	Name string
	/*Timeout
	  Set the HTTP client timeout (in seconds)

	*/
	Timeout *int64

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the plugin enable params
func (o *PluginEnableParams) WithRequestTimeout(timeout time.Duration) *PluginEnableParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the plugin enable params
func (o *PluginEnableParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the plugin enable params
func (o *PluginEnableParams) WithContext(ctx context.Context) *PluginEnableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugin enable params
func (o *PluginEnableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugin enable params
func (o *PluginEnableParams) WithHTTPClient(client *http.Client) *PluginEnableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugin enable params
func (o *PluginEnableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the plugin enable params
func (o *PluginEnableParams) WithName(name string) *PluginEnableParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the plugin enable params
func (o *PluginEnableParams) SetName(name string) {
	o.Name = name
}

// WithTimeout adds the timeout to the plugin enable params
func (o *PluginEnableParams) WithTimeout(timeout *int64) *PluginEnableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugin enable params
func (o *PluginEnableParams) SetTimeout(timeout *int64) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *PluginEnableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout int64
		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := swag.FormatInt64(qrTimeout)
		if qTimeout != "" {
			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
