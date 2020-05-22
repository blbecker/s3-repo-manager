package rpmpackage

type Package struct {
	Text    string `xml:",chardata"`
	Pkgid   string `xml:"pkgid,attr"`
	Name    string `xml:"name,attr"`
	Arch    string `xml:"arch,attr"`
	Version Version  `xml:"version"`
	File string `xml:"file"`
}

type Version struct {
	Text  string `xml:",chardata"`
	Epoch string `xml:"epoch,attr"`
	Ver   string `xml:"ver,attr"`
	Rel   string `xml:"rel,attr"`
}