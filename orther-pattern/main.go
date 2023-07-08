package main

import(
	"fmt"
	"RepositoryPattern/driver"
	"RepositoryPattern/repository/repoimpl"
	models "RepositoryPattern/model"
)
// host,port,user,password, dbname
const (
	host = "localhost"
	port="5432"
	user ="ryan"
	password="123456"
	dbname="demoRepo"
)

func main(){
	db :=driver.Connect(host,port,user,password, dbname)
	
	err := db.SQL.Ping()
	if err != nil{
		panic(err)
	}
	userRepo := repoimpl.NewUserRepo(db.SQL)

	nva := models.User{
		ID: 1,
		Name: "Nguyen van A",
		Gender: "Male",
		Email: "nguyenvana@gmail.com",
	}

	pvc := models.User{
		ID: 2,
		Name: "Phan Van Canh",
		Gender: "Female",
		Email: "canh2k6@gmail.com",
	}
	userRepo.Insert(nva)
	userRepo.Insert(pvc)

	users,_ := userRepo.Select()

	for i:= range users{
		fmt.Println(users[i])
	}
}