package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

//Why not use postFile fs.File instead of postFile io.Reader?
//We only want to use the ReadAll function, so it's better to be more flexible and use the io.Reader interface.
func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func(tagName string) string {
		scanner.Scan()
		if len(scanner.Text()) != 0 {
			return strings.TrimPrefix(scanner.Text(), tagName)
		}
		return ""
	}

	return Post{
		Title:       readLine(titleSeparator),
		Description: readLine(descriptionSeparator),
		Tags:        strings.Split(readLine(tagsSeparator), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	body := bytes.Buffer{}
	scanner.Scan() //ignore "---"

	for scanner.Scan() {
		fmt.Fprintln(&body, scanner.Text())
	}
	return strings.TrimSuffix(body.String(), "\n")
}
