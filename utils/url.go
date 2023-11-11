package utils

import (
	"encoding/base64"
	"net/url"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func B64SafeDecode(encoded []byte) (decoded []byte, err error) {
	escape, err := url.QueryUnescape(string(encoded))
	if err != nil {
		hlog.Error(err, 1)
		return
	}
	decoded, err = base64.StdEncoding.DecodeString(escape)
	if err != nil {
		hlog.Error(err, 2)
		return nil, err
	}
	return
}

func B64Decode(encoded string) (decoded []byte, err error) {
	decoded, err = base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
