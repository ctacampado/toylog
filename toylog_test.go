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
		name: "TestNewToyLogToFile",
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

func TestError(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file bool
		in   string
	}{
		name: "TestError",
		lvl:  ERROR,
		file: false,
		in:   "hello world",
	}

	logger := NewToyLog(in.name, in.lvl, in.file)
	logger.Error(in.in)
	logger.Warning(in.in)
	logger.Info(in.in)
	logger.Debug(in.in)
	logger.Trace(in.in)
}

func TestInfo(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file bool
		in   string
	}{
		name: "TestInfo",
		lvl:  INFO,
		file: false,
		in:   "hello world",
	}

	logger := NewToyLog(in.name, in.lvl, in.file)
	logger.Error(in.in)
	logger.Warning(in.in)
	logger.Info(in.in)
	logger.Debug(in.in)
	logger.Trace(in.in)
}

func TestDebug(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file bool
		in   string
	}{
		name: "TestDebug",
		lvl:  DEBUG,
		file: false,
		in:   "hello world",
	}

	logger := NewToyLog(in.name, in.lvl, in.file)
	logger.Error(in.in)
	logger.Warning(in.in)
	logger.Info(in.in)
	logger.Debug(in.in)
	logger.Trace(in.in)
}

func TestTrace(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file bool
		in   string
	}{
		name: "TestTrace",
		lvl:  TRACE,
		file: false,
		in:   "hello world",
	}

	logger := NewToyLog(in.name, in.lvl, in.file)
	logger.Error(in.in)
	logger.Warning(in.in)
	logger.Info(in.in)
	logger.Debug(in.in)
	logger.Trace(in.in)
}

func TestALL(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file bool
		in   string
	}{
		name: "TestALL",
		lvl:  ALL,
		file: false,
		in:   "hello world",
	}

	logger := NewToyLog(in.name, in.lvl, in.file)
	logger.Error(in.in)
	logger.Warning(in.in)
	logger.Info(in.in)
	logger.Debug(in.in)
	logger.Trace(in.in)
}

func TestOFF(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file bool
		in   string
	}{
		name: "TestOFF",
		lvl:  OFF,
		file: false,
		in:   "hello world",
	}

	logger := NewToyLog(in.name, in.lvl, in.file)
	logger.Error(in.in)
	logger.Warning(in.in)
	logger.Info(in.in)
	logger.Debug(in.in)
	logger.Trace(in.in)
}
