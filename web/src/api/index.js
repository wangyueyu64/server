import request from '@/utils/axios'

export default {
  testPost(data) {
    return request({
      url: '/API/web/testp/',
      method: 'post',
      data
    })
  },
  testGet(data){
    return request({
      url: '/API/web/testg/',
      method: 'get',
      params: { data }
    })
  },

}