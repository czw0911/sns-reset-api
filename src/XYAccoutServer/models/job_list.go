//@Description 职业列表
//@Contact czw@outlook.com


package models

import (
	"github.com/astaxie/beego/orm"
)



type JobList struct {
	JobId string
	JobName string
}

type JobBigClass struct {
	JobId string
	JobName string
	JobSub []*JobSmallClass
}

type JobSmallClass struct {
	JobId string
	JobName string
}






func (r *JobList) GetAll() []*JobBigClass {
	var b []orm.Params
	var s []orm.Params
	db := ConnCommonDB()
	_, _ = db.Raw("SELECT JobID,JobName  FROM  job_list_large").Values(&b)
	_, _ = db.Raw("SELECT JobID, JobName, JobLargeID  FROM  job_list_small ").Values(&s)
	
	all := make([]*JobBigClass,0)
	for _,v := range b {
		pe := new(JobBigClass)
		pe.JobId = v["JobID"].(string)
		pe.JobName = v["JobName"].(string)
		pe.JobSub = make([]*JobSmallClass,0)
		for _,x := range s {
				if pe.JobId == x["JobLargeID"].(string) {
					cy := new(JobSmallClass)
					cy.JobId = x["JobID"].(string)
					cy.JobName = x["JobName"].(string)
					pe.JobSub = append(pe.JobSub, cy)
				}
		}
		all = append(all,pe)
	}
	
	return all
}


