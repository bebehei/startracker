package main

import "context"
import "fmt"
import "log"
import "os"
import "strconv"

import "github.com/google/go-github/github"

func main() {

	gh := github.NewClient(nil)
	ctx := context.Background()

	var u string = ""
	var r string = ""

	if len(os.Args) > 2 {
		u = os.Args[1]
		r = os.Args[2]
	} else if len(os.Args) > 0 {
		log.Fatal("User and Repository name not specified!")
	}

	var opts = github.ListOptions {
		Page: 0,
		PerPage: 1000, /* This option is intended to just set the max of GitHub */
	}

	var maxPage int = 1

	for opts.Page <= maxPage {

		stars, response, e := gh.Activity.ListStargazers(ctx, u, r, &opts)

		if e != nil {
			log.Fatal(e)
		}

		for _, star := range stars {
			user := star.GetUser()
			time := star.GetStarredAt()

			fmt.Println(strconv.FormatInt(time.Time.Unix(), 10) + "\t" + user.GetLogin())
		}

		if opts.Page == maxPage {
			break;
		}

		maxPage = response.LastPage
		opts.Page = response.NextPage
	}

}

