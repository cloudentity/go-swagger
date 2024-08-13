// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloudentity/go-swagger/examples/contributed-templates/stratoscale/models"
)

// PetListOKCode is the HTTP code returned for type PetListOK
const PetListOKCode int = 200

/*
PetListOK successful operation

swagger:response petListOK
*/
type PetListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Pet `json:"body,omitempty"`
}

// NewPetListOK creates PetListOK with default headers values
func NewPetListOK() *PetListOK {

	return &PetListOK{}
}

// WithPayload adds the payload to the pet list o k response
func (o *PetListOK) WithPayload(payload []*models.Pet) *PetListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pet list o k response
func (o *PetListOK) SetPayload(payload []*models.Pet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PetListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Pet, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PetListBadRequestCode is the HTTP code returned for type PetListBadRequest
const PetListBadRequestCode int = 400

/*
PetListBadRequest Invalid status value

swagger:response petListBadRequest
*/
type PetListBadRequest struct {
}

// NewPetListBadRequest creates PetListBadRequest with default headers values
func NewPetListBadRequest() *PetListBadRequest {

	return &PetListBadRequest{}
}

// WriteResponse to the client
func (o *PetListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
