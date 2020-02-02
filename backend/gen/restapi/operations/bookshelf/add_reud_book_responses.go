// Code generated by go-swagger; DO NOT EDIT.

package bookshelf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddReudBookNoContentCode is the HTTP code returned for type AddReudBookNoContent
const AddReudBookNoContentCode int = 204

/*AddReudBookNoContent No Content

swagger:response addReudBookNoContent
*/
type AddReudBookNoContent struct {
}

// NewAddReudBookNoContent creates AddReudBookNoContent with default headers values
func NewAddReudBookNoContent() *AddReudBookNoContent {

	return &AddReudBookNoContent{}
}

// WriteResponse to the client
func (o *AddReudBookNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// AddReudBookBadRequestCode is the HTTP code returned for type AddReudBookBadRequest
const AddReudBookBadRequestCode int = 400

/*AddReudBookBadRequest Bad Request

swagger:response addReudBookBadRequest
*/
type AddReudBookBadRequest struct {
}

// NewAddReudBookBadRequest creates AddReudBookBadRequest with default headers values
func NewAddReudBookBadRequest() *AddReudBookBadRequest {

	return &AddReudBookBadRequest{}
}

// WriteResponse to the client
func (o *AddReudBookBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// AddReudBookUnauthorizedCode is the HTTP code returned for type AddReudBookUnauthorized
const AddReudBookUnauthorizedCode int = 401

/*AddReudBookUnauthorized Unauthorized

swagger:response addReudBookUnauthorized
*/
type AddReudBookUnauthorized struct {
}

// NewAddReudBookUnauthorized creates AddReudBookUnauthorized with default headers values
func NewAddReudBookUnauthorized() *AddReudBookUnauthorized {

	return &AddReudBookUnauthorized{}
}

// WriteResponse to the client
func (o *AddReudBookUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// AddReudBookInternalServerErrorCode is the HTTP code returned for type AddReudBookInternalServerError
const AddReudBookInternalServerErrorCode int = 500

/*AddReudBookInternalServerError Internal Server Error

swagger:response addReudBookInternalServerError
*/
type AddReudBookInternalServerError struct {
}

// NewAddReudBookInternalServerError creates AddReudBookInternalServerError with default headers values
func NewAddReudBookInternalServerError() *AddReudBookInternalServerError {

	return &AddReudBookInternalServerError{}
}

// WriteResponse to the client
func (o *AddReudBookInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}