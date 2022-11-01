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
          <div v-if="scope.row.parameterName == 'number'" style="font-weight: 900">
            {{ scope.row.parameterValue }}
            &nbsp;
            <el-button-group>
              <el-button
                type="primary"
                size="small"
                style="border: 0"
                plain
                @click="moveToBlock(parseInt(blockNumber) - 1)"
              >
                <el-icon><ArrowLeftBold /></el-icon>
              </el-button>
              <el-button
                type="primary"
                size="small"
                style="border: 0"
                plain
                @click="moveToBlock(parseInt(blockNumber) + 1)"
              >
                <el-icon><ArrowRightBold /></el-icon>
              </el-button>
            </el-button-group>
          </div>
          <div class="center-row" v-else-if="scope.row.parameterName == 'createdTime'">
            <el-icon><Clock /></el-icon>&nbsp;{{ getAge(scope.row.parameterValue) }}
          </div>
          <div v-else-if="scope.row.parameterName == 'transactionsTotal'">
            <div v-if="scope.row.parameterValue == 0">0 transaction in this block</div>
            <div v-else>
              <el-tooltip class="box-item" effect="dark" content="Click to view Transactions" placement="right">
                <el-button type="primary" plain size="small" @click="moveToTxs" style="border: 0">
                  {{ scope.row.parameterValue }} transactions
                </el-button>
                &nbsp;in this block
              </el-tooltip>
            </div>
          </div>
          <div v-else-if="scope.row.parameterName == 'miner'">
            <router-link :to="'/address/' + scope.row.parameterValue"> {{ scope.row.parameterValue }}</router-link>
          </div>
          <div v-else>{{ scope.row.parameterValue }}</div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script lang="ts" setup>
import { GetBlockByNumber } from '../../script/service/blockService';
import { getBlockOverviews } from '../../script/model/block';
import { QuestionFilled, ArrowLeftBold, ArrowRightBold, Clock } from '@element-plus/icons-vue';
import { getAge } from '../../script/utils';
import { useRouter } from 'vue-router';
import { watchEffect, reactive } from 'vue';

const props = defineProps({
  blockNumber: Number,
});

const overviews: any[] = reactive([]);

const router = useRouter();

const moveToBlock = function (blockNumber: number) {
  router.push('/block/' + blockNumber);
};

const moveToTxs = function () {
  router.push('/txs/all?block=' + props.blockNumber);
};

watchEffect(async () => {
  overviews.length = 0;
  const res = await GetBlockByNumber(props.blockNumber as number);
  getBlockOverviews(res.data).forEach((element) => overviews.push(element));
  // console.log('overviews', overviews);
});
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>
