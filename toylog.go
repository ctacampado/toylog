// Created on Sat Nov 30 2019
// author: ctacampado

// Package toylog implements a logging package based on
// go's standard log library that features a configurable
// logging level.
package toylog

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

// LogLvl is a custom type based on int.
// It is used to differenciate a valid log level argument
// vs some number
type LogLvl int

//logging levels
const (
	DEBUG LogLvl = iota
	INFO
	NOTICE
	WARNING
	ERR
	CRIT
	ALERT
	EMRG
)

// A ToyLog represents a container object for a
// logger and its current configuration parameters
type ToyLog struct {
	name     string      // name of the logger used for labeling logs for easier tracing
	lvl      LogLvl      // logging level
	File     *os.File    // stdout or file
	FileName string      // log file name
	logger   *log.Logger // pointer to the active logger object
}

func initLoggerName(s string) (lname string) {
	if s != "" {
		lname = fmt.Sprintf("[%s] ", s)
	} else {
		lname = "[toylog] "
	}
	return
}

func initLogFile() (f *os.File, fname string, err error) {
	fname = time.Now().Format("2006_01_02_15_04_05") + ".log"
	f, err = os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, "", err
	}
	return
}

// NewToyLog initializes ToyLog struct. Valid arguments are
// as follows:
// - type string for the logger name
// - type int for the log level
// - type bool set to true if you want to output logs to a file.
//   file name is using the format YYYY_MM_DD_HH_MM_SS.log
//
// If it is an unknown parameter type, it returns nil
func NewToyLog(args ...interface{}) (tl *ToyLog, err error) {
	tl = &ToyLog{
		name: "logger",
		lvl:  0,
		File: os.Stdout,
	}
	for _, arg := range args {
		switch a := arg.(type) {
		case string:
			tl.name = initLoggerName(a)
		case LogLvl:
			if a >= 0 && a <= EMRG {
				tl.lvl = a
			}
		case bool:
			if a == true {
				tl.File, tl.FileName, err = initLogFile()
				if err != nil {
					return nil, err
				}
			}
		default:
			err := fmt.Sprintf("unknown parameter '%v' of type %s\n", arg, reflect.TypeOf(arg))
			return nil, errors.New(err)
		}
	}
	tl.logger = log.New(tl.File, tl.name, log.LstdFlags)
	return tl, nil
}

// Close log file
func (t *ToyLog) Close() {
	if t.File != os.Stdout {
		t.File.Close()
	}
}

// Info level is a logging level that uses log.LstdFlags only.
// Using Info log level also shows Debug level logs
func (t *ToyLog) Info(format string, v ...interface{}) {
	if t.lvl >= INFO {
		t.logger.SetFlags(log.LstdFlags)
		newfmt := fmt.Sprintf("[INFO]: " + format)
		t.logger.Output(2, fmt.Sprintf(newfmt, v...))
	}
}

// Debug level is a logging level that uses log.LstdFlags | log.Lshortfile flags.
// This is the default logging level
func (t *ToyLog) Debug(format string, v ...interface{}) {
	t.logger.SetFlags(log.LstdFlags | log.Lshortfile)
	newfmt := fmt.Sprintf("[DEBUG]: " + format)
	t.logger.Output(2, fmt.Sprintf(newfmt, v...))
}

// Error level is a logging level that uses log.LstdFlags | log.Lshortfile | log.Llongfile flags.
// Using this log level will show all log levels less than or equal to ERR level
func (t *ToyLog) Error(format string, v ...interface{}) {
	t.logger.SetFlags(log.LstdFlags | log.Lshortfile | log.Llongfile)
	if t.lvl >= ERR {
		newfmt := fmt.Sprintf("[ERROR]: " + format)
		t.logger.Output(2, fmt.Sprintf(newfmt, v...))
	}
}
