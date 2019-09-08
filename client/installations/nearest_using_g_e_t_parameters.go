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

// NewNearestUsingGETParams creates a new NearestUsingGETParams object
// with the default values initialized.
func NewNearestUsingGETParams() *NearestUsingGETParams {
	var (
		latDefault           = float64(50.062006)
		lngDefault           = float64(19.940984)
		maxDistanceKMDefault = float64(3)
		maxResultsDefault    = int32(1)
	)
	return &NearestUsingGETParams{
		Lat:           latDefault,
		Lng:           lngDefault,
		MaxDistanceKM: &maxDistanceKMDefault,
		MaxResults:    &maxResultsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewNearestUsingGETParamsWithTimeout creates a new NearestUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNearestUsingGETParamsWithTimeout(timeout time.Duration) *NearestUsingGETParams {
	var (
		latDefault           = float64(50.062006)
		lngDefault           = float64(19.940984)
		maxDistanceKMDefault = float64(3)
		maxResultsDefault    = int32(1)
	)
	return &NearestUsingGETParams{
		Lat:           latDefault,
		Lng:           lngDefault,
		MaxDistanceKM: &maxDistanceKMDefault,
		MaxResults:    &maxResultsDefault,

		timeout: timeout,
	}
}

// NewNearestUsingGETParamsWithContext creates a new NearestUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewNearestUsingGETParamsWithContext(ctx context.Context) *NearestUsingGETParams {
	var (
		latDefault           = float64(50.062006)
		lngDefault           = float64(19.940984)
		maxDistanceKMDefault = float64(3)
		maxResultsDefault    = int32(1)
	)
	return &NearestUsingGETParams{
		Lat:           latDefault,
		Lng:           lngDefault,
		MaxDistanceKM: &maxDistanceKMDefault,
		MaxResults:    &maxResultsDefault,

		Context: ctx,
	}
}

// NewNearestUsingGETParamsWithHTTPClient creates a new NearestUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNearestUsingGETParamsWithHTTPClient(client *http.Client) *NearestUsingGETParams {
	var (
		latDefault           = float64(50.062006)
		lngDefault           = float64(19.940984)
		maxDistanceKMDefault = float64(3)
		maxResultsDefault    = int32(1)
	)
	return &NearestUsingGETParams{
		Lat:           latDefault,
		Lng:           lngDefault,
		MaxDistanceKM: &maxDistanceKMDefault,
		MaxResults:    &maxResultsDefault,
		HTTPClient:    client,
	}
}

/*NearestUsingGETParams contains all the parameters to send to the API endpoint
for the nearest using g e t operation typically these are written to a http.Request
*/
type NearestUsingGETParams struct {

	/*Lat
	  Latitude

	*/
	Lat float64
	/*Lng
	  Longitude

	*/
	Lng float64
	/*MaxDistanceKM
	  Max distance in km. Negative values mean 'infinity' (no distance limit). Optional parameter, if omitted defaults to 3km.

	*/
	MaxDistanceKM *float64
	/*MaxResults
	  Max results to return. Negative values mean 'infinity' (no results limit). Optional parameter, if omitted defaults to 1 item.

	*/
	MaxResults *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the nearest using g e t params
func (o *NearestUsingGETParams) WithTimeout(timeout time.Duration) *NearestUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nearest using g e t params
func (o *NearestUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nearest using g e t params
func (o *NearestUsingGETParams) WithContext(ctx context.Context) *NearestUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nearest using g e t params
func (o *NearestUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nearest using g e t params
func (o *NearestUsingGETParams) WithHTTPClient(client *http.Client) *NearestUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nearest using g e t params
func (o *NearestUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLat adds the lat to the nearest using g e t params
func (o *NearestUsingGETParams) WithLat(lat float64) *NearestUsingGETParams {
	o.SetLat(lat)
	return o
}

// SetLat adds the lat to the nearest using g e t params
func (o *NearestUsingGETParams) SetLat(lat float64) {
	o.Lat = lat
}

// WithLng adds the lng to the nearest using g e t params
func (o *NearestUsingGETParams) WithLng(lng float64) *NearestUsingGETParams {
	o.SetLng(lng)
	return o
}

// SetLng adds the lng to the nearest using g e t params
func (o *NearestUsingGETParams) SetLng(lng float64) {
	o.Lng = lng
}

// WithMaxDistanceKM adds the maxDistanceKM to the nearest using g e t params
func (o *NearestUsingGETParams) WithMaxDistanceKM(maxDistanceKM *float64) *NearestUsingGETParams {
	o.SetMaxDistanceKM(maxDistanceKM)
	return o
}

// SetMaxDistanceKM adds the maxDistanceKM to the nearest using g e t params
func (o *NearestUsingGETParams) SetMaxDistanceKM(maxDistanceKM *float64) {
	o.MaxDistanceKM = maxDistanceKM
}

// WithMaxResults adds the maxResults to the nearest using g e t params
func (o *NearestUsingGETParams) WithMaxResults(maxResults *int32) *NearestUsingGETParams {
	o.SetMaxResults(maxResults)
	return o
}

// SetMaxResults adds the maxResults to the nearest using g e t params
func (o *NearestUsingGETParams) SetMaxResults(maxResults *int32) {
	o.MaxResults = maxResults
}

// WriteToRequest writes these params to a swagger request
func (o *NearestUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param lat
	qrLat := o.Lat
	qLat := swag.FormatFloat64(qrLat)
	if qLat != "" {
		if err := r.SetQueryParam("lat", qLat); err != nil {
			return err
		}
	}

	// query param lng
	qrLng := o.Lng
	qLng := swag.FormatFloat64(qrLng)
	if qLng != "" {
		if err := r.SetQueryParam("lng", qLng); err != nil {
			return err
		}
	}

	if o.MaxDistanceKM != nil {

		// query param maxDistanceKM
		var qrMaxDistanceKM float64
		if o.MaxDistanceKM != nil {
			qrMaxDistanceKM = *o.MaxDistanceKM
		}
		qMaxDistanceKM := swag.FormatFloat64(qrMaxDistanceKM)
		if qMaxDistanceKM != "" {
			if err := r.SetQueryParam("maxDistanceKM", qMaxDistanceKM); err != nil {
				return err
			}
		}

	}

	if o.MaxResults != nil {

		// query param maxResults
		var qrMaxResults int32
		if o.MaxResults != nil {
			qrMaxResults = *o.MaxResults
		}
		qMaxResults := swag.FormatInt32(qrMaxResults)
		if qMaxResults != "" {
			if err := r.SetQueryParam("maxResults", qMaxResults); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}