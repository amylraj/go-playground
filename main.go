package main

import (
	"fmt"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	fmt.Println("Test Random Stuff")
	log.SetOutput(&lumberjack.Logger{
		Filename:   "/var/log/myapp/foo.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
}
