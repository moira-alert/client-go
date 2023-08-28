// Code generated by go-swagger; DO NOT EDIT.

package contact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moira-alert/client-go/pkg/models"
)

// GetAllContactsReader is a Reader for the GetAllContacts structure.
type GetAllContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewGetAllContactsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /contact] get-all-contacts", response, response.Code())
	}
}

// NewGetAllContactsOK creates a GetAllContactsOK with default headers values
func NewGetAllContactsOK() *GetAllContactsOK {
	return &GetAllContactsOK{}
}

/*
GetAllContactsOK describes a response with status code 200, with default header values.

Contacts fetched successfully
*/
type GetAllContactsOK struct {
	Payload *models.DtoContactList
}

// IsSuccess returns true when this get all contacts o k response has a 2xx status code
func (o *GetAllContactsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all contacts o k response has a 3xx status code
func (o *GetAllContactsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all contacts o k response has a 4xx status code
func (o *GetAllContactsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all contacts o k response has a 5xx status code
func (o *GetAllContactsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all contacts o k response a status code equal to that given
func (o *GetAllContactsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all contacts o k response
func (o *GetAllContactsOK) Code() int {
	return 200
}

func (o *GetAllContactsOK) Error() string {
	return fmt.Sprintf("[GET /contact][%d] getAllContactsOK  %+v", 200, o.Payload)
}

func (o *GetAllContactsOK) String() string {
	return fmt.Sprintf("[GET /contact][%d] getAllContactsOK  %+v", 200, o.Payload)
}

func (o *GetAllContactsOK) GetPayload() *models.DtoContactList {
	return o.Payload
}

func (o *GetAllContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtoContactList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllContactsUnprocessableEntity creates a GetAllContactsUnprocessableEntity with default headers values
func NewGetAllContactsUnprocessableEntity() *GetAllContactsUnprocessableEntity {
	return &GetAllContactsUnprocessableEntity{}
}

/*
GetAllContactsUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type GetAllContactsUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this get all contacts unprocessable entity response has a 2xx status code
func (o *GetAllContactsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all contacts unprocessable entity response has a 3xx status code
func (o *GetAllContactsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all contacts unprocessable entity response has a 4xx status code
func (o *GetAllContactsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all contacts unprocessable entity response has a 5xx status code
func (o *GetAllContactsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get all contacts unprocessable entity response a status code equal to that given
func (o *GetAllContactsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get all contacts unprocessable entity response
func (o *GetAllContactsUnprocessableEntity) Code() int {
	return 422
}

func (o *GetAllContactsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /contact][%d] getAllContactsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetAllContactsUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /contact][%d] getAllContactsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetAllContactsUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *GetAllContactsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllContactsInternalServerError creates a GetAllContactsInternalServerError with default headers values
func NewGetAllContactsInternalServerError() *GetAllContactsInternalServerError {
	return &GetAllContactsInternalServerError{}
}

/*
GetAllContactsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAllContactsInternalServerError struct {
	Payload *models.APIErrorInternalServerExample
}

// IsSuccess returns true when this get all contacts internal server error response has a 2xx status code
func (o *GetAllContactsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all contacts internal server error response has a 3xx status code
func (o *GetAllContactsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all contacts internal server error response has a 4xx status code
func (o *GetAllContactsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all contacts internal server error response has a 5xx status code
func (o *GetAllContactsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all contacts internal server error response a status code equal to that given
func (o *GetAllContactsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get all contacts internal server error response
func (o *GetAllContactsInternalServerError) Code() int {
	return 500
}

func (o *GetAllContactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /contact][%d] getAllContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllContactsInternalServerError) String() string {
	return fmt.Sprintf("[GET /contact][%d] getAllContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllContactsInternalServerError) GetPayload() *models.APIErrorInternalServerExample {
	return o.Payload
}

func (o *GetAllContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorInternalServerExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
