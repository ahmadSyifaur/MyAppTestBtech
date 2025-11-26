<template>
  <div>
    <h1>{{ message }}</h1>
    <button @click="logout">Logout</button>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";

const message = ref("");
const router = useRouter();

onMounted(async () => {
  const res = await axios.get("/api/v1/dashboard/", {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  });

  message.value = res.data.message;
});

const logout = () => {
  localStorage.removeItem("token");
  router.push("/login");
};
</script>
