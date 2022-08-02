package blogposts

import (
	"bufio"
	"bytes"
	"errors"
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

//Why not use postFile fs.File instead of postFile io.Reader?
//We only want to use the ReadAll function, so it's better to be more flexible and use the io.Reader interface.
func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	post := Post{}

	//load the properties into our post object
	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			break
		}

		parameter := strings.Split(line, ": ")
		switch parameter[0] {
		case "Title":
			post.Title = parameter[1]
		case "Description":
			post.Description = parameter[1]
		case "Tags":
			post.Tags = strings.Split(parameter[1], ", ")
		default:
			return Post{}, errors.New(fmt.Sprintf("invalid parameter: %v", line))
		}
	}

	post.Body = readBody(scanner)
	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	body := bytes.Buffer{}

	for scanner.Scan() {
		fmt.Fprintln(&body, scanner.Text())
	}
	return strings.TrimSuffix(body.String(), "\n")
}
