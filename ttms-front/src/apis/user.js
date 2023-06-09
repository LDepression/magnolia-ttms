import axios from "axios";
import request from "./base";
import Cookies from "js-cookie";

// const instance = axios.create({
//   baseURL
// })
export const register = (Email, EmailCode, Password, Username, Gender) => {
  return request({
    url: "/register",
    method: "POST",
    data: {
      Email,
      EmailCode,
      Password,
      Username,
      Gender,
    },
  });
};
export const isRepeat = (data) => {
  return request({
    url: `/isRePeat?Username=${data}`,
    method: "GET",
  });
};
export const Login = (Email, Password, LoginType) => {
  return request({
    url: "/login",
    method: "POST",
    data: {
      Email,
      Password,
      LoginType,
    },
  });
};
// export const getCode = (data) => {
//   return axios.post('http://192.168.3.17:9090/api/v1/email/send',data)
// }
export const findUser = (UserName) => {
  // return axios.post(base.findUser,{
  //   Headers:{
  //     "Authorization":Cookies.get('AccessToken')
  //   },params
  // })
  return request({
    url: "/user/findUser",
    method: "POST",
    headers: {
      x_token: Cookies.get("AccessToken"),
    },
    data: {
      UserName
    },
  });
};
export const getToken = () => {
  return request({
    url: "/user/refreshToken",
    method: "POST",
    data: {
      AccessToken: Cookies.get("AccessToken"),
      RefreshToken: localStorage.getItem("RefreshToken"),
    },
  });
};
export const getCode = (email) => {
  return request({
    url: "/email/send",
    method: "POST",
    data: {
      email,
    },
  });
};
export const changeAvatar = (newAvatar) => {
  return request({
    url: "/user/modifyAvatar",
    method: "PUT",
    data: {
      newAvatar,
    },
  });
};
export const getTokenAgain = (AccessToken, RefreshToken) => {
  return request({
    url: "/refreshToken",
    method: "POST",
    data: {
      AccessToken,
      RefreshToken,
    },
  });
};
export const ping = () => {
  return request({
    url: "/user/ping",
    method: "GET",
  });
};
export const changeEmail = (EmailCode, NewEmail) => {
  return request({
    url: "/user/modifyEmail",
    method: "PUT",
    data: {
      EmailCode,
      NewEmail,
    },
  });
};
export const changePassword = (EmailCode, NewPassword) => {
  return request({
    url: "/user/modifyPassword",
    method: "PUT",
    data: {
      EmailCode,
      NewPassword,
    },
  });
};
export const fileUpLoad = (formData) => {
  return request({
    url:'file/upload',
    method:'POST',
   
    // headers:{
    //   'Content-Type':'multipart/form-data'
    // },

      data:formData,
      headers: {
        'Content-Type': 'multipart/form-data'
      },
    
  })
}
export const changeDescription = (gender, signature) => {
  return request({
    url: "/user/updateInfo",
    method: "PUT",
    data: {
      gender,
      signature,
    },
  });
};
