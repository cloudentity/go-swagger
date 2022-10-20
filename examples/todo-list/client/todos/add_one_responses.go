// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/go-swagger/examples/todo-list/models"
)

// AddOneReader is a Reader for the AddOne structure.
type AddOneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddOneCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddOneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddOneCreated creates a AddOneCreated with default headers values
func NewAddOneCreated() *AddOneCreated {
	return &AddOneCreated{}
}

/*
AddOneCreated describes a response with status code 201, with default header values.

Created
*/
type AddOneCreated struct {
	Payload *models.Item
}

// IsSuccess returns true when this add one created response has a 2xx status code
func (o *AddOneCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add one created response has a 3xx status code
func (o *AddOneCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add one created response has a 4xx status code
func (o *AddOneCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add one created response has a 5xx status code
func (o *AddOneCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add one created response a status code equal to that given
func (o *AddOneCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the add one created response
func (o *AddOneCreated) Code() int {
	return 201
}

func (o *AddOneCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /][%d] addOneCreated %s", 201, payload)
}

func (o *AddOneCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /][%d] addOneCreated %s", 201, payload)
}

func (o *AddOneCreated) GetPayload() *models.Item {
	return o.Payload
}

func (o *AddOneCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Item)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOneDefault creates a AddOneDefault with default headers values
func NewAddOneDefault(code int) *AddOneDefault {
	return &AddOneDefault{
		_statusCode: code,
	}
}

/*
AddOneDefault describes a response with status code -1, with default header values.

error
*/
type AddOneDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this add one default response has a 2xx status code
func (o *AddOneDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add one default response has a 3xx status code
func (o *AddOneDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add one default response has a 4xx status code
func (o *AddOneDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add one default response has a 5xx status code
func (o *AddOneDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add one default response a status code equal to that given
func (o *AddOneDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the add one default response
func (o *AddOneDefault) Code() int {
	return o._statusCode
}

func (o *AddOneDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /][%d] addOne default %s", o._statusCode, payload)
}

func (o *AddOneDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /][%d] addOne default %s", o._statusCode, payload)
}

func (o *AddOneDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddOneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
