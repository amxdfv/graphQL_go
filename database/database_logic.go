package database

import (
	"database/sql"
	"errors"
	_ "github.com/iamacarpet/go-sqlite3-dynamic"
	"go-graphql-api/graph/model"
)

func GetUserFromDB(id string) (*model.User, error) {
	db, err := returnDB() // TODO надо привести этот колхоз в нормальный вид
	var sample_user model.User
	if err != nil {
		return nil, err
	}
	getUserRequest := "select * from users where id =?"
	res, err := db.Query(getUserRequest, id)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	if res.Next() == false {
		err = errors.New("Пользователь не найден")
		return nil, err
	}
	err = res.Scan(&sample_user.ID, &sample_user.Name, &sample_user.Age, &sample_user.Address, &sample_user.DocumentType, &sample_user.Document)
	if err != nil {
		return nil, err
	}
	db.Close()
	var user = &sample_user
	return user, nil
}

func GetTransFromDB(id string) (*model.Transaction, error) {
	db, err := returnDB()
	var sample_trans model.Transaction
	if err != nil {
		return nil, err
	}
	getTransRequest := "select * from transactions where id =?"
	res, err := db.Query(getTransRequest, id)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	if res.Next() == false {
		err = errors.New("Транзакция не найдена")
		return nil, err
	}
	err = res.Scan(&sample_trans.ID, &sample_trans.Rrn, &sample_trans.Amount, &sample_trans.Currency, &sample_trans.UserID, &sample_trans.GoodID, &sample_trans.Place, &sample_trans.Time)
	if err != nil {
		return nil, err
	}
	db.Close()
	var trans = &sample_trans
	return trans, nil
}

func GetGoodFromDB(id string) (*model.Good, error) {
	db, err := returnDB()
	var sample_good model.Good
	if err != nil {
		return nil, err
	}
	getGoodRequest := "select * from goods where id =?"
	res, err := db.Query(getGoodRequest, id)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	if res.Next() == false {
		err = errors.New("Товар не найден")
		return nil, err
	}
	err = res.Scan(&sample_good.ID, &sample_good.Name, &sample_good.Price, &sample_good.Currency, &sample_good.CountryOrigin)
	if err != nil {
		return nil, err
	}
	db.Close()
	var good = &sample_good
	return good, nil

}

func returnDB() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "database\\corebase.db")
	return database, err
}
