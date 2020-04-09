<template>
  <div>
    <b-navbar
      toggleable="lg"
      type="dark"
      variant="info"
    >
      <b-container>
        <b-navbar-brand @click="home">主页</b-navbar-brand>
        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
        <b-collapse
          id="nav-collapse"
          is-nav
        >
          <b-navbar-nav class="ml-auto">
            <b-nav-item-dropdown
              right
              v-if="userInfo"
            >
              <template v-slot:button-content>
                <em>{{userInfo.name}}</em>
              </template>
              <b-container>
                <b-dropdown-item @click="$router.replace({name: 'business'})">工作台</b-dropdown-item>
                <b-dropdown-item @click="$router.replace({name: 'private'})">个人主页</b-dropdown-item>
                <b-dropdown-item @click="logout">登出</b-dropdown-item>
              </b-container>
            </b-nav-item-dropdown>
            <b-container v-if="!userInfo">
              <b-nav-item
                v-if="$route.name !== 'login'"
                @click="$router.replace({name: 'login'})"
              >登录</b-nav-item>
              <b-nav-item
                v-else-if="$route.name !== 'register'"
                @click="$router.replace({name: 'register'})"
              >注册</b-nav-item>
            </b-container>
          </b-navbar-nav>
        </b-collapse>
      </b-container>
    </b-navbar>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  name: 'Navbar',
  components: [
  ],
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
  methods: {
    ...mapActions('userModule', ['logout']),
    home() {
      if (this.$route.name !== 'home') {
        this.$router.replace({ name: 'home' });
      }
    },
  },
};
</script>

<style lang="scss">
</style>
