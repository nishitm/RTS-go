package main

import (
	"log"
	"time"

	"github.com/nishitm/RTS-go/config"
	"github.com/nishitm/RTS-go/src/gist"

	"github.com/nishitm/RTS-go/src/github"
	"github.com/nishitm/RTS-go/src/pastie/codepad"
	"github.com/nishitm/RTS-go/src/pastie/kpaste"
	"github.com/nishitm/RTS-go/src/pastie/paste"
	"github.com/nishitm/RTS-go/src/pastie/slexy"
	"github.com/nishitm/RTS-go/src/pastie/snipplr"
	"github.com/nishitm/RTS-go/src/reddit"
	"github.com/nishitm/RTS-go/src/twitter"

	"github.com/tkanos/gonfig"
)

func main() {
	configurations := config.Config{}
	err := gonfig.GetConf("config/config.json", &configurations)

	if err != nil {
		log.Print(err)
		return
	}

	twitterStart := false

	for range time.Tick(configurations.Interval * time.Second) {

		for _, src := range configurations.Sources {
			switch src {
			case "twitter":
				if !twitterStart {
					twitterStart = true
					twitterObj := &twitter.TwitterImplement{}
					go twitterObj.GetSearchedTerm(configurations)
				}
			case "github":
				githubObj := &github.GitImplement{}
				go githubObj.GetSearchedTerm(configurations)
			case "reddit":
				redditObj := &reddit.RedditImplement{}
				go redditObj.GetSearchedTerm(configurations)
			case "codepad":
				codepadObj := &codepad.CodepadImplement{}
				go codepadObj.GetSearchedTerm(configurations)
			case "slexy":
				slexyObj := &slexy.SlexyImplement{}
				go slexyObj.GetSearchedTerm(configurations)
			case "kpaste":
				kpasteObj := &kpaste.KpasteImplement{}
				go kpasteObj.GetSearchedTerm(configurations)
			case "snipplr":
				snipplrObj := &snipplr.SnipplrImplement{}
				go snipplrObj.GetSearchedTerm(configurations)
			case "paste":
				pasteObj := &paste.PasteImplement{}
				go pasteObj.GetSearchedTerm(configurations)
			case "gist":
				gistObj := &gist.GistImplement{}
				go gistObj.GetSearchedTerm(configurations)
			}
		}
	}
}
