<template>
  <div class="LoginView">
    <div class="Content">
      <form @submit.prevent="handleLogin" novalidate>
        <div class="Title">Đăng nhập</div>
        <div class="FormBox">
          <Field name="email" v-slot="{ field }">
            <div class="FormGroup">
              <label class="FormLabel" for="EmailInput">Email</label>
              <Input
                ref="emailInputRef"
                placeholder="Email"
                id="EmailInput"
                v-bind="field"
                :status="bindStatusInput('email')"
              />
              <span v-if="errors.email" class="ErrorMessage">{{ errors.email }}</span>
            </div>
          </Field>
          <Field name="password" v-slot="{ field }">
            <div class="FormGroup">
              <label for="PasswordInput" class="FormLabel">Mật khẩu</label>
              <InputPassword
                placeholder="Mật khẩu"
                type="password"
                id="PasswordInput"
                v-bind="field"
                :status="bindStatusInput('password')"
              />
              <span v-if="errors.password" class="ErrorMessage">{{ errors.password }}</span>
            </div>
          </Field>
          <div class="FormGroup">
            <Button type="primary" html-type="submit" :loading="isLoading">Đăng nhập</Button>
          </div>
          <div class="FormGroup">
            Chưa có tài khoản?
            <router-link class="ButtonRegister" :to="{ name: RouterName.Register }"
              >Đăng ký</router-link
            >
            ngay
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import RouterName from '@/configs/RouterName';
import { Input, Button, InputPassword, message } from 'ant-design-vue';
import { onMounted, ref } from 'vue';
import * as yup from 'yup';
import { useForm, Field } from 'vee-validate';
import UserAPI from '@/apis/users/UserAPI';
import ConfigGlobal from '@/configs/ConfigGlobal';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const emailInputRef = ref<HTMLInputElement | null>(null);
const isLoading = ref<boolean>(false);

const authStore = useAuthStore();
const router = useRouter();

const loginSchema = yup.object({
  email: yup.string().required('Email là bắt buộc').email('Email không hợp lệ').trim(),
  password: yup.string().required('Mật khẩu là bắt buộc').trim(),
});

const { handleSubmit, errors } = useForm<yup.InferType<typeof loginSchema>>({
  validationSchema: loginSchema,
});

const fieldNames = ['email', 'password'] as const;
type FieldName = (typeof fieldNames)[number];

const bindStatusInput = (name: FieldName) => {
  const error = errors.value[name];
  return error ? 'error' : undefined;
};

onMounted(() => {
  focusEmailInput();
});

/**
 * Thực hiện focus cho input email
 */
const focusEmailInput = () => {
  emailInputRef.value?.focus();
};

const handleLogin = handleSubmit(async (payload) => {
  isLoading.value = true;
  const res = await UserAPI.login({
    email: payload.email,
    password: payload.password,
  });
  isLoading.value = false;

  // xử lý kết quả call api
  // trường hợp thất bại
  if (!res.status) {
    // show thông báo lỗi
    switch (res.error?.code) {
      case '204':
        message.error('Thông tin tài khoản không hợp lệ');
        break;
      default:
        message.error('Đăng nhập thất bại. Vui lòng thử lại sau.');
        break;
    }
    return;
  }

  // Thực hiện store token
  let accessToken = res.data?.accessToken;
  if (accessToken && window._storeAuth == ConfigGlobal.StoreAuthLocation.LocalStorage) {
    localStorage.setItem('accessToken', accessToken);
  }

  authStore.isAuthenticated = true;
  router.push({ name: RouterName.Home });
});
</script>

<style lang="scss" scoped>
.LoginView {
  height: 100vh;
  width: 100vw;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 40px;

  .Content {
    min-width: 400px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
    transition: 0.3s;
    border-radius: 12px;
    padding: 16px;

    .Title {
      font-size: 16px;
      font-weight: 700;
      padding-bottom: 8px;
      border-bottom: 1px solid #ccc;
    }

    .FormBox {
      padding-top: 18px;
      display: flex;
      flex-direction: column;
      gap: 16px;

      .FormGroup {
        .FormLabel {
          display: inline-block;
          margin-bottom: 4px;
        }
        .ButtonRegister {
          text-decoration: none;
          color: #000;
        }
        .ErrorMessage {
          margin-top: 6px;
          display: inline-block;
          color: red;
          font-size: 12px;
        }
      }
    }
  }
}
</style>
