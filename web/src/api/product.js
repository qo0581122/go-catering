import service from '@/utils/request'

export function fetchProductList(query) {
  return service({
    url: '/products/page',
    method: 'get',
    params: query,
  })
}


export function createProduct(data) {
    return service({
      url: '/product',
      method: 'put',
      data,
    })
}

export function updateProduct(data) {
    return service({
        url: '/product',
        method: 'post',
        data,
      })
}

export function deleteProduct(id) {
  return service({
    url: '/product/' + id,
    method: 'delete',
  })
}

export function fetchProductBySpecis(specis) {
  return service({
    url: '/products/specis/' + specis,
    method: 'get',
  })
}

export function fetchProductAll() {
  return service({
    url: '/products/specis/0',
    method: 'get'
  })
}