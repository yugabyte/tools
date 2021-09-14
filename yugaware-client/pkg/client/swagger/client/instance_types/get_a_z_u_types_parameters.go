// Code generated by go-swagger; DO NOT EDIT.

package instance_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetAZUTypesParams creates a new GetAZUTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAZUTypesParams() *GetAZUTypesParams {
	return &GetAZUTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAZUTypesParamsWithTimeout creates a new GetAZUTypesParams object
// with the ability to set a timeout on a request.
func NewGetAZUTypesParamsWithTimeout(timeout time.Duration) *GetAZUTypesParams {
	return &GetAZUTypesParams{
		timeout: timeout,
	}
}

// NewGetAZUTypesParamsWithContext creates a new GetAZUTypesParams object
// with the ability to set a context for a request.
func NewGetAZUTypesParamsWithContext(ctx context.Context) *GetAZUTypesParams {
	return &GetAZUTypesParams{
		Context: ctx,
	}
}

// NewGetAZUTypesParamsWithHTTPClient creates a new GetAZUTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAZUTypesParamsWithHTTPClient(client *http.Client) *GetAZUTypesParams {
	return &GetAZUTypesParams{
		HTTPClient: client,
	}
}

/* GetAZUTypesParams contains all the parameters to send to the API endpoint
   for the get a z u types operation.

   Typically these are written to a http.Request.
*/
type GetAZUTypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get a z u types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAZUTypesParams) WithDefaults() *GetAZUTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get a z u types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAZUTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get a z u types params
func (o *GetAZUTypesParams) WithTimeout(timeout time.Duration) *GetAZUTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get a z u types params
func (o *GetAZUTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get a z u types params
func (o *GetAZUTypesParams) WithContext(ctx context.Context) *GetAZUTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get a z u types params
func (o *GetAZUTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get a z u types params
func (o *GetAZUTypesParams) WithHTTPClient(client *http.Client) *GetAZUTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get a z u types params
func (o *GetAZUTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAZUTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
