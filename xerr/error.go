package xerr

import (
	"fmt"
	"errors"
	log "github.com/sirupsen/logrus"
	"strings"
)

func HandleErr(err *error) {
	if r := recover(); r != nil {
		log.Error(err)
		*err = errors.New(fmt.Sprint(r))
	}
}

func ThrowError(err error) {
	if err !=nil {
		panic(err)
	}
}

func ThrowErrorMessage(format string, args ...string) {
	msg := fmt.Sprintf(format,args)
	panic(msg)
}

func ThrowEmptyError(input string,format string, args ...string){
	if strings.TrimSpace(input) == "" {
		msg := fmt.Sprintf(format,args)
		panic(msg)
	}
}