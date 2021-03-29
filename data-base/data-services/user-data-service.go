package data_base

import "context"

func (dataBase *DataBase) CreateUser(context context.Context) {
	dataBase.UserData.Find(context)
}
