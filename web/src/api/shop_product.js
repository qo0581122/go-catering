import service from '@/utils/request'

export function fetchShopProductPage(query) {
  return service({
    url: '/shop/products/page',
    method: 'get',
    params: query,
  })
}


export function createShopProduct(data) {
    return service({
      url: '/shop/product',
      method: 'put',
      data,
    })
}

export function updateShopProduct(data) {
    return service({
        url: '/shop/product',
        method: 'post',
        data,
      })
}

export function deleteShopProduct(id) {
  return service({
    url: '/shop/product/' + id,
    method: 'delete',
  })
}

export function fetchShopProductList(params) {
  return service({
    url: '/shop/products',
    method: 'get',
    params:params,
  })
}