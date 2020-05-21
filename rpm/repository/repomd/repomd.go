package repomd

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Repomd struct {
	XMLName xml.Name `xml:"repomd"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Data    []Data   `xml:"data"`
}

type Data struct {
	Text         string       `xml:",chardata"`
	Type         string       `xml:"type,attr"`
	Location     Location     `xml:"location"`
	Checksum     RepoChecksum `xml:"checksum"`
	Timestamp    string       `xml:"timestamp"`
	OpenChecksum OpenChecksum `xml:"open-checksum"`
}

type Location struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
}

type RepoChecksum struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

type OpenChecksum struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

func FromFile(path string) (*Repomd, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	metadataBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var repomd Repomd

	err = xml.Unmarshal(metadataBytes, &repomd)
	if err != nil {
		return nil, err
	}

	return &repomd, nil
}
