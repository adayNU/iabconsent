package iabconsent_test

import (
	"strconv"
	"strings"

	"github.com/go-check/check"

	"github.com/openx/iabconsent"
)

type MspaSuite struct{}

var _ = check.Suite(&MspaSuite{})

func (s *MspaSuite) TestReadMspaNotice(c *check.C) {
	var r = iabconsent.NewConsentReader([]byte{0b00011011})
	var mns = []iabconsent.MspaNotice{
		iabconsent.NoticeNotApplicable,
		iabconsent.NoticeProvided,
		iabconsent.NoticeNotProvided,
		iabconsent.InvalidNoticeValue,
	}
	for _, i := range mns {
		var mn, err = r.ReadMspaNotice()
		c.Check(err, check.IsNil)
		c.Check(mn, check.Equals, i)
	}
}

func (s *MspaSuite) TestReadMspaOptOut(c *check.C) {
	var r = iabconsent.NewConsentReader([]byte{0b00011011})
	var mos = []iabconsent.MspaOptout{
		iabconsent.OptOutNotApplicable,
		iabconsent.OptedOut,
		iabconsent.NotOptedOut,
		iabconsent.InvalidOptOutValue,
	}
	for _, i := range mos {
		var mo, err = r.ReadMspaOptOut()
		c.Check(err, check.IsNil)
		c.Check(mo, check.Equals, i)
	}
}

func (s *MspaSuite) TestReadMspaConsent(c *check.C) {
	var r = iabconsent.NewConsentReader([]byte{0b00011011})
	var mcs = []iabconsent.MspaConsent{
		iabconsent.ConsentNotApplicable,
		iabconsent.NoConsent,
		iabconsent.Consent,
		iabconsent.InvalidConsentValue,
	}
	for _, i := range mcs {
		var mc, err = r.ReadMspaConsent()
		c.Check(err, check.IsNil)
		c.Check(mc, check.Equals, i)
	}
}

func (s *MspaSuite) TestReadMspaBitfieldConsent(c *check.C) {
	var tcs = []struct {
		testBytes      []byte
		bitfieldLength uint
		expected       map[int]iabconsent.MspaConsent
	}{
		{testBytes: []byte{0b00000000},
			bitfieldLength: 1,
			expected:       map[int]iabconsent.MspaConsent{0: iabconsent.ConsentNotApplicable}},
		{testBytes: []byte{0b00000000},
			bitfieldLength: 4,
			expected: map[int]iabconsent.MspaConsent{
				0: iabconsent.ConsentNotApplicable,
				1: iabconsent.ConsentNotApplicable,
				2: iabconsent.ConsentNotApplicable,
				3: iabconsent.ConsentNotApplicable}},
		{testBytes: []byte{0b00011011},
			bitfieldLength: 4,
			expected: map[int]iabconsent.MspaConsent{
				0: iabconsent.ConsentNotApplicable,
				1: iabconsent.NoConsent,
				2: iabconsent.Consent,
				3: iabconsent.InvalidConsentValue}},
	}

	for _, t := range tcs {
		var r = iabconsent.NewConsentReader(t.testBytes)
		var bc, err = r.ReadMspaBitfieldConsent(t.bitfieldLength)
		c.Check(err, check.IsNil)
		c.Check(bc, check.DeepEquals, t.expected)
	}
}

func (s *MspaSuite) TestReadMspaNaYesNo(c *check.C) {
	var r = iabconsent.NewConsentReader([]byte{0b00011011})
	var mcs = []iabconsent.MspaNaYesNo{
		iabconsent.MspaNotApplicable,
		iabconsent.MspaYes,
		iabconsent.MspaNo,
		iabconsent.InvalidMspaValue,
	}
	for _, i := range mcs {
		var mc, err = r.ReadMspaNaYesNo()
		c.Check(err, check.IsNil)
		c.Check(mc, check.Equals, i)
	}
}

