package evotor

type pagedResponse[T any] struct {
	Items  []T `json:"items"`
	Paging struct {
		NextCursor string `json:"next_cursor"`
	} `json:"paging"`
}

func (p *pagedResponse[T]) HasNext() bool {
	return p.Paging.NextCursor != ""
}
