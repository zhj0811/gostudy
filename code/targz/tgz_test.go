package targz

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
)

func Test_compressTGZ(t *testing.T) {
	err := compressTGZ("test.tar.gz", "tgz_test.go", "../tgz")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("Success")
}

func Test_uncompressTGZ(t *testing.T) {
	err := uncompressTGZ("test.tar.gz", "test")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("Success")
}

func compressTGZ(destFile string, srcs ...string) error {
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
	for _, src := range srcs {
		if err := tarAll(src, tw); err != nil {
			return err
		}
	}
	return nil
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
			return errors.WithMessage(err, "write file to package failed")
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
		return errors.WithMessagef(err, "open file %s failed", localpath)
	}
	defer fd.Close()

	fi, err := fd.Stat()
	if err != nil {
		return errors.WithMessagef(err, "file %s stat error", localpath)
	}

	header, err := tar.FileInfoHeader(fi, localpath)
	if err != nil {
		return errors.WithMessagef(err, "create %s header failed", localpath)
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

func uncompressTGZ(tgzFile, destDir string) error {
	fr, err := os.Open(tgzFile)
	if err != nil {
		return errors.Errorf("open %s failed", tgzFile)
	}
	defer fr.Close()

	// gzip read
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return errors.WithMessage(err, "gzip reader error")
	}
	defer gr.Close()

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
		if err != nil {
			return err
		}
		defer fw.Close()
		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			return err
		}
	}
	return nil
}
