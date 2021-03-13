package accounts

import (
	"errors"
	"fmt"
)

//Account struct
type account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw") //err의 이름은 err~이어야 한다(컨벤션)

//NewAccount creates Account
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

//Deposit x amount on your account
func (a *account) Deposit(amount int) { //(a account) 부분은 Receiver로, a 타입이 account이므로 account의 메소드이다. a(struct의 첫글자 소문자)라고 짓는 것은 컨벤션이다.
	a.balance += amount //*account로 하면 copy하지 않고 원본을 가져와서 조작한다.
}

//Withdraw x amount from your account
func (a *account) WithDraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil //error가 없다는 의미
}

//ChangeOwner of the account
func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Owner of account
func (a account) Owner() string {
	return a.owner
}

//Balance of your account
func (a account) Balance() int { //메소드는 기본적으로 원본을 바꾸지 않고 a라는 copy를 이용하여 조작한다.
	return a.balance //이 함수는 원본을 변경하지 않으므로 상관없다.
}

func (a account) String() string { //toString()같은 함수 재선언
	return fmt.Sprint(a.Owner(), "'s account\nHas: ", a.Balance())
}
