// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetImageProductOKCode is the HTTP code returned for type GetImageProductOK
const GetImageProductOKCode int = 200

/*
GetImageProductOK Returns the product.

swagger:response getImageProductOK
*/
type GetImageProductOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetImageProductOK creates GetImageProductOK with default headers values
func NewGetImageProductOK() *GetImageProductOK {

	return &GetImageProductOK{}
}

// WithPayload adds the payload to the get image product o k response
func (o *GetImageProductOK) WithPayload(payload io.ReadCloser) *GetImageProductOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get image product o k response
func (o *GetImageProductOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImageProductOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
