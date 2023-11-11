package serializer

import (
	"fmt"
	"testing"

	"github.com/elliotchance/phpserialize"
)

func TestPHPSerialize(t *testing.T) {
	s := PHPSerializer{}
	ms := &mapStrInterface{}

	ses := "s:255:\"a:4:{s:6:\"_token\";s:40:\"alypQqmmeqHFv2pYQ0uu9sftg0C241BQdviREO4T\";s:3:\"url\";a:1:{s:8:\"intended\";s:31:\"http://127.0.0.1:8000/dashboard\";}s:9:\"_previous\";a:1:{s:3:\"url\";s:27:\"http://127.0.0.1:8000/login\";}s:6:\"_flash\";a:2:{s:3:\"old\";a:0:{}s:3:\"new\";a:0:{}}}\";"
	var a []byte
	phpserialize.Unmarshal([]byte(ses), &a)

	s.Deserialize(a, ms)
	fmt.Println(ms)
}

func TestPHPDeserialize(t *testing.T) {
}
