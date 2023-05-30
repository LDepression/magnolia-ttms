//接口路径配置
let baseURL = 'https://127.0.0.1:9090/ttms';
const base = {
  register:baseURL+'/register',//注册接口
  login:baseURL+'/login',
  email:baseURL+'/email/send',
  isRePeat:baseURL + '/isRePeat'
}
export default base