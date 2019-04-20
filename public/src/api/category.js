/**
 * Created by xusenlin on 2019/4/13.
 */

import request from '../utils/request.js'

export function getAllList(params) {
    return request({
        url: '/category',
        method: 'get',
        params:params
    })
}