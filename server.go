package palantir

import (
	"net/rpc"
	"fmt"
	"net"
)


type Server struct {

}

func (s Server) FileReport(r Report, ok *bool) error {
	fmt.Printf("%+v", r)
	fmt.Println(string(r.Stack))
	return nil
}

func Start(server *Server, port string) {
  rpc.Register(server)
  listener, e := net.Listen("tcp", ":"+port)
  if e != nil {
    // log.Fatal("listen error:", e)
  } else {
	  for {
	    if conn, err := listener.Accept(); err != nil {
	      // log.Fatal("accept error: " + err.Error())
	    } else {
	      go rpc.ServeConn(conn)
	    }
	  }
  }
}

