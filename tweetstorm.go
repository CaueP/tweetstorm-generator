// Author: Cauê Garcia Polimanti
// Date: 17/06/2017

package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Tweet lenght constant
const tweetLenght int = 140

// convertTextIntoTweets Splits a text into an 140 characters strings array
// It returns a slice of strings containing the Tweets with its respective prefix
func convertTextIntoTweets(text string) []string {

	var tweets []string

	textLenght := len(text)
	if textLenght > tweetLenght {
		// creating a string reader
		reader := strings.NewReader(text)

		totalTweets, maxIndexTextSize := calculateTweetTotalAndIndexSize(textLenght)
		readingOffset := tweetLenght - maxIndexTextSize
		for i := 0; i < totalTweets; i++ {
			// Creates the string index text
			tweetIndex := strconv.Itoa(i+1) + "/" + strconv.Itoa(totalTweets) + " "

			offset, _ := reader.Seek(int64(i*readingOffset), 0)

			var textPart []byte

			if int(offset)+readingOffset > textLenght {
				textPart = make([]byte, int64(textLenght)-offset)
			} else {
				textPart = make([]byte, readingOffset)
			}
			// textPart := make([]byte, readingOffset > textLenght? textLenght: readingOffset)

			// set io to read at least 1 char
			io.ReadAtLeast(reader, textPart, 1)

			// concatenates the index with the actual text to form a tweet
			tweet := tweetIndex + string(textPart)

			// append the tweet to the array of tweets
			tweets = append(tweets, tweet)
		}
	} else {
		tweets = append(tweets, text)
	}
	return tweets
}

// calculateTweetTotalAndIndexSize Calculates the number of tweets and the number of chars on the index
// It returns the total number of tweets that a text will generate and the number of characters needed for prefixing each tweet
func calculateTweetTotalAndIndexSize(textLenght int) (totalTweets int, maxIndexTextSize int) {
	// calculate the number of total tweets
	totalTweets = textLenght / tweetLenght

	// Calculate the tweet index text size for "x/y " format. For example: 1/4, then 2/4 then 3/4 then 4/4
	// Method: multiply the number of chars on the total string by 2 to consider the current and total of parts.
	// add 2 to consider both  "/" and " " characters
	maxIndexTextSize = (len(strconv.Itoa(totalTweets)) * 2) + 2

	// recalculate the totalTweets considering the prefix text for each tweet
	totalTweets = (textLenght + maxIndexTextSize*totalTweets) / tweetLenght
	if textLenght%tweetLenght != 0 {
		totalTweets++
	}

	return totalTweets, maxIndexTextSize
}

// printTweets Print the tweets on a tweet slice, from the last to the first part
func printTweets(tweets []string) {
	for i := len(tweets) - 1; i >= 0; i-- {
		fmt.Println(tweets[i])
	}
}

// Error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}
