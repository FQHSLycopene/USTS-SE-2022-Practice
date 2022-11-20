<template>
    <div class="home_container">
        <div class="home">
            <h1 class="title">欢迎来到物理OJ系统!</h1>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {}
    },

}
</script>

<style lang="less" scoped>
.home_container {
     // 这些能让背景图片好好显示
    background-image: url(../assets/home.jpg);
    background-color: rgb(58, 98, 215);
    width: 100%;
    height: 100%;
    position: fixed;
    background-size: 100% 100%;
    .home {
        .title {
            line-height: 437px;
            text-align: center;
            overflow: auto;
            font-size: 67px;
            margin-right: 90px;
        }
    }
}
</style>