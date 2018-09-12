package models

import (
	"fmt"
)

type RestError struct {
	Err     string
	ErrCode int
}

func (e *RestError) Error() string {
	return fmt.Sprint(e.Err)
}
