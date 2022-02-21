import axios from "axios";

axios.interceptors.response.use (
    response => response,
    error =>{
        if(error.response.status === 401){
            window.location.href = "/"
        }
        return Promise.reject(error)
    }
)
export default{
    install: (app) =>{
        app.config.globalProperties.$axios = axios;
    }
}