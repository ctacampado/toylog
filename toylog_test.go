package toylog

import (
	"log"
	"os/exec"
	"runtime"
	"testing"
)

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

	l := NewToyLog(in.name, in.lvl, in.file)

	// delete created log file
	if runtime.GOOS == "linux" {
		cmd := exec.Command("rm", l.FileName)
		err := cmd.Run()
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

	logger := NewToyLog(in.name, in.lvl, in.file)
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

	logger := NewToyLog(in.name, in.lvl, in.file)
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

	logger := NewToyLog(in.name, in.lvl, in.file)
	logger.Error(in.in)
}
