package iabconsent

// MspaParsedConsent represents data extract from a Multi-State Privacy Agreement (mspa) consent string.
// Format can be found here: https://github.com/InteractiveAdvertisingBureau/Global-Privacy-Platform/blob/main/Sections/US-National/IAB%20Privacy%E2%80%99s%20National%20Privacy%20Technical%20Specification.md#core-segment
type MspaParsedConsent struct {
	// The version of this section specification used to encode the string.
	Version int
	// Notice of the Sharing of the Consumer’s Personal Data with Third Parties.
	// 0 Not Applicable. The Business does not share Personal Data with Third Parties.
	// 1 Yes, notice was provided
	// 2 No, notice was not provided
	// Note: ProcessingNotice and SharingNotice are the same
	SharingNotice MspaNotice
	// Notice of the Opportunity to Opt Out of the Sale of the Consumer’s Personal Data.
	// 0 Not Applicable. The Business does not Sell Personal Data.
	// 1 Yes, notice was provided
	// 2 No, notice was not provided
	SaleOptOutNotice MspaNotice
	// Notice of the Opportunity to Opt Out of the Sharing of the Consumer’s Personal Data.
	// 0 Not Applicable.The Business does not Share Personal Data.
	// 1 Yes, notice was provided
	// 2 No, notice was not provided
	SharingOptOutNotice MspaNotice
	// Notice of the Opportunity to Opt Out of Processing of the Consumer’s Personal Data for Targeted Advertising.
	// 0 Not Applicable.The Business does not Process Personal Data for Targeted Advertising.
	// 1 Yes, notice was provided
	// 2 No, notice was not provided
	TargetedAdvertisingOptOutNotice MspaNotice
	// Notice of the Opportunity to Opt Out of the Processing of the Consumer’s Sensitive Data.
	// 0 Not Applicable. The Business does not Process Sensitive Data.
	// 1 Yes, notice was provided
	// 2 No, notice was not provided
	// Note: SensitiveDataOptOutNotice and SensitiveDataProcessingOptOutNotice are the same
	SensitiveDataProcessingOptOutNotice MspaNotice
	// Notice of the Opportunity to Limit Use or Disclosure of the Consumer’s Sensitive Data.
	// 0 Not Applicable. The Business does not use or disclose Sensitive Data.
	// 1 Yes, notice was provided
	// 2 No, notice was not provided
	SensitiveDataLimitUseNotice MspaNotice
	// Opt-Out of the Sale of the Consumer’s Personal Data.
	// 0 Not Applicable. SaleOptOutNotice value was not applicable or no notice was provided
	// 1 Opted Out
	// 2 Did Not Opt Out
	SaleOptOut MspaOptout
	// Opt-Out of the Sharing of the Consumer’s Personal Data.
	// 0 Not Applicable. SharingOptOutNotice value was not applicable or no notice was provided.
	// 1 Opted Out
	// 2 Did Not Opt Out
	SharingOptOut MspaOptout
	// Opt-Out of Processing the Consumer’s Personal Data for Targeted Advertising.
	// 0 Not Applicable. TargetedAdvertisingOptOutNotice value was not applicable or no notice was provided
	// 1 Opted Out
	// 2 Did Not Opt Out
	TargetedAdvertisingOptOut MspaOptout
	// Two bits for each Data Activity:
	// 0 Not Applicable. The Business does not Process the specific category of Sensitive Data.
	// 1 No Consent
	// 2 Consent
	// Sensitive Data Consents is an optional subsection: the map is nil when
	// the subsection is absent, meaning "not provided" rather than all Not Applicable.
	SensitiveDataProcessingConsents map[int]MspaConsent
	// Two bits for each Data Activity:
	// 0 Not Applicable. SensitiveDataLimitUseNotice value was not applicable or no notice was provided.
	// 1 Opted Out
	// 2 Did Not Opt Out
	SensitiveDataProcessingOptOuts map[int]MspaOptout
	// Two bits for each Data Activity:
	// 0 Not Applicable. The Business does not have actual knowledge that it Processes Personal Data or Sensitive Data of a Consumer who is a known child.
	// 1 No Consent
	// 2 Consent
	// Fields:
	// (1) Consent to Process the Consumer’s Personal Data or Sensitive Data for Consumers from Age 13 to 16.
	// (2) Consent to Process the Consumer’s Personal Data or Sensitive Data for Consumers Younger Than 13 Years of Age.
	KnownChildSensitiveDataConsents map[int]MspaConsent
	// Consent to Collection, Use, Retention, Sale, and/or Sharing of the Consumer’s Personal Data that Is Unrelated to or Incompatible with the Purpose(s) for which the Consumer’s Personal Data Was Collected or Processed.
	// 0 Not Applicable. The Business does not use, retain, Sell, or Share the Consumer’s Personal Data for advertising purposes that are unrelated to or incompatible with the purpose(s) for which the Consumer’s Personal Data was collected or processed.
	// 1 No Consent
	// 2 Consent
	// Note: AdditionalDataProcessingConsent and PersonalDataConsents are the same
	PersonalDataConsents MspaConsent
	// Publisher or Advertiser, as applicable, is a signatory to the IAB Multistate Service Provider Agreement (MSPA), as may be amended from time to time, and declares that the transaction is a “Covered Transaction” as defined in the MSPA.
	// 1 Yes
	// 2 No
	MspaCoveredTransaction MspaNaYesNo
	// Publisher or Advertiser, as applicable, has enabled “Opt-Out Option Mode” for the “Covered Transaction,” as such terms are defined in the MSPA.
	// 0 Not Applicable.
	// 1 Yes
	// 2 No
	MspaOptOutOptionMode MspaNaYesNo
	// Publisher or Advertiser, as applicable, has enabled “Service Provider Mode” for the “Covered Transaction,” as such terms are defined in the MSPA.
	// 0 Not Applicable
	// 1 Yes
	// 2 No
	MspaServiceProviderMode MspaNaYesNo
	// Version of the MSPA. Only present in sections that use a section header
	// (e.g. US Maryland, Indiana, Kentucky, Rhode Island). As of the Fifth Amended
	// and Restated MSPA, this field should always be 0.
	MspaVersion int
	// Publisher or Advertiser, as applicable, has enabled "Opt-Out Option Mode" or "Service Provider Mode"
	// for the "Covered Transaction". Only present in sections that use a section header
	// (e.g. US Maryland, Indiana, Kentucky, Rhode Island), where it replaces the separate
	// MspaOptOutOptionMode and MspaServiceProviderMode fields, which are left unset for those sections.
	// 0 Not Applicable
	// 1 Opt-Out Option Mode
	// 2 Service Provider Mode
	MspaMode MspaMode
	// Subsections added below:
	// Global Privacy Control (GPC) is signaled and set.
	Gpc bool
}

