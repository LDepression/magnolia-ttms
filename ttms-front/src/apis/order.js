import request from "./base";
export const toPay = (params) => {
  return request({
    url:'/tickets/pay',
    method:'POST',
    data:params
  })
}
export const tryLock = (planID,cinemaID,seatIDs) => {
  return request({
    url:'/tickets/soldTicket',
    method:'POST',
    data:{planID,cinemaID,seatIDs}
  })
}
export const isPay = (orderID) => {
  return request({
    url:'/tickets/isPay',
    method:'POST',
    data:{orderID}
  })
}
export const returnTicket = (OrderID,PlanID,SeatIDs) => {
  return request({
    url:'/tickets/backTicket',
    method:'POST',
    data:{OrderID,PlanID,SeatIDs}
  })
}
export const getAllOrder = () => {
  return request({
    url:'/tickets/getOrders',
    method:'get',

  })
}