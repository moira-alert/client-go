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

// GetTeamSettingsReader is a Reader for the GetTeamSettings structure.
type GetTeamSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetTeamSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetTeamSettingsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTeamSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /teams/{teamID}/settings] get-team-settings", response, response.Code())
	}
}

// NewGetTeamSettingsOK creates a GetTeamSettingsOK with default headers values
func NewGetTeamSettingsOK() *GetTeamSettingsOK {
	return &GetTeamSettingsOK{}
}

/*
GetTeamSettingsOK describes a response with status code 200, with default header values.

Team settings
*/
type GetTeamSettingsOK struct {
	Payload *models.DtoTeamSettings
}

// IsSuccess returns true when this get team settings o k response has a 2xx status code
func (o *GetTeamSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team settings o k response has a 3xx status code
func (o *GetTeamSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team settings o k response has a 4xx status code
func (o *GetTeamSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team settings o k response has a 5xx status code
func (o *GetTeamSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team settings o k response a status code equal to that given
func (o *GetTeamSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get team settings o k response
func (o *GetTeamSettingsOK) Code() int {
	return 200
}

func (o *GetTeamSettingsOK) Error() string {
	return fmt.Sprintf("[GET /teams/{teamID}/settings][%d] getTeamSettingsOK  %+v", 200, o.Payload)
}

func (o *GetTeamSettingsOK) String() string {
	return fmt.Sprintf("[GET /teams/{teamID}/settings][%d] getTeamSettingsOK  %+v", 200, o.Payload)
}

func (o *GetTeamSettingsOK) GetPayload() *models.DtoTeamSettings {
	return o.Payload
}

func (o *GetTeamSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoTeamSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamSettingsForbidden creates a GetTeamSettingsForbidden with default headers values
func NewGetTeamSettingsForbidden() *GetTeamSettingsForbidden {
	return &GetTeamSettingsForbidden{}
}

/*
GetTeamSettingsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTeamSettingsForbidden struct {
	Payload *models.APIErrorForbiddenExample
}

// IsSuccess returns true when this get team settings forbidden response has a 2xx status code
func (o *GetTeamSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team settings forbidden response has a 3xx status code
func (o *GetTeamSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team settings forbidden response has a 4xx status code
func (o *GetTeamSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team settings forbidden response has a 5xx status code
func (o *GetTeamSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get team settings forbidden response a status code equal to that given
func (o *GetTeamSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get team settings forbidden response
func (o *GetTeamSettingsForbidden) Code() int {
	return 403
}

func (o *GetTeamSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /teams/{teamID}/settings][%d] getTeamSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamSettingsForbidden) String() string {
	return fmt.Sprintf("[GET /teams/{teamID}/settings][%d] getTeamSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamSettingsForbidden) GetPayload() *models.APIErrorForbiddenExample {
	return o.Payload
}

func (o *GetTeamSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorForbiddenExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamSettingsNotFound creates a GetTeamSettingsNotFound with default headers values
func NewGetTeamSettingsNotFound() *GetTeamSettingsNotFound {
	return &GetTeamSettingsNotFound{}
}

/*
GetTeamSettingsNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type GetTeamSettingsNotFound struct {
	Payload *models.APIErrorNotFoundExample
}

// IsSuccess returns true when this get team settings not found response has a 2xx status code
func (o *GetTeamSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team settings not found response has a 3xx status code
func (o *GetTeamSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team settings not found response has a 4xx status code
func (o *GetTeamSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team settings not found response has a 5xx status code
func (o *GetTeamSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get team settings not found response a status code equal to that given
func (o *GetTeamSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get team settings not found response
func (o *GetTeamSettingsNotFound) Code() int {
	return 404
}

func (o *GetTeamSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /teams/{teamID}/settings][%d] getTeamSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamSettingsNotFound) String() string {
	return fmt.Sprintf("[GET /teams/{teamID}/settings][%d] getTeamSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamSettingsNotFound) GetPayload() *models.APIErrorNotFoundExample {
	return o.Payload
}

func (o *GetTeamSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorNotFoundExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamSettingsUnprocessableEntity creates a GetTeamSettingsUnprocessableEntity with default headers values
func NewGetTeamSettingsUnprocessableEntity() *GetTeamSettingsUnprocessableEntity {
	return &GetTeamSettingsUnprocessableEntity{}
}

/*
GetTeamSettingsUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type GetTeamSettingsUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this get team settings unprocessable entity response has a 2xx status code
func (o *GetTeamSettingsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team settings unprocessable entity response has a 3xx status code
func (o *GetTeamSettingsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team settings unprocessable entity response has a 4xx status code
func (o *GetTeamSettingsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team settings unprocessable entity response has a 5xx status code
func (o *GetTeamSettingsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get team settings unprocessable entity response a status code equal to that given
func (o *GetTeamSettingsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get team settings unprocessable entity response
func (o *GetTeamSettingsUnprocessableEntity) Code() int {
	return 422
}

func (o *GetTeamSettingsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /teams/{teamID}/settings][%d] getTeamSettingsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetTeamSettingsUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /teams/{teamID}/settings][%d] getTeamSettingsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetTeamSettingsUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *GetTeamSettingsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamSettingsInternalServerError creates a GetTeamSettingsInternalServerError with default headers values
func NewGetTeamSettingsInternalServerError() *GetTeamSettingsInternalServerError {
	return &GetTeamSettingsInternalServerError{}
}

/*
GetTeamSettingsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetTeamSettingsInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this get team settings internal server error response has a 2xx status code
func (o *GetTeamSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team settings internal server error response has a 3xx status code
func (o *GetTeamSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team settings internal server error response has a 4xx status code
func (o *GetTeamSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team settings internal server error response has a 5xx status code
func (o *GetTeamSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get team settings internal server error response a status code equal to that given
func (o *GetTeamSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get team settings internal server error response
func (o *GetTeamSettingsInternalServerError) Code() int {
	return 500
}

func (o *GetTeamSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /teams/{teamID}/settings][%d] getTeamSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTeamSettingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /teams/{teamID}/settings][%d] getTeamSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTeamSettingsInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *GetTeamSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
