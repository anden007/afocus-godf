// Code generated by hero.
// source: D:\GoProjects\github.com/anden007/afocus-godf\Backend\src\web\template\api.tpl
// DO NOT EDIT!
package template

import "bytes"

func Api(apiName string, apiPath string, vueName string, buffer *bytes.Buffer) {
	buffer.WriteString(`
// 统一请求路径前缀在libs/axios.js中修改
import { getRequest, postRequest, deleteRequest } from '@/libs/axios';

// 分页获取数据
export const get`)
	buffer.WriteString(apiName)
	buffer.WriteString(`List = (params) => {
    return getRequest('`)
	buffer.WriteString(apiPath)
	buffer.WriteString(`/getByCondition', params)
}
// 添加
export const add`)
	buffer.WriteString(apiName)
	buffer.WriteString(` = (params) => {
    return postRequest('`)
	buffer.WriteString(apiPath)
	buffer.WriteString(`/add', params)
}
// 编辑
export const edit`)
	buffer.WriteString(apiName)
	buffer.WriteString(` = (params) => {
    return postRequest('`)
	buffer.WriteString(apiPath)
	buffer.WriteString(`/edit', params)
}
// 删除
export const delete`)
	buffer.WriteString(apiName)
	buffer.WriteString(` = (ids, params) => {
    return deleteRequest('`)
	buffer.WriteString(apiPath)
	buffer.WriteString(`/delByIds/' + ids, params)
}`)

}
