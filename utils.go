package main

import "os"
import "strings"
import "crypto/sha256"
import "fmt"

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

func sha256File(path string) (hash string, err error) {
	var fcb []byte;
	fcb, err = os.ReadFile (path);
	if err != nil {
		return hash, err;
	}
	hash = fmt.Sprintf ("%x", sha256.Sum256 (fcb));
	return hash, err
}
