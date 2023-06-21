// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetPlaybookExecutionsParams creates a new GetPlaybookExecutionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPlaybookExecutionsParams() *GetPlaybookExecutionsParams {
	return &GetPlaybookExecutionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlaybookExecutionsParamsWithTimeout creates a new GetPlaybookExecutionsParams object
// with the ability to set a timeout on a request.
func NewGetPlaybookExecutionsParamsWithTimeout(timeout time.Duration) *GetPlaybookExecutionsParams {
	return &GetPlaybookExecutionsParams{
		timeout: timeout,
	}
}

// NewGetPlaybookExecutionsParamsWithContext creates a new GetPlaybookExecutionsParams object
// with the ability to set a context for a request.
func NewGetPlaybookExecutionsParamsWithContext(ctx context.Context) *GetPlaybookExecutionsParams {
	return &GetPlaybookExecutionsParams{
		Context: ctx,
	}
}

// NewGetPlaybookExecutionsParamsWithHTTPClient creates a new GetPlaybookExecutionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPlaybookExecutionsParamsWithHTTPClient(client *http.Client) *GetPlaybookExecutionsParams {
	return &GetPlaybookExecutionsParams{
		HTTPClient: client,
	}
}

/* GetPlaybookExecutionsParams contains all the parameters to send to the API endpoint
   for the get playbook executions operation.

   Typically these are written to a http.Request.
*/
type GetPlaybookExecutionsParams struct {

	/* DeviceID.

	   Device ID
	*/
	DeviceID string

	/* Namespace.

	   Namespace where the device resides
	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get playbook executions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlaybookExecutionsParams) WithDefaults() *GetPlaybookExecutionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get playbook executions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlaybookExecutionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get playbook executions params
func (o *GetPlaybookExecutionsParams) WithTimeout(timeout time.Duration) *GetPlaybookExecutionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get playbook executions params
func (o *GetPlaybookExecutionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get playbook executions params
func (o *GetPlaybookExecutionsParams) WithContext(ctx context.Context) *GetPlaybookExecutionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get playbook executions params
func (o *GetPlaybookExecutionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get playbook executions params
func (o *GetPlaybookExecutionsParams) WithHTTPClient(client *http.Client) *GetPlaybookExecutionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get playbook executions params
func (o *GetPlaybookExecutionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get playbook executions params
func (o *GetPlaybookExecutionsParams) WithDeviceID(deviceID string) *GetPlaybookExecutionsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get playbook executions params
func (o *GetPlaybookExecutionsParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithNamespace adds the namespace to the get playbook executions params
func (o *GetPlaybookExecutionsParams) WithNamespace(namespace string) *GetPlaybookExecutionsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get playbook executions params
func (o *GetPlaybookExecutionsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlaybookExecutionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device-id
	if err := r.SetPathParam("device-id", o.DeviceID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
