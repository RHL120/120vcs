package main

import (
	"fmt"
	"os"
)

func main() {
	var added []*Blob = make([]*Blob, 1, 10)
	added[0] = new(Blob)
	added[0].FilePath = "home"
	added[0].Node = "1235"
	repo, err := InitVC(".")
	if err != nil {
		_ = fmt.Errorf("Well Fuck %v\n", err)
		os.Exit(1)
	}
	b := repo.Branches[0]
	CommitVc("RHL120", added, b, repo)
	CommitVc("RHL120", added, b, repo)
	BranchVC("dsa", repo.Branches[0], repo)
	CommitVc("RHL120", added, repo.Branches[1], repo)
	fmt.Println(repo)

}
