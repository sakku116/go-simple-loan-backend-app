package enum

import "strconv"

type LoanTermMonths int

const (
	LoanTermMonths_1 = LoanTermMonths(1)
	LoanTermMonths_2 = LoanTermMonths(2)
	LoanTermMonths_3 = LoanTermMonths(3)
	LoanTermMonths_6 = LoanTermMonths(6)
)

var validLoanTermMonths = []LoanTermMonths{LoanTermMonths_1, LoanTermMonths_2, LoanTermMonths_3, LoanTermMonths_6}

func (loanTermMonths LoanTermMonths) String() string {
	return strconv.Itoa(int(loanTermMonths))
}

func (loanTermMonths LoanTermMonths) IsValid() bool {
	for _, item := range []LoanTermMonths{
		LoanTermMonths_1, LoanTermMonths_2,
		LoanTermMonths_3, LoanTermMonths_6,
	} {
		if item == loanTermMonths {
			return true
		}
	}
	return false
}

type LoanStatus string

const (
	LoanStatus_PENDING  = LoanStatus("PENDING")
	LoanStatus_APPROVED = LoanStatus("APPROVED")
	LoanStatus_REJECTED = LoanStatus("REJECTED")
	LoanStatus_PAID     = LoanStatus("PAID")
)

var ValidLoanStatus = []LoanStatus{
	LoanStatus_PENDING, LoanStatus_APPROVED,
	LoanStatus_REJECTED, LoanStatus_PAID,
}

func (e LoanStatus) String() string {
	return string(e)
}

func (e LoanStatus) IsValid() bool {
	for _, item := range ValidLoanStatus {
		if item == e {
			return true
		}
	}

	return false
}
