// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/go-swagger/examples/todo-list/models"
)

// DestroyOneReader is a Reader for the DestroyOne structure.
type DestroyOneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DestroyOneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDestroyOneNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDestroyOneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDestroyOneNoContent creates a DestroyOneNoContent with default headers values
func NewDestroyOneNoContent() *DestroyOneNoContent {
	return &DestroyOneNoContent{}
}

/*
	DestroyOneNoContent describes a response with status code 204, with default header values.

Deleted
*/
type DestroyOneNoContent struct {
}

// IsSuccess returns true when this destroy one no content response returns a 2xx status code
func (o *DestroyOneNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this destroy one no content response returns a 3xx status code
func (o *DestroyOneNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this destroy one no content response returns a 4xx status code
func (o *DestroyOneNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this destroy one no content response returns a 5xx status code
func (o *DestroyOneNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this destroy one no content response returns a 4xx status code
func (o *DestroyOneNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DestroyOneNoContent) Error() string {
	return fmt.Sprintf("[DELETE /{id}][%d] destroyOneNoContent ", 204)
}

func (o *DestroyOneNoContent) String() string {
	return fmt.Sprintf("[DELETE /{id}][%d] destroyOneNoContent ", 204)
}

func (o *DestroyOneNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDestroyOneDefault creates a DestroyOneDefault with default headers values
func NewDestroyOneDefault(code int) *DestroyOneDefault {
	return &DestroyOneDefault{
		_statusCode: code,
	}
}

/*
	DestroyOneDefault describes a response with status code -1, with default header values.

error
*/
type DestroyOneDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the destroy one default response
func (o *DestroyOneDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this destroy one default response returns a 2xx status code
func (o *DestroyOneDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this destroy one default response returns a 3xx status code
func (o *DestroyOneDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this destroy one default response returns a 4xx status code
func (o *DestroyOneDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this destroy one default response returns a 5xx status code
func (o *DestroyOneDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this destroy one default response returns a 4xx status code
func (o *DestroyOneDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DestroyOneDefault) Error() string {
	return fmt.Sprintf("[DELETE /{id}][%d] destroyOne default  %+v", o._statusCode, o.Payload)
}

func (o *DestroyOneDefault) String() string {
	return fmt.Sprintf("[DELETE /{id}][%d] destroyOne default  %+v", o._statusCode, o.Payload)
}

func (o *DestroyOneDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DestroyOneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
