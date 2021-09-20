package wallet

import "github.com/google/uuid"

type Wallet struct {
	name          string
	balance       int
	canBeNegative bool
	uuid          uuid.UUID
	userUUID      uuid.UUID
	note          string
}
