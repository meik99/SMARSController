<template>
<span class="form-group">
  <button class="checkbox"
          :class="classes"
          @click="toggle">{{ name }}</button>
</span>
</template>

<script lang="ts">
import {Options, Vue} from "vue-class-component";

@Options({
  props: {
    id: String,
    name: String,
    checked: {
      type: Boolean,
      required: false,
      default: false
    }
  },
  emits: [
    'checkbox-toggle',
  ],
  computed: {
    classes() {
      return {
        'checked': this.isChecked
      }
    }
  },
  watch: {
    checked(val) {
      this.isChecked = val;
    }
  },
  data() {
    return {
      isChecked: false
    }
  },
  created() {
    this.isChecked = this.checked;
  },
  methods: {
    toggle() {
      this.isChecked = !this.isChecked;
      this.$emit('checkbox-toggle', this.isChecked);
    }
  }
})
export default class Checkbox extends Vue {
}
</script>

<style scoped>

</style>