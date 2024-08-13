// Code generated by go-swagger; DO NOT EDIT.

package customers

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

	"github.com/cloudentity/go-swagger/examples/authentication/models"
)

// NewGetIDParams creates a new GetIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIDParams() *GetIDParams {
	return &GetIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIDParamsWithTimeout creates a new GetIDParams object
// with the ability to set a timeout on a request.
func NewGetIDParamsWithTimeout(timeout time.Duration) *GetIDParams {
	return &GetIDParams{
		timeout: timeout,
	}
}

// NewGetIDParamsWithContext creates a new GetIDParams object
// with the ability to set a context for a request.
func NewGetIDParamsWithContext(ctx context.Context) *GetIDParams {
	return &GetIDParams{
		Context: ctx,
	}
}

// NewGetIDParamsWithHTTPClient creates a new GetIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIDParamsWithHTTPClient(client *http.Client) *GetIDParams {
	return &GetIDParams{
		HTTPClient: client,
	}
}

/*
GetIDParams contains all the parameters to send to the API endpoint

	for the get Id operation.

	Typically these are written to a http.Request.
*/
type GetIDParams struct {

	// Info.
	Info *models.SocialID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIDParams) WithDefaults() *GetIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Id params
func (o *GetIDParams) WithTimeout(timeout time.Duration) *GetIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Id params
func (o *GetIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Id params
func (o *GetIDParams) WithContext(ctx context.Context) *GetIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Id params
func (o *GetIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Id params
func (o *GetIDParams) WithHTTPClient(client *http.Client) *GetIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Id params
func (o *GetIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the get Id params
func (o *GetIDParams) WithInfo(info *models.SocialID) *GetIDParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the get Id params
func (o *GetIDParams) SetInfo(info *models.SocialID) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *GetIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
