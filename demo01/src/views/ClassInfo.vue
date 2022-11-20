<template>
    <div class="classinfo">
        <el-container>
            <el-header style="height: 55px;">
                <i class="el-icon-back" style="margin-left: 63px;font-size: 42px;" @click="goback"></i>
                <span class="class">班级信息</span>
            </el-header>
            <el-container>
                <el-aside width="200px">
                    <h1>班级信息</h1>
                    <h2>{{ infoForm.name }}</h2>
                    <hr style="height:1px;border:none;border-top:1px solid #555555;" />
                    <h2>{{ infoForm.student_number + "人" }}</h2>
                    <hr style="height:1px;border:none;border-top:1px solid #555555;" />
                    <h2>{{ infoForm.join_code }}</h2>
                    <hr style="height:1px;border:none;border-top:1px solid #555555;" />
                    <el-button type="danger" size="small" class="blueBug" @click="EditHandle">
                        修改班级代码
                    </el-button>
                </el-aside>
                <el-main>
                    <div class="tableBar">
                        <el-button type="primary" size="small" class="blueBug" @click="">
                            查看考试
                        </el-button>
                        <el-input size="mini" v-model="input" style="width: 250px;padding-left: 0%;" clearable
                            @keyup.enter.native="handleQuery">
                            <i slot="suffix" class="el-input__icon el-icon-search"
                                style="cursor: pointer; margin-left: 53px;" @click="handleQuery"></i>
                        </el-input>
                    </div>
                    <el-table :data="tableData" stripe class="tableBox" :header-cell-style="{ 'text-align': 'center' }"
                        :cell-style="{ 'text-align': 'center' }">
                        <el-table-column prop="name" label="班级学生"></el-table-column>
                        <el-table-column label="操作" width="280" align="center">
                            <template slot-scope="scope">

                                <el-button type="danger" size="small" class="blueBug"
                                    @click="Teacher_Class_put(scope.row)">
                                    删除
                                </el-button>

                            </template>
                        </el-table-column>
                    </el-table>
                    <el-pagination class="pageList" :page-sizes="[5, 10, 20, 30]" :page-size="pageSize"
                        layout="total, sizes, prev, pager, next, jumper" :total="counts" :current-page.sync="page"
                        @size-change="handleSizeChange" @current-change="handleCurrentChange"></el-pagination>

                </el-main>
            </el-container>
            <el-dialog :title="classData.title" :visible.sync="classData.dialogVisible" width="30%"
                :before-close="handleClose">
                <el-form class="demo-form-inline" label-width="100px">
                    <el-form-item label="班级代码">
                        <el-input v-model="classData.join_code" placeholder="单行输入" maxlength="14" />
                    </el-form-item>
                </el-form>
                <span slot="footer" class="dialog-footer">
                    <el-button size="medium" @click="classData.dialogVisible = false">取 消</el-button>
                    <el-button type="primary" size="medium" class="continue" @click="Teacher_Class_put">
                        修改</el-button>
                </span>
            </el-dialog>
        </el-container>
    </div>
</template>

<script>
import { teacher_Class_identity, teacher_ClassStudents, teacher_Class_put } from '../api'
export default {
    data() {
        return {
            infoForm: {
                name: '',
                student_number: '',
                join_code: '',
                identity: '',
            },
            action: '',
            input: '',
            counts: 0,
            page: 1,
            pageSize: 5,
            tableData: [],
            identity: '',
            classData: {
                'title': '修改班级代码',
                'dialogVisible': false,
                'join_code': '',
            }
        }
    },
    created() {
        this.init()
    },
    methods: {
        init() {
            console.log("接收到", this.$route.query.identity)
            this.identity = this.$route.query.identity;
            this.Teacher_Class_identity();
            this.Teacher_ClassStudents();

        },
        Teacher_Class_identity() {
            teacher_Class_identity({ params: { identity: this.identity } }).then(({ data }) => {
                console.log(data.code)
                if (data.code === 200) {
                    this.infoForm = data.data
                    this.$message.success("获取班级详情失败成功")
                } else {
                    this.$message.error("获取班级详情失败")
                }
            });
        },
        Teacher_ClassStudents() {
            teacher_ClassStudents({ params: { page: this.page, pageSize: this.pageSize, classIdentity: this.identity } }).then(({ data }) => {
                console.log(data.code)
                if (data.code === 200) {
                    this.tableData = data.data.list
                    this.counts = data.data.total
                    this.$message.success("更新成功")
                } else {
                    this.$message.error("更新失败")
                }
            })
        },
        Teacher_Class_put(row) {
            if (this.action === 'edit') {
                this.action = '';
                teacher_Class_put({
                    "class_identity": this.infoForm.identity,
                    "is_change_code": true,
                    "student_identities": [
                    ]
                }).then(({ data }) => {
                    if (data.code === 200) {
                        this.classData.dialogVisible = false
                        this.Teacher_Class_identity()
                        this.$message.success("修改成功")
                    } else if (data.code === 403) {
                        this.classData.dialogVisible = false
                        this.Teacher_Class_identity()
                        this.$message.error("修改失败")
                    }
                })
            } else {
                teacher_Class_put({
                    "class_identity": this.infoForm.identity,
                    "class_name": this.infoForm.name,
                    "is_change_code": false,
                    "student_identities": [
                        row.identity
                    ]
                }).then(({ data }) => {
                    if (data.code === 200) {
                        this.classData.dialogVisible = false
                        this.Teacher_ClassStudents()
                        this.$message.success("删除成功")
                    } else {
                        this.$message.error("删除失败")
                    }
                })
            }

        }
        ,
        handleQuery() {

        },
        handleSizeChange(val) {
            this.pageSize = val
            this.init()
        },
        handleCurrentChange(val) {
            this.page = val
            this.init()
        },
        // 关闭弹窗
        handleClose(st) {
            this.classData.dialogVisible = false
        },
        goback() {
            this.$router.go(-1)
        },
        EditHandle() {
            this.action = 'edit';
            this.classData.join_code = this.infoForm.join_code;
            this.classData.dialogVisible = true
        },
    }
}
</script>
<style lang="less" scoped>
.classinfo {

    .el-container {
        .el-header {
            background-color: #B3C0D1;
            color: #333;

            .class {
                font-size: 30px;
                margin-left: 567px;
            }
        }

        .el-container {
            .el-main {
                background-color: #E9EEF3;
                color: #333;
                text-align: center;
                height: 586px;

                .tableBar {
                    width: 40%;
                    margin-left: 718px;
                    margin-bottom: 12px;

                    .el-input {
                        padding-left: 0px;
                        margin-left: 53px;
                    }
                }

                .el-pagination {
                    margin: 12px 0px 12px 700px;
                }
            }

            .el-aside {
                background-color: #E9EEF3;
                color: #333;
                text-align: center;
                border-right: 1px dashed;
            }
        }
    }
}
</style>