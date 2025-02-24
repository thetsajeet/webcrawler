package main

import "errors"

func GetArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New(NO_ARG)
	}
	if len(args) >= 2 {
		return "", errors.New(TOO_MANY_ARGS)
	}
	return args[0], nil
}
