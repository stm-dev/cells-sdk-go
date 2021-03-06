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

// FrontPluginsReader is a Reader for the FrontPlugins structure.
type FrontPluginsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FrontPluginsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFrontPluginsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFrontPluginsOK creates a FrontPluginsOK with default headers values
func NewFrontPluginsOK() *FrontPluginsOK {
	return &FrontPluginsOK{}
}

/*FrontPluginsOK handles this case with default header values.

FrontPluginsOK front plugins o k
*/
type FrontPluginsOK struct {
	Payload models.RestFrontPluginsResponse
}

func (o *FrontPluginsOK) Error() string {
	return fmt.Sprintf("[GET /frontend/plugins/{Lang}][%d] frontPluginsOK  %+v", 200, o.Payload)
}

func (o *FrontPluginsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
