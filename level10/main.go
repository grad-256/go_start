package main

import "fmt"

//構造体 classみたいな感じ

type T struct {
	User User
}

type User struct {
	Name string
	Age  int
	// X, Y int
}

// type Users []*User

// func (u User) SayName() {
// 	fmt.Println(u.Name)
// }

// func (u User) SetName(name string) {
// 	u.Name = name
// }

// // レシーバーはポインタ型にする必要がある
// func (u *User) SetName2(name string) {
// 	u.Name = name
// }

// func (u *User) SetName3() {
// 	u.Name = "A"
// }

// // 値渡し
// func UpdateUser(user User) {
// 	user.Name = "A"
// 	user.Age = 1000
// }

// // 参照渡し
// func UpdateUser2(user *User) {
// 	user.Name = "A"
// 	user.Age = 1000
// }

// 型コンストラクタ　コンストラクタ関数
// func NewUser(name string, age int) *User {
// 	return &User{Name: name, Age: age}
// }

func main() {
	// var user1 User
	// fmt.Println(user1)
	// user1.Name = "user1"
	// user1.Age = 10
	// fmt.Println(user1)

	// user2 := User{}
	// fmt.Println(user2)
	// user2.Name = "user2"
	// fmt.Println(user2)

	// user3 := User{Name: "user3", Age: 30}
	// fmt.Println(user3)

	// user4 := User{"user4", 40}
	// fmt.Println(user4)

	// user6 := User{Name: "user6"}
	// fmt.Println(user6)

	// // ポインタ型を返している &{ 0}
	// user7 := new(User)
	// fmt.Println(user7)

	// // 関数の引数として構造体の引数を渡す場合
	// user8 := &User{}
	// fmt.Println(user8)

	// UpdateUser(user6)
	// UpdateUser2(user8)
	// fmt.Println(user6)
	// fmt.Println(user8)

	// user1 := User{Name: "user1"}
	// user1.SayName()

	// user1.SetName("A")
	// user1.SayName()

	// user1.SetName2("A")
	// user1.SayName()

	// t := T{User: User{Name: "user1", Age: 10}}

	// fmt.Println(t)
	// fmt.Println(t.User)
	// fmt.Println(t.User.Name)

	// t.User.SetName3()
	// fmt.Println(t)

	// user1 := NewUser("user1", 10)
	// fmt.Println(user1)

	// デリファレンスする
	// リファレンスの指し示す先の値を取得・参照すること
	// リファレンス
	// 何かの値を参照しているもののこと
	// fmt.Println(*user1)

	// user1 := User{Name: "user1", Age: 10}
	// user2 := User{Name: "user2", Age: 20}
	// user3 := User{Name: "user3", Age: 30}
	// user4 := User{Name: "user4", Age: 40}

	// users := Users{}

	// users = append(users, &user1)
	// users = append(users, &user2)
	// users = append(users, &user3, &user4)

	// // fmt.Println(users)
	// // for _, u := range users {
	// // 	fmt.Println(u)
	// // }

	// users2 := make([]*User, 0)
	// users2 = append(users2, &user1, &user2)

	// for _, u := range users2 {
	// 	fmt.Println(*u)
	// 	fmt.Println(u.Name)
	// 	fmt.Println(u.Age)
	// }

	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
	}
	// for _, u := range m {
	// 	fmt.Println(u)
	// }
	fmt.Println(m)

	m2 := map[User]string{
		{Name: "user1", Age: 10}: "Tokyo",
		{Name: "user2", Age: 20}: "LA",
	}

	fmt.Println(m2)

}
