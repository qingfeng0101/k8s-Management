import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    Tabbarname: 'host',
    Data: [],
    Nodes: [],
    Pods: [],
    Pod: {},
    Deplyment: [],
    isShowtabelbar: true
  },
  mutations: {
    UpdateTabbarname (state, data) {
      state.Tabbarname = data
    },
    Updata (state, data) {
      state.Data = data
    },
    UpdateNamespace (state, data) {
      state.Namespace = data
    },
    UpdatePods (state, data) {
      state.Pods = data
    },
    UpdateDeployment (state, data) {
      state.Deplyment = data
    },
    UpdatePod (state, data) {
      state.Pod = data
    },
    hildShowtabbar (state, data) {
      state.isShowtabelbar = data
      console.log(state.isShowtabelbar)
    },
    Showtabbar (state, data) {
      state.isShowtabelbar = data
    }
  },
  actions: {
    Getnamespace (state) {
      axios({
        url: 'http://192.168.0.105:8080/namespace'
      }).then(res => {
        state.commit('Updata', res.data.message)
      })
    },
    GetHost (state) {
      axios({
        url: 'http://192.168.0.105:8080/nodes'
      }).then(res => {
        state.commit('Updata', res.data.message)
      })
    },
    Postpods (state, data) {
      axios({
        url: 'http://192.168.0.105:8080/pod',
        method: 'POST',
        data: {
          'namespace': data
        },
        transformRequest: [function (data) {
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
    DeletePods (state, data) {
      axios({
        url: 'http://192.168.0.105:8080/deletepod',
        method: 'POST',
        data: {
          'name': data.name,
          'namespace': data.namespace
        },
        transformRequest: [function (data) {
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
    GetPodinfo (state, data) {
      axios({
        url: 'http://192.168.0.105:8080/getpodinfo',
        method: 'POST',
        data: {
          'name': data.name,
          'namespace': data.namespace
        },
        transformRequest: [function (data) {
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
    Postdeployment (state, data) {
      axios({
        url: 'http://192.168.0.105:8080/getdeplyment',
        method: 'POST',
        data: {
          'namespace': data
        },
        transformRequest: [function (data) {
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
    DeleteDeployment (state, data) {
      console.log(data.name, data.namespace)
      axios({
        url: 'http://192.168.0.105:8080/deletedeployment',
        method: 'POST',
        data: {
          'name': data.name,
          'namespace': data.namespace
        },
        transformRequest: [function (data) {
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
    UpdataDeployment (state, data) {
      console.log(data.name, data.namespace)
      axios({
        url: 'http://192.168.0.105:8080/updatadeployment',
        method: 'POST',
        data: {
          'name': data.name,
          'namespace': data.namespace,
          'num': data.num
        },
        transformRequest: [function (data) {
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
