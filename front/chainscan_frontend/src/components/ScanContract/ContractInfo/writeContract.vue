<template lang="">
  <div>
    <div class="center-row" style="font-size: 13px">
      <el-icon><Document /></el-icon>&nbsp;
      <p>Write Contract Information</p>
    </div>
    <p>Please make sure the metamask is installed.</p>
    <div v-for="(functionObject, index) in functionObjectList" :key="index">
      <el-collapse class="method-content" v-model="activeNames">
        <el-collapse-item class="method-object" :title="index + 1 + '.' + functionObject.name" :name="index">
          <div style="padding-right: 0.5rem; padding-left: 0.5rem">
            <div v-for="(input, inputIndex) in functionObject.inputsArg" :key="inputIndex">
              <div style="margin-top: 0.8rem">
                <div style="font-size: 9px">{{ input.name + '(' + input.internalType + ')' }}</div>
                <el-input v-model="input.arg" :placeholder="input.name + '(' + input.internalType + ')'" />
              </div>
            </div>
            <div style="margin-top: 0.8rem" v-if="functionObject.inputsArg.length != 0">
              <el-button type="info" plain @click="write(functionObject)">Write</el-button>
            </div>
            <div style="margin-top: 0.8rem">
              <div v-for="(output, index) in functionObject.outputsRes" :key="index">
                <div class="method-output">
                  <div>{{ output.arg }}</div>
                  &nbsp;&nbsp;
                  <div class="arg-type">{{ output.internalType }}</div>
                </div>
              </div>
            </div>
            <div v-if="functionObject.resMsg != ''">
              <p>{{ functionObject.resMsg }}</p>
            </div>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ref, reactive } from 'vue';
import { ContractContent } from '../../../script/model/contract';
import { Document } from '@element-plus/icons-vue';
import { ethers } from 'ethers';

const props = defineProps({
  contractAddress: String,
  contractInfo: {
    type: Object as () => ContractContent,
  },
});

const functionObjectList = reactive([] as any[]);

const activeNames = ref([] as number[]);

const functionResMap = new Map();

const abi = (props.contractInfo as unknown as ContractContent).abi as string;

const initData = () => {
  if (Object.keys(props.contractInfo as ContractContent).length !== 0) {
    activeNames.value = [];
    let index = 0;
    JSON.parse(abi).forEach((elementFunc: any) => {
      if (elementFunc.stateMutability == 'nonpayable' && elementFunc.type == 'function') {
        // console.log('elementFunc', elementFunc);
        const inputsArg: any[] = [];
        if (elementFunc.inputs) {
          if (elementFunc.inputs.length != 0) {
            elementFunc.inputs.forEach((element: any) => {
              inputsArg.push({
                arg: '',
                name: element.name,
                internalType: element.type,
              });
            });
          }
        }
        const outputsRes: any[] = [];
        if (elementFunc.outputs) {
          if (elementFunc.outputs.length != 0) {
            elementFunc.outputs.forEach((element: any) => {
              outputsRes.push({
                arg: '',
                name: element.name,
                internalType: element.type,
              });
            });
          }
        }
        const functionObject = {
          name: elementFunc.name,
          inputs: elementFunc.inputs,
          inputsArg: inputsArg,
          outputs: elementFunc.outputs,
          outputsRes: outputsRes,
          resMsg: '',
        };
        functionObjectList.push(functionObject);
        functionResMap.set(index, functionObject);
        index += 1;
      }
    });
    activeNames.value = Array.from(new Array(functionObjectList.length).keys());
  }
};

initData();

const write = async (functionObject: any) => {
  // console.log(functionObject);
  if (JSON.parse(abi).length != 0) {
    const provider = new ethers.providers.Web3Provider((window as any).ethereum);
    if ((window as any).ethereum._state.accounts.length == 0) {
      await provider.send('eth_requestAccounts', []);
    }
    const signer = provider.getSigner();
    const contract = new ethers.Contract(props.contractAddress as string, abi, provider);
    const contractWithSigner = contract.connect(signer);
    Reflect.ownKeys(contractWithSigner.functions).forEach(async function (key) {
      if (key == functionObject.name) {
        const requestArgList: any[] = [];
        functionObject.inputsArg.forEach((element: any) => {
          requestArgList.push(element.arg);
        });
        try {
          const tx = await contractWithSigner.functions[key as string](...requestArgList);
          functionObject.resMsg = 'Write succeeded, please wait for confirmation';
          console.log(tx);
        } catch (err: any) {
          // console.log("err", err.reason);
          functionObject.resMsg = err.reason;
        }
      }
    });
  }
};
</script>
<style lang="less" scoped>
@import '../../../css/style.css';
</style>
