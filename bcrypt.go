package main
//               BCRYPT 
// is an excellent way to store password information.
// - as soon as you receive a password from a user,
// you use bcrypt to encrypt it, you store the 
// encrypted value.  You never even know that password.

import "fmt"
// import "bcrypt" 
import "golang.org/x/crypto/bcrypt"



func main() {

	s := `password123`
//	bcrypt.GenerateFromPassword(password []byte, cost int) ([]byte, error)
	bs, err := bcrypt.GenerateFromPassword( []byte(s), bcrypt.MinCost) 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(s)
}