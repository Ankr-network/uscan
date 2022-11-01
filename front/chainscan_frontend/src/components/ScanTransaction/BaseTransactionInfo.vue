<template lang="">
  <div>
    <h4>Status:</h4>
    <div v-if="baseInfo.status == 1" class="center-row">
      <el-icon color="green"><SuccessFilled /></el-icon> &nbsp; Success
    </div>
    <div v-if="baseInfo.status == 0" class="center-row">
      <el-icon color="red"><Failed /></el-icon> &nbsp; Fail
    </div>
    <div v-if="baseInfo.status == 3" class="center-row">
      <el-icon><VideoPause /></el-icon> &nbsp; Pending
    </div>
    <el-divider />
    <h4>Transaction Fee:</h4>
    {{ ethers.utils.formatEther(baseInfo.gasLimit * baseInfo.gasUsed) }} Eth
    <el-divider />
    <h4>Gas Info:</h4>
    {{ baseInfo.gasUsed }} Used From {{ baseInfo.gasLimit }} GasLimit
    <el-divider />
    <h4>Nonce:</h4>
    {{ baseInfo.nonce }}
    <el-divider />

    <router-link class="center-row" :to="'/tx/' + txHash">
      See more details &nbsp; <el-icon size="large"><Link /></el-icon>
    </router-link>
  </div>
</template>
<script lang="ts" setup>
import { GetBaseTxByHash } from '../../script/service/transactionService';
import { Link, SuccessFilled, Failed, VideoPause } from '@element-plus/icons-vue';
import { ethers } from 'ethers';

const props = defineProps({
  txHash: String,
});

const res = await GetBaseTxByHash(props.txHash as string);
const baseInfo = res.data;
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>
