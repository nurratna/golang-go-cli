package main

import(
	"fmt"
	"os"
	"encoding/json"
    "io/ioutil"
)

func UserFromFile(f string) User {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return User{}
	}
	return userFromJson(bs)
}

func (u User) Save(f string) error {
	return ioutil.WriteFile(f, u.toJson(), 0644)
}

func (u User) toJson() []byte {
	bs, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bs
}

func userFromJson(bs []byte) User {
	var u User
	err := json.Unmarshal(bs, &u)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return u
}