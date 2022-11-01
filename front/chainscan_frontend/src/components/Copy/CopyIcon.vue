<template lang="">
  <div class="copy-content">
    <el-tooltip placement="right" :visible="visible">
      <template #content>{{ copyTip }} </template>
      <el-icon
        class="copy-icon"
        @click="copyContent"
        @mouseenter="visible = true"
        @mouseleave="
          visible = false;
          leave();
        "
        ><DocumentCopy
      /></el-icon>
    </el-tooltip>
  </div>
</template>
<script lang="ts" setup>
import { ref } from 'vue';
import useClipboard from 'vue-clipboard3';
import { DocumentCopy } from '@element-plus/icons-vue';

const props = defineProps({
  text: String,
});

const copyTip = ref('Copy to clipboard');
const visible = ref(false);
const { toClipboard } = useClipboard();
const copyContent = async () => {
  copyTip.value = 'Copy Success';
  try {
    await toClipboard(props.text as string);
  } catch (e) {
    console.error(e);
  }
};
const leave = () => {
  copyTip.value = 'Copy to clipboard';
};
</script>
<style lang="less" scoped>
.copy-content {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
.copy-icon {
  color: #9ba2aa;
  cursor: pointer;
}
</style>
