
package utils

import (
	"log"
	"strings"

	"github.com/google/uuid"
)

const (
	VWONamespace := uuid.NewV5(uuid.NamespaceURL, "https://vwo.com")
)

func GenerateFor(userID string, accountID string) string {
	userIdNamespace := Generate(VWONamespace, accountID)
	uuidForAccountUserId := Generate(userIdNamespace, userID)
	desiredUUID = strings.ToUpper(strings.Replace(string(uuidForAccountUserId), "-", "", -1)) //To be confirmed
	log.Println("Log For User")
	return desiredUUID
}

func Generate(namespace uuid.UUID, name string) uuid.UUID {
	return uuid.NewV5(namespace, name)
}
