<template>
  <nav>

<el-tabs v-model="activeName" type="card" @tab-click="handleClick">
    <el-tab-pane label="主机信息" name="GetHost">主机信息</el-tab-pane>
    <el-tab-pane label="命名空间" name="Getnamespace"></el-tab-pane>
  </el-tabs>
      
    <el-radio v-model="radio1" label="prod" @change="changes" border>生产环境</el-radio>
    <el-radio v-model="radio1" label="test" @change="changes" border>测试环境</el-radio>

  </nav>
</template>
<script>

export default {
  data () {
    return {
      activeName: this.$store.state.Tabbarname,
       radio1: 'test',
    }
  },
 mounted () {
   this.$store.commit('UpdateENV', this.radio1)
 },
  
  methods: {
    handleClick (tab, event) {
      console.log(tab.index)
      // this.activeName = tab.name
      localStorage.setItem('activeName', tab.name)
      this.$store.commit('UpdateTabbarname', tab.name)
      this.activeName = this.$store.state.Tabbarname
      this.$router.push(tab.name)

      console.log(tab, event)
    },
    changes(){
      this.$store.commit('UpdateENV', this.radio1)
      localStorage.setItem('ENV', this.radio1)
      if (this.activeName === "GetHost" ){
         if (this.radio1 === "test"  ){
             this.$store.dispatch('Getdata',"nodes")
        }else{
            this.$store.dispatch('Getdata',"prod/nodes")
        }
      }else{
        if (this.radio1 === "test"  ){
             this.$store.dispatch('Getdata',"namespace")
        }else{
            this.$store.dispatch('Getdata',"prod/namespace")
        }
      }
      
      this.$router.push(this.activeName)
      console.log("label: ",this.radio1)
    }

  }
}
</script>
