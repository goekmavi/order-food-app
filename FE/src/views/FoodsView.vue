<script setup lang="ts">
import FoodCard from '@/components/FoodCard.vue';
import type { FoodCardType } from '@/components/FoodCard.vue';
import { onMounted, ref } from 'vue';
import { useRouter, RouterLink } from 'vue-router';

type Foods = {
  recipes: FoodCardType[]
}

const router = useRouter();

const foods = ref<Foods| null>(null)

const fetchProducts = async () => {
  try {
    const response = await fetch("https://dummyjson.com/recipes")

    if (!response.ok) {
      throw new Error("no response")
    }

    const data = await response.json()
    foods.value = {
      recipes: data.recipes.map((recipe: any) => ({
        id: recipe.id,
        headline: recipe.name,
        image: recipe.image,
        rating: recipe.rating,
        anchor: "/foods/" + recipe.id
      })),
    };
  } catch {
    router.push("/404")
  }
}

onMounted(() => {
  fetchProducts();
})
</script>

<template>
  <h1>Foods overview</h1>
  <div class="foods-view">
    <template v-for="food in foods?.recipes" :key="food.id">
      <FoodCard class="foods-view__card" v-bind="food" />
    </template>
  </div>
</template>
  
<style lang="scss">
  .foods-view {
    display: flex;
    flex-wrap: wrap;
    gap: 24px;

    &__card {
      flex-basis: calc(100% / 2 - 12px);
    }

    @media (min-width: 680px) {
      &__card {
        flex-basis: calc(100% / 3 - 16px);
      }
    }
  }
</style>