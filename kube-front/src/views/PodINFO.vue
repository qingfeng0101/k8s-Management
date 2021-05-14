<template>
<div>
  <div class="head"> {{$store.state.Pod.Name}} 详情</div>
 <el-button-group>
  <span v-on:click="back"><el-button size="mini" type="primary" icon="el-icon-arrow-left">返回</el-button></span>
</el-button-group>

<table >
<span>
      <tr >
        <td >Name</td>
        <td >{{$store.state.Pod.Name}}</td>
      </tr>
      <tr >
        <td ><pre>Namespace    </pre></td>
        <td >{{$store.state.Pod.Namespace}}</td>
      </tr>
      <tr >
        <td >Labels</td>
        <td >
        <span v-for="(label,key) in $store.state.Pod.Labels" :key=key>
        {{key}} = {{label}}
        </span>
        </td>
      </tr>
    </span>
    <span   v-for="(Container,index) in $store.state.Pod.Containers" :key=index>
      <tr >
        <td >Container Name</td>
        <td >{{Container.Name}}</td>
      </tr>
      <tr >
        <td >Container ID</td>
        <td>{{Container.ContainerID}}</td>
      </tr>
      <tr >
        <td >Image</td>
        <td>{{Container.Image}}</td>
      </tr>
      <tr >
        <td >Image ID</td>
         <td>{{Container.ImageID}}</td>
      </tr>
       <tr >
        <td >State</td>
        <td>{{status(Container.Status)}}</td>
      </tr>
      <tr >
        <td >Ready</td>
        <td>{{Container.Ready}}</td>
      </tr>
      <tr >
        <td >Restart Count</td>
        <td>{{Container.RestartNum}}</td>
      </tr>
      <tr >
        <td >Limits</td>
        <td>CPU={{Container.Lcpu}} MEM={{Container.Lmem}}</td>
      </tr>
      <tr >
        <td >Requests</td>
        <td>CPU={{Container.Rcpu}} MEM={{Container.Rmem}}</td>
      </tr>
    </span>
    </table>
</div>
   

</template>
<script>
export default {
    beforeMount () {
    // bus.$emit('maizuo', false)
    // this.$store.state.isShowtabelbar = false
    this.$store.commit('hildShowtabbar', false)
  },
  beforeDestroy () {
    // this.$store.state.isShowtabelbar = true
    this.$store.commit('Showtabbar', true)
  },
    methods:{
        status(data){
            for (var k in data) {
               return k
            }

        }, 
        back(){
        this.$router.go(-1);//返回上一层
    },
    }
}
</script>
<style lang="scss" scoped>
span{
margin: 0;
    padding: 0;
    border: 0;
    font-size: 100%;
    font: inherit;
    vertical-align: baseline;
}
       table,table tr th, table tr td {

        border:1px solid #ccc;

         }

       table{

        width: 1100px;

        border-collapse: collapse;

    }  
.head{
height:60px;
width: 100%;
background-color: rgb(54, 46, 46);
text-align:center;
color: #fffefe;
}
    </style>