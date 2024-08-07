// Code generated by go-swagger; DO NOT EDIT.

package contact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moira-alert/client-go/pkg/models"
)

// UpdateContactReader is a Reader for the UpdateContact structure.
type UpdateContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateContactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateContactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateContactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateContactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateContactUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateContactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /contact/{contactID}] update-contact", response, response.Code())
	}
}

// NewUpdateContactOK creates a UpdateContactOK with default headers values
func NewUpdateContactOK() *UpdateContactOK {
	return &UpdateContactOK{}
}

/*
UpdateContactOK describes a response with status code 200, with default header values.

Updated contact
*/
type UpdateContactOK struct {
	Payload *models.DtoContact
}

// IsSuccess returns true when this update contact o k response has a 2xx status code
func (o *UpdateContactOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update contact o k response has a 3xx status code
func (o *UpdateContactOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update contact o k response has a 4xx status code
func (o *UpdateContactOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update contact o k response has a 5xx status code
func (o *UpdateContactOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update contact o k response a status code equal to that given
func (o *UpdateContactOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update contact o k response
func (o *UpdateContactOK) Code() int {
	return 200
}

func (o *UpdateContactOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contact/{contactID}][%d] updateContactOK %s", 200, payload)
}

func (o *UpdateContactOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contact/{contactID}][%d] updateContactOK %s", 200, payload)
}

func (o *UpdateContactOK) GetPayload() *models.DtoContact {
	return o.Payload
}

func (o *UpdateContactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContactBadRequest creates a UpdateContactBadRequest with default headers values
func NewUpdateContactBadRequest() *UpdateContactBadRequest {
	return &UpdateContactBadRequest{}
}

/*
UpdateContactBadRequest describes a response with status code 400, with default header values.

Bad request from client
*/
type UpdateContactBadRequest struct {
	Payload *models.APIErrorInvalidRequestExample
}

// IsSuccess returns true when this update contact bad request response has a 2xx status code
func (o *UpdateContactBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update contact bad request response has a 3xx status code
func (o *UpdateContactBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update contact bad request response has a 4xx status code
func (o *UpdateContactBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update contact bad request response has a 5xx status code
func (o *UpdateContactBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update contact bad request response a status code equal to that given
func (o *UpdateContactBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update contact bad request response
func (o *UpdateContactBadRequest) Code() int {
	return 400
}

func (o *UpdateContactBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contact/{contactID}][%d] updateContactBadRequest %s", 400, payload)
}

func (o *UpdateContactBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contact/{contactID}][%d] updateContactBadRequest %s", 400, payload)
}

func (o *UpdateContactBadRequest) GetPayload() *models.APIErrorInvalidRequestExample {
	return o.Payload
}

func (o *UpdateContactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInvalidRequestExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContactForbidden creates a UpdateContactForbidden with default headers values
func NewUpdateContactForbidden() *UpdateContactForbidden {
	return &UpdateContactForbidden{}
}

/*
UpdateContactForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateContactForbidden struct {
	Payload *models.APIErrorForbiddenExample
}

// IsSuccess returns true when this update contact forbidden response has a 2xx status code
func (o *UpdateContactForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update contact forbidden response has a 3xx status code
func (o *UpdateContactForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update contact forbidden response has a 4xx status code
func (o *UpdateContactForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update contact forbidden response has a 5xx status code
func (o *UpdateContactForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update contact forbidden response a status code equal to that given
func (o *UpdateContactForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update contact forbidden response
func (o *UpdateContactForbidden) Code() int {
	return 403
}

func (o *UpdateContactForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contact/{contactID}][%d] updateContactForbidden %s", 403, payload)
}

func (o *UpdateContactForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contact/{contactID}][%d] updateContactForbidden %s", 403, payload)
}

func (o *UpdateContactForbidden) GetPayload() *models.APIErrorForbiddenExample {
	return o.Payload
}

func (o *UpdateContactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorForbiddenExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContactNotFound creates a UpdateContactNotFound with default headers values
func NewUpdateContactNotFound() *UpdateContactNotFound {
	return &UpdateContactNotFound{}
}

/*
UpdateContactNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type UpdateContactNotFound struct {
	Payload *models.APIErrorNotFoundExample
}

// IsSuccess returns true when this update contact not found response has a 2xx status code
func (o *UpdateContactNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update contact not found response has a 3xx status code
func (o *UpdateContactNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update contact not found response has a 4xx status code
func (o *UpdateContactNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update contact not found response has a 5xx status code
func (o *UpdateContactNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update contact not found response a status code equal to that given
func (o *UpdateContactNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update contact not found response
func (o *UpdateContactNotFound) Code() int {
	return 404
}

func (o *UpdateContactNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contact/{contactID}][%d] updateContactNotFound %s", 404, payload)
}

func (o *UpdateContactNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contact/{contactID}][%d] updateContactNotFound %s", 404, payload)
}

func (o *UpdateContactNotFound) GetPayload() *models.APIErrorNotFoundExample {
	return o.Payload
}

func (o *UpdateContactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorNotFoundExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContactUnprocessableEntity creates a UpdateContactUnprocessableEntity with default headers values
func NewUpdateContactUnprocessableEntity() *UpdateContactUnprocessableEntity {
	return &UpdateContactUnprocessableEntity{}
}

/*
UpdateContactUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type UpdateContactUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this update contact unprocessable entity response has a 2xx status code
func (o *UpdateContactUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update contact unprocessable entity response has a 3xx status code
func (o *UpdateContactUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update contact unprocessable entity response has a 4xx status code
func (o *UpdateContactUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update contact unprocessable entity response has a 5xx status code
func (o *UpdateContactUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update contact unprocessable entity response a status code equal to that given
func (o *UpdateContactUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update contact unprocessable entity response
func (o *UpdateContactUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateContactUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contact/{contactID}][%d] updateContactUnprocessableEntity %s", 422, payload)
}

func (o *UpdateContactUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contact/{contactID}][%d] updateContactUnprocessableEntity %s", 422, payload)
}

func (o *UpdateContactUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *UpdateContactUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateContactInternalServerError creates a UpdateContactInternalServerError with default headers values
func NewUpdateContactInternalServerError() *UpdateContactInternalServerError {
	return &UpdateContactInternalServerError{}
}

/*
UpdateContactInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateContactInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this update contact internal server error response has a 2xx status code
func (o *UpdateContactInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update contact internal server error response has a 3xx status code
func (o *UpdateContactInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update contact internal server error response has a 4xx status code
func (o *UpdateContactInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update contact internal server error response has a 5xx status code
func (o *UpdateContactInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update contact internal server error response a status code equal to that given
func (o *UpdateContactInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update contact internal server error response
func (o *UpdateContactInternalServerError) Code() int {
	return 500
}

func (o *UpdateContactInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contact/{contactID}][%d] updateContactInternalServerError %s", 500, payload)
}

func (o *UpdateContactInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /contact/{contactID}][%d] updateContactInternalServerError %s", 500, payload)
}

func (o *UpdateContactInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *UpdateContactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
