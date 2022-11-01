<template lang="">
  <div>
    <div>
      <div v-for="(page, index) of pages" :key="index" style="margin-top: 20px">
        <el-row :gutter="20">
          <el-col :span="4" v-for="(item, index) of page" :key="index">
            <el-card :body-style="{ padding: '10px' }">
              <div style="height: 100px; width: 100px; background-color: #598df6; border-radius: 0.35rem"></div>
              <div class="text-secondary">
                TokenID:
                <span>
                  <router-link :to="'/token/nfts/' + item.contract + '/' + item.tokenID + '/' + props.ercType">
                    {{ item.tokenID }}
                  </router-link>
                </span>
              </div>
              <div class="text-secondary">
                Owner:
                <span>
                  <router-link :to="'/address/' + item.owner">{{ item.owner.slice(0, 18) + '...' }}</router-link>
                </span>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>
      <div style="margin-top: 1%; display: flex; justify-content: center">
        <el-pagination
          small
          background
          :currentPage="currentPageIndex"
          :page-size="pageSizeNumber"
          :page-sizes="[10, 25, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="Total"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { Token } from '../../script/model/token';
import { GetTokenInventory } from '../../script/service/tokenService';
import { computed, ref, reactive } from 'vue';

const props = defineProps({
  address: {
    type: String,
    require: true,
  },
  ercType: {
    type: String,
    require: true,
  },
});

const currentPageIndex = ref(1);
const pageSizeNumber = ref(25);
const inventory: any = reactive([]);
const Total = ref(1);
const inventoryRes = await GetTokenInventory(
  props.address as string,
  props.ercType as string,
  currentPageIndex.value - 1,
  pageSizeNumber.value
);
Total.value = inventoryRes.data.total;
// console.log('inventoryRes', inventoryRes);
inventoryRes.data.items.forEach((element) => {
  // console.log('element', element);
  inventory.push(element);
});

const pages = computed({
  get() {
    const pagesList: any[] = [];
    (inventory as Token[]).forEach((item, index) => {
      const page = Math.floor(index / 6);
      if (!pagesList[page]) {
        pagesList[page] = [];
      }
      pagesList[page].push(item);
    });
    return pagesList;
  },
  set() {},
});

const handleSizeChange = async (pageSizeArg: number) => {
  inventory.length = 0;
  currentPageIndex.value = 1;
  pageSizeNumber.value = pageSizeArg;
  const res = await GetTokenInventory(
    props.address as string,
    props.ercType as string,
    currentPageIndex.value - 1,
    pageSizeNumber.value
  );
  res.data.items.forEach((element) => {
    inventory.push(element);
  });
};

const handleCurrentChange = async (currentPageArg: number) => {
  inventory.length = 0;
  currentPageIndex.value = currentPageArg;
  const res = await GetTokenInventory(
    props.address as string,
    props.ercType as string,
    currentPageIndex.value - 1,
    pageSizeNumber.value
  );
  res.data.items.forEach((element) => {
    inventory.push(element);
  });
};
</script>
<style lang="less" scoped>
.text-secondary {
  color: #77838f !important;
  font-size: 0.76562rem !important;
  margin-top: 10px;
  margin-bottom: 10px;
}
</style>
