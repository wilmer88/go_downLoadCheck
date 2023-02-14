package models

import(
	"github.com/jinzhu/gorm"
	"github.com/wilmer88/go_downLoadCheck/pkg/config"
)
var db *gorm.DB

type Fammember struct{
gorm.model
FirstName string `gorm:""json: FirstName`
Happiness int `gorm:""json: Happiness`
UrlStr string `gorm:""json: Urlstr`
} 

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Fammember{})
}

func(b *Fammember) CreateMember() *Fammember{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetALLMembers() []Fammember {
	var Members []Fammember
	db.Find(&Members)
	return Members
}

func GetMemberByID(id int64) (*Fammember, *gorm.DB) {
	var getMember Fammember
	db:=db.Where("FamID=?",id).Find(&getMember)
	return &getMember, db
}

func DeleteMember(ID int64) Fammember{
	var fammember Fammember
	db.Where("famID", ID).Delete(fammember)
	return fammember
}

