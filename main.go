package main

import (
	"net/http"
	"log"
)

// type Server struct {
// 	Addr string
// 	Handler Handler
// 	DisableGeneralOptionsHandler bool
// 	TLSConfig *tls.Config
// 	ReadTimeout time.Duration
// 	ReadHeaderTimeout time.Duration
// 	WriteTimeout time.Duration
// 	IdleTimeout time.Duration
// 	MaxHeaderBytes int
// 	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
// 	ConnState func(net.Conn, ConnState)
// 	ErrorLog *log.Logger
// 	BaseContext func(net.Listener) context.Context
// 	ConnContext func(ctx context.Context, c net.Conn) context.Context
// 	HTTP2 *HTTP2Config
// 	Protocols *Protocols
// }

func main() {
	serveMux := http.NewServeMux()
	server := http.Server{
		Addr: ":8080",
		Handler: serveMux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("RIP")
	}
	// var err error = nil
	// for err == nil {
		
	// }
}