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

// GetAllTeamsReader is a Reader for the GetAllTeams structure.
type GetAllTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAllTeamsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllTeamsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAllTeamsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllTeamsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /teams/all] get-all-teams", response, response.Code())
	}
}

// NewGetAllTeamsOK creates a GetAllTeamsOK with default headers values
func NewGetAllTeamsOK() *GetAllTeamsOK {
	return &GetAllTeamsOK{}
}

/*
GetAllTeamsOK describes a response with status code 200, with default header values.

Teams fetched successfully
*/
type GetAllTeamsOK struct {
	Payload *models.DtoTeamsList
}

// IsSuccess returns true when this get all teams o k response has a 2xx status code
func (o *GetAllTeamsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all teams o k response has a 3xx status code
func (o *GetAllTeamsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all teams o k response has a 4xx status code
func (o *GetAllTeamsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all teams o k response has a 5xx status code
func (o *GetAllTeamsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all teams o k response a status code equal to that given
func (o *GetAllTeamsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all teams o k response
func (o *GetAllTeamsOK) Code() int {
	return 200
}

func (o *GetAllTeamsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/all][%d] getAllTeamsOK %s", 200, payload)
}

func (o *GetAllTeamsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/all][%d] getAllTeamsOK %s", 200, payload)
}

func (o *GetAllTeamsOK) GetPayload() *models.DtoTeamsList {
	return o.Payload
}

func (o *GetAllTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoTeamsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTeamsBadRequest creates a GetAllTeamsBadRequest with default headers values
func NewGetAllTeamsBadRequest() *GetAllTeamsBadRequest {
	return &GetAllTeamsBadRequest{}
}

/*
GetAllTeamsBadRequest describes a response with status code 400, with default header values.

Bad request from client
*/
type GetAllTeamsBadRequest struct {
	Payload *models.APIErrorInvalidRequestExample
}

// IsSuccess returns true when this get all teams bad request response has a 2xx status code
func (o *GetAllTeamsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all teams bad request response has a 3xx status code
func (o *GetAllTeamsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all teams bad request response has a 4xx status code
func (o *GetAllTeamsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all teams bad request response has a 5xx status code
func (o *GetAllTeamsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get all teams bad request response a status code equal to that given
func (o *GetAllTeamsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get all teams bad request response
func (o *GetAllTeamsBadRequest) Code() int {
	return 400
}

func (o *GetAllTeamsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/all][%d] getAllTeamsBadRequest %s", 400, payload)
}

func (o *GetAllTeamsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/all][%d] getAllTeamsBadRequest %s", 400, payload)
}

func (o *GetAllTeamsBadRequest) GetPayload() *models.APIErrorInvalidRequestExample {
	return o.Payload
}

func (o *GetAllTeamsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInvalidRequestExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTeamsForbidden creates a GetAllTeamsForbidden with default headers values
func NewGetAllTeamsForbidden() *GetAllTeamsForbidden {
	return &GetAllTeamsForbidden{}
}

/*
GetAllTeamsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllTeamsForbidden struct {
	Payload *models.APIErrorForbiddenExample
}

// IsSuccess returns true when this get all teams forbidden response has a 2xx status code
func (o *GetAllTeamsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all teams forbidden response has a 3xx status code
func (o *GetAllTeamsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all teams forbidden response has a 4xx status code
func (o *GetAllTeamsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all teams forbidden response has a 5xx status code
func (o *GetAllTeamsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all teams forbidden response a status code equal to that given
func (o *GetAllTeamsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get all teams forbidden response
func (o *GetAllTeamsForbidden) Code() int {
	return 403
}

func (o *GetAllTeamsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/all][%d] getAllTeamsForbidden %s", 403, payload)
}

func (o *GetAllTeamsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/all][%d] getAllTeamsForbidden %s", 403, payload)
}

func (o *GetAllTeamsForbidden) GetPayload() *models.APIErrorForbiddenExample {
	return o.Payload
}

func (o *GetAllTeamsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorForbiddenExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTeamsUnprocessableEntity creates a GetAllTeamsUnprocessableEntity with default headers values
func NewGetAllTeamsUnprocessableEntity() *GetAllTeamsUnprocessableEntity {
	return &GetAllTeamsUnprocessableEntity{}
}

/*
GetAllTeamsUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type GetAllTeamsUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this get all teams unprocessable entity response has a 2xx status code
func (o *GetAllTeamsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all teams unprocessable entity response has a 3xx status code
func (o *GetAllTeamsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all teams unprocessable entity response has a 4xx status code
func (o *GetAllTeamsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all teams unprocessable entity response has a 5xx status code
func (o *GetAllTeamsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get all teams unprocessable entity response a status code equal to that given
func (o *GetAllTeamsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get all teams unprocessable entity response
func (o *GetAllTeamsUnprocessableEntity) Code() int {
	return 422
}

func (o *GetAllTeamsUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/all][%d] getAllTeamsUnprocessableEntity %s", 422, payload)
}

func (o *GetAllTeamsUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/all][%d] getAllTeamsUnprocessableEntity %s", 422, payload)
}

func (o *GetAllTeamsUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *GetAllTeamsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTeamsInternalServerError creates a GetAllTeamsInternalServerError with default headers values
func NewGetAllTeamsInternalServerError() *GetAllTeamsInternalServerError {
	return &GetAllTeamsInternalServerError{}
}

/*
GetAllTeamsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAllTeamsInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this get all teams internal server error response has a 2xx status code
func (o *GetAllTeamsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all teams internal server error response has a 3xx status code
func (o *GetAllTeamsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all teams internal server error response has a 4xx status code
func (o *GetAllTeamsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all teams internal server error response has a 5xx status code
func (o *GetAllTeamsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all teams internal server error response a status code equal to that given
func (o *GetAllTeamsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get all teams internal server error response
func (o *GetAllTeamsInternalServerError) Code() int {
	return 500
}

func (o *GetAllTeamsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/all][%d] getAllTeamsInternalServerError %s", 500, payload)
}

func (o *GetAllTeamsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/all][%d] getAllTeamsInternalServerError %s", 500, payload)
}

func (o *GetAllTeamsInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *GetAllTeamsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
