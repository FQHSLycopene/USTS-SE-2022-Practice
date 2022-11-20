<template>
    <div class="header-container">
        <div class="l-content">
            <el-button @click="handleMenu" icon="el-icon-menu" size="mini"></el-button>
            <!-- 面包屑 -->

        </div>
        <div class="r-content">
            <el-dropdown v-if="isLogin" @command="handleClick">
                <span class="el-dropdown-link">
                    <img class="user" src="../assets/user.png" alt="">
                </span>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item>个人中心</el-dropdown-item>
                    <el-dropdown-item command="cancel">注销</el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>

            <el-row v-if="!isLogin">
                <el-button size='mini' @click="login">
                    登录
                </el-button>
                <el-button size='mini' @click="register">
                    注册
                </el-button>
            </el-row>

        </div>
    </div>
</template>

<script>
import Cookie from 'js-cookie'
export default {
    data() {
        return {}
    },
    methods: {
        handleMenu() {
            this.$store.commit('collapseMenu')
        },
        handleClick(command) {
            if (command === 'cancel') {
                Cookie.remove('token');
                Cookie.remove('status')
                this.$router.push('/login')
            }
        },
        login () {
            this.$router.push("/login")
        },
        register () {
            this.$router.push("/register")
        }
    },
    computed: {
        isLogin() {
            console.log(Cookie.get('token')+"用户已登录所以有了token")
            return Cookie.get('token') != undefined
        }
    }
}
</script>
<style lang="less" scoped>
.header-container {
    padding: 0 20px; // 使左右都有20px的边距
    background-color: #333;
    height: 60px;
    display: flex;
    justify-content: space-between;
    align-items: center;

    .text {
        color: #fff;
        font-size: 14px;
        margin-left: 10px;
    }

    .r-content {
        .user {
            width: 40px;
            height: 40px;
            border-radius: 50%; // 这个就是把图像给变圆
        }
    }
}
</style>