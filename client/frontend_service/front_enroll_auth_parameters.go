// Code generated by go-swagger; DO NOT EDIT.

package frontend_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// NewFrontEnrollAuthParams creates a new FrontEnrollAuthParams object
// with the default values initialized.
func NewFrontEnrollAuthParams() *FrontEnrollAuthParams {
	var ()
	return &FrontEnrollAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFrontEnrollAuthParamsWithTimeout creates a new FrontEnrollAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFrontEnrollAuthParamsWithTimeout(timeout time.Duration) *FrontEnrollAuthParams {
	var ()
	return &FrontEnrollAuthParams{

		timeout: timeout,
	}
}

// NewFrontEnrollAuthParamsWithContext creates a new FrontEnrollAuthParams object
// with the default values initialized, and the ability to set a context for a request
func NewFrontEnrollAuthParamsWithContext(ctx context.Context) *FrontEnrollAuthParams {
	var ()
	return &FrontEnrollAuthParams{

		Context: ctx,
	}
}

// NewFrontEnrollAuthParamsWithHTTPClient creates a new FrontEnrollAuthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFrontEnrollAuthParamsWithHTTPClient(client *http.Client) *FrontEnrollAuthParams {
	var ()
	return &FrontEnrollAuthParams{
		HTTPClient: client,
	}
}

/*FrontEnrollAuthParams contains all the parameters to send to the API endpoint
for the front enroll auth operation typically these are written to a http.Request
*/
type FrontEnrollAuthParams struct {

	/*Body*/
	Body *models.RestFrontEnrollAuthRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the front enroll auth params
func (o *FrontEnrollAuthParams) WithTimeout(timeout time.Duration) *FrontEnrollAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the front enroll auth params
func (o *FrontEnrollAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the front enroll auth params
func (o *FrontEnrollAuthParams) WithContext(ctx context.Context) *FrontEnrollAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the front enroll auth params
func (o *FrontEnrollAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the front enroll auth params
func (o *FrontEnrollAuthParams) WithHTTPClient(client *http.Client) *FrontEnrollAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the front enroll auth params
func (o *FrontEnrollAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the front enroll auth params
func (o *FrontEnrollAuthParams) WithBody(body *models.RestFrontEnrollAuthRequest) *FrontEnrollAuthParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the front enroll auth params
func (o *FrontEnrollAuthParams) SetBody(body *models.RestFrontEnrollAuthRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FrontEnrollAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
