package ienv

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
	"os"
)

type Source struct{}

func (*Source) fetch(url string, dir string) {
	log.Info(fmt.Sprintf("Executing `git clone %s %s`", url, dir))
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (*Source) cleanup(dir string) {
	log.WithFields(log.Fields{
		"dir": dir,
	}).Info("Removing directory")
	os.RemoveAll(dir)
}
