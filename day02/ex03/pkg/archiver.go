package archiver

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
)

func CompressToTarGz(filename string, directory string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	timestamp := strconv.FormatInt(int64(info.ModTime().Unix()), 10)
	var fullPath string
	if directory != "" {
		fullPath = path.Join(directory, filename+"_"+timestamp+".tar.gz")
	} else {
		fullPath = filename + "_" + timestamp + ".tar.gz"
	}

	archive, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer archive.Close()

	writer, err := gzip.NewWriterLevel(archive, gzip.BestCompression)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer writer.Close()

	tw := tar.NewWriter(writer)
	defer tw.Close()

	hdr := &tar.Header{
		Name: path.Base(filename),
		Mode: int64(0644),
		Size: int64(info.Size()),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		fmt.Println(err)
	}

	if _, err := io.Copy(tw, file); err != nil {
		fmt.Println(err)
		return
	}
}
