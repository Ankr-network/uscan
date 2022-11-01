<template lang="">
  <div>
    <div class="sub-info">
      <h3 class="h3-title">Verify & Publish Contract Source Code</h3>
      <p class="subtitle">Compiler Type: SINGLE FILE / CONCATENANTED METHOD</p>
    </div>
    <div class="content-sub">
      <el-tabs style="width: 100%">
        <el-tab-pane>
          <template #label>
            <p class="subtitle1">Contract Source Code</p>
          </template>
          <div style="margin: 10px">
            <el-row :gutter="10">
              <el-col :span="7">
                <div class="title-input">
                  <p>Contract Name</p>
                  <el-input v-model="contractName" size="large" />
                  <div style="color: red" v-if="contractNameRequired"><p>Required</p></div>
                </div>
              </el-col>
              <el-col :span="7">
                <div class="title-input">
                  <p>Contract Address</p>
                  <el-input v-model="contractAddress" size="large" :readonly="true" />
                </div>
              </el-col>
              <el-col :span="7">
                <div class="title-input">
                  <p>Compiler</p>
                  <el-select v-model="compilerVersion" size="large" style="width: 100%" disabled> </el-select>
                </div>
              </el-col>
              <el-col :span="3">
                <div v-if="compilerType == 'solidity-single-file'" class="title-input">
                  <p>Optimization</p>
                  <el-select v-model="optimizationValue" placeholder="Select" size="large" style="width: 100%">
                    <el-option :key="1" :label="'Yes'" :value="1" />
                    <el-option :key="2" :label="'No'" :value="0" />
                  </el-select>
                </div>
              </el-col>
            </el-row>
          </div>
          <div v-if="route.query.ct == 'solidity-single-file'" style="margin: 10px">
            <h4>Enter the Solidity Contract Code below</h4>
            <textarea
              class="byte-codes-text"
              rows="10"
              style="margin-top: 0px; background-color: white"
              v-model="sourceCode"
            >
            </textarea>
          </div>
          <div v-else style="margin: 10px">
            <h4>Please select the Standard-Input-Json (*.json) file to upload</h4>
            <div>
              <div>
                <div><p>Click button select file</p></div>
                <div style="width: 30%">
                  <el-upload
                    :auto-upload="false"
                    action="Fake Action"
                    accept=".json"
                    :on-change="handleUploadChange"
                    :file-list="fileList"
                  >
                    <el-button>Select a file</el-button>
                  </el-upload>
                  <div style="color: red" v-if="fileRequired"><p>Required</p></div>
                </div>
                <div v-if="fileList.length == 0">
                  <p>No file selected</p>
                </div>
              </div>
            </div>
          </div>

          <div style="margin: 10px; margin-top: 30px">
            <el-collapse>
              <el-collapse-item>
                <template #title>
                  <p>Misc Settings</p>
                  <p style="color: #77838f">(Runs, EvmVersion & License Type settings)</p>
                </template>
                <div>
                  <el-row :gutter="10">
                    <el-col :span="8">
                      <div v-if="optimizationValue == 1" class="title-input">
                        <p>Runs</p>
                        <el-input v-model.number="runsValue" size="large" oninput="value=value.replace(/[^0-9]/g,'')" />
                      </div>
                    </el-col>
                    <el-col :span="8">
                      <div v-if="route.query.ct == 'solidity-single-file'" class="title-input">
                        <p>EVM Version to target</p>
                        <el-select v-model="evmVersionValue" size="large" style="width: 100%" disabled> </el-select>
                      </div>
                    </el-col>
                    <el-col :span="8">
                      <div class="title-input">
                        <p>LicenseType</p>
                        <el-select v-model="licenseValue" placeholder="Select" size="large" style="width: 100%">
                          <el-option
                            v-for="item in licenseOptions"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                          />
                        </el-select>
                      </div>
                    </el-col>
                  </el-row>
                </div>
              </el-collapse-item>
            </el-collapse>
          </div>
          <div class="button-content">
            <el-button type="primary" size="large" @click="submit" :loading="submitLoading">
              Verify and Publish
            </el-button>
            <el-button type="info" size="large" @click="reset">Reset</el-button>
            <el-button type="info" size="large" @click="returnMain">Return to main</el-button>
          </div>
          <div>
            <div v-if="verifyContractStatus == -1" class="submit-result">
              <div class="subtitle" v-if="submittedStatus >= 200 && submittedStatus <= 300">
                <p>Submitted, please wait for verification</p>
              </div>
              <div class="subtitle" v-else-if="submittedStatus == -1">
                <p>Something wrong, {{ submittedError }}</p>
              </div>
            </div>
            <div v-else class="submit-result">
              <div class="subtitle" v-if="verifyContractStatus == 1">Verify success!</div>
              <div class="subtitle" v-else-if="verifyContractStatus == 2">
                Verify fail!
                <router-link :to="'/address/' + contractAddress">
                  click this back contract {{ contractAddress }} page
                </router-link>
              </div>
              <div class="subtitle" v-else-if="submitRes == 0">Verify handling!</div>
            </div>
          </div>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ref, reactive } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import {
  GetVerifyContractMetadata,
  SubmitVerifyContract,
  GetVerifyContractStatus,
} from '../../script/service/contractService';
import { getTitle } from '../../script/utils';

