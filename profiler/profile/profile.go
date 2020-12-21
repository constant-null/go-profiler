package profile

import (
	"compress/gzip"
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// Parse parses profile file
func Parse(r io.Reader) (*Profile, error) {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(gz)
	if err != nil {
		return nil, err
	}

	var pb Profile
	if err = proto.Unmarshal(data, &pb); err != nil {
		return nil, err
	}

	return &pb, nil
}