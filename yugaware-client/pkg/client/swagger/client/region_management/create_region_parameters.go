// Code generated by go-swagger; DO NOT EDIT.

package region_management

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

// NewCreateRegionParams creates a new CreateRegionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRegionParams() *CreateRegionParams {
	return &CreateRegionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRegionParamsWithTimeout creates a new CreateRegionParams object
// with the ability to set a timeout on a request.
func NewCreateRegionParamsWithTimeout(timeout time.Duration) *CreateRegionParams {
	return &CreateRegionParams{
		timeout: timeout,
	}
}

// NewCreateRegionParamsWithContext creates a new CreateRegionParams object
// with the ability to set a context for a request.
func NewCreateRegionParamsWithContext(ctx context.Context) *CreateRegionParams {
	return &CreateRegionParams{
		Context: ctx,
	}
}

// NewCreateRegionParamsWithHTTPClient creates a new CreateRegionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRegionParamsWithHTTPClient(client *http.Client) *CreateRegionParams {
	return &CreateRegionParams{
		HTTPClient: client,
	}
}

/* CreateRegionParams contains all the parameters to send to the API endpoint
   for the create region operation.

   Typically these are written to a http.Request.
*/
type CreateRegionParams struct {

	// CUUID.
	//
	// Format: uuid
	CUUID strfmt.UUID

	// PUUID.
	//
	// Format: uuid
	PUUID strfmt.UUID

	/* Region.

	   region form data for new region to be created
	*/
	Region *models.RegionFormData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create region params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRegionParams) WithDefaults() *CreateRegionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create region params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRegionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create region params
func (o *CreateRegionParams) WithTimeout(timeout time.Duration) *CreateRegionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create region params
func (o *CreateRegionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create region params
func (o *CreateRegionParams) WithContext(ctx context.Context) *CreateRegionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create region params
func (o *CreateRegionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create region params
func (o *CreateRegionParams) WithHTTPClient(client *http.Client) *CreateRegionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create region params
func (o *CreateRegionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCUUID adds the cUUID to the create region params
func (o *CreateRegionParams) WithCUUID(cUUID strfmt.UUID) *CreateRegionParams {
	o.SetCUUID(cUUID)
	return o
}

// SetCUUID adds the cUuid to the create region params
func (o *CreateRegionParams) SetCUUID(cUUID strfmt.UUID) {
	o.CUUID = cUUID
}

// WithPUUID adds the pUUID to the create region params
func (o *CreateRegionParams) WithPUUID(pUUID strfmt.UUID) *CreateRegionParams {
	o.SetPUUID(pUUID)
	return o
}

// SetPUUID adds the pUuid to the create region params
func (o *CreateRegionParams) SetPUUID(pUUID strfmt.UUID) {
	o.PUUID = pUUID
}

// WithRegion adds the region to the create region params
func (o *CreateRegionParams) WithRegion(region *models.RegionFormData) *CreateRegionParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the create region params
func (o *CreateRegionParams) SetRegion(region *models.RegionFormData) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRegionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cUUID
	if err := r.SetPathParam("cUUID", o.CUUID.String()); err != nil {
		return err
	}

	// path param pUUID
	if err := r.SetPathParam("pUUID", o.PUUID.String()); err != nil {
		return err
	}
	if o.Region != nil {
		if err := r.SetBodyParam(o.Region); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
