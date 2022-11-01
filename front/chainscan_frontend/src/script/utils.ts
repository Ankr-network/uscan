/**
 * Get time diff
 * @param {number} startDateStamp
 * @param {number} endDateStamp
 * @return {string}
 */
const diffTime: (startDateStamp: number, endDateStamp: number) => string = function (
  startDateStamp: number,
  endDateStamp: number
): string {
  const startDate = new Date(startDateStamp * 1000);
  const endDate = new Date(endDateStamp);
  const diff = endDate.getTime() - startDate.getTime(); // .getTime();//时间差的毫秒数

  // 计算出相差天数
  const days = Math.floor(diff / (24 * 3600 * 1000));

  // 计算出小时数
  const leave1 = diff % (24 * 3600 * 1000); // 计算天数后剩余的毫秒数
  const hours = Math.floor(leave1 / (3600 * 1000));
  // 计算相差分钟数
  const leave2 = leave1 % (3600 * 1000); // 计算小时数后剩余的毫秒数
  const minutes = Math.floor(leave2 / (60 * 1000));

  // 计算相差秒数
  const leave3 = leave2 % (60 * 1000); // 计算分钟数后剩余的毫秒数
  const seconds = Math.round(leave3 / 1000);

  let returnStr = seconds + ' secs ago';

  if (seconds > 0) {
    if (seconds == 1) {
      returnStr = seconds + ' sec ago'; // + returnStr;
    }
  }

  if (minutes > 0) {
    if (minutes == 1) {
      returnStr = minutes + ' min ago'; // + returnStr;
    } else {
      returnStr = minutes + ' mins ago'; // + returnStr;
    }
  }
  if (hours > 0) {
    if (hours == 1) {
      returnStr = hours + ' hour ago'; // + returnStr;
    }
    returnStr = hours + ' hours ago'; // + returnStr;
  }
  if (days > 0) {
    if (days == 1) {
      returnStr = days + ' day ago'; // + returnStr;
    } else {
      returnStr = days + ' days ago'; // + returnStr;
    }
  }
  return returnStr;
};

const getAge = (startDateStamp: number): string => {
  return diffTime(startDateStamp, new Date().getTime());
};

const formatNumber = (input: number | bigint): string => {
  return input.toString().replace(/(\d)(?=(?:\d{3})+$)/g, '$1,');
};

const getTitle = import.meta.env.VITE_APP_TITLE;

const getParenthesesStr = (text: string): string[] => {
  const reg = /(?<=\().*(?=\))/;
  const res1 = text.match(reg)![0];
  // console.log('res1', res1);

  const arr = res1.split(/,(?![^\(]*?\))/);
  return arr;
};

export { getAge, formatNumber, getTitle, getParenthesesStr };
