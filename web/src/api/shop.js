import service from '@/utils/request'

export function fetchShopPage(query) {
  return service({
    url: '/shops/page',
    method: 'get',
    params: query,
  })
}


export function createShop(data) {
    return service({
      url: '/shop',
      method: 'put',
      data,
    })
}

export function updateShop(data) {
    return service({
        url: '/shop',
        method: 'post',
        data,
      })
}

export function deleteShop(id) {
  return service({
    url: '/shop/' + id,
    method: 'delete',
  })
}
