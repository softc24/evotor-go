package helpers

import (
	"context"
	"fmt"
	"iter"
	"net/http"
	"net/url"

	"github.com/capcom6/go-restkit"
	"github.com/softc24/evotor-go"
)

func GetPagedData[T any](
	ctx context.Context,
	client *restkit.Client,
	path string,
	params url.Values,
	headers http.Header,
) (iter.Seq[T], func() error) {
	var err error
	return func(yield func(T) bool) {
			for {
				res := new(evotor.PagedResponse[T])

				if err = client.Do(
					ctx,
					http.MethodGet,
					fmt.Sprintf("%s?%s", path, params.Encode()),
					headers,
					nil,
					res,
				); err != nil {
					return
				}

				for _, item := range res.Items {
					if !yield(item) {
						return
					}
				}

				if res.HasNext() {
					params = url.Values{
						"cursor": []string{res.Paging.NextCursor},
					}
				} else {
					return
				}
			}
		}, func() error {
			if err == nil {
				return nil
			}

			return fmt.Errorf("failed to get paged data: %w", err)
		}
}
