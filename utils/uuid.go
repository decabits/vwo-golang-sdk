package utils

import (
	"strconv"
	"strings"

	"github.com/decabits/vwo-golang-sdk/schema"
	guuid "github.com/google/uuid"
	suuid "github.com/satori/go.uuid"
)

// generateFor generates desired UUID
func generateFor(vwoInstance schema.VwoInstance, userID string, accountID int) string {
	/*
		Args:
		    userID : User identifier
		    accountID : Account identifier

		Returns:
			string : Desired Uuid
	*/
	NameSpaceURL, _ := guuid.Parse("6ba7b811-9dad-11d1-80b4-00c04fd430c8")
	VWONamespace := suuid.NewV5(suuid.UUID(NameSpaceURL), "https://vwo.com")
	userIDNamespace := suuid.NewV5(VWONamespace, strconv.Itoa(accountID))
	uuidForAccountUserID := suuid.NewV5(userIDNamespace, userID)
	desiredUUID := strings.ToUpper(strings.Replace(uuidForAccountUserID.String(), "-", "", -1)) //To be confirmed
	vwoInstance.Logger.Info("DEBUG_MESSAGES.UUID_FOR_USER " + desiredUUID)
	return desiredUUID
}
