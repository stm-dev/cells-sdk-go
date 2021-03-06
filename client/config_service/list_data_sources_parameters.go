// Code generated by go-swagger; DO NOT EDIT.

package config_service

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
)

// NewListDataSourcesParams creates a new ListDataSourcesParams object
// with the default values initialized.
func NewListDataSourcesParams() *ListDataSourcesParams {

	return &ListDataSourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListDataSourcesParamsWithTimeout creates a new ListDataSourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListDataSourcesParamsWithTimeout(timeout time.Duration) *ListDataSourcesParams {

	return &ListDataSourcesParams{

		timeout: timeout,
	}
}

// NewListDataSourcesParamsWithContext creates a new ListDataSourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListDataSourcesParamsWithContext(ctx context.Context) *ListDataSourcesParams {

	return &ListDataSourcesParams{

		Context: ctx,
	}
}

// NewListDataSourcesParamsWithHTTPClient creates a new ListDataSourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListDataSourcesParamsWithHTTPClient(client *http.Client) *ListDataSourcesParams {

	return &ListDataSourcesParams{
		HTTPClient: client,
	}
}

/*ListDataSourcesParams contains all the parameters to send to the API endpoint
for the list data sources operation typically these are written to a http.Request
*/
type ListDataSourcesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list data sources params
func (o *ListDataSourcesParams) WithTimeout(timeout time.Duration) *ListDataSourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list data sources params
func (o *ListDataSourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list data sources params
func (o *ListDataSourcesParams) WithContext(ctx context.Context) *ListDataSourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list data sources params
func (o *ListDataSourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list data sources params
func (o *ListDataSourcesParams) WithHTTPClient(client *http.Client) *ListDataSourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list data sources params
func (o *ListDataSourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListDataSourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
