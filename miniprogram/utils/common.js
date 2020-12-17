
/*
一些通用的工具函数

*/



/*
防抖函数
*/
function debounce(fun, delay) {
  var delay_time = delay || 1000
  console.log("防抖")
  return function (args) {
      let that = this;
      let _args = args;
      clearTimeout(fun.id);
      fun.id = setTimeout(function () {
          fun.call(that, _args)
      }, delay_time)
  }
}

export default {
  debounce
};