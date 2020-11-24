<%: func TreeApi(apiName string, apiPath string, vueName string, buffer *bytes.Buffer) %>
// 统一请求路径前缀在libs/axios.js中修改
import { getRequest, postRequest, deleteRequest } from '@/libs/axios';

// 获取一级数据
export const init<%==s apiName%> = (params) => {
    return getRequest('<%==s apiPath%>/getByParentId/0', params)
}
// 加载子级数据
export const load<%==s apiName%> = (id, params) => {
    return getRequest('<%==s apiPath%>/getByParentId/' + id, params)
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
// 搜索
export const search<%==s apiName%> = (params) => {
    return getRequest('<%==s apiPath%>/search', params)
}