// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/wml"
)

func TestWdCT_PosVChoiceConstructor(t *testing.T) {
	v := wml.NewWdCT_PosVChoice()
	if v == nil {
		t.Errorf("wml.NewWdCT_PosVChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_PosVChoice should validate: %s", err)
	}
}

func TestWdCT_PosVChoiceMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_PosVChoice()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_PosVChoice()
	xml.Unmarshal(buf, v2)
}
