<template>
    <div>
      <label>{{userInfo.name}}</label>
      <label>{{userInfo.telephone}}</label>
    </div>
</template>

<script>
import { mapState } from 'vuex';

export default {
  name: 'Private',
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
};
</script>

<style lang="scss">

</style>
