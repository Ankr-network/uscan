<template lang="">
  <div :class="isRolling ? 'rolling' : ''">
    <div v-for="(trans, index) in tokensTransferData" :key="index">
      <div class="center-row">
        <div style="font-weight: bold">From</div>
        &nbsp;&nbsp;&nbsp;
        <router-link :to="'/address/' + trans.from">{{ trans.from.slice(0, 18) + '...' }}</router-link>
        &nbsp;&nbsp;&nbsp;
        <div style="font-weight: bold">To</div>
        &nbsp;&nbsp;&nbsp;
        <router-link :to="'/address/' + trans.to">{{ trans.to.slice(0, 18) + '...' }}</router-link>
        &nbsp;&nbsp;&nbsp;
        <router-link :to="'/address/' + trans.address">
          <div v-if="trans.addressName">{{ trans.addressName }}</div>
          <div v-else>
            {{ trans.address.slice(0, 18) + '...' }}
          </div>
        </router-link>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { TokensTransferred } from '../../script/model/transaction';
import { ref } from 'vue';

const isRolling = ref(false);

const props = defineProps({
  tokensTransferData: {
    type: Array as () => Array<TokensTransferred>,
    require: true,
  },
});

// console.log('tokensTransferData init', props.tokensTransferData);

if (props.tokensTransferData?.length != 0) {
  // console.log('tokensTransferData', props.tokensTransferData);
  if ((props.tokensTransferData as TokensTransferred[]).length >= 3) {
    isRolling.value = true;
  }
}
</script>
<style lang="less" scoped>
@import '../../css/style.css';
.rolling {
  height: 100px;
  overflow: auto;
}
</style>
