<template>
    <el-main>
        <div class="container">
            <div class="search">
                <el-input v-model="keyWord" placeholder="请输入内容"></el-input>
                <el-button @click="fp_btn" slot="append" icon="el-icon-search"></el-button>
            </div>
            <div>
                <el-table :data="pvData" border fit highlight-current-row style="width: 80%;margin: 25px auto;"
                    :header-cell-style="{ 'text-align': 'center' }" :cell-style="{ 'text-align': 'center' }">
                    <el-table-column prop="CreatedAt" label="日期"></el-table-column>
                    <el-table-column prop="name" label="错题名称"></el-table-column>
                </el-table>
            </div>
            <div>
                <el-pagination background @size-change="handleSizeChange" @current-change="handleCurrentChange"
                    :current-page="page" :page-sizes="[5, 10, 20, 40]" :page-size="pageSize"
                    layout="total, sizes, prev, pager, next, jumper" :total="total">
                </el-pagination>
            </div>
        </div>
    </el-main>
</template>
  
<script>
import { student_WrongProblem } from "../api";
export default {
    data() {
        return {
            // 分页
            page: 1, //初始页
            pageSize: 5,
            pvData: [],
            findName: '',
            keyWord: '',
            total: 0
            // form: {

            // }
        }
    },
    created() {
        this.student_WrongProblem()
    },
    methods: {
        student_WrongProblem() {
            student_WrongProblem({ params: { page: this.page, pageSize: this.pageSize } }).then(({ data }) => {
                console.log("更新的数据量：" + data.data.list.length)
                // console.log("更新" + data.data.list)
                console.log(data.code)
                if (data.code === 200) {
                    this.pvData = data.data.list
                    this.total = data.data.total
                    this.$message.success("更新成功")
                } else {
                    this.$message.error("更新失败")
                }
            })
        },
        // 初始页page、初始每页数据数pageSize和数据data
        handleSizeChange: function (size) {
            this.pageSize = size;
            console.log("下拉显示", this.pageSize); //每页下拉显示数据
            this.student_WrongProblem()
        },
        handleCurrentChange: function (page) {
            this.page = page;
            console.log("下拉显示的第几页", this.page); //点击第几页
            this.student_WrongProblem()
        },

        fp_btn() {
            student_WrongProblem({ params: { page: this.page, pageSize: this.pageSize,keyWord: this.keyWord} }).then(({ data }) => {
                console.log("更新" + data)
                if (data.code === 200) {
                    this.pvData = data.data.list
                    this.total = data.data.total
                    this.$message.success("搜索成功",data.data.total)
                } else {
                    this.$message.error("搜索失败")
                }
            })
        }
    }
};
</script>
  
<style lang="less" scoped>
.container {

    .search {
        display: flex;
        width: 35%;
        margin-left: 660px;
    }


    .el-pagination {
        margin: 12px 0px 12px 700px;
    }
}
</style>
  