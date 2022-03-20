import service from '@/utils/request'

export function fetchCityPage(query) {
  return service({
    url: '/citys/page',
    method: 'get',
    params: query,
  })
}


export function createCity(data) {
  return service({
    url: '/city',
    method: 'put',
    data,
  })
}

export function updateCity(data) {
  return service({
    url: '/city',
    method: 'post',
    data,
  })
}

export function deleteCity(id) {
  return service({
    url: '/city/' + id,
    method: 'delete',
  })
}

export function fetchListAll() {
  return service({
    url: '/city/all',
    method: 'get',
  })
}

export function fetchCityList(query) {
  return service({
    url: '/citys',
    method: 'get',
    params: query
  })
}