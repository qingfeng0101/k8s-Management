<template>
  <nav >

<el-tabs   v-model="activeName" type="card" @tab-click="handleClick">
    <el-tab-pane label="主机信息" name='GetHost'>主机信息</el-tab-pane>
    <el-tab-pane label="命名空间" name='Getnamespace'></el-tab-pane>
    <el-tab-pane label="环境增加" name='Uploadenv'></el-tab-pane>
  </el-tabs>
     <div v-show='isshow'>
     <el-radio  v-for="(v,index) in Env" :key="index" v-model="radio1" :label="v.name"   @change="changes" border>{{v.name}}</el-radio>
     
     </div> 
    

  </nav>
</template>
<script>
import { mapState } from 'vuex'
export default {
  data () {
    return {
      activeName: this.Tabbarname,
       radio1: this.ENV,
    }
  },
  computed: {
    ...mapState(['Env','isshow','ENV','Tabbarname','show'])
  },
 mounted () {
   this.$store.commit('UpdateTabbarname', localStorage.getItem('activeName'))
   this.activeName=this.Tabbarname
   this.$store.commit('UpdateENV', localStorage.getItem('env'))
   this.radio1 = this.ENV
   console.log(this.radio1)
  //  this.radio1 = this.$store.state.Env[0].name
  //  if (this.$store.state.Tabbarname === 'Uploadenv'){
  //         this.show=false
  //     }else{
  //        this.show=true
  //     }
      var data = {}
      data['url'] = '/api/names'
      this.$store.dispatch('Getenv',data)
      this.$store.commit('UpdateENV', this.radio1)
 },
   beforeDestroy () {
  // this.$store.state.isShowtabelbar = true
    this.$store.commit('Showtabbar', true)
  },
  methods: {
    handleClick (tab, event) {
      // this.activeName = tab.name
      this.$store.commit('UpdateENV', localStorage.getItem('env'))
      localStorage.setItem('activeName', tab.name)
      this.radio1 = this.ENV
      this.$store.commit('UpdateTabbarname', tab.name)
      this.$router.push(tab.name)
    },
    changes(){
      this.$store.commit('UpdateENV', this.radio1)
      var data = {}
      localStorage.setItem('ENV', this.radio1)
      if (this.activeName === "GetHost" &&  this.show === true){
        data["env"] = this.radio1
        data["namespace"] = null
        data['url'] = '/api/nodes'
        this.$store.dispatch('Postdata',data)
      }else if (this.activeName === 'Uploadenv' ){
         this.show = false
      }else{
        data["env"] = this.radio1
        data["namespace"] = null
        data['url'] = '/api/namespace'
        this.$store.dispatch('Postdata',data)
      }
      this.$router.push(this.activeName)
      console.log("label: ",this.radio1)
    }

  }
}
</script>
