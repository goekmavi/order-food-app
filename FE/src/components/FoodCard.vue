<script setup lang="ts">
import { RouterLink } from 'vue-router';

export type FoodCardType = {
    Headline: string,
    Image: string,
    Id: number,
    Rating: number,
    Anchor?: string
}

defineProps<FoodCardType>();
</script>

<template>
    <component :is="Anchor ? 'RouterLink' : 'div'" class="food-card" :to="Anchor">
        <article class="food-card__wrapper">
            <img class="food-card__image" :src="Image" :alt="Headline" />
            <h3 class="food-card__headline">{{ Headline }}</h3>
            <span class="food-card__rating">{{ Rating }}</span>
        </article>
    </component>
</template>

<style lang="scss">
    .food-card {
        box-sizing: border-box;
        box-shadow: 0px 0px 18px -7px rgba(0,0,0,0.75);
        text-decoration: none;
        color: black;

        &__wrapper {
            padding: 26px;
        }
        &__image {
            width: 100%;
            height: auto;
        }

        &__headline {
            font-size: 1.3rem;
            margin: 8px 0;
        }
    }
</style>