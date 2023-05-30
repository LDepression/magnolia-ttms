import axios from 'axios'
import base from './base'
export const register = () => {
  return axios.post(base.register,params)
};
export const isRepeat = () => {
  return axios.get(base.isRePeat,params)
}
export const Login = () => {
  return axios.post(base.login,params)
}
export const getCode = () => {
  return axios.post(base.email,params)
}