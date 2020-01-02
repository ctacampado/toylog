package toylog

import (
	"log"
	"os/exec"
	"runtime"
	"testing"
)

func TestNewToyLoggerNIL(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file bool
	}{
		name: "logger",
		lvl:  DEBUG,
		file: false,
	}

	_, err := NewToyLog(in.name, in.lvl, 12345)
	if nil != err {
		t.Log(err.Error())
	}
}

func TestNewToyLogToFile(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file bool
	}{
		name: "logger",
		lvl:  DEBUG,
		file: true,
	}

	l, err := NewToyLog(in.name, in.lvl, in.file)
	if nil != err {
		t.Error(err.Error())
	}

	// delete created log file
	if runtime.GOOS == "linux" {
		cmd := exec.Command("rm", l.FileName)
		err = cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	}

}

func TestInfo(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file bool
		in   string
	}{
		name: "",
		lvl:  INFO,
		file: false,
		in:   "hello world",
	}
	logger, err := NewToyLog(in.name, in.lvl, in.file)
	if nil != err {
		t.Error(err.Error())
	}
	logger.Info(in.in)
	logger.Debug(in.in)
}

func TestDebug(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file bool
		in   string
	}{
		name: "logger",
		lvl:  DEBUG,
		file: false,
		in:   "hello world",
	}

	logger, err := NewToyLog(in.name, in.lvl, in.file)
	if nil != err {
		t.Error(err.Error())
	}
	logger.Debug(in.in)
}

func TestError(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file bool
		in   string
	}{
		name: "logger",
		lvl:  ERR,
		file: false,
		in:   "hello world",
	}

	logger, err := NewToyLog(in.name, in.lvl, in.file)
	if nil != err {
		t.Error(err.Error())
	}
	logger.Error(in.in)
}
