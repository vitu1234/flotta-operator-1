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

	commonmodel "github.com/project-flotta/flotta-operator/models"
)

// NewRegisterDeviceParams creates a new RegisterDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegisterDeviceParams() *RegisterDeviceParams {
	return &RegisterDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterDeviceParamsWithTimeout creates a new RegisterDeviceParams object
// with the ability to set a timeout on a request.
func NewRegisterDeviceParamsWithTimeout(timeout time.Duration) *RegisterDeviceParams {
	return &RegisterDeviceParams{
		timeout: timeout,
	}
}

// NewRegisterDeviceParamsWithContext creates a new RegisterDeviceParams object
// with the ability to set a context for a request.
func NewRegisterDeviceParamsWithContext(ctx context.Context) *RegisterDeviceParams {
	return &RegisterDeviceParams{
		Context: ctx,
	}
}

// NewRegisterDeviceParamsWithHTTPClient creates a new RegisterDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewRegisterDeviceParamsWithHTTPClient(client *http.Client) *RegisterDeviceParams {
	return &RegisterDeviceParams{
		HTTPClient: client,
	}
}

/*
RegisterDeviceParams contains all the parameters to send to the API endpoint

	for the register device operation.

	Typically these are written to a http.Request.
*/
type RegisterDeviceParams struct {

	/* DeviceID.

	   Device ID
	*/
	DeviceID string

	/* Namespace.

	   Namespace where the device resides
	*/
	Namespace string

	// RegistrationInfo.
	RegistrationInfo commonmodel.RegistrationInfo

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the register device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterDeviceParams) WithDefaults() *RegisterDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the register device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterDeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the register device params
func (o *RegisterDeviceParams) WithTimeout(timeout time.Duration) *RegisterDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register device params
func (o *RegisterDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register device params
func (o *RegisterDeviceParams) WithContext(ctx context.Context) *RegisterDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register device params
func (o *RegisterDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register device params
func (o *RegisterDeviceParams) WithHTTPClient(client *http.Client) *RegisterDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register device params
func (o *RegisterDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the register device params
func (o *RegisterDeviceParams) WithDeviceID(deviceID string) *RegisterDeviceParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the register device params
func (o *RegisterDeviceParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithNamespace adds the namespace to the register device params
func (o *RegisterDeviceParams) WithNamespace(namespace string) *RegisterDeviceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the register device params
func (o *RegisterDeviceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithRegistrationInfo adds the registrationInfo to the register device params
func (o *RegisterDeviceParams) WithRegistrationInfo(registrationInfo commonmodel.RegistrationInfo) *RegisterDeviceParams {
	o.SetRegistrationInfo(registrationInfo)
	return o
}

// SetRegistrationInfo adds the registrationInfo to the register device params
func (o *RegisterDeviceParams) SetRegistrationInfo(registrationInfo commonmodel.RegistrationInfo) {
	o.RegistrationInfo = registrationInfo
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.RegistrationInfo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
