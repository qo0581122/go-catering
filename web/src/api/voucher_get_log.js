import service from '@/utils/request'

export function fetchVoucherGetLogPage(query) {
  return service({
    url: '/vouchers/getLogs',
    method: 'get',
    params: query,
  })
}
