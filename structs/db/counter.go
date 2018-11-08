package db

type (
	// Counter - Counter DB
	Counter struct {
		BranchCode string `bson:"branch_code"`
		Month      int    `bson:"month"`
		Counter    int    `bson:"counter"`
	}
)
