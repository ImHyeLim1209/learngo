package banking

//Account struct
type Account struct {
	Owner   string
	Balance int
}

// //Account struct
// type Account struct { //export해서 쓰려면 대문자로 시작하고, 위에 Accout로 시작하는 주석이 있어야 한다.
// 	Owner   string //export 하려면 속성도 대문자로 시작해야 한다. (대문자는 public, 소문자를 private인 것!)
// 	Balance int
// }
