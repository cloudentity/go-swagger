// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloudentity/go-swagger/examples/task-tracker/models"
)

// UpdateTaskOKCode is the HTTP code returned for type UpdateTaskOK
const UpdateTaskOKCode int = 200

/*
UpdateTaskOK Task details

swagger:response updateTaskOK
*/
type UpdateTaskOK struct {

	/*
	  In: Body
	*/
	Payload *models.Task `json:"body,omitempty"`
}

// NewUpdateTaskOK creates UpdateTaskOK with default headers values
func NewUpdateTaskOK() *UpdateTaskOK {

	return &UpdateTaskOK{}
}

// WithPayload adds the payload to the update task o k response
func (o *UpdateTaskOK) WithPayload(payload *models.Task) *UpdateTaskOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update task o k response
func (o *UpdateTaskOK) SetPayload(payload *models.Task) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTaskOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateTaskUnprocessableEntityCode is the HTTP code returned for type UpdateTaskUnprocessableEntity
const UpdateTaskUnprocessableEntityCode int = 422

/*
UpdateTaskUnprocessableEntity Validation error

swagger:response updateTaskUnprocessableEntity
*/
type UpdateTaskUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewUpdateTaskUnprocessableEntity creates UpdateTaskUnprocessableEntity with default headers values
func NewUpdateTaskUnprocessableEntity() *UpdateTaskUnprocessableEntity {

	return &UpdateTaskUnprocessableEntity{}
}

// WithPayload adds the payload to the update task unprocessable entity response
func (o *UpdateTaskUnprocessableEntity) WithPayload(payload *models.ValidationError) *UpdateTaskUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update task unprocessable entity response
func (o *UpdateTaskUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTaskUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
UpdateTaskDefault Error response

swagger:response updateTaskDefault
*/
type UpdateTaskDefault struct {
	_statusCode int
	/*

	 */
	XErrorCode string `json:"X-Error-Code"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateTaskDefault creates UpdateTaskDefault with default headers values
func NewUpdateTaskDefault(code int) *UpdateTaskDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateTaskDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update task default response
func (o *UpdateTaskDefault) WithStatusCode(code int) *UpdateTaskDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update task default response
func (o *UpdateTaskDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithXErrorCode adds the xErrorCode to the update task default response
func (o *UpdateTaskDefault) WithXErrorCode(xErrorCode string) *UpdateTaskDefault {
	o.XErrorCode = xErrorCode
	return o
}

// SetXErrorCode sets the xErrorCode to the update task default response
func (o *UpdateTaskDefault) SetXErrorCode(xErrorCode string) {
	o.XErrorCode = xErrorCode
}

// WithPayload adds the payload to the update task default response
func (o *UpdateTaskDefault) WithPayload(payload *models.Error) *UpdateTaskDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update task default response
func (o *UpdateTaskDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTaskDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Error-Code

	xErrorCode := o.XErrorCode
	if xErrorCode != "" {
		rw.Header().Set("X-Error-Code", xErrorCode)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