func (s *MspaSuite) TestParseMSPA(c *check.C) {
	for sid, sections := range mspaConsentFixtures {
		for section, result := range sections {

			c.Log(section)

			var gppSection = iabconsent.NewMspa(sid, section)
			var p, err = gppSection.ParseConsent()

			c.Check(err, check.IsNil)
			c.Check(p, check.DeepEquals, result)
		}
	}
}

func (s *MspaSuite) TestParseMSPAError(c *check.C) {
	var mspaTests = []struct {
		desc string
		key  string
		sid  int
	}{
		{
			desc: "US National",
			key:  "usnat",
			sid:  iabconsent.UsNationalSID,
		},
		{
			desc: "California",
			key:  "usca",
			sid:  iabconsent.UsCaliforniaSID,
		},
		{
			desc: "Virginia",
			key:  "usva",
			sid:  iabconsent.UsVirginiaSID,
		},
		{
			desc: "Colorado",
			key:  "usco",
			sid:  iabconsent.UsColoradoSID,
		},
		{
			desc: "Utan",
			key:  "usut",
			sid:  iabconsent.UsUtahSID,
		},
		{
			desc: "Connecticut",
			key:  "usct",
			sid:  iabconsent.UsConnecticutSID,
		},
		{
			desc: "Florida",
			key:  "usfl",
			sid:  iabconsent.UsFloridaSID,
		},
		{
			desc: "Montana",
			key:  "usmt",
			sid:  iabconsent.UsMontanaSID,
		},
		{
			desc: "Oregon",
			key:  "usor",
			sid:  iabconsent.UsOregonSID,
		},
		{
			desc: "Texas",
			key:  "ustx",
			sid:  iabconsent.UsTexasSID,
		},
		{
			desc: "Delaware",
			key:  "usde",
			sid:  iabconsent.UsDelawareSID,
		},
		{
			desc: "Iowa",
			key:  "usia",
			sid:  iabconsent.UsIowaSID,
		},
		{
			desc: "Nebraska",
			key:  "usne",
			sid:  iabconsent.UsNebraskaSID,
		},
		{
			desc: "New Hampshire",
			key:  "usnh",
			sid:  iabconsent.UsNewHampshireSID,
		},
		{
			desc: "New Jersey",
			key:  "usnj",
			sid:  iabconsent.UsNewJerseySID,
		},
		{
			desc: "Tennessee",
			key:  "ustn",
			sid:  iabconsent.UsTennesseeSID,
		},
		{
			desc: "Minnesota",
			key:  "usmn",
			sid:  iabconsent.UsMinnesotaSID,
		},
	}
	var tcs = []struct {
		desc          string
		consentString string
		expected      string
	}{
		{
			desc:          "Wrong Version.",
			consentString: "DVVqAAEABA",
			expected:      "unsupported version: 3",
		},
		{
			desc:          "Bad Decoding.",
			consentString: "$%&*(",
			expected:      "parse %s consent string: illegal base64 data at input byte 0",
		},
		{
			desc:          "Invalid Section Length. Too long for all v1 strings",
			consentString: "BqqAqqqqqqqqAA",
			expected:      "invalid consent string length for v1",
		},
		{
			desc:          "Invalid Section Length. Too short for all v1 strings",
			consentString: "BqqA",
			expected:      "invalid consent string length for v1",
		},
	}
	for _, s := range mspaTests {
		for _, t := range tcs {
			c.Log(s.desc + " - " + t.desc)

			var gppSection = iabconsent.NewMspa(s.sid, t.consentString)
			var p, err = gppSection.ParseConsent()

			c.Check(p, check.IsNil)
			var expected = strings.Replace(t.expected, "%s", s.key, 1)
			c.Check(err, check.ErrorMatches, expected)
		}
	}
}

