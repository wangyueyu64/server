<template>
  <div class = "demo">
    <h1>设备信息</h1>
    <ul class="bottom">
      <li @click="showinfo()">show</li>
    </ul>
  </div>
</template>

<script>
  //获取用户登录设备
  import * as sysTool from "../utils/systemTool.js";
  export default {
    name: 'demo',
    data () {
      return {
          ip: '1.1.1.1',
          area: '北京市',
          brower: 'chrome',
          os: 'windows7',
          dev: '电脑',
          version: '1.1'
      }
    },
    methods:{
      showinfo() {
        this.ip = sessionStorage.getItem('ip')
        this.area = sessionStorage.getItem('area')
        this.brower = sysTool.GetCurrentBrowser()
        this.os = sysTool.GetOs()
        this.dev = sysTool.IsMobile()
        this.version = sysTool.GetVersion()
  
        console.log('ip，地区，浏览器，操作系统，版本号，设备，：',  this.ip, this.area,this.brower, this.os, this.version, this.dev)
      },
  
    }
  }
  
</script>

<style>
  .demo {
    float:left;
    border: solid 1px #000000;
    width:1000px;
    height: 800px;
    margin:auto;
    top: 0; left: 0; bottom: 0; right: 0;
    position: absolute;
  }
</style>
