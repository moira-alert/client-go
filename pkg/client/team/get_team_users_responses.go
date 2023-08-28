// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moira-alert/client-go/pkg/models"
)

// GetTeamUsersReader is a Reader for the GetTeamUsers structure.
type GetTeamUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetTeamUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetTeamUsersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTeamUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /teams/{teamID}/users] get-team-users", response, response.Code())
	}
}

// NewGetTeamUsersOK creates a GetTeamUsersOK with default headers values
func NewGetTeamUsersOK() *GetTeamUsersOK {
	return &GetTeamUsersOK{}
}

/*
GetTeamUsersOK describes a response with status code 200, with default header values.

Users fetched successfully
*/
type GetTeamUsersOK struct {
	Payload *models.DtoTeamMembers
}

// IsSuccess returns true when this get team users o k response has a 2xx status code
func (o *GetTeamUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team users o k response has a 3xx status code
func (o *GetTeamUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team users o k response has a 4xx status code
func (o *GetTeamUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team users o k response has a 5xx status code
func (o *GetTeamUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team users o k response a status code equal to that given
func (o *GetTeamUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get team users o k response
func (o *GetTeamUsersOK) Code() int {
	return 200
}

func (o *GetTeamUsersOK) Error() string {
	return fmt.Sprintf("[GET /teams/{teamID}/users][%d] getTeamUsersOK  %+v", 200, o.Payload)
}

func (o *GetTeamUsersOK) String() string {
	return fmt.Sprintf("[GET /teams/{teamID}/users][%d] getTeamUsersOK  %+v", 200, o.Payload)
}

func (o *GetTeamUsersOK) GetPayload() *models.DtoTeamMembers {
	return o.Payload
}

func (o *GetTeamUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoTeamMembers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamUsersForbidden creates a GetTeamUsersForbidden with default headers values
func NewGetTeamUsersForbidden() *GetTeamUsersForbidden {
	return &GetTeamUsersForbidden{}
}

/*
GetTeamUsersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTeamUsersForbidden struct {
	Payload *models.APIErrorForbiddenExample
}

// IsSuccess returns true when this get team users forbidden response has a 2xx status code
func (o *GetTeamUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team users forbidden response has a 3xx status code
func (o *GetTeamUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team users forbidden response has a 4xx status code
func (o *GetTeamUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team users forbidden response has a 5xx status code
func (o *GetTeamUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get team users forbidden response a status code equal to that given
func (o *GetTeamUsersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get team users forbidden response
func (o *GetTeamUsersForbidden) Code() int {
	return 403
}

func (o *GetTeamUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /teams/{teamID}/users][%d] getTeamUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamUsersForbidden) String() string {
	return fmt.Sprintf("[GET /teams/{teamID}/users][%d] getTeamUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamUsersForbidden) GetPayload() *models.APIErrorForbiddenExample {
	return o.Payload
}

func (o *GetTeamUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorForbiddenExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamUsersNotFound creates a GetTeamUsersNotFound with default headers values
func NewGetTeamUsersNotFound() *GetTeamUsersNotFound {
	return &GetTeamUsersNotFound{}
}

/*
GetTeamUsersNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type GetTeamUsersNotFound struct {
	Payload *models.APIErrorNotFoundExample
}

// IsSuccess returns true when this get team users not found response has a 2xx status code
func (o *GetTeamUsersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team users not found response has a 3xx status code
func (o *GetTeamUsersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team users not found response has a 4xx status code
func (o *GetTeamUsersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team users not found response has a 5xx status code
func (o *GetTeamUsersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get team users not found response a status code equal to that given
func (o *GetTeamUsersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get team users not found response
func (o *GetTeamUsersNotFound) Code() int {
	return 404
}

func (o *GetTeamUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /teams/{teamID}/users][%d] getTeamUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamUsersNotFound) String() string {
	return fmt.Sprintf("[GET /teams/{teamID}/users][%d] getTeamUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamUsersNotFound) GetPayload() *models.APIErrorNotFoundExample {
	return o.Payload
}

func (o *GetTeamUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorNotFoundExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamUsersUnprocessableEntity creates a GetTeamUsersUnprocessableEntity with default headers values
func NewGetTeamUsersUnprocessableEntity() *GetTeamUsersUnprocessableEntity {
	return &GetTeamUsersUnprocessableEntity{}
}

/*
GetTeamUsersUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type GetTeamUsersUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this get team users unprocessable entity response has a 2xx status code
func (o *GetTeamUsersUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team users unprocessable entity response has a 3xx status code
func (o *GetTeamUsersUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team users unprocessable entity response has a 4xx status code
func (o *GetTeamUsersUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team users unprocessable entity response has a 5xx status code
func (o *GetTeamUsersUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get team users unprocessable entity response a status code equal to that given
func (o *GetTeamUsersUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get team users unprocessable entity response
func (o *GetTeamUsersUnprocessableEntity) Code() int {
	return 422
}

func (o *GetTeamUsersUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /teams/{teamID}/users][%d] getTeamUsersUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetTeamUsersUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /teams/{teamID}/users][%d] getTeamUsersUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetTeamUsersUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *GetTeamUsersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamUsersInternalServerError creates a GetTeamUsersInternalServerError with default headers values
func NewGetTeamUsersInternalServerError() *GetTeamUsersInternalServerError {
	return &GetTeamUsersInternalServerError{}
}

/*
GetTeamUsersInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetTeamUsersInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this get team users internal server error response has a 2xx status code
func (o *GetTeamUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team users internal server error response has a 3xx status code
func (o *GetTeamUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team users internal server error response has a 4xx status code
func (o *GetTeamUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team users internal server error response has a 5xx status code
func (o *GetTeamUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get team users internal server error response a status code equal to that given
func (o *GetTeamUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get team users internal server error response
func (o *GetTeamUsersInternalServerError) Code() int {
	return 500
}

func (o *GetTeamUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /teams/{teamID}/users][%d] getTeamUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTeamUsersInternalServerError) String() string {
	return fmt.Sprintf("[GET /teams/{teamID}/users][%d] getTeamUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTeamUsersInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *GetTeamUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
