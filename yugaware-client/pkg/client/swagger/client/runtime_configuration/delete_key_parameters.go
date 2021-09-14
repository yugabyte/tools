// Code generated by go-swagger; DO NOT EDIT.

package runtime_configuration

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

// NewDeleteKeyParams creates a new DeleteKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteKeyParams() *DeleteKeyParams {
	return &DeleteKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKeyParamsWithTimeout creates a new DeleteKeyParams object
// with the ability to set a timeout on a request.
func NewDeleteKeyParamsWithTimeout(timeout time.Duration) *DeleteKeyParams {
	return &DeleteKeyParams{
		timeout: timeout,
	}
}

// NewDeleteKeyParamsWithContext creates a new DeleteKeyParams object
// with the ability to set a context for a request.
func NewDeleteKeyParamsWithContext(ctx context.Context) *DeleteKeyParams {
	return &DeleteKeyParams{
		Context: ctx,
	}
}

// NewDeleteKeyParamsWithHTTPClient creates a new DeleteKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteKeyParamsWithHTTPClient(client *http.Client) *DeleteKeyParams {
	return &DeleteKeyParams{
		HTTPClient: client,
	}
}

/* DeleteKeyParams contains all the parameters to send to the API endpoint
   for the delete key operation.

   Typically these are written to a http.Request.
*/
type DeleteKeyParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	// Key.
	Key string

	// Scope.
	//
	// Format: uuid
	Scope strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKeyParams) WithDefaults() *DeleteKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete key params
func (o *DeleteKeyParams) WithTimeout(timeout time.Duration) *DeleteKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete key params
func (o *DeleteKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete key params
func (o *DeleteKeyParams) WithContext(ctx context.Context) *DeleteKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete key params
func (o *DeleteKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete key params
func (o *DeleteKeyParams) WithHTTPClient(client *http.Client) *DeleteKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete key params
func (o *DeleteKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the delete key params
func (o *DeleteKeyParams) WithCUUID(cUUID strfmt.UUID) *DeleteKeyParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the delete key params
func (o *DeleteKeyParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithKey adds the key to the delete key params
func (o *DeleteKeyParams) WithKey(key string) *DeleteKeyParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the delete key params
func (o *DeleteKeyParams) SetKey(key string) {
	o.Key = key
}

// WithScope adds the scope to the delete key params
func (o *DeleteKeyParams) WithScope(scope strfmt.UUID) *DeleteKeyParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the delete key params
func (o *DeleteKeyParams) SetScope(scope strfmt.UUID) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	// path param scope
	if err := r.SetPathParam("scope", o.Scope.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
