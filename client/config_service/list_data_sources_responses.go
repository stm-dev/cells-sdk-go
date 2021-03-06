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

// ListDataSourcesReader is a Reader for the ListDataSources structure.
type ListDataSourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDataSourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListDataSourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListDataSourcesOK creates a ListDataSourcesOK with default headers values
func NewListDataSourcesOK() *ListDataSourcesOK {
	return &ListDataSourcesOK{}
}

/*ListDataSourcesOK handles this case with default header values.

ListDataSourcesOK list data sources o k
*/
type ListDataSourcesOK struct {
	Payload *models.RestDataSourceCollection
}

func (o *ListDataSourcesOK) Error() string {
	return fmt.Sprintf("[GET /config/datasource][%d] listDataSourcesOK  %+v", 200, o.Payload)
}

func (o *ListDataSourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestDataSourceCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
