import axios from "axios";
import Cookies from "js-cookie";
import { getTokenAgain } from "./user";
import { getToken, setToken,getExpires} from "./cookie";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router"
const router = useRouter();
console.log(router);
const httpInstance = axios.create({
  baseURL: "http://192.168.3.233:9090/api/v1",
  timeout: 5000,
});
// httpInstance.defaults.headers.common['x_token'] = `${token}`;
//获取cookie 的时间

// 判断token是否即将过期
function getCookieExpirationDate(cookieName) {
  var cookies = document.cookie.split(";");

  for (var i = 0; i < cookies.length; i++) {
    var cookie = cookies[i].trim();
    if (cookie.indexOf(cookieName) == 0) {
      var cookieValue = cookie.substring(cookieName.length + 1);
      var expiresIndex = cookieValue.indexOf("Expires=");
      if (expiresIndex != -1) {
        var expiresValue = cookieValue.substring(expiresIndex + 8).trim();
        return new Date(expiresValue);
      }
    }
  }

  // Cookie not found or no expiration time set
  return null;
}
var x = getCookieExpirationDate('AccessToken')

function isTokenExpired() {
  let curTime = new Date().getTime();
  let expiresTime = Number(getExpires()) - curTime;
  // 还差10分钟即将过期或者已经过期了，但过期时间在20分钟内
if ((expiresTime >= 0 && expiresTime < 600000) || (expiresTime < 0 && Math.abs(expiresTime) <= 1200000)) {
  return true
}
  return false;
}



// 是否正在刷新的标志
window.isRefreshing = false;
// 存储请求的数组
let cacheRequestArr = [];
// 将所有的请求都push到数组中,其实数组是[function(token){}, function(token){},...]
function cacheRequestArrHandle(cb) {
  cacheRequestArr.push(cb);
}
// 数组中的请求得到新的token之后自执行，用新的token去重新发起请求
function afreshRequest(token) {
  cacheRequestArr.map((cb) => cb(token));
  cacheRequestArr = [];
}
let accessToken = localStorage.getItem('AccessToken')
httpInstance.interceptors.request.use(
  (config) => {
    // 如果存在token，将其添加到请求头中
    if(getToken()){
    config.headers["x_token"] = accessToken
   
   
    if ( isTokenExpired() && config.url !== "/refreshToken") {
      // 所有的请求来了，先判断是否正在刷新token，
      // 如果不是，将刷新token标志置为true并请求刷新token.
      // 如果是，则先将请求缓存到数组中
      // 等到刷新完token后再次重新请求之前缓存的请求接口即可
      
      if (!window.isRefreshing) {
        //标志改为true，表示正在刷新
        window.isRefreshing = true;
        console.log(1111);
        getTokenAgain(accessToken,localStorage.getItem("RefreshToken"))
          .then((res) => {
            
            console.log(res);
            if (res.data.status_code == 0) {
              console.log(111);
              //更新cookie
              setToken(res.data.data.AccessToken,new Date(res.data.data.expiresAt));
              setExpires(res.data.data.expiresAt, new Date(res.data.data.expiresAt));
              // 将刷新的token替代老的token
              config.headers["x_token"] = getToken();
              accessToken = getToken()
              // 刷新token完成后重新请求之前的请求
              afreshRequest(getToken());
            }
          })
          .catch((err) => {
            console.log("refreshToken err =>" + err);
          })
          .finally(() => {
            window.isRefreshing = false;
          });
          const token = getToken()
        // 下面这段代码一定要写，不然第一个请求的接口带过去的token还是原来的，要将第一个请求也缓存起来
        let retry = new Promise((resolve) => {
          cacheRequestArrHandle((token) => {
            config.headers["x_token"] = token; // token为刷新完成后传入的token
            // 将请求挂起
            resolve(config);
          });
        });
        return retry;
      } else {
        let retry = new Promise((resolve) => {
          cacheRequestArrHandle((token) => {
            config.headers["x_token"] = token; // token为刷新完成后传入的token
            // 将请求挂起
            resolve(config);
          });
        });
        return retry;
      }
    } else {
      return config;
    }
  }
  if(!getToken()){
    if(config.url!=='/login'){
       
        console.log(router);
        // router.push({path:"/login"})
        ElMessage({
          message: "登录信息过期，请重新登录!!",
          center: true,
          type: "error",
        });
        localStorage.removeItem("AccessToken")
        localStorage.removeItem("RefreshToken")
        localStorage.removeItem("user")
    }
    
  }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// httpInstance.interceptors.response.use(
//   response => {
//       // 当response.data.re为401, 则判断token已经过期
//       // /oauth/token为刷新token的接口, 需要排除掉
//       if ((isTokenExpired() || response.data.ret === 401) && !response.config.url.includes('/user/refreshToken')) {
//           if (!isRefreshing) {
//               isRefreshing = true;
//               // 将刷新token的方法放在vuex中处理了, 可见下面区块代码
//               store.dispatch('refreshToken').then((res) => {
//                   // 当刷新成功后, 重新发送缓存请求
//                   onAccessTokenFetched(res);
//               }).catch(() => {
//                   // 刷新token报错的话, 就需要跳转到登录页面
//                   window.location = '/#/guide/login';
//               }).finally(() => {
//                   isRefreshing = false;
//               });
//           }
//           // 将其他接口缓存起来 -- 这个Promise函数很关键
//           const retryOriginalRequest = new Promise((resolve) => {
//               // 这里是将其他接口缓存起来的关键, 返回Promise并且让其状态一直为等待状态,
//               // 只有当token刷新成功后, 就会调用通过addSubscriber函数添加的缓存接口,
//               // 此时, Promise的状态就会变成resolve
//               addSubscriber((newToken) => {
//                   // 表示用新的token去替换掉原来的token
//                   response.config.headers.Authorization = `bearer ${newToken}`;
//                   // 替换掉url -- 因为baseURL会扩展请求url
//                   response.config.url = response.config.url.replace(response.config.baseURL, '');
//                   // 用重新封装的config去请求, 就会将重新请求后的返回
//                   resolve(service(response.config));
//               });
//           });
//           return retryOriginalRequest;
//       }
//       // ========================

//       // .... 省略其他代码.....
//   })

export default httpInstance;
