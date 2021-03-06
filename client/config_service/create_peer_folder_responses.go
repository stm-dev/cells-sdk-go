// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// CreatePeerFolderReader is a Reader for the CreatePeerFolder structure.
type CreatePeerFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePeerFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreatePeerFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePeerFolderOK creates a CreatePeerFolderOK with default headers values
func NewCreatePeerFolderOK() *CreatePeerFolderOK {
	return &CreatePeerFolderOK{}
}

/*CreatePeerFolderOK handles this case with default header values.

CreatePeerFolderOK create peer folder o k
*/
type CreatePeerFolderOK struct {
	Payload *models.RestCreatePeerFolderResponse
}

func (o *CreatePeerFolderOK) Error() string {
	return fmt.Sprintf("[PUT /config/peers/{PeerAddress}][%d] createPeerFolderOK  %+v", 200, o.Payload)
}

func (o *CreatePeerFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestCreatePeerFolderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
