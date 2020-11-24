package lib

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/anden007/afocus-godf/src/model"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type QueryCondition struct {
	Page      int       `json:"page"`
	Size      int       `json:"size"`
	Sort      string    `json:"sort"`
	Order     string    `json:"order"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

func NewQueryCondition() *QueryCondition {
	return &QueryCondition{}
}

func (m *QueryCondition) GetQuery(tableName string, condition map[string]string, query *gorm.DB, countQuery *gorm.DB) (err error) {
	m.Page = 0
	m.Size = 0
	m.Sort = ""
	m.Order = ""
	m.StartDate = time.Time{}
	m.EndDate = time.Time{}
	//清理前端传入的没有数据的空条件
	for key, value := range condition {
		val := strings.TrimSpace(value)
		if val == "" {
			delete(condition, key)
		}
	}
	//处理分页参数
	if _, hasValue := condition["pageNumber"]; hasValue {
		m.Page = cast.ToInt(condition["pageNumber"])
		delete(condition, "pageNumber")
	}
	if _, hasValue := condition["pageSize"]; hasValue {
		m.Size = cast.ToInt(condition["pageSize"])
		delete(condition, "pageSize")
	}
	if _, hasValue := condition["sort"]; hasValue {
		m.Sort = condition["sort"]
		delete(condition, "sort")
	}
	if _, hasValue := condition["order"]; hasValue {
		m.Order = condition["order"]
		delete(condition, "order")
	}
	if _, hasValue := condition["startDate"]; hasValue {
		m.StartDate, _ = time.ParseInLocation("2006-01-02", condition["startDate"], time.Local)
		delete(condition, "startDate")
	}
	if _, hasValue := condition["endDate"]; hasValue {
		m.EndDate, _ = time.ParseInLocation("2006-01-02", condition["endDate"], time.Local)
		delete(condition, "endDate")
	}
	if !m.StartDate.IsZero() && !m.EndDate.IsZero() {
		query = query.Where("create_time >= ?", m.StartDate).Where("create_time <= ?", m.EndDate)
		countQuery = countQuery.Where("create_time >= ?", m.StartDate).Where("create_time <= ?", m.EndDate)
	}
	DBFieldNames, isExists := model.ModelDBFields[tableName]
	if isExists {
		// 常规参数
		for key, value := range condition {
			val := strings.TrimSpace(value)
			if val != "" {
				fieldName := DBFieldNames[key]
				if fieldName != "" {
					query = query.Where(fmt.Sprintf("%s = ?", fieldName), value)
					countQuery = countQuery.Where(fmt.Sprintf("%s = ?", fieldName), value)
				}
			}
		}
		// 处理排序
		if m.Sort != "" && m.Order != "" {
			fieldName := DBFieldNames[m.Sort]
			if fieldName == "" {
				err = errors.New(fmt.Sprintf("%s不是有效的排序字段", m.Sort))
			} else {
				if m.Page > 0 && m.Size > 0 {
					query.Order(fmt.Sprintf("%s %s", fieldName, m.Order)).Offset((m.Page - 1) * m.Size).Limit(m.Size)
				} else {
					query.Order(fmt.Sprintf("%s %s", fieldName, m.Order))
				}
			}
		} else {
			if m.Page > 0 && m.Size > 0 {
				query.Offset((m.Page - 1) * m.Size).Limit(m.Size)
			}
		}
	} else {
		query = query.Where(fmt.Sprintf("%s = ?", "true"), false)
		countQuery = countQuery.Where(fmt.Sprintf("%s = ?", "true"), false)
	}
	return
}
