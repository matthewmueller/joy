package prompt

import (
	"fmt"
	"strings"
	"time"

	"github.com/matthewmueller/joy/internal/typewriter"
	"github.com/matthewmueller/store"
	isatty "github.com/mattn/go-isatty"
	"github.com/pkg/errors"
	"github.com/rapidloop/skv"
	prompt "github.com/tj/go-prompt"
	govalidator "gopkg.in/asaskevich/govalidator.v4"
)

// Prompt does special things depending on
// the number of times Joy is run
func Prompt(db *skv.KVStore) (done bool, err error) {
	// ignore when not a tty
	if !isatty.IsTerminal(0) {
		return false, nil
	}
	return false, nil
}

// Introduce Joy to first timers
func introduce() error {
	fmt.Println(splash())

	typewriter.Type(strings.TrimSpace(`
Thanks for trying the Joy Compiler!

To get started, create a "main.go" file with the following code:

  package main

  func main() {
    println("Joy to the world!")
  }

Next, run "joy run main.go" to see the result.
Then, run "joy main.go" to see the compiled Javascript.

That's all there is to it!

More tips:

    • Run "joy help" to see what else Joy can do for you
    • Visit mat.tm/joy to read more about Joy's origins
    • Chat with us in #joy-compiler on Slack at gophers.slack.com
    • Star github.com/matthewmueller/joy to follow the development
    • Follow twitter.com/mattmueller for project updates

Joy was a labor of love for me. I hope you'll find Joy as
delightful as I do.

Before jumping into Joy, I had a quick question:

	`), 20*time.Second)
	fmt.Printf(string('\n'))
	fmt.Printf(string('\n'))

ask:
	referrer := prompt.StringRequired(" • How did you hear about Joy? ")
	if referrer == "" {
		goto ask
	}
	fmt.Printf("\nThanks a bunch!\n")

	return nil
}

func promptEmail(db *skv.KVStore) error {
	var email string
	if err := db.Get("email", &email); err != nil && err != store.ErrNotFound {
		return errors.Wrapf(err, "error getting email from store")
	} else if email == "" {
		fmt.Println(splash())

		typewriter.Type(strings.TrimSpace(`
Woah nelly! I hope you're enjoying the Joy compiler so far. Are you ready to
take this relationship to the next level?

I wanted to ask you for your name and email address.

I'd like to be able to reach you when we add features, write tutorials and
spotlight websites using Joy in production. I'm really hoping we can build
a kick-butt community around frontend Go development.

I hate spam as much as you do and if the emails you receive aren’t helpful,
I’d love to hear how I can improve them. As always, you'll have an option
to unsubscribe in each email you receive.
		`), 15*time.Second)
		fmt.Printf(string('\n'))
		fmt.Printf(string('\n'))

	askName:
		name := prompt.StringRequired(" • What is your name? ")
		if len(name) <= 2 {
			goto askName
		}

	askEmail:
		email := prompt.StringRequired(" • What is your email address? ")
		if !govalidator.IsEmail(email) {
			goto askEmail
		}

		if err := db.Put("name", email); err != nil {
			return errors.Wrapf(err, "error storing name")
		}

		if err := db.Put("email", email); err != nil {
			return errors.Wrapf(err, "error storing email")
		}

		fmt.Printf("\nThanks bud! Now where were we")
		typewriter.Type("...", 2*time.Second)
		fmt.Printf(" Ah yes!\n\n")
		return nil
	}

	return nil
}

func splash() string {
	return `
     ________  __
 __ / / __ \ \/ /
/ // / /_/ /\  /   A Delightful Go to Javascript Compiler
\___/\____/ /_/
`
}
