package util

import "github.com/google/uuid"

func CreateUUID() (string, error) {
	uuID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uuID.String(), nil
}
