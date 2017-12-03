package chrome

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
)

// ChromiumRevision chromium revision
const ChromiumRevision = "518818"

// DownloadHost download host
const DownloadHost = "https://storage.googleapis.com"

var downloadURLs = map[string]string{
	"linux": "%s/chromium-browser-snapshots/Linux_x64/%s/chrome-linux.zip",
	"mac":   "%s/chromium-browser-snapshots/Mac/%s/chrome-mac.zip",
	"win32": "%s/chromium-browser-snapshots/Win/%s/chrome-win32.zip",
	"win64": "%s/chromium-browser-snapshots/Win_x64/%s/chrome-win32.zip",
}

// Exists checks if chrome exists on our system or not
func Exists(dir string) (string, error) {
	filepath, err := chromiumPath(dir)
	if err != nil {
		return "", errors.Wrapf(err, "error getting chromium path")
	}

	if _, err := os.Stat(filepath); !os.IsNotExist(err) {
		return filepath, nil
	}

	return "", nil
}

// Download path
func Download(dir string) error {
	var url, platform string

	switch runtime.GOOS {
	case "darwin":
		platform = "mac"
		url = fmt.Sprintf(downloadURLs[platform], DownloadHost, ChromiumRevision)
	case "linux":
		platform = "linux"
		url = fmt.Sprintf(downloadURLs[platform], DownloadHost, ChromiumRevision)
	case "windows":
		platform = "windows"
		// TODO: 32 vs 64
		url = fmt.Sprintf(downloadURLs["win64"], DownloadHost, ChromiumRevision)
	default:
		return errors.New("unsupported operating system")
	}

	res, err := http.Get(url)
	if err != nil {
		return errors.Wrapf(err, "error getting download")
	}
	defer res.Body.Close()

	src, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	folderPath := path.Join(dir, platform+"-"+ChromiumRevision)
	if err := os.MkdirAll(folderPath, 0775); err != nil {
		return errors.Wrapf(err, "error making directory")
	}

	return unzip(src, folderPath)
}

// Path to the executable
// TODO: test on windows
func chromiumPath(dir string) (string, error) {
	var platform string
	switch runtime.GOOS {
	case "darwin":
		platform = "mac"
		return path.Join(dir, platform+"-"+ChromiumRevision, "chrome-mac", "Chromium.app", "Contents", "MacOS", "Chromium"), nil
	case "linux":
		platform = "linux"
		return path.Join(dir, platform+"-"+ChromiumRevision, "chrome-linux", "chrome"), nil
	case "windows":
		platform = "windows"
		return path.Join(dir, platform+"-"+ChromiumRevision, "chrome-win32", "chrome.exe"), nil
	default:
		return "", errors.New("unsupported operating system")
	}
}

// Unzip a file to a destination and preserve the symlinks
// Based on http://stackoverflow.com/a/24792688/842097 with symlink additions
func unzip(src []byte, dest string) error {
	rdr := bytes.NewReader(src)

	r, err := zip.NewReader(rdr, rdr.Size())
	if err != nil {
		return err
	}

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else if f.FileInfo().Mode()&os.ModeSymlink != 0 {
			buffer := make([]byte, f.FileInfo().Size())
			size, err := rc.Read(buffer)
			if err != nil {
				return err
			}

			target := string(buffer[:size])

			err = os.Symlink(target, path)
			if err != nil {
				return err
			}
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err = f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
