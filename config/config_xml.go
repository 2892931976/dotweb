package config

import (
	"encoding/xml"
)

// UnmarshalXml parses the XML-encoded data and stores the result in
// the value pointed to by v, which must be an arbitrary struct,
// slice, or string. Well-formed data that does not fit into v is
// discarded.
func UnmarshalXml(content []byte, v interface{}) error {
	return xml.Unmarshal(content, v)
}

// MarshalXml returns the XML encoding of v.
func MarshalXml(v interface{}) (out []byte, err error) {
	return xml.Marshal(v)
}

// MarshalXmlString returns the XML encoding string format of v.
func MarshalXmlString(v interface{}) (out string) {
	marshal, err := xml.Marshal(v)
	if err != nil {
		return ""
	}
	return string(marshal)
}
