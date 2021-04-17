package main

import (
	"os"
	"strings"
)

func absp(path string) (ret string, err error) {
	if (path[0] == '/') {
		return path, nil;
	}
	ret, err = os.Getwd();
	ret += path;
	return ret, err;
}

func path2vcpath(path string) (ret string, err error) {
	ret, err = absp(path);
	if err != nil {
		return ret, err;
	}
	ret = strings.ReplaceAll (path, ret, "");
	return ret, err;
}
