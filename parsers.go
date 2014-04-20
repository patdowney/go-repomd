package repomd

import (
	"encoding/xml"
	"io"
)

func parse(target interface{}, body io.Reader) error {
	xmlDecoder := xml.NewDecoder(body)
	err := xmlDecoder.Decode(target)

	if err != nil {
		return err
	}

	return nil
}

func ParseMetadata(body io.Reader) (*Metadata, error) {
	metadata := &Metadata{}

	err := parse(metadata, body)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

func ParseRepomd(body io.Reader) (*Repomd, error) {
	repomd := &Repomd{}

	err := parse(repomd, body)
	if err != nil {
		return nil, err
	}
	return repomd, nil
}
