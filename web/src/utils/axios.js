import axios from 'axios'           //记得引入axios库   npm install --save axios
import {Message} from 'element-ui' //弹窗组件


const request = axios.create({
  baseURL: 'localhost:8080', //baseURL:process.env.VUE_APP_BASE_API,   //localhost:8080前端运行端口
  timeout:60000,

  // switch(procees.env.NODE_ENV){   //NODE_ENV依靠package.json中的指令修改
  //   case "production":
  //     //生产环境
  //     axios.default.baseURL = "http://localhost:8000";
  //     break;
  //   case "test":
  //     //测试环境
  //     axios.default.baseURL = "http://localhost:8000";
  //     break;
  //   default:
  //     //开发环境
  //     axios.default.baseURL = "http://localhost:8000";
  // }
  //package.json可作如下配置
  // "serve":"vue-cli-service serve",   //默认开发环境
  // "serve:test":"set NODE_ENV=test&&vue-cli-service serve",
  // "serve:test":"set NODE_ENV=production&&vue-cli-service serve",

})

//请求拦截器（预处理器）
request.interceptors.request.use(config=>{   //两个箭头函数作为参数 config=>{},error=>{}
  let user = localStorage.getItem('userInfo')
  //为请求头添加用户信息
  user && (config.headers['userInfo'] = user )
  return config
},error=>{
  return Promise.reject(error)
})


//响应拦截器（处理器）
request.interceptors.response.use(response=>{ //response=>{},error=>{}
  const res = response.data
  console.log(`"intercept ${response.config.method} response..."`,res)
  if(res.code === 40400)
    Message({message:res.message,type:'error'})
  else if(res.code === 50000)
    Message({message:"服务器出错，请刷新重新尝试",type:'error'})
  else if(res.code === 20000)
    Message({message:"拦截结果：请求成功"})
  return res
},error=>{
  console.log("intercepting...",error)
  Message({
    message:error.message,
    type:'error',
    duration:5 * 1000
  });
  return Promise.reject(error)
})

export default request
