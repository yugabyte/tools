// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// NewAcknowledgeByFilterParams creates a new AcknowledgeByFilterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAcknowledgeByFilterParams() *AcknowledgeByFilterParams {
	return &AcknowledgeByFilterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAcknowledgeByFilterParamsWithTimeout creates a new AcknowledgeByFilterParams object
// with the ability to set a timeout on a request.
func NewAcknowledgeByFilterParamsWithTimeout(timeout time.Duration) *AcknowledgeByFilterParams {
	return &AcknowledgeByFilterParams{
		timeout: timeout,
	}
}

// NewAcknowledgeByFilterParamsWithContext creates a new AcknowledgeByFilterParams object
// with the ability to set a context for a request.
func NewAcknowledgeByFilterParamsWithContext(ctx context.Context) *AcknowledgeByFilterParams {
	return &AcknowledgeByFilterParams{
		Context: ctx,
	}
}

// NewAcknowledgeByFilterParamsWithHTTPClient creates a new AcknowledgeByFilterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAcknowledgeByFilterParamsWithHTTPClient(client *http.Client) *AcknowledgeByFilterParams {
	return &AcknowledgeByFilterParams{
		HTTPClient: client,
	}
}

/* AcknowledgeByFilterParams contains all the parameters to send to the API endpoint
   for the acknowledge by filter operation.

   Typically these are written to a http.Request.
*/
type AcknowledgeByFilterParams struct {

	// AcknowledgeAlertsRequest.
	AcknowledgeAlertsRequest *models.AlertAPIFilter

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the acknowledge by filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcknowledgeByFilterParams) WithDefaults() *AcknowledgeByFilterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the acknowledge by filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcknowledgeByFilterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the acknowledge by filter params
func (o *AcknowledgeByFilterParams) WithTimeout(timeout time.Duration) *AcknowledgeByFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the acknowledge by filter params
func (o *AcknowledgeByFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the acknowledge by filter params
func (o *AcknowledgeByFilterParams) WithContext(ctx context.Context) *AcknowledgeByFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the acknowledge by filter params
func (o *AcknowledgeByFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the acknowledge by filter params
func (o *AcknowledgeByFilterParams) WithHTTPClient(client *http.Client) *AcknowledgeByFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the acknowledge by filter params
func (o *AcknowledgeByFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcknowledgeAlertsRequest adds the acknowledgeAlertsRequest to the acknowledge by filter params
func (o *AcknowledgeByFilterParams) WithAcknowledgeAlertsRequest(acknowledgeAlertsRequest *models.AlertAPIFilter) *AcknowledgeByFilterParams {
	o.SetAcknowledgeAlertsRequest(acknowledgeAlertsRequest)
	return o
}

// SetAcknowledgeAlertsRequest adds the acknowledgeAlertsRequest to the acknowledge by filter params
func (o *AcknowledgeByFilterParams) SetAcknowledgeAlertsRequest(acknowledgeAlertsRequest *models.AlertAPIFilter) {
	o.AcknowledgeAlertsRequest = acknowledgeAlertsRequest
}

// WithCUUID adds the cUUID to the acknowledge by filter params
func (o *AcknowledgeByFilterParams) WithCUUID(cUUID strfmt.UUID) *AcknowledgeByFilterParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the acknowledge by filter params
func (o *AcknowledgeByFilterParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WriteToRequest writes these params to a swagger request
func (o *AcknowledgeByFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.AcknowledgeAlertsRequest != nil {
		if err := r.SetBodyParam(o.AcknowledgeAlertsRequest); err != nil {
			return err
		}
	}

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
