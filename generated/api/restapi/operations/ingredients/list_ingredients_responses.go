// Code generated by go-swagger; DO NOT EDIT.

package ingredients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/BlindGarret/project-nommies/generated/api/models"
)

// ListIngredientsOKCode is the HTTP code returned for type ListIngredientsOK
const ListIngredientsOKCode int = 200

/*ListIngredientsOK Ingredient response object containing list of ingredients

swagger:response listIngredientsOK
*/
type ListIngredientsOK struct {

	/*
	  In: Body
	*/
	Payload *ListIngredientsOKBody `json:"body,omitempty"`
}

// NewListIngredientsOK creates ListIngredientsOK with default headers values
func NewListIngredientsOK() *ListIngredientsOK {

	return &ListIngredientsOK{}
}

// WithPayload adds the payload to the list ingredients o k response
func (o *ListIngredientsOK) WithPayload(payload *ListIngredientsOKBody) *ListIngredientsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list ingredients o k response
func (o *ListIngredientsOK) SetPayload(payload *ListIngredientsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIngredientsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListIngredientsDefault unexpected error

swagger:response listIngredientsDefault
*/
type ListIngredientsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListIngredientsDefault creates ListIngredientsDefault with default headers values
func NewListIngredientsDefault(code int) *ListIngredientsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListIngredientsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list ingredients default response
func (o *ListIngredientsDefault) WithStatusCode(code int) *ListIngredientsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list ingredients default response
func (o *ListIngredientsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list ingredients default response
func (o *ListIngredientsDefault) WithPayload(payload *models.Error) *ListIngredientsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list ingredients default response
func (o *ListIngredientsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIngredientsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
