<template>
  <div v-if="isshow">
   <el-table
    :data="Data"
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
    ...mapState(['Data','isshow','ENV'])
  },
  mounted () {
        var data = {}
      data['url'] = '/api/names'
     this.$store.dispatch('Getenv', data).then((res) => {
        if (localStorage.getItem('ENV') == null){
          
            console.log("22222")
           console.log("show2: ",this.ENV)
          this.data["env"] = this.ENV
          this.data["namespace"] = null
          this.data['url'] = '/api/nodes'
          this.data['name'] = 'null'
          console.log("data: ",this.data)
          this.$store.dispatch('Postdata',this.data)   
        }else{
          console.log("1")
           this.$store.commit('UpdateENV', localStorage.getItem('ENV'))
           console.log("isshow1: ",this.isshow)
            if (this.isshow == true){
           console.log("show2: ",this.ENV)
          this.data["env"] = this.ENV
          this.data["namespace"] = null
          this.data['url'] = '/api/nodes'
          this.data['name'] = 'null'
           console.log("show2: ",this.data)
          this.$store.dispatch('Postdata',this.data)   
        }
        }
     })
      
        
      //}
      
  },
  methods: {
       
      },
}
</script>
