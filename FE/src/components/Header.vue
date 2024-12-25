<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
    <header class="header">
        <h1 class="header__headline">Food order app</h1>
        <nav class="header__nav" aria-label="lorem ipsum">
            <ul class="header__list">
                <li class="header__list-item">
                    <RouterLink class="header__list-item-anchor" to="/">Home</RouterLink>
                </li>
                <li class="header__list-item">
                    <RouterLink class="header__list-item-anchor" to="/foods">Foods overview</RouterLink>
                </li>
            </ul>
        </nav>
    </header>
</template>

<style lang="scss">
    .header {
        max-width: 1200px;
        margin: 2rem auto;
        display: flex;
        flex-direction: column;
        align-items: center;

        &__headline {
            font-size: 2rem;
        }


        &__list {
            margin-top: 1rem;
            list-style-type: none;
            display: flex;
            gap: 1rem;
        }

        &__list-item-anchor {
            font-size: 1.4rem;
            color: black;
            font-size: 1.2rem;
        }
    }
</style>