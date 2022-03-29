import service from '@/utils/request'

export function fetchProductBatchList(query) {
  return service({
    url: '/product/batchs/page',
    method: 'get',
    params: query,
  })
}


export function createProductBatch(data) {
    return service({
      url: '/product/batch',
      method: 'put',
      data,
    })
}

export function updateProductBatch(data) {
    return service({
        url: '/product/batch',
        method: 'post',
        data,
      })
}

export function deleteProductBatch(id) {
  return service({
    url: '/product/batch/' + id,
    method: 'delete',
  })
}

export function fetchProductBatchAll() {
  return service({
    url: '/product/batchs',
    method: 'get',
  })
}