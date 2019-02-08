package util

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"
)

/*
	How user:
	* Insert the next code in the lowest level function that can return a error:
		defer func() { err = util.GetFullErr(err) }()


	* The function will look like:
		func ... {
			defer func() { err = util.GetFullErr(err) }()
			...
		}

*/

func GetFullErr(err error) error {
	if err == nil {
		return nil
	}

	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	fullErr := `
====================================================
				%s
====================================================
::::::::::::::::::: ERROR TRACE ::::::::::::::::::::

*** Full err msg: %s

*** Full err content: %s

*** Full stack trace: %s
`
	fullErrContent, _ := json.Marshal(err)
	stackTrace := string(debug.Stack())
	return fmt.Errorf(fullErr, timeStamp, err.Error(), string(fullErrContent), stackTrace)
}
