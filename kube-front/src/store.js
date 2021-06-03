import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        Tabbarname: 'GetHost',
        Data: [],
        isshow: true,
        show: true,
        Env: [],
        Pod: {},
        Deplyment: [],
        isShowtabelbar: true,
        ENV: ''
    },
    mutations: {
        UpdateTabbarname(state, data) {
            state.Tabbarname = data
        },
        UpdateENV(state, data) {
            state.ENV = data
            console.log("ENV: ", state.ENV)
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
        Updateenv(state, data) {
            state.Env = data
        },
        hildShowtabbar(state, data) {
            state.isShowtabelbar = data
            console.log(state.isShowtabelbar)
        },
        Showtabbar(state, data) {
            state.isShowtabelbar = data
        },
        hildShow(state, data) {
            state.isshow = data
            console.log("sishow: ", state.isshow)
        },
        hildshow(state, data) {
            state.show = data
            console.log("show: ", state.show)
        }
    },
    actions: {
        Postdata(state, data) {
            axios({
                url: `http://127.0.0.1:8080${data.url}`,
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
                url: `http://127.0.0.1:8080/${data.url}`,
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
                url: `http://127.0.0.1:8080${data.url}`,
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
        Uploadenv(state, data) {
            let param = new FormData()
            console.log("data: ", data.data)
            param.append('files', data.data.file)
            param.append('testname', data.name)
                // axios.post(`http://192.168.0.105:8080${data.url}`, data).then(res => {
                //     console.log(res)
                // })
            axios({
                method: 'post',
                url: `http://127.0.0.1:8080${data.url}`,
                headers: { 'Content-Type': 'multipart/form-data' },
                data: param
            }).then(res => {
                state.commit('Updateenv', res.data.message)
            })
        },
        Getenv(state, data) {
            // axios.post(`http://192.168.0.105:8080${data.url}`, data).then(res => {
            //     console.log(res)
            // })
            axios({
                method: 'get',
                url: `http://127.0.0.1:8080${data.url}`,
            }).then(res => {
                console.log('Updateenv: ', res.data.message)
                if (res.data.message === null) {
                    return
                } else {
                    state.commit('Updateenv', res.data.message)
                }

            })
        },
        UpdataDeployment(state, data) {
            console.log(data.name, data.namespace)
            axios({
                url: `http://127.0.0.1:8080/${data.url}`,
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