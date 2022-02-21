<template>
  <div class="auth">
    <form action="" class="auth-form">
      <label for="phone">
        <MazPhoneNumberInput
          v-model="phone"
          show-code-on-list
          color="info"
          no-country-selector="true,"
          default-country-code="RU"
          @update="results = $event"
          :success="results?.isValid"
        />
        <code> {{ result }} </code>
        <button class="btn-submit" type="submit">Send</button>
      </label>
    </form>
  </div>
</template>

<script>
import "maz-ui/css/main.css";
import MazPhoneNumberInput from "maz-ui/components/MazPhoneNumberInput";
import { ref } from "vue";
export default {
  name: "Auth",
  components: {
    MazPhoneNumberInput,
  },

  data() {
    return {
      test: null,
      phone: ref,
      result: ref,
    };
  },

  // mounted() {
  //   this.$axios
  //     .get("/api/account/test")
  //     .then((response) => (this.test = response.data));
  // },
};
</script>

<style>
.auth {
  margin: 0;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}
.phone-lb {
  margin-right: 10px;
}
button {
  width: 70px;
  height: 30px;
}
</style>
