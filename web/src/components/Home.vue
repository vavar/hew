<template>
  <div class="home">
    <h1>Dashboard</h1>
    <order-today></order-today>
  </div>
</template>

<script>
import OrderToday from './OrderToday';

export default {
  name: 'home',
  data() {
    return {};
  },
  components: {
    OrderToday,
  },
};
</script>