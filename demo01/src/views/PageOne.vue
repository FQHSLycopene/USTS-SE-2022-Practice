<template>
    <div class="container">
        <el-form>
            <el-button icon="el-icon-back" @click="back_btn" style="display: inline-block;"></el-button>
            <h1>我的班级</h1>
        </el-form>
        <el-card :inline = "true" label> //这边要加一个接口
            <el-upload class="avatar-uploader" action="这里填入后台的接口地址" :show-file-list="false" :on-success="handleAvatarSuccess" :on-remove="handleRemove" :before-upload="beforeAvatarUpload">
                <img v-if="imageUrl" :src="imageUrl" class="avatar">
                <span v-if="imageUrl" class="el-upload-action" @click.stop="handleRemove()">
                    <i class="el-icon-delete"></i>
                </span>
                <i v-else class="el-icon-upload2 avatar-uploader-icon" stop></i>
            </el-upload>
            <el-tag :hit="true">花生</el-tag>
        </el-card>
        <el-table :data="pvData.slice((currentPage - 1) * pagesize, currentPage * pagesize)" border fit highlight-current-row style="width: 100%">
            <el-table-column prop="datein" label="加入日期"></el-table-column>
            <el-table-column prop="classname" label="班级名称"></el-table-column>
            <el-table-column prop="teathername" label="老师"></el-table-column>
        </el-table>

        <el-pagination style="margin: 12px 480px" background @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="currentPage" :page-sizes="[5, 10, 20, 40]"
            :page-size="pagesize" layout="total, sizes, prev, pager, next, jumper" :total="pvData.length">
        </el-pagination>
        <el-icon-circle-plus></el-icon-circle-plus>
        <el-button icon="el-icon-circle-plus" @click="getintoclass" class = "into_btn">加入班级</el-button>
    </div>
</template>

<script>
export default {
    data() {
        return {
            // 分页
            currentPage: 1, //初始页
            pagesize: 5, //    每页的数据
            total: 0,
            found_contant: "",
            pvData: [],
            imageUrl : ""
        };
    },
    methods: {
        // 初始页currentPage、初始每页数据数pagesize和数据data
        handleSizeChange: function (size) {
            this.pagesize = size;
            console.log(this.pagesize); //每页下拉显示数据
        },
        handleCurrentChange: function (currentPage) {
            this.currentPage = currentPage;
            console.log(this.currentPage); //点击第几页
        },
        // 移除图片
        handleRemove() {
            this.imageUrl = ''
        },
        handleAvatarSuccess(res, file) {
            this.imageUrl = res.data.url
        },
        // 上传前格式和图片大小限制
        beforeAvatarUpload(file) {
            const type = file.type === 'image/jpeg' || 'image/jpg' || 'image/webp' || 'image/png'
            const isLt2M = file.size / 1024 / 1024 < 2
            if (!type) {
                this.$message.error('图片格式不正确!(只能包含jpg，png，webp，JPEG)')
            }
            if (!isLt2M) {
                this.$message.error('上传图片大小不能超过 2MB!')
            }
            return type && isLt2M
        },
        getintoclass() //进入班级的方法
        {

        },
        back_btn(){ //回退键

        }
    }
};
</script>

<style scoped lang='less'>
    .container {
        margin: 30px;
    }

    h1 {
        text-align: center;
        display: inline-block;
        width: 1300px;
    }
    .avatar-uploader {
        display: inline-block;
        width: 60PX;
        height: 60px;
        border-radius: 50%;
        cursor: pointer;
        position: relative;
        overflow: hidden;
        background: no-repeat;
        background-size: 100% 100%;
    }

    .avatar-uploader-icon {
        font-size: 0;
        color: #fff;
        width: 60px;
        height: 60px;
        line-height: 60px;
        text-align: center;
    }

    .avatar-uploader-icon:hover {
        font-size: 28px;
        background-color: rgba(0, 0, 0, .3);
    }

    .avatar {
        position: relative;
        width: 60px;
        height: 60px;
        display: block;
    }

    .el-upload-action {
        position: absolute;
        top: 0;
        left: 0;
        display: block;
        width: 100%;
        height: 100%;
        font-size: 0;
        color: #fff;
        text-align: center;
        line-height: 120px;

    }

    .el-upload-action:hover {
        font-size: 20px;
        background-color: #000;
        background-color: rgba(0, 0, 0, .3)
    }
    .into_btn{
        background-color:#fff;
        float:right;
    }
</style>
