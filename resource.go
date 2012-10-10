package resource

import (
	"compress/gzip"
	"crypto"
	"github.com/gosexy/checksum"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

const PS = string(os.PathSeparator)

var Root = "downloads" + PS

func Allocate(addr string) (*os.File, error) {
	local := Normalize(addr)
	createDirectories(local)
	return os.Create(local)
}

func createDirectories(local string) {
	os.MkdirAll(path.Dir(local), os.ModeDir|0755)
}

func Normalize(addr string) string {
	data, _ := url.Parse(addr)

	basename := path.Base(data.Path)

	hash := checksum.String(addr, crypto.SHA1)

	return Root + PS + strings.Join([]string{hash[0:3], hash[3:], basename}, PS)
}

func Download(addr string) (string, error) {

	var req *http.Request
	var err error

	req, err = http.NewRequest("GET", addr, nil)

	if err != nil {
		return "", err
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	if resp.Header.Get("Content-Encoding") == "gzip" {
		resp.Body, err = gzip.NewReader(resp.Body)
		if err != nil {
			return "", err
		}
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	// TODO: Simultaneous reader-writer

	file, err := Allocate(addr)

	if err != nil {
		return "", err
	}

	defer file.Close()

	file.Write(bytes)

	return file.Name(), nil
}
