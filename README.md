# iabconsent
[![Build Status][ci]](https://travis-ci.com/LiveRamp/iabconsent)

[ci]: https://travis-ci.com/LiveRamp/iabconsent.svg?branch=master "Build Status"

[![Go Report Card][report]](https://goreportcard.com/report/github.com/LiveRamp/iabconsent)

[report]: https://goreportcard.com/badge/github.com/LiveRamp/iabconsent "Go Report Card"

A Golang implementation of the:
- IAB Consent String 1.1 Spec
- IAB Transparency and Consent String v2.0-v2.2
- IAB Tech Lab Global Privacy Platform (GPP) Spec v1.0 Sections:
  - US National Multi-State Privacy Agreement
  - US California Multi-State Privacy Agreement
  - US Virginia Multi-State Privacy Agreement
  - US Colorado Multi-State Privacy Agreement
  - US Utah Multi-State Privacy Agreement
  - US Connecticut Multi-State Privacy Agreement
  - US Florida, Montana, Oregon, Texas, Delaware, Iowa, Nebraska, New Hampshire, New Jersey, Tennessee
  - US Minnesota (23), Maryland (24), Indiana (25), Kentucky (26), Rhode Island (27)
    - Maryland, Indiana, Kentucky and Rhode Island use the section-header encoding
      (`{header}.{core}[.{subsection}]`). Their single `MspaMode` field is exposed as `MspaParsedConsent.MspaMode`;
      `MspaOptOutOptionMode` / `MspaServiceProviderMode` are not set for these sections.

To install:
```
go get -v github.com/LiveRamp/iabconsent
```

# Transparency and Consent Framework v1.1 + v2.0-v2.2

This package defines two structs (`ParsedConsent` and `V2ParsedConsent`) which contain all of the fields of the IAB 
TCF v1.1 Consent String and the IAB Transparency and Consent String v2.0 respectively.

Each spec has their own parsing function (`ParseV1` and `ParseV2`). If the caller is unsure of which version a consent
string is, they can use `TCFVersionFromTCString` to receive a `TCFVersion` enum type to determine which parsing function is appropriate.

Example use:
```go
package main

import "github.com/LiveRamp/iabconsent"

func main() {
    var consent = "COvzTO5OvzTO5BRAAAENAPCoALIAADgAAAAAAewAwABAAlAB6ABBFAAA"

    var c, err = iabconsent.TCFVersionFromTCString(consent)
    if err != nil {
        panic(err)
    }
    
    switch c {
    case iabconsent.V1:
        var v1, err = iabconsent.ParseV1(consent)
        // Use v1 consent.
    case iabconsent.V2:
        var v2, err = iabconsent.ParseV2(consent)
        // Use v2 consent.
    default:
        panic("unknown version")
    }
}
```

The function `Parse(s string)` is deprecated, and should no longer be used.

# Global Privacy Platform v1.0

This package defines two structs (`GPPHeader` and `GppParsedConsent`) which contain the fields of the GPP Header and GPP Sections respectively. 
`GppParsedConsent` itself is broad, as a given GPP String may contain different sections that have their own unique privacy specifications.

All supported sections of the Multi-State Privacy Agreement via GPP have their own struct `MspaParsedConsent`.

There are two ways of working with the GPP string.
1. Getting the Parsing Functions
   - `MapGppSectionToParser` takes the full string, parses and processes the header to get the remaining sections, and maps sections to a parsing function (if supported). This allows the user to determine how/when they want to parse the sections.
2. Parse the Entire String
   - `ParseGppConsent` takes the full string, parses and process the header and all supported sections consecutively, returning the ParsedConsents.


Example use:
```go
package main

import "github.com/LiveRamp/iabconsent"

func main() {
	var consent = "DBABrGA~BVVqAAEABCA~BVoYYZoI~BVoYYYI~BVoYYQg~BVaGGGCA~BVoYYYQg"

	// Parse Entire String via function.
	var gppConsents, err = iabconsent.ParseGppConsent(consent)
	if err != nil {
		panic(err)
	}
	var usNational = gppConsents[7]
	var mspa, ok = usNational.(*iabconsent.MspaParsedConsent)
	if !ok {
		// Not MSPA ParsedConsent
	}
	if mspa.Version == 1 {
		// Can check specific values/fields to determine your own requirements to process.
	}

	// Get Parsing Functions, and parse on your own.
	var gppConsentFunctions, errMap = iabconsent.MapGppSectionToParser(consent)
	if errMap != nil {
		panic(err)
	}
	var usNationalParsed iabconsent.GppParsedConsent
	var errParse error
	for _, gppSection := range gppConsentFunctions {
		if gppSection.GetSectionId() == 7 {
			usNationalParsed, errParse = gppSection.ParseConsent()
			if errParse != nil {
				panic(err)
			}
		}
	}

	var parsed, pOk = usNationalParsed.(*iabconsent.MspaParsedConsent)
	if !pOk {
		// Not MSPA ParsedConsent
	}
	if parsed.Version == 1 {
		// Can check specific values/fields to determine your own requirements to process.
	}
}
```

## Global Privacy Platform custom parsers

Each way also support the ability to customize the parsers used. This is useful if you want to use your own custom
parsers for a given section, or if you want to support a section that is not currently supported by this package.

Example custom sections parser use:
```go
package main

import "github.com/LiveRamp/iabconsent"

func main() {
	var consent = "DBABrGA~BVVqAAEABCA~BVoYYZoI~BVoYYYI~BVoYYQg~BVaGGGCA~BVoYYYQg"

	// Parse Entire String via function using your custom parser's setup
	var gppConsents, err = iabconsent.ParseGppConsent(consent, customParsers())
	if err != nil {
		panic(err)
	}
	var tcfEuConsent = gppConsents[2]
	var tcfEu, ok = tcfEuConsent.(*iabconsent.V2ParsedConsent)
	if !ok {
		// Not TCFv2 EU ParsedConsent
	}
	if tcfEu.Version == 2 {
		// Can check specific values/fields to determine your own requirements to process.
	}
}

func customParsers() *iabconsent.Options {
	return &iabconsent.Options{
		GppSectionParser: func(gppSectionID int, sectionString string) iabconsent.GppSectionParser {
			switch gppSectionID {
			case 2:
				return &TcfV2GppSectionParser{
					GppSectionParser{sectionId: gppSectionID, sectionValue: sectionString},
				}
			}

			return iabconsent.NewMspa(gppSectionID, sectionString)
		},
	}
}

type GppSectionParser struct {
	sectionValue string
	sectionId    int
}

type TcfV2GppSectionParser struct {
	GppSectionParser
}

func (p *GppSectionParser) GetSectionId() int {
	return p.sectionId
}

func (p *TcfV2GppSectionParser) ParseConsent() (iabconsent.GppParsedConsent, error) {
	tcfEuV2Parsed, err := iabconsent.ParseV2(p.sectionValue)

	if err != nil {
		return nil, err
	}

	return tcfEuV2Parsed, nil
}
```