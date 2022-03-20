import service from '@/utils/request'

export function fetchUserIntegrationList(query) {
  return service({
    url: '/user/integrations/page',
    method: 'get',
    params: query,
  })
}

export function fetchUserIntegrationLogList(query) {
  return service({
    url: '/user/integrations/logs',
    method: 'get',
    params: query,
  })
}

export function changeUserIntegration(data) {
  return service({
    url: '/user/integration/change',
    method: 'get',
    data,
  })
}
