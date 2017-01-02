<template>
  <div class="countdown md-toolbar md-accent">
    <md-icon>schedule</md-icon>
    <div class="countdown-block">
      <div class="countdown-title">Closed in</div>
      <div>
      <span v-if="days">{{ days | two_digits }} Days </span>
      <span v-if="hours">{{ hours | two_digits }} Hours </span>
      <span v-if="minutes || hours">{{ minutes | two_digits }} Minutes </span>
      <span >{{ seconds | two_digits }} Seconds</span>
      </div>
    </div>
  </div>
</template>

<script>
  export default {

    /* ready function will be here */

    props: ['expired'],
    created() {
      console.log('>>' , this.expired);
      this.date = Math.trunc(Date.parse(this.expired) / 1000);
      const intervalID = setInterval(() => {
        this.now = Math.trunc((new Date()).getTime() / 1000);
        if ( this.now > this.date ) {
          clearInterval( intervalID );
        }
      }, 1000);
    },
    computed: {
      seconds() {
        return (this.date - this.now) % 60;
      },

      minutes() {
        return Math.trunc((this.date - this.now) / 60) % 60;
      },

      hours() {
        return Math.trunc((this.date - this.now) / 60 / 60) % 24;
      },

      days() {
        return Math.trunc((this.date - this.now) / 60 / 60 / 24);
      }
    },
    data() {
      return {
        now: Math.trunc((new Date()).getTime() / 1000),
        date: 0,
      }
    }

    /* Computed properties will be here */
  }
</script>

<style>
.countdown-title{
  font-weight: bold;
}
.countdown {
  display:flex;
  flex-direction: row;
}
.countdown-block {
    margin-left:10px;
    display: flex;
    flex-direction: column;
}

.text {
    color: #1abc9c;
    font-size: 10px;
    text-align: center;
}

.digit {
    color: #ecf0f1;
    font-size: 14px;
    text-align: center;
}
</style>
