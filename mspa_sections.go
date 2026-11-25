package iabconsent

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const (
	// MspaUsNationalV1StringLength is 60 bits padded with 4 bits of 0s making valid length 64 bits
	MspaUsNationalV1StringLength = 64
	// MspaUsNationalV2StringLength is 70 bits padded with 2 bits of 0s making valid length 72 bits
	MspaUsNationalV2StringLength = 72
	// MspaUsCaV1StringLength is 46 bits padded with 2 bits of 0s making valid length 48 bits
	MspaUsCaV1StringLength = 48
	// MspaUsVaV1StringLength is 40 bits no padding needed
	MspaUsVaV1StringLength = 40
	// MspaUsCoV1StringLength is 38 bits padded with 2 bits of 0s making valid length 40 bits
	MspaUsCoV1StringLength = 40
	// MspaUsUtV1StringLength is 42 bits padded with 6 bits of 0s making valid length 48 bits
	MspaUsUtV1StringLength = 48
	// MspaUsCtV1StringLength is 44 bits padded with 4 bits of 0s making valid length 48 bits
	MspaUsCtV1StringLength = 48
	// MspaUsFlV1StringLength is 46 bits padded with 2 bits of 0s making valid length 48 bits
	MspaUsFlV1StringLength = 48
	// MspaUsMtV1StringLength is 46 bits padded with 2 bits of 0s making valid length 48 bits
	MspaUsMtV1StringLength = 48
	// MspaUsOrV1StringLength is 52 bits padded with 4 bits of 0s making valid length 56 bits
	MspaUsOrV1StringLength = 56
	// MspaUsTxV1StringLength is 42 bits padded with 6 bits of 0s making valid length 48 bits
	MspaUsTxV1StringLength = 48
	// MspaUsDeV1StringLength is 52 bits padded with 4 bits of 0s making valid length 56 bits
	MspaUsDeV1StringLength = 56
	// MspaUsIaV1StringLength is 42 bits padded with 6 bits of 0s making valid length 48 bits
	MspaUsIaV1StringLength = 48
	// MspaUsNeV1StringLength is 42 bits padded with 6 bits of 0s making valid length 48 bits
	MspaUsNeV1StringLength = 48
	// MspaUsNhV1StringLength is 46 bits padded with 2 bits of 0s making valid length 48 bits
	MspaUsNhV1StringLength = 48
	// MspaUsNjV1StringLength is 54 bits padded with 2 bits of 0s making valid length 56 bits
	MspaUsNjV1StringLength = 56
	// MspaUsTnV1StringLength is 42 bits padded with 6 bits of 0s making valid length 48 bits
	MspaUsTnV1StringLength = 48
	// MspaUsMnV1StringLength is 42 bits padded with 6 bits of 0s making valid length 48 bits
	MspaUsMnV1StringLength = 48
	// MspaUsMdV1CoreLength is 22 bits padded with 2 bits of 0s making valid length 24 bits
	MspaUsMdV1CoreLength = 24
	// MspaUsInV1CoreLength is 24 bits no padding needed
	MspaUsInV1CoreLength = 24
	// MspaUsInV1SensitiveDataConsentsLength is 16 bits no padding needed
	MspaUsInV1SensitiveDataConsentsLength = 16
	// MspaUsKyV1CoreLength is 24 bits no padding needed
	MspaUsKyV1CoreLength = 24
	// MspaUsKyV1SensitiveDataConsentsLength is 16 bits no padding needed
	MspaUsKyV1SensitiveDataConsentsLength = 16
	// MspaUsRiV1CoreLength is 24 bits no padding needed
	MspaUsRiV1CoreLength = 24
	// MspaUsRiV1SensitiveDataConsentsLength is 16 bits no padding needed
	MspaUsRiV1SensitiveDataConsentsLength = 16
)

const (
	// UsMdGpcSubSectionId is the ID of the GPC subsection in the US Maryland section header.
	UsMdGpcSubSectionId = 1
	// UsInSensitiveDataConsentsSubSectionId is the ID of the Sensitive Data Consents subsection in the US Indiana section header.
	UsInSensitiveDataConsentsSubSectionId = 1
	// UsKySensitiveDataConsentsSubSectionId is the ID of the Sensitive Data Consents subsection in the US Kentucky section header.
	UsKySensitiveDataConsentsSubSectionId = 1
	// UsRiSensitiveDataConsentsSubSectionId is the ID of the Sensitive Data Consents subsection in the US Rhode Island section header.
	UsRiSensitiveDataConsentsSubSectionId = 1
)

