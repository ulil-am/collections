package cards

import (
	// "collections/helper/constant"
	iUser "collections/models/db/interfaces/user"
	mgUser "collections/models/db/mongo/user"

	"collections/helper/constant"
	iCounter "collections/models/db/interfaces/counter"
	mgCounter "collections/models/db/mongo/counter"
)

var (
	logicName = "@User"
	// DBUser ...
	DBUser iUser.IUser
	// DBCounter ...
	DBCounter iCounter.ICounter
)

func init() {
	if constant.GOENV == constant.DEVCI {
		// DBUser = new(stUser.User)
		// DBCounter = new(stCounter.Counter)

	} else {
		DBUser = new(mgUser.User)
		DBCounter = new(mgCounter.Counter)
	}
}
