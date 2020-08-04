package file

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)

func HomeDir() string {
	p, _ := homedir.Expand("~/")
	return p
}

func WriteDataToFile(filePath string, data []byte) error {
	exists, err := Exists(path.Dir(filePath))
	if err != nil {
		return xerrors.Errorf("check filePath %s failed", filePath)
	}

	if !exists {
		if err := os.MkdirAll(path.Dir(filePath), 0755); err != nil {
			return xerrors.Errorf("mkdirAll failed, %w", err)
		}
	}

	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		return xerrors.Errorf("write file failed, %w", err)
	}

	return nil
}

func FetchDataFromFile(filePath string) ([]byte, error) {
	exists, err := Exists(filePath)
	if err != nil {
		return nil, xerrors.Errorf("check filePath %s failed", filePath)
	}
	if !exists {
		return nil, xerrors.Errorf("file %s not exist", filePath)
	}

	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, xerrors.New("load file failed")
	}

	return contents, err
}

func Exists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
