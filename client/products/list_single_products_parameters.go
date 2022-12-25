// Code generated by go-swagger; DO NOT EDIT.

package products

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
	"github.com/go-openapi/swag"
)

// NewListSingleProductsParams creates a new ListSingleProductsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSingleProductsParams() *ListSingleProductsParams {
	return &ListSingleProductsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSingleProductsParamsWithTimeout creates a new ListSingleProductsParams object
// with the ability to set a timeout on a request.
func NewListSingleProductsParamsWithTimeout(timeout time.Duration) *ListSingleProductsParams {
	return &ListSingleProductsParams{
		timeout: timeout,
	}
}

// NewListSingleProductsParamsWithContext creates a new ListSingleProductsParams object
// with the ability to set a context for a request.
func NewListSingleProductsParamsWithContext(ctx context.Context) *ListSingleProductsParams {
	return &ListSingleProductsParams{
		Context: ctx,
	}
}

// NewListSingleProductsParamsWithHTTPClient creates a new ListSingleProductsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSingleProductsParamsWithHTTPClient(client *http.Client) *ListSingleProductsParams {
	return &ListSingleProductsParams{
		HTTPClient: client,
	}
}

/*
ListSingleProductsParams contains all the parameters to send to the API endpoint

	for the list single products operation.

	Typically these are written to a http.Request.
*/
type ListSingleProductsParams struct {

	/* ID.

	   The id of the product for which the operation relates

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list single products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSingleProductsParams) WithDefaults() *ListSingleProductsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list single products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSingleProductsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list single products params
func (o *ListSingleProductsParams) WithTimeout(timeout time.Duration) *ListSingleProductsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list single products params
func (o *ListSingleProductsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list single products params
func (o *ListSingleProductsParams) WithContext(ctx context.Context) *ListSingleProductsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list single products params
func (o *ListSingleProductsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list single products params
func (o *ListSingleProductsParams) WithHTTPClient(client *http.Client) *ListSingleProductsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list single products params
func (o *ListSingleProductsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list single products params
func (o *ListSingleProductsParams) WithID(id int64) *ListSingleProductsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list single products params
func (o *ListSingleProductsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ListSingleProductsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
