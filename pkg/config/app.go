package config

import(
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dilects/mysgl"
)
var (
    db * gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:morter/fammember?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
    db =d
}
func GetDB() *gorm.DB {
    return db
}