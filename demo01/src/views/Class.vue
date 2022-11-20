<template>
    <div class="container">
        <div class="tableBar">
            <el-input size="mini" v-model="input" style="width: 250px;padding-left: 0%;" clearable
                @keyup.enter.native="handleQuery">
                <i slot="suffix" class="el-input__icon el-icon-search" style="cursor: pointer" @click="handleQuery"></i>
            </el-input>
            <el-button size="small" type="primary" @click="addHandle">
                创建班级
            </el-button>
        </div>
        <el-table :data="tableData" stripe class="tableBox" :header-cell-style="{ 'text-align': 'center' }"
            :cell-style="{ 'text-align': 'center' }">
            <el-table-column prop="CreatedAt" label="创建日期"></el-table-column>
            <el-table-column prop="name" label="班级名称"></el-table-column>

            <el-table-column label="操作" width="280" align="center">
                <template slot-scope="scope">
                    <el-button size="small" class="blueBug" @click="EditHandle(scope.row)">
                        编辑
                    </el-button>
                    <el-button type="danger" size="small" class="blueBug" @click="DelHandle(scope.row)">
                        删除
                    </el-button>
                    <el-button type="primary" size="small" class="blueBug" @click="CatHandle(scope.row)">
                        查看
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-pagination class="pageList" :page-sizes="[2, 10, 20, 30]" :page-size="pageSize"
            layout="total, sizes, prev, pager, next, jumper" :total="counts" :current-page.sync="page"
            @size-change="handleSizeChange" @current-change="handleCurrentChange"></el-pagination>
        <el-dialog :title="classData.title" :visible.sync="classData.dialogVisible" width="30%"
            :before-close="handleClose">
            <el-form class="demo-form-inline" label-width="100px">
                <el-form-item label="创建班级：">
                    <el-input v-model="classData.name" placeholder="单行输入" maxlength="14" />
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button size="medium" @click="classData.dialogVisible = false">取 消</el-button>
                <el-button type="primary" size="medium" class="continue" @click="submitForm">
                    {{ action === 'add' ? '创建' : '编辑' }}</el-button>
            </span>
        </el-dialog>
    </div>
</template>
  
<script>
import { Class, teacher_Class,teacher_Class_put  } from '../api'
export default {
    data() {
        return {
            input: '',
            counts: 0,
            page: 1,
            pageSize: 10,
            tableData: [],
            id: '',
            action: '',
            classData: {
                'title': '编辑信息',
                'dialogVisible': false,
                'CreatedAt': '',
                'name': '',
                'identity': ''
            }
        }
    },
    created() {
        this.init()
    },
    methods: {
        submitForm() {
            if (this.action === 'add') {
                teacher_Class({ name: this.classData.name }).then(({ data }) => {
                    if (data.code === 200) {
                        this.classData.dialogVisible = false
                        this.init()
                        this.$message.success("添加成功")
                    }
                })
            } else if(this.action === 'edit') {
                teacher_Class_put({
                    "class_identity": this.classData.identity,
                    "class_name": this.classData.name,
                    "is_change_code": false,
                    "student_identities": [
                        
                    ]
                }).then(({ data }) => {
                    if (data.code === 200) {
                        this.classData.dialogVisible = false
                        this.init()
                        this.$message.success("编辑成功")
                    } else {
                        this.$message.error("编辑失败")
                    }
                })
            }
        },
        init() {
            Class({ params: { page: this.page, pageSize: this.pageSize } }).then(({ data }) => {
                console.log("更新的数据量：" + data.data.list.length)
                // console.log("更新" + data.data.list)
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
        handleQuery() {

        },
        addHandle() {
            this.classData.title = '创建班级'
            this.action = 'add'
            this.classData.dialogVisible = true
        },
        EditHandle(row) {
            this.classData.title = '编辑信息'
            this.action = 'edit'
            this.classData.name = row.name
            this.classData.CreatedAt = row.CreatedAt
            this.classData.identity = row.identity
            this.classData.dialogVisible = true
        },
        DelHandle(row) {
            console.log(id);
            this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
                'confirmButtonText': '确定',
                'cancelButtonText': '取消',
                'type': 'warning'
            }).then(() => {
                this.$message.success("删除成功")
            })
        },
        CatHandle(row) {
            this.$router.push({name: "classinfo", query: {identity: row.identity}})
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
    },
}
</script>

<style lang="less" scoped>
.tableBar {
    width: 40%;
    margin-left: 818px;
    .el-pagination {
        margin: 12px 0px 12px 700px;
    }
}
</style>