// Code generated by go-swagger; DO NOT EDIT.

package measurements

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

// NewInstallationMeasurementsUsingGETParams creates a new InstallationMeasurementsUsingGETParams object
// with the default values initialized.
func NewInstallationMeasurementsUsingGETParams() *InstallationMeasurementsUsingGETParams {
	var (
		indexTypeDefault      = string("AIRLY_CAQI")
		installationIDDefault = int32(204)
	)
	return &InstallationMeasurementsUsingGETParams{
		IndexType:      &indexTypeDefault,
		InstallationID: installationIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewInstallationMeasurementsUsingGETParamsWithTimeout creates a new InstallationMeasurementsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInstallationMeasurementsUsingGETParamsWithTimeout(timeout time.Duration) *InstallationMeasurementsUsingGETParams {
	var (
		indexTypeDefault      = string("AIRLY_CAQI")
		installationIDDefault = int32(204)
	)
	return &InstallationMeasurementsUsingGETParams{
		IndexType:      &indexTypeDefault,
		InstallationID: installationIDDefault,

		timeout: timeout,
	}
}

// NewInstallationMeasurementsUsingGETParamsWithContext creates a new InstallationMeasurementsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewInstallationMeasurementsUsingGETParamsWithContext(ctx context.Context) *InstallationMeasurementsUsingGETParams {
	var (
		indexTypeDefault      = string("AIRLY_CAQI")
		installationIdDefault = int32(204)
	)
	return &InstallationMeasurementsUsingGETParams{
		IndexType:      &indexTypeDefault,
		InstallationID: installationIdDefault,

		Context: ctx,
	}
}

// NewInstallationMeasurementsUsingGETParamsWithHTTPClient creates a new InstallationMeasurementsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInstallationMeasurementsUsingGETParamsWithHTTPClient(client *http.Client) *InstallationMeasurementsUsingGETParams {
	var (
		indexTypeDefault      = string("AIRLY_CAQI")
		installationIdDefault = int32(204)
	)
	return &InstallationMeasurementsUsingGETParams{
		IndexType:      &indexTypeDefault,
		InstallationID: installationIdDefault,
		HTTPClient:     client,
	}
}

/*InstallationMeasurementsUsingGETParams contains all the parameters to send to the API endpoint
for the installation measurements using g e t operation typically these are written to a http.Request
*/
type InstallationMeasurementsUsingGETParams struct {

	/*IndexType
	  Select index which should be returned in response

	*/
	IndexType *string
	/*InstallationID
	  AirlyInstallation ID

	*/
	InstallationID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the installation measurements using g e t params
func (o *InstallationMeasurementsUsingGETParams) WithTimeout(timeout time.Duration) *InstallationMeasurementsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the installation measurements using g e t params
func (o *InstallationMeasurementsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the installation measurements using g e t params
func (o *InstallationMeasurementsUsingGETParams) WithContext(ctx context.Context) *InstallationMeasurementsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the installation measurements using g e t params
func (o *InstallationMeasurementsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the installation measurements using g e t params
func (o *InstallationMeasurementsUsingGETParams) WithHTTPClient(client *http.Client) *InstallationMeasurementsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the installation measurements using g e t params
func (o *InstallationMeasurementsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndexType adds the indexType to the installation measurements using g e t params
func (o *InstallationMeasurementsUsingGETParams) WithIndexType(indexType *string) *InstallationMeasurementsUsingGETParams {
	o.SetIndexType(indexType)
	return o
}

// SetIndexType adds the indexType to the installation measurements using g e t params
func (o *InstallationMeasurementsUsingGETParams) SetIndexType(indexType *string) {
	o.IndexType = indexType
}

// WithInstallationID adds the installationID to the installation measurements using g e t params
func (o *InstallationMeasurementsUsingGETParams) WithInstallationID(installationID int32) *InstallationMeasurementsUsingGETParams {
	o.SetInstallationID(installationID)
	return o
}

// SetInstallationID adds the installationId to the installation measurements using g e t params
func (o *InstallationMeasurementsUsingGETParams) SetInstallationID(installationID int32) {
	o.InstallationID = installationID
}

// WriteToRequest writes these params to a swagger request
func (o *InstallationMeasurementsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IndexType != nil {

		// query param indexType
		var qrIndexType string
		if o.IndexType != nil {
			qrIndexType = *o.IndexType
		}
		qIndexType := qrIndexType
		if qIndexType != "" {
			if err := r.SetQueryParam("indexType", qIndexType); err != nil {
				return err
			}
		}

	}

	// query param installationId
	qrInstallationID := o.InstallationID
	qInstallationID := swag.FormatInt32(qrInstallationID)
	if qInstallationID != "" {
		if err := r.SetQueryParam("installationId", qInstallationID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
