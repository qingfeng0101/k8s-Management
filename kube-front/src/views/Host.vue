<template>
  <div v-if="show">
   <el-table
    :data=Data
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
  <div v-else>
    请添加环境
  </div> 
</template>
<script>
import { mapState } from 'vuex'
export default {
  data() {
    return {
      //show: true,
      data: {}
    }
  },
  computed: {
    ...mapState(['Data','show','ENV'])
  },
  mounted () {
      
      // var data = {}
      // data['url'] = '/api/names'
      // this.$store.dispatch('Getenv',data)
      //  console.log("length1: ",this.Data)
      // if (this.Data.length === 0){
      //   console.log("length: ",this.Data)
      //   this.show = false
      // }
      // else{
        var show = localStorage.getItem('show')
        if (show == "true"){
          show = true
        }else{
          show = false
        }
        this.$store.commit('hildshow', show)
        this.$store.commit('UpdateENV', localStorage.getItem('env'))
        if (this.show){
           console.log("show2: ",this.show)
          this.data["env"] = this.ENV
          this.data["namespace"] = null
          this.data['url'] = '/api/nodes'
          this.data['name'] = 'null'
          this.$store.dispatch('Postdata',this.data)   
        }
        
      //}
      
  }
}
</script>
