// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	crmRolesFieldNames          = builder.RawFieldNames(&CrmRoles{})
	crmRolesRows                = strings.Join(crmRolesFieldNames, ",")
	crmRolesRowsExpectAutoSet   = strings.Join(stringx.Remove(crmRolesFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	crmRolesRowsWithPlaceHolder = strings.Join(stringx.Remove(crmRolesFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	crmRolesModel interface {
		Insert(ctx context.Context, data *CrmRoles) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*CrmRoles, error)
		Update(ctx context.Context, data *CrmRoles) error
		UpdateEnable(ctx context.Context, id int64 , is_enable int64) error
		Delete(ctx context.Context, id int64) error
		IshaveByName(ctx context.Context, name string,level int64) (bool,error)
		Detail(ctx context.Context,id int64) (DetailWithChilds,error)
		List(ctx context.Context,name string , level int64 , is_enable int64 , parent_id int64, page int64) (PageList,error)
	}

	defaultCrmRolesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CrmRoles struct {
		Id       int64          `db:"id"`
		Name     string         `db:"name"`
		IsEnable int64				  `db:"is_enable"`
		Level    int64				  `db:"level"`
		ParentId int64				  `db:"parent_id"`
		CreateAt string			 	  `db:"create_at"`
		UpdateAt string			    `db:"update_at"`
	}

	DetailWithChilds struct {
		Id       int64
		Name     string
		IsEnable int64
		Level    int64
		ParentId int64
		CreateAt string
		UpdateAt string
		Childs   []CrmRoles
	}

)

func newCrmRolesModel(conn sqlx.SqlConn) *defaultCrmRolesModel {
	return &defaultCrmRolesModel{
		conn:  conn,
		table: "`crm_roles`",
	}
}

func (m *defaultCrmRolesModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCrmRolesModel) FindOne(ctx context.Context, id int64) (*CrmRoles, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", crmRolesRows, m.table)
	var resp CrmRoles
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCrmRolesModel) Insert(ctx context.Context, data *CrmRoles) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, crmRolesRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.IsEnable, data.Level, data.ParentId, data.CreateAt, data.UpdateAt)
	return ret, err
}

func (m *defaultCrmRolesModel) Update(ctx context.Context, data *CrmRoles) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, crmRolesRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.IsEnable, data.Level, data.ParentId, data.CreateAt, data.UpdateAt, data.Id)
	return err
}

func (m *defaultCrmRolesModel) UpdateEnable(ctx context.Context, id int64 , is_enable int64) error {
	query := fmt.Sprintf("update %s set is_enable = ? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query , is_enable , id)
	return err
}

func (m *defaultCrmRolesModel) IshaveByName(ctx context.Context, name string, level int64) (bool,error) {
	// query := fmt.Sprintf("select id from %s where `name` = ? limit 1 ", m.table)
	query := fmt.Sprintf("select count(*) from %s where `name` = ? and level = ?", m.table)
	var resp int64
	err := m.conn.QueryRowCtx(ctx, &resp, query, name , level)
	return resp > 0, err
}


func (m *defaultCrmRolesModel) Detail(ctx context.Context, id int64) (DetailWithChilds,error) {
	var result DetailWithChilds

	query := fmt.Sprintf("select * from %s where id = ?", m.table)
	var resp CrmRoles
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	if err != nil {
		if err == ErrNotFound {
			return  result, errors.New("未找到对应记录")
		}
		return result, errors.New("查询失败")
	}

	// 赋值
	result.Id  = resp.Id
	result.Name = resp.Name
	result.IsEnable = resp.IsEnable
	result.ParentId = resp.ParentId
	result.CreateAt = resp.CreateAt
	result.UpdateAt = resp.UpdateAt

	var resps []CrmRoles
	query1 := fmt.Sprintf("select * from %s where parent_id = ?", m.table)
	err = m.conn.QueryRowsCtx(ctx, &resps, query1, id)
	if err == nil {
		result.Childs = resps
	}

	return result,err
}

func (m *defaultCrmRolesModel) List(ctx context.Context,name string , level int64 , is_enable int64 , parent_id int64, page int64)(result PageList,err error) {
	var resps []CrmRoles
	query := fmt.Sprintf("select * from %s where 1 = 1", m.table)
	query1 := fmt.Sprintf("select count(*) from %s where 1 = 1", m.table)

	if name != "" {
		query += " and name like '%"+ name +"%'"
		query1 += " and name like '%"+ name +"%'"
	}

	if level != 0 {
		query += " and level = "+ strconv.FormatInt(level,10)
		query1 += " and level = "+ strconv.FormatInt(level,10)
	}

	if is_enable != -1 {
		query += " and is_enable = " + strconv.FormatInt(is_enable,10)
		query1 += " and is_enable = " + strconv.FormatInt(is_enable,10)
	}

	if parent_id != 0 {
		query += " and parent_id = "+ strconv.FormatInt( parent_id , 10 )
		query1 += " and parent_id = "+ strconv.FormatInt( parent_id , 10 )
	}

	result.Size = 5
	result.Page = page
	m.conn.QueryRowCtx(ctx, &result.Total , query1)
	result.TotalPage = int64(math.Ceil( float64( result.Total ) / float64( result.Size ) ))

	if result.Total > 0 && result.TotalPage >= result.Page  {
		query += " limit "+ strconv.FormatInt( (result.Page - 1) * result.Size , 10 ) +" ," + strconv.FormatInt(result.Size,10)
		fmt.Println("========")
		fmt.Println(query)
		fmt.Println("========")
		err = m.conn.QueryRowsCtx(ctx, &resps, query)
		result.List = resps	
	} else {
		result.List =  make([]int,0)
		// result.List =  make(map[string]interface{},0)
	}

	

	return result , err
}

func (m *defaultCrmRolesModel) tableName() string {
	return m.table
}
