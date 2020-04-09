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
          title="注册"
          class="mt-5"
        >
          <b-form>
            <b-form-group label="姓名">
              <b-form-input
                v-model="$v.user.name.$model"
                type="text"
                placeholder="输入用户的名称"
              ></b-form-input>
            </b-form-group>
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
            <b-form-group label="重复密码">
              <b-form-input
                v-model="$v.user.confirmPassword.$model"
                type="password"
                placeholder="输入密码"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateStatus('confirmPassword')">
                密码两次输入必须相同
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button
                variant="outline-primary"
                block
                @click="register"
              >注册</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength, sameAs } from 'vuelidate/lib/validators';
import { customValidator } from '@/helper/validator';
import { mapActions } from 'vuex';

export default {
  data() {
    return {
      user: {
        name: '',
        telephone: '',
        password: '',
        confirmPassword: '',
      },
    };
  },
  validations: {
    user: {
      name: {

      },
      telephone: {
        required,
        telephone: customValidator,
      },
      password: {
        required,
        minLength: minLength(6),
      },
      confirmPassword: {
        required,
        minLength: minLength(6),
        sameAs: sameAs('password'),
      },
    },
  },
  methods: {
    ...mapActions('userModule', { userRegister: 'register' }),
    validateStatus(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      this.userRegister(this.user).then(() => {
        this.$router.replace({ name: 'home' });
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });
    },
  },
  name: 'Register',
};
</script>

<style lang="scss">
</style>
