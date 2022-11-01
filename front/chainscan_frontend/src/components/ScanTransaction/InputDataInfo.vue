<template lang="">
  <div>
    <div v-if="isCodeViewFlag">
      <div>
        <el-input
          style="font-family: Monaco"
          v-model="textarea"
          rows="8"
          placeholder="Please input"
          show-word-limit
          type="textarea"
          :readonly="true"
        />
      </div>
      <div style="margin-top: 10px">
        <el-dropdown>
          <el-button type="info">
            View Input As &nbsp; <el-icon><ArrowDownBold /></el-icon>
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item v-if="props.methodName !== ''" @click.native="setDefault()">
                Default View
              </el-dropdown-item>
              <el-dropdown-item @click.native="setOriginal()">Original</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>

        <el-button v-if="props.methodName !== ''" type="info" style="margin-left: 10px" @click="showDecode(false)">
          Decode Input Data &nbsp;<el-icon><Grid /></el-icon>
        </el-button>
      </div>
    </div>
    <div v-else style="margin-top: 10px; margin-bottom: 10px">
      <el-table :data="tableData" border style="width: 85%; font-size: 0.4rem; border-radius: 0.4rem">
        <el-table-column prop="index" label="#" />
        <el-table-column prop="type" label="Type" />
        <el-table-column prop="data" label="Data" width="500" />
      </el-table>
      <el-button style="margin-top: 10px" type="info" @click="showDecode(true)">
        <el-icon><Back /></el-icon>&nbsp; Switch Back
      </el-button>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ref, onMounted, reactive } from 'vue';
import { getParenthesesStr } from '../../script/utils';
import { ethers } from 'ethers';
import { ArrowDownBold, Grid, Back } from '@element-plus/icons-vue';

const props = defineProps({
  inputData: String,
  methodName: String,
});
const textarea = ref('');
const isCodeViewFlag = ref(true);
const tableData: any[] = reactive([]);

const decodeInput = (argList: string[], inputContent: string) => {
  tableData.length = 0;
  const abiCoder = new ethers.utils.AbiCoder();
  const decodeRes = abiCoder.decode(argList, inputContent);
  decodeRes.forEach((element, index) => {
    tableData.push({
      index: index,
      type: argList[index],
      data: element,
    });
  });
};

const showDecode = (flag: boolean) => {
  isCodeViewFlag.value = flag;
};

const setDefault = () => {
  const inputArgData: string = (props.inputData as string).slice(10);
  const functionName: string = 'Function: ' + props.methodName;
  const methodID: string = 'MethodID: ' + (props.inputData as string).slice(0, 10);
  let argData: string = '';
  let index: number = 0;
  for (let i = 0, len = inputArgData.length; i < len; i += 64) {
    // console.log('arg', (props.inputData as string).slice(i, i + 64));
    argData += '[' + index + '] :    ' + inputArgData.slice(i, i + 64) + '\n';
    index += 1;
  }
  textarea.value = functionName + '\n\n' + methodID + '\n' + argData;
};

const setOriginal = () => {
  textarea.value = props.inputData!;
};

onMounted(() => {
  if (props.methodName !== '') {
    const argList: string[] = getParenthesesStr(props.methodName!);
    if (argList.length !== 0) {
      decodeInput(argList, '0x' + (props.inputData as string).slice(10));
    }
    setDefault();
  } else {
    setOriginal();
  }
});
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>
