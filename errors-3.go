package main

import (
	"errors"
	"log"
)

type nonCriticalError struct{}

func (e nonCriticalError) Error() string {
	return "validation error"
}

var (
	errBadConnection = errors.New("bad connection")
	errBadRequest    = errors.New("bad request")
)

const unknownErrorMsg = "unknown error"

var criticalErrs = []error{errBadRequest, errBadConnection}

/* 
Возвращает текст ошибки, если она критичная. 
Пустую строку если некритичная. 
В случае неизвестной ошибки возвращается строка unknown error
*/
func GetErrorMsg(err error) string {
	for _, crErr := range criticalErrs {
		if errors.Is(err, crErr) {
			return crErr.Error()
		}
	}

	if errors.As(err, &nonCriticalError{}) {
		return ""
	}

	return unknownErrorMsg
}

func main() {
	log.Println(GetErrorMsg(errBadConnection))           // "bad connection"
	log.Println(GetErrorMsg(errBadRequest))              // "bad request"
	log.Println(GetErrorMsg(nonCriticalError{}))         // ""
	log.Println(GetErrorMsg(errors.New("random error"))) // "unknown error"
}