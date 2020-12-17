// 封装请求工具类

const app = getApp()

const requests = (url, method, data={}) => {
  return new Promise((resolve, reject) => {
      wx.showLoading({
          title: '加载中',
      })
      wx.request({
          url: `${app.globalData.BaseUrl}${url}`,
          method: method,
          data: data,
          header: {
              'content-type': 'application/json',
          },
          success(request) {
              // console.log(request);
              if (request.data.code == 200) {
                  resolve(request.data)
              } else {
                  reject(request)
              }
          },
          fail(error) {
              console.log(error);
              reject(error.data)
          }, 
          complete: () => {
              setTimeout(() => {
                  wx.hideLoading();
              }, 100);
          }
      })
  })
}


export default requests