<template lang="">
  <div>
    <div class="center-row">
      <el-icon :color="'green'"><CircleCheckFilled /></el-icon> &nbsp;
      <h4>Contract Source Code Verified</h4>
    </div>
    <div class="code-content-overview">
      <div style="width: 50%">
        <el-row>
          <el-col :span="8">Contract Name:</el-col>
          <el-col :span="16" class="bolder">{{ props.contractInfo.contractName }}</el-col>
        </el-row>
        <el-divider />
        <el-row>
          <el-col :span="8">Contract Version:</el-col>
          <el-col :span="16" class="bolder">{{ props.contractInfo.compilerVersion }}</el-col>
        </el-row>
      </div>
      <div style="width: 50%; margin-left: 3%">
        <el-row>
          <el-col :span="8">Optimization Enabled:</el-col>
          <el-col :span="16" class="bolder"
            >{{ props.contractInfo.runs === 0 ? 'No' : 'Yes' }} with {{ props.contractInfo.optimization }} runs</el-col
          >
        </el-row>
        <el-divider />
        <el-row>
          <el-col :span="8">Other Settings:</el-col>
          <el-col :span="16" class="bolder">
            default evmVersion, {{ licenseTypeMap.get(props.contractInfo.licenseType) }}
          </el-col>
        </el-row>
      </div>
    </div>

    <div style="margin-top: 20px">
      <div v-for="(code, index) in contractSourceList" :key="index">
        <div class="center-row">
          <el-icon><Document /></el-icon> &nbsp;
          <h4>Contract Source Code (Solidity)</h4>
        </div>
        <p style="font-size: 13px">File {{ index + 1 }} of {{ contractSourceList.length }} : {{ code.filename }}</p>
        <code-view :codeData="code.codeContent"></code-view>
      </div>
    </div>

    <div>
      <div class="center-row">
        <el-icon><Document /></el-icon> &nbsp;
        <h4>Contract ABI</h4>
      </div>
      <textarea class="byte-codes-text" style="margin: 0px" rows="10" v-model="abi"> </textarea>
    </div>

    <div>
      <div class="center-row">
        <el-icon><Document /></el-icon> &nbsp;
        <h4>Contract Creation Code</h4>
      </div>
      <textarea class="byte-codes-text" style="margin: 0px" rows="10" v-model="creationCode"> </textarea>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ref, watchEffect } from 'vue';
import { ContractContent } from '../../../script/model/contract';
import { CircleCheckFilled, Document } from '@element-plus/icons-vue';
import { GetVerifyContractMetadata } from '../../../script/service/contractService';

const props = defineProps({
  contractAddress: String,
  contractInfo: {
    type: Object as () => ContractContent,
  },
});
const contractSourceList = ref([] as any[]);
const abi = ref('');
const creationCode = ref('');

const metadataRes = await GetVerifyContractMetadata();
const licenseTypeMap = new Map();
metadataRes.data.licenseTypes.forEach((element) => {
  licenseTypeMap.set(element.id, element.name);
});

watchEffect(async () => {
  if (Object.keys(props.contractInfo as ContractContent).length !== 0) {
    // console.log('code', props.contractInfo);
    abi.value = props.contractInfo?.abi as string;
    creationCode.value = props.contractInfo?.object as string;
    Object.keys(props.contractInfo?.metadata).forEach((key) => {
      contractSourceList.value.push({
        filename: key,
        codeContent: props.contractInfo?.metadata[key],
      });
    });
    // console.log('contractSourceList', contractSourceList);
  }
});
</script>
<style lang="less" scoped>
@import '../../../css/style.css';
</style>
