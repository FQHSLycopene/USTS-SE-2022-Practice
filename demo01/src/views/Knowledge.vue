<template>
    <div class="btns">
        <div class="btn" v-for="(item,index) in Knowledges" :key="item.name" :body-style="{ display: 'flex', padding: 0 }">
            <el-button :type="index % 2 === 0 ? 'primary' : 'warning'">{{ item.name }}</el-button>
        </div>
        <!-- <el-button type="primary" v-for="item in Knowledges" :key="item.name" >{{ item.name }}</el-button> -->
        <!-- <el-table :data="Knowledges" style="width: 100%">
            <el-table-column prop="name" width="180">
            </el-table-column>
        </el-table> -->
    </div>

</template>

<script>
import { teacher_Knowledge } from '../api'
export default {
    data() {
        return {
            Knowledges: [],
        }
    },
    mounted() {
        this.Teacher_Knowledge()
    },
    methods: {
        Teacher_Knowledge() {
            teacher_Knowledge().then(({ data }) => {
                console.log("更新的数据量：" + data.data.list.length)
                // console.log("更新" + data.data.list)
                console.log(data.code)
                if (data.code === 200) {
                    this.Knowledges = data.data.list
                    this.$message.success("更新成功")
                } else {
                    this.$message.error("更新失败")
                }
            })
        }
    }
}
</script>
<style lang="less" scoped>
.btns {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    // .el-button {
    //     width: 33%;
    //     margin: 77px auto;
    // }
    .btn {
        width: 32%;
        margin: 55px auto;
        border-bottom: solid;
        .el-button {
            margin: auto 130px;
            padding: 20px 50px;
        }
    }
}
</style>