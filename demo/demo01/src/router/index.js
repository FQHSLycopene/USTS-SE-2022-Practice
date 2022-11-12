import VueRouter from "vue-router";
import Vue from 'vue';
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Main from '../views/Main.vue'
import Home from '../views/Home.vue'
import Practice from '../views/Practice.vue'
import Exam from '../views/Exam.vue'
import Wrong from '../views/Wrong.vue'
// import PageOne from '../views/PageOne.vue'
// import PageTwo from '../views/PageTwo.vue'
Vue.use(VueRouter)

const routes = [
    {
        path: "/", 
        component: Main, 
        redirect: '/home', // 重定向
        children: [
            { path: 'home', name: 'home', component: Home }, // 首页
            { path: 'practice', name: 'practice', component: Practice }, // 练习管理
            { path: 'exam', name: 'exam', component: Exam }, // 考试管理
            { path: 'wrong', name: 'wrong', component: Wrong }, // 考试管理
            // { path: 'page1', name: 'page1', component: PageOne }, // 页面1
            // { path: 'page2', name: 'page2', component: PageTwo }, // 页面2
        ]
    },
    {
        path: '/login',
        name: 'login',
        component: Login,
    },
    {
        path: '/register',
        name: 'register',
        component: Register,
    }
]
const router = new VueRouter({
    routes
})

export default router