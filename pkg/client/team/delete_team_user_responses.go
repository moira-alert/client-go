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

// DeleteTeamUserReader is a Reader for the DeleteTeamUser structure.
type DeleteTeamUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTeamUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTeamUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTeamUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTeamUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTeamUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteTeamUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTeamUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /teams/{teamID}/users/{teamUserID}] delete-team-user", response, response.Code())
	}
}

// NewDeleteTeamUserOK creates a DeleteTeamUserOK with default headers values
func NewDeleteTeamUserOK() *DeleteTeamUserOK {
	return &DeleteTeamUserOK{}
}

/*
DeleteTeamUserOK describes a response with status code 200, with default header values.

Team updated successfully
*/
type DeleteTeamUserOK struct {
	Payload *models.DtoTeamMembers
}

// IsSuccess returns true when this delete team user o k response has a 2xx status code
func (o *DeleteTeamUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete team user o k response has a 3xx status code
func (o *DeleteTeamUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team user o k response has a 4xx status code
func (o *DeleteTeamUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team user o k response has a 5xx status code
func (o *DeleteTeamUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team user o k response a status code equal to that given
func (o *DeleteTeamUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete team user o k response
func (o *DeleteTeamUserOK) Code() int {
	return 200
}

func (o *DeleteTeamUserOK) Error() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}/users/{teamUserID}][%d] deleteTeamUserOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamUserOK) String() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}/users/{teamUserID}][%d] deleteTeamUserOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamUserOK) GetPayload() *models.DtoTeamMembers {
	return o.Payload
}

func (o *DeleteTeamUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoTeamMembers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamUserBadRequest creates a DeleteTeamUserBadRequest with default headers values
func NewDeleteTeamUserBadRequest() *DeleteTeamUserBadRequest {
	return &DeleteTeamUserBadRequest{}
}

/*
DeleteTeamUserBadRequest describes a response with status code 400, with default header values.

Bad request from client
*/
type DeleteTeamUserBadRequest struct {
	Payload *models.APIErrorInvalidRequestExample
}

// IsSuccess returns true when this delete team user bad request response has a 2xx status code
func (o *DeleteTeamUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team user bad request response has a 3xx status code
func (o *DeleteTeamUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team user bad request response has a 4xx status code
func (o *DeleteTeamUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team user bad request response has a 5xx status code
func (o *DeleteTeamUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team user bad request response a status code equal to that given
func (o *DeleteTeamUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete team user bad request response
func (o *DeleteTeamUserBadRequest) Code() int {
	return 400
}

func (o *DeleteTeamUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}/users/{teamUserID}][%d] deleteTeamUserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTeamUserBadRequest) String() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}/users/{teamUserID}][%d] deleteTeamUserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTeamUserBadRequest) GetPayload() *models.APIErrorInvalidRequestExample {
	return o.Payload
}

func (o *DeleteTeamUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInvalidRequestExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamUserForbidden creates a DeleteTeamUserForbidden with default headers values
func NewDeleteTeamUserForbidden() *DeleteTeamUserForbidden {
	return &DeleteTeamUserForbidden{}
}

/*
DeleteTeamUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteTeamUserForbidden struct {
	Payload *models.APIErrorForbiddenExample
}

// IsSuccess returns true when this delete team user forbidden response has a 2xx status code
func (o *DeleteTeamUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team user forbidden response has a 3xx status code
func (o *DeleteTeamUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team user forbidden response has a 4xx status code
func (o *DeleteTeamUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team user forbidden response has a 5xx status code
func (o *DeleteTeamUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team user forbidden response a status code equal to that given
func (o *DeleteTeamUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete team user forbidden response
func (o *DeleteTeamUserForbidden) Code() int {
	return 403
}

func (o *DeleteTeamUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}/users/{teamUserID}][%d] deleteTeamUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamUserForbidden) String() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}/users/{teamUserID}][%d] deleteTeamUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamUserForbidden) GetPayload() *models.APIErrorForbiddenExample {
	return o.Payload
}

func (o *DeleteTeamUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorForbiddenExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamUserNotFound creates a DeleteTeamUserNotFound with default headers values
func NewDeleteTeamUserNotFound() *DeleteTeamUserNotFound {
	return &DeleteTeamUserNotFound{}
}

/*
DeleteTeamUserNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type DeleteTeamUserNotFound struct {
	Payload *models.APIErrorNotFoundExample
}

// IsSuccess returns true when this delete team user not found response has a 2xx status code
func (o *DeleteTeamUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team user not found response has a 3xx status code
func (o *DeleteTeamUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team user not found response has a 4xx status code
func (o *DeleteTeamUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team user not found response has a 5xx status code
func (o *DeleteTeamUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team user not found response a status code equal to that given
func (o *DeleteTeamUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete team user not found response
func (o *DeleteTeamUserNotFound) Code() int {
	return 404
}

func (o *DeleteTeamUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}/users/{teamUserID}][%d] deleteTeamUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamUserNotFound) String() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}/users/{teamUserID}][%d] deleteTeamUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamUserNotFound) GetPayload() *models.APIErrorNotFoundExample {
	return o.Payload
}

func (o *DeleteTeamUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorNotFoundExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamUserUnprocessableEntity creates a DeleteTeamUserUnprocessableEntity with default headers values
func NewDeleteTeamUserUnprocessableEntity() *DeleteTeamUserUnprocessableEntity {
	return &DeleteTeamUserUnprocessableEntity{}
}

/*
DeleteTeamUserUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type DeleteTeamUserUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this delete team user unprocessable entity response has a 2xx status code
func (o *DeleteTeamUserUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team user unprocessable entity response has a 3xx status code
func (o *DeleteTeamUserUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team user unprocessable entity response has a 4xx status code
func (o *DeleteTeamUserUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team user unprocessable entity response has a 5xx status code
func (o *DeleteTeamUserUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team user unprocessable entity response a status code equal to that given
func (o *DeleteTeamUserUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete team user unprocessable entity response
func (o *DeleteTeamUserUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteTeamUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}/users/{teamUserID}][%d] deleteTeamUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteTeamUserUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}/users/{teamUserID}][%d] deleteTeamUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteTeamUserUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *DeleteTeamUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamUserInternalServerError creates a DeleteTeamUserInternalServerError with default headers values
func NewDeleteTeamUserInternalServerError() *DeleteTeamUserInternalServerError {
	return &DeleteTeamUserInternalServerError{}
}

/*
DeleteTeamUserInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteTeamUserInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this delete team user internal server error response has a 2xx status code
func (o *DeleteTeamUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team user internal server error response has a 3xx status code
func (o *DeleteTeamUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team user internal server error response has a 4xx status code
func (o *DeleteTeamUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team user internal server error response has a 5xx status code
func (o *DeleteTeamUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete team user internal server error response a status code equal to that given
func (o *DeleteTeamUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete team user internal server error response
func (o *DeleteTeamUserInternalServerError) Code() int {
	return 500
}

func (o *DeleteTeamUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}/users/{teamUserID}][%d] deleteTeamUserInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTeamUserInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}/users/{teamUserID}][%d] deleteTeamUserInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTeamUserInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *DeleteTeamUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
