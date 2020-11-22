package database

import "database/sql"

type Customer struct{
	customerId int 'json:"customer_id"`
	FristName String `json:"first_name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`


































//Get Customer Data
func Get Customers(db*sql.DB){
	rows, err := db.Query( query:"select * from Customers")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer.rows.Close()

	var result []Customer

	//Iterate data per baris yang didapat dari Query select
	for.rows.Next(){
		var each = Customer{}
		//scan per kolom field
		