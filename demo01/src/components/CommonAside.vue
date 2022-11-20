<template>
    <el-menu default-active="1-4-1" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose"
        :collapse="isCollapse" background-color="#545c64" text-color="#fff" active-text-color="#ffd04b">
        <h3>{{ isCollapse ? 'OJ' : 'phycial OJ system' }}</h3>
        <el-menu-item @click="clickMenu(item)" v-for="item in noChildren" :key="item.name" :index="item.name">
            <i :class="`el-icon-${item.icon}`"></i>
            <span slot="title">{{ item.label }}</span>
        </el-menu-item>
        <!-- 下面为子标签 -->
        <!-- <el-submenu v-for="item in hasChildren" :key="item.label" :index="item.label">
            <template slot="title">
                <i :class="`el-icon-${item.icon}`"></i>
                <span slot="title">{{ item.label }}</span>
            </template>
            <el-menu-item-group v-for="subItem in item.children" :key="subItem.path">
                <el-menu-item @click="clickMenu(subItem)" :index="subItem.path">{{ subItem.label }}</el-menu-item>
            </el-menu-item-group>
        </el-submenu> -->
    </el-menu>
</template>
<style lang="less" scoped>
.el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
}

.el-menu {
    height: 100vh;
    border-right: none;

    h3 {
        color: #fff;
        text-align: center;
        line-height: 48px;
        font-size: 16px;
        font-weight: 400;
    }
}
</style>
  
<script>
import Cookie from 'js-cookie'
export default {
    data() {
        return {
            menuData: [
                {
                    path: '/',
                    name: 'home',
                    label: '首页',
                    icon: 's-home',
                    url: 'Home/Home'
                },
                {
                    path: '/practice',
                    name: 'practice',
                    label: '练习',
                    icon: 'pie-chart',
                    url: 'PracticeManage/PracticeManage'
                },
                {
                    path: '/exam',
                    name: 'exam',
                    label: '考试',
                    icon: 'document',
                    url: 'ExamManage/ExamManage'
                },
                {
                    path: '/wrong',
                    name: 'wrong',
                    label: '错题',
                    icon: 'circle-close',
                    url: 'WrongManage/WrongManage'
                },
                // {
                //     label: '其他',
                //     icon: 'location',
                //     children: [
                //         {
                //             path: '/page1',
                //             name: 'page1',
                //             label: '页面1',
                //             icon: 'setting',
                //             url: 'Other/PageOne'
                //         },
                //         {
                //             path: '/page2',
                //             name: 'page2',
                //             label: '页面2',
                //             icon: 'setting',
                //             url: 'Other/PageTwo'
                //         }
                //     ]
                // }
            ],
            menuDatas: [
                {
                    path: '/',
                    name: 'home',
                    label: '首页',
                    icon: 's-home',
                    url: 'Home/Home'
                },
                {
                    path: '/knowledge',
                    name: 'knowledge',
                    label: '题库',
                    icon: 'pie-chart',
                    url: 'KnowledgeManage/KnowledgeManage'
                },
                {
                    path: '/class',
                    name: 'class',
                    label: '班级',
                    icon: 'document',
                    url: 'ClassManage/ClassManage'
                }
            ]
        };
    },
    methods: {
        handleOpen(key, keyPath) {
            console.log(key, keyPath);
        },
        handleClose(key, keyPath) {
            console.log(key, keyPath);
        },
        clickMenu(item) {
            // console.log(item);
            // 跳转路由与当前路由不一致才允许跳转
            if (this.$route.path != item.path && !(this.$route.path === '/home' && (item.path === '/'))) {
                this.$router.push(item.path);
            }
        }
    },
    computed: {
        noChildren() {
            console.log("登陆是学生还是老师",Cookie.get('status'))
            if (Cookie.get('status') === '1') {
                console.log("学生的侧边栏",Cookie.get('status'))
                return this.menuData.filter(item => !item.children)
            } else if (Cookie.get('status') === '2') {
                console.log("老师的侧边栏",Cookie.get('status'))
                return this.menuDatas.filter(item => !item.children)
            } else {
                console.log("没有登陆的侧边栏",Cookie.get('status'))
                return this.menuData.filter(item => !item.children)
            }
        },
        // 有子菜单
        hasChildren() {
            return this.menuData.filter(item => item.children)
        },
        isCollapse() {
            return this.$store.state.tab.isCollapse
        }
    }
}
</script>