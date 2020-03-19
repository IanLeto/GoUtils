package utils

import (
	"io/ioutil"
	"net/http"
	"transfer/logging"
)

func HelloWorld() {
	res, err := http.Get("http://httpbin.org/anything?hello=world")
	NoErr(err)
	defer res.Body.Close()
	bs, err := ioutil.ReadAll(res.Body)
	NoErr(err)
	logging.GetStdLogger().Info(bs)
}
