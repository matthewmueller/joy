// Package stats provides CLI analytics.
package stats

import (
	"time"

	"github.com/apex/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	analytics "github.com/matthewmueller/firehose-analytics"
	"github.com/matthewmueller/joy/internal/env"
	"github.com/matthewmueller/store"
	"github.com/pkg/errors"
	"github.com/rapidloop/skv"
)

// Client sets up the AWS tracking client
var Client = func() *analytics.Analytics {
	e := env.Get()

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			e.AWSAccessKey,
			e.AWSSecretKey,
			"",
		),
		Region: aws.String(e.AWSRegion),
	})
	if err != nil {
		log.Debugf("aws session not found")
		sess = nil
	}

	return analytics.New(&analytics.Config{
		Stream:  e.FirehoseStream,
		Session: sess,
		Prefix:  "joy:",
		Dir:     "joy",
	})
}()

// Increment runs
func Increment(db *skv.KVStore) (runs int, err error) {
	// defer log.Trace("increment").Stop(&err)

	if err := db.Get("runs", &runs); err != nil && err != store.ErrNotFound {
		return runs, errors.Wrapf(err, "error getting 'first time' field from the db")
	}

	// increment
	runs++

	if err := db.Put("runs", runs); err != nil {
		return runs, errors.Wrapf(err, "error setting 'first time' field")
	}

	return runs, nil
}

// Track event `name` with optional `props`.
func Track(name string, fields map[string]interface{}) error {
	return Client.Track(name, fields)
}

// TrackError tracks event `name` with errors
func TrackError(name string, now time.Time, err *error) error {
	if *err == nil {
		return nil
	}

	return Client.Track(name, map[string]interface{}{
		"duration": time.Since(now).Round(time.Millisecond).String(),
		"error":    (*err).Error(),
	})
}
