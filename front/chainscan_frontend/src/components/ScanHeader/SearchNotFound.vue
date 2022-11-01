<template lang="">
  <div style="display: flex; justify-content: center; text-align: center; height: 360px">
    <div>
      <h1 style="color: #3498db">Not Found</h1>
      <el-button type="primary" size="large" @click="moveToHome">Back Home</el-button>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { useRouter } from 'vue-router';
import { getTitle } from '../../script/utils';
const router = useRouter();

document.title = 'Not Found | The ' + getTitle + ' Explorer';
const moveToHome = () => {
  router.push('/');
};
</script>
<style lang=""></style>
