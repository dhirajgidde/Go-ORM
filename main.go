package main

import (
	"fmt"

	"github.com/astaxie/beego/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      int //default field
	Name    string
	Subject string
	Marks   int
}

func init() {
	orm.RegisterModel(new(User))                                                         //Register A ORM
	orm.RegisterDriver("mysql", orm.DRMySQL)                                             //Register A Driver
	orm.RegisterDataBase("default", "mysql", "dhiraj:dhiraj@tcp(127.0.0.1:3306)/dhiraj") //Register A DataBase
}

func fetchDetails() {
	o := orm.NewOrm()
	user := User{Id: 1}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("No Result Found")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key is found")
	} else {
		fmt.Println(user)
	}
}

func insertQuery() {
	o := orm.NewOrm()
	var user User
	user.Id = 2
	user.Name = "Tata"
	user.Subject = "nano"
	user.Marks = 450

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
		fmt.Println("Record Inserted")
		
	}

}

func updateQuery() {
	o := orm.NewOrm()
	user := User{Id: 1}
	if o.Read(&user) == nil {
		user.Subject = "English"
		if num, err := o.Update(&user); err == nil {
			fmt.Println(num)
			fmt.Println("Record Updated")
		}
	}

}

func deleteQuery() {
	o := orm.NewOrm()
	if num, err := o.Delete(&User{Id: 1}); err == nil {
		fmt.Println(num)
		fmt.Println("Record Deleted")
	}
}

func main() {
	//fetchDetails()
	//insertQuery()
	//deleteQuery()
	updateQuery()
}
