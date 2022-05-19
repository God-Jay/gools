package publisher

import (
	"embed"
	"os"
	"strings"
)

func CopyTo(dstDir string, fs embed.FS) error {
	fsDirPath := "."
	return copyTo(dstDir, fs, fsDirPath)
}

func copyTo(dstDir string, fs embed.FS, fsDirPath string) error {
	fsDirPath = strings.Replace(fsDirPath, "./", "", -1)

	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		err := os.Mkdir(dstDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	fsDir, err := fs.ReadDir(fsDirPath)
	if err != nil {
		return err
	}

	for _, f := range fsDir {
		if f.IsDir() {
			err := copyTo(dstDir+"/"+f.Name(), fs, fsDirPath+"/"+f.Name())
			if err != nil {
				return err
			}
		} else {
			b, err := fs.ReadFile(fsDirPath + "/" + f.Name())
			if err != nil {
				return err
			}
			err = os.WriteFile(dstDir+"/"+f.Name(), b, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
