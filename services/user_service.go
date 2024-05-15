package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/YahyaCengiz/go_deneme/CRUD/models"
)

func LoadUsers(filename string) ([]models.User, error) {
    var users []models.User

    absPath, err := filepath.Abs(filename)
    if err != nil {
        return nil, err
    }

    data, err := ioutil.ReadFile(absPath)
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(data, &users)
    if err != nil {
        return nil, err
    }

    return users, nil
}

func AuthenticateUser(username, password string) (*models.User, error) {
    users, err := LoadUsers("./users.json")
    if err != nil {
        return nil, err
    }

    for _, user := range users {
        if user.Username == username && user.Password == password {
            return &user, nil
        }
    }

    return nil, errors.New("invalid username or password")
}
