<template>
  <div class="siimple-content">
    <div class="siimple-form">
      <div class="siimple-form-title">Passwords</div>
      <div class="siimple-form-detail">List of generated passwords.</div>
      <div id="passwords" class="siimple-list">
        <div
          class="siimple-list-item"
          v-for="(password, ndx) in passwords"
          :key="ndx"
        >{{ password }}</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Passwords",
  props: ["passwords"]
};
</script>

<style>
</style>
