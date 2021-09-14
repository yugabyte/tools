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

// NewGetTaskAuditParams creates a new GetTaskAuditParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTaskAuditParams() *GetTaskAuditParams {
	return &GetTaskAuditParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskAuditParamsWithTimeout creates a new GetTaskAuditParams object
// with the ability to set a timeout on a request.
func NewGetTaskAuditParamsWithTimeout(timeout time.Duration) *GetTaskAuditParams {
	return &GetTaskAuditParams{
		timeout: timeout,
	}
}

// NewGetTaskAuditParamsWithContext creates a new GetTaskAuditParams object
// with the ability to set a context for a request.
func NewGetTaskAuditParamsWithContext(ctx context.Context) *GetTaskAuditParams {
	return &GetTaskAuditParams{
		Context: ctx,
	}
}

// NewGetTaskAuditParamsWithHTTPClient creates a new GetTaskAuditParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTaskAuditParamsWithHTTPClient(client *http.Client) *GetTaskAuditParams {
	return &GetTaskAuditParams{
		HTTPClient: client,
	}
}

/* GetTaskAuditParams contains all the parameters to send to the API endpoint
   for the get task audit operation.

   Typically these are written to a http.Request.
*/
type GetTaskAuditParams struct {

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

// WithDefaults hydrates default values in the get task audit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskAuditParams) WithDefaults() *GetTaskAuditParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get task audit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTaskAuditParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get task audit params
func (o *GetTaskAuditParams) WithTimeout(timeout time.Duration) *GetTaskAuditParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task audit params
func (o *GetTaskAuditParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task audit params
func (o *GetTaskAuditParams) WithContext(ctx context.Context) *GetTaskAuditParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task audit params
func (o *GetTaskAuditParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get task audit params
func (o *GetTaskAuditParams) WithHTTPClient(client *http.Client) *GetTaskAuditParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get task audit params
func (o *GetTaskAuditParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the get task audit params
func (o *GetTaskAuditParams) WithCUUID(cUUID strfmt.UUID) *GetTaskAuditParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the get task audit params
func (o *GetTaskAuditParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithTUUID adds the tUUID to the get task audit params
func (o *GetTaskAuditParams) WithTUUID(tUUID strfmt.UUID) *GetTaskAuditParams {
	o.SetTUUID(tUUID)
	return o
}

// SetTUUID adds the tUuid to the get task audit params
func (o *GetTaskAuditParams) SetTUUID(tUUID strfmt.UUID) {
	o.TUUID = tUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskAuditParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
