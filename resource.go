package resource

import (
	"compress/gzip"
	"crypto"
	"fmt"
	"github.com/gosexy/checksum"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

const PS = string(os.PathSeparator)

// Size of the first directory chunk.
var Offset = 8

// Hasing method.
var HashMethod = crypto.SHA1

// Optional salt.
var Salt = ""

/*
	Given an URL returns a local *os.File.
*/
func Allocate(uri string, basepath string) (*os.File, error) {
	local, err := LocalPath(uri, basepath)
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(path.Dir(local), os.ModeDir|0755)
	if err != nil {
		return nil, err
	}
	return os.Create(local)
}

/*
	Given an URL returns a local file path.
*/
func LocalPath(uri string, basepath string) (string, error) {
	stat, err := os.Stat(basepath)

	if err == nil {
		if stat.IsDir() == false {
			return "", fmt.Errorf("Path %s is a file, not a directory.", basepath)
		}
	}

	data, _ := url.Parse(uri)

	basename := path.Base(data.Path)

	hash := checksum.String(uri+Salt, HashMethod)

	return strings.TrimRight(basepath, PS) + PS + strings.TrimLeft(strings.Join([]string{hash[0:Offset], hash[Offset:], basename}, PS), PS), nil
}

/*
	Downloads the given URI to a file into the base directory.
*/
func Download(uri string, basepath string) (string, error) {

	var req *http.Request
	var err error

	req, err = http.NewRequest("GET", uri, nil)

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

	fp, err := Allocate(uri, basepath)

	if err != nil {
		return "", err
	}

	defer fp.Close()

	io.Copy(fp, resp.Body)

	return fp.Name(), nil
}
