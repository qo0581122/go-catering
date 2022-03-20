import service from '@/utils/request'

export function fetchShopCategoryPage(query) {
  return service({
    url: '/shop/categorys/page',
    method: 'get',
    params: query,
  })
}


export function createShopCategory(data) {
    return service({
      url: '/shop/category',
      method: 'put',
      data,
    })
}

export function updateShopCategory(data) {
    return service({
        url: '/shop/category',
        method: 'post',
        data,
      })
}

export function deleteShopCategory(id) {
  return service({
    url: '/shop/category/' + id,
    method: 'delete',
  })
}

export function fetchShopCategoryList(params) {
  return service({
    url: '/shop/categorys',
    method: 'get',
    params:params,
  })
}