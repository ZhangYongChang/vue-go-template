<template>
  <div id="app">
    <navbar></navbar>
    <b-container>
      <router-view/>
    </b-container>
  </div>
</template>
<script>
import Navbar from './views/layout/Navbar.vue';

export default {
  name: 'App',
  components: {
    Navbar,
  },
  data() {
    return {
      title: '',
    };
  },
};
</script>>

<style lang="scss">
</style>
