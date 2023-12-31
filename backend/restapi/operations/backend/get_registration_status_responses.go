// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/project-flotta/flotta-operator/backend/models"
)

// GetRegistrationStatusOKCode is the HTTP code returned for type GetRegistrationStatusOK
const GetRegistrationStatusOKCode int = 200

/*
GetRegistrationStatusOK OK

swagger:response getRegistrationStatusOK
*/
type GetRegistrationStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.DeviceRegistrationStatusResponse `json:"body,omitempty"`
}

// NewGetRegistrationStatusOK creates GetRegistrationStatusOK with default headers values
func NewGetRegistrationStatusOK() *GetRegistrationStatusOK {

	return &GetRegistrationStatusOK{}
}

// WithPayload adds the payload to the get registration status o k response
func (o *GetRegistrationStatusOK) WithPayload(payload *models.DeviceRegistrationStatusResponse) *GetRegistrationStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get registration status o k response
func (o *GetRegistrationStatusOK) SetPayload(payload *models.DeviceRegistrationStatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRegistrationStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRegistrationStatusUnauthorizedCode is the HTTP code returned for type GetRegistrationStatusUnauthorized
const GetRegistrationStatusUnauthorizedCode int = 401

/*
GetRegistrationStatusUnauthorized Unauthorized

swagger:response getRegistrationStatusUnauthorized
*/
type GetRegistrationStatusUnauthorized struct {
}

// NewGetRegistrationStatusUnauthorized creates GetRegistrationStatusUnauthorized with default headers values
func NewGetRegistrationStatusUnauthorized() *GetRegistrationStatusUnauthorized {

	return &GetRegistrationStatusUnauthorized{}
}

// WriteResponse to the client
func (o *GetRegistrationStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetRegistrationStatusForbiddenCode is the HTTP code returned for type GetRegistrationStatusForbidden
const GetRegistrationStatusForbiddenCode int = 403

/*
GetRegistrationStatusForbidden Forbidden

swagger:response getRegistrationStatusForbidden
*/
type GetRegistrationStatusForbidden struct {
}

// NewGetRegistrationStatusForbidden creates GetRegistrationStatusForbidden with default headers values
func NewGetRegistrationStatusForbidden() *GetRegistrationStatusForbidden {

	return &GetRegistrationStatusForbidden{}
}

// WriteResponse to the client
func (o *GetRegistrationStatusForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

/*
GetRegistrationStatusDefault Error

swagger:response getRegistrationStatusDefault
*/
type GetRegistrationStatusDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.DeviceRegistrationStatusResponse `json:"body,omitempty"`
}

// NewGetRegistrationStatusDefault creates GetRegistrationStatusDefault with default headers values
func NewGetRegistrationStatusDefault(code int) *GetRegistrationStatusDefault {
	if code <= 0 {
		code = 500
	}

	return &GetRegistrationStatusDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get registration status default response
func (o *GetRegistrationStatusDefault) WithStatusCode(code int) *GetRegistrationStatusDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get registration status default response
func (o *GetRegistrationStatusDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get registration status default response
func (o *GetRegistrationStatusDefault) WithPayload(payload *models.DeviceRegistrationStatusResponse) *GetRegistrationStatusDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get registration status default response
func (o *GetRegistrationStatusDefault) SetPayload(payload *models.DeviceRegistrationStatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRegistrationStatusDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
