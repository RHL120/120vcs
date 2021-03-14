package main

import "os"

type Branch struct {
	Name string
	Head *Commit
}

func BranchVC(bname string, parent *Branch, repo *VC120) error {
	var nb *Branch = new(Branch)
	nb.Name = bname
	err := os.MkdirAll(repo.Path+"branches/"+bname+"/commits/", 0775)
	if err != nil {
		return err
	}
	if parent != nil {
		nb.Head = parent.Head
		err := copyCommits(repo.Path+"branches/"+nb.Name+"/commits/", repo.Path+"branches/"+parent.Name+"/commits/")
		if err != nil {
			return err
		}
	}
	repo.Branches = append(repo.Branches, nb)
	repo.CurrentBranch = uint(len(repo.Branches))
	return nil
}
func getIndex(branchn string, repo *VC120) int {
	for i, e := range repo.Branches {
		if e.Name == branchn {
			repo.CurrentBranch = uint(i)
			return i
		}
	}
	return -1
}
func CheckoutVC(branchn string, repo *VC120) {
	repo.CurrentBranch = uint(getIndex(branchn, repo))
}
func DeleteVC(branchn string, repo *VC120) error {
	var bindex int
	bindex = getIndex(branchn, repo)
	err := os.RemoveAll(repo.Path + "branches/" + repo.Branches[bindex].Name)
	repo.Branches = append(repo.Branches[:bindex], repo.Branches[bindex+1:]...)
	return err
}
