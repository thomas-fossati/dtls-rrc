// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

import (
	"encoding/binary"
)

const (
	useRRCHeaderSize = 4
)

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
type UseRRC struct {
	Supported bool
}

// // TypeValue returns the extension TypeValue
func (r UseRRC) TypeValue() TypeValue {
	return UseRRCTypeValue
}

// Marshal encodes the extension
func (r *UseRRC) Marshal() ([]byte, error) {
	if !r.Supported {
		return []byte{}, nil
	}

	out := make([]byte, useRRCHeaderSize)

	binary.BigEndian.PutUint16(out, uint16(r.TypeValue()))
	binary.BigEndian.PutUint16(out[2:], uint16(0)) // length

	return out, nil
}

// Unmarshal populates the extension from encoded data
func (r *UseRRC) Unmarshal(data []byte) error {
	if len(data) < useRRCHeaderSize {
		return errBufferTooSmall
	} else if TypeValue(binary.BigEndian.Uint16(data)) != r.TypeValue() {
		return errInvalidExtensionType
	}

	r.Supported = true

	return nil
}
