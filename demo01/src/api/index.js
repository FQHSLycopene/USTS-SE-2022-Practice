import http from '../utils/request'
export const LoginData = (form) => {
    console.log("登陆",form.name)
    return http.post('/Login',form);
}

export const SendCode = (form) => {
    return http.post('/SendCode',form)
}

export const VerifyEmailCode = (form) => {
    return http.post('/VerifyEmailCode',form)
}

export const Register = (form) => {
    return http.post('/Register',form)
}
// 获取练习列表
export const student_Practise = (params) => {
    console.log("get请求参数" + params.page)
    console.log("get请求参数" + params['page'])
    return http.get('/student/Practise',params)
}
// 老师获取知识点列表
export const teacher_Knowledge = () => {
    return http.get('/teacher/Knowledge')
}

// 错题列表
export const student_WrongProblem = (params) => {
    return http.get('/student/WrongProblem',params)
}

// 获取用户信息
export const User = () => {
    return http.get('/User')
}

// 获取班级列表
export const Class = (params) => {
    return http.get('/Class',params)
}

// 创建班级
export const teacher_Class = (form) => {
    return http.post('/teacher/Class',form)
}
// 获取学生列表
export const student_Exam = (params) => {
    return http.get('/student/Exam',params)
}
// 学生加入班级
export const student_Class = (form) => {
    return http.put('/student/Class',form)
}

// 获取班级详情
export const teacher_Class_identity = (params) => {
    return http.get('/teacher/Class/'+params.params.identity)
}
// 获取班级学生列表
export const teacher_ClassStudents = (params) => {
    return http.get('/teacher/ClassStudents',params)
}
// 修改班级代码
export const teacher_Class_put = (form) => {
    return http.put('/teacher/Class',form)
}

