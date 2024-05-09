import request from '@/utils/request'

// 添加不良记录(公安局)
export function badRecordAdd(data) {
  return request({
    url: '/badRecordAdd',
    method: 'post',
    data
  })
}

// 获取不良记
export function queryBadRecordLatestByIdCard(data) {
  return request({
    url: '/queryBadRecordLatestByIdCard',
    method: 'get',
    data
  })
}
