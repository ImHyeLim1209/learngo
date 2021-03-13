package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string //map[string]string 타입인 새로운 타입 Dictionary를 만듦 cf. type Money int -> Money(1) 처럼 사용

var (
	errNotFound   = errors.New("Not Found") //같은 타입의 변수를 여러번 선언할 때 다음과 같이 선언할 수 있다.
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Can't update non-existing word")
	errCantDelete = errors.New("Can't delete non-existing word")
)

//Search for a word
func (d Dictionary) Search(word string) (string, error) { //type도 메소드를 추가할 수 있다.
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

//Add a word to dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

//Update a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

//Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
