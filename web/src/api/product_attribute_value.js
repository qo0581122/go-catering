import service from '@/utils/request'

export function fetchProductAttributeValueList(query) {
  return service({
    url: '/product/attribute/values/page',
    method: 'get',
    params: query,
  })
}


export function createProductAttributeValue(data) {
    return service({
      url: '/product/attribute/value',
      method: 'put',
      data,
    })
}

export function updateProductAttributeValue(data) {
    return service({
        url: '/product/attribute/value',
        method: 'post',
        data,
      })
}

export function deleteProductAttributeValue(id) {
  return service({
    url: '/product/attribute/value/' + id,
    method: 'delete',
  })
}

export function fetchProductAttributeValueAll() {
  return service({
    url: '/product/attribute/values',
    method: 'get',
  })
}