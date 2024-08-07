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

// GetTriggerThrottlingReader is a Reader for the GetTriggerThrottling structure.
type GetTriggerThrottlingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTriggerThrottlingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTriggerThrottlingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTriggerThrottlingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetTriggerThrottlingUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /trigger/{triggerID}/throttling] get-trigger-throttling", response, response.Code())
	}
}

// NewGetTriggerThrottlingOK creates a GetTriggerThrottlingOK with default headers values
func NewGetTriggerThrottlingOK() *GetTriggerThrottlingOK {
	return &GetTriggerThrottlingOK{}
}

/*
GetTriggerThrottlingOK describes a response with status code 200, with default header values.

Trigger throttle info retrieved
*/
type GetTriggerThrottlingOK struct {
	Payload *models.DtoThrottlingResponse
}

// IsSuccess returns true when this get trigger throttling o k response has a 2xx status code
func (o *GetTriggerThrottlingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get trigger throttling o k response has a 3xx status code
func (o *GetTriggerThrottlingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get trigger throttling o k response has a 4xx status code
func (o *GetTriggerThrottlingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get trigger throttling o k response has a 5xx status code
func (o *GetTriggerThrottlingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get trigger throttling o k response a status code equal to that given
func (o *GetTriggerThrottlingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get trigger throttling o k response
func (o *GetTriggerThrottlingOK) Code() int {
	return 200
}

func (o *GetTriggerThrottlingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /trigger/{triggerID}/throttling][%d] getTriggerThrottlingOK %s", 200, payload)
}

func (o *GetTriggerThrottlingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /trigger/{triggerID}/throttling][%d] getTriggerThrottlingOK %s", 200, payload)
}

func (o *GetTriggerThrottlingOK) GetPayload() *models.DtoThrottlingResponse {
	return o.Payload
}

func (o *GetTriggerThrottlingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoThrottlingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTriggerThrottlingNotFound creates a GetTriggerThrottlingNotFound with default headers values
func NewGetTriggerThrottlingNotFound() *GetTriggerThrottlingNotFound {
	return &GetTriggerThrottlingNotFound{}
}

/*
GetTriggerThrottlingNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type GetTriggerThrottlingNotFound struct {
	Payload *models.APIErrorNotFoundExample
}

// IsSuccess returns true when this get trigger throttling not found response has a 2xx status code
func (o *GetTriggerThrottlingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get trigger throttling not found response has a 3xx status code
func (o *GetTriggerThrottlingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get trigger throttling not found response has a 4xx status code
func (o *GetTriggerThrottlingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get trigger throttling not found response has a 5xx status code
func (o *GetTriggerThrottlingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get trigger throttling not found response a status code equal to that given
func (o *GetTriggerThrottlingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get trigger throttling not found response
func (o *GetTriggerThrottlingNotFound) Code() int {
	return 404
}

func (o *GetTriggerThrottlingNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /trigger/{triggerID}/throttling][%d] getTriggerThrottlingNotFound %s", 404, payload)
}

func (o *GetTriggerThrottlingNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /trigger/{triggerID}/throttling][%d] getTriggerThrottlingNotFound %s", 404, payload)
}

func (o *GetTriggerThrottlingNotFound) GetPayload() *models.APIErrorNotFoundExample {
	return o.Payload
}

func (o *GetTriggerThrottlingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorNotFoundExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTriggerThrottlingUnprocessableEntity creates a GetTriggerThrottlingUnprocessableEntity with default headers values
func NewGetTriggerThrottlingUnprocessableEntity() *GetTriggerThrottlingUnprocessableEntity {
	return &GetTriggerThrottlingUnprocessableEntity{}
}

/*
GetTriggerThrottlingUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type GetTriggerThrottlingUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this get trigger throttling unprocessable entity response has a 2xx status code
func (o *GetTriggerThrottlingUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get trigger throttling unprocessable entity response has a 3xx status code
func (o *GetTriggerThrottlingUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get trigger throttling unprocessable entity response has a 4xx status code
func (o *GetTriggerThrottlingUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get trigger throttling unprocessable entity response has a 5xx status code
func (o *GetTriggerThrottlingUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get trigger throttling unprocessable entity response a status code equal to that given
func (o *GetTriggerThrottlingUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get trigger throttling unprocessable entity response
func (o *GetTriggerThrottlingUnprocessableEntity) Code() int {
	return 422
}

func (o *GetTriggerThrottlingUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /trigger/{triggerID}/throttling][%d] getTriggerThrottlingUnprocessableEntity %s", 422, payload)
}

func (o *GetTriggerThrottlingUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /trigger/{triggerID}/throttling][%d] getTriggerThrottlingUnprocessableEntity %s", 422, payload)
}

func (o *GetTriggerThrottlingUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *GetTriggerThrottlingUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
