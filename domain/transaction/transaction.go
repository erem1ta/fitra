package transaction

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	noteLimit int = 500
)

var (
	ErrNoteTooLong = errors.New("note too long")
)

type Transaction struct {
	uuid     uuid.UUID
	name     string
	amount   int
	user     uuid.UUID
	from     uuid.UUID
	to       uuid.UUID
	category uuid.UUID
	note     string
}

// NewTransaction create transaction with required fields
func NewTransaction(uuid uuid.UUID, name string, amount int, userUUID uuid.UUID, from uuid.UUID, to uuid.UUID) (*Transaction, error) {
	// transaction valid check
	if len(name) == 0 {
		return nil, errors.New("empty name")
	}
	if amount == 0 {
		return nil, errors.New("0 amount")
	}

	// create transaction
	return &Transaction{
		uuid:   uuid,
		name:   name,
		amount: amount,
		user:   userUUID,
		from:   from,
		to:     to,
	}, nil
}

func (t Transaction) UUID() uuid.UUID {
	return t.uuid
}

func (t Transaction) User() uuid.UUID {
	return t.user
}

func (t Transaction) From() uuid.UUID {
	return t.from
}

func (t Transaction) To() uuid.UUID {
	return t.to
}

func (t Transaction) Category() uuid.UUID {
	return t.category
}

func (t *Transaction) AddCategory(category uuid.UUID) {
	t.category = category
}

func (t Transaction) Note() string {
	return t.note
}

func (t *Transaction) UpdateNote(note string) error {
	if len(note) < noteLimit {
		return errors.WithStack(ErrNoteTooLong)
	}

	t.note = note

	return nil
}
