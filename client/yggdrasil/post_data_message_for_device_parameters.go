// Code generated by go-swagger; DO NOT EDIT.

package yggdrasil

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

	"github.com/project-flotta/flotta-operator/models"
)

// NewPostDataMessageForDeviceParams creates a new PostDataMessageForDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDataMessageForDeviceParams() *PostDataMessageForDeviceParams {
	return &PostDataMessageForDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDataMessageForDeviceParamsWithTimeout creates a new PostDataMessageForDeviceParams object
// with the ability to set a timeout on a request.
func NewPostDataMessageForDeviceParamsWithTimeout(timeout time.Duration) *PostDataMessageForDeviceParams {
	return &PostDataMessageForDeviceParams{
		timeout: timeout,
	}
}

// NewPostDataMessageForDeviceParamsWithContext creates a new PostDataMessageForDeviceParams object
// with the ability to set a context for a request.
func NewPostDataMessageForDeviceParamsWithContext(ctx context.Context) *PostDataMessageForDeviceParams {
	return &PostDataMessageForDeviceParams{
		Context: ctx,
	}
}

// NewPostDataMessageForDeviceParamsWithHTTPClient creates a new PostDataMessageForDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDataMessageForDeviceParamsWithHTTPClient(client *http.Client) *PostDataMessageForDeviceParams {
	return &PostDataMessageForDeviceParams{
		HTTPClient: client,
	}
}

/*
PostDataMessageForDeviceParams contains all the parameters to send to the API endpoint

	for the post data message for device operation.

	Typically these are written to a http.Request.
*/
type PostDataMessageForDeviceParams struct {

	/* DeviceID.

	   Device ID
	*/
	DeviceID string

	// Message.
	Message *models.Message

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post data message for device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDataMessageForDeviceParams) WithDefaults() *PostDataMessageForDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post data message for device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDataMessageForDeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post data message for device params
func (o *PostDataMessageForDeviceParams) WithTimeout(timeout time.Duration) *PostDataMessageForDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post data message for device params
func (o *PostDataMessageForDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post data message for device params
func (o *PostDataMessageForDeviceParams) WithContext(ctx context.Context) *PostDataMessageForDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post data message for device params
func (o *PostDataMessageForDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post data message for device params
func (o *PostDataMessageForDeviceParams) WithHTTPClient(client *http.Client) *PostDataMessageForDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post data message for device params
func (o *PostDataMessageForDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the post data message for device params
func (o *PostDataMessageForDeviceParams) WithDeviceID(deviceID string) *PostDataMessageForDeviceParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the post data message for device params
func (o *PostDataMessageForDeviceParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithMessage adds the message to the post data message for device params
func (o *PostDataMessageForDeviceParams) WithMessage(message *models.Message) *PostDataMessageForDeviceParams {
	o.SetMessage(message)
	return o
}

// SetMessage adds the message to the post data message for device params
func (o *PostDataMessageForDeviceParams) SetMessage(message *models.Message) {
	o.Message = message
}

// WriteToRequest writes these params to a swagger request
func (o *PostDataMessageForDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device_id
	if err := r.SetPathParam("device_id", o.DeviceID); err != nil {
		return err
	}
	if o.Message != nil {
		if err := r.SetBodyParam(o.Message); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
