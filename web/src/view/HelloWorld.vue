<template>
  <div class="hello">
    <img src="../assets/599dd2c83233a_610.jpg">
    <h1>网络资产管理系统</h1>
    <el-row>
      <el-button type="primary" @click= "register">注册</el-button>
      <el-button type="success"  @click= "login">登录</el-button>
    </el-row>
    <router-view class="content"/>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
       }
  },
  methods: {
      register () {
        this.$router.push('/register')
      },
      login () {
        this.$router.push('/login')
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
img{
  width: 300px;
  height: 300px;
}
.el-button--primary{
  margin-top: 100px;
  min-height: 50px;
  min-width: 100px;
}
.el-button--success{
  margin-top: 100px;
  min-height: 50px;
  min-width: 100px;
}
.hello {
  float:left;
  border: solid 1px #000000;
  width:1000px;
  height: 800px;
  margin:auto;
  top: 0; left: 0; bottom: 0; right: 0;
  position: absolute;
}
.content{
  width: 100%;
  height: 100%;
  position: absolute;
}
</style>
