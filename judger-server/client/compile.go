package client

import (
	"errors"
	"os"
	"os/exec"

	"strconv"
)

var ErrorCompile = errors.New("compile error")

const DOJToken = "DLNU-DOJ"

var (
	runPath  = "/tmp"
	dataPath = "/tmp"
)

type DLNUJudger struct {
	Token   string
	Workdir string
	Time    int
	Memory  int
}

type Params struct {
	StudentID int64
	ProblemID int64
	Language  string
	Code      string
	Time      int
	Memory    int
}

func (d *DLNUJudger) files(p *Params, workdir string) error {
	var codeFileName string

	switch p.Language {
	case "cpp":
		codeFileName = workdir + "/Main.cpp"
	case "java":
		codeFileName = workdir + "/Main.java"
	}

	codefile, err := os.Create(codeFileName)
	defer codefile.Close()

	if err != nil {
		return err
	}

	_, err = codefile.WriteString(p.Code)

	return err
}

// 穿件提交题目的文件夹
func (d *DLNUJudger) Init(p *Params) error {
	d.Token = DOJToken
	d.Workdir = runPath + "/" + strconv.Itoa(int(p.StudentID)) + "/" + strconv.Itoa(int(p.ProblemID))

	cmd := exec.Command("mkdir", "-p", d.Workdir)
	cmd.Run()

	err := d.files(p, d.Workdir)
	return err
}

func (d *DLNUJudger) Match(token string) bool {
	if DOJToken == token || token == "" {
		return true
	}
	return false
}
