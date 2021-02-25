package tgz

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
)

func Test_Uncompress(t *testing.T) {
	//file := "./chaincode.tar.gz"
	file := "../../code.tar.gz"
	dest := "./chaincode"
	//dest := "/d/share/gopath/src/github.com/zhj0811/gostudy/code/tgz/chaincode"
	err := uncompressTGZ(file, dest)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("success")
}

func uncompressTGZ(tgzFile, destDir string) error {
	fr, err := os.Open(tgzFile)
	defer fr.Close()
	if err != nil {
		return errors.Errorf("open %s failed", tgzFile)
	}
	// gzip read
	gr, err := gzip.NewReader(fr)
	defer gr.Close()
	if err != nil {
		return errors.WithMessage(err, "gzip reader error")
	}

	// tar read
	tr := tar.NewReader(gr)
	// 读取文件

	if err := os.MkdirAll(destDir, 0755); err != nil {
		return errors.WithMessagef(err, "mkdir %s failed", destDir)
	}
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if h.FileInfo().IsDir() {
			err := os.MkdirAll(filepath.Join(destDir, h.Name), 0755)
			if err != nil {
				return errors.WithMessage(err, "test mkdir error")
			}
			continue
		}
		// 打开文件
		fw, err := os.OpenFile(filepath.Join(destDir, h.Name), os.O_CREATE|os.O_WRONLY, os.FileMode(h.Mode))
		defer fw.Close()
		if err != nil {
			return err
		}

		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			return err
		}
	}
	return nil
}

//func createFile(name string) (*os.File, error) {
//	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
//	if err != nil {
//		return nil, err
//	}
//	return os.Create(name)
//}
