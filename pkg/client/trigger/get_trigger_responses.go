// Code generated by go-swagger; DO NOT EDIT.

package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moira-alert/client-go/pkg/models"
)

// GetTriggerReader is a Reader for the GetTrigger structure.
type GetTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTriggerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetTriggerUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTriggerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /trigger/{triggerID}] get-trigger", response, response.Code())
	}
}

// NewGetTriggerOK creates a GetTriggerOK with default headers values
func NewGetTriggerOK() *GetTriggerOK {
	return &GetTriggerOK{}
}

/*
GetTriggerOK describes a response with status code 200, with default header values.

Trigger data
*/
type GetTriggerOK struct {
	Payload *models.DtoTrigger
}

// IsSuccess returns true when this get trigger o k response has a 2xx status code
func (o *GetTriggerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get trigger o k response has a 3xx status code
func (o *GetTriggerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get trigger o k response has a 4xx status code
func (o *GetTriggerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get trigger o k response has a 5xx status code
func (o *GetTriggerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get trigger o k response a status code equal to that given
func (o *GetTriggerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get trigger o k response
func (o *GetTriggerOK) Code() int {
	return 200
}

func (o *GetTriggerOK) Error() string {
	return fmt.Sprintf("[GET /trigger/{triggerID}][%d] getTriggerOK  %+v", 200, o.Payload)
}

func (o *GetTriggerOK) String() string {
	return fmt.Sprintf("[GET /trigger/{triggerID}][%d] getTriggerOK  %+v", 200, o.Payload)
}

func (o *GetTriggerOK) GetPayload() *models.DtoTrigger {
	return o.Payload
}

func (o *GetTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoTrigger)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTriggerNotFound creates a GetTriggerNotFound with default headers values
func NewGetTriggerNotFound() *GetTriggerNotFound {
	return &GetTriggerNotFound{}
}

/*
GetTriggerNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type GetTriggerNotFound struct {
	Payload *models.APIErrorNotFoundExample
}

// IsSuccess returns true when this get trigger not found response has a 2xx status code
func (o *GetTriggerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get trigger not found response has a 3xx status code
func (o *GetTriggerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get trigger not found response has a 4xx status code
func (o *GetTriggerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get trigger not found response has a 5xx status code
func (o *GetTriggerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get trigger not found response a status code equal to that given
func (o *GetTriggerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get trigger not found response
func (o *GetTriggerNotFound) Code() int {
	return 404
}

func (o *GetTriggerNotFound) Error() string {
	return fmt.Sprintf("[GET /trigger/{triggerID}][%d] getTriggerNotFound  %+v", 404, o.Payload)
}

func (o *GetTriggerNotFound) String() string {
	return fmt.Sprintf("[GET /trigger/{triggerID}][%d] getTriggerNotFound  %+v", 404, o.Payload)
}

func (o *GetTriggerNotFound) GetPayload() *models.APIErrorNotFoundExample {
	return o.Payload
}

func (o *GetTriggerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorNotFoundExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTriggerUnprocessableEntity creates a GetTriggerUnprocessableEntity with default headers values
func NewGetTriggerUnprocessableEntity() *GetTriggerUnprocessableEntity {
	return &GetTriggerUnprocessableEntity{}
}

/*
GetTriggerUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type GetTriggerUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this get trigger unprocessable entity response has a 2xx status code
func (o *GetTriggerUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get trigger unprocessable entity response has a 3xx status code
func (o *GetTriggerUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get trigger unprocessable entity response has a 4xx status code
func (o *GetTriggerUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get trigger unprocessable entity response has a 5xx status code
func (o *GetTriggerUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get trigger unprocessable entity response a status code equal to that given
func (o *GetTriggerUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get trigger unprocessable entity response
func (o *GetTriggerUnprocessableEntity) Code() int {
	return 422
}

func (o *GetTriggerUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /trigger/{triggerID}][%d] getTriggerUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetTriggerUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /trigger/{triggerID}][%d] getTriggerUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetTriggerUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *GetTriggerUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTriggerInternalServerError creates a GetTriggerInternalServerError with default headers values
func NewGetTriggerInternalServerError() *GetTriggerInternalServerError {
	return &GetTriggerInternalServerError{}
}

/*
GetTriggerInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetTriggerInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this get trigger internal server error response has a 2xx status code
func (o *GetTriggerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get trigger internal server error response has a 3xx status code
func (o *GetTriggerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get trigger internal server error response has a 4xx status code
func (o *GetTriggerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get trigger internal server error response has a 5xx status code
func (o *GetTriggerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get trigger internal server error response a status code equal to that given
func (o *GetTriggerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get trigger internal server error response
func (o *GetTriggerInternalServerError) Code() int {
	return 500
}

func (o *GetTriggerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /trigger/{triggerID}][%d] getTriggerInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTriggerInternalServerError) String() string {
	return fmt.Sprintf("[GET /trigger/{triggerID}][%d] getTriggerInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTriggerInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *GetTriggerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
