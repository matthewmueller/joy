package generator

import (
	"bytes"
	"encoding/xml"
)

// dom struct
type dom struct {
	WebIDL             xml.Name             `xml:"webidl-xml"`
	CallbackFunctions  []*CallbackFunction  `xml:"callback-functions>callback-function"`
	CallbackInterfaces []*CallbackInterface `xml:"callback-interfaces>interface"`
	Dictionaries       []*Dictionary        `xml:"dictionaries>dictionary"`
	Enums              []*Enum              `xml:"enums>enum"`
	Interfaces         []*Interface         `xml:"interfaces>interface"`
	MixinInterfaces    []*Interface         `xml:"mixin-interfaces>interface"`
	TypeDefs           []*Typedef           `xml:"typedefs>typedef"`
}

func parse(src string) (*dom, error) {
	var d dom
	xmlDecoder := xml.NewDecoder(bytes.NewBufferString(src))
	if e := xmlDecoder.Decode(&d); e != nil {
		return nil, e
	}
	return &d, nil
}
