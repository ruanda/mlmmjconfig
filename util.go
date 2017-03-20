package mlmmjconfig

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readAllLines(filePath string) ([]string, error) {
	var result []string
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func writeAllLines(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(file)
	for _, l := range lines {
		fmt.Fprintln(w, l)
	}
	w.Flush()
	return nil
}

func readFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func writeFile(filePath string, value string) error {
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(file)
	fmt.Fprint(w, value)
	w.Flush()
	return nil
}

func checkFlagFile(filePath string) (bool, error) {
	fi, err := os.Lstat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	}
	if fi.Mode().IsRegular() {
		return true, nil
	}
	return false, fmt.Errorf("Wrong file type at %s", filePath)
}

func setFlagFile(filePath string, value bool) error {
	if !value {
		return os.Remove(filePath)
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
