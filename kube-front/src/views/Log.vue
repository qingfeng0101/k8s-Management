<template>
    <div class="container">
     <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
    <el-tab-pane :label=this.name name="logs"></el-tab-pane>
    <el-tab-pane label="返回" name="pods">返回</el-tab-pane>
  </el-tabs>
     <ul class="infinite-list"  style="overflow:auto">
    <li v-for="(i,index) in this.logs" class="infinite-list-item" :key=index>{{i}}</li>
  </ul>
    </div>
</template>
<script>
export default {
  data () {
    return {
      Socket: null,
      data: {},
      messages: {},
      count: 0,
      logs: [],
      activeName: 'logs',
      name: ''
    }
  },
  beforeMount () {
    // bus.$emit('maizuo', false)
    // this.$store.state.isShowtabelbar = false
    this.$store.commit('hildShowtabbar', false)
  },
  beforeDestroy () {
    // this.$store.state.isShowtabelbar = true
    this.$store.commit('Showtabbar', true)
  },
  mounted () {
    this.$store.commit('hildShowtabbar', false)
    var name = localStorage.getItem('name')
    var namespace = localStorage.getItem('namespace')
    this.name = name
    this.data['name'] = name
    this.data['namespace'] = namespace
    this.initSocket(name, namespace)
  },
  destroyed: function () {
    this.onClose()
  },
  methods: {
    initSocket (name, namespace) {
      const wsUrl = 'ws://192.168.0.105:8080/getpodlog'
      const ws = new WebSocket(wsUrl)
      this.Socket = ws
      this.Socket.onopen = this.onOpen
      this.Socket.onerror = this.onError
      this.Socket.onmessage = this.onMessage
      this.Socket.onclose = this.onClose
    },

    onOpen (name, namespace) {
      var jsondata = JSON.stringify(this.data)
      this.Socket.send(jsondata)
    },
    onMessage (e) {
      this.message = e
      if (this.logs.length > 99) {
        this.logs.shift()
      }

      this.logs.push(this.message.data)
    },
    onError (e) {
      console.log('Socket连接发生错误')
      console.log(e)
    },
    onClose (e) {
      this.Socket.close()
      console.log('Socket连接关闭')
    },
    handleClick (tab, event) {
      this.Socket.close()
      localStorage.setItem('activeName', tab.name)
      this.$store.commit('UpdateTabbarname', tab.name)
      this.$router.push(tab.name)
    }
  }

}
</script>
