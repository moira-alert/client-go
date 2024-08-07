// Code generated by go-swagger; DO NOT EDIT.

package config

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

// GetWebConfigReader is a Reader for the GetWebConfig structure.
type GetWebConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewGetWebConfigUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /config] get-web-config", response, response.Code())
	}
}

// NewGetWebConfigOK creates a GetWebConfigOK with default headers values
func NewGetWebConfigOK() *GetWebConfigOK {
	return &GetWebConfigOK{}
}

/*
GetWebConfigOK describes a response with status code 200, with default header values.

Configuration fetched successfully
*/
type GetWebConfigOK struct {
	Payload *models.APIWebConfig
}

// IsSuccess returns true when this get web config o k response has a 2xx status code
func (o *GetWebConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get web config o k response has a 3xx status code
func (o *GetWebConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get web config o k response has a 4xx status code
func (o *GetWebConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get web config o k response has a 5xx status code
func (o *GetWebConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get web config o k response a status code equal to that given
func (o *GetWebConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get web config o k response
func (o *GetWebConfigOK) Code() int {
	return 200
}

func (o *GetWebConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /config][%d] getWebConfigOK %s", 200, payload)
}

func (o *GetWebConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /config][%d] getWebConfigOK %s", 200, payload)
}

func (o *GetWebConfigOK) GetPayload() *models.APIWebConfig {
	return o.Payload
}

func (o *GetWebConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIWebConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebConfigUnprocessableEntity creates a GetWebConfigUnprocessableEntity with default headers values
func NewGetWebConfigUnprocessableEntity() *GetWebConfigUnprocessableEntity {
	return &GetWebConfigUnprocessableEntity{}
}

/*
GetWebConfigUnprocessableEntity describes a response with status code 422, with default header values.

Render error
*/
type GetWebConfigUnprocessableEntity struct {
	Payload *models.APIErrorRenderExample
}

// IsSuccess returns true when this get web config unprocessable entity response has a 2xx status code
func (o *GetWebConfigUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get web config unprocessable entity response has a 3xx status code
func (o *GetWebConfigUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get web config unprocessable entity response has a 4xx status code
func (o *GetWebConfigUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get web config unprocessable entity response has a 5xx status code
func (o *GetWebConfigUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get web config unprocessable entity response a status code equal to that given
func (o *GetWebConfigUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get web config unprocessable entity response
func (o *GetWebConfigUnprocessableEntity) Code() int {
	return 422
}

func (o *GetWebConfigUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /config][%d] getWebConfigUnprocessableEntity %s", 422, payload)
}

func (o *GetWebConfigUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /config][%d] getWebConfigUnprocessableEntity %s", 422, payload)
}

func (o *GetWebConfigUnprocessableEntity) GetPayload() *models.APIErrorRenderExample {
	return o.Payload
}

func (o *GetWebConfigUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorRenderExample)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
