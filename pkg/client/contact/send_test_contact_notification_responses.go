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

// SendTestContactNotificationReader is a Reader for the SendTestContactNotification structure.
type SendTestContactNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendTestContactNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSendTestContactNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSendTestContactNotificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSendTestContactNotificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSendTestContactNotificationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /contact/{contactID}/test] send-test-contact-notification", response, response.Code())
	}
}

// NewSendTestContactNotificationOK creates a SendTestContactNotificationOK with default headers values
func NewSendTestContactNotificationOK() *SendTestContactNotificationOK {
	return &SendTestContactNotificationOK{}
}

/*
SendTestContactNotificationOK describes a response with status code 200, with default header values.

Test successful
*/
type SendTestContactNotificationOK struct {
}

// IsSuccess returns true when this send test contact notification o k response has a 2xx status code
func (o *SendTestContactNotificationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this send test contact notification o k response has a 3xx status code
func (o *SendTestContactNotificationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send test contact notification o k response has a 4xx status code
func (o *SendTestContactNotificationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this send test contact notification o k response has a 5xx status code
func (o *SendTestContactNotificationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this send test contact notification o k response a status code equal to that given
func (o *SendTestContactNotificationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the send test contact notification o k response
func (o *SendTestContactNotificationOK) Code() int {
	return 200
}

func (o *SendTestContactNotificationOK) Error() string {
	return fmt.Sprintf("[POST /contact/{contactID}/test][%d] sendTestContactNotificationOK", 200)
}

func (o *SendTestContactNotificationOK) String() string {
	return fmt.Sprintf("[POST /contact/{contactID}/test][%d] sendTestContactNotificationOK", 200)
}

func (o *SendTestContactNotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendTestContactNotificationForbidden creates a SendTestContactNotificationForbidden with default headers values
func NewSendTestContactNotificationForbidden() *SendTestContactNotificationForbidden {
	return &SendTestContactNotificationForbidden{}
}

/*
SendTestContactNotificationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SendTestContactNotificationForbidden struct {
	Payload *models.APIErrorForbiddenExample
}

// IsSuccess returns true when this send test contact notification forbidden response has a 2xx status code
func (o *SendTestContactNotificationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send test contact notification forbidden response has a 3xx status code
func (o *SendTestContactNotificationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send test contact notification forbidden response has a 4xx status code
func (o *SendTestContactNotificationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this send test contact notification forbidden response has a 5xx status code
func (o *SendTestContactNotificationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this send test contact notification forbidden response a status code equal to that given
func (o *SendTestContactNotificationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the send test contact notification forbidden response
func (o *SendTestContactNotificationForbidden) Code() int {
	return 403
}

func (o *SendTestContactNotificationForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /contact/{contactID}/test][%d] sendTestContactNotificationForbidden %s", 403, payload)
}

func (o *SendTestContactNotificationForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /contact/{contactID}/test][%d] sendTestContactNotificationForbidden %s", 403, payload)
}

func (o *SendTestContactNotificationForbidden) GetPayload() *models.APIErrorForbiddenExample {
	return o.Payload
}

func (o *SendTestContactNotificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorForbiddenExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendTestContactNotificationNotFound creates a SendTestContactNotificationNotFound with default headers values
func NewSendTestContactNotificationNotFound() *SendTestContactNotificationNotFound {
	return &SendTestContactNotificationNotFound{}
}

/*
SendTestContactNotificationNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type SendTestContactNotificationNotFound struct {
	Payload *models.APIErrorNotFoundExample
}

// IsSuccess returns true when this send test contact notification not found response has a 2xx status code
func (o *SendTestContactNotificationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send test contact notification not found response has a 3xx status code
func (o *SendTestContactNotificationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send test contact notification not found response has a 4xx status code
func (o *SendTestContactNotificationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this send test contact notification not found response has a 5xx status code
func (o *SendTestContactNotificationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this send test contact notification not found response a status code equal to that given
func (o *SendTestContactNotificationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the send test contact notification not found response
func (o *SendTestContactNotificationNotFound) Code() int {
	return 404
}

func (o *SendTestContactNotificationNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /contact/{contactID}/test][%d] sendTestContactNotificationNotFound %s", 404, payload)
}

func (o *SendTestContactNotificationNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /contact/{contactID}/test][%d] sendTestContactNotificationNotFound %s", 404, payload)
}

func (o *SendTestContactNotificationNotFound) GetPayload() *models.APIErrorNotFoundExample {
	return o.Payload
}

func (o *SendTestContactNotificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorNotFoundExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendTestContactNotificationInternalServerError creates a SendTestContactNotificationInternalServerError with default headers values
func NewSendTestContactNotificationInternalServerError() *SendTestContactNotificationInternalServerError {
	return &SendTestContactNotificationInternalServerError{}
}

/*
SendTestContactNotificationInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SendTestContactNotificationInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this send test contact notification internal server error response has a 2xx status code
func (o *SendTestContactNotificationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send test contact notification internal server error response has a 3xx status code
func (o *SendTestContactNotificationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send test contact notification internal server error response has a 4xx status code
func (o *SendTestContactNotificationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this send test contact notification internal server error response has a 5xx status code
func (o *SendTestContactNotificationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this send test contact notification internal server error response a status code equal to that given
func (o *SendTestContactNotificationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the send test contact notification internal server error response
func (o *SendTestContactNotificationInternalServerError) Code() int {
	return 500
}

func (o *SendTestContactNotificationInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /contact/{contactID}/test][%d] sendTestContactNotificationInternalServerError %s", 500, payload)
}

func (o *SendTestContactNotificationInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /contact/{contactID}/test][%d] sendTestContactNotificationInternalServerError %s", 500, payload)
}

func (o *SendTestContactNotificationInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *SendTestContactNotificationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
