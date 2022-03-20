import type { AxiosRequestConfig, AxiosResponse } from 'axios'

export interface YOLORequestInterceptors<T = AxiosResponse> {
  // 发送成功
  requestInterceptor?: (config: AxiosRequestConfig) => AxiosRequestConfig
  // 发送失败
  requestInterceptorCatch?: (error: any) => any
  // 响应成功
  responseInterceptor?: (config: T) => T
  // 响应失败
  responseInterceptorCatch?: (error: any) => any
}

export interface YOLORequestConfig<T = AxiosResponse>
  extends AxiosRequestConfig {
  interceptors?: YOLORequestInterceptors<T>
  showLoading?: boolean
}