type MspaUsNational struct {
	GppSection
}

type MspaUsCA struct {
	GppSection
}

type MspaUsVA struct {
	GppSection
}

type MspaUsCO struct {
	GppSection
}

type MspaUsUT struct {
	GppSection
}

type MspaUsCT struct {
	GppSection
}

type MspaUsFL struct {
	GppSection
}

type MspaUsMT struct {
	GppSection
}

type MspaUsOR struct {
	GppSection
}

type MspaUsTX struct {
	GppSection
}

type MspaUsDE struct {
	GppSection
}

type MspaUsIA struct {
	GppSection
}

type MspaUsNE struct {
	GppSection
}

type MspaUsNH struct {
	GppSection
}

type MspaUsNJ struct {
	GppSection
}

type MspaUsTN struct {
	GppSection
}

type MspaUsMN struct {
	GppSection
}

type MspaUsMD struct {
	GppSection
}

type MspaUsIN struct {
	GppSection
}

type MspaUsKY struct {
	GppSection
}

type MspaUsRI struct {
	GppSection
}

// NewMspa returns a supported parser given a GPP Section ID.
// If the SID is not yet supported, it will be null.
func NewMspa(sid int, section string) GppSectionParser {
	switch sid {
	case UsNationalSID:
		return &MspaUsNational{GppSection{sectionId: UsNationalSID, sectionValue: section}}
	case UsCaliforniaSID:
		return &MspaUsCA{GppSection{sectionId: UsCaliforniaSID, sectionValue: section}}
	case UsVirginiaSID:
		return &MspaUsVA{GppSection{sectionId: UsVirginiaSID, sectionValue: section}}
	case UsColoradoSID:
		return &MspaUsCO{GppSection{sectionId: UsColoradoSID, sectionValue: section}}
	case UsUtahSID:
		return &MspaUsUT{GppSection{sectionId: UsUtahSID, sectionValue: section}}
	case UsConnecticutSID:
		return &MspaUsCT{GppSection{sectionId: UsConnecticutSID, sectionValue: section}}
	case UsFloridaSID:
		return &MspaUsFL{GppSection{sectionId: UsFloridaSID, sectionValue: section}}
	case UsMontanaSID:
		return &MspaUsMT{GppSection{sectionId: UsMontanaSID, sectionValue: section}}
	case UsOregonSID:
		return &MspaUsOR{GppSection{sectionId: UsOregonSID, sectionValue: section}}
	case UsTexasSID:
		return &MspaUsTX{GppSection{sectionId: UsTexasSID, sectionValue: section}}
	case UsDelawareSID:
		return &MspaUsDE{GppSection{sectionId: UsDelawareSID, sectionValue: section}}
	case UsIowaSID:
		return &MspaUsIA{GppSection{sectionId: UsIowaSID, sectionValue: section}}
	case UsNebraskaSID:
		return &MspaUsNE{GppSection{sectionId: UsNebraskaSID, sectionValue: section}}
	case UsNewHampshireSID:
		return &MspaUsNH{GppSection{sectionId: UsNewHampshireSID, sectionValue: section}}
	case UsNewJerseySID:
		return &MspaUsNJ{GppSection{sectionId: UsNewJerseySID, sectionValue: section}}
	case UsTennesseeSID:
		return &MspaUsTN{GppSection{sectionId: UsTennesseeSID, sectionValue: section}}
	case UsMinnesotaSID:
		return &MspaUsMN{GppSection{sectionId: UsMinnesotaSID, sectionValue: section}}
	case UsMarylandSID:
		return &MspaUsMD{GppSection{sectionId: UsMarylandSID, sectionValue: section}}
	case UsIndianaSID:
		return &MspaUsIN{GppSection{sectionId: UsIndianaSID, sectionValue: section}}
	case UsKentuckySID:
		return &MspaUsKY{GppSection{sectionId: UsKentuckySID, sectionValue: section}}
	case UsRhodeIslandSID:
		return &MspaUsRI{GppSection{sectionId: UsRhodeIslandSID, sectionValue: section}}
	}
	// Skip if no matching struct, as Section ID is not supported yet.
	// Any newly supported Section IDs should be added as cases here.
	return nil
}

