package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello function returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// Selects and returns a random greeting message.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v!",
		"Howdy, %v!",
		"Hey, %v!",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
