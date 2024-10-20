<template>
  <a-form :model="form" layout="vertical" :style="{ width: '500px', margin: '0 auto' }" @submit="handleSubmit">
    <a-form-item field="key" tooltip="请输入你的身份证号" label="身份证号">
      <a-input v-model="form.key" placeholder="请输入你的身份证号" />
    </a-form-item>
    <a-form-item>
      <a-button html-type="submit">提交</a-button>
    </a-form-item>
  </a-form>
</template>

<script>
import { reactive } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

export default {
  setup() {
    const form = reactive({
      key: '',
      post: '',
      isRead: false,
    });
    const handleSubmit = (data) => {
      try {
        const response = await fetch('/api/', {
          method: "HEAD",
          cache: "no-cache",
          headers: {
            "key": form.key
          }
        });
        if (!response.ok) {
          alert("身份证号登录失败");
          router.push('/login');
        }
      } catch (error) {
        alert("身份证号登录失败");
        router.push('/login');
      }
      console.log(data);
    };

    return {
      form,
      handleSubmit,
    };
  },
};
</script>