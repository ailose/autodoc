package autodoc

import (
	"testing"
)

func TestSetLang(t *testing.T) {
	SetLang("chinese")

}

func TestGenDoc(t *testing.T) {
	sec := make(map[string]DocSection)
	sec["test"] = DocSection{
		Index: 0,
		Name:  "test",
	}
	var infos []*DocInfo
	info := DocInfo{
		Name: "test",
		Desc: "test des",
	}
	infos = append(infos, &info)
	GenDoc("test.md", "test", sec, infos)
	// _, err := GetNewPassword(pwd)
	// if err != nil {
	// 	t.Fatal(err)
	// }
}

func TestParseModel(t *testing.T) {
	type d struct {
		Test1 string `json:"t1" note:"test1"`
		Test2 int    `json:"t2" note:"test2"`
		Test3 bool   `json:"t3" note:"test3"`
	}
	var t1 d
	r := ParseModel(t1)
	if r == nil {
		t.Error(r)
	}
}

func TestParseStruct(t *testing.T) {
	type d struct {
		Test1 string `json:"t1" note:"test1"`
		Test2 int    `json:"t2" note:"test2"`
		Test3 bool   `json:"t3" note:"test3"`
	}
	var t1 d
	res := make(map[string]X)
	ParseStruct(t1, res)
	if res == nil {
		t.Error(res)
	}
}
