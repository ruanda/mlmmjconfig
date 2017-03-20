package mlmmjconfig

import (
	"path/filepath"
)

type MlmmjList struct {
	baseDir string
}

func (l *MlmmjList) Name() string {
	return filepath.Base(l.baseDir)
}

func (l *MlmmjList) getFilePath(name string) string {
	return filepath.Join(l.baseDir, name)
}

func (l *MlmmjList) Owners() []string {
	fp := l.getFilePath("control/owner")
	result, err := readAllLines(fp)
	if err != nil {
		return []string{}
	}
	return result
}

func (l *MlmmjList) SetOwners(owners []string) error {
	fp := l.getFilePath("control/owner")
	return writeAllLines(fp, owners)
}

func (l *MlmmjList) Prefix() string {
	fp := l.getFilePath("control/prefix")
	prefix, err := readFile(fp)
	if err != nil {
		return ""
	}
	return prefix
}

func (l *MlmmjList) SetPrefix(prefix string) error {
	fp := l.getFilePath("control/prefix")
	return writeFile(fp, prefix)
}

func (l *MlmmjList) ClosedList() bool {
	fp := l.getFilePath("control/closedlist")
	state, err := checkFlagFile(fp)
	if err != nil {
		return false
	}
	return state
}

func (l *MlmmjList) SetClosedList(state bool) error {
	fp := l.getFilePath("control/closedlist")
	return setFlagFile(fp, state)
}
