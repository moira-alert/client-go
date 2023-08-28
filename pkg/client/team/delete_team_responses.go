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

// DeleteTeamReader is a Reader for the DeleteTeam structure.
type DeleteTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTeamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTeamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /teams/{teamID}] delete-team", response, response.Code())
	}
}

// NewDeleteTeamOK creates a DeleteTeamOK with default headers values
func NewDeleteTeamOK() *DeleteTeamOK {
	return &DeleteTeamOK{}
}

/*
DeleteTeamOK describes a response with status code 200, with default header values.

Team has been successfully deleted
*/
type DeleteTeamOK struct {
	Payload *models.DtoSaveTeamResponse
}

// IsSuccess returns true when this delete team o k response has a 2xx status code
func (o *DeleteTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete team o k response has a 3xx status code
func (o *DeleteTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team o k response has a 4xx status code
func (o *DeleteTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team o k response has a 5xx status code
func (o *DeleteTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team o k response a status code equal to that given
func (o *DeleteTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete team o k response
func (o *DeleteTeamOK) Code() int {
	return 200
}

func (o *DeleteTeamOK) Error() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}][%d] deleteTeamOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamOK) String() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}][%d] deleteTeamOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamOK) GetPayload() *models.DtoSaveTeamResponse {
	return o.Payload
}

func (o *DeleteTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoSaveTeamResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamBadRequest creates a DeleteTeamBadRequest with default headers values
func NewDeleteTeamBadRequest() *DeleteTeamBadRequest {
	return &DeleteTeamBadRequest{}
}

/*
DeleteTeamBadRequest describes a response with status code 400, with default header values.

Bad request from client
*/
type DeleteTeamBadRequest struct {
	Payload *models.APIErrorInvalidRequestExample
}

// IsSuccess returns true when this delete team bad request response has a 2xx status code
func (o *DeleteTeamBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team bad request response has a 3xx status code
func (o *DeleteTeamBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team bad request response has a 4xx status code
func (o *DeleteTeamBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team bad request response has a 5xx status code
func (o *DeleteTeamBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team bad request response a status code equal to that given
func (o *DeleteTeamBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete team bad request response
func (o *DeleteTeamBadRequest) Code() int {
	return 400
}

func (o *DeleteTeamBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}][%d] deleteTeamBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTeamBadRequest) String() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}][%d] deleteTeamBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTeamBadRequest) GetPayload() *models.APIErrorInvalidRequestExample {
	return o.Payload
}

func (o *DeleteTeamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInvalidRequestExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamForbidden creates a DeleteTeamForbidden with default headers values
func NewDeleteTeamForbidden() *DeleteTeamForbidden {
	return &DeleteTeamForbidden{}
}

/*
DeleteTeamForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteTeamForbidden struct {
	Payload *models.APIErrorForbiddenExample
}

// IsSuccess returns true when this delete team forbidden response has a 2xx status code
func (o *DeleteTeamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team forbidden response has a 3xx status code
func (o *DeleteTeamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team forbidden response has a 4xx status code
func (o *DeleteTeamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team forbidden response has a 5xx status code
func (o *DeleteTeamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team forbidden response a status code equal to that given
func (o *DeleteTeamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete team forbidden response
func (o *DeleteTeamForbidden) Code() int {
	return 403
}

func (o *DeleteTeamForbidden) Error() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}][%d] deleteTeamForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamForbidden) String() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}][%d] deleteTeamForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamForbidden) GetPayload() *models.APIErrorForbiddenExample {
	return o.Payload
}

func (o *DeleteTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorForbiddenExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamNotFound creates a DeleteTeamNotFound with default headers values
func NewDeleteTeamNotFound() *DeleteTeamNotFound {
	return &DeleteTeamNotFound{}
}

/*
DeleteTeamNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type DeleteTeamNotFound struct {
	Payload *models.APIErrorNotFoundExample
}

// IsSuccess returns true when this delete team not found response has a 2xx status code
func (o *DeleteTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team not found response has a 3xx status code
func (o *DeleteTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team not found response has a 4xx status code
func (o *DeleteTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team not found response has a 5xx status code
func (o *DeleteTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team not found response a status code equal to that given
func (o *DeleteTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete team not found response
func (o *DeleteTeamNotFound) Code() int {
	return 404
}

func (o *DeleteTeamNotFound) Error() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}][%d] deleteTeamNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamNotFound) String() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}][%d] deleteTeamNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamNotFound) GetPayload() *models.APIErrorNotFoundExample {
	return o.Payload
}

func (o *DeleteTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorNotFoundExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamUnprocessableEntity creates a DeleteTeamUnprocessableEntity with default headers values
func NewDeleteTeamUnprocessableEntity() *DeleteTeamUnprocessableEntity {
	return &DeleteTeamUnprocessableEntity{}
}

/*
DeleteTeamUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type DeleteTeamUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this delete team unprocessable entity response has a 2xx status code
func (o *DeleteTeamUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team unprocessable entity response has a 3xx status code
func (o *DeleteTeamUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team unprocessable entity response has a 4xx status code
func (o *DeleteTeamUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team unprocessable entity response has a 5xx status code
func (o *DeleteTeamUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team unprocessable entity response a status code equal to that given
func (o *DeleteTeamUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the delete team unprocessable entity response
func (o *DeleteTeamUnprocessableEntity) Code() int {
	return 422
}

func (o *DeleteTeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}][%d] deleteTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteTeamUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}][%d] deleteTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteTeamUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *DeleteTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamInternalServerError creates a DeleteTeamInternalServerError with default headers values
func NewDeleteTeamInternalServerError() *DeleteTeamInternalServerError {
	return &DeleteTeamInternalServerError{}
}

/*
DeleteTeamInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteTeamInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this delete team internal server error response has a 2xx status code
func (o *DeleteTeamInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team internal server error response has a 3xx status code
func (o *DeleteTeamInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team internal server error response has a 4xx status code
func (o *DeleteTeamInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team internal server error response has a 5xx status code
func (o *DeleteTeamInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete team internal server error response a status code equal to that given
func (o *DeleteTeamInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete team internal server error response
func (o *DeleteTeamInternalServerError) Code() int {
	return 500
}

func (o *DeleteTeamInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}][%d] deleteTeamInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTeamInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /teams/{teamID}][%d] deleteTeamInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTeamInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *DeleteTeamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}