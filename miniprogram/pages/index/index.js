//index.js
//获取应用实例
const app = getApp();

Page({
  data: {
  },
  handleNavigateToLocation() {
    wx.navigateTo({
      url: '../location/location'
    })

  },
  handleToShowWorkTime() {
    wx.showModal({
      title: '温馨提示',
      content: '敬爱的客户们，请注意避免饭点和休息时间哈~',
      showCancel: false,
      confirmText: '好的',
      confirmColor: '#D81E06'
    })
  },
  handleToCallPhone() {
    wx.makePhoneCall({
      phoneNumber: '0755-84665789'
    })
  },
  handleToCallMobilePhone() {
    wx.makePhoneCall({
      phoneNumber: '13715236126'
    })
  }

});