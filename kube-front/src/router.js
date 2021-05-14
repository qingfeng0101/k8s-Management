import Vue from 'vue'
import Router from 'vue-router'
import Test from './views/Test'
import Pods from './views/Pods'
import Podinfo from './views/PodINFO'
import Log from './views/Log'
const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
}

Vue.use(Router)

export default new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [{
            path: '',
            name: 'test',
            component: Test
        },
        {
            path: '/pods',
            name: 'Pods',
            component: Pods
        },
        {
            path: '/podinfo',
            name: 'Podinfo',
            component: Podinfo
        },
        {
            path: '/logs',
            name: 'Log',
            component: Log
        }

    ]
})