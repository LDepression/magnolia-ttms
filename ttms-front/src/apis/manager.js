import axios from "axios";
import request from "./base";
import Cookies from "js-cookie";
export const getUsers = (page) => {
  return request({
    url:'/manager/list',
    method:'GET',
  })
}
export const deleteUser = (userID) => {
  return request({
    url:`/manager/${userID}`,
    method:'DELETE',
    query:{userID}
  })
}
export const ping = () => {
  return request({
    url:'/manager/ping',
    method:'GET',
  })
}
export const updateRole = (Username) => {
  return request({
    url:'/manager/createManager',
    method:'POST',
    data:{Username}
  })
}