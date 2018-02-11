// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewSystemEventsParams creates a new SystemEventsParams object
// with the default values initialized.
func NewSystemEventsParams() *SystemEventsParams {
	var ()
	return &SystemEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSystemEventsParamsWithTimeout creates a new SystemEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSystemEventsParamsWithTimeout(timeout time.Duration) *SystemEventsParams {
	var ()
	return &SystemEventsParams{

		timeout: timeout,
	}
}

// NewSystemEventsParamsWithContext creates a new SystemEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSystemEventsParamsWithContext(ctx context.Context) *SystemEventsParams {
	var ()
	return &SystemEventsParams{

		Context: ctx,
	}
}

// NewSystemEventsParamsWithHTTPClient creates a new SystemEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSystemEventsParamsWithHTTPClient(client *http.Client) *SystemEventsParams {
	var ()
	return &SystemEventsParams{
		HTTPClient: client,
	}
}

/*SystemEventsParams contains all the parameters to send to the API endpoint
for the system events operation typically these are written to a http.Request
*/
type SystemEventsParams struct {

	/*Filters
	  A JSON encoded value of filters (a `map[string][]string`) to process on the event list. Available filters:

	- `config=<string>` config name or ID
	- `container=<string>` container name or ID
	- `daemon=<string>` daemon name or ID
	- `event=<string>` event type
	- `image=<string>` image name or ID
	- `label=<string>` image or container label
	- `network=<string>` network name or ID
	- `node=<string>` node ID
	- `plugin`=<string> plugin name or ID
	- `scope`＝<string> local or swarm
	- `secret=<string>` secret name or ID
	- `service=<string>` service name or ID
	- `type=<string>` object to filter by, one of `container`, `image`, `volume`, `network`, `daemon`, `plugin`, `node`, `service`, `secret` or `config`
	- `volume=<string>` volume name


	*/
	Filters *string
	/*Since
	  Show events created since this timestamp then stream new events.

	*/
	Since *string
	/*Until
	  Show events created until this timestamp then stop streaming.

	*/
	Until *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the system events params
func (o *SystemEventsParams) WithTimeout(timeout time.Duration) *SystemEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system events params
func (o *SystemEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system events params
func (o *SystemEventsParams) WithContext(ctx context.Context) *SystemEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system events params
func (o *SystemEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system events params
func (o *SystemEventsParams) WithHTTPClient(client *http.Client) *SystemEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system events params
func (o *SystemEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the system events params
func (o *SystemEventsParams) WithFilters(filters *string) *SystemEventsParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the system events params
func (o *SystemEventsParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithSince adds the since to the system events params
func (o *SystemEventsParams) WithSince(since *string) *SystemEventsParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the system events params
func (o *SystemEventsParams) SetSince(since *string) {
	o.Since = since
}

// WithUntil adds the until to the system events params
func (o *SystemEventsParams) WithUntil(until *string) *SystemEventsParams {
	o.SetUntil(until)
	return o
}

// SetUntil adds the until to the system events params
func (o *SystemEventsParams) SetUntil(until *string) {
	o.Until = until
}

// WriteToRequest writes these params to a swagger request
func (o *SystemEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filters != nil {

		// query param filters
		var qrFilters string
		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {
			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}

	}

	if o.Since != nil {

		// query param since
		var qrSince string
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if o.Until != nil {

		// query param until
		var qrUntil string
		if o.Until != nil {
			qrUntil = *o.Until
		}
		qUntil := qrUntil
		if qUntil != "" {
			if err := r.SetQueryParam("until", qUntil); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
