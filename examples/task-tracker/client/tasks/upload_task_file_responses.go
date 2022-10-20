// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/go-swagger/examples/task-tracker/models"
)

// UploadTaskFileReader is a Reader for the UploadTaskFile structure.
type UploadTaskFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadTaskFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUploadTaskFileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUploadTaskFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUploadTaskFileCreated creates a UploadTaskFileCreated with default headers values
func NewUploadTaskFileCreated() *UploadTaskFileCreated {
	return &UploadTaskFileCreated{}
}

/*
	UploadTaskFileCreated describes a response with status code 201, with default header values.

File added
*/
type UploadTaskFileCreated struct {
}

// IsSuccess returns true when this upload task file created response returns a 2xx status code
func (o *UploadTaskFileCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload task file created response returns a 3xx status code
func (o *UploadTaskFileCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload task file created response returns a 4xx status code
func (o *UploadTaskFileCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload task file created response returns a 5xx status code
func (o *UploadTaskFileCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this upload task file created response returns a 4xx status code
func (o *UploadTaskFileCreated) IsCode(code int) bool {
	return code == 201
}

func (o *UploadTaskFileCreated) Error() string {
	return fmt.Sprintf("[POST /tasks/{id}/files][%d] uploadTaskFileCreated ", 201)
}

func (o *UploadTaskFileCreated) String() string {
	return fmt.Sprintf("[POST /tasks/{id}/files][%d] uploadTaskFileCreated ", 201)
}

func (o *UploadTaskFileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadTaskFileDefault creates a UploadTaskFileDefault with default headers values
func NewUploadTaskFileDefault(code int) *UploadTaskFileDefault {
	return &UploadTaskFileDefault{
		_statusCode: code,
	}
}

/*
	UploadTaskFileDefault describes a response with status code -1, with default header values.

Error response
*/
type UploadTaskFileDefault struct {
	_statusCode int
	XErrorCode  string

	Payload *models.Error
}

// Code gets the status code for the upload task file default response
func (o *UploadTaskFileDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this upload task file default response returns a 2xx status code
func (o *UploadTaskFileDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this upload task file default response returns a 3xx status code
func (o *UploadTaskFileDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this upload task file default response returns a 4xx status code
func (o *UploadTaskFileDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this upload task file default response returns a 5xx status code
func (o *UploadTaskFileDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this upload task file default response returns a 4xx status code
func (o *UploadTaskFileDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UploadTaskFileDefault) Error() string {
	return fmt.Sprintf("[POST /tasks/{id}/files][%d] uploadTaskFile default  %+v", o._statusCode, o.Payload)
}

func (o *UploadTaskFileDefault) String() string {
	return fmt.Sprintf("[POST /tasks/{id}/files][%d] uploadTaskFile default  %+v", o._statusCode, o.Payload)
}

func (o *UploadTaskFileDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadTaskFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
