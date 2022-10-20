// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/go-swagger/examples/authentication/models"
)

// GetIDReader is a Reader for the GetID structure.
type GetIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIDOK creates a GetIDOK with default headers values
func NewGetIDOK() *GetIDOK {
	return &GetIDOK{}
}

/*
	GetIDOK describes a response with status code 200, with default header values.

OK
*/
type GetIDOK struct {
	Payload *models.Customer
}

// IsSuccess returns true when this get Id o k response returns a 2xx status code
func (o *GetIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Id o k response returns a 3xx status code
func (o *GetIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Id o k response returns a 4xx status code
func (o *GetIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Id o k response returns a 5xx status code
func (o *GetIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Id o k response returns a 4xx status code
func (o *GetIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIDOK) Error() string {
	return fmt.Sprintf("[GET /customers][%d] getIdOK  %+v", 200, o.Payload)
}

func (o *GetIDOK) String() string {
	return fmt.Sprintf("[GET /customers][%d] getIdOK  %+v", 200, o.Payload)
}

func (o *GetIDOK) GetPayload() *models.Customer {
	return o.Payload
}

func (o *GetIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIDUnauthorized creates a GetIDUnauthorized with default headers values
func NewGetIDUnauthorized() *GetIDUnauthorized {
	return &GetIDUnauthorized{}
}

/*
	GetIDUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type GetIDUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Id unauthorized response returns a 2xx status code
func (o *GetIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Id unauthorized response returns a 3xx status code
func (o *GetIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Id unauthorized response returns a 4xx status code
func (o *GetIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Id unauthorized response returns a 5xx status code
func (o *GetIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get Id unauthorized response returns a 4xx status code
func (o *GetIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /customers][%d] getIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /customers][%d] getIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIDNotFound creates a GetIDNotFound with default headers values
func NewGetIDNotFound() *GetIDNotFound {
	return &GetIDNotFound{}
}

/*
	GetIDNotFound describes a response with status code 404, with default header values.

resource not found
*/
type GetIDNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get Id not found response returns a 2xx status code
func (o *GetIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Id not found response returns a 3xx status code
func (o *GetIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Id not found response returns a 4xx status code
func (o *GetIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Id not found response returns a 5xx status code
func (o *GetIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Id not found response returns a 4xx status code
func (o *GetIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetIDNotFound) Error() string {
	return fmt.Sprintf("[GET /customers][%d] getIdNotFound  %+v", 404, o.Payload)
}

func (o *GetIDNotFound) String() string {
	return fmt.Sprintf("[GET /customers][%d] getIdNotFound  %+v", 404, o.Payload)
}

func (o *GetIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIDDefault creates a GetIDDefault with default headers values
func NewGetIDDefault(code int) *GetIDDefault {
	return &GetIDDefault{
		_statusCode: code,
	}
}

/*
	GetIDDefault describes a response with status code -1, with default header values.

error
*/
type GetIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get Id default response
func (o *GetIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get Id default response returns a 2xx status code
func (o *GetIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get Id default response returns a 3xx status code
func (o *GetIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get Id default response returns a 4xx status code
func (o *GetIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get Id default response returns a 5xx status code
func (o *GetIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get Id default response returns a 4xx status code
func (o *GetIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetIDDefault) Error() string {
	return fmt.Sprintf("[GET /customers][%d] getId default  %+v", o._statusCode, o.Payload)
}

func (o *GetIDDefault) String() string {
	return fmt.Sprintf("[GET /customers][%d] getId default  %+v", o._statusCode, o.Payload)
}

func (o *GetIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
