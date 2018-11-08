package genappid

import (
	"collections/helper/constant"
	iGenappid "collections/models/db/interfaces/genappid"
	mgGenappid "collections/models/db/mongo/genappid"
	stGenappid "collections/models/stub/genappid"
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
