func arrayTest() {
	names := [5]string{"i am", "ihl"} //[배열길이]타입{values}
	names[2] = "!!"

	unlimitedNames := []string{"i am", "infinity", "???"} //길이제한이 없는 배열
	unlimitedNames[2] = "!!!"                             //배열길이를 넘는 곳에는 접근할 수 없다.
	unlimitedNames = append(unlimitedNames, "append!")    //원본을 바꾸지 않으므로 value를 다시 넣어준다.

	fmt.Println(names)
	fmt.Println(unlimitedNames)
}

func mapFun() {
	//map: key와 value의 타입이 정해져있는 Dictionary..?
	ihl := map[string]string{"name": "ihl", "age": "12"} //map[key타입]value타입{"key": value...}
	fmt.Println(ihl)
}

func structFun() {
	type person struct {
		name    string
		age     int
		favFood []string
	}

	ihl1 := person{"ihl", 12, []string{"ramen", "kimchi"}}                     //순서대로 value 적기
	ihl2 := person{name: "ihl", age: 14, favFood: []string{"macaron", "cake"}} //key를 명시하고 value 적기

	fmt.Println(ihl1)
	fmt.Println(ihl2)
}