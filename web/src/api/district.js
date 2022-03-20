import service from '@/utils/request'

export function fetchDistinctPage(query) {
  return service({
    url: '/districts/page',
    method: 'get',
    params: query,
  })
}


export function createDistrict(data) {
    return service({
      url: '/district',
      method: 'put',
      data,
    })
}

export function updateDistrict(data) {
    return service({
        url: '/district',
        method: 'post',
        data,
      })
}

export function deleteDistrict(id) {
  return service({
    url: '/district/' + id,
    method: 'delete',
  })
}

export function fetchDistinctList(query) {
  return service({
    url: '/districts',
    method: 'get',
    params: query,
  })
}