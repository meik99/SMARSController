<template>
  <div class="row">
    <div class="col-6">
      <div class="form-group">
        <label for="hours">Hour</label>
        <input type="number" id="hours" v-model="hour">
      </div>
    </div>
    <div class="col-6">
      <div class="form-group">
        <label for="minutes">Minute</label>
        <input type="number" id="minutes" v-model="minute">
      </div>
    </div>
  </div>
  <div class="row">
    <div class="col-12 center">
      <Checkbox v-for="day in days"
                :key="day.id"
                :id="day.id"
                :name="day.name"
                :checked="day.checked"
                @checkbox-toggle="onCheckboxToggled(day)">
      </Checkbox>
    </div>
  </div>
  <div class="row">
    <div class="col-2">
      <button @click="setAlarm">Set</button>
    </div>
    <div class="col-2">
      <button class="danger" @click="deleteAlarm">Delete</button>
    </div>
  </div>
</template>

<script lang="ts">
import {Options, Vue} from "vue-class-component";
import Checkbox from '@/components/Checkbox.vue';
import axios from 'axios';
import {Authentication} from "@/authentication/Authentication";

@Options({
  components: {
    Checkbox,
  },
  watch: {
    hour() {
      if (this.hour < 0) {
        this.hour = 0;
      } else if (this.hour > 23) {
        this.hour = 23;
      }
    },
    minute() {
      if (this.minute < 0) {
        this.minute = 0;
      } else if (this.minute > 59) {
        this.minute = 59;
      }
    }
  },
  data() {
    return {
      id: null,
      minute: 0,
      hour: 0,
      days: [
        {id: 'mon', name: 'Monday', checked: true},
        {id: 'tue', name: 'Tuesday', checked: true},
        {id: 'wed', name: 'Wednesday', checked: true},
        {id: 'thu', name: 'Thursday', checked: true},
        {id: 'fri', name: 'Friday', checked: true},
        {id: 'sat', name: 'Saturday', checked: true},
        {id: 'sun', name: 'Sunday', checked: true},
      ]
    };
  },
  created() {
    const authentication = new Authentication();
    authentication.checkAuthStatusForApp(this);

    const code = authentication.getAuthCode();
    console.log(process.env);

    axios.get(process.env.VUE_APP_COFFEE_ALARM_API, {
      headers: {
        Authorization: `Bearer ${code}`
      }
    })
        .then(result => {
          if (result.data && result.data.length > 0) {
            const alarm = result.data[0];
            if (alarm) {
              this.id = alarm._id;
              this.hour = alarm.hour;
              this.minute = alarm.minute;
              for(let i = 0; i < this.days.length; i++) {
                this.days[i].checked = false;
              }
              for(const day of alarm.days) {
                this.days.find((d: any) => d.id === day).checked = true;
              }
            }
          }
        })
        .catch(err => {
          if (err && err.response && err.response.status && err.response.status === 401) {
            authentication.clear();
            authentication.checkAuthStatusForApp(this);
          } else {
            console.log(err.response);
          }
        })
  },
  methods: {
    setAlarm() {
      const authentication = new Authentication();
      authentication.checkAuthStatusForApp(this);

      const code = authentication.getAuthCode();
      const data: any = {
        hour: this.hour,
        minute: this.minute,
        days: []
      };
      const config = {
        headers: {
          Authorization: `Bearer ${code}`
        }
      };

      for(const day of this.days) {
        if(day.checked) {
          data.days.push(day.id);
        }
      }

      if (!this.id) {
        axios.post(process.env.VUE_APP_COFFEE_ALARM_API, data, config)
      } else {
        data._id = this.id;
        axios.put(process.env.VUE_APP_COFFEE_ALARM_API, data, config)
      }
    },
    deleteAlarm(){
      const authentication = new Authentication();
      authentication.checkAuthStatusForApp(this);

      const code = authentication.getAuthCode();
      const config = {
        headers: {
          Authorization: `Bearer ${code}`
        }
      };

      if (this.id) {
        axios.delete(`${process.env.VUE_APP_COFFEE_ALARM_API}/${this.id}`, config)
        .then(result => console.log(result))
        .catch(err => console.log(err));
      }
    },
    onCheckboxToggled(day: any) {
      console.log(day);
      const dayIndex = this.days.indexOf(this.days.find((d: any) => d.id === day.id));
      console.log(dayIndex);
      if (dayIndex >= 0) {
        this.days[dayIndex].checked = !this.days[dayIndex].checked;
      }
    }
  }
})
export default class Alarm extends Vue {
}
</script>

<style scoped>

</style>