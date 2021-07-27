<template>
  <nav >

<el-tabs   v-model="activeName" type="card" @tab-click="handleClick">
    <el-tab-pane label="集群信息" name='ClusterInfo'>主机信息</el-tab-pane>
    <el-tab-pane label="主机信息" name='GetHost'>主机信息</el-tab-pane>
    <el-tab-pane label="命名空间" name='Getnamespace'></el-tab-pane>
    <el-tab-pane label="环境增加" name='Uploadenv'></el-tab-pane>
  </el-tabs>
     <div v-show='isshow'>
     <el-radio  v-for="(v,index) in data" :key="index" v-model="radio1" :label="v.name"   @change="changes" border>{{v.name}}</el-radio>
     
     </div> 
    

  </nav>
</template>
<script>
import { mapState } from 'vuex'
export default {
  data () {
    return {
      activeName: this.Tabbarname,
       radio1: '',
       data: []
    }
  },
  computed: {
    ...mapState(['Env','isshow','ENV','Tabbarname','show'])
  },
 
 created () {
     var data = {}
      data['url'] = '/api/names'
     this.$store.dispatch('Getenv', data).then((res) => {
       if (localStorage.getItem('ENV') === null  && localStorage.getItem('activeName') === null ){
          if (res.data.message != null){
             this.data = res.data.message
             this.radio1 =  this.data[0].name
             this.activeName=this.Tabbarname
             localStorage.setItem('ENV',this.radio1)
             localStorage.setItem('activeName',this.activeName)
             this.$store.commit('UpdateENV', this.radio1)
              if (this.activeName ==='Uploadenv' ){
             this.$store.commit('hildShow', false)
          }else{
             this.$store.commit('hildShow', true)
          }
           
        }else{
           this.activeName = localStorage.getItem('activeName')
          this.$store.commit('hildShow', false)
          this.data = []
        }
       }else{
         if (res.data.message != null){
           this.data = res.data.message
         this.radio1 = localStorage.getItem('ENV')
         this.activeName = localStorage.getItem('activeName')
          this.$store.commit('UpdateENV', this.radio1)
          if (this.activeName ==='Uploadenv' ){
             this.$store.commit('hildShow', false)
          }else{
             this.$store.commit('hildShow', true)
          }
         
         
         }else{
            this.activeName = localStorage.getItem('activeName')
          this.$store.commit('hildShow', false)
          this.data = []
        }
       }
     })
     
   
    
 },
   beforeDestroy () {
  // this.$store.state.isShowtabelbar = true
    this.$store.commit('Showtabbar', true)
  },
 
  methods: {
    handleClick (tab, event) {
      // this.activeName = tab.name
      if (tab.name === 'Uploadenv'){
           this.$store.commit('hildShow', false)
      }else{
        this.$store.commit('hildShow', true)
      }
      var data = {}
      data['url'] = '/api/names'
     this.$store.dispatch('Getenv', data).then((res) => {
       if (res.data.message != null){
         this.data = res.data.message
       }else{
         this.data = [];
         this.$store.commit('hildShow', false)
       }
     })
      console.log("test1")
      this.$store.commit('UpdateENV', localStorage.getItem('ENV'))
      localStorage.setItem('activeName', tab.name)
      this.radio1 = this.ENV
      this.$store.commit('UpdateTabbarname', tab.name)
      this.$router.push(tab.name)
    },
    changes(){
      
      console.log("test2")
      console.log("111: ",this.activeName)
      console.log("2222: ",this.activeName)
      this.$store.commit('UpdateENV', this.radio1)
      var data = {}
      localStorage.setItem('ENV', this.radio1)
      if (this.activeName === "GetHost"){
        data["env"] = this.radio1
        data["namespace"] = null
        data['url'] = '/api/nodes'
        this.$store.dispatch('Postdata',data)
      }else if (this.activeName === 'Uploadenv' ){
         this.show = false
      }else if(this.activeName === 'ClusterInfo'){
            this.$router.push(this.activeName)
      }else{
        data["env"] = this.radio1
        data["namespace"] = null
        data['url'] = '/api/namespace'
        this.$store.dispatch('Postdata',data)
      }
      this.$router.push(this.activeName)
      console.log("label: ",this.radio1)
    },
  

  }
}
</script>
