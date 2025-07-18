// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package model

import (
	"github.com/google/uuid"
	"github.com/onexstack/onexstack/pkg/authn"
	"gorm.io/gorm"

	known "github.com/onexstack/onex/internal/pkg/known/usercenter"
	"github.com/onexstack/onex/internal/pkg/rid"
)

// BeforeCreate runs before creating a SecretM database record and initializes various fields.
func (m *SecretM) BeforeCreate(tx *gorm.DB) (err error) {
	// Supports custom SecretKey, but users need to ensure the uniqueness of the SecretKey themselves.
	// onex-cacheserver will use this feature to set secret.
	if m.SecretID == "" {
		// Generate a new UUID for SecretKey.
		m.SecretID = uuid.New().String()
	}

	// Generate a new UUID for SecretID.
	m.SecretKey = uuid.New().String()

	// Set the default status for the secret as normal.
	m.Status = known.SecretStatusNormal

	return nil
}

// AfterCreate runs after creating a UserM database record and updates the UserID field.
func (m *UserM) AfterCreate(tx *gorm.DB) (err error) {
	m.UserID = rid.User.New(uint64(m.ID)) // Generate and set a new user ID.

	return tx.Save(m).Error // Save the updated user record to the database.
}

// BeforeCreate runs before creating a UserM database record and initializes various fields.
func (m *UserM) BeforeCreate(tx *gorm.DB) (err error) {
	m.Password, err = authn.Encrypt(m.Password) // Encrypt the user password.
	if err != nil {
		return err // Return error if there's a problem with encryption.
	}

	m.Status = known.UserStatusActived // Set the default status for the user as active.

	return nil
}
