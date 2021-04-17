package main

import (
	"fmt"
	"io"
	"os"
)
type Blob struct {
	FilePath string
	Node     string
}

func newBlob(repo *VC120, path string) (blob *Blob, err error) {
	var cfile, bfile *os.File;
	blob = new (Blob);
	cfile, err = os.Open (path);
	if err != nil {
		return blob, err;
	}
	defer cfile.Close()
	blob.FilePath, err = path2vcpath (path);
	if err != nil {
		return blob, err;
	}
	blob.Node, err = sha256File (path);
	if err != nil {
		return blob, err;
	}
	fmt.Println(blob.Node);
	bfile, err = os.Create (repo.Path + "/blobs/" + blob.Node);
	if err != nil {
		return blob, err;
	}
	_, err = io.Copy (bfile, cfile);
	return blob, err;
}
