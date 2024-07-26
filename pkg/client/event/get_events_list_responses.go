// Code generated by go-swagger; DO NOT EDIT.

package event

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

// GetEventsListReader is a Reader for the GetEventsList structure.
type GetEventsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEventsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEventsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetEventsListUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEventsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /event/{triggerID}] get-events-list", response, response.Code())
	}
}

// NewGetEventsListOK creates a GetEventsListOK with default headers values
func NewGetEventsListOK() *GetEventsListOK {
	return &GetEventsListOK{}
}

/*
GetEventsListOK describes a response with status code 200, with default header values.

Events fetched successfully
*/
type GetEventsListOK struct {
	Payload *models.DtoEventsList
}

// IsSuccess returns true when this get events list o k response has a 2xx status code
func (o *GetEventsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get events list o k response has a 3xx status code
func (o *GetEventsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events list o k response has a 4xx status code
func (o *GetEventsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get events list o k response has a 5xx status code
func (o *GetEventsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get events list o k response a status code equal to that given
func (o *GetEventsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get events list o k response
func (o *GetEventsListOK) Code() int {
	return 200
}

func (o *GetEventsListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /event/{triggerID}][%d] getEventsListOK %s", 200, payload)
}

func (o *GetEventsListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /event/{triggerID}][%d] getEventsListOK %s", 200, payload)
}

func (o *GetEventsListOK) GetPayload() *models.DtoEventsList {
	return o.Payload
}

func (o *GetEventsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoEventsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsListBadRequest creates a GetEventsListBadRequest with default headers values
func NewGetEventsListBadRequest() *GetEventsListBadRequest {
	return &GetEventsListBadRequest{}
}

/*
GetEventsListBadRequest describes a response with status code 400, with default header values.

Bad request from client
*/
type GetEventsListBadRequest struct {
	Payload *models.APIErrorInvalidRequestExample
}

// IsSuccess returns true when this get events list bad request response has a 2xx status code
func (o *GetEventsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get events list bad request response has a 3xx status code
func (o *GetEventsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events list bad request response has a 4xx status code
func (o *GetEventsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get events list bad request response has a 5xx status code
func (o *GetEventsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get events list bad request response a status code equal to that given
func (o *GetEventsListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get events list bad request response
func (o *GetEventsListBadRequest) Code() int {
	return 400
}

func (o *GetEventsListBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /event/{triggerID}][%d] getEventsListBadRequest %s", 400, payload)
}

func (o *GetEventsListBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /event/{triggerID}][%d] getEventsListBadRequest %s", 400, payload)
}

func (o *GetEventsListBadRequest) GetPayload() *models.APIErrorInvalidRequestExample {
	return o.Payload
}

func (o *GetEventsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInvalidRequestExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsListNotFound creates a GetEventsListNotFound with default headers values
func NewGetEventsListNotFound() *GetEventsListNotFound {
	return &GetEventsListNotFound{}
}

/*
GetEventsListNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type GetEventsListNotFound struct {
	Payload *models.APIErrorNotFoundExample
}

// IsSuccess returns true when this get events list not found response has a 2xx status code
func (o *GetEventsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get events list not found response has a 3xx status code
func (o *GetEventsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events list not found response has a 4xx status code
func (o *GetEventsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get events list not found response has a 5xx status code
func (o *GetEventsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get events list not found response a status code equal to that given
func (o *GetEventsListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get events list not found response
func (o *GetEventsListNotFound) Code() int {
	return 404
}

func (o *GetEventsListNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /event/{triggerID}][%d] getEventsListNotFound %s", 404, payload)
}

func (o *GetEventsListNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /event/{triggerID}][%d] getEventsListNotFound %s", 404, payload)
}

func (o *GetEventsListNotFound) GetPayload() *models.APIErrorNotFoundExample {
	return o.Payload
}

func (o *GetEventsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorNotFoundExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsListUnprocessableEntity creates a GetEventsListUnprocessableEntity with default headers values
func NewGetEventsListUnprocessableEntity() *GetEventsListUnprocessableEntity {
	return &GetEventsListUnprocessableEntity{}
}

/*
GetEventsListUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type GetEventsListUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this get events list unprocessable entity response has a 2xx status code
func (o *GetEventsListUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get events list unprocessable entity response has a 3xx status code
func (o *GetEventsListUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events list unprocessable entity response has a 4xx status code
func (o *GetEventsListUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get events list unprocessable entity response has a 5xx status code
func (o *GetEventsListUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get events list unprocessable entity response a status code equal to that given
func (o *GetEventsListUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get events list unprocessable entity response
func (o *GetEventsListUnprocessableEntity) Code() int {
	return 422
}

func (o *GetEventsListUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /event/{triggerID}][%d] getEventsListUnprocessableEntity %s", 422, payload)
}

func (o *GetEventsListUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /event/{triggerID}][%d] getEventsListUnprocessableEntity %s", 422, payload)
}

func (o *GetEventsListUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *GetEventsListUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsListInternalServerError creates a GetEventsListInternalServerError with default headers values
func NewGetEventsListInternalServerError() *GetEventsListInternalServerError {
	return &GetEventsListInternalServerError{}
}

/*
GetEventsListInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetEventsListInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this get events list internal server error response has a 2xx status code
func (o *GetEventsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get events list internal server error response has a 3xx status code
func (o *GetEventsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get events list internal server error response has a 4xx status code
func (o *GetEventsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get events list internal server error response has a 5xx status code
func (o *GetEventsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get events list internal server error response a status code equal to that given
func (o *GetEventsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get events list internal server error response
func (o *GetEventsListInternalServerError) Code() int {
	return 500
}

func (o *GetEventsListInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /event/{triggerID}][%d] getEventsListInternalServerError %s", 500, payload)
}

func (o *GetEventsListInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /event/{triggerID}][%d] getEventsListInternalServerError %s", 500, payload)
}

func (o *GetEventsListInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *GetEventsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
