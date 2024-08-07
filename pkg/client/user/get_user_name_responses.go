// Code generated by go-swagger; DO NOT EDIT.

package user

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

// GetUserNameReader is a Reader for the GetUserName structure.
type GetUserNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewGetUserNameUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /user] get-user-name", response, response.Code())
	}
}

// NewGetUserNameOK creates a GetUserNameOK with default headers values
func NewGetUserNameOK() *GetUserNameOK {
	return &GetUserNameOK{}
}

/*
GetUserNameOK describes a response with status code 200, with default header values.

User name fetched successfully
*/
type GetUserNameOK struct {
	Payload *models.DtoUser
}

// IsSuccess returns true when this get user name o k response has a 2xx status code
func (o *GetUserNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user name o k response has a 3xx status code
func (o *GetUserNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user name o k response has a 4xx status code
func (o *GetUserNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user name o k response has a 5xx status code
func (o *GetUserNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user name o k response a status code equal to that given
func (o *GetUserNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user name o k response
func (o *GetUserNameOK) Code() int {
	return 200
}

func (o *GetUserNameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user][%d] getUserNameOK %s", 200, payload)
}

func (o *GetUserNameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user][%d] getUserNameOK %s", 200, payload)
}

func (o *GetUserNameOK) GetPayload() *models.DtoUser {
	return o.Payload
}

func (o *GetUserNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserNameUnprocessableEntity creates a GetUserNameUnprocessableEntity with default headers values
func NewGetUserNameUnprocessableEntity() *GetUserNameUnprocessableEntity {
	return &GetUserNameUnprocessableEntity{}
}

/*
GetUserNameUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type GetUserNameUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this get user name unprocessable entity response has a 2xx status code
func (o *GetUserNameUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user name unprocessable entity response has a 3xx status code
func (o *GetUserNameUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user name unprocessable entity response has a 4xx status code
func (o *GetUserNameUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user name unprocessable entity response has a 5xx status code
func (o *GetUserNameUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get user name unprocessable entity response a status code equal to that given
func (o *GetUserNameUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get user name unprocessable entity response
func (o *GetUserNameUnprocessableEntity) Code() int {
	return 422
}

func (o *GetUserNameUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user][%d] getUserNameUnprocessableEntity %s", 422, payload)
}

func (o *GetUserNameUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user][%d] getUserNameUnprocessableEntity %s", 422, payload)
}

func (o *GetUserNameUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *GetUserNameUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
