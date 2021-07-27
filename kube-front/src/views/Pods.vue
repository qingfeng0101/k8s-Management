<template>
  <div>
  <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
    <el-tab-pane label="pods列表" name="pods">pods列表</el-tab-pane>
    <el-tab-pane label="返回" name="Getnamespace">返回</el-tab-pane>
  </el-tabs>
  <el-table
    :data=Data
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
        @click.native.prevent="Delete(scope.$index, Data)"
          type="text"
          size="small">
          <el-button type="text" @click="open(scope.$index, Data)">删除</el-button>
        </el-button>
        <el-button
          @click.native.prevent="Podinfo(scope.$index, Data)"
          type="text"
          size="small">
          查看pod详情
        </el-button>
        <el-button type="text" @click.native.prevent="logs(scope.$index, Data)">查看pod日志</el-button>
        
       <el-dialog
          :modal="false"
          title="提示"
          :visible.sync="centerDialogVisible"
           width="30%"
           center>
           <div v-for="(item,index) in Pod.containers" :key=index>
             <el-radio    v-model=radio :label=index>{{item.Name}}</el-radio>
           </div>
           <span slot="footer" class="dialog-footer">
         <el-button @click="centerDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click.native.prevent="confirm">确 定</el-button>
       </span>
      </el-dialog>
      <el-button type="text" @click.native.prevent="exec(scope.$index, Data)">进入pod终端</el-button>
        
       <el-dialog
          :modal="false"
          title="提示"
          :visible.sync="centerDialogVisible"
           width="30%"
           center>
           <div v-for="(item,index) in Pod.containers" :key=index>
             <el-radio    v-model=radio :label=index>{{item.Name}}</el-radio>
           </div>
           <span slot="footer" class="dialog-footer">
         <el-button @click="centerDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click.native.prevent="execs">确 定</el-button>
       </span>
      </el-dialog>
      </template>
    </el-table-column>
  </el-table>
  </div>
</template>
<script>
import { mapState } from 'vuex'
export default {
  data () {
    return {
      namespace: '',
      activeName: 'pods',
      podname: 'test',
      pods: [],
      data: {},
      radio: 0,
      centerDialogVisible: false,
      modal:false,
      names:{}
      
    }
  },
  computed: {
    ...mapState(['Data','Pod'])
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
    var namespace = localStorage.getItem('namespace')
    var data = {}
    data['env']= localStorage.getItem('ENV')
    data['url'] = '/api/pod'
    data['namespace'] = namespace
    this.$store.dispatch('Postdata', data)
      
  },

  methods: {
    Delete(index, rows){
      console.log(index)
    },
    open (index, rows) {
      this.data['env'] = localStorage.getItem('ENV')
      this.data['url'] = '/api/deletepod'
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
        this.data['url'] = '/api/pod'
        this.$store.dispatch('Postdata', this.data)
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    Podinfo (index, rows) {
      console.log("rows: ",rows)
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
    },
    test(){
      return this.Pod.containers
    },
    logs (index, rows) {
      this.names = rows[index]
      this.centerDialogVisible =true
      var radio1 = "1"
      this.data['env'] = localStorage.getItem('ENV')
      this.data['url'] = '/api/getpodinfo'
      this.data['name'] = rows[index].name
      this.data['namespace'] = rows[index].namespace
      this.$store.dispatch('GetPodinfo', this.data)
      localStorage.setItem('name',rows[index].name)
      console.log("openpodlogs",this.Pod)
    },
    confirm(){
       this.centerDialogVisible =false
       localStorage.setItem('cname',this.Pod.containers[this.radio].Name)
       this.$router.push('/logs')
    },
    exec (index, rows) {
      this.names = rows[index]
      this.centerDialogVisible =true
      var radio1 = "1"
      this.data['env'] = localStorage.getItem('ENV')
      this.data['url'] = '/api/getpodinfo'
      this.data['name'] = rows[index].name
      this.data['namespace'] = rows[index].namespace
      this.$store.dispatch('GetPodinfo', this.data)
      localStorage.setItem('name',rows[index].name)
      console.log("openpodlogs",this.Pod)
    },
    execs(){
      this.centerDialogVisible =false
       localStorage.setItem('cname',this.Pod.containers[this.radio].Name)
       this.$router.push('/podexec')
    }
}}
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