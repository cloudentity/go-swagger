// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/go-swagger/examples/task-tracker/models"
)

// ListTasksReader is a Reader for the ListTasks structure.
type ListTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewListTasksUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListTasksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListTasksOK creates a ListTasksOK with default headers values
func NewListTasksOK() *ListTasksOK {
	return &ListTasksOK{}
}

/*
ListTasksOK describes a response with status code 200, with default header values.

Successful response
*/
type ListTasksOK struct {

	/* the last task id known to the application

	   Format: int64
	*/
	XLastTaskID int64

	Payload []*models.TaskCard
}

// IsSuccess returns true when this list tasks o k response has a 2xx status code
func (o *ListTasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list tasks o k response has a 3xx status code
func (o *ListTasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tasks o k response has a 4xx status code
func (o *ListTasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list tasks o k response has a 5xx status code
func (o *ListTasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list tasks o k response a status code equal to that given
func (o *ListTasksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list tasks o k response
func (o *ListTasksOK) Code() int {
	return 200
}

func (o *ListTasksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tasks][%d] listTasksOK %s", 200, payload)
}

func (o *ListTasksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tasks][%d] listTasksOK %s", 200, payload)
}

func (o *ListTasksOK) GetPayload() []*models.TaskCard {
	return o.Payload
}

func (o *ListTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Last-Task-Id
	hdrXLastTaskID := response.GetHeader("X-Last-Task-Id")

	if hdrXLastTaskID != "" {
		valxLastTaskId, err := swag.ConvertInt64(hdrXLastTaskID)
		if err != nil {
			return errors.InvalidType("X-Last-Task-Id", "header", "int64", hdrXLastTaskID)
		}
		o.XLastTaskID = valxLastTaskId
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksUnprocessableEntity creates a ListTasksUnprocessableEntity with default headers values
func NewListTasksUnprocessableEntity() *ListTasksUnprocessableEntity {
	return &ListTasksUnprocessableEntity{}
}

/*
ListTasksUnprocessableEntity describes a response with status code 422, with default header values.

Validation error
*/
type ListTasksUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this list tasks unprocessable entity response has a 2xx status code
func (o *ListTasksUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list tasks unprocessable entity response has a 3xx status code
func (o *ListTasksUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tasks unprocessable entity response has a 4xx status code
func (o *ListTasksUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this list tasks unprocessable entity response has a 5xx status code
func (o *ListTasksUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this list tasks unprocessable entity response a status code equal to that given
func (o *ListTasksUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the list tasks unprocessable entity response
func (o *ListTasksUnprocessableEntity) Code() int {
	return 422
}

func (o *ListTasksUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tasks][%d] listTasksUnprocessableEntity %s", 422, payload)
}

func (o *ListTasksUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tasks][%d] listTasksUnprocessableEntity %s", 422, payload)
}

func (o *ListTasksUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *ListTasksUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksDefault creates a ListTasksDefault with default headers values
func NewListTasksDefault(code int) *ListTasksDefault {
	return &ListTasksDefault{
		_statusCode: code,
	}
}

/*
ListTasksDefault describes a response with status code -1, with default header values.

Error response
*/
type ListTasksDefault struct {
	_statusCode int
	XErrorCode  string

	Payload *models.Error
}

// IsSuccess returns true when this list tasks default response has a 2xx status code
func (o *ListTasksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list tasks default response has a 3xx status code
func (o *ListTasksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list tasks default response has a 4xx status code
func (o *ListTasksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list tasks default response has a 5xx status code
func (o *ListTasksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list tasks default response a status code equal to that given
func (o *ListTasksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list tasks default response
func (o *ListTasksDefault) Code() int {
	return o._statusCode
}

func (o *ListTasksDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tasks][%d] listTasks default %s", o._statusCode, payload)
}

func (o *ListTasksDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tasks][%d] listTasks default %s", o._statusCode, payload)
}

func (o *ListTasksDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListTasksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Error-Code
	hdrXErrorCode := response.GetHeader("X-Error-Code")

	if hdrXErrorCode != "" {
		o.XErrorCode = hdrXErrorCode
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
