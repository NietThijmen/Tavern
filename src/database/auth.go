package database

func GetUserById(id int) (User, error) {
	row := Connection.QueryRow("SELECT * FROM users WHERE id = ? LIMIT 1", id)
	var user User
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	row := Connection.QueryRow("SELECT * FROM users WHERE email = ? LIMIT 1", email)
	var user User
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func CreateUser(email string, password string) (User, error) {
	result, err := Connection.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, password)
	if err != nil {
		return User{}, err
	}

	id, _ := result.LastInsertId()
	user, _ := GetUserById(int(id))

	return user, nil
}

func DeleteUser(id int) error {
	_, err := Connection.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func Login(email string, password string) {

}

func CreateApiKey(userId int) {

}
