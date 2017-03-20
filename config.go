package mlmmjconfig

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type MlmmjConfig struct {
	baseDir string
}

func Open(baseDir string) (*MlmmjConfig, error) {
	f, err := os.Stat(baseDir)
	if err != nil {
		return nil, err
	}
	if !f.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", baseDir)
	}
	c := MlmmjConfig{baseDir}
	return &c, nil
}

func (c *MlmmjConfig) Lists() ([]MlmmjList, error) {
	var result []MlmmjList
	files, err := ioutil.ReadDir(c.baseDir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			result = append(result, MlmmjList{filepath.Join(c.baseDir, f.Name())})
		}
	}
	return result, nil
}
