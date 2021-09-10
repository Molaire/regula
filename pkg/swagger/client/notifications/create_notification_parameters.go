// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

	"github.com/fugue/regula/pkg/swagger/models"
)

// NewCreateNotificationParams creates a new CreateNotificationParams object
// with the default values initialized.
func NewCreateNotificationParams() *CreateNotificationParams {
	var ()
	return &CreateNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNotificationParamsWithTimeout creates a new CreateNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateNotificationParamsWithTimeout(timeout time.Duration) *CreateNotificationParams {
	var ()
	return &CreateNotificationParams{

		timeout: timeout,
	}
}

// NewCreateNotificationParamsWithContext creates a new CreateNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateNotificationParamsWithContext(ctx context.Context) *CreateNotificationParams {
	var ()
	return &CreateNotificationParams{

		Context: ctx,
	}
}

// NewCreateNotificationParamsWithHTTPClient creates a new CreateNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateNotificationParamsWithHTTPClient(client *http.Client) *CreateNotificationParams {
	var ()
	return &CreateNotificationParams{
		HTTPClient: client,
	}
}

/*CreateNotificationParams contains all the parameters to send to the API endpoint
for the create notification operation typically these are written to a http.Request
*/
type CreateNotificationParams struct {

	/*Notification
	  Configuration options for the new notification.

	*/
	Notification *models.CreateNotificationInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create notification params
func (o *CreateNotificationParams) WithTimeout(timeout time.Duration) *CreateNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create notification params
func (o *CreateNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create notification params
func (o *CreateNotificationParams) WithContext(ctx context.Context) *CreateNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create notification params
func (o *CreateNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create notification params
func (o *CreateNotificationParams) WithHTTPClient(client *http.Client) *CreateNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create notification params
func (o *CreateNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNotification adds the notification to the create notification params
func (o *CreateNotificationParams) WithNotification(notification *models.CreateNotificationInput) *CreateNotificationParams {
	o.SetNotification(notification)
	return o
}

// SetNotification adds the notification to the create notification params
func (o *CreateNotificationParams) SetNotification(notification *models.CreateNotificationInput) {
	o.Notification = notification
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Notification != nil {
		if err := r.SetBodyParam(o.Notification); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}