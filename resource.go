package resource

import (
	"compress/gzip"
	"crypto"
	"errors"
	"io"
	"menteslibres.net/gosexy/checksum"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

// Path separator string
const pathSeparator = string(os.PathSeparator)

// Hasing method.
var hashingFunc = crypto.SHA1

var (
	ErrNotADirectory = errors.New(`Path %s is a file, expecting a directory.`)
)

// Creates a local *os.File based on the given URI. The user is responsible for
// closing the file.
func allocate(uri string, basepath string) (*os.File, error) {
	var local string
	var err error

	if local, err = localPath(uri, basepath); err != nil {
		return nil, err
	}

	if err = os.MkdirAll(path.Dir(local), os.ModeDir|0755); err != nil {
		return nil, err
	}

	return os.Create(local)
}

// Returns a unique file path for a given URL.
func localPath(uri string, basepath string) (string, error) {
	var stat os.FileInfo
	var err error
	var urlData *url.URL

	stat, err = os.Stat(basepath)

	if err == nil {
		// Path exists.
		if stat.IsDir() == false {
			// Path is not a directory.
			return "", ErrNotADirectory
		}
	}

	if urlData, err = url.Parse(uri); err != nil {
		return "", err
	}

	basename := path.Base(urlData.Path)

	hash := checksum.String(uri, hashingFunc)

	return strings.TrimRight(basepath, pathSeparator) + pathSeparator + strings.TrimLeft(strings.Join([]string{hash[0:4], hash[4:8], hash[8:12], hash[12:], basename}, pathSeparator), pathSeparator), nil
}

// Downloads the given URI into a local file and returns the file path.
func Download(uri string, basepath string) (string, error) {
	var req *http.Request
	var resp *http.Response
	var err error
	var fp *os.File

	if req, err = http.NewRequest("GET", uri, nil); err != nil {
		return "", err
	}

	client := &http.Client{}

	if resp, err = client.Do(req); err != nil {
		return "", err
	}

	if resp.Header.Get("Content-Encoding") == "gzip" {
		if resp.Body, err = gzip.NewReader(resp.Body); err != nil {
			return "", err
		}
	}

	defer resp.Body.Close()

	if fp, err = allocate(uri, basepath); err != nil {
		return "", err
	}

	defer fp.Close()

	if _, err = io.Copy(fp, resp.Body); err != nil {
		return "", err
	}

	return fp.Name(), nil
}
