import axios from 'axios'

const request = axios.create({
    baseURL: 'http://localhost:18703',
    timeout: 10000
})

request.interceptors.response.use(
    response => {
        return response.data
    }, () => {
        window.$message.error('网络异常')
        throw "网络异常";
    }
)

export default request
