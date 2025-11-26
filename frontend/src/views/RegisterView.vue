<template>
  <div class="container">
    <h1>Register</h1>

    <form @submit.prevent="submit">
      <input v-model="email" type="email" placeholder="Email" required />
      <input v-model="password" type="password" placeholder="Password" required />
      <input v-model="confirmPassword" type="password" placeholder="Confirm Password" required />
      <button type="submit">Register</button>
    </form>

    <p>
      Sudah punya akun?
      <router-link to="/login">Login</router-link>
    </p>

    <p v-if="message" style="color:green">{{ message }}</p>
    <p v-if="error" style="color:red">{{ error }}</p>
  </div>
</template>

<script setup>
import { ref } from "vue";
import axios from "axios";

const email = ref("");
const password = ref("");
const confirmPassword = ref(""); // <-- kamu lupa ini
const error = ref("");
const message = ref("");

const submit = async () => {
  error.value = "";
  message.value = "";

  try {
    const res = await axios.post("/api/v1/register", {
      email: email.value,
      password: password.value,
      confirmPassword: confirmPassword.value,
    });

    // backend response:
    // { responseCode: "00", responseMessage: "User registered successfully" }

    if (res.data.responseCode === "00") {
      message.value = res.data.responseMessage;
    } else {
      error.value = res.data.responseMessage || "Gagal register";
    }

  } catch (err) {
    error.value = err.response?.data?.responseMessage || "Gagal register";
  }
};
</script>