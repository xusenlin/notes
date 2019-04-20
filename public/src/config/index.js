/**
 * Created by xusenlin on 2019/4/13.
 */

const devApiUrl = 'http://localhost:8081';

//正式环境变量,注意修改
const proApiUrl = 'https://pro.web.com';


const nodeDevEnv = process.env.NODE_ENV=='development' ? true : false;
export default {
    apiUrl : nodeDevEnv ? devApiUrl : proApiUrl,
    apiPrefix : "api",
    timeout:1000,
    accessTokenKey:'ACCESS_TOKEN',
    requestRetry:4,
    requestRetryDelay:800,
}
