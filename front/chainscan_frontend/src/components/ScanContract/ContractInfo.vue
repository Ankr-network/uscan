<template lang="">
  <div>
    <div class="center-row" style="font-size: 14px">
      <el-icon size="large"><InfoFilled /></el-icon> &nbsp; Are you the contract creator?&nbsp;
      <router-link :to="'/verifyContract/input?a=' + props.contractAddress" style="font-size: 15px; font-weight: bold">
        Verify and Publish
      </router-link>
      &nbsp;your contract source code today!
    </div>
    <br />

    <textarea class="byte-codes-text" rows="15" v-model="code"> </textarea>
  </div>
</template>
<script lang="ts" setup>
import { InfoFilled } from '@element-plus/icons-vue';

const props = defineProps({
  contractAddress: String,
  codeContent: String,
});

// eslint-disable-next-line vue/no-setup-props-destructure
const code = props.codeContent;

// console.log(props.codeContent);
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>
