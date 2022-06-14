<script>
  import { PasswordInput, TextInput, Form, Button, Tile } from 'carbon-components-svelte'
  import 'carbon-components-svelte/css/white.css'
</script>
  
  <Tile light>
    <h3>Log in to Banketeer</h3>
    <Form on:submit>
      <TextInput style="margin-bottom: 10px;" labelText="Username" placeholder="Enter username..." required />
      <PasswordInput required type="password" labelText="Password" placeholder="Enter password..." />
      <Button type="submit" style="margin-top: 10px;">Login</Button>
    </Form>
  </Tile>
<Tile></Tile>
<Tile></Tile>
<Tile></Tile>
<Tile></Tile>
<Tile></Tile>
<Tile></Tile>
<Tile></Tile>


<style>

  h3 {
    margin-bottom: 20px;
  }

</style>