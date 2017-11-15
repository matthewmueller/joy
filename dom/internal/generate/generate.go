//+build ignore

package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"

	"github.com/apex/log"
	"github.com/kr/pretty"
)

type CallbackFunction struct {
	XMLName  xml.Name `xml:"callback-function"`
	Name     string   `xml:"name,attr"`
	Callback bool     `xml:"callback,attr,omitempty"`
	Type     string   `xml:"type,attr"`
	Params   []Param  `xml:"param"`
}

type Param struct {
	Name     string `xml:"name,attr"`
	Type     string `xml:"type,attr,omitempty"`
	Optional bool   `xml:"optional,attr,omitempty"`
	Variadic bool   `xml:"variadic,attr,omitempty"`
}

type Method struct {
	Name                     string  `xml:"name,attr"`
	Type                     string  `xml:"type,attr"`
	Creator                  bool    `xml:"creator,attr"`
	Setter                   bool    `xml:"setter,attr"`
	DoNotCheckDomainSecurity bool    `xml:"do-not-check-domain-security,attr"`
	Params                   []Param `xml:"param"`
}

type Interface struct {
	Name              string   `xml:"name,attr"`
	Extends           string   `xml:"extends,attr"`
	Implements        []string `xml:"implements,omitempty"`
	NoInterfaceObject bool     `xml:"no-interface-object,attr"`
	Element           []struct {
		Name            string `xml:"name,attr"`
		HTMLSelfClosing bool   `xml:"html-self-closing,attr"`
		Namespace       string `xml:"namespace,attr"`
	} `xml:"element"`
	Methods          []Method `xml:"methods>method"`
	AnonymousMethods []Method `xml:"anonymous-methods>method"`
	Properties       []struct {
		Name                              string `xml:"name,attr"`
		Type                              string `xml:"type,attr"`
		ReadOnly                          bool   `xml:"read-only,attr"`
		ContentAttribute                  string `xml:"content-attribute,attr"`
		ContentAttributeReflects          bool   `xml:"content-attribute-reflects,attr"`
		ContentAttributeValueSyntax       string `xml:"content-attribute-value-syntax,attr"`
		ContentAttributeEnumValues        string `xml:"content-attribute-enum-values,attr"`
		EventHandler                      string `xml:"event-handler,attr"`
		TreatNullAs                       string `xml:"treat-null-as,attr"`
		Nullable                          bool   `xml:"nullable,attr"`
		Replaceable                       bool   `xml:"replaceable,attr`
		PropertyDescriptorNotConfigurable bool   `xml:"property-descriptor-not-configurable,attr"`
		DoNotCheckDomainSecurity          bool   `xml:"do-not-check-domain-security,attr"`
		Unforgeable                       bool   `xml:"unforgeable,attr"`
	} `xml:"properties>property"`
	Constants []struct {
		Name         string `xml:"name,attr"`
		Type         string `xml:"type,attr"`
		TypeOriginal string `xml:"type-original,attr`
		Value        string `xml:"value,attr"`
	} `xml:"constants>constant"`
	Constructor struct {
		Params []Param `xml:"param"`
	} `xml:"constructor,omitempty"`
	NamedConstructor struct {
		Name   string  `xml:"name,attr"`
		Params []Param `xml:"param"`
	} `xml:"named-constructor"`
	Events []struct {
		Name        string `xml:"name,attr"`
		Bubbles     bool   `xml:"bubbles,attr"`
		Cancelable  bool   `xml:"cancelable,attr"`
		Tags        string `xml:"tags,attr"`
		Dispatch    string `xml:"dispatch,attr"`
		Follows     string `xml:"follows,attr"`
		Precedes    string `xml:"precedes,attr"`
		SkipsWindow bool   `xml:"skips-window,attr"`
		Type        string `xml:"type,attr"`
		Aliases     string `xml:"aliases,attr"`
	} `xml:"events>event"`
	AnonymousContentAttributes []struct {
		Name        string `xml:"name,attr"`
		EnumValues  string `xml:"enum-values,attr"`
		ValueSyntax string `xml:"value-syntax,attr"`
	} `xml:"anonymous-content-attributes>parsedattribute"`
}

type Dictionary struct {
	Name    string `xml:"name,attr"`
	Extends string `xml:"extends,attr"`
	Members []struct {
		Name     string `xml:"name,attr"`
		Type     string `xml:"type,attr,omitempty"`
		Required bool   `xml:"required,attr,omitempty"`
		Default  string `xml:"default,attr,omitempty"`
		Nullable bool   `xml:"nullable,attr,omitempty"`
	} `xml:"members>member"`
}

type Enum struct {
	Name   string   `xml:"name,attr"`
	Values []string `xml:"value"`
}

type Typedef struct {
	NewType string `xml:"new-type,attr"`
	Type    string `xml:"type,attr"`
}

type DOM struct {
	WebIDL             xml.Name           `xml:"webidl-xml"`
	CallbackFunctions  []CallbackFunction `xml:"callback-functions>callback-function"`
	CallbackInterfaces []Interface        `xml:"callback-interfaces>interface"`
	Dictionaries       []Dictionary       `xml:"dictionaries>dictionary"`
	Enums              []Enum             `xml:"enums>enum"`
	Interfaces         []Interface        `xml:"interfaces>interface"`
	MixinInterfaces    []Interface        `xml:"mixin-interfaces>interface"`
	TypeDefs           []Typedef          `xml:"typedefs>typedef"`
}

func main() {
	src, err := ioutil.ReadFile("./browser.webidl.xml")
	if err != nil {
		log.WithError(err).Errorf("error reading file")
		return
	}

	var dom DOM
	decoder := xml.NewDecoder(bytes.NewBuffer(src))
	if e := decoder.Decode(&dom); e != nil {
		log.WithError(e).Errorf("error decoding")
		return
	}

	pretty.Print(dom)

	// 	cwd, err := os.Getwd()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Infof("cwd=%s", cwd)
}
