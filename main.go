package main

import (
	"fmt"
	"os"
)

func example() {
	repo, err := InitVC(".")
	var added []*Blob = make([]*Blob, 1, 10)
	added[0], err = newBlob(repo, "./add.go");
	if (err != nil) {
		fmt.Println (err);
	}
	if err != nil {
		_ = fmt.Errorf("Well Fuck %v\n", err)
		os.Exit(1)
	}
	b := repo.Branches[0]
	CommitVc("RHL120", added, b, repo)
	CommitVc("RHL120", added, b, repo)
	BranchVC("dsa", repo.Branches[0], repo)
	CommitVc("RHL120", added, repo.Branches[1], repo)
	fmt.Println(repo.Branches)
	DeleteVC("dsa", repo)
	fmt.Println(repo.Branches)
	fmt.Println(repo)

}
func main() {
	example()
}
