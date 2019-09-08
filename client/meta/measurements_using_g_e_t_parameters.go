// Code generated by go-swagger; DO NOT EDIT.

package meta

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewMeasurementsUsingGETParams creates a new MeasurementsUsingGETParams object
// with the default values initialized.
func NewMeasurementsUsingGETParams() *MeasurementsUsingGETParams {

	return &MeasurementsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMeasurementsUsingGETParamsWithTimeout creates a new MeasurementsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMeasurementsUsingGETParamsWithTimeout(timeout time.Duration) *MeasurementsUsingGETParams {

	return &MeasurementsUsingGETParams{

		timeout: timeout,
	}
}

// NewMeasurementsUsingGETParamsWithContext creates a new MeasurementsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewMeasurementsUsingGETParamsWithContext(ctx context.Context) *MeasurementsUsingGETParams {

	return &MeasurementsUsingGETParams{

		Context: ctx,
	}
}

// NewMeasurementsUsingGETParamsWithHTTPClient creates a new MeasurementsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMeasurementsUsingGETParamsWithHTTPClient(client *http.Client) *MeasurementsUsingGETParams {

	return &MeasurementsUsingGETParams{
		HTTPClient: client,
	}
}

/*MeasurementsUsingGETParams contains all the parameters to send to the API endpoint
for the measurements using g e t operation typically these are written to a http.Request
*/
type MeasurementsUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the measurements using g e t params
func (o *MeasurementsUsingGETParams) WithTimeout(timeout time.Duration) *MeasurementsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the measurements using g e t params
func (o *MeasurementsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the measurements using g e t params
func (o *MeasurementsUsingGETParams) WithContext(ctx context.Context) *MeasurementsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the measurements using g e t params
func (o *MeasurementsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the measurements using g e t params
func (o *MeasurementsUsingGETParams) WithHTTPClient(client *http.Client) *MeasurementsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the measurements using g e t params
func (o *MeasurementsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MeasurementsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}