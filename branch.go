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
	return nil
}
