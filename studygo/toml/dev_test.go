package conf

import "testing"

func TestReadConf(t *testing.T) {
	p, err := ReadConf("./my.conf")
	if err != nil {
		t.Logf("%v", err)
	}

	t.Logf("person %v", p.Books)
	t.Logf("person.friend %v", p.Friend)
}