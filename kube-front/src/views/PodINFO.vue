<template>
<div>
<el-tabs v-model="activeName" type="card" @tab-click="handleClick">
    <el-tab-pane label="pods列表" name="pod">详情</el-tab-pane>
    <el-tab-pane label="返回" name="pods">返回</el-tab-pane>
  </el-tabs>

<table >
<span>
      <tr >
        <td >Name</td>
        <td >{{$store.state.Pod.name}}</td>
      </tr>
      <tr >
        <td ><pre>Namespace    </pre></td>
        <td >{{$store.state.Pod.namespace}}</td>
      </tr>
      <tr >
        <td >Labels</td>
        <td >
        <span v-for="(label,key) in $store.state.Pod.labels" :key=key>
        {{key}} = {{label}}
        </span>
        </td>
      </tr>
    </span>
    <span   v-for="(Container,index) in $store.state.Pod.containers" :key=index>
      <tr >
        <td >Container Name</td>
        <td >{{Container.Name}}</td>
      </tr>
      <tr >
        <td >Container ID</td>
        <td>{{Container.ContainerID}}</td>
      </tr>
      <tr >
        <td >Image</td>
        <td>{{Container.Image}}</td>
      </tr>
      <tr >
        <td >Image ID</td>
         <td>{{Container.ImageID}}</td>
      </tr>
       <tr >
        <td >State</td>
        <td>{{status(Container.Status)}}</td>
      </tr>
      <tr >
        <td >Ready</td>
        <td>{{Container.Ready}}</td>
      </tr>
      <tr >
        <td >Restart Count</td>
        <td>{{Container.RestartNum}}</td>
      </tr>
      <tr >
        <td >Limits</td>
        <td>CPU={{Container.Lcpu}} MEM={{Container.Lmem}}</td>
      </tr>
      <tr >
        <td >Requests</td>
        <td>CPU={{Container.Rcpu}} MEM={{Container.Rmem}}</td>
      </tr>
    </span>
    </table>

</div>

</template>
<script>
export default {
  data () {
    return {
      activeName: 'pod',
      namespace: '',
      name: '',
      data: {}
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
    // console.log("mounted data",$store.state.Pods)
    this.$store.commit('hildShowtabbar', false)
     var ENV = localStorage.getItem('ENV')
    this.namespace = localStorage.getItem('namespace')
    this.podname = localStorage.getItem('podname')
    this.data['name'] = this.podname
    this.data['namespace'] = this.namespace
    if (ENV === "test"){
      this.data['url'] = 'getpodinfo'
    }else{
      this.data['url'] = 'prod/getpodinfo'
    }
    this.$store.dispatch('GetPodinfo', this.data)
    // this.$router.push("/pods")
     
  },
  methods: {
    status (data) {
      for (var k in data) {
        return k
      }
    },
    back () {
      this.$router.go(-1)// 返回上一层
    },
    handleClick (tab, event) {
      localStorage.setItem('activeName', tab.name)
      this.$store.commit('UpdateTabbarname', tab.name)
      this.$router.push(tab.name)
    }
  }
}
</script>
<style lang="scss" scoped>
span{
margin: 0;
    padding: 0;
    border: 0;
    font-size: 100%;
    font: inherit;
    vertical-align: baseline;
}
       table,table tr th, table tr td {

        border:1px solid #ccc;

         }

       table{

        width: 1100px;

        border-collapse: collapse;

    }

</style>
