package test

import (
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
)

func FakeRelay(f http.HandlerFunc) (*httptest.Server, string) {
	// custom listener to start up on a random port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	s := httptest.NewUnstartedServer(f)
	_ = s.Listener.Close()
	s.Listener = l
	s.Start()

	port := l.Addr().(*net.TCPAddr).Port
	portString := strconv.FormatInt(int64(port), 10)

	return s, portString
}
