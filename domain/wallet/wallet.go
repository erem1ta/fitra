package wallet

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

type Wallet struct {
	name        string
	balance     int
	canNegative bool
	uuid        uuid.UUID
	user        uuid.UUID
	note        string
}

func NewWallet(name string, balance int, canNegative bool, uuid uuid.UUID, userUUID uuid.UUID) (*Wallet, error) {
	if len(name) == 0 {
		return nil, errors.New("empty name")
	}
	if !canNegative && balance < 0 {
		return nil, errors.New("negative balance")
	}

	return &Wallet{
		name:        name,
		balance:     balance,
		canNegative: canNegative,
		uuid:        uuid,
		user:        userUUID,
	}, nil
}

func (w Wallet) Name() string {
	return w.name
}

func (w Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) UpdateBalance(amount int) error {
	if !w.canNegative && (w.balance+amount < 0) {
		return errors.New("can not set balance to negative")
	}

	w.balance = w.balance + amount
	return nil
}

func (w Wallet) UUID() uuid.UUID {
	return w.uuid
}

func (w Wallet) User() uuid.UUID {
	return w.user
}

func (w Wallet) Note() string {
	return w.note
}

func (w *Wallet) UpdateNote(note string) error {
	if len(note) > noteLimit {
		return errors.WithStack(ErrNoteTooLong)
	}

	w.note = note

	return nil
}
