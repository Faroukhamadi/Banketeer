<script lang="ts">
	import Button from 'carbon-components-svelte/src/Button/Button.svelte';
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<h1>Index is a big yikes</h1>
<Button on:click={() => console.log('hello')}>Call thingie</Button>
