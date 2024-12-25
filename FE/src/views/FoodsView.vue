<script setup lang="ts">
import FoodCard from '@/components/FoodCard.vue';
import type { FoodCardType } from '@/components/FoodCard.vue';
import { onMounted, ref } from 'vue';
import { useRouter, RouterLink } from 'vue-router';

type Foods = {
  foods: FoodCardType[]
}

const router = useRouter();

const foods = ref<Foods| null>(null)

const fetchProducts = async () => {
  try {
    const response = await fetch("http://127.0.0.1:3000/foods")

    if (!response.ok) {
      throw new Error("no response")
    }

    const data = await response.json()
    foods.value = data
  } catch(err) {
    router.push("/404")
  }
}

onMounted(() => {
  fetchProducts();
})
</script>

<template>
  <h1 class="foods-view__headline">Foods overview</h1>
  <div class="foods-view__wrapper">
    <template v-for="food in foods?.foods" :key="food.Id">
      <FoodCard class="foods-view__card" v-bind="food" />
    </template>
  </div>
</template>
  
<style lang="scss">
  .foods-view {
    &__headline {
      font-size: 2rem;
      margin-bottom: 32px;
    }

    &__wrapper {
      display: flex;
      flex-wrap: wrap;
      gap: 24px;
    }

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