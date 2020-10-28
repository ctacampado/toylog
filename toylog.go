// Created on Sat Nov 30 2019
// author: ctacampado

// Package toylog implements a logging package based on
// go's standard log library that features a configurable
// logging level.
package toylog

import (
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
	PROD LogLvl = iota
	TRACE
	DEBUG
	INFO
	WARNING
	ERR
	FATAL
	PANIC
)

// A ToyLog represents a container object for a
// logger and its current configuration parameters
type ToyLog struct {
	// name of the logger used for labeling
	// logs for easier tracing
	name string
	// logging level
	lvl LogLvl
	// stdout or file
	File *os.File
	// log file name
	FileName string
	// pointer to the active logger object
	l *log.Logger
}

func initLoggerName(s string) (lname string) {
	if s != "" {
		return fmt.Sprintf("[%s] ", s)
	}
	return "[toylog] "
}

// file name is using the format YYYY_MM_DD_HH_MM_SS.log
func initLogFile() (f *os.File, fname string) {
	fname = time.Now().Format("2006_01_02_15_04_05") + ".log"
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// NewToyLog initializes ToyLog struct.
// Valid arguments are as follows:
// - type string for the logger name
// - type int for the log level
// - type bool set to true if you want to output logs to a file.
// If it is an unknown parameter type, it returns nil
func NewToyLog(args ...interface{}) (tl *ToyLog) {
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
			if a >= 0 && a <= PANIC {
				tl.lvl = a
			}
		case bool:
			if a == true {
				tl.File, tl.FileName = initLogFile()
			}
		default:
			err := fmt.Sprintf("unknown parameter '%v' of type %s\n", arg, reflect.TypeOf(arg))
			log.Fatal(err)
		}
	}
	tl.l = log.New(tl.File, tl.name, log.LstdFlags)
	return tl
}

// Close log file
func (t *ToyLog) Close() {
	if t.File != os.Stdout {
		t.File.Close()
	}
}

// Trace level
func (t *ToyLog) Trace(format string, v ...interface{}) {
	if t.lvl >= TRACE {
		t.l.SetFlags(log.LstdFlags)
		t.l.Output(2, fmt.Sprintf("[TRACE] {\"trace\":\""+format+"\"}", v...))
	}
}

// Debug level
func (t *ToyLog) Debug(format string, v ...interface{}) {
	if t.lvl >= DEBUG {
		t.l.SetFlags(log.LstdFlags)
		t.l.Output(2, fmt.Sprintf("[DEBUG] {\"debug\":\""+format+"\"}", v...))
	}
}

// Info level
func (t *ToyLog) Info(format string, v ...interface{}) {
	if t.lvl >= INFO {
		t.l.SetFlags(log.LstdFlags)
		t.l.Output(2, fmt.Sprintf("[INFO] {\"info\":\""+format+"\"}", v...))
	}
}

// Warning level
func (t *ToyLog) Warning(format string, v ...interface{}) {
	if t.lvl >= INFO {
		t.l.SetFlags(log.LstdFlags)
		t.l.Output(2, fmt.Sprintf("[WARN] {\"info\":\""+format+"\"}", v...))
	}
}

// Error level
func (t *ToyLog) Error(format string, v ...interface{}) {
	if t.lvl >= ERR {
		t.l.SetFlags(log.LstdFlags)
		t.l.Output(2, fmt.Sprintf("[ERROR] {\"error\":\""+format+"\"}", v...))
	}
}

// Fatal level
func (t *ToyLog) Fatal(format string, v ...interface{}) {
	if t.lvl >= ERR {
		t.l.SetFlags(log.LstdFlags)
		t.l.Output(2, fmt.Sprintf("[FATAL] {\"error\":\""+format+"\"}", v...))
	}
	os.Exit(1)
}

// Panic level
func (t *ToyLog) Panic(format string, v ...interface{}) {
	if t.lvl >= ERR {
		t.l.SetFlags(log.LstdFlags)
		t.l.Panic(2, fmt.Sprintf("[FATAL] {\"error\":\""+format+"\"}", v...))
	}
}
