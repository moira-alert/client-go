// Code generated by go-swagger; DO NOT EDIT.

package team

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

// GetTeamReader is a Reader for the GetTeam structure.
type GetTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTeamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /teams/{teamID}] get-team", response, response.Code())
	}
}

// NewGetTeamOK creates a GetTeamOK with default headers values
func NewGetTeamOK() *GetTeamOK {
	return &GetTeamOK{}
}

/*
GetTeamOK describes a response with status code 200, with default header values.

Team updated successfully
*/
type GetTeamOK struct {
	Payload *models.DtoTeamModel
}

// IsSuccess returns true when this get team o k response has a 2xx status code
func (o *GetTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team o k response has a 3xx status code
func (o *GetTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team o k response has a 4xx status code
func (o *GetTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team o k response has a 5xx status code
func (o *GetTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team o k response a status code equal to that given
func (o *GetTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get team o k response
func (o *GetTeamOK) Code() int {
	return 200
}

func (o *GetTeamOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{teamID}][%d] getTeamOK %s", 200, payload)
}

func (o *GetTeamOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{teamID}][%d] getTeamOK %s", 200, payload)
}

func (o *GetTeamOK) GetPayload() *models.DtoTeamModel {
	return o.Payload
}

func (o *GetTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoTeamModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamForbidden creates a GetTeamForbidden with default headers values
func NewGetTeamForbidden() *GetTeamForbidden {
	return &GetTeamForbidden{}
}

/*
GetTeamForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTeamForbidden struct {
	Payload *models.APIErrorForbiddenExample
}

// IsSuccess returns true when this get team forbidden response has a 2xx status code
func (o *GetTeamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team forbidden response has a 3xx status code
func (o *GetTeamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team forbidden response has a 4xx status code
func (o *GetTeamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team forbidden response has a 5xx status code
func (o *GetTeamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get team forbidden response a status code equal to that given
func (o *GetTeamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get team forbidden response
func (o *GetTeamForbidden) Code() int {
	return 403
}

func (o *GetTeamForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{teamID}][%d] getTeamForbidden %s", 403, payload)
}

func (o *GetTeamForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{teamID}][%d] getTeamForbidden %s", 403, payload)
}

func (o *GetTeamForbidden) GetPayload() *models.APIErrorForbiddenExample {
	return o.Payload
}

func (o *GetTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorForbiddenExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamNotFound creates a GetTeamNotFound with default headers values
func NewGetTeamNotFound() *GetTeamNotFound {
	return &GetTeamNotFound{}
}

/*
GetTeamNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type GetTeamNotFound struct {
	Payload *models.APIErrorNotFoundExample
}

// IsSuccess returns true when this get team not found response has a 2xx status code
func (o *GetTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team not found response has a 3xx status code
func (o *GetTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team not found response has a 4xx status code
func (o *GetTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team not found response has a 5xx status code
func (o *GetTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get team not found response a status code equal to that given
func (o *GetTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get team not found response
func (o *GetTeamNotFound) Code() int {
	return 404
}

func (o *GetTeamNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{teamID}][%d] getTeamNotFound %s", 404, payload)
}

func (o *GetTeamNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{teamID}][%d] getTeamNotFound %s", 404, payload)
}

func (o *GetTeamNotFound) GetPayload() *models.APIErrorNotFoundExample {
	return o.Payload
}

func (o *GetTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorNotFoundExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamUnprocessableEntity creates a GetTeamUnprocessableEntity with default headers values
func NewGetTeamUnprocessableEntity() *GetTeamUnprocessableEntity {
	return &GetTeamUnprocessableEntity{}
}

/*
GetTeamUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type GetTeamUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this get team unprocessable entity response has a 2xx status code
func (o *GetTeamUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team unprocessable entity response has a 3xx status code
func (o *GetTeamUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team unprocessable entity response has a 4xx status code
func (o *GetTeamUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team unprocessable entity response has a 5xx status code
func (o *GetTeamUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get team unprocessable entity response a status code equal to that given
func (o *GetTeamUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get team unprocessable entity response
func (o *GetTeamUnprocessableEntity) Code() int {
	return 422
}

func (o *GetTeamUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{teamID}][%d] getTeamUnprocessableEntity %s", 422, payload)
}

func (o *GetTeamUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{teamID}][%d] getTeamUnprocessableEntity %s", 422, payload)
}

func (o *GetTeamUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *GetTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamInternalServerError creates a GetTeamInternalServerError with default headers values
func NewGetTeamInternalServerError() *GetTeamInternalServerError {
	return &GetTeamInternalServerError{}
}

/*
GetTeamInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetTeamInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this get team internal server error response has a 2xx status code
func (o *GetTeamInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team internal server error response has a 3xx status code
func (o *GetTeamInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team internal server error response has a 4xx status code
func (o *GetTeamInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team internal server error response has a 5xx status code
func (o *GetTeamInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get team internal server error response a status code equal to that given
func (o *GetTeamInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get team internal server error response
func (o *GetTeamInternalServerError) Code() int {
	return 500
}

func (o *GetTeamInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{teamID}][%d] getTeamInternalServerError %s", 500, payload)
}

func (o *GetTeamInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{teamID}][%d] getTeamInternalServerError %s", 500, payload)
}

func (o *GetTeamInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *GetTeamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
