import request from "./base";
import axios from "axios"
export const addMovie2 = (params) => {
  return axios.post("http://127.0.0.1:9090/api/v1/movie/create",{params})
}
export const addMovie = (params) => {
  return request({
    url:'/movie/create',
    method:'POST',
    data:params
  })
}
export const getAllMovie = () => {
  return request({
    url:'/movie/list',
    method:'GET',
  })
}
export const deleteMovie = (movie_id) => {
  return request({
    url:'/movie',
    method:'DELETE',
    data:{movie_id}
  })
}
//根据阅读数进行排序
export const sortByRead = () => {
  return request({
    url:'/movie/byReadCount',
    method:'GET'
   })
}
export const sortByExpect = (page=1) => {
  return request({
    url:`/movie/byExpectedNums/${page}`,
    method:'GET',
    
  })
}

export const updateMovie = () => {
  return request({
    url:'/movie/update',
    method:'put',
    data:{movieID,name,avatar,actors,director,area,content}
  })
}
// export const addMovie = (form) => {
//   return request({
//     url:'/movie/create',
//     method:'POST',
//     params:{
//       actors:form.actors,
//       area:form.area,
//       avatar:form.avatar,
//       content:form.content,
//       director:form.director,
//       duration:form.duration,
//       name:form.name,
//       show_time:form.show_time,
//       tags:form.tags
//     }

//   })
// }
export const getMovieDetail = (movie_id) => {
  return request({
    url:'/movie/details',
    method:'POST',
    data:{movie_id}
    
  })
}
export const getMovieInfoByTag = (tag,area,order_type,start_time,end_time) => {
  return request({
    url:'/movie/byTagAreaPeriod',
    method:'POST',
    data:{tag,area,order_type,start_time,end_time}
  })
}
export const getMovieByName = (key) => {
  return request({
    url:'/movie/byNameOrContent', 
    method:'POST',
    data:{
      key
    }
  })
}