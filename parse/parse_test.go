package parse

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestYamlParse(t *testing.T) {
	b, err := ioutil.ReadFile("./yaml_test.yaml")
	if err != nil {
		t.Errorf("error read: %v", err.Error())
	}
	yaml := YamlParse{}
	m, err := yaml.Parse(b)
	if err != nil {
		t.Fatalf("parse yaml error! %v", err.Error())
	}
	result := map[string]interface{}{
		"k1": map[string]interface{}{
			"k2": "v2",
			"k3": "v3",
		},
	}
	if !reflect.DeepEqual(m, result) {
		t.Fatalf("parse data is not correct")
	}
}
