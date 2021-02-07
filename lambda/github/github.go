package github

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"github.com/google/go-github/github"
	"github.com/gopetbot/tidus/help"
	errs "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

type Github struct {
	auth    *Auth
	client  *github.Client
	repos   []Repo
	version string
	context context.Context
	logger  *help.Logging
}

func (g *Github) Initialize(accessKey string) error {
	auth, err := g.auth.authenticate(accessKey)
	if err != nil {
		return errs.Wrap(err, "Fail to initialize connection...")
	}
	g.client = github.NewClient(auth)
	return nil

}
func (g *Github) listRepository() error {

	repositories, _, err := g.client.Repositories.List(g.auth.ctx, author, nil)
	if err != nil {
		return errs.New("Was not possible to list repositories..")
	}

	for _, repo := range repositories {
		newRepo := Repo{
			Name:        repo.Name,
			FullName:    repo.FullName,
			Description: repo.Description,
			GitURL:      repo.GitURL,
		}
		g.repos = append(g.repos, newRepo)
	}
	return nil
}

func (g *Github) UpdateVersion() error {

	err := g.NewVersionValidate()
	if err != nil {
		return err
	}

	g.logger.WithFields(logrus.Fields{
		"oldVersion": g.version,
	}).Info("init bumping version.....")

	blob, err := g.blobContent(fileName)
	if err != nil {
		blob = g.generateSHA(message)
		if help.IsEmpty(blob) {
			return errs.Wrap(err, errors.New("was not possible generated a hash").Error())
		}

	}

	g.logger.WithFields(logrus.Fields{
		"blobFile": blob,
	}).Info("finish bumping version.....")

	file := &github.RepositoryContentFileOptions{
		Message: &message,
		Content: []byte(g.version),
		Branch:  github.String(branch),
		Author: &github.CommitAuthor{
			Name:  github.String(author),
			Email: github.String(email),
		},
		SHA: github.String(blob),
		Committer: &github.CommitAuthor{
			Name:  github.String(author),
			Email: github.String(email),
		},
	}

	_, _, err = g.client.Repositories.
		UpdateFile(
			g.auth.ctx, author, repo, fileName, file)

	if err != nil {
		return err
	}

	g.logger.WithFields(logrus.Fields{
		"NewVersion": g.version,
	}).Info("finish bumping version.....")

	return nil
}
func (g *Github) generateSHA(message string) string {
	h := sha1.New()
	h.Write([]byte(message))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

func (g *Github) blobContent(file string) (string, error) {

	fileContent, _, _, err := g.client.Repositories.GetContents(g.context, author, repo, file, &github.RepositoryContentGetOptions{})
	if err != nil {
		return help.Empty(), err
	}
	blob := *fileContent.SHA

	return blob, nil

}
func (g *Github) NewVersionValidate() error {

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
		nil,
		help.Empty(),
		context.Background(),
		help.NewLog(),
	}

}
