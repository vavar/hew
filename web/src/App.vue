<template>
  <div class="app-viewport" id="app" v-if="$auth.ready()">
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
html,
body,
.app-viewport {
  height: 100%;
  overflow: hidden;
}

.app-viewport {
  display: flex;
  flex-flow: column;
}
.wrapper {
  margin-top: 20px;
}
</style>
