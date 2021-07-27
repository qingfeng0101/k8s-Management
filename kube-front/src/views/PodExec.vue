<template>
<div>
  <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
    <el-tab-pane :label=this.name name="podexec"></el-tab-pane>
    <el-tab-pane label="返回" name="pods">返回</el-tab-pane>
  </el-tabs>
 <div id="xterm" class="xterm" />
</div>
   
</template>
<script>
import 'xterm/css/xterm.css'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'


export default {
    data () {
        return {
            data: {},
            Socket: null,
             messages: {},
             term: null,
             name: null,
             activeName: "podexec", 
             fontsize: 24
        }
    },
  name: 'Xterm',
  props: {
    socketURI: {
      type: String,
      default: ''
    },
  },
  beforeMount () {
    // bus.$emit('maizuo', false)
    // this.$store.state.isShowtabelbar = false
    this.$store.commit('hildShowtabbar', false)
  },
  beforeDestroy () {
    // this.$store.state.isShowtabelbar = true
     this.socket.close()
    this.term.dispose()
    this.$store.commit('Showtabbar', true)
  },
  mounted() {
      this.$store.commit('hildShowtabbar', false)
    this.data["env"] = localStorage.getItem('ENV')
    this.data["cname"] = localStorage.getItem('cname')
    var name = localStorage.getItem('name')
    var namespace = localStorage.getItem('namespace')
    this.name = name
    this.data['name'] = name
    this.data['namespace'] = namespace
    this.data['url'] = '/api/exec'
    this.data['cmd'] = ["sh"]
    this.initSocket()
  },
 
  methods: {
    initXterm() {
      this.term = new Terminal({
        rendererType: "canvas", //渲染类型
        rows: 35, //行数
        convertEol: true, //启用时，光标将设置为下一行的开头
        scrollback: 10, //终端中的回滚量
        disableStdin: false, //是否应禁用输入
        cursorStyle: "underline", //光标样式
        cursorBlink: true, //光标闪烁
        fontSize: this.fontsize,
        theme: {
          foreground: "yellow", //字体
          background: "#060101", //背景色
          cursor: "help", //设置光标
          
        }
      });
      this.term.open(document.getElementById("xterm"));
      const fitAddon = new FitAddon();
      this.term.loadAddon(fitAddon);
      // 支持输入与粘贴方法
      let _this = this; //一定要重新定义一个this，不然this指向会出问题
      this.term.onData(function(key) {
      //let order = ["stdin",key]; //这里key值是你输入的值，数据格式一定要找后端要！！！！
      let order = {operation:"stdin",data:key}
        _this.socket.onsend(JSON.stringify(order)); //转换为字符串
      });
    },
    initSocket() {
      const wsUrl = `ws://127.0.0.1:8080${this.data.url}`
      this.socket = new WebSocket(wsUrl);
      // 监听socket连接
      this.socket.onopen = this.open;
      // 监听socket错误信息
      this.socket.onerror = this.error;
      // 监听socket消息
      this.socket.onmessage = this.getMessage;
      // 发送socket消息
      this.socket.onsend = this.send;
     
    },
     open: function() {
       var jsondata = JSON.stringify(this.data)
      this.socket.send(jsondata)
      console.log("socket连接成功");
      this.initXterm();
    },
    error: function() {
      console.log("连接错误");
    },
    close: function() {
      this.socket.close();
      console.log("socket已经关闭");
    },
    getMessage: function(msg) {
      console.log(JSON.parse(msg.data))
      let mess = JSON.parse(msg.data)
      console.log(mess.data)
      this.term.write(mess.data);
    },
    send: function(order) {
      console.log(order)
      this.socket.send(order);
    },
    handleClick (tab, event) {
       this.socket.close()
      localStorage.setItem('activeName', tab.name)
      this.$store.commit('UpdateTabbarname', tab.name)
      this.$router.push(tab.name)
    }
      
    }
}
</script>