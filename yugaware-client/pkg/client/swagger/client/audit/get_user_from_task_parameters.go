// Code generated by go-swagger; DO NOT EDIT.

package audit

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

// NewGetUserFromTaskParams creates a new GetUserFromTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserFromTaskParams() *GetUserFromTaskParams {
	return &GetUserFromTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserFromTaskParamsWithTimeout creates a new GetUserFromTaskParams object
// with the ability to set a timeout on a request.
func NewGetUserFromTaskParamsWithTimeout(timeout time.Duration) *GetUserFromTaskParams {
	return &GetUserFromTaskParams{
		timeout: timeout,
	}
}

// NewGetUserFromTaskParamsWithContext creates a new GetUserFromTaskParams object
// with the ability to set a context for a request.
func NewGetUserFromTaskParamsWithContext(ctx context.Context) *GetUserFromTaskParams {
	return &GetUserFromTaskParams{
		Context: ctx,
	}
}

// NewGetUserFromTaskParamsWithHTTPClient creates a new GetUserFromTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserFromTaskParamsWithHTTPClient(client *http.Client) *GetUserFromTaskParams {
	return &GetUserFromTaskParams{
		HTTPClient: client,
	}
}

/* GetUserFromTaskParams contains all the parameters to send to the API endpoint
   for the get user from task operation.

   Typically these are written to a http.Request.
*/
type GetUserFromTaskParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	// TUUID.
	//
	// Format: uuid
	TUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user from task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserFromTaskParams) WithDefaults() *GetUserFromTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user from task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserFromTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user from task params
func (o *GetUserFromTaskParams) WithTimeout(timeout time.Duration) *GetUserFromTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user from task params
func (o *GetUserFromTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user from task params
func (o *GetUserFromTaskParams) WithContext(ctx context.Context) *GetUserFromTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user from task params
func (o *GetUserFromTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user from task params
func (o *GetUserFromTaskParams) WithHTTPClient(client *http.Client) *GetUserFromTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user from task params
func (o *GetUserFromTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the get user from task params
func (o *GetUserFromTaskParams) WithCUUID(cUUID strfmt.UUID) *GetUserFromTaskParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the get user from task params
func (o *GetUserFromTaskParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithTUUID adds the tUUID to the get user from task params
func (o *GetUserFromTaskParams) WithTUUID(tUUID strfmt.UUID) *GetUserFromTaskParams {
	o.SetTUUID(tUUID)
	return o
}

// SetTUUID adds the tUuid to the get user from task params
func (o *GetUserFromTaskParams) SetTUUID(tUUID strfmt.UUID) {
	o.TUUID = tUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserFromTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	// path param tUUID
	if err := r.SetPathParam("tUUID", o.TUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
