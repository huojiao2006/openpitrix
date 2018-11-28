// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeAppVersionAuditsParams creates a new DescribeAppVersionAuditsParams object
// with the default values initialized.
func NewDescribeAppVersionAuditsParams() *DescribeAppVersionAuditsParams {
	var ()
	return &DescribeAppVersionAuditsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeAppVersionAuditsParamsWithTimeout creates a new DescribeAppVersionAuditsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeAppVersionAuditsParamsWithTimeout(timeout time.Duration) *DescribeAppVersionAuditsParams {
	var ()
	return &DescribeAppVersionAuditsParams{

		timeout: timeout,
	}
}

// NewDescribeAppVersionAuditsParamsWithContext creates a new DescribeAppVersionAuditsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeAppVersionAuditsParamsWithContext(ctx context.Context) *DescribeAppVersionAuditsParams {
	var ()
	return &DescribeAppVersionAuditsParams{

		Context: ctx,
	}
}

// NewDescribeAppVersionAuditsParamsWithHTTPClient creates a new DescribeAppVersionAuditsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeAppVersionAuditsParamsWithHTTPClient(client *http.Client) *DescribeAppVersionAuditsParams {
	var ()
	return &DescribeAppVersionAuditsParams{
		HTTPClient: client,
	}
}

/*DescribeAppVersionAuditsParams contains all the parameters to send to the API endpoint
for the describe app version audits operation typically these are written to a http.Request
*/
type DescribeAppVersionAuditsParams struct {

	/*AppID*/
	AppID string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*Operator*/
	Operator []string
	/*Reverse*/
	Reverse *bool
	/*Role*/
	Role []string
	/*SearchWord*/
	SearchWord *string
	/*SortKey*/
	SortKey *string
	/*Status*/
	Status []string
	/*VersionID*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) WithTimeout(timeout time.Duration) *DescribeAppVersionAuditsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) WithContext(ctx context.Context) *DescribeAppVersionAuditsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) WithHTTPClient(client *http.Client) *DescribeAppVersionAuditsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) WithAppID(appID string) *DescribeAppVersionAuditsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithLimit adds the limit to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) WithLimit(limit *int64) *DescribeAppVersionAuditsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) WithOffset(offset *int64) *DescribeAppVersionAuditsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOperator adds the operator to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) WithOperator(operator []string) *DescribeAppVersionAuditsParams {
	o.SetOperator(operator)
	return o
}

// SetOperator adds the operator to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) SetOperator(operator []string) {
	o.Operator = operator
}

// WithReverse adds the reverse to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) WithReverse(reverse *bool) *DescribeAppVersionAuditsParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithRole adds the role to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) WithRole(role []string) *DescribeAppVersionAuditsParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) SetRole(role []string) {
	o.Role = role
}

// WithSearchWord adds the searchWord to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) WithSearchWord(searchWord *string) *DescribeAppVersionAuditsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) WithSortKey(sortKey *string) *DescribeAppVersionAuditsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) WithStatus(status []string) *DescribeAppVersionAuditsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) SetStatus(status []string) {
	o.Status = status
}

// WithVersionID adds the versionID to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) WithVersionID(versionID string) *DescribeAppVersionAuditsParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the describe app version audits params
func (o *DescribeAppVersionAuditsParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeAppVersionAuditsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOperator := o.Operator

	joinedOperator := swag.JoinByFormat(valuesOperator, "multi")
	// query array param operator
	if err := r.SetQueryParam("operator", joinedOperator...); err != nil {
		return err
	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool
		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {
			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}

	}

	valuesRole := o.Role

	joinedRole := swag.JoinByFormat(valuesRole, "multi")
	// query array param role
	if err := r.SetQueryParam("role", joinedRole...); err != nil {
		return err
	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string
		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {
			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	// path param version_id
	if err := r.SetPathParam("version_id", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
