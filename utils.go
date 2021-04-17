package main

import "os"
import "strings"
import "crypto/sha256"
import "fmt"
import "io/ioutil"

func absp(path string) (ret string, err error) {
	if path[0] == '/' {
		return path, nil
	}
	ret, err = os.Getwd()
	ret += path
	return ret, err
}

func path2vcpath(path string) (ret string, err error) {
	ret, err = absp(path)
	if err != nil {
		return ret, err
	}
	ret = strings.ReplaceAll(path, ret, "")
	return ret, err
}

func sha256File(path string) (hash string, err error) {
	var fcb []byte
	fcb, err = os.ReadFile(path)
	if err != nil {
		return hash, err
	}
	hash = fmt.Sprintf("%x", sha256.Sum256(fcb))
	return hash, err
}

func lsr(dst *[]string, path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Println(file.Name())
		if file.Name() == "." || file.Name() == ".." {
			continue
		}
		*dst = append(*dst, path+"/"+file.Name())
		if !file.IsDir() {
			continue
		}
		err = lsr(dst, (*dst)[len(*dst)-1])
		if err != nil {
			return err
		}
	}
	return nil
}

func list_recursive(path string) ([]string, error) {
	files := make([]string, 5)
	err := lsr(&files, path)
	return files, err
}
