import service from '@/utils/request'

export function fetchUserAddressList(query) {
  return service({
    url: '/user/addresses',
    method: 'get',
    params: query,
  })
}
