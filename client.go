package palantir

import (
	"net/rpc"
	"fmt"
)

type PalantirClient struct{
	Host string
	Port string
	Level int
}

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

func (p PalantirClient) Fatal(s string, v ...interface{}) {
	if cl, err := p.client(); err == nil {
		r := NewReport("Fatal", fmt.Sprintf(s, v...))
		var reply bool
		go func() {
			cl.Call("Server.FileReport", r, &reply)
		}()
	}
}

func (p PalantirClient) Error(s string, v ...interface{}) {
	if p.Level < 5 {
		if cl, err := p.client(); err == nil {
			r := NewReport("Error", fmt.Sprintf(s, v...))
			var reply bool
			go func() {
		  	cl.Call("Server.FileReport", r, &reply)
		  }()
		}
	}
}

func (p PalantirClient) Warn(s string, v ...interface{})  {
	if p.Level < 4 {
		if cl, err := p.client(); err == nil {
			r := NewReport("Warn", fmt.Sprintf(s, v...))
		  var reply bool
		  go func() {
		  	cl.Call("Server.FileReport", r, &reply)
		  }()
		}
	}
}

func (p PalantirClient) Info(s string, v ...interface{})  {
	if p.Level < 3 {
		if cl, err := p.client(); err == nil {
			r := NewReport("Info", fmt.Sprintf(s, v...))
		  var reply bool
		  go func() {
		  	cl.Call("Server.FileReport", r, &reply)
		  }()
		}
	}
}

func (p PalantirClient) Debug(s string, v ...interface{}) {
	if p.Level < 2 {
		if cl, err := p.client(); err == nil {
			r := NewReport("Debug", fmt.Sprintf(s, v...))
		  var reply bool
		  go func() {
		  	cl.Call("Server.FileReport", r, &reply)
		  }()
		}
	}
}

func (p PalantirClient) Trace(s string, v ...interface{}) {
	if p.Level < 1 {
		if cl, err := p.client(); err == nil {
			r := NewReport("Trace", fmt.Sprintf(s, v...))
		  var reply bool
		  go func() {
		  	cl.Call("Server.FileReport", r, &reply)
		  }()
		}
	}
}



func (p PalantirClient) client() (*rpc.Client, error) {
  return rpc.Dial("tcp", p.Host+":"+p.Port)
}