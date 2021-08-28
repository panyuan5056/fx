package xstrategy

import (
	"fx/models"
)

type Strategy struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Class1   string `form:"class1" json:"class1" binding:"required"`
	Class2   string `form:"class2" json:"class2" binding:"required"`
	Class3   string `form:"class3" json:"class3" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Dept     string `form:"dept" json:"dept" binding:"required"`
	Category string `form:"category" json:"category" binding:"required"`
	Plan     string `form:"plan" json:"plan" binding:"required"`
	Desc     string `form:"desc" json:"desc" binding:"required"`
	Forward  string `form:"forward" json:"forward" binding:"required"`
}

func (st *Strategy) valid() (int, string) {
	if st.Class1 == "" {
		return 400, "class1不得为空"
	} else if st.Class2 == "" {
		return 400, "class2不得为空"
	} else if st.Class3 == "" {
		return 400, "class3不得为空"
	} else if st.Dept == "" {
		return 400, "dept不得为空"
	} else if st.Name == "" {
		return 400, "name不得为空"
	}
	return 200, "成功"
}

func (st *Strategy) dict() map[string]string {
	l := make(map[string]string)
	l["class1"] = st.Class1
	l["class2"] = st.Class2
	l["class3"] = st.Class3
	l["dept"] = st.Dept
	l["name"] = st.Name
	return l
}

func (st *Strategy) save() {
	if sid := st.get(); sid > -1 {
		var strategy Strategy
		models.DB.Model(&strategy).Select("id", sid).Updates(st)
	} else {
		models.DB.Create(&st)
	}
}

func (st *Strategy) get() int {
	var strategy Strategy
	models.DB.Where("class3 = ? AND name = ?", st.Class3, st.Name).First(&strategy)
	if strategy.ID > 0 {
		return strategy.ID
	}
	return -1
}

func Find() []*Strategy {
	var strategys []*Strategy
	models.DB.Find(&strategys)
	return strategys
}

func init() {
	models.DB.AutoMigrate(&Strategy{})
}
