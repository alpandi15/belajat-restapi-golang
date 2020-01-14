package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type student struct {
	ID    string `form:"id" json:"id"`
	Name  string `form:"name" json:"name"`
	Age   int    `form:"age" json:"age"`
	Grade int    `form:"grade" json:"grade"`
}

type meta struct {
	TotalData int `json:"totalData"`
}

type response struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Meta    meta      `json:"meta"`
	Data    []student `json:"data"`
}

func getAll(w http.ResponseWriter, r *http.Request) {
	var users student
	var arrUser []student
	var response response
	var meta meta

	db, err := connect()
	defer db.Close()

	rows, err := db.Query("select * from tb_student")
	if err != nil {
		log.Print(err)
	}
	columns, err := rows.Columns()
	if err != nil {
		log.Print(err)
	}
	count := len(columns)
	meta.TotalData = count

	for rows.Next() {
		if err := rows.Scan(&users.ID, &users.Name, &users.Age, &users.Grade); err != nil {
			log.Fatal(err.Error())

		} else {
			arrUser = append(arrUser, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrUser
	response.Meta = meta
	fmt.Println(arrUser)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
