import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        Test: [],
        Pods: [],
        Pod: {},
        isShowtabelbar: true,
    },
    mutations: {
        UpdateTest(state, data) {
            state.Test = data
        },
        UpdatePods(state, data) {
            state.Pods = data
        },
        UpdatePod(state, data) {
            state.Pod = data
        },
        hildShowtabbar(state, data) {
            state.isShowtabelbar = data
        },
        Showtabbar(state, data) {
            state.isShowtabelbar = data
        },
    },
    actions: {
        Gettest(state) {
            axios({
                url: 'http://localhost:8080/namespace'
            }).then(res => {
                console.log(res.data.message)
                state.commit('UpdateTest', res.data.message)
            })
        },
        Postpods(state, data) {
            axios({
                url: 'http://localhost:8080/pod',
                method: 'POST',
                data: {
                    'namespace': data
                },
                transformRequest: [function(data) {
                    // Do whatever you want to transform the data
                    let ret = ''
                    for (let it in data) {
                        // 如果要发送中文 编码
                        ret += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
                    }
                    return ret
                }],
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                }
            }).then(res => {
                state.commit('UpdatePods', res.data.message)
            })
        },
        DeletePods(state, { name, namespace }) {
            axios({
                url: 'http://localhost:8080/deletepod',
                method: 'POST',
                data: {
                    'name': name,
                    'namespace': namespace
                },
                transformRequest: [function(data) {
                    // Do whatever you want to transform the data
                    let ret = ''
                    for (let it in data) {
                        // 如果要发送中文 编码
                        ret += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
                    }
                    return ret
                }],
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                }
            }).then(res => {
                state.commit('UpdatePods', res.data.message)
            })
        },
        GetPodinfo(state, { name, namespace }) {
            axios({
                url: 'http://localhost:8080/getpodinfo',
                method: 'POST',
                data: {
                    'name': name,
                    'namespace': namespace
                },
                transformRequest: [function(data) {
                    // Do whatever you want to transform the data
                    let ret = ''
                    for (let it in data) {
                        // 如果要发送中文 编码
                        ret += encodeURIComponent(it) + '=' + encodeURIComponent(data[it]) + '&'
                    }
                    return ret
                }],
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                }
            }).then(res => {
                state.commit('UpdatePod', res.data.message)
                console.log(res.data.message)
            })
        }
    }
})