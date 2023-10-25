// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package protocol

import (
	"encoding/binary"
	"errors"
)

var (
	errRrcUnknownMsgType = errors.New("unknown rrc_msg_type")
)

// RRC messages
// https://www.ietf.org/archive/id/draft-ietf-tls-dtls-rrc-10.html#section-4
type RrcMsgType uint8

const (
	RrcPathChallenge RrcMsgType = iota
	RrcPathResponse
	RrcPathDrop
)

const (
	RRCMsgSize = 9
)

type RRC struct {
	typ    RrcMsgType
	cookie uint64
}

// ContentType returns the RRC ContentType
func (a RRC) ContentType() ContentType {
	return ContentTypeRRC
}

// Marshal encodes the RRC to binary
func (a *RRC) Marshal() ([]byte, error) {
	out := make([]byte, RRCMsgSize)

	switch a.typ {
	case RrcPathChallenge, RrcPathResponse, RrcPathDrop:
		out[0] = byte(a.typ)
	default:
		return nil, errRrcUnknownMsgType
	}

	binary.BigEndian.PutUint64(out[1:], uint64(a.cookie))

	return out, nil
}

// Unmarshal populates the RRC from binary
func (a *RRC) Unmarshal(data []byte) error {
	if len(data) < RRCMsgSize {
		return errBufferTooSmall
	}

	t := RrcMsgType(data[0])

	switch t {
	case RrcPathChallenge, RrcPathResponse, RrcPathDrop:
		a.typ = t
	default:
		return errRrcUnknownMsgType
	}

	a.cookie = binary.BigEndian.Uint64(data[1:])

	return nil
}
