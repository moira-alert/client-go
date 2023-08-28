// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moira-alert/client-go/pkg/models"
)

// GetAllTagsReader is a Reader for the GetAllTags structure.
type GetAllTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewGetAllTagsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /tag] get-all-tags", response, response.Code())
	}
}

// NewGetAllTagsOK creates a GetAllTagsOK with default headers values
func NewGetAllTagsOK() *GetAllTagsOK {
	return &GetAllTagsOK{}
}

/*
GetAllTagsOK describes a response with status code 200, with default header values.

Tags fetched successfully
*/
type GetAllTagsOK struct {
	Payload *models.DtoTagsData
}

// IsSuccess returns true when this get all tags o k response has a 2xx status code
func (o *GetAllTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all tags o k response has a 3xx status code
func (o *GetAllTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tags o k response has a 4xx status code
func (o *GetAllTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all tags o k response has a 5xx status code
func (o *GetAllTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all tags o k response a status code equal to that given
func (o *GetAllTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all tags o k response
func (o *GetAllTagsOK) Code() int {
	return 200
}

func (o *GetAllTagsOK) Error() string {
	return fmt.Sprintf("[GET /tag][%d] getAllTagsOK  %+v", 200, o.Payload)
}

func (o *GetAllTagsOK) String() string {
	return fmt.Sprintf("[GET /tag][%d] getAllTagsOK  %+v", 200, o.Payload)
}

func (o *GetAllTagsOK) GetPayload() *models.DtoTagsData {
	return o.Payload
}

func (o *GetAllTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoTagsData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTagsUnprocessableEntity creates a GetAllTagsUnprocessableEntity with default headers values
func NewGetAllTagsUnprocessableEntity() *GetAllTagsUnprocessableEntity {
	return &GetAllTagsUnprocessableEntity{}
}

/*
GetAllTagsUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type GetAllTagsUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this get all tags unprocessable entity response has a 2xx status code
func (o *GetAllTagsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all tags unprocessable entity response has a 3xx status code
func (o *GetAllTagsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tags unprocessable entity response has a 4xx status code
func (o *GetAllTagsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all tags unprocessable entity response has a 5xx status code
func (o *GetAllTagsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get all tags unprocessable entity response a status code equal to that given
func (o *GetAllTagsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get all tags unprocessable entity response
func (o *GetAllTagsUnprocessableEntity) Code() int {
	return 422
}

func (o *GetAllTagsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /tag][%d] getAllTagsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetAllTagsUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /tag][%d] getAllTagsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetAllTagsUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *GetAllTagsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTagsInternalServerError creates a GetAllTagsInternalServerError with default headers values
func NewGetAllTagsInternalServerError() *GetAllTagsInternalServerError {
	return &GetAllTagsInternalServerError{}
}

/*
GetAllTagsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAllTagsInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this get all tags internal server error response has a 2xx status code
func (o *GetAllTagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all tags internal server error response has a 3xx status code
func (o *GetAllTagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tags internal server error response has a 4xx status code
func (o *GetAllTagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all tags internal server error response has a 5xx status code
func (o *GetAllTagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all tags internal server error response a status code equal to that given
func (o *GetAllTagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get all tags internal server error response
func (o *GetAllTagsInternalServerError) Code() int {
	return 500
}

func (o *GetAllTagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tag][%d] getAllTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllTagsInternalServerError) String() string {
	return fmt.Sprintf("[GET /tag][%d] getAllTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllTagsInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *GetAllTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}