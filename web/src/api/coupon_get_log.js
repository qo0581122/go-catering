import service from '@/utils/request'

export function fetchCouponGetLogPage(query) {
  return service({
    url: '/coupons/getLogs',
    method: 'get',
    params: query,
  })
}
