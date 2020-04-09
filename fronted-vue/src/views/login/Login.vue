<template>
  <div>
    <b-row>
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card
          title="登录"
          class="mt-5"
        >
          <b-form>
            <b-form-group label="手机号">
              <b-form-input
                v-model="$v.user.telephone.$model"
                type="number"
                placeholder="输入手机号"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateStatus('telephone')">
                手机号错误
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                placeholder="输入密码"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateStatus('password')">
                密码必须大于6位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button
                variant="outline-primary"
                block
                @click="login"
              >登录</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength } from 'vuelidate/lib/validators';
import { customValidator } from '@/helper/validator';
import { mapActions } from 'vuex';

export default {
  data() {
    return {
      user: {
        telephone: '',
        password: '',
      },
    };
  },
  validations: {
    user: {
      telephone: {
        required,
        telephone: customValidator,
      },
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    ...mapActions('userModule', { userLogin: 'login' }),
    validateStatus(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    login() {
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      this.userLogin(this.user).then(() => {
        this.$router.replace({ name: 'home' });
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '登录失败',
          variant: 'danger',
          solid: true,
        });
      });
    },
  },
  name: 'Login',
};
</script>

<style lang="scss">
</style>
