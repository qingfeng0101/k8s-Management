<template>
  <div class="font" v-if=$store.state.Pods>
   <table >
  <tr class="tabbar"  style="border-left: none;border-right: none;">
  <td >NAME</td>
   <td >NAMESPACE</td>
  <td >READY</td>
  <td >STATUS</td>
  <td >RESTARTS</td>
  <td >AGE</td>
  <td >操作</td>
  </tr>
  <tr class="font" v-for="(data,index) in $store.state.Pods" :key=index>
  <td >{{data.Name}}</td>
  <td >{{data.Namespace}}</td>
  <td >{{data.Ready}}/{{data.ContainerNum}}</td>
  <td >{{data.Status}}</td>
  <td >{{data.RestartNum}}</td>
  <td >{{data.Age}}</td>
  <td >

  <el-dropdown >
    <span class="el-dropdown-link">
    操作选项<i class="el-icon-arrow-down el-icon--right"></i>
  </span>
  <el-dropdown-menu slot="dropdown"  >
    <el-dropdown-item  @click.native="Delete(data.Name,data.Namespace)">删除</el-dropdown-item>
    <el-dropdown-item @click.native="Podinfo(data.Name,data.Namespace)">查看详情</el-dropdown-item>
    <el-dropdown-item @click.native="GetPodLog(data.Name,data.Namespace)">查询日志</el-dropdown-item>
  </el-dropdown-menu>
</el-dropdown>

  </td>
  </tr>
</table>

  </div>
</template>
<script>
export default {

  methods: {
    Delete (name, namespace) {
      this.$store.dispatch('DeletePods', { name, namespace })
      console.log(name, namespace)
    },
    Podinfo(name, namespace){
      this.$store.dispatch('GetPodinfo', { name, namespace })
      
      console.log(name, namespace)
      this.$router.push("/podinfo")
    },
    GetPodLog(name, namespace){
       localStorage.setItem("name",name)
       localStorage.setItem("namespace",namespace)
       this.$router.push("/logs")
    }
  },

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
