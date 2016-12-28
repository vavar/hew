<template>
  <div id="app">
    <navbar></navbar>
    <div class="wrapper">
      <transition name="md-router" appear>
        <router-view></router-view>
      </transition>
    </div>
  </div>
</template>

<script>
import Navbar from './components/Navbar';

export default {
  name: 'app',
  components: {
    Navbar,
  },
};
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
.wrapper {
  margin-top: 20px;
}
</style>
