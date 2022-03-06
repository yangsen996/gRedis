package lib

import (
	"encoding/json"
	"github.com/yangsen996/gRedis/greids"
)

func DbGetNews(id string) greids.DBGetter {
	return func() string {
		newModle := NewModel{}
		GORMDB.Table("news").Where("id=?", id).Find(&newModle)
		b, _ := json.Marshal(newModle)
		return string(b)
	}
}
