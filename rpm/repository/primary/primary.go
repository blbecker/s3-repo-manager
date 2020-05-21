package primary

import (
	"compress/gzip"
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Metadata struct {
	XMLName  xml.Name `xml:"metadata"`
	Text     string   `xml:",chardata"`
	Xmlns    string   `xml:"xmlns,attr"`
	Rpm      string   `xml:"rpm,attr"`
	Packages string   `xml:"packages,attr"`
	Package  Package  `xml:"package"`
}

type Package struct {
	Text        string          `xml:",chardata"`
	Type        string          `xml:"type,attr"`
	Name        string          `xml:"name"`
	Arch        string          `xml:"arch"`
	Version     Version         `xml:"version"`
	Checksum    PackageChecksum `xml:"checksum"`
	Summary     string          `xml:"summary"`
	Description string          `xml:"description"`
	Packager    string          `xml:"packager"`
	URL         string          `xml:"url"`
	Time        Time            `xml:"time"`
	Size        Size            `xml:"size"`
	Location    PackageLocation `xml:"location"`
	Format      Format          `xml:"format"`
}

type Version struct {
	Text  string `xml:",chardata"`
	Epoch string `xml:"epoch,attr"`
	Rel   string `xml:"rel,attr"`
	Ver   string `xml:"ver,attr"`
}

type PackageChecksum struct {
	Text  string `xml:",chardata"`
	Pkgid string `xml:"pkgid,attr"`
	Type  string `xml:"type,attr"`
}

type Time struct {
	Text  string `xml:",chardata"`
	Build string `xml:"build,attr"`
	File  string `xml:"file,attr"`
}

type Size struct {
	Text      string `xml:",chardata"`
	Archive   string `xml:"archive,attr"`
	Installed string `xml:"installed,attr"`
	Package   string `xml:"package,attr"`
}

type PackageLocation struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
}

type Format struct {
	Text        string      `xml:",chardata"`
	License     string      `xml:"license"`
	Vendor      string      `xml:"vendor"`
	Group       string      `xml:"group"`
	Buildhost   string      `xml:"buildhost"`
	Sourcerpm   string      `xml:"sourcerpm"`
	HeaderRange HeaderRange `xml:"header-range"`
	Provides    Provides    `xml:"provides"`
	Requires    Require     `xml:"requires"`
	File        string      `xml:"file"`
}

type HeaderRange struct {
	Text  string `xml:",chardata"`
	End   string `xml:"end,attr"`
	Start string `xml:"start,attr"`
}

type Provides struct {
	Text  string  `xml:",chardata"`
	Entry []Entry `xml:"entry"`
}

type Entry struct {
	Text  string `xml:",chardata"`
	Epoch string `xml:"epoch,attr"`
	Flags string `xml:"flags,attr"`
	Name  string `xml:"name,attr"`
	Rel   string `xml:"rel,attr"`
	Ver   string `xml:"ver,attr"`
}

type Require struct {
	Text  string  `xml:",chardata"`
	Entry []Entry `xml:"entry"`
}

func FromFile(path string) (*Metadata, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	defer archive.Close()
	metadataBytes, err := ioutil.ReadAll(archive)
	if err != nil {
		return nil, err
	}

	var metadata Metadata
	err = xml.Unmarshal(metadataBytes, &metadata)
	if err != nil {
		return nil, err
	}

	return &metadata, nil
}
