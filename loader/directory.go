package loader

import (
	"github.com/henomis/lingoose/document"

	"os"
	"path/filepath"
	"regexp"
)

type directoryLoader struct {
	dirname string
	regExp  *regexp.Regexp
}

func NewDirectoryLoader(dirname string, regExPathMatch string) (*directoryLoader, error) {

	regExp, err := regexp.Compile(regExPathMatch)
	if err != nil {
		return nil, err
	}

	return &directoryLoader{
		dirname: dirname,
		regExp:  regExp,
	}, nil

}

func (t *directoryLoader) Load() ([]document.Document, error) {
	docs := []document.Document{}

	err := filepath.Walk(t.dirname, func(path string, info os.FileInfo, err error) error {
		if err == nil && t.regExp.MatchString(info.Name()) {

			l, err := NewTextLoader(path, nil)
			if err != nil {
				return err
			}

			d, err := l.Load()
			if err != nil {
				return err
			}

			docs = append(docs, d...)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return docs, nil
}