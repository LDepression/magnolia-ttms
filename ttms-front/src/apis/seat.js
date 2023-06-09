import request from "./base";
export const getSeat = (planID) => {
  return request({
    url:`/tickets/seats/${planID}`,
    method:'GET',
    data:planID
  })
}
export const updateSeat = (cinemaID,seatID,status) => {
  return request({
    url:'/seat/update',
    method:'PUT',
    data:{
      cinemaID,seatID,status
    }
  })
}
