package datastore

import "fmt"

const usersTable = "Users"

func UsersGetAllByCityId(id int) []User {
	var entities []User

	//rows, err := db.Query("SELECT * FROM Users")
	stmt, err := db.Prepare(fmt.Sprintf("SELECT * FROM %s WHERE City_id=?", usersTable))
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
