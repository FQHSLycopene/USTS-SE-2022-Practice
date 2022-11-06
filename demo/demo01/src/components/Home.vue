<template>
    <div>
        <h1>Home!!!!!!!!!!!!!</h1>
        <h1>{{ msg }}</h1>
    </div>

</template>

<script>
export default {
    name: 'HelloWorld',
    props: {
        msg: String
    }
}
</script>