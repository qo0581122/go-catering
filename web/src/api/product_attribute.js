import service from '@/utils/request'

export function fetchProductAttributeList(query) {
  return service({
    url: '/product/attributes/page',
    method: 'get',
    params: query,
  })
}


export function createProductAttribute(data) {
    return service({
      url: '/product/attribute',
      method: 'put',
      data,
    })
}

export function updateProductAttribute(data) {
    return service({
        url: '/product/attribute',
        method: 'post',
        data,
      })
}

export function deleteProductAttribute(id) {
  return service({
    url: '/product/attribute/' + id,
    method: 'delete',
  })
}

export function fetchProductAttributeAll() {
  return service({
    url: '/product/attributes',
    method: 'get',
  })
}