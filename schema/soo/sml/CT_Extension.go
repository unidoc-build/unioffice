// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_Extension struct {
	// URI
	UriAttr *string
	Any     unioffice.Any
}

func NewCT_Extension() *CT_Extension {
	ret := &CT_Extension{}
	return ret
}

func (m *CT_Extension) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UriAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uri"},
			Value: fmt.Sprintf("%v", *m.UriAttr)})
	}
	e.EncodeToken(start)
	if m.Any != nil {
		m.Any.MarshalXML(e, xml.StartElement{})
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Extension) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uri" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UriAttr = &parsed
			continue
		}
	}
lCT_Extension:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			default:
				if anyEl, err := unioffice.CreateElement(el); err != nil {
					return err
				} else {
					if err := d.DecodeElement(anyEl, &el); err != nil {
						return err
					}
					m.Any = anyEl
				}
			}
		case xml.EndElement:
			break lCT_Extension
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Extension and its children
func (m *CT_Extension) Validate() error {
	return m.ValidateWithPath("CT_Extension")
}

// ValidateWithPath validates the CT_Extension and its children, prefixing error messages with path
func (m *CT_Extension) ValidateWithPath(path string) error {
	return nil
}
