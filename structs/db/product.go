package db

// Product - ProductStruct
type Product struct {
	CampaignCode                 string `bson:"campaign_code"`
	CampaignDesc                 string `bson:"campaign_desc"`
	MaxLoanAmount                int    `bson:"max_loan_amount"`
	MaxTerm                      int    `bson:"max_term"`
	InterestRate                 int    `bson:"interest_rate"`
	LoanCategoryCode             string `bson:"loan_catagory_code"`
	PaymentAmountRounding        int    `bson:"payment_amount_rounding"`
	InstallmentCalculationMethod string `bson:"installment_calculation_method"`
}