func (s *MspaSuite) TestParseMSPASectionWithHeaderError(c *check.C) {
	var mspaTests = []struct {
		desc   string
		key    string
		sid    int
		header string
	}{
		{desc: "Maryland", key: "usmd", sid: iabconsent.UsMarylandSID, header: "YBAA"},
		{desc: "Indiana", key: "usin", sid: iabconsent.UsIndianaSID, header: "ZBAA"},
		{desc: "Kentucky", key: "usky", sid: iabconsent.UsKentuckySID, header: "aBAA"},
		{desc: "Rhode Island", key: "usri", sid: iabconsent.UsRhodeIslandSID, header: "bBAA"},
	}
	var tcs = []struct {
		desc          string
		consentString string
		// mutateHeader, if set, modifies the section's valid header before it replaces %h in consentString.
		mutateHeader func(header string) string
		expected     string
	}{
		{
			desc:          "Missing core subsection.",
			consentString: "%h",
			expected:      "parse %s consent string: missing core subsection",
		},
		{
			desc:          "Wrong section id.",
			consentString: "HBAA.AQAA",
			expected:      "parse %s section header: section id mismatch: expected %d, got 7",
		},
		{
			desc:          "Wrong Version.",
			consentString: "%h.AQAA",
			// Version 3 instead of 1.
			mutateHeader: func(h string) string { return h[:1] + "D" + h[2:] },
			expected:     "parse %s consent string: unsupported version: 3",
		},
		{
			desc:          "Bad header decoding.",
			consentString: "$%&*.AQAA",
			expected:      "parse %s section header: decode section header: illegal base64 data at input byte 0",
		},
		{
			desc:          "Bad core decoding.",
			consentString: "%h.$%&*",
			expected:      "parse %s consent string: illegal base64 data at input byte 0",
		},
		{
			desc:          "Core too short.",
			consentString: "%h.AQ",
			expected:      "parse %s consent string: invalid consent string length for v1",
		},
		{
			desc:          "Subsection listed in header, but missing.",
			consentString: "%hBYA.AQAA",
			// Header with a single subsection: drop the empty subsections range.
			mutateHeader: func(h string) string { return h[:3] },
			expected:     "parse %s section header: mismatch number of subsections: header lists 1, found 0",
		},
		{
			desc:          "Segment not listed in header.",
			consentString: "%h.AQAA.AQAA",
			expected:      "parse %s section header: mismatch number of subsections: header lists 0, found 1",
		},
	}
	for _, m := range mspaTests {
		for _, t := range tcs {
			c.Log(m.desc + " - " + t.desc)

			var header = m.header
			if t.mutateHeader != nil {
				header = t.mutateHeader(header)
			}
			var consentString = strings.Replace(t.consentString, "%h", header, 1)
			var gppSection = iabconsent.NewMspa(m.sid, consentString)
			var p, err = gppSection.ParseConsent()

			c.Check(p, check.IsNil)
			var expected = strings.Replace(t.expected, "%s", m.key, 1)
			expected = strings.Replace(expected, "%d", strconv.Itoa(m.sid), 1)
			c.Check(err, check.ErrorMatches, expected)
		}
	}
}

// Segments must have the exact length of their fields plus padding, like the other US state sections.
func (s *MspaSuite) TestParseMSPASectionWithHeaderTooLong(c *check.C) {
	var tcs = []struct {
		desc          string
		sid           int
		consentString string
		expected      string
		// partial is whether a partially parsed consent is returned with the error.
		partial bool
	}{
		{
			desc:          "Maryland core too long.",
			sid:           iabconsent.UsMarylandSID,
			consentString: "YBAA.AQAAAA",
			expected:      "parse usmd consent string: invalid consent string length for v1",
		},
		{
			desc:          "Indiana core too long.",
			sid:           iabconsent.UsIndianaSID,
			consentString: "ZBAA.AQAAAAAA",
			expected:      "parse usin consent string: invalid consent string length for v1",
		},
		{
			desc:          "Kentucky sensitive data consents subsection too long.",
			sid:           iabconsent.UsKentuckySID,
			consentString: "aBABYA.BVVV.AAAAAA",
			expected:      "parse usky subsection 1: invalid sensitive data consents subsection length",
			partial:       true,
		},
	}
	for _, t := range tcs {
		c.Log(t.desc)

		var p, err = iabconsent.NewMspa(t.sid, t.consentString).ParseConsent()
		c.Check(err, check.ErrorMatches, t.expected)
		if t.partial {
			c.Check(p, check.NotNil)
		} else {
			c.Check(p, check.IsNil)
		}
	}
}

