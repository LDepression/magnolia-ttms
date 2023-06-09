import request from "./base";
import axios from "axios";
export const createCinema = (name, rows, cols) => {
  return request({
    url: "/cinema/create",
    method: "post",
    data: {
      name,
      rows,
      cols,
    },
  });
};
export const updateCinema = (name, cinemaID) => {
  return request({
    url: "/cinema/update",
    method: "put",
    data: {
      name,
      cinemaID,
    },
  });
};
export const getCinemaList = (page=1) => {
  return request({
    url: `/cinema/list/${page}`,
    method: "GET",

  });
};
export const deleteCinema = (cinema_id) => {
  return request({
    url:'/cinema/delete',
    method:'DELETE',
    data:{cinema_id}
  })
}
export const getCinemaDetail = (cinemaID) => {
  return request({
    url:"/cinema/details",
    method:'GET',
    params:cinemaID
  })
}