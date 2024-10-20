<script setup lang="ts">
import { RouterLink, RouterView, useRouter, useRoute } from 'vue-router'
import { ref, onMounted, onBeforeRouteUpdate } from 'vue';

const username = ref('');
const loggedIn = ref(true);

async function fakeFetch(url: string) {
  await new Promise(resolve => { setTimeout(resolve, 100); });
  return {
    json: async () => ({ username: '乔某' }),
  };
}

onMounted(async () => {
  try {
    const response = await fakeFetch('https://api.example.com/username');
    const data = await response.json();
    if (!response.ok) {
      router.push('/login');
    }
    username.value = data.username;
    loggedIn.value = true;
  } catch (error) {
    console.error('Error loading username:', error);
    loggedIn.value = false;
  }
});

</script>

<template>
  <div class="layout-demo">
    <a-layout-header>
      <div class="menu-demo">
        <a-menu mode="horizontal" :default-selected-keys="['1']">
          <a-menu-item key="0" :style="{ padding: 0, marginRight: '38px' }" disabled>{{ username }}</a-menu-item>
          <RouterLink to="/"><a-menu-item key="1">主页</a-menu-item></RouterLink>
          <RouterLink to="/login"><a-menu-item key="2">登录</a-menu-item></RouterLink>
          <RouterLink to="/status"><a-menu-item key="3">状态</a-menu-item></RouterLink>
        </a-menu>
      </div>
    </a-layout-header>
    <a-layout-content>
      <RouterView :propa="1"/>
    </a-layout-content>
  </div>
</template>

<style scoped>
.menu-demo {
  box-sizing: border-box;
  width: 100%;
  padding: 40px;
  background-color: var(--color-neutral-2);
}

.layout-demo :deep(.arco-layout-header),
.layout-demo :deep(.arco-layout-content) {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: var(--color-white);
  font-size: 16px;
  font-stretch: condensed;
  text-align: center;
}

.layout-demo :deep(.arco-layout-content) {
  width: 100%;
  padding: 40px;
}
</style>
