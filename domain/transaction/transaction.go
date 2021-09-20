package transaction

import "github.com/google/uuid"

type Transaction struct {
	name     string
	amount   int
	uuid     uuid.UUID
	userUUID uuid.UUID
	from     uuid.UUID
	to       uuid.UUID
	category uuid.UUID
	note     string
}
