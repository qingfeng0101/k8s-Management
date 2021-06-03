<template>
    <div>
     <el-alert
     v-show='show'
    title="已存在相同环境"
    type="error"
    description="请重新输入不同名称的环境，避免重复环境"
    @close="Close"
    show-icon>
  </el-alert>
    <div style="width: 180px;">
   <el-input v-model="input" placeholder="请输入环境名称" ><el-button type="primary">主要按钮</el-button></el-input>
    </div>
        <el-upload
  class="upload-demo"
  action="http://192.168.0.105:8080/api/upload"
  :on-remove="handleRemove"
  :before-remove="beforeRemove"
  :before-upload="beforeUpload"
  :http-request="submitUpload" 
  :file-list="Env">
  <el-button size="small" type="primary" >点击上传</el-button>
</el-upload>

    </div>
</template>
<style lang="scss" scoped>
.test{
    height: 40px;
    width: 180px;
}
</style>
<script>
import { mapState } from 'vuex'
  export default {
    data() {
      return {
        fileList: [],
        input: '',
        text: '',
        Title:'',
        show: false,
      };
    },
    computed: {
    ...mapState(['Env','isshow'])
  },
  beforeMount () {
    // bus.$emit('maizuo', false)
    // this.$store.state.isShowtabelbar = false
    this.$store.commit('hildShow', false)
  },
  beforeDestroy () {
    // this.$store.state.isShowtabelbar = true
    this.$store.commit('hildShow', true)
  },
    mounted() {
         var data = {}
         data['url'] = '/api/names'
         this.$store.dispatch('Getenv',data)
         console.log(this.Env)

    }, 
    methods: {
      handleRemove(file, fileList) {
          var data = {}
          data['url'] = '/api/delname'
          data['name'] = file.name
          this.$store.dispatch('Postdata', data)
           var a = this.Env.findIndex(item => item.name == file.name);
           this.Env.splice(a, 1);
           if (this.Env.length != 0){
               localStorage.setItem('env',this.Env[0].name); 
           }else if (this.Env.length === 0){
             console.log("11111")
             this.$store.commit('hildshow', false)
             localStorage.setItem('show',false);
           }
           
      },
      beforeRemove(file, fileList) {
        return this.$confirm(`确定移除 ${ file.name }？`);
      },
      beforeUpload(file){
          if (this.Env.length === 0){
              return true
          }else{
           for (var i=0;i<this.Env.length;i++)
        { 
            if (this.Env[i].name === this.input){
                this.show = true;
                return false;
            }
            return true
        }}
      },
      submitUpload(file){

          console.log("sub: ",file)
        //   const formData = new FormData()
        //   formData.append('file', file)
         if (this.Env.length === 0){
             var data = {}
                   data['url'] = '/api/upload'
                   data['data'] = file
                   data['name'] = this.input
                   this.$store.dispatch('Uploadenv', data)
                   this.$store.commit('UpdateENV', this.input)
                   localStorage.setItem('env',this.input)
                   localStorage.setItem('show',true);
         }else{
             console.log("2")
        for (var i=0;i<this.Env.length;i++)
        { 
            if (this.Env[i].name === this.input){
                console.log("1")
                this.show = true;
            }else{
                console.log("3")
                   var data = {}
                   data['url'] = '/api/upload'
                   data['data'] = file
                   data['name'] = this.input
                   this.$store.dispatch('Uploadenv', data)
                   this.$store.commit('UpdateENV', this.input)
                   localStorage.setItem('env',this.input)
                   localStorage.setItem('show',true);
            }
        }}

        //   var data = {}
        //   data['url'] = '/api/upload'
        //   data['data'] = file
        //   data['name'] = this.input
        //   this.$store.dispatch('Uploadenv', data)
        //   console.log(param)
      },
      Close(){
          this.show = false;
      }

      
    }
  }
</script>
