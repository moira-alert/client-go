// Code generated by go-swagger; DO NOT EDIT.

package trigger

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

// UpdateTriggerReader is a Reader for the UpdateTrigger structure.
type UpdateTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTriggerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTriggerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateTriggerUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTriggerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateTriggerServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /trigger/{triggerID}] update-trigger", response, response.Code())
	}
}

// NewUpdateTriggerOK creates a UpdateTriggerOK with default headers values
func NewUpdateTriggerOK() *UpdateTriggerOK {
	return &UpdateTriggerOK{}
}

/*
UpdateTriggerOK describes a response with status code 200, with default header values.

Updated trigger
*/
type UpdateTriggerOK struct {
	Payload *models.DtoSaveTriggerResponse
}

// IsSuccess returns true when this update trigger o k response has a 2xx status code
func (o *UpdateTriggerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update trigger o k response has a 3xx status code
func (o *UpdateTriggerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger o k response has a 4xx status code
func (o *UpdateTriggerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trigger o k response has a 5xx status code
func (o *UpdateTriggerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update trigger o k response a status code equal to that given
func (o *UpdateTriggerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update trigger o k response
func (o *UpdateTriggerOK) Code() int {
	return 200
}

func (o *UpdateTriggerOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trigger/{triggerID}][%d] updateTriggerOK %s", 200, payload)
}

func (o *UpdateTriggerOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trigger/{triggerID}][%d] updateTriggerOK %s", 200, payload)
}

func (o *UpdateTriggerOK) GetPayload() *models.DtoSaveTriggerResponse {
	return o.Payload
}

func (o *UpdateTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoSaveTriggerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTriggerBadRequest creates a UpdateTriggerBadRequest with default headers values
func NewUpdateTriggerBadRequest() *UpdateTriggerBadRequest {
	return &UpdateTriggerBadRequest{}
}

/*
UpdateTriggerBadRequest describes a response with status code 400, with default header values.

Bad request from client
*/
type UpdateTriggerBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this update trigger bad request response has a 2xx status code
func (o *UpdateTriggerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trigger bad request response has a 3xx status code
func (o *UpdateTriggerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger bad request response has a 4xx status code
func (o *UpdateTriggerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trigger bad request response has a 5xx status code
func (o *UpdateTriggerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update trigger bad request response a status code equal to that given
func (o *UpdateTriggerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update trigger bad request response
func (o *UpdateTriggerBadRequest) Code() int {
	return 400
}

func (o *UpdateTriggerBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trigger/{triggerID}][%d] updateTriggerBadRequest %s", 400, payload)
}

func (o *UpdateTriggerBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trigger/{triggerID}][%d] updateTriggerBadRequest %s", 400, payload)
}

func (o *UpdateTriggerBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTriggerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTriggerNotFound creates a UpdateTriggerNotFound with default headers values
func NewUpdateTriggerNotFound() *UpdateTriggerNotFound {
	return &UpdateTriggerNotFound{}
}

/*
UpdateTriggerNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type UpdateTriggerNotFound struct {
	Payload *models.APIErrorNotFoundExample
}

// IsSuccess returns true when this update trigger not found response has a 2xx status code
func (o *UpdateTriggerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trigger not found response has a 3xx status code
func (o *UpdateTriggerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger not found response has a 4xx status code
func (o *UpdateTriggerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trigger not found response has a 5xx status code
func (o *UpdateTriggerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update trigger not found response a status code equal to that given
func (o *UpdateTriggerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update trigger not found response
func (o *UpdateTriggerNotFound) Code() int {
	return 404
}

func (o *UpdateTriggerNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trigger/{triggerID}][%d] updateTriggerNotFound %s", 404, payload)
}

func (o *UpdateTriggerNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trigger/{triggerID}][%d] updateTriggerNotFound %s", 404, payload)
}

func (o *UpdateTriggerNotFound) GetPayload() *models.APIErrorNotFoundExample {
	return o.Payload
}

func (o *UpdateTriggerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorNotFoundExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTriggerUnprocessableEntity creates a UpdateTriggerUnprocessableEntity with default headers values
func NewUpdateTriggerUnprocessableEntity() *UpdateTriggerUnprocessableEntity {
	return &UpdateTriggerUnprocessableEntity{}
}

/*
UpdateTriggerUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type UpdateTriggerUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this update trigger unprocessable entity response has a 2xx status code
func (o *UpdateTriggerUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trigger unprocessable entity response has a 3xx status code
func (o *UpdateTriggerUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger unprocessable entity response has a 4xx status code
func (o *UpdateTriggerUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trigger unprocessable entity response has a 5xx status code
func (o *UpdateTriggerUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update trigger unprocessable entity response a status code equal to that given
func (o *UpdateTriggerUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update trigger unprocessable entity response
func (o *UpdateTriggerUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateTriggerUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trigger/{triggerID}][%d] updateTriggerUnprocessableEntity %s", 422, payload)
}

func (o *UpdateTriggerUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trigger/{triggerID}][%d] updateTriggerUnprocessableEntity %s", 422, payload)
}

func (o *UpdateTriggerUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *UpdateTriggerUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTriggerInternalServerError creates a UpdateTriggerInternalServerError with default headers values
func NewUpdateTriggerInternalServerError() *UpdateTriggerInternalServerError {
	return &UpdateTriggerInternalServerError{}
}

/*
UpdateTriggerInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateTriggerInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this update trigger internal server error response has a 2xx status code
func (o *UpdateTriggerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trigger internal server error response has a 3xx status code
func (o *UpdateTriggerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger internal server error response has a 4xx status code
func (o *UpdateTriggerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trigger internal server error response has a 5xx status code
func (o *UpdateTriggerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update trigger internal server error response a status code equal to that given
func (o *UpdateTriggerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update trigger internal server error response
func (o *UpdateTriggerInternalServerError) Code() int {
	return 500
}

func (o *UpdateTriggerInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trigger/{triggerID}][%d] updateTriggerInternalServerError %s", 500, payload)
}

func (o *UpdateTriggerInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trigger/{triggerID}][%d] updateTriggerInternalServerError %s", 500, payload)
}

func (o *UpdateTriggerInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *UpdateTriggerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTriggerServiceUnavailable creates a UpdateTriggerServiceUnavailable with default headers values
func NewUpdateTriggerServiceUnavailable() *UpdateTriggerServiceUnavailable {
	return &UpdateTriggerServiceUnavailable{}
}

/*
UpdateTriggerServiceUnavailable describes a response with status code 503, with default header values.

Remote server unavailable
*/
type UpdateTriggerServiceUnavailable struct {
	Payload *models.APIErrorRemoteServerUnavailableExample
}

// IsSuccess returns true when this update trigger service unavailable response has a 2xx status code
func (o *UpdateTriggerServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trigger service unavailable response has a 3xx status code
func (o *UpdateTriggerServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger service unavailable response has a 4xx status code
func (o *UpdateTriggerServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trigger service unavailable response has a 5xx status code
func (o *UpdateTriggerServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this update trigger service unavailable response a status code equal to that given
func (o *UpdateTriggerServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the update trigger service unavailable response
func (o *UpdateTriggerServiceUnavailable) Code() int {
	return 503
}

func (o *UpdateTriggerServiceUnavailable) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trigger/{triggerID}][%d] updateTriggerServiceUnavailable %s", 503, payload)
}

func (o *UpdateTriggerServiceUnavailable) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trigger/{triggerID}][%d] updateTriggerServiceUnavailable %s", 503, payload)
}

func (o *UpdateTriggerServiceUnavailable) GetPayload() *models.APIErrorRemoteServerUnavailableExample {
	return o.Payload
}

func (o *UpdateTriggerServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRemoteServerUnavailableExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
