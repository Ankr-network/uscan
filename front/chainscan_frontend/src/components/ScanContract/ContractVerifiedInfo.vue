<template lang="">
  <div>
    <el-row style="margin-top: 0.5%">
      <el-button :type="codeFlag ? 'primary' : 'info'" @click="moveToCode">Code</el-button>
      <el-button :type="readFlag ? 'primary' : 'info'" @click="moveToRead">Read Contract</el-button>
      <el-button :type="writeFlag ? 'primary' : 'info'" @click="moveToWrite">Write Contract</el-button>
    </el-row>
    <br />
    <component :is="comName" :contractAddress="props.contractAddress" :contractInfo="props.contractInfo"></component>
  </div>
</template>
<script lang="ts" setup>
import { shallowRef, ref } from 'vue';
import { ContractContent } from '../../script/model/contract';
import codeContractVue from './ContractInfo/codeContract.vue';
import readContractVue from './ContractInfo/readContract.vue';
import writeContractVue from './ContractInfo/writeContract.vue';

const props = defineProps({
  contractAddress: String,
  contractInfo: {
    type: Object as () => ContractContent,
  },
});

const comName = shallowRef(codeContractVue);
const codeFlag = ref(true);
const readFlag = ref(false);
const writeFlag = ref(false);

const moveToCode = () => {
  comName.value = codeContractVue;
  codeFlag.value = true;
  readFlag.value = false;
  writeFlag.value = false;
};

const moveToRead = () => {
  comName.value = readContractVue;
  readFlag.value = true;
  codeFlag.value = false;
  writeFlag.value = false;
  console.log('codeFlag:', codeFlag, 'readFlag:', readFlag, 'writeFlag:', writeFlag);
};

const moveToWrite = () => {
  comName.value = writeContractVue;
  codeFlag.value = false;
  readFlag.value = false;
  writeFlag.value = true;
};
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>
