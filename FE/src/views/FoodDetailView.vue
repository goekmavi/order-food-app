<script setup lang="ts">
import FoodCard, { type FoodCardType } from '@/components/FoodCard.vue';
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();

const id = route.params.id;
const food = ref<FoodCardType | null>(null)

const fetchFood = async () => {
    try {
        const response = await fetch(`https://dummyjson.com/recipes/${id}`)

        if (!response.ok) {
            throw new Error("no response")
        }

        const data = await response.json()

        food.value = {
            Headline: data.name,
            Image: data.image,
            Id: data.id,
            Rating: data.rating
        };
    } catch {
        router.push("/404")
    }
}

onMounted(() => {
  fetchFood();
});
</script>

<template>
    <div class="food-detail-view" v-if="food">
        <FoodCard class="food-detail-view__food-card" v-bind="food" />
        <button class="food-detail-view__button">Add to cart</button>
    </div>
</template>
  
<style lang="scss">
    .food-detail-view {
        display: flex;
        gap: 32px;
        align-items: flex-start;

        &__food-card {
            flex-basis: 50%;
        }
    }
</style>
  