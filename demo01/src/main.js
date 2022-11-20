import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import router from "./router"
import store from "./store"
import Cookie from 'js-cookie'
Vue.config.productionTip = false
// 全局引入
Vue.use(ElementUI)
// 守卫导航
router.beforeEach((to, from, next) => {
  const token = Cookie.get('token')
  const status = Cookie.get('status')
  // token不存在 说明没有登陆，页面应该跳转到登录页
  if (to.name === 'register' || to.name === 'home' || to.name === 'login') {
    next()
  } else {
    if (token) {
      // if(to.name === 'knowledge' || to.name === 'class') {
      //   if(status === 2) {
      //     next()
      //   } else {
      //     next({ name: 'home' })
      //   }
      // } else if(to.name === 'practice' || to.name === 'exam' || to.name === 'wrong') {
      //   if(status === 1) {
      //     next()
      //   } else {
      //     next({ name: 'home' })
      //   }
      // } else {
      //   next()
      // }
      next()
    } else {
      next({ name: 'login' })
    }
  }

  //else {
  //   if(!token && to.name != 'login') {
  //     next({name: 'login'})
  //   } else if (token && to.name === 'login') {
  //     next({name: 'home'})  // 这里本来我是想进入了主页要删除token才能进入Login页面
  //   } else {
  //     next()
  //   }
  // }
  // if (token && Cookie.get('status') === 1) {
  //   if (to.name === 'practice' || to.name === 'exam' || to.name === 'wrong') {
  //     next()
  //   } else {
  //     next({ name: 'home' })
  //   }
  // } else if (token && status === 2) {
  //   if (to.name === 'knowledge' || to.name === 'class') {
  //     next()
  //   } else {
  //     next({ name: 'home' })
  //   }
  // }
  // else {
  //   next({ name: 'login' })
  // }

})
new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')

