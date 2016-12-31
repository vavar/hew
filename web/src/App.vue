<template>
  <div id="app" v-if="$auth.ready()">
    <navbar></navbar>
    <div class="wrapper">
      <transition name="md-router" appear>
        <router-view></router-view>
      </transition>
    </div>
  </div>
  <div id="loading" v-else>loading</div>
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
}
.wrapper {
  margin-top: 20px;
}
</style>
