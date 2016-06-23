package tools

import(
	"ttapi/models"
)

var (
	CHBool chan bool
	CHUser chan []models.User
	CHRelationship chan []models.Relationship
)
type ChannelHelper struct {}

func init() {
	CHBool = make(chan bool, 20)
	CHUser = make(chan []models.User, 10)
	CHRelationship = make(chan []models.Relationship, 10)
}
