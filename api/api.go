package api

import (
	"github.com/jmcvetta/napping"
	"net/http"
)

func Setup() napping::Session {
	s := napping.Session{ Log: false }
	s.Client = &http.Client{}
	return s
}
