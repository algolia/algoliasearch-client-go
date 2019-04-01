package search

import "time"

type CreateKeyRes struct {
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"createdAt"`
	wait      func() error
}

type RestoreKeyRes struct {
	CreatedAt time.Time `json:"createdAt"`
	wait      func() error
}

type UpdateKeyRes struct {
	Key       string    `json:"key"`
	UpdatedAt time.Time `json:"updatedAt"`
	wait      func() error
}

type DeleteKeyRes struct {
	DeletedAt time.Time `json:"deletedAt"`
	wait      func() error
}

type ListAPIKeysRes struct {
	Keys []Key `json:"keys"`
}

func (r CreateKeyRes) Wait() error  { return r.wait() }
func (r RestoreKeyRes) Wait() error { return r.wait() }
func (r UpdateKeyRes) Wait() error  { return r.wait() }
func (r DeleteKeyRes) Wait() error  { return r.wait() }
