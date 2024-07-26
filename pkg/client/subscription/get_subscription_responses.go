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

// GetSubscriptionReader is a Reader for the GetSubscription structure.
type GetSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSubscriptionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetSubscriptionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSubscriptionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /subscription/{subscriptionID}] get-subscription", response, response.Code())
	}
}

// NewGetSubscriptionOK creates a GetSubscriptionOK with default headers values
func NewGetSubscriptionOK() *GetSubscriptionOK {
	return &GetSubscriptionOK{}
}

/*
GetSubscriptionOK describes a response with status code 200, with default header values.

Subscription fetched successfully
*/
type GetSubscriptionOK struct {
	Payload *models.DtoSubscription
}

// IsSuccess returns true when this get subscription o k response has a 2xx status code
func (o *GetSubscriptionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get subscription o k response has a 3xx status code
func (o *GetSubscriptionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subscription o k response has a 4xx status code
func (o *GetSubscriptionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get subscription o k response has a 5xx status code
func (o *GetSubscriptionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get subscription o k response a status code equal to that given
func (o *GetSubscriptionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get subscription o k response
func (o *GetSubscriptionOK) Code() int {
	return 200
}

func (o *GetSubscriptionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /subscription/{subscriptionID}][%d] getSubscriptionOK %s", 200, payload)
}

func (o *GetSubscriptionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /subscription/{subscriptionID}][%d] getSubscriptionOK %s", 200, payload)
}

func (o *GetSubscriptionOK) GetPayload() *models.DtoSubscription {
	return o.Payload
}

func (o *GetSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoSubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionForbidden creates a GetSubscriptionForbidden with default headers values
func NewGetSubscriptionForbidden() *GetSubscriptionForbidden {
	return &GetSubscriptionForbidden{}
}

/*
GetSubscriptionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSubscriptionForbidden struct {
	Payload *models.APIErrorForbiddenExample
}

// IsSuccess returns true when this get subscription forbidden response has a 2xx status code
func (o *GetSubscriptionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get subscription forbidden response has a 3xx status code
func (o *GetSubscriptionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subscription forbidden response has a 4xx status code
func (o *GetSubscriptionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get subscription forbidden response has a 5xx status code
func (o *GetSubscriptionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get subscription forbidden response a status code equal to that given
func (o *GetSubscriptionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get subscription forbidden response
func (o *GetSubscriptionForbidden) Code() int {
	return 403
}

func (o *GetSubscriptionForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /subscription/{subscriptionID}][%d] getSubscriptionForbidden %s", 403, payload)
}

func (o *GetSubscriptionForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /subscription/{subscriptionID}][%d] getSubscriptionForbidden %s", 403, payload)
}

func (o *GetSubscriptionForbidden) GetPayload() *models.APIErrorForbiddenExample {
	return o.Payload
}

func (o *GetSubscriptionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorForbiddenExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionUnprocessableEntity creates a GetSubscriptionUnprocessableEntity with default headers values
func NewGetSubscriptionUnprocessableEntity() *GetSubscriptionUnprocessableEntity {
	return &GetSubscriptionUnprocessableEntity{}
}

/*
GetSubscriptionUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type GetSubscriptionUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this get subscription unprocessable entity response has a 2xx status code
func (o *GetSubscriptionUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get subscription unprocessable entity response has a 3xx status code
func (o *GetSubscriptionUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subscription unprocessable entity response has a 4xx status code
func (o *GetSubscriptionUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get subscription unprocessable entity response has a 5xx status code
func (o *GetSubscriptionUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get subscription unprocessable entity response a status code equal to that given
func (o *GetSubscriptionUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get subscription unprocessable entity response
func (o *GetSubscriptionUnprocessableEntity) Code() int {
	return 422
}

func (o *GetSubscriptionUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /subscription/{subscriptionID}][%d] getSubscriptionUnprocessableEntity %s", 422, payload)
}

func (o *GetSubscriptionUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /subscription/{subscriptionID}][%d] getSubscriptionUnprocessableEntity %s", 422, payload)
}

func (o *GetSubscriptionUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *GetSubscriptionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionInternalServerError creates a GetSubscriptionInternalServerError with default headers values
func NewGetSubscriptionInternalServerError() *GetSubscriptionInternalServerError {
	return &GetSubscriptionInternalServerError{}
}

/*
GetSubscriptionInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSubscriptionInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this get subscription internal server error response has a 2xx status code
func (o *GetSubscriptionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get subscription internal server error response has a 3xx status code
func (o *GetSubscriptionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subscription internal server error response has a 4xx status code
func (o *GetSubscriptionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get subscription internal server error response has a 5xx status code
func (o *GetSubscriptionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get subscription internal server error response a status code equal to that given
func (o *GetSubscriptionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get subscription internal server error response
func (o *GetSubscriptionInternalServerError) Code() int {
	return 500
}

func (o *GetSubscriptionInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /subscription/{subscriptionID}][%d] getSubscriptionInternalServerError %s", 500, payload)
}

func (o *GetSubscriptionInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /subscription/{subscriptionID}][%d] getSubscriptionInternalServerError %s", 500, payload)
}

func (o *GetSubscriptionInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *GetSubscriptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
