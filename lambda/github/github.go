package github

import (
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
)

type Github struct {
	auth    *Auth
	client  *github.Client
	repos   []*github.Repository
	version string
}

func (g *Github) Initialize(accessKey string) error {
	auth, err := g.auth.authenticate(accessKey)
	if err != nil {
		return errors.Wrap(err, "Fail to initialize connection...")
	}
	g.client = github.NewClient(auth)
	return nil

}
func (g *Github) ListRepository() error {

	repositories, _, err := g.client.Repositories.List(g.auth.ctx, "", nil)
	if err != nil {
		return errors.New("Was not possible to list repositories..")
	}

	for _, repo := range repositories {
		g.repos = append(g.repos, repo)
	}
	return nil
}

func (g *Github) UpdateVersion() error {
	err := g.Version()
	if err != nil {
		return err
	}
	file := &github.RepositoryContentFileOptions{
		Message: &message,
		Content: []byte(g.version),
		Branch:  github.String(branch),
		Author: &github.CommitAuthor{
			Name:  github.String(author),
			Email: github.String(email),
		},
		Committer: &github.CommitAuthor{
			Name: github.String(author),
		},
	}

	g.client.Repositories.
		UpdateFile(
			g.auth.ctx, "", "", "", file)

	return nil
}
func (g *Github) Version() error {

	var file, err = os.OpenFile("./VERSION", os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	var version = make([]byte, 1024)
	_, err = file.Read(version)

	if err != nil {
		return err
	}
	newVersion, err := strconv.Atoi(strings.Split(string(version), "")[4])
	if err != nil {
		return err
	}
	newVersion += 1

	_, err = file.WriteAt([]byte(strconv.Itoa(newVersion)), 4)
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}
	g.version = string(version)
	return nil
}

//NewGithub
func NewGithub() *Github {
	return &Github{
		NewAuthentication(),
		github.NewClient(nil),
		[]*github.Repository{},
		string("version"),
	}
}
