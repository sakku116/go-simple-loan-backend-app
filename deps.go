package main

import ucase "backend/usecase"

type CommonDeps struct {
	AuthUcase ucase.IAuthUcase
	UserUcase ucase.IUserUcase
	LoanUcase ucase.ILoanUcase
}
