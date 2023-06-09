import request from "./base";
export const getPlan = (page) => {
  return request({
    url: `/plan/list/${page}`,
    method: "GET",
    data: page,
  });
};
export const deletePlan = (planID) => {
  return request({
    url: "/plan/delete",
    method: "DELETE",
    data: { planID },
  });
};
export const createPlan = (movieID, cinemaID, price, startTime) => {
  return request({
    url: "/plan/create",
    method: "POST",
    data: {
      movieID,
      cinemaID,
      price,
      startTime,
    },
  });
};
