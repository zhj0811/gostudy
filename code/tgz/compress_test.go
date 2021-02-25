package tgz

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
)

func Test_Compress(t *testing.T) {
	src := "../assert.go"
	//src := "../../code"
	destFile := "../../code.tar.gz"
	err := compress(src, destFile)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("success")
}

func compress(src, destFile string) error {
	fw, err := os.Create(destFile)
	if err != nil {
		return errors.WithMessagef(err, "create %s failed", destFile)
	}
	defer fw.Close()
	// gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()
	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()
	return tarAll(src, tw)
}

func tarAll(src string, tw *tar.Writer) error {
	rootDirectory := filepath.Dir(src)
	walkFn := func(localpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relpath, err := filepath.Rel(rootDirectory, localpath)
		if err != nil {
			return err
		}
		packagepath := filepath.ToSlash(relpath)
		if info.IsDir() {
			header, err := tar.FileInfoHeader(info, localpath)
			header.Name = packagepath
			if err = tw.WriteHeader(header); err != nil {
				return err
			}
			return nil
		}
		err = tarFile(localpath, packagepath, tw)
		if err != nil {
			return fmt.Errorf("Error writing file to package: %s", err)
		}
		return nil
	}
	if err := filepath.Walk(src, walkFn); err != nil {
		return errors.WithMessage(err, "walk error")
	}
	return nil
}

func tarFile(localpath, packagepath string, tw *tar.Writer) error {
	fd, err := os.Open(localpath)
	if err != nil {
		return fmt.Errorf("%s: %s", localpath, err)
	}
	defer fd.Close()

	fi, err := fd.Stat()
	if err != nil {
		return fmt.Errorf("%s: %s", localpath, err)
	}

	header, err := tar.FileInfoHeader(fi, localpath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	header.Name = packagepath
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}
	if _, err = io.Copy(tw, bufio.NewReader(fd)); err != nil {
		return err
	}
	return nil
}
