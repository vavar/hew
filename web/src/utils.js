import moment from 'moment-timezone';

const TIMEZONE = 'Asia/Bangkok';
const FORMAT = 'YYYY/MM/DD HH:mm';

export default {
  formatDate(dateStr) {
    return moment(dateStr).tz(TIMEZONE).format(FORMAT);
  },
};
