package models

type Hash struct {
	ID int

	Status HashStatus
	Hash   int
	Error  string
}

type HashStatus string

const (
	PendindStatus  HashStatus = "pending"
	FinishedStatus HashStatus = "finished"
)

func (h Hash) IsPending() bool {
	return h.Status == PendindStatus
}
