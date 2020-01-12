package utils

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type TestServiceSuit struct {
	suite.Suite
	opt HttpServiceOption
}

func (s *TestServiceSuit) SetupTest() {
	s.opt.Port = "8999"
	s.opt.Addr = "127.0.0.1"
	go NewHttpServer(&s.opt)
}

func (s *TestServiceSuit) TestNewHttpService() {
	resp, err := http.Get("http://" + s.opt.Addr + ":" + s.opt.Port + "/hello")
	s.NoError(err)
	defer resp.Body.Close()
	s.Exactly("200 OK", resp.Status)
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(TestServiceSuit))
}
