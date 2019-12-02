package protobuf

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	"github.com/emicklei/proto"
	"github.com/pkg/errors"
)

// Validate will ensure the bytes passed are a valid gzipped proto schema
func Validate(schema []byte) error {
	r, err := gzip.NewReader(bytes.NewBuffer(schema))
	if err != nil {
		return errors.Wrap(err, "invalid gzip: failed to create gzip reader")
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.Wrap(err, "failed to read from gzip buffer")
	}

	p := proto.NewParser(bytes.NewReader(b))
	pp, err := p.Parse()
	if err != nil {
		return errors.Wrap(err, "invalid proto")
	}

	hasRecord := false
	for _, e := range pp.Elements {
		switch v := e.(type) {
		case *proto.Message:
			if v.Name == "record" {
				hasRecord = true
				goto breakOut
			}
		}
	}

breakOut:
	if !hasRecord {
		return errors.Errorf("proto has no record message")
	}

	return nil
}
