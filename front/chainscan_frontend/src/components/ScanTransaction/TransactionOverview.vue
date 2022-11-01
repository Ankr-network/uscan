<template lang="">
  <div>
    <el-table class="table-border" :data="overviews" empty-text="loading..." :row-style="{ height: '50px' }">
      <el-table-column width="240">
        <template v-slot:default="scope">
          <div class="center-row">
            <el-tooltip effect="dark" placement="top">
              <template #content>
                <div style="max-width: 250px">{{ scope.row.parameterExplain }}</div>
              </template>
              <el-icon><QuestionFilled /></el-icon>
            </el-tooltip>
            &nbsp;{{ scope.row.parameterDisplay }}
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="parameterValue">
        <template v-slot:default="scope">
          <div v-if="scope.row.parameterName == 'blockNumber'" style="font-weight: 900">
            <router-link :to="'/block/' + parseInt(scope.row.parameterValue)">
              {{ parseInt(scope.row.parameterValue) }}</router-link
            >
          </div>
          <div class="center-row" v-else-if="scope.row.parameterName == 'createTime'">
            <el-icon><Clock /></el-icon>&nbsp;{{ getAge(scope.row.parameterValue) }}&nbsp; ({{
              new Date(scope.row.parameterValue * 1000).toUTCString()
            }})
          </div>
          <div class="center-row" v-else-if="scope.row.parameterName == 'hash'">
            {{ scope.row.parameterValue }} &nbsp;
            <copy-icon :text="scope.row.parameterValue"></copy-icon>
          </div>
          <div class="center-row" v-else-if="scope.row.parameterName == 'from'">
            <router-link :to="'/address/' + scope.row.parameterValue.from">
              {{ scope.row.parameterValue.from }}
            </router-link>
            &nbsp;
            <copy-icon :text="scope.row.parameterValue.from"></copy-icon>
          </div>
          <div v-else-if="scope.row.parameterName == 'to'">
            <div class="center-row" v-if="scope.row.parameterValue.to == ''">
              Contract &nbsp;
              <router-link :to="'/address/' + scope.row.parameterValue.contractAddress">{{
                scope.row.parameterValue.contractAddress
              }}</router-link>
              Created &nbsp;
              <copy-icon :text="scope.row.parameterValue.contractAddress"></copy-icon>
            </div>
            <div class="center-row" v-else-if="scope.row.parameterValue.toCode != ''">
              Contract &nbsp;
              <router-link :to="'/address/' + scope.row.parameterValue.to">
                {{ scope.row.parameterValue.to }} &nbsp; {{ scope.row.parameterValue.toName }}
              </router-link>
              &nbsp;
              <copy-icon :text="scope.row.parameterValue.to"></copy-icon>
            </div>
            <div class="center-row" v-else>
              <router-link :to="'/address/' + scope.row.parameterValue.to">{{
                scope.row.parameterValue.to
              }}</router-link>
              &nbsp;
              <copy-icon :text="scope.row.parameterValue.to"></copy-icon>
            </div>
          </div>
          <div v-else-if="scope.row.parameterName == 'value'">
            {{ ethers.utils.formatEther(scope.row.parameterValue) }} Eth
          </div>
          <div v-else-if="scope.row.parameterName == 'tokensTransferred'">
            <div :class="scope.row.parameterValue.length >= 3 ? 'rolling' : ''">
              <div v-for="(trans, index) in scope.row.parameterValue" :key="index">
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
          </div>
          <div v-else-if="scope.row.parameterName == 'gas'">
            {{ scope.row.parameterValue.gasUsed }} Gwei | {{ scope.row.parameterValue.gasLimit }} Gwei({{
              scope.row.parameterValue.percent
            }})
          </div>
          <div v-else-if="scope.row.parameterName == 'gasPrice'">
            {{ ethers.utils.formatEther(scope.row.parameterValue) }} Eth
          </div>
          <div v-else-if="scope.row.parameterName == 'maxPriorityFeePerGas'">
            Base: {{ ethers.utils.formatEther(scope.row.parameterValue.baseFeePerGas, 'gwei') }} Gwei | Max:
            {{ ethers.utils.formatEther(scope.row.parameterValue.maxFeePerGas, 'gwei') }} Gwei | MaxPriority:
            {{ ethers.utils.formatEther(scope.row.parameterValue.maxPriorityFeePerGas, 'gwei') }} Gwei
          </div>
          <div v-else-if="scope.row.parameterName == 'input'">
            <!-- <textarea
              class="byte-codes-text"
              style="margin: 0px"
              rows="6"
              v-model="scope.row.parameterValue.inputContent"
            >
            </textarea> -->
            <Input-Data-Info
              :inputData="scope.row.parameterValue.inputContent"
              :methodName="scope.row.parameterValue.methodName"
            ></Input-Data-Info>
          </div>
          <div v-else-if="scope.row.parameterName == 'status'">
            <div v-if="scope.row.parameterValue.status == 1">
              <div class="success-status">
                <el-icon color="green"><SuccessFilled /></el-icon> &nbsp; Success
              </div>
            </div>
            <div v-if="scope.row.parameterValue.status == 0">
              <div class="fail-status">
                <el-icon color="red"><Failed /></el-icon> &nbsp; Fail
              </div>
              Fail reason : {{ scope.row.parameterValue.errorMsg }}
            </div>
            <div v-if="scope.row.parameterValue.status == 3">
              <div class="pending-status">
                <el-icon><VideoPause /></el-icon> &nbsp; Pending
              </div>
            </div>
          </div>

          <div v-else>{{ scope.row.parameterValue }}</div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script lang="ts" setup>
import { GetTxByHash } from '../../script/service/transactionService';
import { getTxOverviews } from '../../script/model/transaction';
import { QuestionFilled, Clock, SuccessFilled, Failed, VideoPause } from '@element-plus/icons-vue';
import { getAge } from '../../script/utils';
import { reactive, watch } from 'vue';
import { ethers } from 'ethers';
import { useRoute } from 'vue-router';
import { getTitle } from '../../script/utils';

document.title = 'Transaction overview | The ' + getTitle + ' Explorer';

const props = defineProps({
  txHash: String,
});

const route = useRoute();

const overviews: any[] = reactive([]);

const initData = async (txHash: String) => {
  overviews.length = 0;

  const res = await GetTxByHash(txHash as string);
  getTxOverviews(res.data).forEach((element) => {
    overviews.push(element);
  });
};

initData(props.txHash as string);

watch(
  () => route.params,
  async (val) => {
    // console.log('watchsssssss', val.txHash);
    if (val.txHash) {
      initData(val.txHash as unknown as string);
    }
  }
);
</script>
<style lang="less" scoped>
@import '../../css/style.css';
.rolling {
  height: 100px;
  overflow: auto;
}
</style>
