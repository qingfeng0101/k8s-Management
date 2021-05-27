<template>
    <div>
    <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
    <el-tab-pane label="deployments列表" name="deployments">deployments列表</el-tab-pane>
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
      width="200">
    </el-table-column>
    <el-table-column
      prop="desired"
      label="DESIRED"
      width="120">
    </el-table-column>
    <el-table-column
      prop="current"
      label="CURRENT"
      width="120">
    </el-table-column>
    <el-table-column
      prop="uptodate"
      label="UP-TO-DATE"
      width="120">
    </el-table-column>
    <el-table-column
      prop="available"
      label="AVAILABLE"
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
      width="500">
      <template slot-scope="scope">
        <el-button
          type="text"
          size="small">
          <el-button type="text" @click="open(scope.$index, $store.state.Data)">删除</el-button>
        </el-button>

        <el-button
          @click.native.prevent="Add(scope.$index, $store.state.Data)"
          type="text"
          size="small">
          扩容

        </el-button>
        <el-input  maxlength="10" v-model=num placeholder="请输入内容" v-show=$store.state.Data[scope.$index].isshow></el-input>
        <el-button
          v-show=$store.state.Data[scope.$index].isshow
          @click.native.prevent="OK(scope.$index, $store.state.Data)"
          type="text"
          size="small">

          确认
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
      activeName: 'deployments',
      podname: 'test',
      data: {},
      isshow: false,
      num: 0,
      index: null
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
    // console.log("mounted data",$store.state.Pods)
    var ENV = localStorage.getItem('ENV')
    this.namespace = localStorage.getItem('namespace')
      this.data['url'] = '/api/getdeplyment'
       this.data['namespace'] = this.namespace
       this.data["env"] = localStorage.getItem('ENV')
       this.$store.dispatch('Postdata',this.data)
  },

  methods: {
    Add (index, rows) {
      // var namespace = rows[index].name
      // this.$store.dispatch('DeletePods', { name, namespace })
      if (this.index === null) {
        this.index = index
      } else {
        rows[this.index].isshow = false
        this.index = index
      }
      this.num = rows[index].desired
      rows[index].isshow = true
    },
    OK (index, rows) {
      console.log('OK: ', this.num)
      var data = {}
      data["env"] = localStorage.getItem('ENV')
      data['name'] = rows[index].name
      data['namespace'] = rows[index].namespace
      data['num'] = this.num
      data['url'] = '/api/updatadeployment'
      this.$store.dispatch('Postdata', data)
      rows[index].isshow = false
    },
    open (index, rows) {
      this.data['env'] = localStorage.getItem('ENV')
      this.data['url'] = '/api/deletedeployment'
      this.data['name'] = rows[index].name
      this.data['namespace'] = rows[index].namespace
      this.$confirm('是否删除' + rows[index].name, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$store.dispatch('Postdata', this.data)
        this.$message({
          type: 'success',
          message: '删除成功!'
        })
        this.data['url'] = '/api/getdeplyment'
        this.$store.dispatch('Postdata', this.data)
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    handleClick (tab, event) {
      localStorage.setItem('activeName', tab.name)
      this.$store.commit('UpdateTabbarname', tab.name)
      this.$router.push(tab.name)
    }
  }

}
</script>
