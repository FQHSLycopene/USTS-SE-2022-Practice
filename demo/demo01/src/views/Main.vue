<template>
    <div>
        <el-container>
            <!-- 自适应 -->
            <el-aside width="auto">  
                <common-aside/>
            </el-aside>
            <el-container>
                <el-header>
                    <common-header/>
                </el-header>
                <el-main>
                    <router-view></router-view>
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>
<script>
import CommonAside from '@/components/CommonAside.vue';
import CommonHeader from '@/components/CommonHeader.vue';
export default {
    data() {
        return {}
    },
    components: {
        CommonAside,
        CommonHeader
    }
}
</script>
<style lang="less" scoped>
.el-header {
    padding: 0;
}
.el-main {
    padding: 0;
}
</style>