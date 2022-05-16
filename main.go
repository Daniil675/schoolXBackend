package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"schoolXBackend/datastore"
	"schoolXBackend/server"
	"sync"
)

func main() {
	//getUserByCityId(1)
	var wg sync.WaitGroup

	wg.Add(1)
	s := server.Server{}
	go s.Start("8080")

	id := datastore.CityAdd(datastore.City{
		Name: "San Francisco",
	})

	fmt.Println(datastore.CityGet(id))
	datastore.CityEdit(datastore.City{
		ID:   id,
		Name: "Palo Alto",
	})
	fmt.Println(datastore.CityGet(id))

	fmt.Println(datastore.CitiesGetAll())
	datastore.CityDelete(id)
	fmt.Println(datastore.CitiesGetAll())

	fmt.Println(datastore.CitiesGetWithOffsetAndLimit(0, 2))
	fmt.Println(datastore.CitiesGetWithOffsetAndLimit(1, 3))

	wg.Wait()
}

func getUserByCityId(id int) { //TODO: Deprecation
	users := datastore.UsersGetAllByCityId(id)

	if len(users) == 0 {
		fmt.Println("Пользователь с таким городом не найден")
		return
	}
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}
