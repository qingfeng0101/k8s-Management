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
      this.data["env"] = localStorage.getItem('ENV')
      this.data["namespace"] = null
      this.data['url'] = '/api/nodes'
      this.data['name'] = 'null'
      this.$store.dispatch('Postdata',this.data)
  }
}
</script>
