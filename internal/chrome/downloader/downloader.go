package downloader

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/mholt/archiver"
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

// Find locally or remotely
func Find(dir string) (string, error) {
	filepath, err := chromiumPath(dir)
	if err != nil {
		return "", errors.Wrapf(err, "error getting chromium path")
	}

	if _, err := os.Stat(filepath); !os.IsNotExist(err) {
		return filepath, nil
	}

	if err := download(dir); err != nil {
		return "", errors.Wrapf(err, "error downloading")
	}

	return filepath, nil
}

// Download path
func download(dir string) error {
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

	folderPath := path.Join(dir, platform+"-"+ChromiumRevision)
	if err := os.MkdirAll(folderPath, 0775); err != nil {
		return errors.Wrapf(err, "error making directory")
	}

	return archiver.Zip.Read(res.Body, folderPath)
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
