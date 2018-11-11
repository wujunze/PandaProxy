package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func clearConfigFile() {
	os.Remove(configPath)
}

func TestConfigSaveConfig(t *testing.T) {
	clearConfigFile()
	config := Config{
		ListenAddr: ":7474",
		RemoteAddr: "wujunze.com:45234",
		Password:   "cGfU96a9o/sPM0bJ6xRPJiQCOSmfDF9UYNsAKhsxgfK2HWS1oAvQDXM10fVX5oz8g0Rr5WP2WkMGhupKXS83qDrTVSwJU2kwUT6Z7o90w5SJfpxJFnI9+oCwcZYY2QiIf1YixPMSEd2bs4XHk2IERdfc8RcofF6Hu6+Y6LoB1eyxeJ40dXsaJ1ClCkKEwZWRt8jvBYK0THoDqS4j6eQlYRDYB2w/59btHhydolt5ofBlwFkhOyCOH6zLjeOLxUCnd2hYR63KikGuvOFO3g4TS88rMpBNxjbgzH2Xbs4VSGqqwv8t0s1S+TiSvvQ8pKttXLni32a/2v12mrL+uPhvGQ==",
	}
	config.SaveConfig()

	file, err := os.Open(configPath)
	if err != nil {
		t.Errorf("打开配置文件 %s 出错:%s", configPath, err)
	}
	defer file.Close()

	tmp := make(map[string]string)
	err = json.NewDecoder(file).Decode(&tmp)
	if err != nil {
		t.Error(err)
	}

	if tmp["listen"] != ":7474" || tmp["password"] != "cGfU96a9o/sPM0bJ6xRPJiQCOSmfDF9UYNsAKhsxgfK2HWS1oAvQDXM10fVX5oz8g0Rr5WP2WkMGhupKXS83qDrTVSwJU2kwUT6Z7o90w5SJfpxJFnI9+oCwcZYY2QiIf1YixPMSEd2bs4XHk2IERdfc8RcofF6Hu6+Y6LoB1eyxeJ40dXsaJ1ClCkKEwZWRt8jvBYK0THoDqS4j6eQlYRDYB2w/59btHhydolt5ofBlwFkhOyCOH6zLjeOLxUCnd2hYR63KikGuvOFO3g4TS88rMpBNxjbgzH2Xbs4VSGqqwv8t0s1S+TiSvvQ8pKttXLni32a/2v12mrL+uPhvGQ==" || tmp["remote"] != "wujunze.com:45234" {
		t.Error("数据保存后不一致")
	}
}

func TestConfigReadConfig(t *testing.T) {
	clearConfigFile()
	jsonConfig := []byte(`
{
	"listen": ":7474",
	"remote": "wujunze.com:45234",
	"password": "cGfU96a9o/sPM0bJ6xRPJiQCOSmfDF9UYNsAKhsxgfK2HWS1oAvQDXM10fVX5oz8g0Rr5WP2WkMGhupKXS83qDrTVSwJU2kwUT6Z7o90w5SJfpxJFnI9+oCwcZYY2QiIf1YixPMSEd2bs4XHk2IERdfc8RcofF6Hu6+Y6LoB1eyxeJ40dXsaJ1ClCkKEwZWRt8jvBYK0THoDqS4j6eQlYRDYB2w/59btHhydolt5ofBlwFkhOyCOH6zLjeOLxUCnd2hYR63KikGuvOFO3g4TS88rMpBNxjbgzH2Xbs4VSGqqwv8t0s1S+TiSvvQ8pKttXLni32a/2v12mrL+uPhvGQ=="
}`)

	err := ioutil.WriteFile(configPath, jsonConfig, 0644)
	if err != nil {
		t.Error(err)
	}

	config := Config{}
	config.ReadConfig()

	if config.ListenAddr != ":7474" || config.RemoteAddr != "wujunze.com:45234" || config.Password != "cGfU96a9o/sPM0bJ6xRPJiQCOSmfDF9UYNsAKhsxgfK2HWS1oAvQDXM10fVX5oz8g0Rr5WP2WkMGhupKXS83qDrTVSwJU2kwUT6Z7o90w5SJfpxJFnI9+oCwcZYY2QiIf1YixPMSEd2bs4XHk2IERdfc8RcofF6Hu6+Y6LoB1eyxeJ40dXsaJ1ClCkKEwZWRt8jvBYK0THoDqS4j6eQlYRDYB2w/59btHhydolt5ofBlwFkhOyCOH6zLjeOLxUCnd2hYR63KikGuvOFO3g4TS88rMpBNxjbgzH2Xbs4VSGqqwv8t0s1S+TiSvvQ8pKttXLni32a/2v12mrL+uPhvGQ==" {
		t.Error("读取的配置数据不一致")
	}

}
