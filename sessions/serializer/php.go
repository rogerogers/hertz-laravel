package serializer

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/elliotchance/phpserialize"
	"github.com/gorilla/sessions"
)

type PHPSerializer struct{}

func (s PHPSerializer) Serialize(ss *sessions.Session) ([]byte, error) {
	m := make(map[string]interface{}, len(ss.Values))
	for k, v := range ss.Values {
		ks, ok := k.(string)
		if !ok {
			err := fmt.Errorf("non-string key value, cannot serialize session to JSON: %v", k)
			hlog.Errorf("PHPSerializer.serialize() Error: %v", err)
			return nil, err
		}
		m[ks] = v
	}
	return phpserialize.Marshal(m, nil)
}

type mapStrInterface map[interface{}]interface{}

func (s PHPSerializer) Deserialize(d []byte, ms *mapStrInterface) error {
	m := make(map[interface{}]interface{})
	err := phpserialize.Unmarshal(d, &m)
	if err != nil {
		hlog.Errorf("serializer.PHPSerializer.deserialize() Error: %v", err)
		return err
	}
	ss := *ms
	for k, v := range m {
		ss[k] = v
	}
	return nil
}
