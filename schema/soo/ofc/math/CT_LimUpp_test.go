// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/ofc/math"
)

func TestCT_LimUppConstructor(t *testing.T) {
	v := math.NewCT_LimUpp()
	if v == nil {
		t.Errorf("math.NewCT_LimUpp must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_LimUpp should validate: %s", err)
	}
}

func TestCT_LimUppMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_LimUpp()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_LimUpp()
	xml.Unmarshal(buf, v2)
}
