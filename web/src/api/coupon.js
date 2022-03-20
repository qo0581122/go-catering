import service from '@/utils/request'

export function fetchCouponPage(query) {
  return service({
    url: '/coupons/page',
    method: 'get',
    params: query,
  })
}


export function createCoupon(data) {
  return service({
    url: '/coupon',
    method: 'put',
    data,
  })
}

export function updateCoupon(data) {
  return service({
    url: '/coupon',
    method: 'post',
    data,
  })
}

export function deleteCoupon(id) {
  return service({
    url: '/coupon/' + id,
    method: 'delete',
  })
}
