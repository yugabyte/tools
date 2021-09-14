// Code generated by go-swagger; DO NOT EDIT.

package universe_information

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

// NewGetLiveQueriesParams creates a new GetLiveQueriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLiveQueriesParams() *GetLiveQueriesParams {
	return &GetLiveQueriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLiveQueriesParamsWithTimeout creates a new GetLiveQueriesParams object
// with the ability to set a timeout on a request.
func NewGetLiveQueriesParamsWithTimeout(timeout time.Duration) *GetLiveQueriesParams {
	return &GetLiveQueriesParams{
		timeout: timeout,
	}
}

// NewGetLiveQueriesParamsWithContext creates a new GetLiveQueriesParams object
// with the ability to set a context for a request.
func NewGetLiveQueriesParamsWithContext(ctx context.Context) *GetLiveQueriesParams {
	return &GetLiveQueriesParams{
		Context: ctx,
	}
}

// NewGetLiveQueriesParamsWithHTTPClient creates a new GetLiveQueriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLiveQueriesParamsWithHTTPClient(client *http.Client) *GetLiveQueriesParams {
	return &GetLiveQueriesParams{
		HTTPClient: client,
	}
}

/* GetLiveQueriesParams contains all the parameters to send to the API endpoint
   for the get live queries operation.

   Typically these are written to a http.Request.
*/
type GetLiveQueriesParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	// UniUUID.
	//
	// Format: uuid
	UniUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get live queries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLiveQueriesParams) WithDefaults() *GetLiveQueriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get live queries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLiveQueriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get live queries params
func (o *GetLiveQueriesParams) WithTimeout(timeout time.Duration) *GetLiveQueriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get live queries params
func (o *GetLiveQueriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get live queries params
func (o *GetLiveQueriesParams) WithContext(ctx context.Context) *GetLiveQueriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get live queries params
func (o *GetLiveQueriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get live queries params
func (o *GetLiveQueriesParams) WithHTTPClient(client *http.Client) *GetLiveQueriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get live queries params
func (o *GetLiveQueriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the get live queries params
func (o *GetLiveQueriesParams) WithCUUID(cUUID strfmt.UUID) *GetLiveQueriesParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the get live queries params
func (o *GetLiveQueriesParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithUniUUID adds the uniUUID to the get live queries params
func (o *GetLiveQueriesParams) WithUniUUID(uniUUID strfmt.UUID) *GetLiveQueriesParams {
	o.SetUniUUID(uniUUID)
	return o
}

// SetUniUUID adds the uniUuid to the get live queries params
func (o *GetLiveQueriesParams) SetUniUUID(uniUUID strfmt.UUID) {
	o.UniUUID = uniUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLiveQueriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	// path param uniUUID
	if err := r.SetPathParam("uniUUID", o.UniUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