document.title = 'Verify & Publish Contract Source Code | The ' + getTitle + ' Explorer';

const route = useRoute();
const router = useRouter();

const contractAddress = route.query.a;
const contractName = ref('');
const contractNameRequired = ref(false);
const compilerVersion = route.query.cv;
const licenseValue = ref(parseInt(route.query.lictype as string));
const compilerType = route.query.ct;
const compilerFileName = route.query.cf;
const optimizationValue = ref(0);
const runsValue = ref(0);
const sourceCode = ref('');
const evmVersionValue = 'default';
const submitLoading = ref(false);
const submittedStatus = ref(0);
const verifyContractStatus = ref(-1);
const submittedError = ref('');
const fileList = reactive([]);
const fileRequired = ref(false);
const licenseOptions: Option[] = [];

interface Option {
  value: any;
  label: string;
}

const metadataRes = await GetVerifyContractMetadata();
metadataRes.data.licenseTypes.forEach((element) => {
  licenseOptions.push({
    value: element.id,
    label: element.name,
  } as Option);
});

// console.log(route.query);

const submit = async () => {
  if (fileRequired.value || contractNameRequired.value) {
    return;
  }

  submitLoading.value = true;
  submittedStatus.value = 0;
  contractName.value.trim() == '' ? (contractNameRequired.value = true) : (contractNameRequired.value = false);
  const formdata = new FormData();
  formdata.append('contractAddress', contractAddress as string);
  formdata.append('contractName', contractName.value);
  formdata.append('compilerType', compilerType as string);
  formdata.append('compilerVersion', compilerVersion as string);
  formdata.append('compilerFileName', compilerFileName as string);
  formdata.append('licenseType', licenseValue.value as unknown as string);
  if (fileList.length == 1) {
    formdata.append('file', fileList[0]['raw'], fileList[0]['name']);
  }
  if (compilerType == 'solidity-single-file') {
    formdata.append('sourceCode', sourceCode.value);
    formdata.append('optimization', optimizationValue.value as unknown as string);
    formdata.append('runs', runsValue.value as unknown as string);
  } else {
    fileList.length == 0 ? (fileRequired.value = true) : (fileRequired.value = false);
  }

  // console.log('submittedStatus', submittedStatus);

  const submitRes = await SubmitVerifyContract(contractAddress as string, formdata);
  // console.log('submitRes', submitRes);
  submittedStatus.value = submitRes.code;
  if (submitRes.code == 200) {
    setTimeout(() => {
      CheckVerifyContractStatus(submitRes.data.id);
    }, 3 * 1000);
  } else {
    submittedError.value = submitRes.msg;
    submitLoading.value = false;
  }
};

const CheckVerifyContractStatus = async (submitId: string) => {
  const contractStatusRes = await GetVerifyContractStatus(submitId);
  // console.log('contractStatusRes', contractStatusRes);
  verifyContractStatus.value = contractStatusRes.data.status;
  if (contractStatusRes.data.status == 1) {
    router.push(('/address/' + contractAddress) as string);
  }
  submitLoading.value = false;
};

const handleUploadChange = (fileList: any[]) => {
  if (fileList.length > 1) {
    fileList.splice(0, 1);
  }
};

const reset = () => {
  router.go(0);
};

const returnMain = () => {
  router.push('/verifyContract/input?a=' + route.query.a);
};
</script>
<style lang="less" scoped>
@import '../../css/style.css';
</style>
