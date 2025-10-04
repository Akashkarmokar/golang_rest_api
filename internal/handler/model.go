package handler

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

type NewsPostReqBody struct {
	Author    string   `json:"author"`
	Title     string   `json:"title"`
	Summary   string   `json:"summary"`
	CreatedAt string   `json:"created_at"`
	Content   string   `json:"content"`
	Source    string   `json:"source"`
	Tags      []string `jsong:"tags"`
}

func (n NewsPostReqBody) Validate() (errs error) {
	if n.Author == "" {
		errs = errors.Join(errs, fmt.Errorf("Author is empty %s", n.Author))
	}
	if n.Title == "" {
		errs = errors.Join(errs, fmt.Errorf("Title is empty %s", n.Author))
	}
	if n.Summary == "" {
		errs = errors.Join(errs, fmt.Errorf("Summary is empty %s", n.Summary))
	}
	if _, err := time.Parse(time.RFC3339, n.CreatedAt); err != nil {
		errs = errors.Join(errs, err)
	}
	if _, err := url.Parse(n.Source); err != nil {
		errs = errors.Join(errs, err)
	}

	if len(n.Tags) == 0 {
		errs = errors.Join(errs, errors.New("Tags can not be empty"))
	}
	return errs
}
