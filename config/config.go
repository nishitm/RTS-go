package config

import "time"

// Config struct contains all the source configurations
type Config struct {
	Twitter  TwitterStruct `json:twitter`
	Reddit   RedditStruct  `json:reddit`
	Github   GithubStruct  `json:github`
	Codepad  CodepadStruct `json:codepad`
	Slexy    SlexyStruct   `json:slexy`
	Kpaste   KpasteStruct  `json:kpaste`
	Snipplr  SnipplrStruct `json:snipplr`
	Interval time.Duration `json:interval` //This interval is consecutive scan timings in seconds
	Sources  []string      `json:sources`
}

// TwitterStruct is the struct for twitter configurations
type TwitterStruct struct {
	ConsumerKey    string   `json:consumerKey`
	ConsumerSecret string   `json:consumerSecret`
	AccessToken    string   `json:accessToken`
	AccessSecret   string   `json:accessSecret`
	SearchTerms    []string `json:searchTerms`
}

// RedditStruct is the struct for reddit configurations
type RedditStruct struct {
	URL         string   `json:url`
	Interval    string   `json:interval` //This interval is the time interval till before the daemon will search the reddit posts
	SearchTerms []string `json:searchTerms`
}

// GithubStruct is the struct for the github configurations
type GithubStruct struct {
	SearchTerms []string `json:searchTerms`
}

// CodepadStruct is the struct for codepad.com pastie website configurations
type CodepadStruct struct {
	URL         string   `json:url`
	SearchTerms []string `json:searchTerms`
}

// SlexyStruct is the struct for the https://slexy.org/ pastie website configurations
type SlexyStruct struct {
	URL         string   `json:url`
	URLBase     string   `json:urlbase`
	Regex       string   `json:regex`
	SearchTerms []string `json:searchTerms`
}

// KpasteStruct is the struct for the https://kpaste.net/ pastie website configurations
type KpasteStruct struct {
	URL         string   `json:url`
	URLBase     string   `json:urlbase`
	Regex       string   `json:regex`
	SearchTerms []string `json:searchTerms`
}

// SnipplrStruct is the struct for the https://snipplr.com pastie website configurations
type SnipplrStruct struct {
	URL         string   `json:url`
	URLBase     string   `json:urlbase`
	Regex       string   `json:regex`
	SearchTerms []string `json:searchTerms`
}
