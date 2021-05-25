<template>
  <div>
  <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
    <el-tab-pane label="pods列表" name="pods">pods列表</el-tab-pane>
    <el-tab-pane label="返回" name="Getnamespace">返回</el-tab-pane>
  </el-tabs>
  <el-table
    :data=$store.state.Data
    style="width: 100%"
    max-height="800">
    <el-table-column
      fixed
      prop="name"
      label="NAME"
      width="400">
    </el-table-column>
    <el-table-column
      prop="namespace"
      label="NAMESPACE"
      width="120">
    </el-table-column>
    <el-table-column
      prop="ready_str"
      label="READY"
      width="120">
    </el-table-column>
    <el-table-column
      prop="status"
      label="STATUS"
      width="200">
    </el-table-column>
    <el-table-column
      prop="restart_num"
      label="RESTARTS"
      width="120">
    </el-table-column>
    <el-table-column
      prop="age"
      label="AGE"
      width="300">
    </el-table-column>
    <el-table-column
      fixed="right"
      label="操作"
      width="300">
      <template slot-scope="scope">
        <el-button
        @click.native.prevent="Delete(scope.$index, $store.state.Data)"
          type="text"
          size="small">
          <el-button type="text" @click="open(scope.$index, $store.state.Data)">删除</el-button>
        </el-button>
        <el-button
          @click.native.prevent="Podinfo(scope.$index, $store.state.Data)"
          type="text"
          size="small">
          查看pod详情
        </el-button>
        <el-button
          @click.native.prevent="GetPodLog(scope.$index, $store.state.Data)"
          type="text"
          size="small">
          查看pod日志
        </el-button>
      </template>
    </el-table-column>
  </el-table>

  </div>
</template>
<script>
export default {
  data () {
    return {
      namespace: '',
      activeName: 'pods',
      podname: 'test',
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
    this.$store.commit('hildShowtabbar', false)
    var ENV = localStorage.getItem('ENV')
    var namespace = localStorage.getItem('namespace')
      var data = {}
      if (ENV === "test"){
          data['url'] = 'pod'
          data['namespace'] = namespace
           this.$store.dispatch('Postdata', data)
      }else{
           data['url'] = 'prod/pod'
          data['namespace'] = namespace
           this.$store.dispatch('Postdata', data)
      }
  },

  methods: {
    Delete(index, rows){
      console.log(index)
    },
    open (index, rows) {
      var ENV = localStorage.getItem('ENV')
      var namespace = localStorage.getItem('namespace')
      if (ENV === "test"){
          this.data['url'] = 'deletepod'
        }else{
          this.data['url'] = 'prod/deletepod'
        }
      this.podname = rows[index].name
      this.data['name'] = this.podname
      this.data['namespace'] = namespace
      this.$confirm('是否删除' + rows[index].name, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$store.dispatch('DeleteData', this.data)
        this.$message({
          type: 'success',
          message: '删除成功!'
        })
        if (ENV === "test"){
          this.data['url'] = 'pod'
      }else{
          this.data['url'] = 'prod/pod'
      }
        this.$store.dispatch('Postdata', this.data)
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    Podinfo (index, rows) {
      
      this.podname = rows[index].name
      this.data['name'] = this.podname
      this.data['namespace'] = rows[index].namespace
      localStorage.setItem('podname', this.podname)
      localStorage.setItem('namespace', rows[index].namespace)
      this.$router.push('/podinfo')
    },
    GetPodLog (index, rows) {
      localStorage.setItem('name', rows[index].name)
      localStorage.setItem('namespace', rows[index].namespace)
      this.$router.push('/logs')
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
.font{
font-size: 5px;
background-color:#666666;
}

.tabbar{
background-color:#888888;
}
.el-dropdown-link {
    cursor: pointer;
    color:#87CEFA;
  }
  .el-icon-arrow-down {
    font-size: 12px;
  }
</style>
