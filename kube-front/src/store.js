import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        Tabbarname: localStorage.getItem('activeName') || 'GetHost',
        Data: [],
        Nodes: [],
        Pods: [],
        Pod: {},
        Deplyment: [],
        isShowtabelbar: true,
        ENV: ''
    },
    mutations: {
        UpdateTabbarname(state, data) {
            state.Tabbarname = data
            console.log("state", state.Tabbarname)
        },
        UpdateENV(state, data) {
            state.ENV = data
            console.log("state", state.ENV)
        },
        Updata(state, data) {
            if (data === null) {
                state.Data = []
            } else {
                state.Data = data
            }
        },
        UpdatePod(state, data) {
            state.Pod = data
        },
        hildShowtabbar(state, data) {
            state.isShowtabelbar = data
            console.log(state.isShowtabelbar)
        },
        Showtabbar(state, data) {
            state.isShowtabelbar = data
        }
    },
    actions: {
        Postdata(state, data) {
            axios({
                url: `http://192.168.0.105:8080${data.url}`,
                method: 'POST',
                data: {
                    'name': data.name,
                    'namespace': data.namespace,
                    'env': data.env,
                    'num': data.num
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
                // localStorage.removeItem("data")
                // localStorage.setItem("data", JSON.stringify(res.data.message))
                state.commit('Updata', res.data.message)
            })
        },
        DeleteData(state, data) {
            console.log("delete data: ", data)
            axios({
                url: `http://192.168.0.105:8080/${data.url}`,
                method: 'POST',
                data: {
                    'name': data.name,
                    'namespace': data.namespace,
                    'env': data.env
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
                state.commit('Updata', res.data.message)
            })
        },
        GetPodinfo(state, data) {
            axios({
                url: `http://192.168.0.105:8080${data.url}`,
                method: 'POST',
                data: {
                    'name': data.name,
                    'namespace': data.namespace,
                    'env': data.env
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
            })
        },
        DeleteDeployment(state, data) {
            console.log(data.name, data.namespace)
            axios({
                url: 'http://192.168.0.105:8080/deletedeployment',
                method: 'POST',
                data: {
                    'name': data.name,
                    'namespace': data.namespace
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
                state.commit('Updata', res.data.message)
            })
        },
        UpdataDeployment(state, data) {
            console.log(data.name, data.namespace)
            axios({
                url: `http://192.168.0.105:8080/${data.url}`,
                method: 'POST',
                data: {
                    'name': data.name,
                    'namespace': data.namespace,
                    'num': data.num
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
                state.commit('Updata', res.data.message)
            })
        }
    }

})