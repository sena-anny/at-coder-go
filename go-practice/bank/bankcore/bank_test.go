package bank

import "testing"

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Jhon",
			Address: "LA",
			Phone:   "xxx",
		},
		Number:  1001,
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("can't create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Jhon",
			Address: "LA",
			Phone:   "xxx",
		},
		Number:  1001,
		Balance: 0,
	}
	account.Deposit(10)
	if account.Balance != 10 {
		t.Error("balance is not being updated after a depoist")
	}
}
func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Jhon",
			Address: "LA",
			Phone:   "xxx",
		},
		Number:  1001,
		Balance: 0,
	}
	if err := account.Deposit(-20); err == nil {
		t.Error("only positive numbers should be allowed to deposit\n", err)
	}
}
func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("balance is not being updated after withdraw")
	}
}
func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()
	if statement != "1001 - John - 100" {
		t.Error("statement doesn't have the proper format")
	}
}

func TestSendMoney(t *testing.T) {
	requester := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 1000,
	}
	receiver := &Account{
		Customer: Customer{
			Name:    "Bob",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1002,
		Balance: 0,
	}

	err := requester.Transfer(100, receiver)
	if err != nil {
		t.Error("Can't send money", err)
	}
	statement := requester.Statement()
	if statement != "1001 - John - 900" {
		t.Error("statement doesn't have the proper format", statement)
	}
	statement = receiver.Statement()
	if statement != "1002 - Bob - 100" {
		t.Error("statement doesn't have the proper format", statement)
	}
}
