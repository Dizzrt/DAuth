<template>
  <a-layout style="width: 100%; height: 100%">
    <!-- <a-layout-header style="height: 64px">
    
    </a-layout-header> -->

    <a-layout-content style="display: flex; justify-content: center; align-items: center">
      <div class="login-card">
        <div style="width: 300px; border-right: 1px solid var(--color-neutral-3); box-sizing: border-box">
          <a-image
            width="100%"
            height="100%"
            fit="cover"
            :src="'https://pic.bbtkids.cn/Foc1XGREfk2YaEFJQ9qsofb_agLN'" />
        </div>

        <div style="flex: 1; display: flex; flex-direction: column">
          <div style="display: flex; justify-content: center; align-items: center">
            <a-form :model="formData" :style="{ width: '500px' }" size="large" :layout="'vertical'">
              <a-form-item field="account" label="Account">
                <a-input v-model="formData.account" />
              </a-form-item>

              <a-form-item field="password" label="Password">
                <a-input v-model="formData.password" type="password" />
              </a-form-item>

              <a-form-item>
                <a-checkbox v-model="formData.isRemeber">remeber me</a-checkbox>
              </a-form-item>

              <a-form-item>
                <a-button type="primary" @click="login">Submit</a-button>
              </a-form-item>
            </a-form>
          </div>
        </div>
      </div>
    </a-layout-content>

    <!-- <a-layout-footer style="height: 20px">
    
    </a-layout-footer> -->
  </a-layout>
</template>

<style scoped>
.login-card {
  display: flex;
  width: 840px;
  height: 520px;
  background-color: #fff;
  border: 1px solid var(--color-neutral-3);
  border-radius: 5px;
  box-shadow: 0 0 10px rgb(0 0 0 / 10%);
}

.login-card div {
  height: 100%;
}
</style>

<script setup lang="ts">
import { reactive } from 'vue';

const formData = reactive({
  account: '',
  password: '',
  isRemeber: false,
});

const login = async () => {
  console.log(formData);
};
</script>
