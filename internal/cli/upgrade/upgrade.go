package upgrade

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/apex/log"
	"github.com/google/go-github/github"
	"github.com/matthewmueller/joy/internal/stats"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// New build command
func New(ctx context.Context, root *kingpin.Application, version string) {
	cmd := root.Command("upgrade", "upgrade joy to the latest version")
	cmd.Action(func(_ *kingpin.ParseContext) (err error) {
		defer stats.TrackError("upgrade", time.Now(), &err)
		return upgrade(ctx, version)
	})
}

// https://github.com/apex/apex/blob/master/upgrade/upgrade.go
// Thanks TJ
func upgrade(ctx context.Context, version string) (err error) {
	start := time.Now()
	log.Infof("current release is %s", version)

	// fetch releases
	gh := github.NewClient(nil)
	releases, _, err := gh.Repositories.ListReleases(ctx, "matthewmueller", "joy", nil)
	if err != nil {
		return err
	}

	// see if it's new
	latest := releases[0]
	log.Infof("latest release is %s", *latest.TagName)
	latestVersion := (*latest.TagName)[1:]

	// stats
	defer func() {
		if err == nil {
			stats.Track("upgrade", map[string]interface{}{
				"duration": time.Since(start).Round(time.Millisecond).String(),
				"from":     version,
				"to":       latestVersion,
			})
		}
	}()

	// TODO: ignore if semver
	if latestVersion == version {
		log.Infof("you're up to date :)")
		return nil
	}

	asset := findAsset(latest)
	if asset == nil {
		return errors.New("cannot find binary for your system")
	}

	// get the executable's path
	cmdPath, err := exec.LookPath("joy")
	if err != nil {
		return err
	}
	cmdDir := filepath.Dir(cmdPath)

	// create tmp file
	tmpPath := filepath.Join(cmdDir, "joy-upgrade")
	f, err := os.OpenFile(tmpPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	// download binary
	log.Infof("downloading %s", *asset.BrowserDownloadURL)
	res, err := http.Get(*asset.BrowserDownloadURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// copy it down
	_, err = io.Copy(f, res.Body)
	if err != nil {
		return err
	}

	// replace it
	log.Infof("replacing %s", cmdPath)
	err = os.Rename(tmpPath, cmdPath)
	if err != nil {
		return err
	}

	log.Infof("visit https://github.com/matthewmueller/joy/releases for the changelog")
	return nil
}

// findAsset returns the binary for this platform.
func findAsset(release *github.RepositoryRelease) *github.ReleaseAsset {
	for _, asset := range release.Assets {
		if *asset.Name == fmt.Sprintf("joy_%s_%s", runtime.GOOS, runtime.GOARCH) {
			return &asset
		}
	}
	return nil
}
