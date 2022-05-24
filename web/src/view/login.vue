<template>
	<p class="login">
				<el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
					<el-form-item label="账户" prop="name"><el-input v-model="ruleForm.name" placeholder="请输入账号"></el-input></el-form-item>
					<el-form-item label="密码" prop="pass"><el-input type="password" v-model="ruleForm.pass"placeholder="请输入密码" auto-complete="off"></el-input></el-form-item>
          <el-form-item label="验证码" prop="code"><el-input v-model="ruleForm.code" placeholder="请输入验证码"></el-input></el-form-item>
					<el-form-item>
						<el-button type="primary" @click="submitForm('ruleForm')">登录</el-button>
            <el-button type="primary" @click="sendcode" :disabled="this.isUsed">发送验证码</el-button>
						<el-button @click="resetForm('ruleForm')">重置</el-button>
					</el-form-item>
				</el-form>
	</p>
</template>

<script>
import api from '../api/index.js'
export default {
	data() {
		var validatePass = (rule, value, callback) => {
			if (value === '') {
				callback(new Error('请输入密码'));
			} else {
				if (this.ruleForm.checkPass !== '') {
					this.$refs.ruleForm.validateField('checkPass');
				}

				callback();
			}
		};

		return {
			activeName: 'first',
			ruleForm: {
				name: '',
				pass: '',
        code: '',
				checkPass: ''
			},
      isUsed : false,
			rules: {
				name: [{ required: true, message: '请输入您的名称', trigger: 'blur' }, { min: 1, max: 15, message: '长度在 1 到 15 个字符', trigger: 'blur' }],
				pass: [{ required: true, message: '请输入您的密码', validator: validatePass, trigger: 'blur' }]
			}
		};
	},

	methods: {
		//重置表单
		resetForm(formName) {
			this.$refs[formName].resetFields();
		},
		//提交表单
		submitForm(formName) {
			this.$refs[formName].validate(valid => {
				if (valid) {
					this.login();
				} else {

				}
			});
		},
    //发送登录请求
    async login (){
      let res = await api.login({'account': this.ruleForm.name, 'password':this.ruleForm.pass, 'code': this.ruleForm.code});
      //console.log(res.data);
      if(res.data === '登陆成功')
        this.$router.replace('/home');
      else if (res.data === '验证码错误！'){
          this.$message({message: res.data,type:'error'});
          this.ruleForm.code = '';
       }
      else{
        this.ruleForm.name = '';
        this.ruleForm.pass = '';
        this.ruleForm.code = '';
        this.$message({message: res.data,type:'error'});
      }
    },
    //发送验证码
    async sendcode(){
      this.isUsed = true;
      setInterval(()=>{
        this.isUsed= false;
      }, 5000)
      let res = await api.sendCode();
       //console.log(res);
    }

	},
};
</script>

<style lang="scss" scoped>
.login {
	width: 400px;
	margin: auto;

  margin-top: 300px;
  top: 0; left: 0; bottom: 0; right: 0;
}

.el-tabsitem {
	text-align: center;
	width: 60px;
}
</style>
