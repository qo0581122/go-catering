import service from '@/utils/request'

export function fetchUserAddressTagList(query) {
  return service({
    url: '/user/address/tags/page',
    method: 'get',
    params: query,
  })
}


export function createUserAddressTag(data) {
  return service({
    url: '/user/address/tag',
    method: 'put',
    data,
  })
}

export function updateUserAddressTag(data) {
  return service({
      url: '/user/address/tag',
      method: 'post',
      data,
    })
}

export function deleteUserAddressTag(id) {
return service({
  url: '/user/address/tag/' + id,
  method: 'delete',
})
}
