/**
 * Created by xusenlin on 2019/4/13.
 */

import Axios from 'axios'
import Config from '../config/index.js'

const service = Axios.create({
    baseURL: Config.apiUrl + '/' + Config.apiPrefix,
    headers: {
        'Accept': '*/*'
    },
    timeout: Config.timeout
})

service.defaults.retry = Config.requestRetry;
service.defaults.retryDelay = Config.requestRetryDelay;


service.interceptors.response.use(
        response => {

            const res = response
            if (res.status !== 200) {
                //Toast('数据返回出错');
                return Promise.reject('error')
            } else {
                return res.data
            }
        },
        error => {
            //Toast("请求未响应");
            return Promise.reject(error)
        }
)

export default service