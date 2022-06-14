<script>
  import 'carbon-components-svelte/css/white.css'
  // import { Tile } from 'carbon-components-svelte'54.
  import Tile from 'carbon-components-svelte/src/Tile/Tile.svelte'
</script>

<Tile light><h4>Banketeer</h4></Tile>
<Tile></Tile>
<Tile></Tile>
<Tile></Tile>


<main>
  <slot />
</main>

<style>

  h4 {
    font-weight: 600;
    font-size: x-large;
  }
</style>