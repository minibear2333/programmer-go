import type { App } from 'vue'
// 引入base.css
import 'ant-design-vue/dist/antd.css'
import { DatePicker } from 'ant-design-vue/lib/components'

const components = [DatePicker]

export default function (app: App): void {
  for (const component of components) {
    app.component(component.name, component)
  }
}
