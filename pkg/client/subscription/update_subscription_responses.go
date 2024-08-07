// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// UpdateSubscriptionReader is a Reader for the UpdateSubscription structure.
type UpdateSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateSubscriptionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSubscriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateSubscriptionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSubscriptionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /subscription/{subscriptionID}] update-subscription", response, response.Code())
	}
}

// NewUpdateSubscriptionOK creates a UpdateSubscriptionOK with default headers values
func NewUpdateSubscriptionOK() *UpdateSubscriptionOK {
	return &UpdateSubscriptionOK{}
}

/*
UpdateSubscriptionOK describes a response with status code 200, with default header values.

Subscription updated successfully
*/
type UpdateSubscriptionOK struct {
	Payload *models.DtoSubscription
}

// IsSuccess returns true when this update subscription o k response has a 2xx status code
func (o *UpdateSubscriptionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update subscription o k response has a 3xx status code
func (o *UpdateSubscriptionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update subscription o k response has a 4xx status code
func (o *UpdateSubscriptionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update subscription o k response has a 5xx status code
func (o *UpdateSubscriptionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update subscription o k response a status code equal to that given
func (o *UpdateSubscriptionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update subscription o k response
func (o *UpdateSubscriptionOK) Code() int {
	return 200
}

func (o *UpdateSubscriptionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /subscription/{subscriptionID}][%d] updateSubscriptionOK %s", 200, payload)
}

func (o *UpdateSubscriptionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /subscription/{subscriptionID}][%d] updateSubscriptionOK %s", 200, payload)
}

func (o *UpdateSubscriptionOK) GetPayload() *models.DtoSubscription {
	return o.Payload
}

func (o *UpdateSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoSubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubscriptionBadRequest creates a UpdateSubscriptionBadRequest with default headers values
func NewUpdateSubscriptionBadRequest() *UpdateSubscriptionBadRequest {
	return &UpdateSubscriptionBadRequest{}
}

/*
UpdateSubscriptionBadRequest describes a response with status code 400, with default header values.

Bad request from client
*/
type UpdateSubscriptionBadRequest struct {
	Payload *models.APIErrorInvalidRequestExample
}

// IsSuccess returns true when this update subscription bad request response has a 2xx status code
func (o *UpdateSubscriptionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update subscription bad request response has a 3xx status code
func (o *UpdateSubscriptionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update subscription bad request response has a 4xx status code
func (o *UpdateSubscriptionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update subscription bad request response has a 5xx status code
func (o *UpdateSubscriptionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update subscription bad request response a status code equal to that given
func (o *UpdateSubscriptionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update subscription bad request response
func (o *UpdateSubscriptionBadRequest) Code() int {
	return 400
}

func (o *UpdateSubscriptionBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /subscription/{subscriptionID}][%d] updateSubscriptionBadRequest %s", 400, payload)
}

func (o *UpdateSubscriptionBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /subscription/{subscriptionID}][%d] updateSubscriptionBadRequest %s", 400, payload)
}

func (o *UpdateSubscriptionBadRequest) GetPayload() *models.APIErrorInvalidRequestExample {
	return o.Payload
}

func (o *UpdateSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInvalidRequestExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubscriptionForbidden creates a UpdateSubscriptionForbidden with default headers values
func NewUpdateSubscriptionForbidden() *UpdateSubscriptionForbidden {
	return &UpdateSubscriptionForbidden{}
}

/*
UpdateSubscriptionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateSubscriptionForbidden struct {
	Payload *models.APIErrorForbiddenExample
}

// IsSuccess returns true when this update subscription forbidden response has a 2xx status code
func (o *UpdateSubscriptionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update subscription forbidden response has a 3xx status code
func (o *UpdateSubscriptionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update subscription forbidden response has a 4xx status code
func (o *UpdateSubscriptionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update subscription forbidden response has a 5xx status code
func (o *UpdateSubscriptionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update subscription forbidden response a status code equal to that given
func (o *UpdateSubscriptionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update subscription forbidden response
func (o *UpdateSubscriptionForbidden) Code() int {
	return 403
}

func (o *UpdateSubscriptionForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /subscription/{subscriptionID}][%d] updateSubscriptionForbidden %s", 403, payload)
}

func (o *UpdateSubscriptionForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /subscription/{subscriptionID}][%d] updateSubscriptionForbidden %s", 403, payload)
}

func (o *UpdateSubscriptionForbidden) GetPayload() *models.APIErrorForbiddenExample {
	return o.Payload
}

func (o *UpdateSubscriptionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorForbiddenExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubscriptionNotFound creates a UpdateSubscriptionNotFound with default headers values
func NewUpdateSubscriptionNotFound() *UpdateSubscriptionNotFound {
	return &UpdateSubscriptionNotFound{}
}

/*
UpdateSubscriptionNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type UpdateSubscriptionNotFound struct {
	Payload *models.APIErrorNotFoundExample
}

// IsSuccess returns true when this update subscription not found response has a 2xx status code
func (o *UpdateSubscriptionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update subscription not found response has a 3xx status code
func (o *UpdateSubscriptionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update subscription not found response has a 4xx status code
func (o *UpdateSubscriptionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update subscription not found response has a 5xx status code
func (o *UpdateSubscriptionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update subscription not found response a status code equal to that given
func (o *UpdateSubscriptionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update subscription not found response
func (o *UpdateSubscriptionNotFound) Code() int {
	return 404
}

func (o *UpdateSubscriptionNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /subscription/{subscriptionID}][%d] updateSubscriptionNotFound %s", 404, payload)
}

func (o *UpdateSubscriptionNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /subscription/{subscriptionID}][%d] updateSubscriptionNotFound %s", 404, payload)
}

func (o *UpdateSubscriptionNotFound) GetPayload() *models.APIErrorNotFoundExample {
	return o.Payload
}

func (o *UpdateSubscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorNotFoundExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubscriptionUnprocessableEntity creates a UpdateSubscriptionUnprocessableEntity with default headers values
func NewUpdateSubscriptionUnprocessableEntity() *UpdateSubscriptionUnprocessableEntity {
	return &UpdateSubscriptionUnprocessableEntity{}
}

/*
UpdateSubscriptionUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type UpdateSubscriptionUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this update subscription unprocessable entity response has a 2xx status code
func (o *UpdateSubscriptionUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update subscription unprocessable entity response has a 3xx status code
func (o *UpdateSubscriptionUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update subscription unprocessable entity response has a 4xx status code
func (o *UpdateSubscriptionUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update subscription unprocessable entity response has a 5xx status code
func (o *UpdateSubscriptionUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update subscription unprocessable entity response a status code equal to that given
func (o *UpdateSubscriptionUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update subscription unprocessable entity response
func (o *UpdateSubscriptionUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateSubscriptionUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /subscription/{subscriptionID}][%d] updateSubscriptionUnprocessableEntity %s", 422, payload)
}

func (o *UpdateSubscriptionUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /subscription/{subscriptionID}][%d] updateSubscriptionUnprocessableEntity %s", 422, payload)
}

func (o *UpdateSubscriptionUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *UpdateSubscriptionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSubscriptionInternalServerError creates a UpdateSubscriptionInternalServerError with default headers values
func NewUpdateSubscriptionInternalServerError() *UpdateSubscriptionInternalServerError {
	return &UpdateSubscriptionInternalServerError{}
}

/*
UpdateSubscriptionInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateSubscriptionInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this update subscription internal server error response has a 2xx status code
func (o *UpdateSubscriptionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update subscription internal server error response has a 3xx status code
func (o *UpdateSubscriptionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update subscription internal server error response has a 4xx status code
func (o *UpdateSubscriptionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update subscription internal server error response has a 5xx status code
func (o *UpdateSubscriptionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update subscription internal server error response a status code equal to that given
func (o *UpdateSubscriptionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update subscription internal server error response
func (o *UpdateSubscriptionInternalServerError) Code() int {
	return 500
}

func (o *UpdateSubscriptionInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /subscription/{subscriptionID}][%d] updateSubscriptionInternalServerError %s", 500, payload)
}

func (o *UpdateSubscriptionInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /subscription/{subscriptionID}][%d] updateSubscriptionInternalServerError %s", 500, payload)
}

func (o *UpdateSubscriptionInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *UpdateSubscriptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
