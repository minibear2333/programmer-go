import axios from 'axios'
import type { AxiosInstance } from 'axios'
import type { YOLORequestInterceptors, YOLORequestConfig } from './type'
// 用ant的loading, 这里待开发
// import { ElLoading } from 'element-plus/lib/components/loading/index'
// import { LoadingInstance } from 'element-plus/lib/components/loading/src/loading'

const DEFAULT_LOADING = true

class YOLORequest {
  instance: AxiosInstance
  interceptors?: YOLORequestInterceptors
  // loading?: LoadingInstance
  showLoading: boolean
  // 使用构造器初始化
  constructor(config: YOLORequestConfig) {
    // 创建axios实例
    this.instance = axios.create(config)
    this.interceptors = config.interceptors
    this.showLoading = config.showLoading ?? DEFAULT_LOADING
    // 传入接受拦截器
    this.instance.interceptors.request.use(
      this.interceptors?.requestInterceptor,
      this.interceptors?.requestInterceptorCatch
    )
    // 传入响应拦截器
    this.instance.interceptors.response.use(
      this.interceptors?.responseInterceptor,
      this.interceptors?.responseInterceptorCatch
    )

    // 公共拦截器
    this.instance.interceptors.request.use(
      (config) => {
        if (this.showLoading) {
          // this.loading = Spin.service({
          //   lock: true,
          //   text: '正在加载中喔...',
          //   background: 'rgba(0, 0, 0, 0.5)'
          // })
        }

        return config
      },
      (err) => {
        return err
      }
    )
    this.instance.interceptors.response.use(
      (res) => {
        // 将loading移除
        // this.loading?.close()

        // 方式二
        const data = res.data
        if (data.returnCode === '-1001') {
          console.log('请求失败')
        } else {
          return data
        }
      },
      (err) => {
        // 将loading移除
        // this.loading?.close()
        // 方式一
        if (err.response.status === 404) {
          console.log('404错误')
        }
        return err
      }
    )
  }

  request<T>(config: YOLORequestConfig<T>): Promise<T> {
    return new Promise((resolve, reject) => {
      // 判断请求是否有拦截器, 若有则调用里面的函数, 再拿到返回的config
      if (config.interceptors?.requestInterceptor) {
        config = config.interceptors.requestInterceptor(config)
      }
      // 当不需要loading时
      if (config.showLoading === false) {
        this.showLoading = config.showLoading
      }

      this.instance
        .request<any, T>(config)
        .then((res) => {
          if (config.interceptors?.responseInterceptor) {
            res = config.interceptors.responseInterceptor(res)
          }
          this.showLoading = DEFAULT_LOADING

          resolve(res)
        })
        .catch((err) => {
          this.showLoading = DEFAULT_LOADING
          reject(err)
          return err
        })
    })
  }

  get<T>(config: YOLORequestConfig<T>): Promise<T> {
    return this.request<T>({
      ...config,
      method: 'GET'
    })
  }
  post<T>(config: YOLORequestConfig<T>): Promise<T> {
    return this.request<T>({
      ...config,
      method: 'POST'
    })
  }
  delete<T>(config: YOLORequestConfig<T>): Promise<T> {
    return this.request<T>({
      ...config,
      method: 'DELETE'
    })
  }
  patch<T>(config: YOLORequestConfig<T>): Promise<T> {
    return this.request<T>({
      ...config,
      method: 'PATCH'
    })
  }
}

// 暴露封装类
export default YOLORequest
