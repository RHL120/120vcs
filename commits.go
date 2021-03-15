package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type Blob struct {
	FilePath string
	Node     string
}
type Commit struct {
	Prev      *Commit
	Blobs     []*Blob
	Author    string
	Timestamp int64
	Hash      string
	Next      []*Commit
}

func copyCommits(dst string, src string) error {
	sfiles, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, i := range sfiles {
		d, err := os.Create(dst + "/" + i.Name())
		if err != nil {
			return err
		}
		s, err := os.Open(src + "/" + i.Name())
		if err != nil {
			return err
		}
		_, err = io.Copy(d, s)
		if err != nil {
			return err
		}
	}
	return err
}
func WriteCommit(c *Commit) (output string, hash string) {
	output = fmt.Sprintf("authour:%s\ntimestamp:%d", c.Author, c.Timestamp)
	for _, i := range c.Blobs {
		output = fmt.Sprintf("%s\n%s:%s", output, i.Node, i.FilePath)
	}
	hash = fmt.Sprintf("%x", sha256.Sum256([]byte(output)))
	if c.Prev != nil {
		output = fmt.Sprintf("%s\nprev:%s", output, c.Prev.Hash)
	}
	return output, hash
}
func CommitVc(auth string, added []*Blob, parent *Branch, repo *VC120) error {
	var c *Commit = new(Commit)
	var output string
	c.Author = auth
	c.Blobs = added
	c.Timestamp = time.Now().UnixNano()
	c.Prev = parent.Head
	parent.Head = c
	output, c.Hash = WriteCommit(c)
	err := ioutil.WriteFile(repo.Path+"/branches/"+parent.Name+"/commits/"+c.Hash, []byte(output), 0664)
	if err != nil {
		return err
	}
	if c.Prev != nil {
		c.Prev.Next = append(c.Prev.Next, c)
		pfile, err := os.OpenFile(repo.Path+"branches/"+parent.Name+"/commits/"+c.Prev.Hash, os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
			fmt.Println(err)
			return err
		}
		pfile.WriteString("\nnext:" + c.Hash)
		pfile.Close()
	}
	return nil
}
