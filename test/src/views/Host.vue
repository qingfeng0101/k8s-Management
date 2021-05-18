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
      <!-- <template slot-scope="scope">
        <el-button

          type="text"
          size="small">
          移除
        </el-button>
        <el-button
          @click.native.prevent="getdeployment(scope.$index, $store.state.Data)"
          type="text"
          size="small">
          查看deployment
        </el-button>
        <el-button
          @click.native.prevent="getpods(scope.$index, $store.state.Data)"
          type="text"
          size="small">
          查看pods
        </el-button>
      </template> -->
    </el-table-column>
  </el-table>
  </div>
</template>
<script>
export default {
  mounted () {
    this.$store.dispatch('GetHost')
  }
  // methods: {
  //   deleteRow(index, rows) {
  //     rows.splice(index, 1);
  //   },
  //   getpods(index,rows){
  //     var namespace = rows[index].name
  //     localStorage.setItem("namespace",namespace)
  //     // this.$store.dispatch('Postpods',namespace)
  //     this.$router.push("/pods")
  //   },
  //   getdeployment(index,rows){
  //     var namespace = rows[index].name
  //     localStorage.setItem("namespace",namespace)
  //     // this.$store.dispatch('Postpods',namespace)
  //     this.$router.push("/deployment")
  //   }

  // },

}
</script>
