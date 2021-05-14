<template>
  <nav >

<el-menu
  :default-active="activeIndex2"
  class="el-menu-demo"
  mode="horizontal"
  @select="handleSelect"
  background-color="#545c64"
  text-color="#fff"
  active-text-color="#ffd04b">
  <el-menu-item index="1">处理中心</el-menu-item>
  <el-submenu index="2">
    <template slot="title"  >命名空间</template>
    <el-menu-item  :index=data.Name v-for="(data,index) in $store.state.Test" :key=index>
    {{data.Name}}
    </el-menu-item>

    <el-submenu index="2-4">
      <template slot="title">选项4</template>
      <el-menu-item index="2-4-1">选项1</el-menu-item>
      <el-menu-item index="2-4-2">选项2</el-menu-item>
      <el-menu-item index="2-4-3">选项3</el-menu-item>
    </el-submenu>
  </el-submenu>
  <el-menu-item index="3" disabled>消息中心</el-menu-item>
  <el-menu-item index="4"><a href="https://www.ele.me" target="_blank">订单管理</a></el-menu-item>
</el-menu>

  </nav>
</template>
<script>
export default {
  data () {
    return {
      activeIndex: '1',
      activeIndex2: '1'

    }
  },
  mounted () {
    this.$store.dispatch('Gettest')
  },
  methods: {
    handleSelect (key, keyPath) {
      this.$store.dispatch('Postpods', keyPath[1])
      this.$router.push(`/pods`)
      console.log(key, keyPath)
    }

  }
}
</script>
