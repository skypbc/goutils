package guuid

import (
	"crypto/rand"

	"github.com/google/uuid"
)

type UUID = uuid.UUID

var Nil = uuid.Nil
var Max = uuid.Max

var NewUuid = newUuid
var NewRandom = newRandom
var Parse = uuid.Parse
var MustParse = uuid.MustParse
var FromBytes = uuid.FromBytes

func newUuid() (UUID, error) {
	return uuid.NewV7()
}

func newRandom() (uuid UUID) {
	rand.Read(uuid[:]) //nolint:errcheck
	return
}
