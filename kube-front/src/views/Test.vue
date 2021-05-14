<template>
  <div >
    <table>
      <tr v-for="(data,index) in $store.state.Test" :key=index>
        <td class="font"><pre>{{data}}</pre></td>
      </tr>
   </table>

  </div>
</template>
<script>
export default {
  mounted () {
    this.$store.dispatch('Gettest')
  }
}
</script>

<style lang="scss" scoped>
.font{
font-size: 12px;
}
</style>
