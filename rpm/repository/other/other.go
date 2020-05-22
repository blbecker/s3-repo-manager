package other

import "encoding/xml"

type Otherdata struct {
	XMLName  xml.Name `xml:"otherdata"`
	Text     string   `xml:",chardata"`
	Xmlns    string   `xml:"xmlns,attr"`
	Packages string   `xml:"packages,attr"`
	Package  struct {
		Text    string `xml:",chardata"`
		Arch    string `xml:"arch,attr"`
		Name    string `xml:"name,attr"`
		Pkgid   string `xml:"pkgid,attr"`
		Version struct {
			Text  string `xml:",chardata"`
			Epoch string `xml:"epoch,attr"`
			Rel   string `xml:"rel,attr"`
			Ver   string `xml:"ver,attr"`
		} `xml:"version"`
	} `xml:"rpmpackage"`
}
