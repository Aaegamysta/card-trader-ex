<template>
    <div>
        <header class="header">
            <div>
                <img class="header__logo" src="/logos/logo.png" alt="">
            </div>
            <h1 class="header__title">Card Trader EX</h1>
            <nav class="header__menu">
                <a class="header__link">Play</a>
                <a class="header__link">Hiscores</a>
                <a class="header__link">Login</a>
                <a class="header__link">About</a>
            </nav>
        </header>
    </div>
</template>

<script setup lang="ts">

</script>

<style scoped>
.header {
    border: 1px solid red;
    width: 90%;
    height: 4rem;
    margin: 0 auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 0;
}

.header__logo {
    height: 5.5em;
    width: auto;
}

.header__title {
    flex: 1;
    font-size: 1.5rem;
    text-transform: uppercase;
}

.header__menu {
    justify-content: flex-end;
    flex: 1;
    display: flex;
    gap: 1rem;
}

.header__link {
    text-decoration: none;
    color: #000;
}
</style>