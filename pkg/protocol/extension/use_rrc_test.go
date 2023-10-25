// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

import (
	"testing"
)

func TestExtensionRRC(t *testing.T) {
	out := UseRRC{
		Supported: true,
	}

	b, err := out.Marshal()
	if err != nil {
		t.Errorf("marshaling RRC extension failed: %s", err)
	}

	var in UseRRC

	err = in.Unmarshal(b)
	if err != nil {
		t.Errorf("unmarshaling RRC extension failed: %s", err)
	}

	if !in.Supported {
		t.Error("expecting RRC supported, got RRC unsupported")
	}
}
