package enum

type LoanTermMonths int

const (
	LoanTermMonths_1 = LoanTermMonths(1)
	LoanTermMonths_2 = LoanTermMonths(2)
	LoanTermMonths_3 = LoanTermMonths(3)
	LoanTermMonths_6 = LoanTermMonths(6)
)

var validLoanTermMonths = []LoanTermMonths{LoanTermMonths_1, LoanTermMonths_2, LoanTermMonths_3, LoanTermMonths_6}

func (loanTermMonths LoanTermMonths) Int() int {
	return int(loanTermMonths)
}

func (loanTermMonths LoanTermMonths) IsValid() bool {
	for _, item := range validLoanTermMonths {
		if item == loanTermMonths {
			return true
		}
	}
	return false
}
