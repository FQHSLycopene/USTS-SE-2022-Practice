<template>
    <div class="register-container">
        <!-- 注册内容 -->
        <div class="register">
            <el-card class="box-card">
                <h3>注册账号</h3>
                <router-link to="/login">
                    <span class="go">已有账号,立即登录</span>
                </router-link>
                <el-form ref="form" :model="form" label-width="80px" :rules="rules">
                    <el-form-item prop="name">
                        <el-input v-model="form.name" placeholder="设置用户名称" prefix-icon="el-icon-user-solid">
                        </el-input>
                    </el-form-item>
                    <el-form-item prop="password">
                        <el-input type="password" show-password v-model="form.password" placeholder="设置6-18位登录密码"
                            prefix-icon="el-icon-lock">
                        </el-input>
                    </el-form-item>
                    <el-form-item >
                        <el-input type="password" show-password v-model="password1" placeholder="再次输入登录密码"
                            prefix-icon="el-icon-lock"></el-input>
                    </el-form-item>
                    <el-form-item prop="phone">
                        <el-input v-model="form.phone" placeholder="输入手机号" prefix-icon="el-icon-phone"></el-input>
                    </el-form-item>
                    <el-form-item prop="email">
                        <el-input v-model="form.email" placeholder="输入邮箱" prefix-icon="el-icon-message"></el-input>
                    </el-form-item>
                    <el-form-item prop="radio">
                        <el-radio v-model="form.status" label="1">我是学生</el-radio>
                        <el-radio v-model="form.status" label="2">我是老师</el-radio>
                    </el-form-item>
                    <el-form-item>
                        <el-button @click="Register" type="primary">注册</el-button>
                    </el-form-item>
                </el-form>
            </el-card>

        </div>
    </div>
</template>
  
<script>
import { Register } from '../api'
export default {
    name: "Register",
    data() {
        return {
            form: {
                name: "",
                password: "",
                phone: "",
                email: "",
                status: "1",
            },
            password1: "",
            // 这个radio 就很妙 这个跟label对应 能够默认 
            rules: {
                name: [
                    { required: true, trigger: 'blur', message: '请输入用户名' }
                ],
                password: [
                    { required: true, trigger: 'blur', message: '请输入密码8-16位', min: 8, max: 16 }
                ],
                phone: [
                    { required: true, trigger: 'blur', message: '请输入手机号' }
                ],
                email: [
                    { required: true, trigger: 'blur', message: '请输入邮箱' }
                ]
            }
        };
    },
    methods: {
        Register() {
            this.$refs.form.validate((valid) => {
                if (this.form.password !== this.password1) {
                    this.$message.error("两个密码不一致")
                    console.log("两个密码不一致")
                    return
                }
                Register(this.form).then(({ data }) => {
                    console.log("注册")
                    if (data.code === 200) {
                        this.$message.success("注册成功")
                    } else {
                        this.$message.error("注册失败")
                    }
                    this.$router.push("/login")
                })
            })
        }

    },
};
</script>
  
<style lang="less" scoped>
.register-container {
    background-image: url(../assets/Login_background.jpg);
    background-color: rgb(58, 98, 215);
    width: 100%;
    height: 100%;
    position: fixed;
    background-size: 100% 100%;

    .register {
        width: 500px;
        height: 445px;
        border: 1px solid #dfdfdf;
        margin: 97px auto;

        .box-card {
            height: 445px;

            .go {
                Color: #1891FF;
                margin-left: 353px;
                font-size: 1px;
            }
        }

        .el-form-item {
            margin-bottom: 13px;
        }
    }
}
</style>
  