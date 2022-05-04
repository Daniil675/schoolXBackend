package main

import (
	"fmt"
)

func UsersGetAllById(id int) []User {
	var entities []User

	//rows, err := db.Query("SELECT * FROM Users")
	stmt, err := db.Prepare("SELECT * FROM Users WHERE City_id=?")
	if err != nil {
		fmt.Println(err)
	}
	
	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var (
			entity User
		)
		err = rows.Scan(&entity.ID, &entity.Name, &entity.Email, &entity.Phone, &entity.CityId)
		if err != nil {
			fmt.Println(err)
		}

		entities = append(entities, entity)
	}

	rows.Close() //good habit to close
	return entities
}
