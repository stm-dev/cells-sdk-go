// Code generated by go-swagger; DO NOT EDIT.

package frontend_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// SettingsMenuReader is a Reader for the SettingsMenu structure.
type SettingsMenuReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SettingsMenuReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSettingsMenuOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSettingsMenuOK creates a SettingsMenuOK with default headers values
func NewSettingsMenuOK() *SettingsMenuOK {
	return &SettingsMenuOK{}
}

/*SettingsMenuOK handles this case with default header values.

SettingsMenuOK settings menu o k
*/
type SettingsMenuOK struct {
	Payload *models.RestSettingsMenuResponse
}

func (o *SettingsMenuOK) Error() string {
	return fmt.Sprintf("[GET /frontend/settings-menu][%d] settingsMenuOK  %+v", 200, o.Payload)
}

func (o *SettingsMenuOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestSettingsMenuResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}