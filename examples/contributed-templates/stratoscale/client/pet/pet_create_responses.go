// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/go-swagger/examples/contributed-templates/stratoscale/models"
)

// PetCreateReader is a Reader for the PetCreate structure.
type PetCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PetCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPetCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 405:
		result := NewPetCreateMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pet] PetCreate", response, response.Code())
	}
}

// NewPetCreateCreated creates a PetCreateCreated with default headers values
func NewPetCreateCreated() *PetCreateCreated {
	return &PetCreateCreated{}
}

/*
PetCreateCreated describes a response with status code 201, with default header values.

Created
*/
type PetCreateCreated struct {
	Payload *models.Pet
}

// IsSuccess returns true when this pet create created response has a 2xx status code
func (o *PetCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pet create created response has a 3xx status code
func (o *PetCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pet create created response has a 4xx status code
func (o *PetCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this pet create created response has a 5xx status code
func (o *PetCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this pet create created response a status code equal to that given
func (o *PetCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the pet create created response
func (o *PetCreateCreated) Code() int {
	return 201
}

func (o *PetCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pet][%d] petCreateCreated %s", 201, payload)
}

func (o *PetCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pet][%d] petCreateCreated %s", 201, payload)
}

func (o *PetCreateCreated) GetPayload() *models.Pet {
	return o.Payload
}

func (o *PetCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPetCreateMethodNotAllowed creates a PetCreateMethodNotAllowed with default headers values
func NewPetCreateMethodNotAllowed() *PetCreateMethodNotAllowed {
	return &PetCreateMethodNotAllowed{}
}

/*
PetCreateMethodNotAllowed describes a response with status code 405, with default header values.

Invalid input
*/
type PetCreateMethodNotAllowed struct {
}

// IsSuccess returns true when this pet create method not allowed response has a 2xx status code
func (o *PetCreateMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pet create method not allowed response has a 3xx status code
func (o *PetCreateMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pet create method not allowed response has a 4xx status code
func (o *PetCreateMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this pet create method not allowed response has a 5xx status code
func (o *PetCreateMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this pet create method not allowed response a status code equal to that given
func (o *PetCreateMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

// Code gets the status code for the pet create method not allowed response
func (o *PetCreateMethodNotAllowed) Code() int {
	return 405
}

func (o *PetCreateMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /pet][%d] petCreateMethodNotAllowed", 405)
}

func (o *PetCreateMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /pet][%d] petCreateMethodNotAllowed", 405)
}

func (o *PetCreateMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
