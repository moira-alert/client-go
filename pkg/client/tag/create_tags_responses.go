// Code generated by go-swagger; DO NOT EDIT.

package tag

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

// CreateTagsReader is a Reader for the CreateTags structure.
type CreateTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateTagsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /tag] create-tags", response, response.Code())
	}
}

// NewCreateTagsOK creates a CreateTagsOK with default headers values
func NewCreateTagsOK() *CreateTagsOK {
	return &CreateTagsOK{}
}

/*
CreateTagsOK describes a response with status code 200, with default header values.

Create tags successfully
*/
type CreateTagsOK struct {
}

// IsSuccess returns true when this create tags o k response has a 2xx status code
func (o *CreateTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create tags o k response has a 3xx status code
func (o *CreateTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tags o k response has a 4xx status code
func (o *CreateTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create tags o k response has a 5xx status code
func (o *CreateTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create tags o k response a status code equal to that given
func (o *CreateTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create tags o k response
func (o *CreateTagsOK) Code() int {
	return 200
}

func (o *CreateTagsOK) Error() string {
	return fmt.Sprintf("[POST /tag][%d] createTagsOK", 200)
}

func (o *CreateTagsOK) String() string {
	return fmt.Sprintf("[POST /tag][%d] createTagsOK", 200)
}

func (o *CreateTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTagsBadRequest creates a CreateTagsBadRequest with default headers values
func NewCreateTagsBadRequest() *CreateTagsBadRequest {
	return &CreateTagsBadRequest{}
}

/*
CreateTagsBadRequest describes a response with status code 400, with default header values.

Bad request from client
*/
type CreateTagsBadRequest struct {
	Payload *models.APIErrorInvalidRequestExample
}

// IsSuccess returns true when this create tags bad request response has a 2xx status code
func (o *CreateTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tags bad request response has a 3xx status code
func (o *CreateTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tags bad request response has a 4xx status code
func (o *CreateTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tags bad request response has a 5xx status code
func (o *CreateTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create tags bad request response a status code equal to that given
func (o *CreateTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create tags bad request response
func (o *CreateTagsBadRequest) Code() int {
	return 400
}

func (o *CreateTagsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tag][%d] createTagsBadRequest %s", 400, payload)
}

func (o *CreateTagsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tag][%d] createTagsBadRequest %s", 400, payload)
}

func (o *CreateTagsBadRequest) GetPayload() *models.APIErrorInvalidRequestExample {
	return o.Payload
}

func (o *CreateTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInvalidRequestExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagsUnprocessableEntity creates a CreateTagsUnprocessableEntity with default headers values
func NewCreateTagsUnprocessableEntity() *CreateTagsUnprocessableEntity {
	return &CreateTagsUnprocessableEntity{}
}

/*
CreateTagsUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type CreateTagsUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this create tags unprocessable entity response has a 2xx status code
func (o *CreateTagsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tags unprocessable entity response has a 3xx status code
func (o *CreateTagsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tags unprocessable entity response has a 4xx status code
func (o *CreateTagsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tags unprocessable entity response has a 5xx status code
func (o *CreateTagsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create tags unprocessable entity response a status code equal to that given
func (o *CreateTagsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create tags unprocessable entity response
func (o *CreateTagsUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateTagsUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tag][%d] createTagsUnprocessableEntity %s", 422, payload)
}

func (o *CreateTagsUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tag][%d] createTagsUnprocessableEntity %s", 422, payload)
}

func (o *CreateTagsUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *CreateTagsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagsInternalServerError creates a CreateTagsInternalServerError with default headers values
func NewCreateTagsInternalServerError() *CreateTagsInternalServerError {
	return &CreateTagsInternalServerError{}
}

/*
CreateTagsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateTagsInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this create tags internal server error response has a 2xx status code
func (o *CreateTagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tags internal server error response has a 3xx status code
func (o *CreateTagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tags internal server error response has a 4xx status code
func (o *CreateTagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create tags internal server error response has a 5xx status code
func (o *CreateTagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create tags internal server error response a status code equal to that given
func (o *CreateTagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create tags internal server error response
func (o *CreateTagsInternalServerError) Code() int {
	return 500
}

func (o *CreateTagsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tag][%d] createTagsInternalServerError %s", 500, payload)
}

func (o *CreateTagsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tag][%d] createTagsInternalServerError %s", 500, payload)
}

func (o *CreateTagsInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *CreateTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
