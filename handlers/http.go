package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mdwhatcott/calc-lib/calc"
)

func NewHTTPRouter() http.Handler {
	router := http.NewServeMux()
	router.Handle("/add", NewHTTPHandler(calc.Addition{}))
	router.Handle("/sub", NewHTTPHandler(calc.Subtraction{}))
	router.Handle("/mul", NewHTTPHandler(calc.Multiplication{}))
	router.Handle("/div", NewHTTPHandler(calc.Division{}))
	return router
}

type HTTPHandler struct {
	calculator Calculator
}

func NewHTTPHandler(calculator Calculator) *HTTPHandler {
	return &HTTPHandler{calculator: calculator}
}

func (this *HTTPHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/plain; charset=utf-8")

	query := request.URL.Query()
	a, err := strconv.Atoi(query.Get("a"))
	if err != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = fmt.Fprint(response, "invalid parameter: 'a'")
		return
	}
	b, err := strconv.Atoi(query.Get("b"))
	if err != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = fmt.Fprint(response, "invalid parameter: 'b'")
		return
	}
	result := this.calculator.Calculate(a, b)
	response.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(response, "%d", result)
	if err != nil {
		log.Println(err)
	}
}
