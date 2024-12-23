<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

type Food = {
    id: number
    name: string
}

const route = useRoute();
const router = useRouter();

const id = route.params.id;
const food = ref<Food | null>(null)

const fetchFood = async () => {
    try {
        const response = await fetch(`https://dummyjson.com/recipes/${id}`)

        if (!response.ok) {
            throw new Error("no response")
        }

        const data = await response.json()
        food.value = { id: data.id, name: data.name };
    } catch {
        router.push("/404")
    }
}

onMounted(() => {
  fetchFood();
});
</script>

<template>
    <template v-if="food">
        <h2>{{ food.name }}</h2>
        <span>{{ food.id }}</span>
    </template>
</template>
  
<style>
</style>
  