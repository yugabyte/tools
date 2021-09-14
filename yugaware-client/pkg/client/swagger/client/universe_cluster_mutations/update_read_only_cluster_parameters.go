// Code generated by go-swagger; DO NOT EDIT.

package universe_cluster_mutations

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

// NewUpdateReadOnlyClusterParams creates a new UpdateReadOnlyClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateReadOnlyClusterParams() *UpdateReadOnlyClusterParams {
	return &UpdateReadOnlyClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateReadOnlyClusterParamsWithTimeout creates a new UpdateReadOnlyClusterParams object
// with the ability to set a timeout on a request.
func NewUpdateReadOnlyClusterParamsWithTimeout(timeout time.Duration) *UpdateReadOnlyClusterParams {
	return &UpdateReadOnlyClusterParams{
		timeout: timeout,
	}
}

// NewUpdateReadOnlyClusterParamsWithContext creates a new UpdateReadOnlyClusterParams object
// with the ability to set a context for a request.
func NewUpdateReadOnlyClusterParamsWithContext(ctx context.Context) *UpdateReadOnlyClusterParams {
	return &UpdateReadOnlyClusterParams{
		Context: ctx,
	}
}

// NewUpdateReadOnlyClusterParamsWithHTTPClient creates a new UpdateReadOnlyClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateReadOnlyClusterParamsWithHTTPClient(client *http.Client) *UpdateReadOnlyClusterParams {
	return &UpdateReadOnlyClusterParams{
		HTTPClient: client,
	}
}

/* UpdateReadOnlyClusterParams contains all the parameters to send to the API endpoint
   for the update read only cluster operation.

   Typically these are written to a http.Request.
*/
type UpdateReadOnlyClusterParams struct {

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

// WithDefaults hydrates default values in the update read only cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateReadOnlyClusterParams) WithDefaults() *UpdateReadOnlyClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update read only cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateReadOnlyClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update read only cluster params
func (o *UpdateReadOnlyClusterParams) WithTimeout(timeout time.Duration) *UpdateReadOnlyClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update read only cluster params
func (o *UpdateReadOnlyClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update read only cluster params
func (o *UpdateReadOnlyClusterParams) WithContext(ctx context.Context) *UpdateReadOnlyClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update read only cluster params
func (o *UpdateReadOnlyClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update read only cluster params
func (o *UpdateReadOnlyClusterParams) WithHTTPClient(client *http.Client) *UpdateReadOnlyClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update read only cluster params
func (o *UpdateReadOnlyClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the update read only cluster params
func (o *UpdateReadOnlyClusterParams) WithCUUID(cUUID strfmt.UUID) *UpdateReadOnlyClusterParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the update read only cluster params
func (o *UpdateReadOnlyClusterParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithUniUUID adds the uniUUID to the update read only cluster params
func (o *UpdateReadOnlyClusterParams) WithUniUUID(uniUUID strfmt.UUID) *UpdateReadOnlyClusterParams {
	o.SetUniUUID(uniUUID)
	return o
}

// SetUniUUID adds the uniUuid to the update read only cluster params
func (o *UpdateReadOnlyClusterParams) SetUniUUID(uniUUID strfmt.UUID) {
	o.UniUUID = uniUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateReadOnlyClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
