<template>
  <div class="container mt-10p border">
    <div class="row">
      <div class="col-12">
        <h1>Alarm</h1>
        <hr>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <Alarm></Alarm>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import Alarm from "@/components/Alarm.vue";
import {Authentication} from "@/authentication/Authentication";

@Options({
  components: {
    Alarm,
  },
  data() {
    return {
      alarm: null
    }
  },
  created() {
    new Authentication().checkAuthStatusForApp(this);
  }
})
export default class Home extends Vue {}
</script>

<style>
.mt-10p {
  margin-top: 10%;
}
.border {
  border: 1px solid #eee;
  border-radius: 16px;
}
</style>