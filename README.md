# Startracker

Tracking GitHub Stars via git

## Rationale

Stars on a project are a smart variable to track the project's popularity. Albeit the amount of stars can decrease, it usually only increases. It's the same as a "Like" on facebook. People usually don't "unlike" pages, they just get off the radar.

So in [dunst-project](https://github.com/dunst-project/dunst/), where I develop for, stars often started to flap for a longer time. It's interesting for me to get access to the actual data and see, why and what happened.

Let's record the history to be later able to trace back the reasons.

# Usage

(This is my first golang project. It's not the best, but it works ;-).)

1. `go get`
1. `go build`
1. `./startracker <user> <repo> > data.txt`

If you track the file `data.txt` with git on every invocation of startracker, you can record the full history.

Keep in Mind: GitHub currently has a rate limit for 60 unauthenticated requests per hour per IPv4. 
