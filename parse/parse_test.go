package parse

import (
	"reflect"
	"testing"
)

func TestYamlParse(t *testing.T) {
	str := "k1:v1\rk2:v2\rk3:1"
	yaml := YamlParse{}
	m, err := yaml.Parse([]byte(str))
	if err != nil {
		t.Fatalf("parse yaml error!")
	}
	result := map[string]interface{}{
		"k1": "v1",
		"k2": "v2",
		"k3": "1",
	}
	if !reflect.DeepEqual(m, result) {
		t.Fatalf("parse data is not correct")
	}
}
