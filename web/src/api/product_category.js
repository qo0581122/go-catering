import service from '@/utils/request'

export function fetchProductCategoryList(query) {
  return service({
    url: '/product/categorys/page',
    method: 'get',
    params: query,
  })
}


export function createProductCategory(data) {
    return service({
      url: '/product/category',
      method: 'put',
      data,
    })
}

export function updateProductCategory(data) {
    return service({
        url: '/product/category',
        method: 'post',
        data,
      })
}

export function deleteProductCategory(id) {
  return service({
    url: '/product/category/' + id,
    method: 'delete',
  })
}

export function fetchProductCategoryAll() {
  return service({
    url: '/product/categorys',
    method: 'get',
  })
}