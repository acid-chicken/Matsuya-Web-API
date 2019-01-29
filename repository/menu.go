package repository

import (
	"math/rand"
	"time"

	"github.com/makotia/Matsuya-Web-API/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MenuRepository メニューリポジトリ
type MenuRepository interface {
	FindRandom() (*model.MatsuyaMenu, error)
}

// MenuRepositoryImpl メニューリポジトリインターフェイスの実装
type MenuRepositoryImpl struct {
	db *mgo.Database
}

// NewMenuRepository メニューリポジトリの実装を生成
func NewMenuRepository(db *mgo.Database) MenuRepository {
	return &MenuRepositoryImpl{
		db: db,
	}
}

// FindRandom ランダムでメニューを返す
func (r *MenuRepositoryImpl) FindRandom() (*model.MatsuyaMenu, error) {
	var menu []*model.MatsuyaMenu
	err := r.db.C("menus").Find(bson.M{}).All(&menu)
	if err != nil {
		return nil, err
	}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(menu))
	return menu[i], nil
}
