const formatTime = date => {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()

  return [year, month, day].map(formatNumber).join('/') + ' ' + [hour, minute, second].map(formatNumber).join(':')
}

const getYear = date => {
  return [date.getFullYear()].map(formatNumber)
}

const getMonth = date => {
  return [date.getMonth()+1].map(formatNumber)
}

const getDay = date => {
  return [date.getDate()].map(formatNumber)
}

const getDateName = date => {
  var weekDay = ["星期天", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"]; 
  return weekDay[date.getDay()]
}

const formatNumber = n => {
  n = n.toString()
  return n[1] ? n : '0' + n
}

module.exports = {
  formatTime: formatTime,
  getYear: getYear,
  getMonth: getMonth,
  getDay: getDay,
  getDateName: getDateName
}
