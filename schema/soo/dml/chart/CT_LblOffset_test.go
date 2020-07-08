// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml/chart"
)

func TestCT_LblOffsetConstructor(t *testing.T) {
	v := chart.NewCT_LblOffset()
	if v == nil {
		t.Errorf("chart.NewCT_LblOffset must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_LblOffset should validate: %s", err)
	}
}

func TestCT_LblOffsetMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_LblOffset()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_LblOffset()
	xml.Unmarshal(buf, v2)
}
