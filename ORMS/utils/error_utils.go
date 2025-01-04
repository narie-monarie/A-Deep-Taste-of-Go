package utils

import "errors"

func HandleError(err error) {
	if err != nil {
		errors.Unwrap(err)
	}
}
