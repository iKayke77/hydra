// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// UpdateJSONWebKeyReader is a Reader for the UpdateJSONWebKey structure.
type UpdateJSONWebKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateJSONWebKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateJSONWebKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateJSONWebKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateJSONWebKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateJSONWebKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateJSONWebKeyOK creates a UpdateJSONWebKeyOK with default headers values
func NewUpdateJSONWebKeyOK() *UpdateJSONWebKeyOK {
	return &UpdateJSONWebKeyOK{}
}

/*UpdateJSONWebKeyOK handles this case with default header values.

JSONWebKey
*/
type UpdateJSONWebKeyOK struct {
	Payload *models.JSONWebKey
}

func (o *UpdateJSONWebKeyOK) Error() string {
	return fmt.Sprintf("[PUT /keys/{set}/{kid}][%d] updateJsonWebKeyOK  %+v", 200, o.Payload)
}

func (o *UpdateJSONWebKeyOK) GetPayload() *models.JSONWebKey {
	return o.Payload
}

func (o *UpdateJSONWebKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONWebKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJSONWebKeyUnauthorized creates a UpdateJSONWebKeyUnauthorized with default headers values
func NewUpdateJSONWebKeyUnauthorized() *UpdateJSONWebKeyUnauthorized {
	return &UpdateJSONWebKeyUnauthorized{}
}

/*UpdateJSONWebKeyUnauthorized handles this case with default header values.

genericError
*/
type UpdateJSONWebKeyUnauthorized struct {
	Payload *models.GenericError
}

func (o *UpdateJSONWebKeyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /keys/{set}/{kid}][%d] updateJsonWebKeyUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateJSONWebKeyUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateJSONWebKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJSONWebKeyForbidden creates a UpdateJSONWebKeyForbidden with default headers values
func NewUpdateJSONWebKeyForbidden() *UpdateJSONWebKeyForbidden {
	return &UpdateJSONWebKeyForbidden{}
}

/*UpdateJSONWebKeyForbidden handles this case with default header values.

genericError
*/
type UpdateJSONWebKeyForbidden struct {
	Payload *models.GenericError
}

func (o *UpdateJSONWebKeyForbidden) Error() string {
	return fmt.Sprintf("[PUT /keys/{set}/{kid}][%d] updateJsonWebKeyForbidden  %+v", 403, o.Payload)
}

func (o *UpdateJSONWebKeyForbidden) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateJSONWebKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJSONWebKeyInternalServerError creates a UpdateJSONWebKeyInternalServerError with default headers values
func NewUpdateJSONWebKeyInternalServerError() *UpdateJSONWebKeyInternalServerError {
	return &UpdateJSONWebKeyInternalServerError{}
}

/*UpdateJSONWebKeyInternalServerError handles this case with default header values.

genericError
*/
type UpdateJSONWebKeyInternalServerError struct {
	Payload *models.GenericError
}

func (o *UpdateJSONWebKeyInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /keys/{set}/{kid}][%d] updateJsonWebKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateJSONWebKeyInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateJSONWebKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
