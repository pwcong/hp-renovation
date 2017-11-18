// pages/location/location.js

const LOCATION = {
  latitude: 22.694398,
  longitude: 114.112426
}

Page({

  data: {
    map: null,
    latitude: LOCATION.latitude,
    longitude: LOCATION.longitude,
    markers: [{
      id: 0,
      latitude: 22.694398,
      longitude: 114.112426,
      width: 50,
      height: 50
    }],
    controls: [
      {
        id: 1,
        iconPath: '../../assets/imgs/icon_location.png',
        position: {
          left: 10,
          top: 425,
          width: 65,
          height: 65
        },
        clickable: true
      },
      {
        id: 2,
        iconPath: '../../assets/imgs/icon_target.png',
        position: {
          left: 10,
          top: 360,
          width: 65,
          height: 65
        },
        clickable: true
      },
    ]
  },
  markertap(e) {
    // console.log(e.markerId)
  },
  controltap(e) {
    const that = this
    // console.log(e.controlId)
    switch (e.controlId) {

      case 1:
        that.data.map.moveToLocation()
        break;
      case 2:
        that.data.map.includePoints({
          points: [{ latitude: LOCATION.latitude, longitude: LOCATION.longitude }]
        })
        break;
      default: break;

    }
  },
  regionchange(e) {
    this.data.map.getCenterLocation({
      success: function (res) {
        console.log(res)
      }
    })
  },
  onLoad() {
    const that = this
    const map = wx.createMapContext("myMap", that)
    that.setData({
      map
    })
  }
})