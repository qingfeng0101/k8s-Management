import Vue from 'vue'
import Router from 'vue-router'
import Pods from './views/Pods'
import Podinfo from './views/PodINFO'
import Log from './views/Log'
import Namespace from './views/Namespace'
import Host from './views/Host'
import Deployment from './views/Deployment'
import Uploadenv from './views/Uploadenv'
import Nodeinfo from './views/nodeinfo'
import Deploymentinfo from './views/DeployemntINFO'
import Podexec from './views/PodExec'
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
            component: Host
        },

        {
            path: '/GetHost',
            name: 'Host',
            component: Host
        },
        {
            path: '/Getnamespace',
            name: 'Namespace',
            component: Namespace
        },
        {
            path: '/pods',
            name: 'Pods',
            component: Pods
        },
        {
            path: '/deployment',
            name: 'Deployment',
            component: Deployment
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
        },
        {
            path: '/uploadenv',
            name: 'Uploadenv',
            component: Uploadenv
        },
        {
            path: '/ClusterInfo',
            name: 'Nodeinfo',
            component: Nodeinfo
        },
        {
            path: '/deploymentInfo',
            name: 'Deploymentinfo',
            component: Deploymentinfo
        },
        {
            path: '/podexec',
            name: 'Podexec',
            component: Podexec
        },
    ]
})