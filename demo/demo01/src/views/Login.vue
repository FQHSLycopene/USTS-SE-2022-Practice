<template>
    <div class="login_container">
        <div class="login">
            <el-card style="padding: 0;height: 290px;width: 350px;background-color: rgb(58, 98, 215);display: flex;">
                <div class="box">
                    <span style="color: #fff;font-size: 20px;margin-left: 10px;">Physical OJ System</span>
                    <img src="../assets/Login_left.png" style="height: 174px;width: 234px;" alt="">
                </div>
            </el-card>
            <el-card style="height: 290px;width: 350px;">
                <el-tabs :stretch="true">
                    <el-tab-pane label="密码登录">
                        <el-form ref="form" label-width="70px" style="height: 185px;" :inline="true"
                            class="el-form login-container el-form--inline" :rules="rules" :model="form">
                            <el-form-item prop="name">
                                <el-input class="username" v-model="form.name" placeholder="用户名:"></el-input>
                            </el-form-item>
                            <el-form-item prop="password">
                                <el-input type="password" show-password  v-model="form.password" placeholder="密码:" style="height: 30px;"></el-input>
                            </el-form-item>
                            <el-form-item>
                                <el-button class="submit_button" @click="submit" type="primary">登录
                                </el-button>
                            </el-form-item>
                        </el-form>
                        <router-link class="register" to="/register">立即注册</router-link>
                    </el-tab-pane>
                    <el-tab-pane label="邮箱登录">
                        <el-form :model="emailForm">
                            <el-form-item>
                                <el-input v-model="emailForm.email" placeholder="邮箱" prefix-icon="el-icon-message">
                                </el-input>
                            </el-form-item>
                            <!-- 邮箱验证码 -->
                            <el-form-item>
                                <el-input v-model="emailForm.code" placeholder="验证码" prefix-icon="el-icon-key">
                                    <template #append>
                                        <el-button size='mini' @click="SendCode" v-if="isSend">{{sendmsg}}</el-button>
                                        <el-button size='mini' v-if="!isSend">{{ sendmsg }}</el-button>
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-form-item>
                                <el-button class="submit_button" @click="VerifyEmailCode" type="primary">登录
                                </el-button>
                            </el-form-item>
                        </el-form>
                        <router-link class="register" to="/register">立即注册</router-link>
                    </el-tab-pane>
                </el-tabs>
            </el-card>
        </div>
    </div>
</template>

<script>
import { LoginData, SendCode, VerifyEmailCode } from '../api'
import Cookie from 'js-cookie'
export default {
    data() {
        return {
            sendmsg: '发送验证码',
            isSend: true,
            form: {
                name: '',
                password: ''
            },
            rules: {
                name: [
                    { required: true, trigger: 'blur', message: '请输入用户名' }
                ],
                password: [
                    { required: true, trigger: 'blur', message: '请输入密码' }
                ]
            },
            emailForm: {
                code: '',
                email: ''
            }
        }
    },
    methods: {
        submit() {
            this.$refs.form.validate((valid) => {
                LoginData(this.form).then(({ data }) => {
                    console.log(data)
                    if (data.code === 200) {
                        this.$message.success(data.msg)
                        Cookie.set('token', data.data);
                        this.$router.push('/home');
                    } else {
                        this.$message.error(data.msg)
                    }
                })
            })
        },
        SendCode() {
            SendCode({ email: this.emailForm.email }).then(({ data }) => {
                if (data.code === 200) {
                    console.log("邮箱发送验证码")
                    console.log(data)
                    this.$message.success("验证码发送成功")
                    this.isSend = false;
                    let timer = 60;
                    this.sendmsg = timer + "s";
                    this.timeFun = setInterval(() => {
                        --timer;
                        this.sendmsg = timer + "s";
                        if (timer === 0) {
                            this.isSend = true;
                            this.sendmsg = "重新发送";
                            clearInterval(this.timeFun);
                        }
                    }, 1000);
                } else {
                    this.$message.error("验证码发送失败")
                }
            })
        },
        VerifyEmailCode() {
            VerifyEmailCode(this.emailForm).then(({ data }) => {
                console.log("邮箱验证登录")
                if (data.code === 200) {
                    this.$message.success(data.msg)
                    Cookie.set('token', data.data);
                    this.$router.push('/home');
                } else {
                    this.$message.error(data.msg)
                }
            })
        }
    }

}
</script>

<style lang="less" scoped>
.login_container {
    // 这些能让背景图片好好显示
    background-image: url(../assets/Login_background.jpg);
    background-color: rgb(58, 98, 215);
    width: 100%;
    height: 100%;
    position: fixed;
    background-size: 100% 100%;

    .login {
        width: 700px;
        border: 1px solid #eaeaea;
        margin: 180px auto;
        display: flex;

        .box {
            width: 200px;
            height: 220px;
            margin-top: 40px;

        }

        .username {
            margin-top: 0px; // 这个是name 输入框距离上面的元素的外边距
        }

        .el-input {
            width: 281px;
            margin-left: 15px;
        }

        .btn {
            margin-left: 44px;

            .btn_email {
                margin-left: 106px;
            }
        }

        .submit_button {
            width: 282px;
            margin-left: 15px;
            margin-top: 0px;
            color: #3A62D7
        }

        .register {
            font-size: 10px;
            margin-left: 247px;
            margin-top: 0;
        }

    }
}
</style>