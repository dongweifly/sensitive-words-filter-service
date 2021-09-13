package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type Matcher interface {
	//Build build Matcher
	Build(words []string)

	//Match return match sensitive words
	Match(text string) ([]string, string)
}

var (
	GMatchService = NewMatchService(NewDFAMather())
)

type MatchService struct {
	dictDir string
	matcher Matcher
}

func NewMatchService(matcher Matcher) *MatchService {
	return &MatchService{
		matcher: matcher,
	}
}

func (m *MatchService) Init(dictPath string) error {
	m.dictDir = dictPath

	var files []string
	//只关注rootDir下面的文件，暂时不支持子文件夹的形式
	err := filepath.Walk(m.dictDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Warnf("load %s fail %s", m.dictDir, err.Error())
			return nil
		}

		if !info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	})

	if err != nil {
		return err
	}

	var words []string
	for _, file := range files {
		if w, err := readWordsFromFile(m.dictDir + "/" + file); err == nil {
			log.Infof("load senvitive files  %s words size %d", m.dictDir+"/"+file, len(w))
			words = append(words, w...)
		} else {
			return err
		}
	}

	m.matcher.Build(words)
	return nil
}

func (m *MatchService) Match(text string) ([]string, string) {
	return m.matcher.Match(text)
}