func (m *MspaUsNational) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usnat consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/blob/main/Sections/US-National/IAB%20Privacy%E2%80%99s%20Multi-State%20Privacy%20Agreement%20(MSPA)%20US%20National%20Technical%20Specification.md
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	// Support both v1 and v2
	if p.Version != 1 && p.Version != 2 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1 and v2

	if p.Version == 1 && r.Size() != MspaUsNationalV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	} else if p.Version == 2 && r.Size() != MspaUsNationalV2StringLength {
		return nil, errors.New("invalid consent string length for v2")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.SharingOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SensitiveDataProcessingOptOutNotice, _ = r.ReadMspaNotice()
	p.SensitiveDataLimitUseNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.SharingOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()

	// see spec in IAB GPP repo for differences between v1 and v2
	if p.Version == 1 {
		p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(12)
		p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(2)
	} else if p.Version == 2 {
		p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(16)
		p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(3)
	}
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsCA) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usca consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/CA
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsCaV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.SharingOptOutNotice, _ = r.ReadMspaNotice()
	p.SensitiveDataLimitUseNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.SharingOptOut, _ = r.ReadMspaOptOut()
	// SensitiveDataProcessingOptOuts, as opposed to Consent.
	p.SensitiveDataProcessingOptOuts, _ = r.ReadMspaBitfieldOptOut(9)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(2)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsVA) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usva consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/VA
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsVaV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsCO) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usco consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/CO
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsCoV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(7)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsUT) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usut consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/UT
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsUtV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SensitiveDataProcessingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingOptOuts, _ = r.ReadMspaBitfieldOptOut(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsCT) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usct consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/CT
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsCtV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(3)
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsFL) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usfl consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/FL
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsFlV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(3)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsMT) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usmt consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/MT
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsMtV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(3)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsOR) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usor consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/OR
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsOrV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(11)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(3)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsTX) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse ustx consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/TX
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsTxV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsDE) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usde consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/DE
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsDeV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(9)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(5)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	// 0 is not a valid value according to the docs for MspaCoveredTransaction. Instead of erroring,
	// return the value of the string, and let downstream processing handle if the value is 0.
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

// Fix Iowa implementation
func (m *MspaUsIA) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usia consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/IA
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsIaV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SensitiveDataProcessingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingOptOuts, _ = r.ReadMspaBitfieldOptOut(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsNE) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usne consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/NE
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsNeV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsNH) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usnh consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/NH
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsNhV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(3)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsNJ) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usnj consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/NJ
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsNjV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(10)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(5)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsTN) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse ustn consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/TN
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if p.Version == 1 && r.Size() != MspaUsTnV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