type MspaNotice int

const (
	NoticeNotApplicable MspaNotice = iota
	NoticeProvided
	NoticeNotProvided
	InvalidNoticeValue
)

type MspaOptout int

const (
	OptOutNotApplicable MspaOptout = iota
	OptedOut
	NotOptedOut
	InvalidOptOutValue
)

type MspaConsent int

const (
	ConsentNotApplicable MspaConsent = iota
	NoConsent
	Consent
	InvalidConsentValue
)

// MspaNaYesNo represents common values for MSPA values representing
// answers, Not Applicable, Yes, No (in that order).
type MspaNaYesNo int

const (
	MspaNotApplicable MspaNaYesNo = iota
	MspaYes
	MspaNo
	InvalidMspaValue
)

// MspaMode represents which MSPA mode is enabled for a Covered Transaction in
// sections that encode a single MspaMode field instead of separate
// MspaOptOutOptionMode and MspaServiceProviderMode fields.
type MspaMode int

const (
	MspaModeNotApplicable MspaMode = iota
	MspaModeOptOutOption
	MspaModeServiceProvider
	InvalidMspaModeValue
)

// ReadMspaMode reads integers into MSPA Mode values of
// 0: Not Applicable, 1: Opt-Out Option Mode, 2: Service Provider Mode
func (r *ConsentReader) ReadMspaMode() (MspaMode, error) {
	var mm, err = r.ReadInt(2)
	return MspaMode(mm), err
}

// ReadMspaNotice reads integers into standard MSPA Notice values of
// 0: Not applicable, 1: Yes, notice was provided, 2: No, notice was not provided.
func (r *ConsentReader) ReadMspaNotice() (MspaNotice, error) {
	var mn, err = r.ReadInt(2)
	return MspaNotice(mn), err
}

// ReadMspaOptOut reads integers into standard MSPA OptOut values of
// 0: Not Applicable, 1: Opted out, 2: Did not opt out
func (r *ConsentReader) ReadMspaOptOut() (MspaOptout, error) {
	var mo, err = r.ReadInt(2)
	return MspaOptout(mo), err
}

// ReadMspaConsent reads integers into standard Consent values of
// 0: Not Applicable, 1: Not Consent, 2: Consent
func (r *ConsentReader) ReadMspaConsent() (MspaConsent, error) {
	var mc, err = r.ReadInt(2)
	return MspaConsent(mc), err
}

// ReadMspaBitfieldConsent reads n-bitfield values, and converts the values into
// MSPA Consent values.
func (r *ConsentReader) ReadMspaBitfieldConsent(l uint) (map[int]MspaConsent, error) {
	var bc, err = r.ReadNBitField(2, l)
	var consentBitfield = make(map[int]MspaConsent, len(bc))
	if err != nil {
		return nil, err
	}
	for i, b := range bc {
		consentBitfield[i] = MspaConsent(b)
	}
	return consentBitfield, err
}

// ReadMspaBitfieldOptOut reads n-bitfield values, and converts the values into
// MSPA OptOut values.
func (r *ConsentReader) ReadMspaBitfieldOptOut(l uint) (map[int]MspaOptout, error) {
	var bc, err = r.ReadNBitField(2, l)
	var consentBitfield = make(map[int]MspaOptout, len(bc))
	if err != nil {
		return nil, err
	}
	for i, b := range bc {
		consentBitfield[i] = MspaOptout(b)
	}
	return consentBitfield, err
}

// ReadMspaNaYesNo is a helper function to handle the responses to standard MSPA
// values that are in the same format of 0: Not Applicable, 1: Yes, 2: No.
func (r *ConsentReader) ReadMspaNaYesNo() (MspaNaYesNo, error) {
	var nyn, err = r.ReadInt(2)
	return MspaNaYesNo(nyn), err
}