func (s *MspaSuite) TestParseMSPASectionWithHeaderSubsectionErrors(c *check.C) {
	var tcs = []struct {
		desc          string
		sid           int
		consentString string
		expected      string
	}{
		{
			desc:          "Indiana sensitive data consents subsection too short.",
			sid:           iabconsent.UsIndianaSID,
			consentString: "ZBABYA.BVVV.kk",
			expected:      "parse usin subsection 1: invalid sensitive data consents subsection length",
		},
		{
			desc:          "Kentucky sensitive data consents subsection bad decoding.",
			sid:           iabconsent.UsKentuckySID,
			consentString: "aBABYA.BVVV.$%&",
			expected:      "parse usky subsection 1: parse sensitive data consents subsection: illegal base64 data at input byte 0",
		},
		{
			desc:          "Maryland GPC subsection bad decoding.",
			sid:           iabconsent.UsMarylandSID,
			consentString: "YBABYA.BVVU.$%",
			expected:      "parse usmd subsection 1: parse gpp subsection string: illegal base64 data at input byte 0",
		},
	}
	for _, t := range tcs {
		c.Log(t.desc)

		var gppSection = iabconsent.NewMspa(t.sid, t.consentString)
		var p, err = gppSection.ParseConsent()

		c.Check(err, check.ErrorMatches, t.expected)
		// Like the other US state sections, the partially parsed consent (core subsection) is returned on
		// subsection errors.
		c.Assert(p, check.NotNil)
		var mspa = p.(*iabconsent.MspaParsedConsent)
		c.Check(mspa.Version, check.Equals, 1)
		c.Check(mspa.MspaCoveredTransaction, check.Equals, iabconsent.MspaYes)
	}
}

// Subsections listed in the section header without a parser for the section are skipped.
func (s *MspaSuite) TestParseMSPASectionWithHeaderUnsupportedSubsection(c *check.C) {
	// Indiana header listing subsection ID 2, which Indiana does not define.
	var p, err = iabconsent.NewMspa(iabconsent.UsIndianaSID, "ZBABMA.BVVV.$%&").ParseConsent()
	c.Assert(err, check.IsNil)
	var mspa = p.(*iabconsent.MspaParsedConsent)
	c.Check(mspa.SensitiveDataProcessingConsents, check.IsNil)
	c.Check(mspa.Gpc, check.Equals, false)
}

func (s *MspaSuite) TestParseMSPAMspaMode(c *check.C) {
	var tcs = []struct {
		consentString string
		mode          iabconsent.MspaMode
	}{
		// MspaVersion=0, MspaCoveredTransaction=1, MspaMode=0
		{"ZBAA.AQAA", iabconsent.MspaModeNotApplicable},
		// MspaVersion=0, MspaCoveredTransaction=1, MspaMode=1
		{"ZBAA.AUAA", iabconsent.MspaModeOptOutOption},
		// MspaVersion=0, MspaCoveredTransaction=1, MspaMode=2
		{"ZBAA.AYAA", iabconsent.MspaModeServiceProvider},
		// MspaVersion=0, MspaCoveredTransaction=1, MspaMode=3
		{"ZBAA.AcAA", iabconsent.InvalidMspaModeValue},
	}
	for _, t := range tcs {
		c.Log(t.consentString)

		var p, err = iabconsent.NewMspa(iabconsent.UsIndianaSID, t.consentString).ParseConsent()
		c.Assert(err, check.IsNil)
		var mspa = p.(*iabconsent.MspaParsedConsent)
		c.Check(mspa.MspaMode, check.Equals, t.mode)
		// Sections with MspaMode do not contain MspaOptOutOptionMode and MspaServiceProviderMode.
		c.Check(mspa.MspaOptOutOptionMode, check.Equals, iabconsent.MspaNotApplicable)
		c.Check(mspa.MspaServiceProviderMode, check.Equals, iabconsent.MspaNotApplicable)
	}
}
