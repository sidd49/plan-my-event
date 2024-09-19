package models

type User struct {
	ID       string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// func (u *User) Save() error {
// 	query := `
// 	INSERT INTO users(email, password) VALUES (?, ?)
// 	`
// 	stmt, err := db.DB.Prepare(query)

// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	hashPassword, err := utils.HashPassword(u.Password)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := stmt.Exec(u.Email, hashPassword)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = result.LastInsertId()

// 	return err
// }

// func (u *User) ValidateCredentials() error {
// 	query := "SELECT id, password FROM users WHERE email = ?"
// 	row := db.DB.QueryRow(query, u.Email)

// 	var retrievedPassword string
// 	var retrievedID int64
// 	err := row.Scan(&retrievedID, &retrievedPassword)

// 	if err != nil {
// 		return errors.New("credentials invalid")
// 	}
// 	if !utils.CheckPasswordHash(retrievedPassword, u.Password) {
// 		return errors.New("credentials invalid")
// 	}
// 	u.ID = retrievedID
// 	return nil
// }
