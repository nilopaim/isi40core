package isi40core

import (
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

func (p *Protocol) InternalLogError(message string) {

	_, filename, line, _ := runtime.Caller(1)
	p.Log.Printf("[CORE] [ERRO] %s:%d - %s", filepath.Base(filename), line, message)

}
