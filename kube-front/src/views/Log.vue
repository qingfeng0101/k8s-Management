<template>
    <div class="container">
     <el-button-group>
  <span v-on:click="back"><el-button size="mini" type="primary" icon="el-icon-arrow-left">返回</el-button></span>
</el-button-group>
     <ul class="infinite-list"  style="overflow:auto">
    <li v-for="(i,index) in this.logs" class="infinite-list-item" :key=index>{{i}}</li>
  </ul>
    </div>
</template>
<script>
export default {
    data() {
        return {
            Socket: null,
            data:{},
            messages: {},
            count: 0,
            logs: []
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
        var name = localStorage.getItem("name")
        var namespace = localStorage.getItem("namespace")
        this.data["name"] = name
        this.data["namespace"] = namespace
         this.initSocket(name,namespace)
    },
    destroyed: function() {
        this.onClose()
    },
    methods: {
        initSocket(name,namespace) {
            const wsUrl = 'ws://localhost:8080/getpodlog';
            const ws = new WebSocket(wsUrl)
            this.Socket = ws;
            this.Socket.onopen = this.onOpen;
            this.Socket.onerror = this.onError;
            this.Socket.onmessage = this.onMessage;
            this.Socket.onclose = this.onClose;
        },
        
        onOpen(name,namespace){
           
            var jsondata = JSON.stringify(this.data)
            this.Socket.send(jsondata)
        },
        onMessage(e){
            this.message = e
            if (this.logs.length > 99){
                this.logs.shift()
            }

            this.logs.push(this.message.data)
           
        },
        onError(e){
            console.log("Socket连接发生错误")
            console.log(e)
        },
        onClose(e){
            this.Socket.close()
            console.log("Socket连接关闭")
        },
        back(){
        this.Socket.onclose = this.onClose;
        this.$router.go(-1);//返回上一层
    },
    
   
        
    }
}
</script>