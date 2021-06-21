import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import echarts from 'echarts' //引入echarts
Vue.prototype.$echarts = echarts //挂载在原型，组件内使用直接this.$echarts调用
Vue.config.productionTip = false
Vue.use(ElementUI)
new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')