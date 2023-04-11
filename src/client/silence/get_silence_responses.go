// Code generated by go-swagger; DO NOT EDIT.

package silence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"geekxflood/alertmanager-cli/src/models"
)

// GetSilenceReader is a Reader for the GetSilence structure.
type GetSilenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSilenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSilenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSilenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSilenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSilenceOK creates a GetSilenceOK with default headers values
func NewGetSilenceOK() *GetSilenceOK {
	return &GetSilenceOK{}
}

/*
GetSilenceOK describes a response with status code 200, with default header values.

Get silence response
*/
type GetSilenceOK struct {
	Payload *models.GettableSilence
}

// IsSuccess returns true when this get silence o k response has a 2xx status code
func (o *GetSilenceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get silence o k response has a 3xx status code
func (o *GetSilenceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get silence o k response has a 4xx status code
func (o *GetSilenceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get silence o k response has a 5xx status code
func (o *GetSilenceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get silence o k response a status code equal to that given
func (o *GetSilenceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSilenceOK) Error() string {
	return fmt.Sprintf("[GET /silence/{silenceID}][%d] getSilenceOK  %+v", 200, o.Payload)
}

func (o *GetSilenceOK) String() string {
	return fmt.Sprintf("[GET /silence/{silenceID}][%d] getSilenceOK  %+v", 200, o.Payload)
}

func (o *GetSilenceOK) GetPayload() *models.GettableSilence {
	return o.Payload
}

func (o *GetSilenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GettableSilence)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSilenceNotFound creates a GetSilenceNotFound with default headers values
func NewGetSilenceNotFound() *GetSilenceNotFound {
	return &GetSilenceNotFound{}
}

/*
GetSilenceNotFound describes a response with status code 404, with default header values.

A silence with the specified ID was not found
*/
type GetSilenceNotFound struct {
}

// IsSuccess returns true when this get silence not found response has a 2xx status code
func (o *GetSilenceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get silence not found response has a 3xx status code
func (o *GetSilenceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get silence not found response has a 4xx status code
func (o *GetSilenceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get silence not found response has a 5xx status code
func (o *GetSilenceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get silence not found response a status code equal to that given
func (o *GetSilenceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSilenceNotFound) Error() string {
	return fmt.Sprintf("[GET /silence/{silenceID}][%d] getSilenceNotFound ", 404)
}

func (o *GetSilenceNotFound) String() string {
	return fmt.Sprintf("[GET /silence/{silenceID}][%d] getSilenceNotFound ", 404)
}

func (o *GetSilenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSilenceInternalServerError creates a GetSilenceInternalServerError with default headers values
func NewGetSilenceInternalServerError() *GetSilenceInternalServerError {
	return &GetSilenceInternalServerError{}
}

/*
GetSilenceInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSilenceInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get silence internal server error response has a 2xx status code
func (o *GetSilenceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get silence internal server error response has a 3xx status code
func (o *GetSilenceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get silence internal server error response has a 4xx status code
func (o *GetSilenceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get silence internal server error response has a 5xx status code
func (o *GetSilenceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get silence internal server error response a status code equal to that given
func (o *GetSilenceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSilenceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /silence/{silenceID}][%d] getSilenceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSilenceInternalServerError) String() string {
	return fmt.Sprintf("[GET /silence/{silenceID}][%d] getSilenceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSilenceInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetSilenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
