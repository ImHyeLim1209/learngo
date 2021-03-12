package basicfun

import (
	"fmt"
	"strings"
)

//export하고 싶은 함수 등이 있다면 대문자로 시작하게 하면된다!
func Multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) { //2개 이상의 value 리턴하기
	return len(name), strings.ToUpper(name)
}

func nakedLenAndUpper(name string) (length int, uppercase string) { //함수 선언에서 리턴되는 변수 미리 만들어주고 리턴하기
	defer fmt.Println("function 끝나면 실행된다!") //callback과 비슷한 느낌 함수가 끝나면 바로 실행 defer키워드
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) { //여러 개의 인자 받기
	fmt.Println(words)
}

func superAdd(numbers ...int) { //for문 사용하기1
	for index, number := range numbers { //항상 index, number 순이다
		fmt.Println(index, number)
	}

	for i := 0; i < len(numbers); i++ { //for문 사용하기2
		fmt.Println(numbers[i])
	}
}

func canIDrink(age int) bool { //if문
	if koreanAge := age + 2; koreanAge < 18 { //if를 하면서 koreanAge라는 변수를 만들어서 사용할 수 있다!(안 써도 상관은 없음)
		return false
	}
	return true
}

func canIDrinkSwitch(age int) bool {
	switch koreanAge := age + 2; koreanAge { //여기 변수 안 넣고도 switch 할 수 있다. and 여기서도 변수를 새롭게 선언하여 사용할 수 있다.
	case 10: //switch에 변수를 안 넣었을 경우엔 age > 10 이런 식으로 case의 조건을 넣을 수도 있다.
		return false
	case 18:
		return true
	case 50:
		return false
	}
	return false
}

func useFun() {
	//fmt.Println(multiply(2, 2))

	//totalLength, upperName := lenAndUpper("ihl") //2개의 value를 리턴하는 함수는 2개의 value 모두 받아야 한다. 받기 싫으면 해당 위치를 _로 무시할 수 있다.
	//fmt.Println(totalLength, upperName)

	//totalLength, upperName := nakedLenAndUpper("ihl")
	//fmt.Println(totalLength, upperName)

	//repeatMe("im", "hye", "lim", "!!!")

	//superAdd(1, 2, 3, 4, 5)

	//fmt.Println(canIDrink(16))

	fmt.Println(canIDrinkSwitch(16))
}
