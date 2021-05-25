<template>
  <div >
   <el-table
    :data=$store.state.Data
    style="width: 100%"
    max-height="800px">
    <el-table-column
      fixed
      prop="name"
      label="NAME"
      width="300">
    </el-table-column>
    <el-table-column
      prop="status"
      label="STATUS"
      width="300">
    </el-table-column>
    <el-table-column
      prop="age"
      label="AGE"
      width="300">
    </el-table-column>
    <el-table-column
      prop="version"
      label="VERSION"
      width="300">
    </el-table-column>
    <el-table-column
      fixed="right"
      label="操作"
      width="300">
    </el-table-column>
  </el-table>
  </div>
  <!-- <div v-else>
     请选择环境
  </div> -->
</template>
<script>
export default {
  data() {
    return {
      show: false,
      data: {}
    }
  },
  
  mounted () {
       if (this.$store.state.ENV === "test"){
         this.$store.dispatch('Getdata',"nodes")
       }else{
         this.$store.dispatch('Getdata',"prod/nodes")
       }
       this.data = localStorage.getItem('envinfo')
       console.log("ENV: ",this.$store.state.ENV)
       console.log("host: ",this.$store.state.Tabbarname)
       
    
    
  }
}
</script>
