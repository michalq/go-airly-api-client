// Code generated by go-swagger; DO NOT EDIT.

package installations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewByIDUsingGETParams creates a new ByIDUsingGETParams object
// with the default values initialized.
func NewByIDUsingGETParams() *ByIDUsingGETParams {
	var (
		installationIDDefault = int32(204)
	)
	return &ByIDUsingGETParams{
		InstallationID: installationIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewByIDUsingGETParamsWithTimeout creates a new ByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewByIDUsingGETParamsWithTimeout(timeout time.Duration) *ByIDUsingGETParams {
	var (
		installationIDDefault = int32(204)
	)
	return &ByIDUsingGETParams{
		InstallationID: installationIDDefault,

		timeout: timeout,
	}
}

// NewByIDUsingGETParamsWithContext creates a new ByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewByIDUsingGETParamsWithContext(ctx context.Context) *ByIDUsingGETParams {
	var (
		installationIdDefault = int32(204)
	)
	return &ByIDUsingGETParams{
		InstallationID: installationIdDefault,

		Context: ctx,
	}
}

// NewByIDUsingGETParamsWithHTTPClient creates a new ByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewByIDUsingGETParamsWithHTTPClient(client *http.Client) *ByIDUsingGETParams {
	var (
		installationIdDefault = int32(204)
	)
	return &ByIDUsingGETParams{
		InstallationID: installationIdDefault,
		HTTPClient:     client,
	}
}

/*ByIDUsingGETParams contains all the parameters to send to the API endpoint
for the by Id using g e t operation typically these are written to a http.Request
*/
type ByIDUsingGETParams struct {

	/*InstallationID
	  Installation ID

	*/
	InstallationID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the by Id using g e t params
func (o *ByIDUsingGETParams) WithTimeout(timeout time.Duration) *ByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the by Id using g e t params
func (o *ByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the by Id using g e t params
func (o *ByIDUsingGETParams) WithContext(ctx context.Context) *ByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the by Id using g e t params
func (o *ByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the by Id using g e t params
func (o *ByIDUsingGETParams) WithHTTPClient(client *http.Client) *ByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the by Id using g e t params
func (o *ByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallationID adds the installationID to the by Id using g e t params
func (o *ByIDUsingGETParams) WithInstallationID(installationID int32) *ByIDUsingGETParams {
	o.SetInstallationID(installationID)
	return o
}

// SetInstallationID adds the installationId to the by Id using g e t params
func (o *ByIDUsingGETParams) SetInstallationID(installationID int32) {
	o.InstallationID = installationID
}

// WriteToRequest writes these params to a swagger request
func (o *ByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param installationId
	if err := r.SetPathParam("installationId", swag.FormatInt32(o.InstallationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
