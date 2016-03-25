//@Description 乡音范例
//@Contact czw@outlook.com


package models

import (
	"github.com/astaxie/beego/orm"
)



type HomeVoiceExample struct {
	ProvinceID string
	Example []string
}



func (r *HomeVoiceExample) GetAll() []*HomeVoiceExample {
	var b []orm.Params
	db := ConnCommonDB()
	_, _ = db.Raw("SELECT  ProvinceID,Example  FROM  home_voice_example ORDER BY ProvinceID ASC").Values(&b)
	test := make(map[string]string,50)
	all := make([]*HomeVoiceExample,0)
	for _,v := range b {
		p := v["ProvinceID"].(string)
		if _,ok := test[p];ok {
			continue
		}
		test[p] = p
		pe := new(HomeVoiceExample)
		pe.ProvinceID = p
		pe.Example = make([]string,0)
		for _,x := range b {
				if pe.ProvinceID == x["ProvinceID"].(string) {
					
					pe.Example = append(pe.Example,  x["Example"].(string))
				}
		}
		all = append(all,pe)
	}
	return all
}


