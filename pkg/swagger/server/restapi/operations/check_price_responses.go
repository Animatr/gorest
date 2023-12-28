// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CheckPriceOKCode is the HTTP code returned for type CheckPriceOK
const CheckPriceOKCode int = 200

/*
CheckPriceOK OK message.

swagger:response checkPriceOK
*/
type CheckPriceOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCheckPriceOK creates CheckPriceOK with default headers values
func NewCheckPriceOK() *CheckPriceOK {

	return &CheckPriceOK{}
}

// WithPayload adds the payload to the check price o k response
func (o *CheckPriceOK) WithPayload(payload string) *CheckPriceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check price o k response
func (o *CheckPriceOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckPriceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}