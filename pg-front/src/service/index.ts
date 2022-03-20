import YOLORequest from './request'
// 拿到不同环境下定义的常量
import { BASE_URL, TIME_OUT } from './request/config'
import localCache from '@/utils/cache'

const yoloRequest = new YOLORequest({
  baseURL: BASE_URL,
  timeout: TIME_OUT,
  interceptors: {
    requestInterceptor: (config) => {
      // 携带token
      const token = localCache.getCache('token')
      if (token) {
        config.headers!.Authorization = `Bearer ${token}`
      }
      return config
    },
    requestInterceptorCatch: (err) => {
      return err
    },
    responseInterceptor: (res) => {
      return res
    },
    responseInterceptorCatch: (err) => {
      return err
    }
  }
})

export default yoloRequest
