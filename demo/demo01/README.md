# demo01

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).


整个项目的搭建过程为: 
1. 首先准备好Node.js版本环境(v16.16.0) 这个为我的版本，当然你想要其他版本的话，可以通过nvm 工具快速切换环境，vue框架创建项目的时候会自动安装一些依赖存放在node_modules文件当中

2. 通过vue-cli架子 搭建好项目整体框架，本人使用vue2

3. 项目启动的过程:
    3.1 : main.js中通过创建vue实列，并且将App.vue组件挂载到index.html的一个class="app"的元素上面，通过在浏览器访问终端打印的地址，Vue实列会渲染在页面上面
    3.2 : App.vue 中引入了Home.vue 组件，并且将Home.vue组件嵌入到App.vue当中，一同显示，当然你可能会遇到报错，原因是组件的命名必须遵守驼峰命名的准则，解决办法在vue.config.js中加入lintOnSave: false // 关闭eslint校验

