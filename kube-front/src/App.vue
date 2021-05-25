<template>
  <div>
    <Tabbar v-show="isShowtabelbar" ></Tabbar>
    <router-view></router-view>
       
   
     
   
  </div>
</template>
<script>
import Tabbar from '@/components/Tabbar'
import { mapState } from 'vuex'
export default {
  data () {
    return {
      isshow: true
    }
  },
  computed: {
    ...mapState(['isShowtabelbar'])
  },
  components: {
    Tabbar: Tabbar

  }
}

</script>
 <style lang="scss">
* {
margin: 0px;
padding: 0px;
}
html {
height: 100px;
}
li {
list-style: none;
}

</style>
