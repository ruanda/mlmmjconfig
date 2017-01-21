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

