package main

import "context"
import "fmt"
import "log"
import "strconv"

import "github.com/google/go-github/github"

var u string = "dunst-project"
var r string = "dunst"

func main() {

	gh := github.NewClient(nil)
	ctx := context.Background()

	var opts = github.ListOptions {
		Page: 0,
		PerPage: 1000, /* This option is intended to just set the max of GitHub */
	}

	var maxPage int = 0

	for opts.Page <= maxPage {

		stars, response, e := gh.Activity.ListStargazers(ctx, u, r, &opts)

		if e != nil {
			log.Fatal(e)
		}

		maxPage = response.LastPage
		opts.Page = response.NextPage

		for _, star := range stars {
			user := star.GetUser()
			time := star.GetStarredAt()

			fmt.Println(strconv.FormatInt(time.Time.Unix(), 10) + "\t" + user.GetLogin())
		}
	}

}

