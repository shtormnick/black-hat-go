package metadata

import (
	"encoding/xml"
	"strings"
)

type OfficeCoreProperty struct {
	XMLName        xml.Name `xml:"CoreProperties"`
	Creator        string   `xml:"creator"`
	LastModifiedBy string   `xml:"lastModifiedBy"`
}

type OfficeAppProperty struct {
	XMLName     xml.Name `xml:"Properties"`
	Application string   `xml:"Application"`
	Company     string   `xml:"Company"`
	Version     string   `xml:"AppVersion"`
}

var OfficeVersions = map[string]string{
	"16": "2016",
	"15": "2013",
	"14": "2010",
	"12": "2007",
	"11": "2003",
}

func (a *OfficeAppProperty) GetMajorVersion() string {
	tokens := strings.Split(a.Version, ".")

	if len(tokens) < 2 {
		return "Unknown"
	}
	v, ok := OfficeVersions[tokens[0]]
	if !ok {
		return "Unknown"
	}
	return v
}
