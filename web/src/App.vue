<template>
  <div id="app">
    <div class="siimple-grid">
      <div class="siimple-grid-row">
        <div class="siimple-grid-col siimple-grid-col--2"></div>
        <div class="siimple-grid-col siimple-grid-col--8">
          <div class="siimple-box siimple-box--primary">
            <div class="siimple-box-title">Password generator</div>
            <div
              class="siimple-box-subtitle"
            >Generate easy to remember, but hard to guess passwords.</div>
            <div class="siimple-box-detail">
              This handy tool will generate passwords that you won't forget
              anymore.
            </div>
          </div>
        </div>
        <div class="siimple-grid-col siimple-grid-col--2"></div>
      </div>
      <div
        id="error"
        class="siimple-grid-row"
        :class="{ 'siimple--display-none': !isNetworkError }"
      >
        <div class="siimple-grid-col siimple-grid-col--2"></div>
        <div class="siimple-grid-col siimple-grid-col--8">
          <div class="siimple-alert siimple-alert--error">Network error.</div>
        </div>
        <div class="siimple-grid-col siimple-grid-col--2"></div>
      </div>
      <div
        id="toomany"
        class="siimple-grid-row"
        :class="{ 'siimple--display-none': !tooManyRequests }"
      >
        <div class="siimple-grid-col siimple-grid-col--2"></div>
        <div class="siimple-grid-col siimple-grid-col--8">
          <div class="siimple-alert siimple-alert--warning">Too many requests.</div>
        </div>
        <div class="siimple-grid-col siimple-grid-col--2"></div>
      </div>
      <div class="siimple-grid-row">
        <div class="siimple-grid-col siimple-grid-col--2"></div>
        <div class="siimple-grid-col siimple-grid-col--4">
          <div class="siimple-content">
            <Parameters :parameters="parameters" @change="onGenerate"/>
            <div class="siimple-form">
              <div class="siimple-form-field">
                <div
                  id="submit-btn"
                  class="siimple-btn siimple-btn--blue siimple-btn--fluid"
                  @click="onGenerate"
                >Generate</div>
              </div>
            </div>
          </div>
        </div>
        <div class="siimple-grid-col siimple-grid-col--4">
          <Passwords :passwords="passwords"/>
        </div>
        <div class="siimple-grid-col siimple-grid-col--2"></div>
      </div>
    </div>
    <div class="siimple-footer siimple-footer--dark siimple--text-center siimple--text-small">
      <p>
        This website does not use any cookies at all.
        <a
          href="https://github.com/sgrzywna/passphrase-web"
          class="siimple-link"
        >See by yourself.</a>
      </p>
      <p>
        Favicon courtesy of the
        <a
          href="http://www.fatcow.com/free-icons"
          class="siimple-link"
        >http://www.fatcow.com/free-icons</a>
      </p>
    </div>
  </div>
</template>

<script>
import axios from "axios";

import Parameters from "./components/Parameters.vue";
import Passwords from "./components/Passwords.vue";

export default {
  name: "app",
  components: {
    Parameters,
    Passwords
  },
  data() {
    return {
      parameters: {
        dicts: [],
        selectedDict: "",
        numOfWords: [1, 2, 3, 4, 5],
        selectedNumOfWords: 3,
        numOfPasswords: [1, 2, 3, 5, 7, 11],
        selectedNumOfPasswords: 7
      },
      passwords: [],
      isNetworkError: false,
      tooManyRequests: false
    };
  },
  created() {
    this.getDictionaries();
  },
  methods: {
    getDictionaries() {
      axios
        .get("/api/dicts.json")
        .then(response => {
          this.parameters.dicts = response.data;
          if (this.parameters.dicts.length > 0) {
            this.parameters.selectedDict = this.parameters.dicts[0];
            this.onGenerate();
          }
        })
        .catch(() => {});
    },
    onGenerate() {
      axios
        .get("/api/passwords.json", {
          params: {
            d: this.parameters.selectedDict,
            w: this.parameters.selectedNumOfWords,
            p: this.parameters.selectedNumOfPasswords
          }
        })
        .then(response => {
          this.clearErrors();
          this.passwords = response.data;
        })
        .catch(err => {
          if (err.response && err.response.status == 429) {
            this.tooManyRequests = true;
          } else {
            this.isNetworkError = true;
          }
        });
    },
    clearErrors() {
      this.isNetworkError = false;
      this.tooManyRequests = false;
    }
  }
};
</script>

<style>
@import "./assets/siimple.min.css";
</style>
