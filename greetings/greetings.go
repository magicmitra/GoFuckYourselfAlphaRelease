package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}

	// Create a message using a random format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos return a map that associates each of the named people
// with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// This is how you make a map -> make(map[key-type]val-type)
	messages := make(map[string]string)

	// Loop through the received slice of names,
	// calling the Hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved messages
		// with the name
		messages[name] = message
	}
	return messages, nil
}

// init sets the initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages.
// The returned message is selected at random
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Fuck you, %v",
		"Stick this in your ass, %v",
		"Hey %v, suck my cock",
	}

	// return a randomly selected message format by specifying
	// a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
