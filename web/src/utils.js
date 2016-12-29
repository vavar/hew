import moment from 'moment-timezone';

const TIMEZONE = 'Asia/Bangkok';
const DISPLAY_FORMAT = 'YYYY/MM/DD hh:mm A';
const GO_TIME_FORMAT = 'YYYY-MM-DD[T]HH:mm:ss[Z]';

export default {
  formatDate(dateStr) {
    return moment(dateStr).tz(TIMEZONE).format(DISPLAY_FORMAT);
  },
  getLocalTime() {
    return moment().tz(TIMEZONE).format(GO_TIME_FORMAT);
  },
  formatDateForAPI(dateStr) {
    return moment.utc(moment(dateStr).tz(TIMEZONE)).format(GO_TIME_FORMAT);
  }
};
