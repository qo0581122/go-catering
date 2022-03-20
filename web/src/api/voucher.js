import service from '@/utils/request'

export function fetchVoucherPage(query) {
  return service({
    url: '/vouchers/page',
    method: 'get',
    params: query,
  })
}


export function createVoucher(data) {
    return service({
      url: '/voucher',
      method: 'put',
      data,
    })
}

export function updateVoucher(data) {
    return service({
        url: '/voucher',
        method: 'post',
        data,
      })
}

export function deleteVoucher(id) {
  return service({
    url: '/voucher/' + id,
    method: 'delete',
  })
}
