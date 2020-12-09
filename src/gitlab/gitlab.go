package gitlab

import (
	"fmt"

	gl "github.com/xanzy/go-gitlab"
)

// Gitlab strut
type Gitlab struct {
	token     string
	Client    *gl.Client
	idProject int
}

// Commit Commit
func (g *Gitlab) Commit(nameFile *string, content *string, commitMessage string, nameBranch string) *gl.Branch {
	branch, _, _ := g.Client.Branches.GetBranch(g.idProject, nameBranch)
	var action gl.FileAction = gl.FileUpdate
	_, _, err := g.Client.Commits.CreateCommit(g.idProject, &gl.CreateCommitOptions{
		Branch:        &nameBranch,
		CommitMessage: &commitMessage,
		Actions: []*gl.CommitActionOptions{
			{
				Action:   &action,
				FilePath: nameFile,
				Content:  content,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	return branch
}

// CreateMergeRequest create Merge
func (g *Gitlab) CreateMergeRequest(sourceBranch string, targetBranch string, title string) {
	_, _, err := g.Client.MergeRequests.CreateMergeRequest(g.idProject, &gl.CreateMergeRequestOptions{
		Title:        &title,
		SourceBranch: &sourceBranch,
		TargetBranch: &targetBranch,
	})

	if err != nil {
		fmt.Println(err)
	}
}

// CreateBranch create new branch
func (g *Gitlab) CreateBranch(branchName string, ref string) {

	_, _, err := g.Client.Branches.CreateBranch(g.idProject, &gl.CreateBranchOptions{
		Branch: &branchName,
		Ref:    &ref,
	})

	fmt.Println(err)
}

// StartClient client start
func (g *Gitlab) StartClient(idProject int, baseURL string) *Gitlab {
	git, err := gl.NewClient(g.token, gl.WithBaseURL(baseURL))
	if err != nil {
		fmt.Println(err)
	}

	g.Client = git
	g.idProject = idProject
	return g
}

// NewGitlab Ioc
func NewGitlab(token string) *Gitlab {
	return &Gitlab{
		token: token,
	}
}
