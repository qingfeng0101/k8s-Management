<template>
<div>
<el-tabs v-model="activeName" type="card" @tab-click="handleClick">
    <el-tab-pane label="pods详情" name="pod">详情</el-tab-pane>
    <el-tab-pane label="返回" name="pods">返回</el-tab-pane>
  </el-tabs>

<table >
<span>
      <tr >
        <td >Name</td>
        <td >{{Pod.name}}</td>
      </tr>
      <tr >
        <td ><pre>Namespace    </pre></td>
        <td >{{Pod.namespace}}</td>
      </tr>
      <tr >
        <td >Labels</td>
        <td >
        <span v-for="(label,key) in Pod.labels" :key=key>
        {{key}} = {{label}}
        </span>
        </td>
      </tr>
    </span>
    <span   v-for="(Container,index) in Pod.containers" :key=index>
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
        <td  v-show=data_show>{{Container.Image}}</td>
        <td v-show=label_show><el-input v-model="input" :placeholder=Container.Image></el-input></td>
        <td ><el-button size="medium" @click=showdata>编辑</el-button></td>
        <td ><el-button size="medium" @click=showlabel(index) >保存</el-button></td>
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
import { mapState } from 'vuex'
export default {
  data () {
    return {
      activeName: 'pod',
      namespace: '',
      name: '',
      data: {},
      input: '',
      data_show: true,
      label_show: false,
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
  computed: {
    ...mapState(['Pod'])
  },
  mounted () {
    // console.log("mounted data",$store.state.Pods)
    this.$store.commit('hildShowtabbar', false)
    this.data['env'] = localStorage.getItem('ENV')
    this.namespace = localStorage.getItem('namespace')
    this.podname = localStorage.getItem('podname')
    this.data['name'] = this.podname
    this.data['namespace'] = this.namespace
    this.data['url'] = '/api/getpodinfo'
    this.$store.dispatch('GetPodinfo', this.data)
    console.log("pod: ",this.Pod)
    //this.$router.push("/pods")
     
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
    },
    showdata(){
      this.data_show = false
      this.label_show = true
    },
    showlabel(index){
      this.data_show = true
      this.label_show = false
      this.data['num'] = index
      this.data['namespace'] = this.namespace
      this.data['name'] = this.podname
      this.data['images'] = this.input
      this.data['url'] = '/api/updateimages'
      this.$store.dispatch('Postdata', this.data)
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
