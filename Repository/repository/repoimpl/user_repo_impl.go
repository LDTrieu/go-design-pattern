package repoimpl

import(
	"fmt"
	"database/sql"
	repo "RepositoryPattern/repository"
	models "RepositoryPattern/model"
)
type UserRepoImpl struct{
	Db * sql.DB
}

func NewUserRepo(db *sql.DB) repo.UserRepo{
	return &UserRepoImpl{
		Db : db,
	}
}

func (u *UserRepoImpl) Select()([]models.User, error){
	users := make([]models.User,0)
	rows, err := u.Db.Query("SELECT * FROM users")
	if err != nil{
		return users, err
	}
	for rows.Next(){
		user := models.User{}
		// SELECT id....
		err := rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Email)
		if err != nil {
			break
		}

		users=append(users, user)
	}
	err = rows.Err()
	if err != nil{
		return users, err
	}
	return users, nil
}

func (u *UserRepoImpl) Insert(user models.User) (error){
	insertStatement := `
	INSERT INTO users  (id, name, gender, email)
	VALUES ($1, $2, $3, $4)`

	_, err := u.Db.Exec(insertStatement, user.ID, user.Name, user.Gender, user.Email)
	if err != nil {
		return err
	}
	fmt.Println("record added: ", user)
	return nil
}