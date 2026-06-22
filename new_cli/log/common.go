package mylog

import "log"

var DebugTag bool = false

func Info(format string, args ...any) {
	log.Printf(format, args...)
}

func Debug(format string, args ...any){
	if DebugTag {
		log.Printf(format, args...)
	}
}

func Error(format string, args ...any) {
	log.Printf(format, args...)
}