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
      fixed="right"
      label="操作"
      width="300">
      <template slot-scope="scope">
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
      </template>
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
    data () {
        return {
          //show: true,
        }
    },
    computed: {
    ...mapState(['Data','show','ENV'])
  },
  mounted () {
    var show = localStorage.getItem('show')
    if (show == "true"){
          show = true
        }else{
          show = false
        }
    this.$store.commit('hildshow', show)
    this.$store.commit('UpdateENV', localStorage.getItem('env'))
    if (this.show){
      var data = {}
      data["env"] = this.ENV
      data["namespace"] = null
      data['url'] = '/api/namespace'
      data['name'] = 'null'
      this.$store.dispatch('Postdata',data)
      }
    //   if (this.$store.state.ENV === "test"){
    //      this.$store.dispatch('Getdata',"namespace")
    //    }else{
    //        this.$store.dispatch('Getdata',"prod/namespace")
    //    }
    
  },
  methods: {
    deleteRow (index, rows) {
      rows.splice(index, 1)
    },
    getpods (index, rows) {
      var namespace = rows[index].name
      localStorage.setItem('namespace', namespace)
      var data = {}
       data['url'] = '/api/pod'
       data['namespace'] = namespace
       data["env"] = this.ENV
       this.$store.dispatch('Postdata',data)
      this.$router.push('/pods')
      },
    getdeployment (index, rows) {
        var namespace = rows[index].name
      localStorage.setItem('namespace', namespace)
      var data = {}
       data['url'] = '/api/getdeplyment'
       data['namespace'] = namespace
       data["env"] = this.ENV
       this.$store.dispatch('Postdata',data)
      this.$router.push('/deployment')
    }
  }

}
</script>
