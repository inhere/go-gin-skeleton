package handler

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/gcli/v2"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/jsonutil"
	"github.com/inhere/go-gin-skeleton/model"
)

var gitOpts GitOpts

type GitOpts struct {
	output string
}

// GitCommand
func GitCommand() *gcli.Command {
	cmd := gcli.Command{
		Name:        "git:info",
		Aliases:     []string{"gitinfo"},
		UseFor: "collect project info by git info",

		Func: gitExecute,
	}

	gitOpts = GitOpts{}

	cmd.StrOpt(&gitOpts.output, "output", "o", "static/app.json", "output file of the git info")

	return &cmd
}

// arg test:
// 	go build cliapp.go && ./cliapp git:info
func gitExecute(_ *gcli.Command, _ []string) error {
	info := model.AppInfo{}

	// latest commit id by: git log --pretty=%H -n1 HEAD
	cid, err := cliutil.QuickExec("git log --pretty=%H -n1 HEAD")
	if err != nil {
		return err
	}

	cid = strings.TrimSpace(cid)
	fmt.Printf("commit id: %s\n", cid)
	info.Version = cid

	// latest commit date by: git log -n1 --pretty=%ci HEAD
	cDate, err := cliutil.QuickExec("git log -n1 --pretty=%ci HEAD")
	if err != nil {
		return err
	}

	cDate = strings.TrimSpace(cDate)
	info.ReleaseAt = cDate
	fmt.Printf("commit date: %s\n", cDate)

	// get tag: git describe --tags --exact-match HEAD
	tag, err := cliutil.QuickExec("git describe --tags --exact-match HEAD")
	if err != nil {
		// get branch: git branch -a | grep "*"
		br, err := cliutil.QuickExec(`git branch -a | grep "*"`)
		if err != nil {
			return err
		}
		br = strings.TrimSpace(strings.Trim(br, "*"))
		info.Tag = br
		fmt.Printf("current branch: %s\n", br)
	} else {
		tag = strings.TrimSpace(tag)
		info.Tag = tag
		fmt.Printf("latest tag: %s\n", tag)
	}

	err = jsonutil.WriteFile(gitOpts.output, &info)
	if err != nil {
		return err
	}

	color.Green.Println("\nOk, project info collect completed!")
	return nil
}
