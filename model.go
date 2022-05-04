package main

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"Name"`
	Email  string `json:"Email"`
	Phone  string `json:"Phone"`
	CityId string `json:"City_id"`
}

//id INT PRIMARY KEY AUTO_INCREMENT,
//Name VARCHAR(255),
//Email VARCHAR(255),
//Phone VARCHAR(15),
//City_id INT
