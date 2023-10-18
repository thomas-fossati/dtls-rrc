// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

import "errors"

// RRC is a DTLS extension that provides a minimal, but extensible, path-layer
// sub-protocol. The main use case for RRC is as a companion to Connection ID
// (CID), to check that path migrations are genuine and that it is safe to
// update CID-related transport bindings.
//
// Code points for experimentation
// rrc extension (TBD1): 61
// return_routability_check content type (TBD2): 27
//
// https://datatracker.ietf.org/doc/draft-ietf-tls-dtls-rrc
type RRC struct{}

// // TypeValue returns the extension TypeValue
func (r RRC) TypeValue() TypeValue {
	return ConnectionIDTypeValue
}

// Marshal encodes the extension
func (r *RRC) Marshal() ([]byte, error) {
	return nil, errors.New("TODO")
}

// Unmarshal populates the extension from encoded data
func (r *RRC) Unmarshal(data []byte) error {
	return errors.New("TODO")
}
