<script setup lang="ts">
import { RouterLink } from 'vue-router';

export type FoodCardType = {
    headline: string,
    image: string,
    id: number,
    rating: number,
    anchor: string
}

defineProps<FoodCardType>();
</script>

<template>
    <RouterLink class="food-card" :to="anchor">
        <article class="food-card__wrapper">
            <img class="food-card__image" :src="image" :alt="headline" />
            <h3 class="food-card__headline">{{ headline }}</h3>
            <span class="food-card__rating">{{ rating }}</span>
        </article>
    </RouterLink>
</template>

<style lang="scss">
    .food-card {
        box-sizing: border-box;
        box-shadow: 0px 0px 18px -7px rgba(0,0,0,0.75);

        &__wrapper {
            padding: 26px;
        }
        &__image {
            width: 100%;
            height: auto;
        }
    }
</style>