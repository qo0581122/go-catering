import service from '@/utils/request'
export function fetchList(query) {
  return service({
    url: '/provinces/page',
    method: 'get',
    params: query,
  })
}


export function createProvince(data) {
    return service({
      url: '/province',
      method: 'put',
      data,
    })
}

export function updateProvince(data) {
    return service({
        url: '/province',
        method: 'post',
        data,
      })
}

export function deleteProvince(id) {
  return service({
    url: '/province/' + id,
    method: 'delete',
  })
}

export function fetchProvinceList() {
  return service({
    url: '/provinces',
    method: 'get',
  })
}