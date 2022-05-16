package datastore

import (
	"fmt"
)

const cityTable = "City"

func CityAdd(e City) int {
	//rows, err := db.Query("SELECT * FROM Users")
	stmt, err := db.Prepare(fmt.Sprintf("INSERT INTO %s(Name) VALUES (?)", cityTable))
	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(e.Name)
	if err != nil {
		fmt.Println(err)
	}

	id, _ := res.LastInsertId()
	return int(id)
}

func CityGet(id int) City {
	stmt, err := db.Prepare(fmt.Sprintf("SELECT * FROM %s WHERE id=?", cityTable)) //db.Prepare(`SELECT id,serialized_item FROM categories WHERE id=?`)
	if err != nil {
		fmt.Println(err)
	}

	var entity City
	err = stmt.QueryRow(id).Scan(&entity.ID, &entity.Name)
	if err != nil {
		fmt.Println(err)
	}

	return entity
}

func CitiesGetAll() []City {
	var entities []City

	//rows, err := db.Query("SELECT * FROM Users")
	stmt, err := db.Prepare(fmt.Sprintf("SELECT * FROM %s", cityTable))
	if err != nil {
		fmt.Println(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var (
			entity City
		)
		err = rows.Scan(&entity.ID, &entity.Name)
		if err != nil {
			fmt.Println(err)
		}

		entities = append(entities, entity)
	}

	rows.Close() //good habit to close
	return entities
}

func CitiesGetWithOffsetAndLimit(offset, limit int) []City {
	var entities []City

	//rows, err := db.Query("SELECT * FROM Users")
	stmt, err := db.Prepare(fmt.Sprintf("SELECT * FROM %s ORDER BY id LIMIT ? OFFSET ?;", cityTable))
	if err != nil {
		fmt.Println(err)
	}

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var (
			entity City
		)
		err = rows.Scan(&entity.ID, &entity.Name)
		if err != nil {
			fmt.Println(err)
		}

		entities = append(entities, entity)
	}

	rows.Close() //good habit to close
	return entities
}

func CityEdit(e City) {
	stmt, err := db.Prepare(fmt.Sprintf("UPDATE %s SET Name=? WHERE id=?", cityTable)) //db.Prepare("UPDATE categories set serialized_item = ? where id = ?")
	if err != nil {
		fmt.Println(err)
	}

	_, err = stmt.Exec(e.Name, e.ID)
	if err != nil {
		fmt.Println(err)
	}
}

func CityDelete(id int) {
	stmt, err := db.Prepare(fmt.Sprintf("DELETE FROM %s WHERE id=?", cityTable)) //db.Prepare("delete from categories where id=?")
	if err != nil {
		fmt.Println(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
}
