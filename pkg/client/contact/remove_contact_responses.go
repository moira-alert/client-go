// Code generated by go-swagger; DO NOT EDIT.

package contact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moira-alert/client-go/pkg/models"
)

// RemoveContactReader is a Reader for the RemoveContact structure.
type RemoveContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveContactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveContactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveContactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveContactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveContactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /contact/{contactID}] remove-contact", response, response.Code())
	}
}

// NewRemoveContactOK creates a RemoveContactOK with default headers values
func NewRemoveContactOK() *RemoveContactOK {
	return &RemoveContactOK{}
}

/*
RemoveContactOK describes a response with status code 200, with default header values.

Contact has been deleted
*/
type RemoveContactOK struct {
}

// IsSuccess returns true when this remove contact o k response has a 2xx status code
func (o *RemoveContactOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove contact o k response has a 3xx status code
func (o *RemoveContactOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove contact o k response has a 4xx status code
func (o *RemoveContactOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove contact o k response has a 5xx status code
func (o *RemoveContactOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove contact o k response a status code equal to that given
func (o *RemoveContactOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove contact o k response
func (o *RemoveContactOK) Code() int {
	return 200
}

func (o *RemoveContactOK) Error() string {
	return fmt.Sprintf("[DELETE /contact/{contactID}][%d] removeContactOK ", 200)
}

func (o *RemoveContactOK) String() string {
	return fmt.Sprintf("[DELETE /contact/{contactID}][%d] removeContactOK ", 200)
}

func (o *RemoveContactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveContactBadRequest creates a RemoveContactBadRequest with default headers values
func NewRemoveContactBadRequest() *RemoveContactBadRequest {
	return &RemoveContactBadRequest{}
}

/*
RemoveContactBadRequest describes a response with status code 400, with default header values.

Bad request from client
*/
type RemoveContactBadRequest struct {
	Payload *models.APIErrorInvalidRequestExample
}

// IsSuccess returns true when this remove contact bad request response has a 2xx status code
func (o *RemoveContactBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove contact bad request response has a 3xx status code
func (o *RemoveContactBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove contact bad request response has a 4xx status code
func (o *RemoveContactBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove contact bad request response has a 5xx status code
func (o *RemoveContactBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove contact bad request response a status code equal to that given
func (o *RemoveContactBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove contact bad request response
func (o *RemoveContactBadRequest) Code() int {
	return 400
}

func (o *RemoveContactBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /contact/{contactID}][%d] removeContactBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveContactBadRequest) String() string {
	return fmt.Sprintf("[DELETE /contact/{contactID}][%d] removeContactBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveContactBadRequest) GetPayload() *models.APIErrorInvalidRequestExample {
	return o.Payload
}

func (o *RemoveContactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInvalidRequestExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveContactForbidden creates a RemoveContactForbidden with default headers values
func NewRemoveContactForbidden() *RemoveContactForbidden {
	return &RemoveContactForbidden{}
}

/*
RemoveContactForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RemoveContactForbidden struct {
	Payload *models.APIErrorForbiddenExample
}

// IsSuccess returns true when this remove contact forbidden response has a 2xx status code
func (o *RemoveContactForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove contact forbidden response has a 3xx status code
func (o *RemoveContactForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove contact forbidden response has a 4xx status code
func (o *RemoveContactForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove contact forbidden response has a 5xx status code
func (o *RemoveContactForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this remove contact forbidden response a status code equal to that given
func (o *RemoveContactForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the remove contact forbidden response
func (o *RemoveContactForbidden) Code() int {
	return 403
}

func (o *RemoveContactForbidden) Error() string {
	return fmt.Sprintf("[DELETE /contact/{contactID}][%d] removeContactForbidden  %+v", 403, o.Payload)
}

func (o *RemoveContactForbidden) String() string {
	return fmt.Sprintf("[DELETE /contact/{contactID}][%d] removeContactForbidden  %+v", 403, o.Payload)
}

func (o *RemoveContactForbidden) GetPayload() *models.APIErrorForbiddenExample {
	return o.Payload
}

func (o *RemoveContactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorForbiddenExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveContactNotFound creates a RemoveContactNotFound with default headers values
func NewRemoveContactNotFound() *RemoveContactNotFound {
	return &RemoveContactNotFound{}
}

/*
RemoveContactNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type RemoveContactNotFound struct {
	Payload *models.APIErrorNotFoundExample
}

// IsSuccess returns true when this remove contact not found response has a 2xx status code
func (o *RemoveContactNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove contact not found response has a 3xx status code
func (o *RemoveContactNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove contact not found response has a 4xx status code
func (o *RemoveContactNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove contact not found response has a 5xx status code
func (o *RemoveContactNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this remove contact not found response a status code equal to that given
func (o *RemoveContactNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the remove contact not found response
func (o *RemoveContactNotFound) Code() int {
	return 404
}

func (o *RemoveContactNotFound) Error() string {
	return fmt.Sprintf("[DELETE /contact/{contactID}][%d] removeContactNotFound  %+v", 404, o.Payload)
}

func (o *RemoveContactNotFound) String() string {
	return fmt.Sprintf("[DELETE /contact/{contactID}][%d] removeContactNotFound  %+v", 404, o.Payload)
}

func (o *RemoveContactNotFound) GetPayload() *models.APIErrorNotFoundExample {
	return o.Payload
}

func (o *RemoveContactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorNotFoundExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveContactInternalServerError creates a RemoveContactInternalServerError with default headers values
func NewRemoveContactInternalServerError() *RemoveContactInternalServerError {
	return &RemoveContactInternalServerError{}
}

/*
RemoveContactInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type RemoveContactInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this remove contact internal server error response has a 2xx status code
func (o *RemoveContactInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove contact internal server error response has a 3xx status code
func (o *RemoveContactInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove contact internal server error response has a 4xx status code
func (o *RemoveContactInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove contact internal server error response has a 5xx status code
func (o *RemoveContactInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove contact internal server error response a status code equal to that given
func (o *RemoveContactInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove contact internal server error response
func (o *RemoveContactInternalServerError) Code() int {
	return 500
}

func (o *RemoveContactInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /contact/{contactID}][%d] removeContactInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveContactInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /contact/{contactID}][%d] removeContactInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveContactInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *RemoveContactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}