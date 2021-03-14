package main

import "os"

type VC120 struct {
	Path          string
	CurrentBranch uint
	Branches      []*Branch
}

func InitVC(path string) (*VC120, error) {
	var repo *VC120 = new(VC120)
	repo.Path = path + "/.vcs/"
	repo.Branches = make([]*Branch, 0, 10)
	err := os.MkdirAll(repo.Path+"/blobs", 0775)
	if err != nil {
		return nil, err
	}
	BranchVC("master", nil, repo)
	return repo, nil
}
