<template lang="">
  <div class="code-content">
    <codemirror
      v-model="code"
      placeholder="Code here..."
      :style="{ height: '300px' }"
      :autofocus="true"
      :indent-with-tab="true"
      :tabSize="2"
      :extensions="extensions"
    />
  </div>
</template>
<script lang="ts" setup>
import { Codemirror } from 'vue-codemirror';
import { javascript } from '@codemirror/lang-javascript';
const props = defineProps({
  codeData: String,
});
// eslint-disable-next-line vue/no-setup-props-destructure
const code = props.codeData;
const extensions = [javascript()];
</script>
<style lang="less" scoped>
.code-content {
  font-size: 13px;
  border: 1px solid #e7eaf3;
  border-radius: 0.35rem;
}
</style>
