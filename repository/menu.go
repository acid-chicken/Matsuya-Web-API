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
	GetAllMenu() ([]*model.MatsuyaMenu, error)
	FindRandom() (*model.MatsuyaMenu, error)
	FindByType(typeName string) ([]*model.MatsuyaMenu, error)
	FindByName(name string) (*model.MatsuyaMenu, error)
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

// GetAllMenu すべてのメニューを返す
func (r *MenuRepositoryImpl) GetAllMenu() ([]*model.MatsuyaMenu, error) {
	var menu []*model.MatsuyaMenu
	err := r.db.C("menus").Find(bson.M{}).All(&menu)
	return menu, err
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

// FindByType ランダムでメニューを返す(種類でフィルタリング)
func (r *MenuRepositoryImpl) FindByType(typeName string) ([]*model.MatsuyaMenu, error) {
	var menu []*model.MatsuyaMenu
	err := r.db.C("menus").Find(bson.M{
		"type": typeName,
	}).All(&menu)
	return menu, err
}

// FindByName メニューを返す(名前でフィルタリング)
func (r *MenuRepositoryImpl) FindByName(name string) (*model.MatsuyaMenu, error) {
	var menu *model.MatsuyaMenu
	err := r.db.C("menus").Find(bson.M{
		"name": name,
	}).One(&menu)
	return menu, err
}
