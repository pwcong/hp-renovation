//index.js
//获取应用实例
const app = getApp();

Page({
  data: {
    userInfo: {}
  },

  onLoad: function () {
    app.getUserInfo(userInfo => {

      this.setData({
        userInfo: userInfo
      });

    });
  }

});