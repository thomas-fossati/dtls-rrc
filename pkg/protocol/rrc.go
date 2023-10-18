// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package protocol

import "errors"

// RRC messages
// https://www.ietf.org/archive/id/draft-ietf-tls-dtls-rrc-10.html#section-4
type RRC struct{}

// ContentType returns the RRC ContentType
func (a RRC) ContentType() ContentType {
	return ContentTypeRRC
}

// Marshal encodes the ApplicationData to binary
func (a *RRC) Marshal() ([]byte, error) {
	return nil, errors.New("TODO")
}

// Unmarshal populates the ApplicationData from binary
func (a *RRC) Unmarshal(data []byte) error {
	return errors.New("TODO")
}
