<template>
  <div class="container">
    <h1>Login</h1>

    <form @submit.prevent="submit">
      <input v-model="email" type="email" placeholder="Email" required />
      <input v-model="password" type="password" placeholder="Password" required />

      <button type="submit">Login</button>

      <p v-if="error" style="color:red">{{ error }}</p>
    </form>
  </div>
</template>

<script setup>
import { ref } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";

const router = useRouter();

const email = ref("");
const password = ref("");
const error = ref("");

const submit = async () => {
  try {
    const res = await axios.post("/api/v1/login", {
      email: email.value,
      password: password.value
    });

    if (res.data.responseCode !== "00") {
      error.value = res.data.responseMessage || "Login failed";
      return;
    }

    // SIMPAN TOKEN
    localStorage.setItem("token", res.data.token);

    // Redirect ke dashboard
    router.push("/dashboard");

  } catch (err) {
    error.value = err.response?.data?.message || "Server error!";
  }
};
</script>