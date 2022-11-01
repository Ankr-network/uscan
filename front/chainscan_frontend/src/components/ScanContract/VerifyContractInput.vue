<template lang="">
  <div style="display: flex; flex-direction: column">
    <div style="text-align: center">
      <h3 class="content-sub">Verify & Publish Contract Source Code</h3>
      <p style="color: #8c98a4; ont-weight: 700; font-size: 80%; font-weight: bold">
        COMPILER TYPE AND VERSION SELECTION
      </p>
    </div>
    <el-divider />
    <div style="display: flex; justify-content: center">
      <p style="color: #77838f; width: 80%">
        &nbsp; &nbsp; Source code verification provides transparency for users interacting with smart contracts. By
        uploading the source code, Etherscan will match the compiled code with that on the blockchain. Just like
        contracts, a "smart contract" should provide end users with more information on what they are "digitally
        signing" for and give users an opportunity to audit the code to independently verify that it actually does what
        it is supposed to do.
      </p>
    </div>
    <br />
    <div style="display: flex; justify-content: center">
      <div class="content-sub1">
        <div style="width: 700px; font-size: 0.875rem; margin-bottom: 10px">
          <p>Please enter the Contract Address you would like to verify</p>
          <el-input v-model="addressInput" size="large" placeholder="0x..." clearable />
          <div style="color: red" v-if="addressInputRequire"><p>Required</p></div>
        </div>
      </div>
    </div>
    <div style="display: flex; justify-content: center">
      <div class="content-sub1">
        <p>Please select Compiler Type</p>
        <el-select v-model="compilerType" placeholder="Select" size="large" style="width: 100%">
          <el-option v-for="item in compilerTypeOptions" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
        <div style="color: red" v-if="compilerTypeRequired"><p>Required</p></div>
      </div>
    </div>
    <div style="display: flex; justify-content: center">
      <div class="content-sub1">
        <p>Please select Compiler Version</p>
        <el-select v-model="compilerVersion" placeholder="Select" size="large" style="width: 100%">
          <el-option
            v-for="item in compilerVersionOptions"
            :key="item.value.name"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <div style="color: red" v-if="compilerVersionRequired"><p>Required</p></div>
      </div>
    </div>
    <div style="display: flex; justify-content: center">
      <div class="content-sub1">
        <p>Please select Open Source License Type</p>
        <el-select v-model="license" placeholder="Select" size="large" style="width: 100%">
          <el-option v-for="item in licenseOptions" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
        <div style="color: red" v-if="licenseRequired"><p>Required</p></div>
      </div>
    </div>
    <div style="display: flex; justify-content: center">
      <div style="display: flex; justify-content: center; margin-top: 10px">
        <el-button type="primary" size="large" @click="moveToSubmit">Continue</el-button>
        <el-button type="info" size="large" @click="reset">Reset</el-button>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ref } from 'vue';
import { GetVerifyContractMetadata } from '../../script/service/contractService';
import { useRoute, useRouter } from 'vue-router';
import { getTitle } from '../../script/utils';

const router = useRouter();
const route = useRoute();

document.title = 'Verify & Publish Contract Source Code | The ' + getTitle + ' Explorer';

interface Option {
  value: any;
  label: string;
}

const addressInput = ref('');
const addressInputRequire = ref(false);

const compilerType = ref('');
const compilerTypeRequired = ref(false);
const compilerTypeOptions: Option[] = [];

const compilerVersion = ref('');
const compilerVersionRequired = ref(false);
const compilerVersionOptions: Option[] = [];

const license = ref('');
const licenseRequired = ref(false);
const licenseOptions: Option[] = [];

if (route.query.a) {
  addressInput.value = route.query.a as string;
}

const metadataRes = await GetVerifyContractMetadata();
// console.log(metadataRes);
compilerTypeOptions.push(
  { value: 'solidity-single-file', label: 'Solidity (Single file)' } as Option,
  { value: 'solidity-standard-json-input', label: 'Solidity (Standard-Json-Input)' } as Option
);
const compilerVersionOptionsNameMap = new Map();
const compilerVersionOptionsFileMap = new Map();

metadataRes.data.compilerVersions.forEach((element) => {
  // console.log(element);
  compilerVersionOptionsNameMap.set(element.id, element.name);
  compilerVersionOptionsFileMap.set(element.id, element.fileName);
  compilerVersionOptions.push({
    value: element.id,
    label: element.name,
  } as Option);
});

metadataRes.data.licenseTypes.forEach((element) => {
  licenseOptions.push({
    value: element.id,
    label: element.name,
  } as Option);
});

const reset = () => {
  addressInput.value = '';
  addressInputRequire.value = false;

  compilerType.value = '';
  compilerTypeRequired.value = false;

  compilerVersion.value = '';
  compilerVersionRequired.value = false;

  license.value = '';
  licenseRequired.value = false;
};

const moveToSubmit = () => {
  addressInput.value === '' ? (addressInputRequire.value = true) : (addressInputRequire.value = false);
  compilerType.value === '' ? (compilerTypeRequired.value = true) : (compilerTypeRequired.value = false);
  compilerVersion.value === '' ? (compilerVersionRequired.value = true) : (compilerVersionRequired.value = false);
  license.value === '' ? (licenseRequired.value = true) : (licenseRequired.value = false);
  if (addressInput.value !== '' && compilerType.value !== '' && compilerVersion.value !== '' && license.value !== '') {
    // console.log('will push');
    // console.log('compilerVersion', compilerVersionOptionsNameMap.get(compilerVersion.value));
    // console.log('compilerVersion', compilerVersionOptionsFileMap.get(compilerVersion.value));
    const url =
      '/verifyContract/submit?a=' +
      addressInput.value +
      '&ct=' +
      compilerType.value +
      '&cv=' +
      compilerVersionOptionsNameMap.get(compilerVersion.value) +
      '&cf=' +
      compilerVersionOptionsFileMap.get(compilerVersion.value) +
      '&lictype=' +
      license.value;
    router.push(url);
  }
};
</script>
<style lang="less" scoped>
.content-sub {
  font-size: 1.53125rem;
  font-weight: 400;
  color: #4a4f55;
}
.content-sub1 {
  width: 700px;
  font-size: 0.875rem;
  margin-bottom: 10px;
}
</style>
