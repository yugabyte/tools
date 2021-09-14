// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlertChannelEmailParams alert channel email params
//
// swagger:model AlertChannelEmailParams
type AlertChannelEmailParams struct {
	textTemplateField *string

	titleTemplateField *string

	// default recipients
	// Required: true
	DefaultRecipients *bool `json:"defaultRecipients"`

	// default Smtp settings
	// Required: true
	DefaultSMTPSettings *bool `json:"defaultSmtpSettings"`

	// recipients
	// Required: true
	Recipients []string `json:"recipients"`

	// smtp data
	// Required: true
	SMTPData *SMTPData `json:"smtpData"`
}

// TextTemplate gets the text template of this subtype
func (m *AlertChannelEmailParams) TextTemplate() *string {
	return m.textTemplateField
}

// SetTextTemplate sets the text template of this subtype
func (m *AlertChannelEmailParams) SetTextTemplate(val *string) {
	m.textTemplateField = val
}

// TitleTemplate gets the title template of this subtype
func (m *AlertChannelEmailParams) TitleTemplate() *string {
	return m.titleTemplateField
}

// SetTitleTemplate sets the title template of this subtype
func (m *AlertChannelEmailParams) SetTitleTemplate(val *string) {
	m.titleTemplateField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AlertChannelEmailParams) UnmarshalJSON(raw []byte) error {
	var data struct {

		// default recipients
		// Required: true
		DefaultRecipients *bool `json:"defaultRecipients"`

		// default Smtp settings
		// Required: true
		DefaultSMTPSettings *bool `json:"defaultSmtpSettings"`

		// recipients
		// Required: true
		Recipients []string `json:"recipients"`

		// smtp data
		// Required: true
		SMTPData *SMTPData `json:"smtpData"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		TextTemplate *string `json:"textTemplate"`

		TitleTemplate *string `json:"titleTemplate"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result AlertChannelEmailParams

	result.textTemplateField = base.TextTemplate

	result.titleTemplateField = base.TitleTemplate

	result.DefaultRecipients = data.DefaultRecipients
	result.DefaultSMTPSettings = data.DefaultSMTPSettings
	result.Recipients = data.Recipients
	result.SMTPData = data.SMTPData

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AlertChannelEmailParams) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// default recipients
		// Required: true
		DefaultRecipients *bool `json:"defaultRecipients"`

		// default Smtp settings
		// Required: true
		DefaultSMTPSettings *bool `json:"defaultSmtpSettings"`

		// recipients
		// Required: true
		Recipients []string `json:"recipients"`

		// smtp data
		// Required: true
		SMTPData *SMTPData `json:"smtpData"`
	}{

		DefaultRecipients: m.DefaultRecipients,

		DefaultSMTPSettings: m.DefaultSMTPSettings,

		Recipients: m.Recipients,

		SMTPData: m.SMTPData,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		TextTemplate *string `json:"textTemplate"`

		TitleTemplate *string `json:"titleTemplate"`
	}{

		TextTemplate: m.TextTemplate(),

		TitleTemplate: m.TitleTemplate(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this alert channel email params
func (m *AlertChannelEmailParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTextTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitleTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultSMTPSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSMTPData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertChannelEmailParams) validateTextTemplate(formats strfmt.Registry) error {

	if err := validate.Required("textTemplate", "body", m.TextTemplate()); err != nil {
		return err
	}

	return nil
}

func (m *AlertChannelEmailParams) validateTitleTemplate(formats strfmt.Registry) error {

	if err := validate.Required("titleTemplate", "body", m.TitleTemplate()); err != nil {
		return err
	}

	return nil
}

func (m *AlertChannelEmailParams) validateDefaultRecipients(formats strfmt.Registry) error {

	if err := validate.Required("defaultRecipients", "body", m.DefaultRecipients); err != nil {
		return err
	}

	return nil
}

func (m *AlertChannelEmailParams) validateDefaultSMTPSettings(formats strfmt.Registry) error {

	if err := validate.Required("defaultSmtpSettings", "body", m.DefaultSMTPSettings); err != nil {
		return err
	}

	return nil
}

func (m *AlertChannelEmailParams) validateRecipients(formats strfmt.Registry) error {

	if err := validate.Required("recipients", "body", m.Recipients); err != nil {
		return err
	}

	return nil
}

func (m *AlertChannelEmailParams) validateSMTPData(formats strfmt.Registry) error {

	if err := validate.Required("smtpData", "body", m.SMTPData); err != nil {
		return err
	}

	if m.SMTPData != nil {
		if err := m.SMTPData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smtpData")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this alert channel email params based on the context it is used
func (m *AlertChannelEmailParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSMTPData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertChannelEmailParams) contextValidateSMTPData(ctx context.Context, formats strfmt.Registry) error {

	if m.SMTPData != nil {
		if err := m.SMTPData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smtpData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertChannelEmailParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertChannelEmailParams) UnmarshalBinary(b []byte) error {
	var res AlertChannelEmailParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
