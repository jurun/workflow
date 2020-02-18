package workflow

import (
	"github.com/jurun/workflow/form/layout"
	"log"
	"testing"
)

var wfk *Workflow

func init() {
	var err error
	wfk, err = NewWorkflow(
		Mysql(
			"127.0.0.1:3307",
			"root",
			"111111",
			"xunray_bpm",
		),
		MongoDB(
			"127.0.0.1:27017",
			"bpm",
		),
		Debug(),
	)

	if err != nil {
		log.Panic(err)
	}
}

func TestFormCreate(t *testing.T) {
	attr, err := wfk.Form.Create(1, 1, "测试表按钮", true)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", attr)
}

func TestFormDetail(t *testing.T) {
	attr, _, err := wfk.Form.Get(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", attr)
}

func TestFormDelete(t *testing.T) {
	err := wfk.Form.Delete(1)
	if err != nil {
		t.Error(err)
	}
}

func TestFormUpdate(t *testing.T) {
	attr, _, err := wfk.Form.Get(1)
	if err != nil {
		t.Error(err)
		return
	}

	attr.Name = "xxxxx"
	attr.Layout = layout.Layout{}.Init()
	err = wfk.Form.Update(attr)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestProcessCreate(t *testing.T) {
	process, err := wfk.Process.Create(1)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", process)
}

func TestProcessGet(t *testing.T) {
	process, found, err := wfk.Process.Get(62)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Found: ", found)
	t.Logf("%+v\n", process)
}

func TestProcessList(t *testing.T) {
	ps, err := wfk.Process.List(1, false)
	if err != nil {
		t.Error(err)
		return
	}

	for _, p := range ps {
		t.Logf("%+v\n", p)
	}
}

func TestProcessEnable(t *testing.T) {
	if err := wfk.Process.Enable(44); err != nil {
		t.Error(err)
		return
	}
}

func TestProcessDelete(t *testing.T) {
	if err := wfk.Process.Delete(62); err != nil {
		t.Error(err)
		return
	}
}
