<template>
    <div class="Exam">
        <div v-if="isStart" class="start" align="middle">
            <span class="start_title">提示：请先加入班级再来考试!</span>
            <div class="btns_exam">
                <el-button @click="back" type="primary" size="medium">返回</el-button>
                <el-button @click="join" class="join" type="danger" size="medium">加入班级</el-button>
            </div>
        </div>
        <div v-if="isStartJoin" style="height: 254px;" class="start_join" align="middle">
            <span class="start_join_title">加入班级</span>
            <el-form label-width="80px">
                <el-form-item label="班级码：">
                    <el-input class="start_join_input" v-model="joinCode" placeholder="请输入班级码"></el-input>
                </el-form-item>
                <el-button @click="Student_Class" type="primary" size="medium" style="padding: 10px 40px;">保存
                </el-button>
            </el-form>
        </div>
        <div v-if="isJoin" class="container">
            <div class="tableBar">
                <el-input type="mini" v-model="input" style="width: 250px;padding-left: 0%;" clearable
                    @keyup.enter.native="handleQuery">
                    <i slot="suffix" class="el-input__icon el-icon-search" style="cursor: pointer"
                        @click="handleQuery"></i>
                </el-input>
            </div>
            <el-table :data="tableData" stripe class="tableBox" :header-cell-style="{ 'text-align': 'center' }"
                :cell-style="{ 'text-align': 'center' }">
                <el-table-column prop="CreatedAt" label="日期"></el-table-column>
                <el-table-column prop="name" label="分值"></el-table-column>
                <el-table-column prop="name" label="分数"></el-table-column>
                <el-table-column prop="name" label="考试"></el-table-column>
                <el-table-column prop="name" label="考试状态"></el-table-column>
                <el-table-column label="操作" width="280" align="center">
                    <template slot-scope="scope">
                        <el-link size="small" class="blueBug" @click="EditHandle(scope.row)">
                            编辑
                        </el-link>
                    </template>
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>

<script>
import { student_Exam, student_Class } from "../api";
export default {
    data() {
        return {
            input: '',
            page: 1, //初始页
            pageSize: 5,// 分页
            tableData: [],
            total: 0,
            isStart: false,
            isStartJoin: false,
            isJoin: false,
            joinCode: ''
        }
    },
    created() {
        this.Student_Exam()
    },
    methods: {
        Student_Exam() {
            student_Exam(({ params: { page: this.page, pageSize: this.pageSize } })).then(({ data }) => {
                console.log("更新的数据量：" + data)
                // console.log("更新" + data.data.list)
                console.log(data.code)
                if (data.code === 200) {
                    if (data.msg === "该学生尚未加入班级") {
                        this.isStart = true
                        this.$message.success("需要先加入班级哦！")
                    } else if (data.msg === "success") {
                        this.isStart = false
                        this.isStartJoin = false
                        this.isJoin = true
                        this.$message.success("已经加入过了")
                    }
                } else if(data.code === 402) {
                    this.isStart = true
                    this.$message.error("还没有加入班级")
                }
            })
        },
        back() {
            this.$router.push("/home")
        },
        join() {
            this.isStart = false;
            this.isStartJoin = true;
        },
        handleQuery() {

        },
        Student_Class() {
            student_Class({ joinCode: this.joinCode }).then(({ data }) => {
                console.log("更新的数据量：" + data)
                // console.log("更新" + data.data.list)
                console.log(data.code)
                if (data.code === 200) {
                    this.isStartJoin = false
                    this.isJoin = true
                    this.$message.success("加入班级成功")
                } else if (data.code === 401) {
                    this.$message.success("已经加入过了")
                }
                else {
                    this.$message.error("更新失败")
                }
            })
        }
    }
}
</script>

<style lang="less" scoped>
.Exam {
    .start {
        width: 480px;
        margin: 252px auto;

        .start_title {
            font-size: 30px;

        }

        .btns_exam {
            .join {
                margin-left: 85px;
            }
        }
    }

    .start_join {
        width: 480px;
        height: 500px;
        margin: 160px auto;

        .start_join_title {
            font-size: 30px;
        }

        .el-form-item {
            margin-top: 50px;
        }
    }

    .container {
        margin: auto 30px;
        .tableBar {
            width: 40%;
            margin-left: 818px;
            .el-input {
                padding-left: 0%;
            }
        }
    }
}
</style>