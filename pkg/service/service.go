package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(fpath string, url string) error {

	fmt.Println(fpath, url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(fpath)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
