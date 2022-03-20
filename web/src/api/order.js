import service from '@/utils/request'

export function fetchOrderPage(query) {
    return service({
      url: '/orders',
      method: 'get',
      params: query,
    })
}

export function fetchOrderProduct(query) {
  return service({
    url: '/orders/ids',
    method: 'get',
    params: query,
  })
}