// Code generated by go-swagger; DO NOT EDIT.

package pattern

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

// GetAllPatternsReader is a Reader for the GetAllPatterns structure.
type GetAllPatternsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllPatternsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllPatternsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAllPatternsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAllPatternsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllPatternsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pattern] get-all-patterns", response, response.Code())
	}
}

// NewGetAllPatternsOK creates a GetAllPatternsOK with default headers values
func NewGetAllPatternsOK() *GetAllPatternsOK {
	return &GetAllPatternsOK{}
}

/*
GetAllPatternsOK describes a response with status code 200, with default header values.

Patterns fetched successfully
*/
type GetAllPatternsOK struct {
	Payload *models.DtoPatternList
}

// IsSuccess returns true when this get all patterns o k response has a 2xx status code
func (o *GetAllPatternsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all patterns o k response has a 3xx status code
func (o *GetAllPatternsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all patterns o k response has a 4xx status code
func (o *GetAllPatternsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all patterns o k response has a 5xx status code
func (o *GetAllPatternsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all patterns o k response a status code equal to that given
func (o *GetAllPatternsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all patterns o k response
func (o *GetAllPatternsOK) Code() int {
	return 200
}

func (o *GetAllPatternsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pattern][%d] getAllPatternsOK %s", 200, payload)
}

func (o *GetAllPatternsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pattern][%d] getAllPatternsOK %s", 200, payload)
}

func (o *GetAllPatternsOK) GetPayload() *models.DtoPatternList {
	return o.Payload
}

func (o *GetAllPatternsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoPatternList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPatternsForbidden creates a GetAllPatternsForbidden with default headers values
func NewGetAllPatternsForbidden() *GetAllPatternsForbidden {
	return &GetAllPatternsForbidden{}
}

/*
GetAllPatternsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllPatternsForbidden struct {
	Payload *models.APIErrorForbiddenExample
}

// IsSuccess returns true when this get all patterns forbidden response has a 2xx status code
func (o *GetAllPatternsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all patterns forbidden response has a 3xx status code
func (o *GetAllPatternsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all patterns forbidden response has a 4xx status code
func (o *GetAllPatternsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all patterns forbidden response has a 5xx status code
func (o *GetAllPatternsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all patterns forbidden response a status code equal to that given
func (o *GetAllPatternsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get all patterns forbidden response
func (o *GetAllPatternsForbidden) Code() int {
	return 403
}

func (o *GetAllPatternsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pattern][%d] getAllPatternsForbidden %s", 403, payload)
}

func (o *GetAllPatternsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pattern][%d] getAllPatternsForbidden %s", 403, payload)
}

func (o *GetAllPatternsForbidden) GetPayload() *models.APIErrorForbiddenExample {
	return o.Payload
}

func (o *GetAllPatternsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorForbiddenExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPatternsUnprocessableEntity creates a GetAllPatternsUnprocessableEntity with default headers values
func NewGetAllPatternsUnprocessableEntity() *GetAllPatternsUnprocessableEntity {
	return &GetAllPatternsUnprocessableEntity{}
}

/*
GetAllPatternsUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type GetAllPatternsUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this get all patterns unprocessable entity response has a 2xx status code
func (o *GetAllPatternsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all patterns unprocessable entity response has a 3xx status code
func (o *GetAllPatternsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all patterns unprocessable entity response has a 4xx status code
func (o *GetAllPatternsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all patterns unprocessable entity response has a 5xx status code
func (o *GetAllPatternsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get all patterns unprocessable entity response a status code equal to that given
func (o *GetAllPatternsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get all patterns unprocessable entity response
func (o *GetAllPatternsUnprocessableEntity) Code() int {
	return 422
}

func (o *GetAllPatternsUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pattern][%d] getAllPatternsUnprocessableEntity %s", 422, payload)
}

func (o *GetAllPatternsUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pattern][%d] getAllPatternsUnprocessableEntity %s", 422, payload)
}

func (o *GetAllPatternsUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *GetAllPatternsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPatternsInternalServerError creates a GetAllPatternsInternalServerError with default headers values
func NewGetAllPatternsInternalServerError() *GetAllPatternsInternalServerError {
	return &GetAllPatternsInternalServerError{}
}

/*
GetAllPatternsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAllPatternsInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this get all patterns internal server error response has a 2xx status code
func (o *GetAllPatternsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all patterns internal server error response has a 3xx status code
func (o *GetAllPatternsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all patterns internal server error response has a 4xx status code
func (o *GetAllPatternsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all patterns internal server error response has a 5xx status code
func (o *GetAllPatternsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all patterns internal server error response a status code equal to that given
func (o *GetAllPatternsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get all patterns internal server error response
func (o *GetAllPatternsInternalServerError) Code() int {
	return 500
}

func (o *GetAllPatternsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pattern][%d] getAllPatternsInternalServerError %s", 500, payload)
}

func (o *GetAllPatternsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pattern][%d] getAllPatternsInternalServerError %s", 500, payload)
}

func (o *GetAllPatternsInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *GetAllPatternsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
