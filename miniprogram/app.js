//app.js
App({
  globalData: {
    userInfo: null
  },
  getUserInfo(cb) {

    const userInfo = this.globalData.userInfo;

    if (userInfo) {
      cb && cb(userInfo);
      return;
    }

    wx.getUserInfo({
      success: res => {
        // 可以将 res 发送给后台解码出 unionId
        this.globalData.userInfo = res.userInfo;
        cb && cb(res.userInfo);

      }
    });

  },
  onLaunch() {

  },
});