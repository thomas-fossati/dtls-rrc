// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package protocol

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestRRCMsg(t *testing.T) {
	msgTypes := []RrcMsgType{
		RrcPathChallenge,
		RrcPathResponse,
		RrcPathDrop,
	}

	for _, typ := range msgTypes {
		out := RRC{
			typ:    typ,
			cookie: rand.Uint64(),
		}

		b, err := out.Marshal()
		if err != nil {
			t.Errorf("marshaling RRC failed: %s", err)
		}

		fmt.Printf("marshalled => %x\n", b)

		var in RRC

		err = in.Unmarshal(b)
		if err != nil {
			t.Errorf("unmarshaling RRC failed: %s", err)
		}

		if !reflect.DeepEqual(out, in) {
			t.Errorf("want: %#v, got: %#v", in, out)
		}
	}
}
