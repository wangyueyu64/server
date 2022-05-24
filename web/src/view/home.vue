<template>
  <div class="about">
    <!-- 菜单 -->
    <div class="mymenu">
    <mymenu/>
    </div>

    <!-- 表格 -->
    <div class="table">
    <assetstable/>
    </div>



  </div>
</template>

<script>
//1.引入组件
import assetstable from '../components/assetstable.vue'
import mymenu from '../components/menu.vue'
export default {
    //2.映射组件标签
    components: {
      assetstable,
      mymenu
      }
};
</script>

<style lang="scss" scoped>
.table{
	width: 1500px;
	margin: auto;
  //margin-top: 100px;
  top: 0; left: 0; bottom: 0; right: 0;
}
.mymenu{
  width: 1500px;
  margin: auto;
  margin-top: 0px;
}
</style>
