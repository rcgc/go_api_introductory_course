package main

import (
	"net/http"
	"sync"
)

type carHandler struct {
	sync.Mutex
}

func (h *carHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if idFromUrl(r) == "-1"{
			h.getAll(w, r)
		} //else {
			//h.getById(w, r)
		//}
		
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "invalid method")
	}
}

func newCarHandler() *carHandler {
	
	preload := []Car{
		{"JHk290Xj",	"Ford",		"F10",		"Base",		"Silver",		2010,	"Truck",	12012,	19999}, 
		{"fWl37la",		"Toyota",	"Camry",	"SE",		"White",		2019,	"Sedan",	3999,	28990},
		{"1i3xjRllc",	"Toyota",	"Rav4",		"XSE",		"Red",			2018,	"SUV",		24001,	22750},
		{"dku43920s",	"Ford",		"Bronco",	"Badlands",	"Burnt Orange",	2022,	"SUV",		1,		44990},
	}

	for _, v := range preload {
		v.createCar()
	}
	
	return &carHandler{}
}