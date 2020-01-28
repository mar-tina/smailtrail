<script>
  import { link } from "svelte-spa-router";
</script>

<style>
  .container {
    display: flex;
    font-size: 20px;
    justify-content: space-evenly;
    margin: 20px;
    padding: 20px;
    flex-wrap: wrap;
  }
</style>

<div class="container">
  <div>
    <a href="/" use:link>
      <b>Home</b>
    </a>
  </div>

  <div>
    <a href="/msgs" use:link>
      <b>Messages</b>
    </a>
  </div>

  <div>
    <a href="/un-sub" use:link>
      <b>Un-sub</b>
    </a>
  </div>

  <div>
    <a href="/auth" use:link>
      <b>Auth</b>
    </a>
  </div>

</div>
