package evotor

type PagedResponse[T any] struct {
	Items  []T `json:"items"`
	Paging struct {
		NextCursor string `json:"next_cursor"`
	} `json:"paging"`
}

func (p *PagedResponse[T]) HasNext() bool {
	return p.Paging.NextCursor != ""
}
