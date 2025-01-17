// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package protocol

import (
	"encoding/binary"
	"errors"
	"fmt"
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

func (r RrcMsgType) String() string {
	switch r {
	case RrcPathChallenge:
		return "path_challenge"
	case RrcPathResponse:
		return "path_response"
	case RrcPathDrop:
		return "path_drop"
	}

	return fmt.Sprintf("unknown_rrc_msg(%d)", r)
}

const (
	RRCMsgSize = 9
)

type RRC struct {
	Type   RrcMsgType
	Cookie uint64
}

// ContentType returns the RRC ContentType
func (a RRC) ContentType() ContentType {
	return ContentTypeRRC
}

// Marshal encodes the RRC to binary
func (a *RRC) Marshal() ([]byte, error) {
	out := make([]byte, RRCMsgSize)

	switch a.Type {
	case RrcPathChallenge, RrcPathResponse, RrcPathDrop:
		out[0] = byte(a.Type)
	default:
		return nil, errRrcUnknownMsgType
	}

	binary.BigEndian.PutUint64(out[1:], uint64(a.Cookie))

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
		a.Type = t
	default:
		return errRrcUnknownMsgType
	}

	a.Cookie = binary.BigEndian.Uint64(data[1:])

	return nil
}
