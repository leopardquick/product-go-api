package customhandler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Home struct {
	l *log.Logger
}

func NewHome(l *log.Logger) *Home {
	return &Home{l}
}

func (h *Home) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "opps", http.StatusBadRequest)
	}
	fmt.Fprintf(rw, "message %s", data)

}
