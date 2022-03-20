import service from '@/utils/request'

export function fetchUserVipLevelList(query) {
  return service({
    url: '/user/vip/levels/page',
    method: 'get',
    params: query,
  })
}

export function fetchUserVipLevelLogList(query) {
  return service({
    url: '/user/vip/levels/logs',
    method: 'get',
    params: query,
  })
}

export function createUserVipLevels(data) {
  return service({
    url: '/user/vip/level',
    method: 'put',
    data,
  })
}

export function createUserVipLevel(data) {
  return service({
    url: '/user/vip/level',
    method: 'put',
    data,
  })
}
export function updateUserVipLevel(data) {
  return service({
    url: '/user/vip/level',
    method: 'post',
    data,
  })
}

export function deleteUserVipLevel(id) {
  return service({
    url: '/user/vip/level/' + id,
    method: 'delete',
  })
}
