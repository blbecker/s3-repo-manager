package filelists

import (
	"compress/gzip"
	"encoding/xml"
	"io/ioutil"
	"os"

	"github.com/blbecker/s3-repo-manager/rpm/repository/rpmpackage"
)

type Filelists struct {
	XMLName  xml.Name `xml:"filelists"`
	Text     string   `xml:",chardata"`
	Xmlns    string   `xml:"xmlns,attr"`
	Packages string   `xml:"packages,attr"`
	Package  rpmpackage.Package `xml:"rpmpackage"`
}

func FromFile(path string) (*Filelists, error) {
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

	var filelists Filelists
	err = xml.Unmarshal(metadataBytes, &filelists)
	if err != nil {
		return nil, err
	}

	return &filelists, nil
}