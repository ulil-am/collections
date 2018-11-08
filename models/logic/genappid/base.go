package genappid

import (
	"medium/helper/constant"
	iGenappid "medium/models/db/interfaces/genappid"
	mgGenappid "medium/models/db/mongo/genappid"
	stGenappid "medium/models/stub/genappid"
)

var (
	logicName = "@genappid"

	// DBGenappid DBname
	DBGenappid iGenappid.IGenappid
)

func init() {
	if constant.GOENV == constant.DEVCI {
		DBGenappid = new(stGenappid.Genappid)
	} else {
		DBGenappid = new(mgGenappid.Genappid)
	}
}
