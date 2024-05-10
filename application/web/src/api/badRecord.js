import request from '@/utils/request'

// 添加不良记录(公安局)
export function badRecordAdd(data) {
  return request({
    url: '/badRecordAdd',
    method: 'post',
    data
  })
}

// 获取用户最新不良记录（单条）
export function queryBadRecordLatestByIdCard(data) {
  return request({
    url: '/queryBadRecordLatestByIdCard',
    method: 'get',
    data
  })
}

// 获取用户最新不良记录（所有）
export function queryBadRecordListByIdCard(data) {
  return request({
    url: '/queryBadRecordListByIdCard',
    method: 'get',
    params:data
  })
}