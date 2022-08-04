package blogrenderer_test

import (
	"blogrenderer"
	"bytes"
	"github.com/approvals/go-approval-tests"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := blogrenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		//this will test by comparing the output of our function with the approved file (of the same name as our test)
		approvals.VerifyString(t, buf.String())
	})
}