func (m *MspaUsMN) ParseConsent() (GppParsedConsent, error) {
	var segments = strings.Split(m.sectionValue, ".")

	var b, err = base64.RawURLEncoding.DecodeString(segments[0])
	if err != nil {
		return nil, errors.Wrap(err, "parse usmn consent string")
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/MN
	var p = &MspaParsedConsent{}
	p.Version, _ = r.ReadInt(6)

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	// validate the length of the bit string for v1
	if r.Size() != MspaUsMnV1StringLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	// ProcessingNotice
	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.SensitiveDataProcessingConsents, _ = r.ReadMspaBitfieldConsent(8)
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	// AdditionalDataProcessingConsent
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaOptOutOptionMode, _ = r.ReadMspaNaYesNo()
	p.MspaServiceProviderMode, _ = r.ReadMspaNaYesNo()

	if len(segments) > 1 {
		var gppSubsectionConsent *GppSubSection
		gppSubsectionConsent, err = ParseGppSubSections(segments[1:])
		if err != nil {
			return p, err
		}
		p.Gpc = gppSubsectionConsent.Gpc
	}

	return p, r.Err
}

// MspaSectionHeader is the header of sections which encode their own section header
// (e.g. US Maryland, Indiana, Kentucky, Rhode Island). Format is:
// SectionID	Int(6)	Section ID from the Section IDs list
// Version	Int(6)	The version of this section specification used to encode the string
// SubSections	Range(Fibonacci)	List of subsection IDs that are contained in this section, in the
// order they appear in the string. The Core subsection (ID 0) is always present first and not listed.
type MspaSectionHeader struct {
	SectionId   int
	Version     int
	SubSections []int
}

// ParseMspaSectionHeader parses the section header of sections which encode their own section header.
func ParseMspaSectionHeader(s string) (*MspaSectionHeader, error) {
	var b, err = decodeGppSegment(s)
	if err != nil {
		return nil, err
	}
	var r = NewConsentReader(b)
	var h = &MspaSectionHeader{}
	if h.SectionId, err = r.ReadInt(6); err != nil {
		return nil, errors.WithMessage(err, "read section id")
	}
	if h.Version, err = r.ReadInt(6); err != nil {
		return nil, errors.WithMessage(err, "read version")
	}
	if h.SubSections, err = r.ReadFibonacciRange(); err != nil {
		return nil, errors.WithMessage(err, "read subsections")
	}
	return h, nil
}

// splitMspaSectionWithHeader splits a section of format {section header}.{core}[.{subsection}...] into
// its segments, and parses and validates the section header against the expected Section ID.
// It guarantees that the core segment and all subsections listed in the header are present in segments.
// Version validation is left to the section specific parsers.
func splitMspaSectionWithHeader(sid int, key string, sectionValue string) (*MspaSectionHeader, []string, error) {
	var segments = strings.Split(sectionValue, ".")
	if len(segments) < 2 {
		return nil, nil, errors.New("parse " + key + " consent string: missing core subsection")
	}

	var header, err = ParseMspaSectionHeader(segments[0])
	if err != nil {
		return nil, nil, errors.Wrap(err, "parse "+key+" section header")
	}
	if header.SectionId != sid {
		return nil, nil, errors.Errorf("parse %s section header: section id mismatch: %d", key, header.SectionId)
	}
	if len(header.SubSections) > len(segments)-2 {
		return nil, nil, errors.New("mismatch number of subsections")
	}
	return header, segments, nil
}

// mspaSubSectionParser parses a single base64 encoded subsection segment into p.
type mspaSubSectionParser func(p *MspaParsedConsent, subSection string) error

// parseMspaSubSections parses the subsections of a section which encodes its own section header.
// Subsections appear in the string in the order listed in the section header, and each one is parsed
// with the parser registered for its subsection ID. Subsections without a parser are skipped.
// segments must be validated by splitMspaSectionWithHeader.
func parseMspaSubSections(p *MspaParsedConsent, header *MspaSectionHeader, segments []string,
	parsers map[int]mspaSubSectionParser) error {
	for i, subSectionId := range header.SubSections {
		if parse, ok := parsers[subSectionId]; ok {
			if err := parse(p, segments[i+2]); err != nil {
				return err
			}
		}
	}
	return nil
}

// parseMspaGpcSubSection parses a GPC subsection and sets p.Gpc.
func parseMspaGpcSubSection(p *MspaParsedConsent, subSection string) error {
	var gppSubsectionConsent, err = ParseGppSubSections([]string{subSection})
	if err != nil {
		return err
	}
	p.Gpc = gppSubsectionConsent.Gpc
	return nil
}

// newMspaSensitiveDataConsentsSubSectionParser returns a parser of a Sensitive Data Consents subsection
// with the given number of 2 bit consent fields and expected (padded) length in bits.
func newMspaSensitiveDataConsentsSubSectionParser(fields uint, length int) mspaSubSectionParser {
	return func(p *MspaParsedConsent, subSection string) error {
		var b, err = base64.RawURLEncoding.DecodeString(subSection)
		if err != nil {
			return errors.Wrap(err, "parse sensitive data consents subsection")
		}
		var r = NewConsentReader(b)
		if r.Size() != length {
			return errors.New("invalid sensitive data consents subsection length")
		}
		p.SensitiveDataProcessingConsents, err = r.ReadMspaBitfieldConsent(fields)
		return err
	}
}

func (m *MspaUsMD) ParseConsent() (GppParsedConsent, error) {
	var header, segments, err = splitMspaSectionWithHeader(m.sectionId, "usmd", m.sectionValue)
	if err != nil {
		return nil, err
	}

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/MD
	var p = &MspaParsedConsent{}
	p.Version = header.Version

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	var b []byte
	b, err = base64.RawURLEncoding.DecodeString(segments[1])
	if err != nil {
		return nil, errors.Wrap(err, "parse usmd consent string")
	}

	var r = NewConsentReader(b)

	// validate the length of the core subsection for v1
	if r.Size() != MspaUsMdV1CoreLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.MspaVersion, _ = r.ReadInt(6)
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaMode, _ = r.ReadMspaMode()
	// ProcessingNotice
	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	// AdditionalDataProcessingConsent
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	if r.Err != nil {
		return nil, r.Err
	}

	err = parseMspaSubSections(p, header, segments, map[int]mspaSubSectionParser{
		UsMdGpcSubSectionId: parseMspaGpcSubSection,
	})
	if err != nil {
		return p, err
	}

	return p, nil
}

func (m *MspaUsIN) ParseConsent() (GppParsedConsent, error) {
	var header, segments, err = splitMspaSectionWithHeader(m.sectionId, "usin", m.sectionValue)
	if err != nil {
		return nil, err
	}

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/IN
	var p = &MspaParsedConsent{}
	p.Version = header.Version

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	var b []byte
	b, err = base64.RawURLEncoding.DecodeString(segments[1])
	if err != nil {
		return nil, errors.Wrap(err, "parse usin consent string")
	}

	var r = NewConsentReader(b)

	// validate the length of the core subsection for v1
	if r.Size() != MspaUsInV1CoreLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.MspaVersion, _ = r.ReadInt(6)
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaMode, _ = r.ReadMspaMode()
	// ProcessingNotice
	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	// AdditionalDataProcessingConsent
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	if r.Err != nil {
		return nil, r.Err
	}

	err = parseMspaSubSections(p, header, segments, map[int]mspaSubSectionParser{
		UsInSensitiveDataConsentsSubSectionId: newMspaSensitiveDataConsentsSubSectionParser(8, MspaUsInV1SensitiveDataConsentsLength),
	})
	if err != nil {
		return p, err
	}

	return p, nil
}

func (m *MspaUsKY) ParseConsent() (GppParsedConsent, error) {
	var header, segments, err = splitMspaSectionWithHeader(m.sectionId, "usky", m.sectionValue)
	if err != nil {
		return nil, err
	}

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/KY
	var p = &MspaParsedConsent{}
	p.Version = header.Version

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	var b []byte
	b, err = base64.RawURLEncoding.DecodeString(segments[1])
	if err != nil {
		return nil, errors.Wrap(err, "parse usky consent string")
	}

	var r = NewConsentReader(b)

	// validate the length of the core subsection for v1
	if r.Size() != MspaUsKyV1CoreLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.MspaVersion, _ = r.ReadInt(6)
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaMode, _ = r.ReadMspaMode()
	// ProcessingNotice
	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	// AdditionalDataProcessingConsent
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	if r.Err != nil {
		return nil, r.Err
	}

	err = parseMspaSubSections(p, header, segments, map[int]mspaSubSectionParser{
		UsKySensitiveDataConsentsSubSectionId: newMspaSensitiveDataConsentsSubSectionParser(8, MspaUsKyV1SensitiveDataConsentsLength),
	})
	if err != nil {
		return p, err
	}

	return p, nil
}

func (m *MspaUsRI) ParseConsent() (GppParsedConsent, error) {
	var header, segments, err = splitMspaSectionWithHeader(m.sectionId, "usri", m.sectionValue)
	if err != nil {
		return nil, err
	}

	// This block of code directly describes the format of the payload.
	// The spec for the consent string can be found here:
	// https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/tree/main/Sections/US-States/RI
	var p = &MspaParsedConsent{}
	p.Version = header.Version

	if p.Version != 1 {
		return nil, errors.New("unsupported version: " + fmt.Sprint(p.Version))
	}

	var b []byte
	b, err = base64.RawURLEncoding.DecodeString(segments[1])
	if err != nil {
		return nil, errors.Wrap(err, "parse usri consent string")
	}

	var r = NewConsentReader(b)

	// validate the length of the core subsection for v1
	if r.Size() != MspaUsRiV1CoreLength {
		return nil, errors.New("invalid consent string length for v1")
	}

	p.MspaVersion, _ = r.ReadInt(6)
	p.MspaCoveredTransaction, _ = r.ReadMspaNaYesNo()
	p.MspaMode, _ = r.ReadMspaMode()
	// ProcessingNotice
	p.SharingNotice, _ = r.ReadMspaNotice()
	p.SaleOptOutNotice, _ = r.ReadMspaNotice()
	p.TargetedAdvertisingOptOutNotice, _ = r.ReadMspaNotice()
	p.SaleOptOut, _ = r.ReadMspaOptOut()
	p.TargetedAdvertisingOptOut, _ = r.ReadMspaOptOut()
	p.KnownChildSensitiveDataConsents, _ = r.ReadMspaBitfieldConsent(1)
	// AdditionalDataProcessingConsent
	p.PersonalDataConsents, _ = r.ReadMspaConsent()
	if r.Err != nil {
		return nil, r.Err
	}

	err = parseMspaSubSections(p, header, segments, map[int]mspaSubSectionParser{
		UsRiSensitiveDataConsentsSubSectionId: newMspaSensitiveDataConsentsSubSectionParser(8, MspaUsRiV1SensitiveDataConsentsLength),
	})
	if err != nil {
		return p, err
	}

	return p, nil
}
