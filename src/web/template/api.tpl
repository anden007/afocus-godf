<%: func Api(apiName string, apiPath string, vueName string, buffer *bytes.Buffer) %>
// 统一请求路径前缀在libs/axios.js中修改
import { getRequest, postRequest, deleteRequest } from '@/libs/axios';

// 分页获取数据
export const get<%==s apiName%>List = (params) => {
    return getRequest('<%==s apiPath%>/getByCondition', params)
}
// 添加
export const add<%==s apiName%> = (params) => {
    return postRequest('<%==s apiPath%>/add', params)
}
// 编辑
export const edit<%==s apiName%> = (params) => {
    return postRequest('<%==s apiPath%>/edit', params)
}
// 删除
export const delete<%==s apiName%> = (ids, params) => {
    return deleteRequest('<%==s apiPath%>/delByIds/' + ids, params)
}