<template>
  <div class="siimple-form">
    <div class="siimple-form-title">Generator parameters</div>
    <div class="siimple-form-detail">Choose generator parameters.</div>
    <div class="siimple-form-field">
      <label class="siimple-label">Dictionary:</label>
      <select
        id="dict-name"
        class="siimple-select siimple-select--fluid"
        v-model="parameters.selectedDict"
        @change="onChange"
      >
        <option v-for="dict in parameters.dicts" :key="dict" :value="dict">{{ dict | capitalize }}</option>
      </select>
    </div>
    <div class="siimple-form-field">
      <label class="siimple-label">Number of words:</label>
      <select
        id="words-num"
        class="siimple-select siimple-select--fluid"
        v-model="parameters.selectedNumOfWords"
        @change="onChange"
      >
        <option v-for="(num, ndx) in parameters.numOfWords" :key="ndx" :value="num">{{ num }}</option>
      </select>
    </div>
    <div class="siimple-form-field">
      <label class="siimple-label">Number of passwords:</label>
      <select
        id="passwords-num"
        class="siimple-select siimple-select--fluid"
        v-model="parameters.selectedNumOfPasswords"
        @change="onChange"
      >
        <option v-for="(num, ndx) in parameters.numOfPasswords" :key="ndx" :value="num">{{ num }}</option>
      </select>
    </div>
  </div>
</template>

<script>
export default {
  name: "Parameters",
  props: ["parameters"],
  data() {
    return {};
  },
  methods: {
    onChange() {
      this.$emit('change')
    }
  }
};
</script>

<style>
</style>
