<template>
  <nav >

<el-tabs v-model="activeName" type="card" @tab-click="handleClick">
    <el-tab-pane label="主机信息" name="host">主机信息</el-tab-pane>
    <el-tab-pane label="命名空间" name="namespace"></el-tab-pane>
  </el-tabs>

  </nav>
</template>
<script>
export default {
  data () {
    return {
      activeName: 'host'

    }
  },
  mounted () {
    this.activeName = localStorage.getItem('activeName')
  },

  methods: {
    handleClick (tab, event) {
      console.log(tab.index)
      console.log()
      localStorage.setItem('activeName', tab.name)
      this.$router.push(tab.name)

      console.log(tab, event)
    }

  }
}
</script>
