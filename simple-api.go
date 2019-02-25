package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type DataStudent struct {
	ID      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Address string `json:"address,omitempty"`
}

func getMahasiswa(w http.ResponseWriter, r *http.Request) {
	//-- Set Database Connection --\\
	db, err := sql.Open("mysql", "user:password@/db")
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	//-- Query set select data
	rows, err := db.Query("SELECT * FROM mahasiswa")
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	//-- Retrieve data
	for rows.Next() {
		var id int64
		var name string
		var phone string
		var address string
		err = rows.Scan(&id, &name, &phone, &address)
		if err != nil {
			fmt.Printf("%s", err.Error())
		}

		dataStudent := DataStudent{id, name, phone, address}

		js, err := json.Marshal(dataStudent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Tpe", "application/json")
		w.Write(js)
	}
}

func insertMahasiswa(w http.ResponseWriter, r *http.Request) {
	//-- Set Database Connection --\\
	db, err := sql.Open("mysql", "user:password@/db")
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		}

		data := DataStudent{}
		err = json.Unmarshal(body, &data)

		Name := data.Name
		Phone := data.Phone
		Address := data.Address

		//-- Query set insert data
		stmt, err := db.Prepare("INSERT INTO mahasiswa (name,phone,address) VALUES (?,?,?)")
		if err != nil {
			fmt.Printf("%s", err.Error())
		}

		//-- Define the dataset
		res, err := stmt.Exec(Name, Phone, Address)
		if err != nil {
			fmt.Printf("%s", err.Error())
		}

		//-- Insert data next id
		id, err := res.LastInsertId()
		if err != nil {
			fmt.Printf("%s", err.Error())
		} else {
			fmt.Println(id)
			w.Header().Set("Content-Tpe", "application/json")
			w.Write([]byte("Insert data is successfully"))
		}

	} else {
		fmt.Println("this endpoint just POST method")
	}

}

func updateMahasiswa(w http.ResponseWriter, r *http.Request) {
	//-- Set Database Connection --\\
	db, err := sql.Open("mysql", "user:password@/db")
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		}

		data := DataStudent{}
		err = json.Unmarshal(body, &data)

		ID := data.ID
		Name := data.Name
		Phone := data.Phone
		Address := data.Address

		//-- Query set update data
		stmt, err := db.Prepare("UPDATE mahasiswa SET name=?, phone=?, address=? WHERE id=?")
		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		}

		//-- Define the dataset
		res, err := stmt.Exec(Name, Phone, Address, ID)
		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		}

		//-- Return the execute update data
		affect, err := res.RowsAffected()
		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		} else {
			fmt.Println(affect)
			w.Header().Set("Content-Tpe", "application/json")
			w.Write([]byte("Update data is successfully"))
		}

	}
}

// func deleteMahasiswa(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)

// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("value", strings.Join(v, ""))
// 	}

// }

func main() {
	http.HandleFunc("/getmhs", getMahasiswa)
	http.HandleFunc("/insertmhs", insertMahasiswa)
	http.HandleFunc("/updatemhs", updateMahasiswa)
	// http.HandleFunc("/deletemhs", deleteMahasiswa)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
