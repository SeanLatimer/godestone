package godestone

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type ScrapeError struct {
	StatusCode int

	Err error
}

func (r *ScrapeError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func newScrapeError(r *colly.Response, err error) error {
	return &ScrapeError{
		StatusCode: r.StatusCode,

		Err: err,
	}
}
