// Author: Cauê Garcia Polimanti
// Date: 17/06/2017

package main

import (
	"os"
	"strings"
)

func main() {
	// text received from args
	var text string
	// array of tweets
	var tweets []string

	// checking args
	args := os.Args[1:]
	if len(args) >= 1 {
		switch os.Args[1] {
		case "--help":
			showHelp()
			return
		default:
			text = strings.Join(os.Args[1:], " ")
		}
	} else {
		showUsageInstructions()
		return
	}

	tweets = convertTextIntoTweets(text)
	printTweets(tweets)
}
