package mlmmjconfig

import (
    "fmt"
    "os"
)

type MlmmjConfig struct {
    baseDir string
}

func Open(baseDir string) (*MlmmjConfig, error) {
    f, err := os.Stat(baseDir)
    if err != nil {
        return nil, err
    }
    if !f.IsDir()  {
        return nil, fmt.Errorf("%s is not a directory", baseDir)
    }
    c := MlmmjConfig{baseDir}
    return &c, nil
}

