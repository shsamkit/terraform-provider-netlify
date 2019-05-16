// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// User user
// swagger:model user
type User struct {

	// affiliate id
	AffiliateID string `json:"affiliate_id,omitempty"`

	// avatar url
	AvatarURL string `json:"avatar_url,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last login
	LastLogin string `json:"last_login,omitempty"`

	// login providers
	LoginProviders []string `json:"login_providers"`

	// onboarding progress
	OnboardingProgress *UserOnboardingProgress `json:"onboarding_progress,omitempty"`

	// site count
	SiteCount int64 `json:"site_count,omitempty"`

	// support priority
	SupportPriority int64 `json:"support_priority,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOnboardingProgress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateOnboardingProgress(formats strfmt.Registry) error {

	if swag.IsZero(m.OnboardingProgress) { // not required
		return nil
	}

	if m.OnboardingProgress != nil {
		if err := m.OnboardingProgress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("onboarding_progress")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
