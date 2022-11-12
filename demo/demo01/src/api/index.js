import http from '../utils/request'

export const LoginData = (form) => {
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