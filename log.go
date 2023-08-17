package isi40core

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func (p *Protocol) LogInfo(message string) {

	_, filename, line, _ := runtime.Caller(1)
	p.Log.Printf("[INFO] %s:%d - %v", filepath.Base(filename), line, message)

}

func (p *Protocol) LogWarning(message string) {

	_, filename, line, _ := runtime.Caller(1)
	p.Log.Printf("[WARN] %s:%d - %v", filepath.Base(filename), line, message)

}

func (p *Protocol) LogError(message string) {

	_, filename, line, _ := runtime.Caller(1)
	p.Log.Printf("[ERRO] %s:%d - %v", filepath.Base(filename), line, message)

}

func (p *Protocol) InternalLogInfo(message string) {

	_, filename, line, _ := runtime.Caller(1)
	p.Log.Printf("[CORE] [INFO] %s:%d - %v", filepath.Base(filename), line, message)

}

func (p *Protocol) InternalLogWarning(message string) {

	_, filename, line, _ := runtime.Caller(1)
	p.Log.Printf("[WARN] [WARN] %s:%d - %v", filepath.Base(filename), line, message)

}

func (p *Protocol) InternalLogError(message string, args ...any) {

	format := "[CORE] [ERRO] %s:%d - %s"

	format2 := ""

	if len(args) != 0 {
		for range args {
			format2 += " %s"
		}
	}

	s := fmt.Sprintf(format2, args)

	_, filename, line, _ := runtime.Caller(1)
	p.Log.Printf(format, filepath.Base(filename), line, message, s)

}
